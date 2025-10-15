// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetBillboardBillboardSettingsGridOptions struct {
	// Number of columns in the grid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.0/docs/resources/one_dashboard#columns OneDashboard#columns}
	Columns *float64 `field:"optional" json:"columns" yaml:"columns"`
	// Grid label setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.0/docs/resources/one_dashboard#label OneDashboard#label}
	Label *float64 `field:"optional" json:"label" yaml:"label"`
	// Grid value setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.0/docs/resources/one_dashboard#value OneDashboard#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

