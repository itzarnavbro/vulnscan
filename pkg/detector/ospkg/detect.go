package ospkg

import (
	"context"
	"maps"
	"slices"

	"github.com/samber/lo"
	"golang.org/x/xerrors"

	"vulnscan/pkg/detector/ospkg/alma"
	"vulnscan/pkg/detector/ospkg/alpine"
	"vulnscan/pkg/detector/ospkg/amazon"
	"vulnscan/pkg/detector/ospkg/azure"
	"vulnscan/pkg/detector/ospkg/bottlerocket"
	"vulnscan/pkg/detector/ospkg/chainguard"
	"vulnscan/pkg/detector/ospkg/coreos"
	"vulnscan/pkg/detector/ospkg/debian"
	"vulnscan/pkg/detector/ospkg/driver"
	"vulnscan/pkg/detector/ospkg/echo"
	"vulnscan/pkg/detector/ospkg/minimos"
	"vulnscan/pkg/detector/ospkg/oracle"
	"vulnscan/pkg/detector/ospkg/photon"
	"vulnscan/pkg/detector/ospkg/rapidfort"
	"vulnscan/pkg/detector/ospkg/redhat"
	"vulnscan/pkg/detector/ospkg/rocky"
	"vulnscan/pkg/detector/ospkg/rootio"
	"vulnscan/pkg/detector/ospkg/seal"
	"vulnscan/pkg/detector/ospkg/suse"
	"vulnscan/pkg/detector/ospkg/ubuntu"
	"vulnscan/pkg/detector/ospkg/wolfi"
	ftypes "vulnscan/pkg/fanal/types"
	"vulnscan/pkg/log"
	"vulnscan/pkg/types"
)

// Detector detects OS package vulnerabilities.
type Detector struct {
	target types.ScanTarget
	driver driver.Driver
}

var (
	// ErrUnsupportedOS defines error for unsupported OS
	ErrUnsupportedOS = xerrors.New("unsupported os")

	drivers = map[ftypes.OSType]driver.Driver{
		ftypes.Alpine:             alpine.NewScanner(),
		ftypes.Alma:               alma.NewScanner(),
		ftypes.Amazon:             amazon.NewScanner(),
		ftypes.Azure:              azure.NewAzureScanner(),
		ftypes.Bottlerocket:       bottlerocket.NewScanner(),
		ftypes.CBLMariner:         azure.NewMarinerScanner(),
		ftypes.Debian:             debian.NewScanner(),
		ftypes.Ubuntu:             ubuntu.NewScanner(),
		ftypes.RedHat:             redhat.NewScanner(),
		ftypes.CentOS:             redhat.NewScanner(),
		ftypes.Rocky:              rocky.NewScanner(),
		ftypes.Oracle:             oracle.NewScanner(),
		ftypes.OpenSUSETumbleweed: suse.NewScanner(suse.OpenSUSETumbleweed),
		ftypes.OpenSUSELeap:       suse.NewScanner(suse.OpenSUSE),
		ftypes.SLES:               suse.NewScanner(suse.SUSEEnterpriseLinux),
		ftypes.SLEMicro:           suse.NewScanner(suse.SUSEEnterpriseLinuxMicro),
		ftypes.Photon:             photon.NewScanner(),
		ftypes.Wolfi:              wolfi.NewScanner(),
		ftypes.Chainguard:         chainguard.NewScanner(),
		ftypes.Echo:               echo.NewScanner(),
		ftypes.MinimOS:            minimos.NewScanner(),
		ftypes.CoreOS:             coreos.NewScanner(),
	}

	// suppliers dynamically generate drivers based on package information
	// and environment detection. They are tried before standard OS-specific drivers.
	suppliers = []driver.Supplier{
		rapidfort.Supplier,
		rootio.Supplier,
		seal.Supplier,
	}
)

// resolver holds the candidate drivers and suppliers and resolves one for a scan target.
type resolver struct {
	drivers   map[ftypes.OSType]driver.Driver
	suppliers []driver.Supplier
}

// Option configures a Detector. Options are provided for extensibility by users
// of Trivy as a library and are not used within Trivy itself.
type Option func(*resolver)

// WithDriver registers a driver for the given OS family, overriding the default one.
func WithDriver(family ftypes.OSType, drv driver.Driver) Option {
	return func(r *resolver) {
		r.drivers[family] = drv
	}
}

// WithSupplier registers an additional supplier. It takes priority over the
// built-in suppliers and the standard OS-specific drivers. When called multiple
// times, the most recently registered supplier is tried first.
func WithSupplier(supplier driver.Supplier) Option {
	return func(r *resolver) {
		r.suppliers = slices.Insert(r.suppliers, 0, supplier)
	}
}

// NewDetector creates a new Detector for the given scan target
func NewDetector(target types.ScanTarget, opts ...Option) (*Detector, error) {
	r := &resolver{
		drivers:   maps.Clone(drivers),
		suppliers: slices.Clone(suppliers),
	}
	for _, opt := range opts {
		opt(r)
	}

	drv, err := r.resolve(target.OS, target.Packages)
	if err != nil {
		return nil, err
	}
	return &Detector{
		target: target,
		driver: drv,
	}, nil
}

// Detect detects the vulnerabilities
func (d *Detector) Detect(ctx context.Context) ([]types.DetectedVulnerability, bool, error) {
	ctx = log.WithContextPrefix(ctx, string(d.target.OS.Family))

	eosl := !d.driver.IsSupportedVersion(ctx, d.target.OS.Family, d.target.OS.Name)

	// gpg-pubkey does not carry a real version.
	// Matching it against any advisory is meaningless, for every driver.
	pkgs := lo.Filter(d.target.Packages, func(pkg ftypes.Package, _ int) bool {
		return pkg.Name != "gpg-pubkey"
	})

	// By default, drop packages installed from third-party repositories such as EPEL or
	// Docker: an OS vendor's advisories do not describe them. A driver whose own feed
	// covers those packages (Echo, Seal, RapidFort) implements PackageFilter to keep them instead.
	filterFunc := driver.DropThirdPartyPackages
	if f, ok := d.driver.(driver.PackageFilter); ok {
		filterFunc = f.FilterPackages
	}

	vulns, err := d.driver.Detect(ctx, d.target.OS.Name, d.target.Repository, filterFunc(ctx, pkgs))
	if err != nil {
		return nil, false, xerrors.Errorf("failed detection: %w", err)
	}

	return vulns, eosl, nil
}

func (r *resolver) resolve(os ftypes.OS, pkgs []ftypes.Package) (driver.Driver, error) {
	// Try suppliers first
	for _, supplier := range r.suppliers {
		if d := supplier(os, pkgs); d != nil {
			return d, nil
		}
	}

	// Fall back to standard drivers
	if d, ok := r.drivers[os.Family]; ok {
		return d, nil
	}

	log.Warn("Unsupported os", log.String("family", string(os.Family)))
	return nil, ErrUnsupportedOS
}
