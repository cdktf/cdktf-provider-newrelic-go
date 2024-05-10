// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationchannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationChannelConfig struct {
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
	// (Required) The id of the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/notification_channel#destination_id NotificationChannel#destination_id}
	DestinationId *string `field:"required" json:"destinationId" yaml:"destinationId"`
	// (Required) The name of the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/notification_channel#name NotificationChannel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// (Required) The type of the channel product. One of: (DISCUSSIONS, ERROR_TRACKING, IINT).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/notification_channel#product NotificationChannel#product}
	Product *string `field:"required" json:"product" yaml:"product"`
	// property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/notification_channel#property NotificationChannel#property}
	Property interface{} `field:"required" json:"property" yaml:"property"`
	// (Required) The type of the channel. One of: (WEBHOOK, EMAIL, SERVICENOW_INCIDENTS, PAGERDUTY_ACCOUNT_INTEGRATION, PAGERDUTY_SERVICE_INTEGRATION, JIRA_CLASSIC, SLACK, SLACK_COLLABORATION, SLACK_LEGACY, MOBILE_PUSH, EVENT_BRIDGE).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/notification_channel#type NotificationChannel#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The account id of the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/notification_channel#account_id NotificationChannel#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Indicates whether the channel is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/notification_channel#active NotificationChannel#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.36.0/docs/resources/notification_channel#id NotificationChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

