// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationdestination


type NotificationDestinationAuthBasic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.75.2/docs/resources/notification_destination#password NotificationDestination#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.75.2/docs/resources/notification_destination#user NotificationDestination#user}.
	User *string `field:"required" json:"user" yaml:"user"`
}

