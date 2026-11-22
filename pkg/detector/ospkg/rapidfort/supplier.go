package rapidfort

import (
	"vulnscan/pkg/detector/ospkg/driver"
	ftypes "vulnscan/pkg/fanal/types"
)

// Supplier creates a RapidFort driver for curated images, which the RapidFort
// fanal analyzer marks by setting OS.Supplier.
func Supplier(os ftypes.OS, _ []ftypes.Package) driver.Driver {
	if os.Supplier != ftypes.SupplierRapidFort {
		return nil
	}
	switch os.Family {
	case ftypes.Ubuntu, ftypes.Alpine, ftypes.Debian, ftypes.RedHat, ftypes.Oracle, ftypes.Rocky, ftypes.Alma, ftypes.Amazon:
		return NewScanner(os.Family)
	}
	return nil
}
