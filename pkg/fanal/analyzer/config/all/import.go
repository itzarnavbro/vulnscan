package all

import (
	_ "vulnscan/pkg/fanal/analyzer/config/ansible"
	_ "vulnscan/pkg/fanal/analyzer/config/azurearm"
	_ "vulnscan/pkg/fanal/analyzer/config/cloudformation"
	_ "vulnscan/pkg/fanal/analyzer/config/dockerfile"
	_ "vulnscan/pkg/fanal/analyzer/config/helm"
	_ "vulnscan/pkg/fanal/analyzer/config/json"
	_ "vulnscan/pkg/fanal/analyzer/config/k8s"
	_ "vulnscan/pkg/fanal/analyzer/config/terraform"
	_ "vulnscan/pkg/fanal/analyzer/config/terraformplan/json"
	_ "vulnscan/pkg/fanal/analyzer/config/terraformplan/snapshot"
	_ "vulnscan/pkg/fanal/analyzer/config/yaml"
)
