package cloudawsgovcloudintegrations

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudAwsGovcloudIntegrationsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the linked AwsGovCloud account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#linked_account_id CloudAwsGovcloudIntegrations#linked_account_id}
	LinkedAccountId *float64 `field:"required" json:"linkedAccountId" yaml:"linkedAccountId"`
	// The ID of the account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#account_id CloudAwsGovcloudIntegrations#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// alb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#alb CloudAwsGovcloudIntegrations#alb}
	Alb *CloudAwsGovcloudIntegrationsAlb `field:"optional" json:"alb" yaml:"alb"`
	// api_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#api_gateway CloudAwsGovcloudIntegrations#api_gateway}
	ApiGateway *CloudAwsGovcloudIntegrationsApiGateway `field:"optional" json:"apiGateway" yaml:"apiGateway"`
	// auto_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#auto_scaling CloudAwsGovcloudIntegrations#auto_scaling}
	AutoScaling *CloudAwsGovcloudIntegrationsAutoScaling `field:"optional" json:"autoScaling" yaml:"autoScaling"`
	// aws_direct_connect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#aws_direct_connect CloudAwsGovcloudIntegrations#aws_direct_connect}
	AwsDirectConnect *CloudAwsGovcloudIntegrationsAwsDirectConnect `field:"optional" json:"awsDirectConnect" yaml:"awsDirectConnect"`
	// aws_states block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#aws_states CloudAwsGovcloudIntegrations#aws_states}
	AwsStates *CloudAwsGovcloudIntegrationsAwsStates `field:"optional" json:"awsStates" yaml:"awsStates"`
	// cloudtrail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#cloudtrail CloudAwsGovcloudIntegrations#cloudtrail}
	Cloudtrail *CloudAwsGovcloudIntegrationsCloudtrail `field:"optional" json:"cloudtrail" yaml:"cloudtrail"`
	// dynamo_db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#dynamo_db CloudAwsGovcloudIntegrations#dynamo_db}
	DynamoDb *CloudAwsGovcloudIntegrationsDynamoDb `field:"optional" json:"dynamoDb" yaml:"dynamoDb"`
	// ebs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#ebs CloudAwsGovcloudIntegrations#ebs}
	Ebs *CloudAwsGovcloudIntegrationsEbs `field:"optional" json:"ebs" yaml:"ebs"`
	// ec2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#ec2 CloudAwsGovcloudIntegrations#ec2}
	Ec2 *CloudAwsGovcloudIntegrationsEc2 `field:"optional" json:"ec2" yaml:"ec2"`
	// elastic_search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#elastic_search CloudAwsGovcloudIntegrations#elastic_search}
	ElasticSearch *CloudAwsGovcloudIntegrationsElasticSearch `field:"optional" json:"elasticSearch" yaml:"elasticSearch"`
	// elb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#elb CloudAwsGovcloudIntegrations#elb}
	Elb *CloudAwsGovcloudIntegrationsElb `field:"optional" json:"elb" yaml:"elb"`
	// emr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#emr CloudAwsGovcloudIntegrations#emr}
	Emr *CloudAwsGovcloudIntegrationsEmr `field:"optional" json:"emr" yaml:"emr"`
	// iam block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#iam CloudAwsGovcloudIntegrations#iam}
	Iam *CloudAwsGovcloudIntegrationsIam `field:"optional" json:"iam" yaml:"iam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#id CloudAwsGovcloudIntegrations#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#lambda CloudAwsGovcloudIntegrations#lambda}
	Lambda *CloudAwsGovcloudIntegrationsLambda `field:"optional" json:"lambda" yaml:"lambda"`
	// rds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#rds CloudAwsGovcloudIntegrations#rds}
	Rds *CloudAwsGovcloudIntegrationsRds `field:"optional" json:"rds" yaml:"rds"`
	// red_shift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#red_shift CloudAwsGovcloudIntegrations#red_shift}
	RedShift *CloudAwsGovcloudIntegrationsRedShift `field:"optional" json:"redShift" yaml:"redShift"`
	// route53 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#route53 CloudAwsGovcloudIntegrations#route53}
	Route53 *CloudAwsGovcloudIntegrationsRoute53 `field:"optional" json:"route53" yaml:"route53"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#s3 CloudAwsGovcloudIntegrations#s3}
	S3 *CloudAwsGovcloudIntegrationsS3 `field:"optional" json:"s3" yaml:"s3"`
	// sns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#sns CloudAwsGovcloudIntegrations#sns}
	Sns *CloudAwsGovcloudIntegrationsSns `field:"optional" json:"sns" yaml:"sns"`
	// sqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_aws_govcloud_integrations#sqs CloudAwsGovcloudIntegrations#sqs}
	Sqs *CloudAwsGovcloudIntegrationsSqs `field:"optional" json:"sqs" yaml:"sqs"`
}

