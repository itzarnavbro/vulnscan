package synapse

import (
	iacTypes "vulnscan/pkg/iac/types"
)

type Synapse struct {
	Workspaces []Workspace
}

type Workspace struct {
	Metadata                    iacTypes.Metadata
	EnableManagedVirtualNetwork iacTypes.BoolValue
}
