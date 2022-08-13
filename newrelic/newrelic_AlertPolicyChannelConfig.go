// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlertPolicyChannelConfig struct {
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
	// Array of channel IDs to apply to the specified policy.
	//
	// We recommended sorting channel IDs in ascending order to avoid drift your Terraform state.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/alert_policy_channel#channel_ids AlertPolicyChannel#channel_ids}
	ChannelIds *[]*float64 `field:"required" json:"channelIds" yaml:"channelIds"`
	// The ID of the policy.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/alert_policy_channel#policy_id AlertPolicyChannel#policy_id}
	PolicyId *float64 `field:"required" json:"policyId" yaml:"policyId"`
	// The New Relic account ID where you want to link the channel to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/alert_policy_channel#account_id AlertPolicyChannel#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/alert_policy_channel#id AlertPolicyChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

