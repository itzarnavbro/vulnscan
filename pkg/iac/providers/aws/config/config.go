package config

import (
	iacTypes "vulnscan/pkg/iac/types"
)

type Config struct {
	ConfigurationAggregrator ConfigurationAggregrator
}

type ConfigurationAggregrator struct {
	Metadata         iacTypes.Metadata
	SourceAllRegions iacTypes.BoolValue
}
