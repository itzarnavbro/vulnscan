package github

import (
	"vulnscan/pkg/iac/adapters/terraform/github/branch_protections"
	"vulnscan/pkg/iac/adapters/terraform/github/repositories"
	"vulnscan/pkg/iac/adapters/terraform/github/secrets"
	"vulnscan/pkg/iac/providers/github"
	"vulnscan/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) github.GitHub {
	return github.GitHub{
		Repositories:       repositories.Adapt(modules),
		EnvironmentSecrets: secrets.Adapt(modules),
		BranchProtections:  branch_protections.Adapt(modules),
	}
}
