package eks

import (
	"vulnscan/pkg/iac/providers/aws/eks"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an EKS instance
func Adapt(cfFile parser.FileContext) eks.EKS {
	return eks.EKS{
		Clusters: getClusters(cfFile),
	}
}
