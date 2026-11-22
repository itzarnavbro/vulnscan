package azure

import (
	"vulnscan/pkg/iac/adapters/terraform/azure/appservice"
	"vulnscan/pkg/iac/adapters/terraform/azure/authorization"
	"vulnscan/pkg/iac/adapters/terraform/azure/compute"
	"vulnscan/pkg/iac/adapters/terraform/azure/container"
	"vulnscan/pkg/iac/adapters/terraform/azure/cosmosdb"
	"vulnscan/pkg/iac/adapters/terraform/azure/database"
	"vulnscan/pkg/iac/adapters/terraform/azure/datafactory"
	"vulnscan/pkg/iac/adapters/terraform/azure/datalake"
	"vulnscan/pkg/iac/adapters/terraform/azure/keyvault"
	"vulnscan/pkg/iac/adapters/terraform/azure/monitor"
	"vulnscan/pkg/iac/adapters/terraform/azure/network"
	"vulnscan/pkg/iac/adapters/terraform/azure/securitycenter"
	"vulnscan/pkg/iac/adapters/terraform/azure/storage"
	"vulnscan/pkg/iac/adapters/terraform/azure/synapse"
	"vulnscan/pkg/iac/providers/azure"
	"vulnscan/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) azure.Azure {
	return azure.Azure{
		AppService:     appservice.Adapt(modules),
		Authorization:  authorization.Adapt(modules),
		Compute:        compute.Adapt(modules),
		Container:      container.Adapt(modules),
		CosmosDB:       cosmosdb.Adapt(modules),
		Database:       database.Adapt(modules),
		DataFactory:    datafactory.Adapt(modules),
		DataLake:       datalake.Adapt(modules),
		KeyVault:       keyvault.Adapt(modules),
		Monitor:        monitor.Adapt(modules),
		Network:        network.Adapt(modules),
		SecurityCenter: securitycenter.Adapt(modules),
		Storage:        storage.Adapt(modules),
		Synapse:        synapse.Adapt(modules),
	}
}
