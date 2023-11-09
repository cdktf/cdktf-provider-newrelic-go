// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datanewrelicalertchannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataNewrelicAlertChannelConfig struct {
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
	// The name of the alert channel in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/data-sources/alert_channel#name DataNewrelicAlertChannel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The New Relic account ID where you want to retrieve the alert channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/data-sources/alert_channel#account_id DataNewrelicAlertChannel#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.7/docs/data-sources/alert_channel#id DataNewrelicAlertChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

