package cloudstack

import (
	"vulnscan/pkg/iac/adapters/terraform/cloudstack/compute"
	"vulnscan/pkg/iac/providers/cloudstack"
	"vulnscan/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) cloudstack.CloudStack {
	return cloudstack.CloudStack{
		Compute: compute.Adapt(modules),
	}
}
