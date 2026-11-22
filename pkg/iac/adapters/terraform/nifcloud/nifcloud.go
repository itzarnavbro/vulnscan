package nifcloud

import (
	"vulnscan/pkg/iac/adapters/terraform/nifcloud/computing"
	"vulnscan/pkg/iac/adapters/terraform/nifcloud/dns"
	"vulnscan/pkg/iac/adapters/terraform/nifcloud/nas"
	"vulnscan/pkg/iac/adapters/terraform/nifcloud/network"
	"vulnscan/pkg/iac/adapters/terraform/nifcloud/rdb"
	"vulnscan/pkg/iac/adapters/terraform/nifcloud/sslcertificate"
	"vulnscan/pkg/iac/providers/nifcloud"
	"vulnscan/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) nifcloud.Nifcloud {
	return nifcloud.Nifcloud{
		Computing:      computing.Adapt(modules),
		DNS:            dns.Adapt(modules),
		NAS:            nas.Adapt(modules),
		Network:        network.Adapt(modules),
		RDB:            rdb.Adapt(modules),
		SSLCertificate: sslcertificate.Adapt(modules),
	}
}
