package aws

import (
	"vulnscan/pkg/iac/adapters/cloudformation/aws/apigateway"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/athena"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/cloudfront"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/cloudtrail"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/cloudwatch"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/codebuild"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/config"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/documentdb"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/dynamodb"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/ec2"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/ecr"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/ecs"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/efs"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/eks"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/elasticache"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/elasticsearch"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/elb"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/iam"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/kinesis"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/lambda"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/mq"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/msk"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/neptune"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/rds"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/redshift"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/s3"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/sam"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/sns"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/sqs"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/ssm"
	"vulnscan/pkg/iac/adapters/cloudformation/aws/workspaces"
	"vulnscan/pkg/iac/providers/aws"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts a Cloudformation AWS instance
func Adapt(cfFile parser.FileContext) aws.AWS {
	return aws.AWS{
		APIGateway:    apigateway.Adapt(cfFile),
		Athena:        athena.Adapt(cfFile),
		Cloudfront:    cloudfront.Adapt(cfFile),
		CloudTrail:    cloudtrail.Adapt(cfFile),
		CloudWatch:    cloudwatch.Adapt(cfFile),
		CodeBuild:     codebuild.Adapt(cfFile),
		Config:        config.Adapt(cfFile),
		DocumentDB:    documentdb.Adapt(cfFile),
		DynamoDB:      dynamodb.Adapt(cfFile),
		EC2:           ec2.Adapt(cfFile),
		ECR:           ecr.Adapt(cfFile),
		ECS:           ecs.Adapt(cfFile),
		EFS:           efs.Adapt(cfFile),
		IAM:           iam.Adapt(cfFile),
		EKS:           eks.Adapt(cfFile),
		ElastiCache:   elasticache.Adapt(cfFile),
		Elasticsearch: elasticsearch.Adapt(cfFile),
		ELB:           elb.Adapt(cfFile),
		MSK:           msk.Adapt(cfFile),
		MQ:            mq.Adapt(cfFile),
		Kinesis:       kinesis.Adapt(cfFile),
		Lambda:        lambda.Adapt(cfFile),
		Neptune:       neptune.Adapt(cfFile),
		RDS:           rds.Adapt(cfFile),
		Redshift:      redshift.Adapt(cfFile),
		S3:            s3.Adapt(cfFile),
		SAM:           sam.Adapt(cfFile),
		SNS:           sns.Adapt(cfFile),
		SQS:           sqs.Adapt(cfFile),
		SSM:           ssm.Adapt(cfFile),
		WorkSpaces:    workspaces.Adapt(cfFile),
	}
}
