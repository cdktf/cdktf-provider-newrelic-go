package datanewrelicnotificationdestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataNewrelicNotificationDestinationConfig struct {
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
	// The ID of the destination.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/notification_destination#id DataNewrelicNotificationDestination#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The account ID under which to put the destination.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/notification_destination#account_id DataNewrelicNotificationDestination#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Indicates whether the destination is active.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/notification_destination#active DataNewrelicNotificationDestination#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// auth_basic block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/notification_destination#auth_basic DataNewrelicNotificationDestination#auth_basic}
	AuthBasic *DataNewrelicNotificationDestinationAuthBasic `field:"optional" json:"authBasic" yaml:"authBasic"`
	// auth_token block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/notification_destination#auth_token DataNewrelicNotificationDestination#auth_token}
	AuthToken *DataNewrelicNotificationDestinationAuthToken `field:"optional" json:"authToken" yaml:"authToken"`
	// The name of the destination.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/notification_destination#name DataNewrelicNotificationDestination#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// property block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/notification_destination#property DataNewrelicNotificationDestination#property}
	Property interface{} `field:"optional" json:"property" yaml:"property"`
	// The type of the destination. One of: (WEBHOOK, EMAIL, SERVICE_NOW, PAGERDUTY_ACCOUNT_INTEGRATION, PAGERDUTY_SERVICE_INTEGRATION, JIRA, SLACK, SLACK_COLLABORATION, SLACK_LEGACY, MOBILE_PUSH, EVENT_BRIDGE).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/notification_destination#type DataNewrelicNotificationDestination#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

