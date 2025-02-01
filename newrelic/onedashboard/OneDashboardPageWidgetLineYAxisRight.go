// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetLineYAxisRight struct {
	// Minimum value of the range to be specified with the Y-Axis on the right of the line widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/one_dashboard#y_axis_right_max OneDashboard#y_axis_right_max}
	YAxisRightMax *float64 `field:"optional" json:"yAxisRightMax" yaml:"yAxisRightMax"`
	// Minimum value of the range to be specified with the Y-Axis on the right of the line widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/one_dashboard#y_axis_right_min OneDashboard#y_axis_right_min}
	YAxisRightMin *float64 `field:"optional" json:"yAxisRightMin" yaml:"yAxisRightMin"`
	// A set of series that helps specify the Y-Axis on the right of the line widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/one_dashboard#y_axis_right_series OneDashboard#y_axis_right_series}
	YAxisRightSeries *[]*string `field:"optional" json:"yAxisRightSeries" yaml:"yAxisRightSeries"`
	// An attribute that helps specify the Y-Axis on the right of the line widget.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/one_dashboard#y_axis_right_zero OneDashboard#y_axis_right_zero}
	YAxisRightZero interface{} `field:"optional" json:"yAxisRightZero" yaml:"yAxisRightZero"`
}

