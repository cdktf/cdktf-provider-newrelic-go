// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertmutingrule


type AlertMutingRuleConditionConditions struct {
	// The attribute on an incident.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.74.0/docs/resources/alert_muting_rule#attribute AlertMutingRule#attribute}
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The operator used to compare the attribute's value with the supplied value(s).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.74.0/docs/resources/alert_muting_rule#operator AlertMutingRule#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The value(s) to compare against the attribute's value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.74.0/docs/resources/alert_muting_rule#values AlertMutingRule#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

