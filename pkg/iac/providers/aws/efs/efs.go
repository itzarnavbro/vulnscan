package efs

import (
	iacTypes "vulnscan/pkg/iac/types"
)

type EFS struct {
	FileSystems []FileSystem
}

type FileSystem struct {
	Metadata  iacTypes.Metadata
	Encrypted iacTypes.BoolValue
}
