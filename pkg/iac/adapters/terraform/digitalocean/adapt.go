package digitalocean

import (
	"vulnscan/pkg/iac/adapters/terraform/digitalocean/compute"
	"vulnscan/pkg/iac/adapters/terraform/digitalocean/spaces"
	"vulnscan/pkg/iac/providers/digitalocean"
	"vulnscan/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) digitalocean.DigitalOcean {
	return digitalocean.DigitalOcean{
		Compute: compute.Adapt(modules),
		Spaces:  spaces.Adapt(modules),
	}
}
