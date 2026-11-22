package ansible

import (
	"vulnscan/pkg/iac/adapters/ansible/aws/s3"
	"vulnscan/pkg/iac/providers/aws"
	"vulnscan/pkg/iac/scanners/ansible/parser"
	"vulnscan/pkg/iac/state"
)

func Adapt(tasks parser.ResolvedTasks) state.State {
	return state.State{
		AWS: aws.AWS{
			S3: s3.Adapt(tasks),
			// TODO(simar): Add other AWS services
		},
	}
}
