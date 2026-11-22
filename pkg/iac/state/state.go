package state

import (
	"reflect"

	"vulnscan/pkg/iac/providers/aws"
	"vulnscan/pkg/iac/providers/azure"
	"vulnscan/pkg/iac/providers/cloudstack"
	"vulnscan/pkg/iac/providers/digitalocean"
	"vulnscan/pkg/iac/providers/github"
	"vulnscan/pkg/iac/providers/google"
	"vulnscan/pkg/iac/providers/kubernetes"
	"vulnscan/pkg/iac/providers/nifcloud"
	"vulnscan/pkg/iac/providers/openstack"
	"vulnscan/pkg/iac/providers/oracle"
	"vulnscan/pkg/iac/rego/convert"
)

type State struct {
	AWS          aws.AWS
	Azure        azure.Azure
	CloudStack   cloudstack.CloudStack
	DigitalOcean digitalocean.DigitalOcean
	GitHub       github.GitHub
	Google       google.Google
	Kubernetes   kubernetes.Kubernetes
	OpenStack    openstack.OpenStack
	Oracle       oracle.Oracle
	Nifcloud     nifcloud.Nifcloud
}

func (a *State) ToRego() any {
	return convert.StructToRego(reflect.ValueOf(a))
}
