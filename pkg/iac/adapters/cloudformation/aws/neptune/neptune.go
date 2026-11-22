package neptune

import (
	"vulnscan/pkg/iac/providers/aws/neptune"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a Neptune instance
func Adapt(cfFile parser.FileContext) neptune.Neptune {
	return neptune.Neptune{
		Clusters: getClusters(cfFile),
	}
}
