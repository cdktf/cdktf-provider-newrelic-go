// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nrqlalertcondition

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertCondition",
		reflect.TypeOf((*NrqlAlertCondition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "aggregationDelay", GoGetter: "AggregationDelay"},
			_jsii_.MemberProperty{JsiiProperty: "aggregationDelayInput", GoGetter: "AggregationDelayInput"},
			_jsii_.MemberProperty{JsiiProperty: "aggregationMethod", GoGetter: "AggregationMethod"},
			_jsii_.MemberProperty{JsiiProperty: "aggregationMethodInput", GoGetter: "AggregationMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "aggregationTimer", GoGetter: "AggregationTimer"},
			_jsii_.MemberProperty{JsiiProperty: "aggregationTimerInput", GoGetter: "AggregationTimerInput"},
			_jsii_.MemberProperty{JsiiProperty: "aggregationWindow", GoGetter: "AggregationWindow"},
			_jsii_.MemberProperty{JsiiProperty: "aggregationWindowInput", GoGetter: "AggregationWindowInput"},
			_jsii_.MemberProperty{JsiiProperty: "baselineDirection", GoGetter: "BaselineDirection"},
			_jsii_.MemberProperty{JsiiProperty: "baselineDirectionInput", GoGetter: "BaselineDirectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "closeViolationsOnExpiration", GoGetter: "CloseViolationsOnExpiration"},
			_jsii_.MemberProperty{JsiiProperty: "closeViolationsOnExpirationInput", GoGetter: "CloseViolationsOnExpirationInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "critical", GoGetter: "Critical"},
			_jsii_.MemberProperty{JsiiProperty: "criticalInput", GoGetter: "CriticalInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "entityGuid", GoGetter: "EntityGuid"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationDelay", GoGetter: "EvaluationDelay"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationDelayInput", GoGetter: "EvaluationDelayInput"},
			_jsii_.MemberProperty{JsiiProperty: "expirationDuration", GoGetter: "ExpirationDuration"},
			_jsii_.MemberProperty{JsiiProperty: "expirationDurationInput", GoGetter: "ExpirationDurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "fillOption", GoGetter: "FillOption"},
			_jsii_.MemberProperty{JsiiProperty: "fillOptionInput", GoGetter: "FillOptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fillValue", GoGetter: "FillValue"},
			_jsii_.MemberProperty{JsiiProperty: "fillValueInput", GoGetter: "FillValueInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nrql", GoGetter: "Nrql"},
			_jsii_.MemberProperty{JsiiProperty: "nrqlInput", GoGetter: "NrqlInput"},
			_jsii_.MemberProperty{JsiiProperty: "openViolationOnExpiration", GoGetter: "OpenViolationOnExpiration"},
			_jsii_.MemberProperty{JsiiProperty: "openViolationOnExpirationInput", GoGetter: "OpenViolationOnExpirationInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyId", GoGetter: "PolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "policyIdInput", GoGetter: "PolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putCritical", GoMethod: "PutCritical"},
			_jsii_.MemberMethod{JsiiMethod: "putNrql", GoMethod: "PutNrql"},
			_jsii_.MemberMethod{JsiiMethod: "putTerm", GoMethod: "PutTerm"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "putWarning", GoMethod: "PutWarning"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountId", GoMethod: "ResetAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAggregationDelay", GoMethod: "ResetAggregationDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetAggregationMethod", GoMethod: "ResetAggregationMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetAggregationTimer", GoMethod: "ResetAggregationTimer"},
			_jsii_.MemberMethod{JsiiMethod: "resetAggregationWindow", GoMethod: "ResetAggregationWindow"},
			_jsii_.MemberMethod{JsiiMethod: "resetBaselineDirection", GoMethod: "ResetBaselineDirection"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloseViolationsOnExpiration", GoMethod: "ResetCloseViolationsOnExpiration"},
			_jsii_.MemberMethod{JsiiMethod: "resetCritical", GoMethod: "ResetCritical"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetEvaluationDelay", GoMethod: "ResetEvaluationDelay"},
			_jsii_.MemberMethod{JsiiMethod: "resetExpirationDuration", GoMethod: "ResetExpirationDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetFillOption", GoMethod: "ResetFillOption"},
			_jsii_.MemberMethod{JsiiMethod: "resetFillValue", GoMethod: "ResetFillValue"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOpenViolationOnExpiration", GoMethod: "ResetOpenViolationOnExpiration"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRunbookUrl", GoMethod: "ResetRunbookUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetSlideBy", GoMethod: "ResetSlideBy"},
			_jsii_.MemberMethod{JsiiMethod: "resetTerm", GoMethod: "ResetTerm"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resetViolationTimeLimit", GoMethod: "ResetViolationTimeLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetViolationTimeLimitSeconds", GoMethod: "ResetViolationTimeLimitSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetWarning", GoMethod: "ResetWarning"},
			_jsii_.MemberProperty{JsiiProperty: "runbookUrl", GoGetter: "RunbookUrl"},
			_jsii_.MemberProperty{JsiiProperty: "runbookUrlInput", GoGetter: "RunbookUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "slideBy", GoGetter: "SlideBy"},
			_jsii_.MemberProperty{JsiiProperty: "slideByInput", GoGetter: "SlideByInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "term", GoGetter: "Term"},
			_jsii_.MemberProperty{JsiiProperty: "termInput", GoGetter: "TermInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "violationTimeLimit", GoGetter: "ViolationTimeLimit"},
			_jsii_.MemberProperty{JsiiProperty: "violationTimeLimitInput", GoGetter: "ViolationTimeLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "violationTimeLimitSeconds", GoGetter: "ViolationTimeLimitSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "violationTimeLimitSecondsInput", GoGetter: "ViolationTimeLimitSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "warning", GoGetter: "Warning"},
			_jsii_.MemberProperty{JsiiProperty: "warningInput", GoGetter: "WarningInput"},
		},
		func() interface{} {
			j := jsiiProxy_NrqlAlertCondition{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionConfig",
		reflect.TypeOf((*NrqlAlertConditionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionCritical",
		reflect.TypeOf((*NrqlAlertConditionCritical)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionCriticalOutputReference",
		reflect.TypeOf((*NrqlAlertConditionCriticalOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "duration", GoGetter: "Duration"},
			_jsii_.MemberProperty{JsiiProperty: "durationInput", GoGetter: "DurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "operatorInput", GoGetter: "OperatorInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDuration", GoMethod: "ResetDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetOperator", GoMethod: "ResetOperator"},
			_jsii_.MemberMethod{JsiiMethod: "resetThresholdDuration", GoMethod: "ResetThresholdDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetThresholdOccurrences", GoMethod: "ResetThresholdOccurrences"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeFunction", GoMethod: "ResetTimeFunction"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdDuration", GoGetter: "ThresholdDuration"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdDurationInput", GoGetter: "ThresholdDurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdInput", GoGetter: "ThresholdInput"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdOccurrences", GoGetter: "ThresholdOccurrences"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdOccurrencesInput", GoGetter: "ThresholdOccurrencesInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeFunction", GoGetter: "TimeFunction"},
			_jsii_.MemberProperty{JsiiProperty: "timeFunctionInput", GoGetter: "TimeFunctionInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NrqlAlertConditionCriticalOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionNrql",
		reflect.TypeOf((*NrqlAlertConditionNrql)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionNrqlOutputReference",
		reflect.TypeOf((*NrqlAlertConditionNrqlOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataAccountId", GoGetter: "DataAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "dataAccountIdInput", GoGetter: "DataAccountIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationOffset", GoGetter: "EvaluationOffset"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationOffsetInput", GoGetter: "EvaluationOffsetInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "query", GoGetter: "Query"},
			_jsii_.MemberProperty{JsiiProperty: "queryInput", GoGetter: "QueryInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataAccountId", GoMethod: "ResetDataAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetEvaluationOffset", GoMethod: "ResetEvaluationOffset"},
			_jsii_.MemberMethod{JsiiMethod: "resetSinceValue", GoMethod: "ResetSinceValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sinceValue", GoGetter: "SinceValue"},
			_jsii_.MemberProperty{JsiiProperty: "sinceValueInput", GoGetter: "SinceValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NrqlAlertConditionNrqlOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionTerm",
		reflect.TypeOf((*NrqlAlertConditionTerm)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionTermList",
		reflect.TypeOf((*NrqlAlertConditionTermList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_NrqlAlertConditionTermList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionTermOutputReference",
		reflect.TypeOf((*NrqlAlertConditionTermOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "duration", GoGetter: "Duration"},
			_jsii_.MemberProperty{JsiiProperty: "durationInput", GoGetter: "DurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "operatorInput", GoGetter: "OperatorInput"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "priorityInput", GoGetter: "PriorityInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDuration", GoMethod: "ResetDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetOperator", GoMethod: "ResetOperator"},
			_jsii_.MemberMethod{JsiiMethod: "resetPriority", GoMethod: "ResetPriority"},
			_jsii_.MemberMethod{JsiiMethod: "resetThresholdDuration", GoMethod: "ResetThresholdDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetThresholdOccurrences", GoMethod: "ResetThresholdOccurrences"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeFunction", GoMethod: "ResetTimeFunction"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdDuration", GoGetter: "ThresholdDuration"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdDurationInput", GoGetter: "ThresholdDurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdInput", GoGetter: "ThresholdInput"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdOccurrences", GoGetter: "ThresholdOccurrences"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdOccurrencesInput", GoGetter: "ThresholdOccurrencesInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeFunction", GoGetter: "TimeFunction"},
			_jsii_.MemberProperty{JsiiProperty: "timeFunctionInput", GoGetter: "TimeFunctionInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NrqlAlertConditionTermOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionTimeouts",
		reflect.TypeOf((*NrqlAlertConditionTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionTimeoutsOutputReference",
		reflect.TypeOf((*NrqlAlertConditionTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NrqlAlertConditionTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionWarning",
		reflect.TypeOf((*NrqlAlertConditionWarning)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionWarningOutputReference",
		reflect.TypeOf((*NrqlAlertConditionWarningOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "duration", GoGetter: "Duration"},
			_jsii_.MemberProperty{JsiiProperty: "durationInput", GoGetter: "DurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "operatorInput", GoGetter: "OperatorInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDuration", GoMethod: "ResetDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetOperator", GoMethod: "ResetOperator"},
			_jsii_.MemberMethod{JsiiMethod: "resetThresholdDuration", GoMethod: "ResetThresholdDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetThresholdOccurrences", GoMethod: "ResetThresholdOccurrences"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeFunction", GoMethod: "ResetTimeFunction"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "threshold", GoGetter: "Threshold"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdDuration", GoGetter: "ThresholdDuration"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdDurationInput", GoGetter: "ThresholdDurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdInput", GoGetter: "ThresholdInput"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdOccurrences", GoGetter: "ThresholdOccurrences"},
			_jsii_.MemberProperty{JsiiProperty: "thresholdOccurrencesInput", GoGetter: "ThresholdOccurrencesInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeFunction", GoGetter: "TimeFunction"},
			_jsii_.MemberProperty{JsiiProperty: "timeFunctionInput", GoGetter: "TimeFunctionInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NrqlAlertConditionWarningOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
