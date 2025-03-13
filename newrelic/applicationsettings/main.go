// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationsettings

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettings",
		reflect.TypeOf((*ApplicationSettings)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "appApdexThreshold", GoGetter: "AppApdexThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "appApdexThresholdInput", GoGetter: "AppApdexThresholdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enableRealUserMonitoring", GoGetter: "EnableRealUserMonitoring"},
			_jsii_.MemberProperty{JsiiProperty: "enableRealUserMonitoringInput", GoGetter: "EnableRealUserMonitoringInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableSlowSql", GoGetter: "EnableSlowSql"},
			_jsii_.MemberProperty{JsiiProperty: "enableSlowSqlInput", GoGetter: "EnableSlowSqlInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableThreadProfiler", GoGetter: "EnableThreadProfiler"},
			_jsii_.MemberProperty{JsiiProperty: "enableThreadProfilerInput", GoGetter: "EnableThreadProfilerInput"},
			_jsii_.MemberProperty{JsiiProperty: "endUserApdexThreshold", GoGetter: "EndUserApdexThreshold"},
			_jsii_.MemberProperty{JsiiProperty: "endUserApdexThresholdInput", GoGetter: "EndUserApdexThresholdInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorCollector", GoGetter: "ErrorCollector"},
			_jsii_.MemberProperty{JsiiProperty: "errorCollectorInput", GoGetter: "ErrorCollectorInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "guid", GoGetter: "Guid"},
			_jsii_.MemberProperty{JsiiProperty: "guidInput", GoGetter: "GuidInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isImported", GoGetter: "IsImported"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putErrorCollector", GoMethod: "PutErrorCollector"},
			_jsii_.MemberMethod{JsiiMethod: "putTransactionTracer", GoMethod: "PutTransactionTracer"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppApdexThreshold", GoMethod: "ResetAppApdexThreshold"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableRealUserMonitoring", GoMethod: "ResetEnableRealUserMonitoring"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableSlowSql", GoMethod: "ResetEnableSlowSql"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableThreadProfiler", GoMethod: "ResetEnableThreadProfiler"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndUserApdexThreshold", GoMethod: "ResetEndUserApdexThreshold"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorCollector", GoMethod: "ResetErrorCollector"},
			_jsii_.MemberMethod{JsiiMethod: "resetGuid", GoMethod: "ResetGuid"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTracerType", GoMethod: "ResetTracerType"},
			_jsii_.MemberMethod{JsiiMethod: "resetTransactionTracer", GoMethod: "ResetTransactionTracer"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseServerSideConfig", GoMethod: "ResetUseServerSideConfig"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "tracerType", GoGetter: "TracerType"},
			_jsii_.MemberProperty{JsiiProperty: "tracerTypeInput", GoGetter: "TracerTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "transactionTracer", GoGetter: "TransactionTracer"},
			_jsii_.MemberProperty{JsiiProperty: "transactionTracerInput", GoGetter: "TransactionTracerInput"},
			_jsii_.MemberProperty{JsiiProperty: "useServerSideConfig", GoGetter: "UseServerSideConfig"},
			_jsii_.MemberProperty{JsiiProperty: "useServerSideConfigInput", GoGetter: "UseServerSideConfigInput"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationSettings{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsConfig",
		reflect.TypeOf((*ApplicationSettingsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsErrorCollector",
		reflect.TypeOf((*ApplicationSettingsErrorCollector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsErrorCollectorList",
		reflect.TypeOf((*ApplicationSettingsErrorCollectorList)(nil)).Elem(),
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
			j := jsiiProxy_ApplicationSettingsErrorCollectorList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsErrorCollectorOutputReference",
		reflect.TypeOf((*ApplicationSettingsErrorCollectorOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "expectedErrorClasses", GoGetter: "ExpectedErrorClasses"},
			_jsii_.MemberProperty{JsiiProperty: "expectedErrorClassesInput", GoGetter: "ExpectedErrorClassesInput"},
			_jsii_.MemberProperty{JsiiProperty: "expectedErrorCodes", GoGetter: "ExpectedErrorCodes"},
			_jsii_.MemberProperty{JsiiProperty: "expectedErrorCodesInput", GoGetter: "ExpectedErrorCodesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "ignoredErrorClasses", GoGetter: "IgnoredErrorClasses"},
			_jsii_.MemberProperty{JsiiProperty: "ignoredErrorClassesInput", GoGetter: "IgnoredErrorClassesInput"},
			_jsii_.MemberProperty{JsiiProperty: "ignoredErrorCodes", GoGetter: "IgnoredErrorCodes"},
			_jsii_.MemberProperty{JsiiProperty: "ignoredErrorCodesInput", GoGetter: "IgnoredErrorCodesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetExpectedErrorClasses", GoMethod: "ResetExpectedErrorClasses"},
			_jsii_.MemberMethod{JsiiMethod: "resetExpectedErrorCodes", GoMethod: "ResetExpectedErrorCodes"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoredErrorClasses", GoMethod: "ResetIgnoredErrorClasses"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoredErrorCodes", GoMethod: "ResetIgnoredErrorCodes"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationSettingsErrorCollectorOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsTransactionTracer",
		reflect.TypeOf((*ApplicationSettingsTransactionTracer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsTransactionTracerExplainQueryPlans",
		reflect.TypeOf((*ApplicationSettingsTransactionTracerExplainQueryPlans)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsTransactionTracerExplainQueryPlansList",
		reflect.TypeOf((*ApplicationSettingsTransactionTracerExplainQueryPlansList)(nil)).Elem(),
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
			j := jsiiProxy_ApplicationSettingsTransactionTracerExplainQueryPlansList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsTransactionTracerExplainQueryPlansOutputReference",
		reflect.TypeOf((*ApplicationSettingsTransactionTracerExplainQueryPlansOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
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
			_jsii_.MemberProperty{JsiiProperty: "queryPlanThresholdType", GoGetter: "QueryPlanThresholdType"},
			_jsii_.MemberProperty{JsiiProperty: "queryPlanThresholdTypeInput", GoGetter: "QueryPlanThresholdTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryPlanThresholdValue", GoGetter: "QueryPlanThresholdValue"},
			_jsii_.MemberProperty{JsiiProperty: "queryPlanThresholdValueInput", GoGetter: "QueryPlanThresholdValueInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryPlanThresholdType", GoMethod: "ResetQueryPlanThresholdType"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryPlanThresholdValue", GoMethod: "ResetQueryPlanThresholdValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationSettingsTransactionTracerExplainQueryPlansOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsTransactionTracerList",
		reflect.TypeOf((*ApplicationSettingsTransactionTracerList)(nil)).Elem(),
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
			j := jsiiProxy_ApplicationSettingsTransactionTracerList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsTransactionTracerOutputReference",
		reflect.TypeOf((*ApplicationSettingsTransactionTracerOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "explainQueryPlans", GoGetter: "ExplainQueryPlans"},
			_jsii_.MemberProperty{JsiiProperty: "explainQueryPlansInput", GoGetter: "ExplainQueryPlansInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putExplainQueryPlans", GoMethod: "PutExplainQueryPlans"},
			_jsii_.MemberMethod{JsiiMethod: "putSql", GoMethod: "PutSql"},
			_jsii_.MemberMethod{JsiiMethod: "resetExplainQueryPlans", GoMethod: "ResetExplainQueryPlans"},
			_jsii_.MemberMethod{JsiiMethod: "resetSql", GoMethod: "ResetSql"},
			_jsii_.MemberMethod{JsiiMethod: "resetStackTraceThresholdValue", GoMethod: "ResetStackTraceThresholdValue"},
			_jsii_.MemberMethod{JsiiMethod: "resetTransactionThresholdType", GoMethod: "ResetTransactionThresholdType"},
			_jsii_.MemberMethod{JsiiMethod: "resetTransactionThresholdValue", GoMethod: "ResetTransactionThresholdValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sql", GoGetter: "Sql"},
			_jsii_.MemberProperty{JsiiProperty: "sqlInput", GoGetter: "SqlInput"},
			_jsii_.MemberProperty{JsiiProperty: "stackTraceThresholdValue", GoGetter: "StackTraceThresholdValue"},
			_jsii_.MemberProperty{JsiiProperty: "stackTraceThresholdValueInput", GoGetter: "StackTraceThresholdValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transactionThresholdType", GoGetter: "TransactionThresholdType"},
			_jsii_.MemberProperty{JsiiProperty: "transactionThresholdTypeInput", GoGetter: "TransactionThresholdTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "transactionThresholdValue", GoGetter: "TransactionThresholdValue"},
			_jsii_.MemberProperty{JsiiProperty: "transactionThresholdValueInput", GoGetter: "TransactionThresholdValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationSettingsTransactionTracerOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsTransactionTracerSql",
		reflect.TypeOf((*ApplicationSettingsTransactionTracerSql)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsTransactionTracerSqlOutputReference",
		reflect.TypeOf((*ApplicationSettingsTransactionTracerSqlOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
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
			_jsii_.MemberProperty{JsiiProperty: "recordSql", GoGetter: "RecordSql"},
			_jsii_.MemberProperty{JsiiProperty: "recordSqlInput", GoGetter: "RecordSqlInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ApplicationSettingsTransactionTracerSqlOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
