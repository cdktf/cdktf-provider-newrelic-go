// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationdestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationDestinationConfig struct {
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
	// (Required) The name of the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.68.0/docs/resources/notification_destination#name NotificationDestination#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.68.0/docs/resources/notification_destination#property NotificationDestination#property}
	Property interface{} `field:"required" json:"property" yaml:"property"`
	// (Required) The type of the destination.
	//
	// One of: (WEBHOOK, EMAIL, SERVICE_NOW, SERVICE_NOW_APP, PAGERDUTY_ACCOUNT_INTEGRATION, PAGERDUTY_SERVICE_INTEGRATION, JIRA, SLACK, SLACK_COLLABORATION, SLACK_LEGACY, MOBILE_PUSH, EVENT_BRIDGE, MICROSOFT_TEAMS, WORKFLOW_AUTOMATION).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.68.0/docs/resources/notification_destination#type NotificationDestination#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The account ID under which to put the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.68.0/docs/resources/notification_destination#account_id NotificationDestination#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Indicates whether the destination is active.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.68.0/docs/resources/notification_destination#active NotificationDestination#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// auth_basic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.68.0/docs/resources/notification_destination#auth_basic NotificationDestination#auth_basic}
	AuthBasic *NotificationDestinationAuthBasic `field:"optional" json:"authBasic" yaml:"authBasic"`
	// auth_custom_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.68.0/docs/resources/notification_destination#auth_custom_header NotificationDestination#auth_custom_header}
	AuthCustomHeader interface{} `field:"optional" json:"authCustomHeader" yaml:"authCustomHeader"`
	// auth_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.68.0/docs/resources/notification_destination#auth_token NotificationDestination#auth_token}
	AuthToken *NotificationDestinationAuthToken `field:"optional" json:"authToken" yaml:"authToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.68.0/docs/resources/notification_destination#id NotificationDestination#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// secure_url block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.68.0/docs/resources/notification_destination#secure_url NotificationDestination#secure_url}
	SecureUrl *NotificationDestinationSecureUrl `field:"optional" json:"secureUrl" yaml:"secureUrl"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.68.0/docs/resources/notification_destination#timeouts NotificationDestination#timeouts}
	Timeouts *NotificationDestinationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

