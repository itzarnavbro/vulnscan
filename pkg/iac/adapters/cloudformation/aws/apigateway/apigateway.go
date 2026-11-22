package apigateway

import (
	"vulnscan/pkg/iac/providers/aws/apigateway"
	v1 "vulnscan/pkg/iac/providers/aws/apigateway/v1"
	v2 "vulnscan/pkg/iac/providers/aws/apigateway/v2"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an APIGateway instance
func Adapt(cfFile parser.FileContext) apigateway.APIGateway {
	return apigateway.APIGateway{
		V1: v1.APIGateway{
			APIs:        adaptAPIsV1(cfFile),
			DomainNames: adaptDomainNamesV1(cfFile),
		},
		V2: v2.APIGateway{
			APIs:        adaptAPIsV2(cfFile),
			DomainNames: adaptDomainNamesV2(cfFile),
		},
	}
}
