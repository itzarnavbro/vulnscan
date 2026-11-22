package pipenv

import (
	"context"
	"os"
	"path/filepath"
	"slices"

	"golang.org/x/xerrors"

	"vulnscan/pkg/dependency/parser/python/pipenv"
	"vulnscan/pkg/fanal/analyzer"
	"vulnscan/pkg/fanal/analyzer/language"
	"vulnscan/pkg/fanal/types"
)

func init() {
	analyzer.RegisterAnalyzer(&pipenvLibraryAnalyzer{})
}

const version = 1

var requiredFiles = []string{types.PipfileLock}

type pipenvLibraryAnalyzer struct{}

func (a pipenvLibraryAnalyzer) Analyze(ctx context.Context, input analyzer.AnalysisInput) (*analyzer.AnalysisResult, error) {
	res, err := language.Analyze(ctx, types.Pipenv, input.FilePath, input.Content, pipenv.NewParser())
	if err != nil {
		return nil, xerrors.Errorf("unable to parse Pipfile.lock: %w", err)
	}
	return res, nil
}

func (a pipenvLibraryAnalyzer) Required(filePath string, _ os.FileInfo) bool {
	fileName := filepath.Base(filePath)
	return slices.Contains(requiredFiles, fileName)
}

func (a pipenvLibraryAnalyzer) Type() analyzer.Type {
	return analyzer.TypePipenv
}

func (a pipenvLibraryAnalyzer) Version() int {
	return version
}
