// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicelevel


type ServiceLevelEventsBadEvents struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.44.0/docs/resources/service_level#from ServiceLevel#from}.
	From *string `field:"required" json:"from" yaml:"from"`
	// select block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.44.0/docs/resources/service_level#select ServiceLevel#select}
	Select *ServiceLevelEventsBadEventsSelect `field:"optional" json:"select" yaml:"select"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.44.0/docs/resources/service_level#where ServiceLevel#where}.
	Where *string `field:"optional" json:"where" yaml:"where"`
}

