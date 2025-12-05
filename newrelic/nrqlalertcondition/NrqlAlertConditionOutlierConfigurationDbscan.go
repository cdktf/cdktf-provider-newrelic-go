// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nrqlalertcondition


type NrqlAlertConditionOutlierConfigurationDbscan struct {
	// BETA PREVIEW: the `epsilon` field is in limited release and only enabled for preview on a per-account basis.
	//
	// - Radius (distance threshold) for DBSCAN in the units of the query result. Smaller values tighten clusters; larger values broaden them. Must be > 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.3/docs/resources/nrql_alert_condition#epsilon NrqlAlertCondition#epsilon}
	Epsilon *float64 `field:"required" json:"epsilon" yaml:"epsilon"`
	// BETA PREVIEW: the `minimum_points` field is in limited release and only enabled for preview on a per-account basis.
	//
	// - Minimum number of neighboring points needed to form a cluster. Must be >= 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.3/docs/resources/nrql_alert_condition#minimum_points NrqlAlertCondition#minimum_points}
	MinimumPoints *float64 `field:"required" json:"minimumPoints" yaml:"minimumPoints"`
	// BETA PREVIEW: the `evaluation_group_facet` field is in limited release and only enabled for preview on a per-account basis.
	//
	// - Optional NRQL facet attribute used to segment data into groups (e.g. `host`, `region`) before running outlier detection. Omit to evaluate all results together.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.3/docs/resources/nrql_alert_condition#evaluation_group_facet NrqlAlertCondition#evaluation_group_facet}
	EvaluationGroupFacet *string `field:"optional" json:"evaluationGroupFacet" yaml:"evaluationGroupFacet"`
}

