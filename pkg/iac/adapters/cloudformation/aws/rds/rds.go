package rds

import (
	"vulnscan/pkg/iac/providers/aws/rds"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an RDS instance
func Adapt(cfFile parser.FileContext) rds.RDS {
	clusters, orphans := getClustersAndInstances(cfFile)
	return rds.RDS{
		Instances:       orphans,
		Clusters:        clusters,
		Classic:         getClassic(cfFile),
		ParameterGroups: getParameterGroups(cfFile),
		Snapshots:       nil,
	}
}
