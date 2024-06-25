// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicelevel


type ServiceLevelObjective struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.38.1/docs/resources/service_level#target ServiceLevel#target}.
	Target *float64 `field:"required" json:"target" yaml:"target"`
	// time_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.38.1/docs/resources/service_level#time_window ServiceLevel#time_window}
	TimeWindow *ServiceLevelObjectiveTimeWindow `field:"required" json:"timeWindow" yaml:"timeWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.38.1/docs/resources/service_level#description ServiceLevel#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.38.1/docs/resources/service_level#name ServiceLevel#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

