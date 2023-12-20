// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboardraw

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OneDashboardRawConfig struct {
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
	// The dashboard's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.28.1/docs/resources/one_dashboard_raw#name OneDashboardRaw#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// page block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.28.1/docs/resources/one_dashboard_raw#page OneDashboardRaw#page}
	Page interface{} `field:"required" json:"page" yaml:"page"`
	// The New Relic account ID where you want to create the dashboard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.28.1/docs/resources/one_dashboard_raw#account_id OneDashboardRaw#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// The dashboard's description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.28.1/docs/resources/one_dashboard_raw#description OneDashboardRaw#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.28.1/docs/resources/one_dashboard_raw#id OneDashboardRaw#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Determines who can see or edit the dashboard. Valid values are private, public_read_only, public_read_write. Defaults to public_read_only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.28.1/docs/resources/one_dashboard_raw#permissions OneDashboardRaw#permissions}
	Permissions *string `field:"optional" json:"permissions" yaml:"permissions"`
}

