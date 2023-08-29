// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datanewrelicservicelevelalerthelper

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.dataNewrelicServiceLevelAlertHelper.DataNewrelicServiceLevelAlertHelper",
		reflect.TypeOf((*DataNewrelicServiceLevelAlertHelper)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alertType", GoGetter: "AlertType"},
			_jsii_.MemberProperty{JsiiProperty: "alertTypeInput", GoGetter: "AlertTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customEvaluationPeriod", GoGetter: "CustomEvaluationPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "customEvaluationPeriodInput", GoGetter: "CustomEvaluationPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "customToleratedBudgetConsumption", GoGetter: "CustomToleratedBudgetConsumption"},
			_jsii_.MemberProperty{JsiiProperty: "customToleratedBudgetConsumptionInput", GoGetter: "CustomToleratedBudgetConsumptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationPeriod", GoGetter: "EvaluationPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isBadEvents", GoGetter: "IsBadEvents"},
			_jsii_.MemberProperty{JsiiProperty: "isBadEventsInput", GoGetter: "IsBadEventsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nrql", GoGetter: "Nrql"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomEvaluationPeriod", GoMethod: "ResetCustomEvaluationPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomToleratedBudgetConsumption", GoMethod: "ResetCustomToleratedBudgetConsumption"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsBadEvents", GoMethod: "ResetIsBadEvents"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "sliGuid", GoGetter: "SliGuid"},
			_jsii_.MemberProperty{JsiiProperty: "sliGuidInput", GoGetter: "SliGuidInput"},
			_jsii_.MemberProperty{JsiiProperty: "sloPeriod", GoGetter: "SloPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "sloPeriodInput", GoGetter: "SloPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "sloTarget", GoGetter: "SloTarget"},
			_jsii_.MemberProperty{JsiiProperty: "sloTargetInput", GoGetter: "SloTargetInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
			_jsii_.MemberProperty{JsiiProperty: "toleratedBudgetConsumption", GoGetter: "ToleratedBudgetConsumption"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataNewrelicServiceLevelAlertHelper{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-newrelic.dataNewrelicServiceLevelAlertHelper.DataNewrelicServiceLevelAlertHelperConfig",
		reflect.TypeOf((*DataNewrelicServiceLevelAlertHelperConfig)(nil)).Elem(),
	)
}
