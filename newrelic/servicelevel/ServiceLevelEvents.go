// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicelevel


type ServiceLevelEvents struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.64.0/docs/resources/service_level#account_id ServiceLevel#account_id}.
	AccountId *float64 `field:"required" json:"accountId" yaml:"accountId"`
	// valid_events block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.64.0/docs/resources/service_level#valid_events ServiceLevel#valid_events}
	ValidEvents *ServiceLevelEventsValidEvents `field:"required" json:"validEvents" yaml:"validEvents"`
	// bad_events block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.64.0/docs/resources/service_level#bad_events ServiceLevel#bad_events}
	BadEvents *ServiceLevelEventsBadEvents `field:"optional" json:"badEvents" yaml:"badEvents"`
	// good_events block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.64.0/docs/resources/service_level#good_events ServiceLevel#good_events}
	GoodEvents *ServiceLevelEventsGoodEvents `field:"optional" json:"goodEvents" yaml:"goodEvents"`
}

