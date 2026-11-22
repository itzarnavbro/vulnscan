package cloudfront

import (
	"vulnscan/pkg/iac/providers/aws/cloudfront"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a CloudFront instance
func Adapt(cfFile parser.FileContext) cloudfront.Cloudfront {
	return cloudfront.Cloudfront{
		Distributions: getDistributions(cfFile),
	}
}
