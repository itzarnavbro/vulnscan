package tftestutil

import (
	"testing"

	"github.com/stretchr/testify/require"

	"vulnscan/internal/testutil"
	"vulnscan/pkg/iac/scanners/terraform/parser"
	"vulnscan/pkg/iac/terraform"
)

func CreateModulesFromSource(t *testing.T, source, ext string) terraform.Modules {
	fs := testutil.CreateFS(map[string]string{
		"source" + ext: source,
	})
	p := parser.New(fs, "", parser.OptionStopOnHCLError(true))
	require.NoError(t, p.ParseFS(t.Context(), "."))
	modules, err := p.EvaluateAll(t.Context())
	require.NoError(t, err)
	return modules
}
