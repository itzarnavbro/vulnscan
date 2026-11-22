package compute

import (
	iacTypes "vulnscan/pkg/iac/types"
)

type SubNetwork struct {
	Metadata              iacTypes.Metadata
	Name                  iacTypes.StringValue
	Purpose               iacTypes.StringValue
	EnableFlowLogs        iacTypes.BoolValue
	PrivateIPGoogleAccess iacTypes.BoolValue
}
