package sslcertificate

import (
	"vulnscan/pkg/iac/providers/nifcloud/sslcertificate"
	"vulnscan/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) sslcertificate.SSLCertificate {
	return sslcertificate.SSLCertificate{
		ServerCertificates: adaptServerCertificates(modules),
	}
}
