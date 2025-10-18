// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPage struct {
	// The dashboard page's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#name OneDashboard#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The dashboard page's description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#description OneDashboard#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// widget_area block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#widget_area OneDashboard#widget_area}
	WidgetArea interface{} `field:"optional" json:"widgetArea" yaml:"widgetArea"`
	// widget_bar block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#widget_bar OneDashboard#widget_bar}
	WidgetBar interface{} `field:"optional" json:"widgetBar" yaml:"widgetBar"`
	// widget_billboard block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#widget_billboard OneDashboard#widget_billboard}
	WidgetBillboard interface{} `field:"optional" json:"widgetBillboard" yaml:"widgetBillboard"`
	// widget_bullet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#widget_bullet OneDashboard#widget_bullet}
	WidgetBullet interface{} `field:"optional" json:"widgetBullet" yaml:"widgetBullet"`
	// widget_funnel block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#widget_funnel OneDashboard#widget_funnel}
	WidgetFunnel interface{} `field:"optional" json:"widgetFunnel" yaml:"widgetFunnel"`
	// widget_heatmap block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#widget_heatmap OneDashboard#widget_heatmap}
	WidgetHeatmap interface{} `field:"optional" json:"widgetHeatmap" yaml:"widgetHeatmap"`
	// widget_histogram block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#widget_histogram OneDashboard#widget_histogram}
	WidgetHistogram interface{} `field:"optional" json:"widgetHistogram" yaml:"widgetHistogram"`
	// widget_json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#widget_json OneDashboard#widget_json}
	WidgetJson interface{} `field:"optional" json:"widgetJson" yaml:"widgetJson"`
	// widget_line block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#widget_line OneDashboard#widget_line}
	WidgetLine interface{} `field:"optional" json:"widgetLine" yaml:"widgetLine"`
	// widget_log_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#widget_log_table OneDashboard#widget_log_table}
	WidgetLogTable interface{} `field:"optional" json:"widgetLogTable" yaml:"widgetLogTable"`
	// widget_markdown block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#widget_markdown OneDashboard#widget_markdown}
	WidgetMarkdown interface{} `field:"optional" json:"widgetMarkdown" yaml:"widgetMarkdown"`
	// widget_pie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#widget_pie OneDashboard#widget_pie}
	WidgetPie interface{} `field:"optional" json:"widgetPie" yaml:"widgetPie"`
	// widget_stacked_bar block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#widget_stacked_bar OneDashboard#widget_stacked_bar}
	WidgetStackedBar interface{} `field:"optional" json:"widgetStackedBar" yaml:"widgetStackedBar"`
	// widget_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.3/docs/resources/one_dashboard#widget_table OneDashboard#widget_table}
	WidgetTable interface{} `field:"optional" json:"widgetTable" yaml:"widgetTable"`
}

