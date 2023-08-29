// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v10/workload/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/workload newrelic_workload}.
type Workload interface {
	cdktf.TerraformResource
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CompositeEntitySearchQuery() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EntityGuids() *[]*string
	SetEntityGuids(val *[]*string)
	EntityGuidsInput() *[]*string
	EntitySearchQuery() WorkloadEntitySearchQueryList
	EntitySearchQueryInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Guid() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Permalink() *string
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
	ScopeAccountIds() *[]*float64
	SetScopeAccountIds(val *[]*float64)
	ScopeAccountIdsInput() *[]*float64
	StatusConfigAutomatic() WorkloadStatusConfigAutomaticOutputReference
	StatusConfigAutomaticInput() *WorkloadStatusConfigAutomatic
	StatusConfigStatic() WorkloadStatusConfigStaticOutputReference
	StatusConfigStaticInput() *WorkloadStatusConfigStatic
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WorkloadId() *float64
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
	PutEntitySearchQuery(value interface{})
	PutStatusConfigAutomatic(value *WorkloadStatusConfigAutomatic)
	PutStatusConfigStatic(value *WorkloadStatusConfigStatic)
	ResetAccountId()
	ResetDescription()
	ResetEntityGuids()
	ResetEntitySearchQuery()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScopeAccountIds()
	ResetStatusConfigAutomatic()
	ResetStatusConfigStatic()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Workload
type jsiiProxy_Workload struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Workload) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) CompositeEntitySearchQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compositeEntitySearchQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) EntityGuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entityGuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) EntityGuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entityGuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) EntitySearchQuery() WorkloadEntitySearchQueryList {
	var returns WorkloadEntitySearchQueryList
	_jsii_.Get(
		j,
		"entitySearchQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) EntitySearchQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entitySearchQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Guid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Permalink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permalink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) ScopeAccountIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"scopeAccountIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) ScopeAccountIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"scopeAccountIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) StatusConfigAutomatic() WorkloadStatusConfigAutomaticOutputReference {
	var returns WorkloadStatusConfigAutomaticOutputReference
	_jsii_.Get(
		j,
		"statusConfigAutomatic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) StatusConfigAutomaticInput() *WorkloadStatusConfigAutomatic {
	var returns *WorkloadStatusConfigAutomatic
	_jsii_.Get(
		j,
		"statusConfigAutomaticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) StatusConfigStatic() WorkloadStatusConfigStaticOutputReference {
	var returns WorkloadStatusConfigStaticOutputReference
	_jsii_.Get(
		j,
		"statusConfigStatic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) StatusConfigStaticInput() *WorkloadStatusConfigStatic {
	var returns *WorkloadStatusConfigStatic
	_jsii_.Get(
		j,
		"statusConfigStaticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workload) WorkloadId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workloadId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/workload newrelic_workload} Resource.
func NewWorkload(scope constructs.Construct, id *string, config *WorkloadConfig) Workload {
	_init_.Initialize()

	if err := validateNewWorkloadParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Workload{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.workload.Workload",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.26.1/docs/resources/workload newrelic_workload} Resource.
func NewWorkload_Override(w Workload, scope constructs.Construct, id *string, config *WorkloadConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.workload.Workload",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_Workload)SetAccountId(val *float64) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_Workload)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Workload)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Workload)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Workload)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Workload)SetEntityGuids(val *[]*string) {
	if err := j.validateSetEntityGuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityGuids",
		val,
	)
}

func (j *jsiiProxy_Workload)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Workload)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Workload)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Workload)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Workload)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Workload)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Workload)SetScopeAccountIds(val *[]*float64) {
	if err := j.validateSetScopeAccountIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopeAccountIds",
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
func Workload_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkload_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.workload.Workload",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Workload_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkload_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.workload.Workload",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Workload_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkload_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.workload.Workload",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Workload_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.workload.Workload",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_Workload) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_Workload) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_Workload) PutEntitySearchQuery(value interface{}) {
	if err := w.validatePutEntitySearchQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putEntitySearchQuery",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Workload) PutStatusConfigAutomatic(value *WorkloadStatusConfigAutomatic) {
	if err := w.validatePutStatusConfigAutomaticParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putStatusConfigAutomatic",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Workload) PutStatusConfigStatic(value *WorkloadStatusConfigStatic) {
	if err := w.validatePutStatusConfigStaticParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putStatusConfigStatic",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Workload) ResetAccountId() {
	_jsii_.InvokeVoid(
		w,
		"resetAccountId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workload) ResetDescription() {
	_jsii_.InvokeVoid(
		w,
		"resetDescription",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workload) ResetEntityGuids() {
	_jsii_.InvokeVoid(
		w,
		"resetEntityGuids",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workload) ResetEntitySearchQuery() {
	_jsii_.InvokeVoid(
		w,
		"resetEntitySearchQuery",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workload) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workload) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workload) ResetScopeAccountIds() {
	_jsii_.InvokeVoid(
		w,
		"resetScopeAccountIds",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workload) ResetStatusConfigAutomatic() {
	_jsii_.InvokeVoid(
		w,
		"resetStatusConfigAutomatic",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workload) ResetStatusConfigStatic() {
	_jsii_.InvokeVoid(
		w,
		"resetStatusConfigStatic",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workload) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workload) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

