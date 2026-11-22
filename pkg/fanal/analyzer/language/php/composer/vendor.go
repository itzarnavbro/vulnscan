package composer

import (
	"context"
	"os"
	"path/filepath"

	"vulnscan/pkg/dependency/parser/php/composer"
	"vulnscan/pkg/fanal/analyzer"
	"vulnscan/pkg/fanal/analyzer/language"
	"vulnscan/pkg/fanal/types"
)

func init() {
	analyzer.RegisterAnalyzer(&composerVendorAnalyzer{})
}

const (
	composerInstalledAnalyzerVersion = 1
)

// composerVendorAnalyzer analyzes 'installed.json'
type composerVendorAnalyzer struct{}

func (a composerVendorAnalyzer) Analyze(ctx context.Context, input analyzer.AnalysisInput) (*analyzer.AnalysisResult, error) {
	return language.Analyze(ctx, types.ComposerVendor, input.FilePath, input.Content, composer.NewParser())
}

func (a composerVendorAnalyzer) Required(filePath string, _ os.FileInfo) bool {
	return filepath.Base(filePath) == types.ComposerInstalledJson
}

func (a composerVendorAnalyzer) Type() analyzer.Type {
	return analyzer.TypeComposerVendor
}

func (a composerVendorAnalyzer) Version() int {
	return composerInstalledAnalyzerVersion
}
