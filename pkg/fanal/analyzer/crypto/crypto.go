package crypto

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"

	"vulnscan/pkg/crypto/parser/x509"
	"vulnscan/pkg/fanal/analyzer"
	"vulnscan/pkg/log"
	"vulnscan/pkg/set"
	xio "vulnscan/pkg/x/io"
)

func init() {
	analyzer.RegisterAnalyzer(&cryptoAnalyzer{})
}

const (
	version = 2

	// maxFileSize bounds what is read into memory. The eligible extensions are shared with
	// unrelated formats, such as Keynote presentations, which grow far beyond any file of
	// cryptographic material.
	maxFileSize = 10 << 20 // 10MB
)

var requiredExtensions = set.NewCaseInsensitive(".pem", ".der", ".crt", ".cer", ".key")

// cryptoAnalyzer inventories cryptographic assets.
type cryptoAnalyzer struct{}

func (a *cryptoAnalyzer) Analyze(ctx context.Context, input analyzer.AnalysisInput) (*analyzer.AnalysisResult, error) {
	// A file matched by --file-patterns reaches the analyzer without passing Required.
	content, err := xio.ReadAllWithLimit(input.Content, maxFileSize)
	if errors.Is(err, xio.ErrLimitExceeded) {
		logSkippedFile(input.FilePath)
		return nil, nil
	} else if err != nil {
		return nil, xerrors.Errorf("read %s: %w", input.FilePath, err)
	}

	assets, err := x509.Parse(ctx, input.FilePath, content)
	if err != nil {
		return nil, xerrors.Errorf("parse %s: %w", input.FilePath, err)
	}
	if len(assets) == 0 {
		return nil, nil
	}
	return &analyzer.AnalysisResult{CryptoAssets: assets}, nil
}

// Required selects a file by extension and size, without reading its content.
func (a *cryptoAnalyzer) Required(filePath string, info os.FileInfo) bool {
	if !requiredExtensions.Contains(filepath.Ext(filePath)) {
		return false
	}
	if info.Size() > maxFileSize {
		logSkippedFile(filePath)
		return false
	}
	return true
}

// logSkippedFile reports a file left out of the inventory for its size.
func logSkippedFile(filePath string) {
	log.WithPrefix("crypto").Debug("File is too large to read", log.FilePath(filePath))
}

func (a *cryptoAnalyzer) Type() analyzer.Type {
	return analyzer.TypeCrypto
}

func (a *cryptoAnalyzer) Version() int {
	return version
}
