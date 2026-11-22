package nas

import (
	iacTypes "vulnscan/pkg/iac/types"
)

type NASInstance struct {
	Metadata  iacTypes.Metadata
	NetworkID iacTypes.StringValue
}
