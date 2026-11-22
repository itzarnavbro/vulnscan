package compute

import (
	iacTypes "vulnscan/pkg/iac/types"
)

type Disk struct {
	Metadata   iacTypes.Metadata
	Name       iacTypes.StringValue
	Encryption DiskEncryption
}

type DiskEncryption struct {
	Metadata   iacTypes.Metadata
	RawKey     iacTypes.BytesValue
	KMSKeyLink iacTypes.StringValue
}
