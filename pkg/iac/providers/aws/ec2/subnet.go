package ec2

import (
	iacTypes "vulnscan/pkg/iac/types"
)

type Subnet struct {
	Metadata            iacTypes.Metadata
	MapPublicIpOnLaunch iacTypes.BoolValue
}
