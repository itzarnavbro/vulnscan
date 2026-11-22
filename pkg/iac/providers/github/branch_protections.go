package github

import (
	iacTypes "vulnscan/pkg/iac/types"
)

type BranchProtection struct {
	Metadata             iacTypes.Metadata
	RequireSignedCommits iacTypes.BoolValue
}
