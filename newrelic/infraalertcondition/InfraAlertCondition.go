// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package infraalertcondition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/infraalertcondition/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/infra_alert_condition newrelic_infra_alert_condition}.
type InfraAlertCondition interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Comparison() *string
	SetComparison(val *string)
	ComparisonInput() *string
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
	CreatedAt() *float64
	Critical() InfraAlertConditionCriticalOutputReference
	CriticalInput() *InfraAlertConditionCritical
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EntityGuid() *string
	Event() *string
	SetEvent(val *string)
	EventInput() *string
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
	IntegrationProvider() *string
	SetIntegrationProvider(val *string)
	IntegrationProviderInput() *string
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
	ProcessWhere() *string
	SetProcessWhere(val *string)
	ProcessWhereInput() *string
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
	Select() *string
	SetSelect(val *string)
	SelectInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UpdatedAt() *float64
	ViolationCloseTimer() *float64
	SetViolationCloseTimer(val *float64)
	ViolationCloseTimerInput() *float64
	Warning() InfraAlertConditionWarningOutputReference
	WarningInput() *InfraAlertConditionWarning
	Where() *string
	SetWhere(val *string)
	WhereInput() *string
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
	PutCritical(value *InfraAlertConditionCritical)
	PutWarning(value *InfraAlertConditionWarning)
	ResetComparison()
	ResetCritical()
	ResetDescription()
	ResetEnabled()
	ResetEvent()
	ResetId()
	ResetIntegrationProvider()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProcessWhere()
	ResetRunbookUrl()
	ResetSelect()
	ResetViolationCloseTimer()
	ResetWarning()
	ResetWhere()
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

// The jsii proxy struct for InfraAlertCondition
type jsiiProxy_InfraAlertCondition struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_InfraAlertCondition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Comparison() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparison",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) ComparisonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparisonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) CreatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Critical() InfraAlertConditionCriticalOutputReference {
	var returns InfraAlertConditionCriticalOutputReference
	_jsii_.Get(
		j,
		"critical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) CriticalInput() *InfraAlertConditionCritical {
	var returns *InfraAlertConditionCritical
	_jsii_.Get(
		j,
		"criticalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) EntityGuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityGuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Event() *string {
	var returns *string
	_jsii_.Get(
		j,
		"event",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) EventInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) IntegrationProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) IntegrationProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) PolicyId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) PolicyIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) ProcessWhere() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processWhere",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) ProcessWhereInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processWhereInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) RunbookUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runbookUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) RunbookUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runbookUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Select() *string {
	var returns *string
	_jsii_.Get(
		j,
		"select",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) SelectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) UpdatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) ViolationCloseTimer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"violationCloseTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) ViolationCloseTimerInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"violationCloseTimerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Warning() InfraAlertConditionWarningOutputReference {
	var returns InfraAlertConditionWarningOutputReference
	_jsii_.Get(
		j,
		"warning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) WarningInput() *InfraAlertConditionWarning {
	var returns *InfraAlertConditionWarning
	_jsii_.Get(
		j,
		"warningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) Where() *string {
	var returns *string
	_jsii_.Get(
		j,
		"where",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InfraAlertCondition) WhereInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"whereInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/infra_alert_condition newrelic_infra_alert_condition} Resource.
func NewInfraAlertCondition(scope constructs.Construct, id *string, config *InfraAlertConditionConfig) InfraAlertCondition {
	_init_.Initialize()

	if err := validateNewInfraAlertConditionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_InfraAlertCondition{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.infraAlertCondition.InfraAlertCondition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.65.0/docs/resources/infra_alert_condition newrelic_infra_alert_condition} Resource.
func NewInfraAlertCondition_Override(i InfraAlertCondition, scope constructs.Construct, id *string, config *InfraAlertConditionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.infraAlertCondition.InfraAlertCondition",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetComparison(val *string) {
	if err := j.validateSetComparisonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparison",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetEvent(val *string) {
	if err := j.validateSetEventParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"event",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetIntegrationProvider(val *string) {
	if err := j.validateSetIntegrationProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationProvider",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetPolicyId(val *float64) {
	if err := j.validateSetPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetProcessWhere(val *string) {
	if err := j.validateSetProcessWhereParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processWhere",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetRunbookUrl(val *string) {
	if err := j.validateSetRunbookUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runbookUrl",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetSelect(val *string) {
	if err := j.validateSetSelectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"select",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetViolationCloseTimer(val *float64) {
	if err := j.validateSetViolationCloseTimerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"violationCloseTimer",
		val,
	)
}

func (j *jsiiProxy_InfraAlertCondition)SetWhere(val *string) {
	if err := j.validateSetWhereParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"where",
		val,
	)
}

// Generates CDKTF code for importing a InfraAlertCondition resource upon running "cdktf plan <stack-name>".
func InfraAlertCondition_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateInfraAlertCondition_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.infraAlertCondition.InfraAlertCondition",
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
func InfraAlertCondition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInfraAlertCondition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.infraAlertCondition.InfraAlertCondition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func InfraAlertCondition_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInfraAlertCondition_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.infraAlertCondition.InfraAlertCondition",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func InfraAlertCondition_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateInfraAlertCondition_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.infraAlertCondition.InfraAlertCondition",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func InfraAlertCondition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.infraAlertCondition.InfraAlertCondition",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_InfraAlertCondition) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_InfraAlertCondition) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_InfraAlertCondition) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_InfraAlertCondition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_InfraAlertCondition) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_InfraAlertCondition) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_InfraAlertCondition) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_InfraAlertCondition) PutCritical(value *InfraAlertConditionCritical) {
	if err := i.validatePutCriticalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCritical",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InfraAlertCondition) PutWarning(value *InfraAlertConditionWarning) {
	if err := i.validatePutWarningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putWarning",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_InfraAlertCondition) ResetComparison() {
	_jsii_.InvokeVoid(
		i,
		"resetComparison",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InfraAlertCondition) ResetCritical() {
	_jsii_.InvokeVoid(
		i,
		"resetCritical",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InfraAlertCondition) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InfraAlertCondition) ResetEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InfraAlertCondition) ResetEvent() {
	_jsii_.InvokeVoid(
		i,
		"resetEvent",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InfraAlertCondition) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InfraAlertCondition) ResetIntegrationProvider() {
	_jsii_.InvokeVoid(
		i,
		"resetIntegrationProvider",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InfraAlertCondition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InfraAlertCondition) ResetProcessWhere() {
	_jsii_.InvokeVoid(
		i,
		"resetProcessWhere",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InfraAlertCondition) ResetRunbookUrl() {
	_jsii_.InvokeVoid(
		i,
		"resetRunbookUrl",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InfraAlertCondition) ResetSelect() {
	_jsii_.InvokeVoid(
		i,
		"resetSelect",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InfraAlertCondition) ResetViolationCloseTimer() {
	_jsii_.InvokeVoid(
		i,
		"resetViolationCloseTimer",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InfraAlertCondition) ResetWarning() {
	_jsii_.InvokeVoid(
		i,
		"resetWarning",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InfraAlertCondition) ResetWhere() {
	_jsii_.InvokeVoid(
		i,
		"resetWhere",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InfraAlertCondition) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InfraAlertCondition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

