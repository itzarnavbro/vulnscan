package elasticache

import (
	"vulnscan/pkg/iac/providers/aws/elasticache"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an ElasticCache instance
func Adapt(cfFile parser.FileContext) elasticache.ElastiCache {
	return elasticache.ElastiCache{
		Clusters:          getClusterGroups(cfFile),
		ReplicationGroups: getReplicationGroups(cfFile),
		SecurityGroups:    getSecurityGroups(cfFile),
	}
}
