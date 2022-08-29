// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-newrelic-go/newrelic/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-newrelic-go/newrelic/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/newrelic/r/alert_condition newrelic_alert_condition}.
type AlertCondition interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConditionScope() *string
	SetConditionScope(val *string)
	ConditionScopeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Entities() *[]*float64
	SetEntities(val *[]*float64)
	EntitiesInput() *[]*float64
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcMetric() *string
	SetGcMetric(val *string)
	GcMetricInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Metric() *string
	SetMetric(val *string)
	MetricInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PolicyId() *float64
	SetPolicyId(val *float64)
	PolicyIdInput() *float64
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
	RunbookUrl() *string
	SetRunbookUrl(val *string)
	RunbookUrlInput() *string
	Term() AlertConditionTermList
	TermInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UserDefinedMetric() *string
	SetUserDefinedMetric(val *string)
	UserDefinedMetricInput() *string
	UserDefinedValueFunction() *string
	SetUserDefinedValueFunction(val *string)
	UserDefinedValueFunctionInput() *string
	ViolationCloseTimer() *float64
	SetViolationCloseTimer(val *float64)
	ViolationCloseTimerInput() *float64
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
	PutTerm(value interface{})
	ResetConditionScope()
	ResetEnabled()
	ResetGcMetric()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRunbookUrl()
	ResetUserDefinedMetric()
	ResetUserDefinedValueFunction()
	ResetViolationCloseTimer()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AlertCondition
type jsiiProxy_AlertCondition struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AlertCondition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) ConditionScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) ConditionScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) Entities() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"entities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) EntitiesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"entitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) GcMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) GcMetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) Metric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) MetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) PolicyId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) PolicyIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) RunbookUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runbookUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) RunbookUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runbookUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) Term() AlertConditionTermList {
	var returns AlertConditionTermList
	_jsii_.Get(
		j,
		"term",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) TermInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"termInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) UserDefinedMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDefinedMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) UserDefinedMetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDefinedMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) UserDefinedValueFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDefinedValueFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) UserDefinedValueFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDefinedValueFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) ViolationCloseTimer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"violationCloseTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertCondition) ViolationCloseTimerInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"violationCloseTimerInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/newrelic/r/alert_condition newrelic_alert_condition} Resource.
func NewAlertCondition(scope constructs.Construct, id *string, config *AlertConditionConfig) AlertCondition {
	_init_.Initialize()

	j := jsiiProxy_AlertCondition{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.AlertCondition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/newrelic/r/alert_condition newrelic_alert_condition} Resource.
func NewAlertCondition_Override(a AlertCondition, scope constructs.Construct, id *string, config *AlertConditionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.AlertCondition",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AlertCondition) SetConditionScope(val *string) {
	_jsii_.Set(
		j,
		"conditionScope",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetEntities(val *[]*float64) {
	_jsii_.Set(
		j,
		"entities",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetGcMetric(val *string) {
	_jsii_.Set(
		j,
		"gcMetric",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetMetric(val *string) {
	_jsii_.Set(
		j,
		"metric",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetPolicyId(val *float64) {
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetRunbookUrl(val *string) {
	_jsii_.Set(
		j,
		"runbookUrl",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetUserDefinedMetric(val *string) {
	_jsii_.Set(
		j,
		"userDefinedMetric",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetUserDefinedValueFunction(val *string) {
	_jsii_.Set(
		j,
		"userDefinedValueFunction",
		val,
	)
}

func (j *jsiiProxy_AlertCondition) SetViolationCloseTimer(val *float64) {
	_jsii_.Set(
		j,
		"violationCloseTimer",
		val,
	)
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
func AlertCondition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.AlertCondition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AlertCondition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.AlertCondition",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AlertCondition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AlertCondition) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertCondition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertCondition) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertCondition) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertCondition) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertCondition) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertCondition) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertCondition) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertCondition) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertCondition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertCondition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AlertCondition) PutTerm(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putTerm",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlertCondition) ResetConditionScope() {
	_jsii_.InvokeVoid(
		a,
		"resetConditionScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertCondition) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertCondition) ResetGcMetric() {
	_jsii_.InvokeVoid(
		a,
		"resetGcMetric",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertCondition) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertCondition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertCondition) ResetRunbookUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetRunbookUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertCondition) ResetUserDefinedMetric() {
	_jsii_.InvokeVoid(
		a,
		"resetUserDefinedMetric",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertCondition) ResetUserDefinedValueFunction() {
	_jsii_.InvokeVoid(
		a,
		"resetUserDefinedValueFunction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertCondition) ResetViolationCloseTimer() {
	_jsii_.InvokeVoid(
		a,
		"resetViolationCloseTimer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertCondition) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertCondition) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertCondition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertCondition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
