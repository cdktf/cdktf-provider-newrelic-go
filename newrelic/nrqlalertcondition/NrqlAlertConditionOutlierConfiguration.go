// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nrqlalertcondition


type NrqlAlertConditionOutlierConfiguration struct {
	// dbscan block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.3/docs/resources/nrql_alert_condition#dbscan NrqlAlertCondition#dbscan}
	Dbscan *NrqlAlertConditionOutlierConfigurationDbscan `field:"required" json:"dbscan" yaml:"dbscan"`
}

