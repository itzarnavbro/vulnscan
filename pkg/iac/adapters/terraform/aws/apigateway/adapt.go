package apigateway

import (
	"vulnscan/pkg/iac/providers/aws/apigateway"
	v1 "vulnscan/pkg/iac/providers/aws/apigateway/v1"
	v2 "vulnscan/pkg/iac/providers/aws/apigateway/v2"
	"vulnscan/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) apigateway.APIGateway {
	return apigateway.APIGateway{
		V1: v1.APIGateway{
			APIs:        adaptAPIsV1(modules),
			DomainNames: adaptDomainNamesV1(modules),
		},
		V2: v2.APIGateway{
			APIs:        adaptAPIsV2(modules),
			DomainNames: adaptDomainNamesV2(modules),
		},
	}
}
