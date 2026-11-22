package cloudformation

import (
	"vulnscan/pkg/iac/adapters/cloudformation/aws"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
	"vulnscan/pkg/iac/state"
)

// Adapt adapts the Cloudformation instance
func Adapt(cfFile parser.FileContext) *state.State {
	return &state.State{
		AWS: aws.Adapt(cfFile),
	}
}
