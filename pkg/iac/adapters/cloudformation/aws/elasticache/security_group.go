package elasticache

import (
	"vulnscan/pkg/iac/providers/aws/elasticache"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

func getSecurityGroups(ctx parser.FileContext) (securityGroups []elasticache.SecurityGroup) {

	sgResources := ctx.GetResourcesByType("AWS::ElastiCache::SecurityGroup")

	for _, r := range sgResources {

		sg := elasticache.SecurityGroup{
			Metadata:    r.Metadata(),
			Description: r.GetStringProperty("Description"),
		}
		securityGroups = append(securityGroups, sg)
	}

	return securityGroups
}
