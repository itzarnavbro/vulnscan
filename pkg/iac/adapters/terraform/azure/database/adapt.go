package database

import (
	"vulnscan/pkg/iac/providers/azure/database"
	"vulnscan/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) database.Database {
	return database.Database{
		MSSQLServers:      adaptMSSQLServers(modules),
		MariaDBServers:    adaptMariaDBServers(modules),
		MySQLServers:      adaptMySQLServers(modules),
		PostgreSQLServers: adaptPostgreSQLServers(modules),
	}
}
