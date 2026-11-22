package arm

import (
	"context"

	"vulnscan/pkg/iac/adapters/arm/appservice"
	"vulnscan/pkg/iac/adapters/arm/authorization"
	"vulnscan/pkg/iac/adapters/arm/compute"
	"vulnscan/pkg/iac/adapters/arm/container"
	"vulnscan/pkg/iac/adapters/arm/database"
	"vulnscan/pkg/iac/adapters/arm/datafactory"
	"vulnscan/pkg/iac/adapters/arm/datalake"
	"vulnscan/pkg/iac/adapters/arm/keyvault"
	"vulnscan/pkg/iac/adapters/arm/monitor"
	"vulnscan/pkg/iac/adapters/arm/network"
	"vulnscan/pkg/iac/adapters/arm/securitycenter"
	"vulnscan/pkg/iac/adapters/arm/storage"
	"vulnscan/pkg/iac/adapters/arm/synapse"
	"vulnscan/pkg/iac/providers/azure"
	scanner "vulnscan/pkg/iac/scanners/azure"
	"vulnscan/pkg/iac/state"
)

// Adapt adapts an azure arm instance
func Adapt(_ context.Context, deployment scanner.Deployment) *state.State {
	return &state.State{
		Azure: adaptAzure(deployment),
	}
}

func adaptAzure(deployment scanner.Deployment) azure.Azure {

	return azure.Azure{
		AppService:     appservice.Adapt(deployment),
		Authorization:  authorization.Adapt(deployment),
		Compute:        compute.Adapt(deployment),
		Container:      container.Adapt(deployment),
		Database:       database.Adapt(deployment),
		DataFactory:    datafactory.Adapt(deployment),
		DataLake:       datalake.Adapt(deployment),
		KeyVault:       keyvault.Adapt(deployment),
		Monitor:        monitor.Adapt(deployment),
		Network:        network.Adapt(deployment),
		SecurityCenter: securitycenter.Adapt(deployment),
		Storage:        storage.Adapt(deployment),
		Synapse:        synapse.Adapt(deployment),
	}

}
