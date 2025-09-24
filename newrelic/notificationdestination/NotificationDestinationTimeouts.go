// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationdestination


type NotificationDestinationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.69.1/docs/resources/notification_destination#create NotificationDestination#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.69.1/docs/resources/notification_destination#update NotificationDestination#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

