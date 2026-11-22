package mq

import (
	"vulnscan/pkg/iac/providers/aws/mq"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an MQ instance
func Adapt(cfFile parser.FileContext) mq.MQ {
	return mq.MQ{
		Brokers: getBrokers(cfFile),
	}
}
