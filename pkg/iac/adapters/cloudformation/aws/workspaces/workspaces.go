package workspaces

import (
	"vulnscan/pkg/iac/providers/aws/workspaces"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a Workspaces instance
func Adapt(cfFile parser.FileContext) workspaces.WorkSpaces {
	return workspaces.WorkSpaces{
		WorkSpaces: getWorkSpaces(cfFile),
	}
}
