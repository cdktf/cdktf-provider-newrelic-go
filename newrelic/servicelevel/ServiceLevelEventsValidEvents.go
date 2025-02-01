// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicelevel


type ServiceLevelEventsValidEvents struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/service_level#from ServiceLevel#from}.
	From *string `field:"required" json:"from" yaml:"from"`
	// select block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/service_level#select ServiceLevel#select}
	Select *ServiceLevelEventsValidEventsSelect `field:"optional" json:"select" yaml:"select"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/service_level#where ServiceLevel#where}.
	Where *string `field:"optional" json:"where" yaml:"where"`
}

