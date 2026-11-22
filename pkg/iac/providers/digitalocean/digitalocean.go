package digitalocean

import (
	"vulnscan/pkg/iac/providers/digitalocean/compute"
	"vulnscan/pkg/iac/providers/digitalocean/spaces"
)

type DigitalOcean struct {
	Compute compute.Compute
	Spaces  spaces.Spaces
}
