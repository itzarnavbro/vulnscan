package ecr

import (
	"vulnscan/pkg/iac/providers/aws/ecr"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an ECR instance
func Adapt(cfFile parser.FileContext) ecr.ECR {
	return ecr.ECR{
		Repositories: getRepositories(cfFile),
	}
}
