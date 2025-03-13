// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datanewrelicnotificationdestination

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataNewrelicNotificationDestinationConfig struct {
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
	// The account ID under which to put the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/data-sources/notification_destination#account_id DataNewrelicNotificationDestination#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// The ID of the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/data-sources/notification_destination#id DataNewrelicNotificationDestination#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/data-sources/notification_destination#name DataNewrelicNotificationDestination#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// secure_url block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/data-sources/notification_destination#secure_url DataNewrelicNotificationDestination#secure_url}
	SecureUrl interface{} `field:"optional" json:"secureUrl" yaml:"secureUrl"`
}

