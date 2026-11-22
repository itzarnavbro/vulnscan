package datalake

import (
	iacTypes "vulnscan/pkg/iac/types"
)

type DataLake struct {
	Stores []Store
}

type Store struct {
	Metadata         iacTypes.Metadata
	EnableEncryption iacTypes.BoolValue
}
