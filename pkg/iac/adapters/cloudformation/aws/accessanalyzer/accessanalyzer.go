package accessanalyzer

import (
	"vulnscan/pkg/iac/providers/aws/accessanalyzer"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an AccessAnalyzer instance
func Adapt(cfFile parser.FileContext) accessanalyzer.AccessAnalyzer {
	return accessanalyzer.AccessAnalyzer{
		Analyzers: getAccessAnalyzer(cfFile),
	}
}
