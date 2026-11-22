package dns

import (
	"vulnscan/pkg/iac/providers/nifcloud/dns"
	"vulnscan/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) dns.DNS {
	return dns.DNS{
		Records: adaptRecords(modules),
	}
}
