package nifcloud

import (
	"vulnscan/pkg/iac/providers/nifcloud/computing"
	"vulnscan/pkg/iac/providers/nifcloud/dns"
	"vulnscan/pkg/iac/providers/nifcloud/nas"
	"vulnscan/pkg/iac/providers/nifcloud/network"
	"vulnscan/pkg/iac/providers/nifcloud/rdb"
	"vulnscan/pkg/iac/providers/nifcloud/sslcertificate"
)

type Nifcloud struct {
	Computing      computing.Computing
	DNS            dns.DNS
	NAS            nas.NAS
	Network        network.Network
	RDB            rdb.RDB
	SSLCertificate sslcertificate.SSLCertificate
}
