package notificationdestination

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
	// (Required) The name of the destination.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#name NotificationDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// (Required) The type of the destination. One of: (WEBHOOK, EMAIL, SERVICE_NOW, PAGERDUTY_ACCOUNT_INTEGRATION, PAGERDUTY_SERVICE_INTEGRATION).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#type NotificationDestination#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// A set of key-value pairs to represent a Notification destination auth.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#auth NotificationDestination#auth}
	Auth *map[string]*string `field:"optional" json:"auth" yaml:"auth"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#id NotificationDestination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_destination#properties NotificationDestination#properties}
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
}

