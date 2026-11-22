package msk

import (
	"vulnscan/pkg/iac/providers/aws/msk"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an MSK instance
func Adapt(cfFile parser.FileContext) msk.MSK {
	return msk.MSK{
		Clusters: getClusters(cfFile),
	}
}
