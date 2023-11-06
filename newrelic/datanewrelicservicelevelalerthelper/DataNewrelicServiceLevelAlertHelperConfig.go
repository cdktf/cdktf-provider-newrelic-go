// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datanewrelicservicelevelalerthelper

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataNewrelicServiceLevelAlertHelperConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.6/docs/data-sources/service_level_alert_helper#alert_type DataNewrelicServiceLevelAlertHelper#alert_type}.
	AlertType *string `field:"required" json:"alertType" yaml:"alertType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.6/docs/data-sources/service_level_alert_helper#sli_guid DataNewrelicServiceLevelAlertHelper#sli_guid}.
	SliGuid *string `field:"required" json:"sliGuid" yaml:"sliGuid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.6/docs/data-sources/service_level_alert_helper#slo_period DataNewrelicServiceLevelAlertHelper#slo_period}.
	SloPeriod *float64 `field:"required" json:"sloPeriod" yaml:"sloPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.6/docs/data-sources/service_level_alert_helper#slo_target DataNewrelicServiceLevelAlertHelper#slo_target}.
	SloTarget *float64 `field:"required" json:"sloTarget" yaml:"sloTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.6/docs/data-sources/service_level_alert_helper#custom_evaluation_period DataNewrelicServiceLevelAlertHelper#custom_evaluation_period}.
	CustomEvaluationPeriod *float64 `field:"optional" json:"customEvaluationPeriod" yaml:"customEvaluationPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.6/docs/data-sources/service_level_alert_helper#custom_tolerated_budget_consumption DataNewrelicServiceLevelAlertHelper#custom_tolerated_budget_consumption}.
	CustomToleratedBudgetConsumption *float64 `field:"optional" json:"customToleratedBudgetConsumption" yaml:"customToleratedBudgetConsumption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.6/docs/data-sources/service_level_alert_helper#id DataNewrelicServiceLevelAlertHelper#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.6/docs/data-sources/service_level_alert_helper#is_bad_events DataNewrelicServiceLevelAlertHelper#is_bad_events}.
	IsBadEvents interface{} `field:"optional" json:"isBadEvents" yaml:"isBadEvents"`
}

