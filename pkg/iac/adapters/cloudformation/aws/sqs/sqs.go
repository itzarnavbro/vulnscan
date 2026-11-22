package sqs

import (
	"vulnscan/pkg/iac/providers/aws/sqs"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an SQS instance
func Adapt(cfFile parser.FileContext) sqs.SQS {
	return sqs.SQS{
		Queues: getQueues(cfFile),
	}
}
