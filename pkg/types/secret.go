package types

import (
	ftypes "vulnscan/pkg/fanal/types"
)

type DetectedSecret ftypes.SecretFinding

func (DetectedSecret) FindingType() FindingType { return FindingTypeSecret }
