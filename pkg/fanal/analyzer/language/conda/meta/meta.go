package meta

import (
	"context"
	"os"
	"path/filepath"
	"regexp"

	"vulnscan/pkg/dependency/parser/conda/meta"
	"vulnscan/pkg/fanal/analyzer"
	"vulnscan/pkg/fanal/analyzer/language"
	"vulnscan/pkg/fanal/types"
)

func init() {
	analyzer.RegisterAnalyzer(&metaAnalyzer{})
}

const version = 1

var fileRegex = regexp.MustCompile(`.*/envs/.+/conda-meta/.+-.+-.+\.json`)

type metaAnalyzer struct{}

func (a metaAnalyzer) Analyze(ctx context.Context, input analyzer.AnalysisInput) (*analyzer.AnalysisResult, error) {
	p := meta.NewParser()
	return language.AnalyzePackage(ctx, types.CondaPkg, input.FilePath, input.Content, p, input.Options.FileChecksum)
}
func (a metaAnalyzer) Required(filePath string, _ os.FileInfo) bool {
	return fileRegex.MatchString(filepath.ToSlash(filePath))
}

func (a metaAnalyzer) Type() analyzer.Type {
	return analyzer.TypeCondaPkg
}

func (a metaAnalyzer) Version() int {
	return version
}
