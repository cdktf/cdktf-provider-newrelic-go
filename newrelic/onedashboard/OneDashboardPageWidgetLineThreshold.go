// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard


type OneDashboardPageWidgetLineThreshold struct {
	// The number from which the range starts in thresholds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.57.0/docs/resources/one_dashboard#from OneDashboard#from}
	From *string `field:"optional" json:"from" yaml:"from"`
	// Name of the threshold created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.57.0/docs/resources/one_dashboard#name OneDashboard#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Severity of the threshold, which would reflect in the widget, in the range of the threshold specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.57.0/docs/resources/one_dashboard#severity OneDashboard#severity}
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
	// The number at which the range ends in thresholds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.57.0/docs/resources/one_dashboard#to OneDashboard#to}
	To *string `field:"optional" json:"to" yaml:"to"`
}

