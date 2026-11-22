package azure

import (
	"vulnscan/pkg/iac/providers/azure/appservice"
	"vulnscan/pkg/iac/providers/azure/authorization"
	"vulnscan/pkg/iac/providers/azure/compute"
	"vulnscan/pkg/iac/providers/azure/container"
	"vulnscan/pkg/iac/providers/azure/cosmosdb"
	"vulnscan/pkg/iac/providers/azure/database"
	"vulnscan/pkg/iac/providers/azure/datafactory"
	"vulnscan/pkg/iac/providers/azure/datalake"
	"vulnscan/pkg/iac/providers/azure/keyvault"
	"vulnscan/pkg/iac/providers/azure/monitor"
	"vulnscan/pkg/iac/providers/azure/network"
	"vulnscan/pkg/iac/providers/azure/securitycenter"
	"vulnscan/pkg/iac/providers/azure/storage"
	"vulnscan/pkg/iac/providers/azure/synapse"
)

type Azure struct {
	AppService     appservice.AppService
	Authorization  authorization.Authorization
	Compute        compute.Compute
	Container      container.Container
	CosmosDB       cosmosdb.CosmosDB
	Database       database.Database
	DataFactory    datafactory.DataFactory
	DataLake       datalake.DataLake
	KeyVault       keyvault.KeyVault
	Monitor        monitor.Monitor
	Network        network.Network
	SecurityCenter securitycenter.SecurityCenter
	Storage        storage.Storage
	Synapse        synapse.Synapse
}
