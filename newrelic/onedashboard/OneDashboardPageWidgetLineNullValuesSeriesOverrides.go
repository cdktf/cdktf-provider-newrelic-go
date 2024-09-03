// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetLineNullValuesSeriesOverrides struct {
	// Null value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.44.0/docs/resources/one_dashboard#null_value OneDashboard#null_value}
	NullValue *string `field:"optional" json:"nullValue" yaml:"nullValue"`
	// Series name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.44.0/docs/resources/one_dashboard#series_name OneDashboard#series_name}
	SeriesName *string `field:"optional" json:"seriesName" yaml:"seriesName"`
}

