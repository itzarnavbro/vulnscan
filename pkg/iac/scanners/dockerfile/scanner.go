package dockerfile

import (
	"vulnscan/pkg/iac/providers/dockerfile"
	"vulnscan/pkg/iac/scanners/dockerfile/parser"
	"vulnscan/pkg/iac/scanners/generic"
	"vulnscan/pkg/iac/scanners/options"
	"vulnscan/pkg/iac/types"
)

func NewScanner(opts ...options.ScannerOption) *generic.GenericScanner[*dockerfile.Dockerfile] {
	defaultOpts := []options.ScannerOption{
		generic.WithSupportsInlineIgnore[*dockerfile.Dockerfile](true),
	}

	p := parser.NewParser()
	return generic.NewScanner("Dockerfile", types.SourceDockerfile, p, append(defaultOpts, opts...)...,
	)
}
