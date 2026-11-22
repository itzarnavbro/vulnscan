package accessanalyzer

import (
	"vulnscan/pkg/iac/providers/aws/accessanalyzer"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
	"vulnscan/pkg/iac/types"
)

func getAccessAnalyzer(ctx parser.FileContext) (analyzers []accessanalyzer.Analyzer) {

	analyzersList := ctx.GetResourcesByType("AWS::AccessAnalyzer::Analyzer")

	for _, r := range analyzersList {
		aa := accessanalyzer.Analyzer{
			Metadata: r.Metadata(),
			Name:     r.GetStringProperty("AnalyzerName"),
			ARN:      r.StringDefault(""),
			Active:   types.BoolDefault(false, r.Metadata()),
		}

		analyzers = append(analyzers, aa)
	}
	return analyzers
}
