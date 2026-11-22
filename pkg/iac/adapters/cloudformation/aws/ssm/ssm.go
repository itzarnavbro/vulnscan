package ssm

import (
	"vulnscan/pkg/iac/providers/aws/ssm"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an SSM instance
func Adapt(cfFile parser.FileContext) ssm.SSM {
	return ssm.SSM{
		Secrets: getSecrets(cfFile),
	}
}
