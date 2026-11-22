package kubernetes

import (
	"context"
	"io"

	"vulnscan/pkg/iac/scanners/generic"
	"vulnscan/pkg/iac/scanners/kubernetes/parser"
	"vulnscan/pkg/iac/scanners/options"
	"vulnscan/pkg/iac/types"
)

func NewScanner(opts ...options.ScannerOption) *generic.GenericScanner[*parser.Manifest] {
	p := generic.ParseFunc[*parser.Manifest](parse)
	return generic.NewScanner("Kubernetes", types.SourceKubernetes, p, opts...)
}

func parse(ctx context.Context, r io.Reader, path string) ([]*parser.Manifest, error) {
	return parser.Parse(ctx, r, path)
}
