package datafactory

import (
	iacTypes "vulnscan/pkg/iac/types"
)

type DataFactory struct {
	DataFactories []Factory
}

type Factory struct {
	Metadata            iacTypes.Metadata
	EnablePublicNetwork iacTypes.BoolValue
}
