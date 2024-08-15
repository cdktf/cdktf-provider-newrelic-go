// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datanewrelicservicelevelalerthelper

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/datanewrelicservicelevelalerthelper/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.42.1/docs/data-sources/service_level_alert_helper newrelic_service_level_alert_helper}.
type DataNewrelicServiceLevelAlertHelper interface {
	cdktf.TerraformDataSource
	AlertType() *string
	SetAlertType(val *string)
	AlertTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomEvaluationPeriod() *float64
	SetCustomEvaluationPeriod(val *float64)
	CustomEvaluationPeriodInput() *float64
	CustomToleratedBudgetConsumption() *float64
	SetCustomToleratedBudgetConsumption(val *float64)
	CustomToleratedBudgetConsumptionInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EvaluationPeriod() *float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsBadEvents() interface{}
	SetIsBadEvents(val interface{})
	IsBadEventsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Nrql() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	SliGuid() *string
	SetSliGuid(val *string)
	SliGuidInput() *string
	SloPeriod() *float64
	SetSloPeriod(val *float64)
	SloPeriodInput() *float64
	SloTarget() *float64
	SetSloTarget(val *float64)
	SloTargetInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Threshold() *float64
	ToleratedBudgetConsumption() *float64
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetCustomEvaluationPeriod()
	ResetCustomToleratedBudgetConsumption()
	ResetId()
	ResetIsBadEvents()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataNewrelicServiceLevelAlertHelper
type jsiiProxy_DataNewrelicServiceLevelAlertHelper struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) AlertType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) AlertTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) CustomEvaluationPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customEvaluationPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) CustomEvaluationPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customEvaluationPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) CustomToleratedBudgetConsumption() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customToleratedBudgetConsumption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) CustomToleratedBudgetConsumptionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customToleratedBudgetConsumptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) EvaluationPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) IsBadEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isBadEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) IsBadEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isBadEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) Nrql() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nrql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) SliGuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sliGuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) SliGuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sliGuidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) SloPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sloPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) SloPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sloPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) SloTarget() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sloTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) SloTargetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sloTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper) ToleratedBudgetConsumption() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"toleratedBudgetConsumption",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.42.1/docs/data-sources/service_level_alert_helper newrelic_service_level_alert_helper} Data Source.
func NewDataNewrelicServiceLevelAlertHelper(scope constructs.Construct, id *string, config *DataNewrelicServiceLevelAlertHelperConfig) DataNewrelicServiceLevelAlertHelper {
	_init_.Initialize()

	if err := validateNewDataNewrelicServiceLevelAlertHelperParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataNewrelicServiceLevelAlertHelper{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.dataNewrelicServiceLevelAlertHelper.DataNewrelicServiceLevelAlertHelper",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.42.1/docs/data-sources/service_level_alert_helper newrelic_service_level_alert_helper} Data Source.
func NewDataNewrelicServiceLevelAlertHelper_Override(d DataNewrelicServiceLevelAlertHelper, scope constructs.Construct, id *string, config *DataNewrelicServiceLevelAlertHelperConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.dataNewrelicServiceLevelAlertHelper.DataNewrelicServiceLevelAlertHelper",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper)SetAlertType(val *string) {
	if err := j.validateSetAlertTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertType",
		val,
	)
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper)SetCustomEvaluationPeriod(val *float64) {
	if err := j.validateSetCustomEvaluationPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customEvaluationPeriod",
		val,
	)
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper)SetCustomToleratedBudgetConsumption(val *float64) {
	if err := j.validateSetCustomToleratedBudgetConsumptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customToleratedBudgetConsumption",
		val,
	)
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper)SetIsBadEvents(val interface{}) {
	if err := j.validateSetIsBadEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isBadEvents",
		val,
	)
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper)SetSliGuid(val *string) {
	if err := j.validateSetSliGuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sliGuid",
		val,
	)
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper)SetSloPeriod(val *float64) {
	if err := j.validateSetSloPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sloPeriod",
		val,
	)
}

func (j *jsiiProxy_DataNewrelicServiceLevelAlertHelper)SetSloTarget(val *float64) {
	if err := j.validateSetSloTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sloTarget",
		val,
	)
}

// Generates CDKTF code for importing a DataNewrelicServiceLevelAlertHelper resource upon running "cdktf plan <stack-name>".
func DataNewrelicServiceLevelAlertHelper_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataNewrelicServiceLevelAlertHelper_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.dataNewrelicServiceLevelAlertHelper.DataNewrelicServiceLevelAlertHelper",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DataNewrelicServiceLevelAlertHelper_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataNewrelicServiceLevelAlertHelper_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.dataNewrelicServiceLevelAlertHelper.DataNewrelicServiceLevelAlertHelper",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataNewrelicServiceLevelAlertHelper_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataNewrelicServiceLevelAlertHelper_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.dataNewrelicServiceLevelAlertHelper.DataNewrelicServiceLevelAlertHelper",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataNewrelicServiceLevelAlertHelper_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataNewrelicServiceLevelAlertHelper_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.dataNewrelicServiceLevelAlertHelper.DataNewrelicServiceLevelAlertHelper",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataNewrelicServiceLevelAlertHelper_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.dataNewrelicServiceLevelAlertHelper.DataNewrelicServiceLevelAlertHelper",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) ResetCustomEvaluationPeriod() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomEvaluationPeriod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) ResetCustomToleratedBudgetConsumption() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomToleratedBudgetConsumption",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) ResetIsBadEvents() {
	_jsii_.InvokeVoid(
		d,
		"resetIsBadEvents",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataNewrelicServiceLevelAlertHelper) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

