package cloudawsintegrations

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudAwsIntegrationsConfig struct {
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
	// The ID of the linked AWS account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#linked_account_id CloudAwsIntegrations#linked_account_id}
	LinkedAccountId *float64 `field:"required" json:"linkedAccountId" yaml:"linkedAccountId"`
	// The ID of the account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#account_id CloudAwsIntegrations#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// alb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#alb CloudAwsIntegrations#alb}
	Alb *CloudAwsIntegrationsAlb `field:"optional" json:"alb" yaml:"alb"`
	// api_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#api_gateway CloudAwsIntegrations#api_gateway}
	ApiGateway *CloudAwsIntegrationsApiGateway `field:"optional" json:"apiGateway" yaml:"apiGateway"`
	// auto_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#auto_scaling CloudAwsIntegrations#auto_scaling}
	AutoScaling *CloudAwsIntegrationsAutoScaling `field:"optional" json:"autoScaling" yaml:"autoScaling"`
	// aws_app_sync block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#aws_app_sync CloudAwsIntegrations#aws_app_sync}
	AwsAppSync *CloudAwsIntegrationsAwsAppSync `field:"optional" json:"awsAppSync" yaml:"awsAppSync"`
	// aws_athena block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#aws_athena CloudAwsIntegrations#aws_athena}
	AwsAthena *CloudAwsIntegrationsAwsAthena `field:"optional" json:"awsAthena" yaml:"awsAthena"`
	// aws_cognito block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#aws_cognito CloudAwsIntegrations#aws_cognito}
	AwsCognito *CloudAwsIntegrationsAwsCognito `field:"optional" json:"awsCognito" yaml:"awsCognito"`
	// aws_connect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#aws_connect CloudAwsIntegrations#aws_connect}
	AwsConnect *CloudAwsIntegrationsAwsConnect `field:"optional" json:"awsConnect" yaml:"awsConnect"`
	// aws_direct_connect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#aws_direct_connect CloudAwsIntegrations#aws_direct_connect}
	AwsDirectConnect *CloudAwsIntegrationsAwsDirectConnect `field:"optional" json:"awsDirectConnect" yaml:"awsDirectConnect"`
	// aws_fsx block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#aws_fsx CloudAwsIntegrations#aws_fsx}
	AwsFsx *CloudAwsIntegrationsAwsFsx `field:"optional" json:"awsFsx" yaml:"awsFsx"`
	// billing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#billing CloudAwsIntegrations#billing}
	Billing *CloudAwsIntegrationsBilling `field:"optional" json:"billing" yaml:"billing"`
	// cloudtrail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#cloudtrail CloudAwsIntegrations#cloudtrail}
	Cloudtrail *CloudAwsIntegrationsCloudtrail `field:"optional" json:"cloudtrail" yaml:"cloudtrail"`
	// doc_db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#doc_db CloudAwsIntegrations#doc_db}
	DocDb *CloudAwsIntegrationsDocDb `field:"optional" json:"docDb" yaml:"docDb"`
	// ebs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#ebs CloudAwsIntegrations#ebs}
	Ebs *CloudAwsIntegrationsEbs `field:"optional" json:"ebs" yaml:"ebs"`
	// elasticache block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#elasticache CloudAwsIntegrations#elasticache}
	Elasticache *CloudAwsIntegrationsElasticache `field:"optional" json:"elasticache" yaml:"elasticache"`
	// health block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#health CloudAwsIntegrations#health}
	Health *CloudAwsIntegrationsHealth `field:"optional" json:"health" yaml:"health"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#id CloudAwsIntegrations#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#s3 CloudAwsIntegrations#s3}
	S3 *CloudAwsIntegrationsS3 `field:"optional" json:"s3" yaml:"s3"`
	// sqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#sqs CloudAwsIntegrations#sqs}
	Sqs *CloudAwsIntegrationsSqs `field:"optional" json:"sqs" yaml:"sqs"`
	// trusted_advisor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#trusted_advisor CloudAwsIntegrations#trusted_advisor}
	TrustedAdvisor *CloudAwsIntegrationsTrustedAdvisor `field:"optional" json:"trustedAdvisor" yaml:"trustedAdvisor"`
	// vpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#vpc CloudAwsIntegrations#vpc}
	Vpc *CloudAwsIntegrationsVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// x_ray block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations#x_ray CloudAwsIntegrations#x_ray}
	XRay *CloudAwsIntegrationsXRay `field:"optional" json:"xRay" yaml:"xRay"`
}

