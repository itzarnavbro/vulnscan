package network

import (
	iacTypes "vulnscan/pkg/iac/types"
)

type VpnGateway struct {
	Metadata      iacTypes.Metadata
	SecurityGroup iacTypes.StringValue
}
