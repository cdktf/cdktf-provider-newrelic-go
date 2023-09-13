// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsmultilocationalertcondition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v10/syntheticsmultilocationalertcondition/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/synthetics_multilocation_alert_condition newrelic_synthetics_multilocation_alert_condition}.
type SyntheticsMultilocationAlertCondition interface {
	cdktf.TerraformResource
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
	Critical() SyntheticsMultilocationAlertConditionCriticalOutputReference
	CriticalInput() *SyntheticsMultilocationAlertConditionCritical
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Entities() *[]*string
	SetEntities(val *[]*string)
	EntitiesInput() *[]*string
	EntityGuid() *string
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ViolationTimeLimitSeconds() *float64
	SetViolationTimeLimitSeconds(val *float64)
	ViolationTimeLimitSecondsInput() *float64
	Warning() SyntheticsMultilocationAlertConditionWarningOutputReference
	WarningInput() *SyntheticsMultilocationAlertConditionWarning
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
	PutCritical(value *SyntheticsMultilocationAlertConditionCritical)
	PutWarning(value *SyntheticsMultilocationAlertConditionWarning)
	ResetEnabled()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRunbookUrl()
	ResetWarning()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SyntheticsMultilocationAlertCondition
type jsiiProxy_SyntheticsMultilocationAlertCondition struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) Critical() SyntheticsMultilocationAlertConditionCriticalOutputReference {
	var returns SyntheticsMultilocationAlertConditionCriticalOutputReference
	_jsii_.Get(
		j,
		"critical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) CriticalInput() *SyntheticsMultilocationAlertConditionCritical {
	var returns *SyntheticsMultilocationAlertConditionCritical
	_jsii_.Get(
		j,
		"criticalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) Entities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) EntitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"entitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) EntityGuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityGuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) PolicyId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) PolicyIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) RunbookUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runbookUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) RunbookUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runbookUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) ViolationTimeLimitSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"violationTimeLimitSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) ViolationTimeLimitSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"violationTimeLimitSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) Warning() SyntheticsMultilocationAlertConditionWarningOutputReference {
	var returns SyntheticsMultilocationAlertConditionWarningOutputReference
	_jsii_.Get(
		j,
		"warning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition) WarningInput() *SyntheticsMultilocationAlertConditionWarning {
	var returns *SyntheticsMultilocationAlertConditionWarning
	_jsii_.Get(
		j,
		"warningInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/synthetics_multilocation_alert_condition newrelic_synthetics_multilocation_alert_condition} Resource.
func NewSyntheticsMultilocationAlertCondition(scope constructs.Construct, id *string, config *SyntheticsMultilocationAlertConditionConfig) SyntheticsMultilocationAlertCondition {
	_init_.Initialize()

	if err := validateNewSyntheticsMultilocationAlertConditionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsMultilocationAlertCondition{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.syntheticsMultilocationAlertCondition.SyntheticsMultilocationAlertCondition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.27.1/docs/resources/synthetics_multilocation_alert_condition newrelic_synthetics_multilocation_alert_condition} Resource.
func NewSyntheticsMultilocationAlertCondition_Override(s SyntheticsMultilocationAlertCondition, scope constructs.Construct, id *string, config *SyntheticsMultilocationAlertConditionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.syntheticsMultilocationAlertCondition.SyntheticsMultilocationAlertCondition",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition)SetEntities(val *[]*string) {
	if err := j.validateSetEntitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entities",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition)SetPolicyId(val *float64) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition)SetRunbookUrl(val *string) {
	if err := j.validateSetRunbookUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runbookUrl",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMultilocationAlertCondition)SetViolationTimeLimitSeconds(val *float64) {
	if err := j.validateSetViolationTimeLimitSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"violationTimeLimitSeconds",
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
func SyntheticsMultilocationAlertCondition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsMultilocationAlertCondition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.syntheticsMultilocationAlertCondition.SyntheticsMultilocationAlertCondition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsMultilocationAlertCondition_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsMultilocationAlertCondition_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.syntheticsMultilocationAlertCondition.SyntheticsMultilocationAlertCondition",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsMultilocationAlertCondition_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsMultilocationAlertCondition_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.syntheticsMultilocationAlertCondition.SyntheticsMultilocationAlertCondition",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SyntheticsMultilocationAlertCondition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.syntheticsMultilocationAlertCondition.SyntheticsMultilocationAlertCondition",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) PutCritical(value *SyntheticsMultilocationAlertConditionCritical) {
	if err := s.validatePutCriticalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCritical",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) PutWarning(value *SyntheticsMultilocationAlertConditionWarning) {
	if err := s.validatePutWarningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putWarning",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) ResetRunbookUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetRunbookUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) ResetWarning() {
	_jsii_.InvokeVoid(
		s,
		"resetWarning",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMultilocationAlertCondition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

