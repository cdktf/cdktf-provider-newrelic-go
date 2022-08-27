// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationDestinationConfig struct {
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
	// The account ID under which to put the destination.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#account_id NotificationDestination#account_id}
	AccountId *float64 `field:"required" json:"accountId" yaml:"accountId"`
	// (Required) The name of the destination.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#name NotificationDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// property block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#property NotificationDestination#property}
	Property interface{} `field:"required" json:"property" yaml:"property"`
	// (Required) The type of the destination. One of: (WEBHOOK, EMAIL, SERVICE_NOW, PAGERDUTY_ACCOUNT_INTEGRATION, PAGERDUTY_SERVICE_INTEGRATION, JIRA).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#type NotificationDestination#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Indicates whether the destination is active.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#active NotificationDestination#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// auth_basic block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#auth_basic NotificationDestination#auth_basic}
	AuthBasic *NotificationDestinationAuthBasic `field:"optional" json:"authBasic" yaml:"authBasic"`
	// auth_token block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#auth_token NotificationDestination#auth_token}
	AuthToken *NotificationDestinationAuthToken `field:"optional" json:"authToken" yaml:"authToken"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#id NotificationDestination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

