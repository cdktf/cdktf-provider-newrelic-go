// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetBillboardBillboardSettingsVisual struct {
	// Billboard alignment type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.3/docs/resources/one_dashboard#alignment OneDashboard#alignment}
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// Billboard display type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.3/docs/resources/one_dashboard#display OneDashboard#display}
	Display *string `field:"optional" json:"display" yaml:"display"`
}

