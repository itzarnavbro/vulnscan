package google

import (
	"vulnscan/pkg/iac/providers/google/bigquery"
	"vulnscan/pkg/iac/providers/google/compute"
	"vulnscan/pkg/iac/providers/google/dns"
	"vulnscan/pkg/iac/providers/google/gke"
	"vulnscan/pkg/iac/providers/google/iam"
	"vulnscan/pkg/iac/providers/google/kms"
	"vulnscan/pkg/iac/providers/google/sql"
	"vulnscan/pkg/iac/providers/google/storage"
)

type Google struct {
	BigQuery bigquery.BigQuery
	Compute  compute.Compute
	DNS      dns.DNS
	GKE      gke.GKE
	KMS      kms.KMS
	IAM      iam.IAM
	SQL      sql.SQL
	Storage  storage.Storage
}
