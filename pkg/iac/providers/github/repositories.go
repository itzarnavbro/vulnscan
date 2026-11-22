package github

import (
	iacTypes "vulnscan/pkg/iac/types"
)

type Repository struct {
	Metadata            iacTypes.Metadata
	Public              iacTypes.BoolValue
	VulnerabilityAlerts iacTypes.BoolValue
	Archived            iacTypes.BoolValue
}
