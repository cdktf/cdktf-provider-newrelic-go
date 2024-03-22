// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationchannel


type NotificationChannelProperty struct {
	// Notification property key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.33.0/docs/resources/notification_channel#key NotificationChannel#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// Notification property value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.33.0/docs/resources/notification_channel#value NotificationChannel#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Notification property display key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.33.0/docs/resources/notification_channel#display_value NotificationChannel#display_value}
	DisplayValue *string `field:"optional" json:"displayValue" yaml:"displayValue"`
	// Notification property label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.33.0/docs/resources/notification_channel#label NotificationChannel#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
}

