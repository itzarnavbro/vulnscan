package ec2

import iacTypes "vulnscan/pkg/iac/types"

type RequestedAMI struct {
	Metadata iacTypes.Metadata
	Owners   iacTypes.StringValueList
}
