// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsmultilocationalertcondition


type SyntheticsMultilocationAlertConditionWarning struct {
	// The minimum number of monitor locations that must be concurrently failing before an incident is opened.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.40.0/docs/resources/synthetics_multilocation_alert_condition#threshold SyntheticsMultilocationAlertCondition#threshold}
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
}

