// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetTableThreshold struct {
	// Name of the column in the table, to which the threshold would be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.38.0/docs/resources/one_dashboard#column_name OneDashboard#column_name}
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// The number from which the range starts in thresholds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.38.0/docs/resources/one_dashboard#from OneDashboard#from}
	From *float64 `field:"optional" json:"from" yaml:"from"`
	// Severity of the threshold, which would reflect in the widget, in the range of the threshold specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.38.0/docs/resources/one_dashboard#severity OneDashboard#severity}
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
	// The number at which the range ends in thresholds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.38.0/docs/resources/one_dashboard#to OneDashboard#to}
	To *float64 `field:"optional" json:"to" yaml:"to"`
}

