package aws

import (
	"vulnscan/pkg/iac/adapters/terraform/aws/apigateway"
	"vulnscan/pkg/iac/adapters/terraform/aws/athena"
	"vulnscan/pkg/iac/adapters/terraform/aws/cloudfront"
	"vulnscan/pkg/iac/adapters/terraform/aws/cloudtrail"
	"vulnscan/pkg/iac/adapters/terraform/aws/cloudwatch"
	"vulnscan/pkg/iac/adapters/terraform/aws/codebuild"
	"vulnscan/pkg/iac/adapters/terraform/aws/config"
	"vulnscan/pkg/iac/adapters/terraform/aws/documentdb"
	"vulnscan/pkg/iac/adapters/terraform/aws/dynamodb"
	"vulnscan/pkg/iac/adapters/terraform/aws/ec2"
	"vulnscan/pkg/iac/adapters/terraform/aws/ecr"
	"vulnscan/pkg/iac/adapters/terraform/aws/ecs"
	"vulnscan/pkg/iac/adapters/terraform/aws/efs"
	"vulnscan/pkg/iac/adapters/terraform/aws/eks"
	"vulnscan/pkg/iac/adapters/terraform/aws/elasticache"
	"vulnscan/pkg/iac/adapters/terraform/aws/elasticsearch"
	"vulnscan/pkg/iac/adapters/terraform/aws/elb"
	"vulnscan/pkg/iac/adapters/terraform/aws/emr"
	"vulnscan/pkg/iac/adapters/terraform/aws/iam"
	"vulnscan/pkg/iac/adapters/terraform/aws/kinesis"
	"vulnscan/pkg/iac/adapters/terraform/aws/kms"
	"vulnscan/pkg/iac/adapters/terraform/aws/lambda"
	"vulnscan/pkg/iac/adapters/terraform/aws/mq"
	"vulnscan/pkg/iac/adapters/terraform/aws/msk"
	"vulnscan/pkg/iac/adapters/terraform/aws/neptune"
	"vulnscan/pkg/iac/adapters/terraform/aws/provider"
	"vulnscan/pkg/iac/adapters/terraform/aws/rds"
	"vulnscan/pkg/iac/adapters/terraform/aws/redshift"
	"vulnscan/pkg/iac/adapters/terraform/aws/s3"
	"vulnscan/pkg/iac/adapters/terraform/aws/sns"
	"vulnscan/pkg/iac/adapters/terraform/aws/sqs"
	"vulnscan/pkg/iac/adapters/terraform/aws/ssm"
	"vulnscan/pkg/iac/adapters/terraform/aws/workspaces"
	"vulnscan/pkg/iac/providers/aws"
	"vulnscan/pkg/iac/terraform"
)

func Adapt(modules terraform.Modules) aws.AWS {
	return aws.AWS{
		Meta: aws.Meta{
			TFProviders: provider.Adapt(modules),
		},
		APIGateway:    apigateway.Adapt(modules),
		Athena:        athena.Adapt(modules),
		Cloudfront:    cloudfront.Adapt(modules),
		CloudTrail:    cloudtrail.Adapt(modules),
		CloudWatch:    cloudwatch.Adapt(modules),
		CodeBuild:     codebuild.Adapt(modules),
		Config:        config.Adapt(modules),
		DocumentDB:    documentdb.Adapt(modules),
		DynamoDB:      dynamodb.Adapt(modules),
		EC2:           ec2.Adapt(modules),
		ECR:           ecr.Adapt(modules),
		ECS:           ecs.Adapt(modules),
		EFS:           efs.Adapt(modules),
		EKS:           eks.Adapt(modules),
		ElastiCache:   elasticache.Adapt(modules),
		Elasticsearch: elasticsearch.Adapt(modules),
		ELB:           elb.Adapt(modules),
		EMR:           emr.Adapt(modules),
		IAM:           iam.Adapt(modules),
		Kinesis:       kinesis.Adapt(modules),
		KMS:           kms.Adapt(modules),
		Lambda:        lambda.Adapt(modules),
		MQ:            mq.Adapt(modules),
		MSK:           msk.Adapt(modules),
		Neptune:       neptune.Adapt(modules),
		RDS:           rds.Adapt(modules),
		Redshift:      redshift.Adapt(modules),
		S3:            s3.Adapt(modules),
		SNS:           sns.Adapt(modules),
		SQS:           sqs.Adapt(modules),
		SSM:           ssm.Adapt(modules),
		WorkSpaces:    workspaces.Adapt(modules),
	}
}
