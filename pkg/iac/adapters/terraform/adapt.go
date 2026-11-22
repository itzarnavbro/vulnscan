package terraform

import (
	"vulnscan/pkg/iac/adapters/terraform/aws"
	"vulnscan/pkg/iac/adapters/terraform/azure"
	"vulnscan/pkg/iac/adapters/terraform/cloudstack"
	"vulnscan/pkg/iac/adapters/terraform/digitalocean"
	"vulnscan/pkg/iac/adapters/terraform/github"
	"vulnscan/pkg/iac/adapters/terraform/google"
	"vulnscan/pkg/iac/adapters/terraform/kubernetes"
	"vulnscan/pkg/iac/adapters/terraform/nifcloud"
	"vulnscan/pkg/iac/adapters/terraform/openstack"
	"vulnscan/pkg/iac/adapters/terraform/oracle"
	"vulnscan/pkg/iac/state"
	"vulnscan/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) *state.State {
	return &state.State{
		AWS:          aws.Adapt(modules),
		Azure:        azure.Adapt(modules),
		CloudStack:   cloudstack.Adapt(modules),
		DigitalOcean: digitalocean.Adapt(modules),
		GitHub:       github.Adapt(modules),
		Google:       google.Adapt(modules),
		Kubernetes:   kubernetes.Adapt(modules),
		Nifcloud:     nifcloud.Adapt(modules),
		OpenStack:    openstack.Adapt(modules),
		Oracle:       oracle.Adapt(modules),
	}
}
