// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetHeatmapUnitsSeriesOverrides struct {
	// Series name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.75.2/docs/resources/one_dashboard#series_name OneDashboard#series_name}
	SeriesName *string `field:"optional" json:"seriesName" yaml:"seriesName"`
	// Unit name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.75.2/docs/resources/one_dashboard#unit OneDashboard#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

