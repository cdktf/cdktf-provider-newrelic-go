// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nrqlalertcondition


type NrqlAlertConditionCriticalPrediction struct {
	// BETA PREVIEW: the `predict_by` field is in limited release and only enabled for preview on a per-account basis.
	//
	// - The duration, in seconds, that the prediction should look into the future.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.68.0/docs/resources/nrql_alert_condition#predict_by NrqlAlertCondition#predict_by}
	PredictBy *float64 `field:"optional" json:"predictBy" yaml:"predictBy"`
	// BETA PREVIEW: the `prefer_prediction_violation` field is in limited release and only enabled for preview on a per-account basis.
	//
	// - If a prediction incident is open when a term's static threshold is breached by the actual signal, default behavior is to close the prediction incident and open a static incident. Setting `prefer_prediction_violation` to `true` overrides this behavior leaving the prediction incident open and preventing a static incident from opening.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.68.0/docs/resources/nrql_alert_condition#prefer_prediction_violation NrqlAlertCondition#prefer_prediction_violation}
	PreferPredictionViolation interface{} `field:"optional" json:"preferPredictionViolation" yaml:"preferPredictionViolation"`
}

