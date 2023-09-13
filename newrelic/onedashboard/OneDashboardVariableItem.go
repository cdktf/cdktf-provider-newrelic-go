// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardVariableItem struct {
	// A possible variable value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#value OneDashboard#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// A human-friendly display string for this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#title OneDashboard#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

