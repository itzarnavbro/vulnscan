package google

import (
	"vulnscan/pkg/iac/adapters/terraform/google/bigquery"
	"vulnscan/pkg/iac/adapters/terraform/google/compute"
	"vulnscan/pkg/iac/adapters/terraform/google/dns"
	"vulnscan/pkg/iac/adapters/terraform/google/gke"
	"vulnscan/pkg/iac/adapters/terraform/google/iam"
	"vulnscan/pkg/iac/adapters/terraform/google/kms"
	"vulnscan/pkg/iac/adapters/terraform/google/sql"
	"vulnscan/pkg/iac/adapters/terraform/google/storage"
	"vulnscan/pkg/iac/providers/google"
	"vulnscan/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) google.Google {
	return google.Google{
		BigQuery: bigquery.Adapt(modules),
		Compute:  compute.Adapt(modules),
		DNS:      dns.Adapt(modules),
		GKE:      gke.Adapt(modules),
		KMS:      kms.Adapt(modules),
		IAM:      iam.Adapt(modules),
		SQL:      sql.Adapt(modules),
		Storage:  storage.Adapt(modules),
	}
}
