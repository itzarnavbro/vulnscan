package redshift

import (
	"vulnscan/pkg/iac/providers/aws/redshift"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a RedShift instance
func Adapt(cfFile parser.FileContext) redshift.Redshift {
	return redshift.Redshift{
		Clusters:          getClusters(cfFile),
		SecurityGroups:    getSecurityGroups(cfFile),
		ClusterParameters: getParameters(cfFile),
		ReservedNodes:     nil,
	}
}
