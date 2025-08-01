// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/applicationsettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/application_settings newrelic_application_settings}.
type ApplicationSettings interface {
	cdktf.TerraformResource
	AppApdexThreshold() *float64
	SetAppApdexThreshold(val *float64)
	AppApdexThresholdInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableRealUserMonitoring() interface{}
	SetEnableRealUserMonitoring(val interface{})
	EnableRealUserMonitoringInput() interface{}
	EnableSlowSql() interface{}
	SetEnableSlowSql(val interface{})
	EnableSlowSqlInput() interface{}
	EnableThreadProfiler() interface{}
	SetEnableThreadProfiler(val interface{})
	EnableThreadProfilerInput() interface{}
	EndUserApdexThreshold() *float64
	SetEndUserApdexThreshold(val *float64)
	EndUserApdexThresholdInput() *float64
	ErrorCollector() ApplicationSettingsErrorCollectorList
	ErrorCollectorInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Guid() *string
	SetGuid(val *string)
	GuidInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IsImported() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TracerType() *string
	SetTracerType(val *string)
	TracerTypeInput() *string
	TransactionTracer() ApplicationSettingsTransactionTracerList
	TransactionTracerInput() interface{}
	UseServerSideConfig() interface{}
	SetUseServerSideConfig(val interface{})
	UseServerSideConfigInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutErrorCollector(value interface{})
	PutTransactionTracer(value interface{})
	ResetAppApdexThreshold()
	ResetEnableRealUserMonitoring()
	ResetEnableSlowSql()
	ResetEnableThreadProfiler()
	ResetEndUserApdexThreshold()
	ResetErrorCollector()
	ResetGuid()
	ResetId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTracerType()
	ResetTransactionTracer()
	ResetUseServerSideConfig()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
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

// The jsii proxy struct for ApplicationSettings
type jsiiProxy_ApplicationSettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApplicationSettings) AppApdexThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appApdexThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) AppApdexThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appApdexThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EnableRealUserMonitoring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRealUserMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EnableRealUserMonitoringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRealUserMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EnableSlowSql() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSlowSql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EnableSlowSqlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSlowSqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EnableThreadProfiler() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableThreadProfiler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EnableThreadProfilerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableThreadProfilerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EndUserApdexThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endUserApdexThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) EndUserApdexThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endUserApdexThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ErrorCollector() ApplicationSettingsErrorCollectorList {
	var returns ApplicationSettingsErrorCollectorList
	_jsii_.Get(
		j,
		"errorCollector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ErrorCollectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorCollectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Guid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) GuidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) IsImported() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isImported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TracerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TracerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TransactionTracer() ApplicationSettingsTransactionTracerList {
	var returns ApplicationSettingsTransactionTracerList
	_jsii_.Get(
		j,
		"transactionTracer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) TransactionTracerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transactionTracerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UseServerSideConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useServerSideConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettings) UseServerSideConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useServerSideConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/application_settings newrelic_application_settings} Resource.
func NewApplicationSettings(scope constructs.Construct, id *string, config *ApplicationSettingsConfig) ApplicationSettings {
	_init_.Initialize()

	if err := validateNewApplicationSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationSettings{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/application_settings newrelic_application_settings} Resource.
func NewApplicationSettings_Override(a ApplicationSettings, scope constructs.Construct, id *string, config *ApplicationSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettings",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetAppApdexThreshold(val *float64) {
	if err := j.validateSetAppApdexThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appApdexThreshold",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEnableRealUserMonitoring(val interface{}) {
	if err := j.validateSetEnableRealUserMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableRealUserMonitoring",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEnableSlowSql(val interface{}) {
	if err := j.validateSetEnableSlowSqlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSlowSql",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEnableThreadProfiler(val interface{}) {
	if err := j.validateSetEnableThreadProfilerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableThreadProfiler",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetEndUserApdexThreshold(val *float64) {
	if err := j.validateSetEndUserApdexThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endUserApdexThreshold",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetGuid(val *string) {
	if err := j.validateSetGuidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guid",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetTracerType(val *string) {
	if err := j.validateSetTracerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tracerType",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettings)SetUseServerSideConfig(val interface{}) {
	if err := j.validateSetUseServerSideConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useServerSideConfig",
		val,
	)
}

// Generates CDKTF code for importing a ApplicationSettings resource upon running "cdktf plan <stack-name>".
func ApplicationSettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateApplicationSettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettings",
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
func ApplicationSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ApplicationSettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateApplicationSettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApplicationSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApplicationSettings) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ApplicationSettings) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApplicationSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ApplicationSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApplicationSettings) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ApplicationSettings) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ApplicationSettings) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApplicationSettings) PutErrorCollector(value interface{}) {
	if err := a.validatePutErrorCollectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putErrorCollector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationSettings) PutTransactionTracer(value interface{}) {
	if err := a.validatePutTransactionTracerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTransactionTracer",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetAppApdexThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetAppApdexThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEnableRealUserMonitoring() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableRealUserMonitoring",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEnableSlowSql() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableSlowSql",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEnableThreadProfiler() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableThreadProfiler",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetEndUserApdexThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetEndUserApdexThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetErrorCollector() {
	_jsii_.InvokeVoid(
		a,
		"resetErrorCollector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetGuid() {
	_jsii_.InvokeVoid(
		a,
		"resetGuid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetTracerType() {
	_jsii_.InvokeVoid(
		a,
		"resetTracerType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetTransactionTracer() {
	_jsii_.InvokeVoid(
		a,
		"resetTransactionTracer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) ResetUseServerSideConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetUseServerSideConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

