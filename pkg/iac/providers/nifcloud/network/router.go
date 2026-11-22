package network

import (
	iacTypes "vulnscan/pkg/iac/types"
)

type Router struct {
	Metadata          iacTypes.Metadata
	SecurityGroup     iacTypes.StringValue
	NetworkInterfaces []NetworkInterface
}
