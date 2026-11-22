package rdb

import (
	"vulnscan/pkg/iac/providers/nifcloud/rdb"
	"vulnscan/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) rdb.RDB {
	return rdb.RDB{
		DBSecurityGroups: adaptDBSecurityGroups(modules),
		DBInstances:      adaptDBInstances(modules),
	}
}
