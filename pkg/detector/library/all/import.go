// Package all registers all built-in library vulnerability suppliers.
// Import it for side effects to enable supplier-specific advisory detection
// (e.g., Seal Security, Echo):
//
//	import _ "vulnscan/pkg/detector/library/all"
package all

import (
	_ "vulnscan/pkg/detector/library/echo"
	_ "vulnscan/pkg/detector/library/seal"
)
