package sslcertificate

import (
	iacTypes "vulnscan/pkg/iac/types"
)

type ServerCertificate struct {
	Metadata   iacTypes.Metadata
	Expiration iacTypes.TimeValue
}
