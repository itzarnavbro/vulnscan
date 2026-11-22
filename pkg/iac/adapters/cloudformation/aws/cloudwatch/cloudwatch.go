package cloudwatch

import (
	"vulnscan/pkg/iac/providers/aws/cloudwatch"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a Cloudwatch instance
func Adapt(cfFile parser.FileContext) cloudwatch.CloudWatch {
	return cloudwatch.CloudWatch{
		LogGroups: getLogGroups(cfFile),
	}
}
