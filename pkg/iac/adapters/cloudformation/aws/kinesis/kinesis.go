package kinesis

import (
	"vulnscan/pkg/iac/providers/aws/kinesis"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a Kinesis instance
func Adapt(cfFile parser.FileContext) kinesis.Kinesis {
	return kinesis.Kinesis{
		Streams: getStreams(cfFile),
	}
}
