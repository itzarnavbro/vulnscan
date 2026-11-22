package codebuild

import (
	"vulnscan/pkg/iac/providers/aws/codebuild"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a CodeBuild instance
func Adapt(cfFile parser.FileContext) codebuild.CodeBuild {
	return codebuild.CodeBuild{
		Projects: getProjects(cfFile),
	}
}
