// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetTableColorsSeriesOverrides struct {
	// Color code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.69.1/docs/resources/one_dashboard#color OneDashboard#color}
	Color *string `field:"optional" json:"color" yaml:"color"`
	// Series name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.69.1/docs/resources/one_dashboard#series_name OneDashboard#series_name}
	SeriesName *string `field:"optional" json:"seriesName" yaml:"seriesName"`
}

