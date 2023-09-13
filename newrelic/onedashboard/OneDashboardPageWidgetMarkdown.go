// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetMarkdown struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#column OneDashboard#column}.
	Column *float64 `field:"required" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#row OneDashboard#row}.
	Row *float64 `field:"required" json:"row" yaml:"row"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#text OneDashboard#text}.
	Text *string `field:"required" json:"text" yaml:"text"`
	// A title for the widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#title OneDashboard#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// colors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#colors OneDashboard#colors}
	Colors interface{} `field:"optional" json:"colors" yaml:"colors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#facet_show_other_series OneDashboard#facet_show_other_series}.
	FacetShowOtherSeries interface{} `field:"optional" json:"facetShowOtherSeries" yaml:"facetShowOtherSeries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#height OneDashboard#height}.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#ignore_time_range OneDashboard#ignore_time_range}.
	IgnoreTimeRange interface{} `field:"optional" json:"ignoreTimeRange" yaml:"ignoreTimeRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#legend_enabled OneDashboard#legend_enabled}.
	LegendEnabled interface{} `field:"optional" json:"legendEnabled" yaml:"legendEnabled"`
	// null_values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#null_values OneDashboard#null_values}
	NullValues interface{} `field:"optional" json:"nullValues" yaml:"nullValues"`
	// units block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#units OneDashboard#units}
	Units interface{} `field:"optional" json:"units" yaml:"units"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#width OneDashboard#width}.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#y_axis_left_max OneDashboard#y_axis_left_max}.
	YAxisLeftMax *float64 `field:"optional" json:"yAxisLeftMax" yaml:"yAxisLeftMax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/one_dashboard#y_axis_left_min OneDashboard#y_axis_left_min}.
	YAxisLeftMin *float64 `field:"optional" json:"yAxisLeftMin" yaml:"yAxisLeftMin"`
}

