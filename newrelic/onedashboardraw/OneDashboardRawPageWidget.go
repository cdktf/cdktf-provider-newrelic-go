// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboardraw


type OneDashboardRawPageWidget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/one_dashboard_raw#column OneDashboardRaw#column}.
	Column *float64 `field:"required" json:"column" yaml:"column"`
	// The configuration of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/one_dashboard_raw#configuration OneDashboardRaw#configuration}
	Configuration *string `field:"required" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/one_dashboard_raw#row OneDashboardRaw#row}.
	Row *float64 `field:"required" json:"row" yaml:"row"`
	// A title for the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/one_dashboard_raw#title OneDashboardRaw#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// The visualization ID of the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/one_dashboard_raw#visualization_id OneDashboardRaw#visualization_id}
	VisualizationId *string `field:"required" json:"visualizationId" yaml:"visualizationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/one_dashboard_raw#height OneDashboardRaw#height}.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// (Optional) Related entity GUIDs. Currently only supports Dashboard entity GUIDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/one_dashboard_raw#linked_entity_guids OneDashboardRaw#linked_entity_guids}
	LinkedEntityGuids *[]*string `field:"optional" json:"linkedEntityGuids" yaml:"linkedEntityGuids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.39.1/docs/resources/one_dashboard_raw#width OneDashboardRaw#width}.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

