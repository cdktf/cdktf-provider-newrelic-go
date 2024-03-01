// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetLogTableNullValues struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.32.0/docs/resources/one_dashboard#null_value OneDashboard#null_value}.
	NullValue *string `field:"optional" json:"nullValue" yaml:"nullValue"`
	// series_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.32.0/docs/resources/one_dashboard#series_overrides OneDashboard#series_overrides}
	SeriesOverrides interface{} `field:"optional" json:"seriesOverrides" yaml:"seriesOverrides"`
}

