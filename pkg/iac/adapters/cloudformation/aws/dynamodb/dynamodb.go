package dynamodb

import (
	"vulnscan/pkg/iac/providers/aws/dynamodb"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a dynamodb instance
func Adapt(cfFile parser.FileContext) dynamodb.DynamoDB {
	return dynamodb.DynamoDB{
		DAXClusters: getClusters(cfFile),
		Tables:      getTables(cfFile),
	}
}
