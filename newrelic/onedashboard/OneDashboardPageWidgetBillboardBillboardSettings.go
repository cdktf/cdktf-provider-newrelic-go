// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetBillboardBillboardSettings struct {
	// grid_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/one_dashboard#grid_options OneDashboard#grid_options}
	GridOptions *OneDashboardPageWidgetBillboardBillboardSettingsGridOptions `field:"optional" json:"gridOptions" yaml:"gridOptions"`
	// link block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/one_dashboard#link OneDashboard#link}
	Link *OneDashboardPageWidgetBillboardBillboardSettingsLink `field:"optional" json:"link" yaml:"link"`
	// visual block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.4/docs/resources/one_dashboard#visual OneDashboard#visual}
	Visual *OneDashboardPageWidgetBillboardBillboardSettingsVisual `field:"optional" json:"visual" yaml:"visual"`
}

