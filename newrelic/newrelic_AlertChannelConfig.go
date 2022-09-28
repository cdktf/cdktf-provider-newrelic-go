// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlertChannelConfig struct {
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
	// (Required) The name of the channel.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/alert_channel#name AlertChannel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// (Required) The type of channel. One of: (opsgenie, pagerduty, slack, user, victorops, webhook, email).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/alert_channel#type AlertChannel#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The New Relic account ID where you want to create alert channels.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/alert_channel#account_id AlertChannel#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/alert_channel#config AlertChannel#config}
	Config *AlertChannelConfigA `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/alert_channel#id AlertChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

