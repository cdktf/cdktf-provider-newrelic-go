package notificationchannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationChannelConfig struct {
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
	// (Required) The id of the destination.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_channel#destination_id NotificationChannel#destination_id}
	DestinationId *string `field:"required" json:"destinationId" yaml:"destinationId"`
	// (Required) The name of the channel.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_channel#name NotificationChannel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// (Required) The type of the channel product. One of: (ALERTS, DISCUSSIONS, ERROR_TRACKING, NTFC, SHARING, PD, IINT).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_channel#product NotificationChannel#product}
	Product *string `field:"required" json:"product" yaml:"product"`
	// (Required) The type of the channel. One of: (WEBHOOK, EMAIL, SERVICENOW_INCIDENTS, PAGERDUTY_ACCOUNT_INTEGRATION, PAGERDUTY_SERVICE_INTEGRATION).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_channel#type NotificationChannel#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_channel#id NotificationChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/notification_channel#properties NotificationChannel#properties}
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
}
