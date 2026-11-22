package dynamodb

import (
	"vulnscan/pkg/iac/providers/aws/dynamodb"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
	xslices "vulnscan/pkg/x/slices"
)

func getTables(fctx parser.FileContext) []dynamodb.Table {
	return xslices.Map(fctx.GetResourcesByType("AWS::DynamoDB::Table"), func(
		resource *parser.Resource,
	) dynamodb.Table {
		sseSpec := resource.GetProperty("SSESpecification")
		return dynamodb.Table{
			Metadata: resource.Metadata(),
			ServerSideEncryption: dynamodb.ServerSideEncryption{
				Metadata: sseSpec.Metadata(),
				Enabled:  sseSpec.GetBoolProperty("SSEEnabled"),
				KMSKeyID: sseSpec.GetStringProperty("KMSMasterKeyId", dynamodb.DefaultKMSKeyID),
			},
			PointInTimeRecovery: resource.GetBoolProperty("PointInTimeRecoverySpecification.PointInTimeRecoveryEnabled"),
		}
	})
}
