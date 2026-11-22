package k8s

import (
	"context"

	"github.com/aquasecurity/trivy-db/pkg/db"
	"vulnscan/pkg/cache"
	"vulnscan/pkg/fanal/applier"
	ftypes "vulnscan/pkg/fanal/types"
	"vulnscan/pkg/scan/langpkg"
	"vulnscan/pkg/scan/local"
	"vulnscan/pkg/scan/ospkg"
	"vulnscan/pkg/types"
	"vulnscan/pkg/vulnerability"
)

// ScanKubernetes implements the scanner
type ScanKubernetes struct {
	localScanner local.Service
}

// NewScanKubernetes is the factory method for scanner
func NewScanKubernetes(s local.Service) *ScanKubernetes {
	return &ScanKubernetes{localScanner: s}
}

// NewKubernetesScanner is the factory method for scanner
func NewKubernetesScanner() *ScanKubernetes {
	return initializeScanK8s(nil)
}

// initializeScanK8s creates a new Kubernetes scanner with the provided cache.
// If cache is nil, it will create the scanner without cache dependency.
func initializeScanK8s(localArtifactCache cache.LocalArtifactCache) *ScanKubernetes {
	applier := applier.NewApplier(localArtifactCache)
	osScanner := ospkg.NewScanner()
	langScanner := langpkg.NewScanner()
	vulnClient := vulnerability.NewClient(db.Config{})

	localService := local.NewService(applier, osScanner, langScanner, vulnClient)
	return NewScanKubernetes(localService)
}

// Scan scans k8s core components and return it findings
func (sk ScanKubernetes) Scan(ctx context.Context, target types.ScanTarget, options types.ScanOptions) (types.Results, ftypes.OS, error) {
	return sk.localScanner.ScanTarget(ctx, target, options)
}
