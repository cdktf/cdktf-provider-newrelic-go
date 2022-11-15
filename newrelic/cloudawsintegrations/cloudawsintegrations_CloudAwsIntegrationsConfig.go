package cloudawsintegrations

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudAwsIntegrationsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_integrations#linked_account_id CloudAwsIntegrations#linked_account_id}
	LinkedAccountId *float64 `field:"required" json:"linkedAccountId" yaml:"linkedAccountId"`
	// The ID of the account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_integrations#account_id CloudAwsIntegrations#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// billing block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_integrations#billing CloudAwsIntegrations#billing}
	Billing *CloudAwsIntegrationsBilling `field:"optional" json:"billing" yaml:"billing"`
	// cloudtrail block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_integrations#cloudtrail CloudAwsIntegrations#cloudtrail}
	Cloudtrail *CloudAwsIntegrationsCloudtrail `field:"optional" json:"cloudtrail" yaml:"cloudtrail"`
	// doc_db block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_integrations#doc_db CloudAwsIntegrations#doc_db}
	DocDb *CloudAwsIntegrationsDocDb `field:"optional" json:"docDb" yaml:"docDb"`
	// health block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_integrations#health CloudAwsIntegrations#health}
	Health *CloudAwsIntegrationsHealth `field:"optional" json:"health" yaml:"health"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_integrations#id CloudAwsIntegrations#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_integrations#s3 CloudAwsIntegrations#s3}
	S3 *CloudAwsIntegrationsS3 `field:"optional" json:"s3" yaml:"s3"`
	// trusted_advisor block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_integrations#trusted_advisor CloudAwsIntegrations#trusted_advisor}
	TrustedAdvisor *CloudAwsIntegrationsTrustedAdvisor `field:"optional" json:"trustedAdvisor" yaml:"trustedAdvisor"`
	// vpc block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_integrations#vpc CloudAwsIntegrations#vpc}
	Vpc *CloudAwsIntegrationsVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// x_ray block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_aws_integrations#x_ray CloudAwsIntegrations#x_ray}
	XRay *CloudAwsIntegrationsXRay `field:"optional" json:"xRay" yaml:"xRay"`
}

