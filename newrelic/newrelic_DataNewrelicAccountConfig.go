// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataNewrelicAccountConfig struct {
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
	// The ID of the account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/account#account_id DataNewrelicAccount#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/account#id DataNewrelicAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/account#name DataNewrelicAccount#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The scope of the account in New Relic.  Valid values are "global" and "in_region".  Defaults to "in_region".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/account#scope DataNewrelicAccount#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
}
