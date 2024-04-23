// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicelevel


type ServiceLevelEventsBadEventsSelect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.0/docs/resources/service_level#function ServiceLevel#function}.
	Function *string `field:"required" json:"function" yaml:"function"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.0/docs/resources/service_level#attribute ServiceLevel#attribute}.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// The event threshold to use in the SELECT clause.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.35.0/docs/resources/service_level#threshold ServiceLevel#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
}

