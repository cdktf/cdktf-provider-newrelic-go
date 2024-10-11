// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationchannel


type NotificationChannelTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.50.0/docs/resources/notification_channel#create NotificationChannel#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.50.0/docs/resources/notification_channel#update NotificationChannel#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

