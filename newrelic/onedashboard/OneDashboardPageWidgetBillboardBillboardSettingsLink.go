// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetBillboardBillboardSettingsLink struct {
	// Whether to open the link in a new tab.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/one_dashboard#new_tab OneDashboard#new_tab}
	NewTab interface{} `field:"optional" json:"newTab" yaml:"newTab"`
	// The title for the link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/one_dashboard#title OneDashboard#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// The URL for the link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/one_dashboard#url OneDashboard#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

