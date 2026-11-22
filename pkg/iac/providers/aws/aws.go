package aws

import (
	"vulnscan/pkg/iac/providers/aws/accessanalyzer"
	"vulnscan/pkg/iac/providers/aws/apigateway"
	"vulnscan/pkg/iac/providers/aws/athena"
	"vulnscan/pkg/iac/providers/aws/cloudfront"
	"vulnscan/pkg/iac/providers/aws/cloudtrail"
	"vulnscan/pkg/iac/providers/aws/cloudwatch"
	"vulnscan/pkg/iac/providers/aws/codebuild"
	"vulnscan/pkg/iac/providers/aws/config"
	"vulnscan/pkg/iac/providers/aws/documentdb"
	"vulnscan/pkg/iac/providers/aws/dynamodb"
	"vulnscan/pkg/iac/providers/aws/ec2"
	"vulnscan/pkg/iac/providers/aws/ecr"
	"vulnscan/pkg/iac/providers/aws/ecs"
	"vulnscan/pkg/iac/providers/aws/efs"
	"vulnscan/pkg/iac/providers/aws/eks"
	"vulnscan/pkg/iac/providers/aws/elasticache"
	"vulnscan/pkg/iac/providers/aws/elasticsearch"
	"vulnscan/pkg/iac/providers/aws/elb"
	"vulnscan/pkg/iac/providers/aws/emr"
	"vulnscan/pkg/iac/providers/aws/iam"
	"vulnscan/pkg/iac/providers/aws/kinesis"
	"vulnscan/pkg/iac/providers/aws/kms"
	"vulnscan/pkg/iac/providers/aws/lambda"
	"vulnscan/pkg/iac/providers/aws/mq"
	"vulnscan/pkg/iac/providers/aws/msk"
	"vulnscan/pkg/iac/providers/aws/neptune"
	"vulnscan/pkg/iac/providers/aws/rds"
	"vulnscan/pkg/iac/providers/aws/redshift"
	"vulnscan/pkg/iac/providers/aws/s3"
	"vulnscan/pkg/iac/providers/aws/sam"
	"vulnscan/pkg/iac/providers/aws/sns"
	"vulnscan/pkg/iac/providers/aws/sqs"
	"vulnscan/pkg/iac/providers/aws/ssm"
	"vulnscan/pkg/iac/providers/aws/workspaces"
)

type AWS struct {
	Meta           Meta
	AccessAnalyzer accessanalyzer.AccessAnalyzer
	APIGateway     apigateway.APIGateway
	Athena         athena.Athena
	Cloudfront     cloudfront.Cloudfront
	CloudTrail     cloudtrail.CloudTrail
	CloudWatch     cloudwatch.CloudWatch
	CodeBuild      codebuild.CodeBuild
	Config         config.Config
	DocumentDB     documentdb.DocumentDB
	DynamoDB       dynamodb.DynamoDB
	EC2            ec2.EC2
	ECR            ecr.ECR
	ECS            ecs.ECS
	EFS            efs.EFS
	EKS            eks.EKS
	ElastiCache    elasticache.ElastiCache
	Elasticsearch  elasticsearch.Elasticsearch
	ELB            elb.ELB
	EMR            emr.EMR
	IAM            iam.IAM
	Kinesis        kinesis.Kinesis
	KMS            kms.KMS
	Lambda         lambda.Lambda
	MQ             mq.MQ
	MSK            msk.MSK
	Neptune        neptune.Neptune
	RDS            rds.RDS
	Redshift       redshift.Redshift
	SAM            sam.SAM
	S3             s3.S3
	SNS            sns.SNS
	SQS            sqs.SQS
	SSM            ssm.SSM
	WorkSpaces     workspaces.WorkSpaces
}

type Meta struct {
	TFProviders []TerraformProvider
}
