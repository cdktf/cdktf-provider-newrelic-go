// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsscriptmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/syntheticsscriptmonitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.44.0/docs/resources/synthetics_script_monitor newrelic_synthetics_script_monitor}.
type SyntheticsScriptMonitor interface {
	cdktf.TerraformResource
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
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
	DeviceOrientation() *string
	SetDeviceOrientation(val *string)
	DeviceOrientationInput() *string
	DeviceType() *string
	SetDeviceType(val *string)
	DeviceTypeInput() *string
	EnableScreenshotOnFailureAndScript() interface{}
	SetEnableScreenshotOnFailureAndScript(val interface{})
	EnableScreenshotOnFailureAndScriptInput() interface{}
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
	LocationPrivate() SyntheticsScriptMonitorLocationPrivateList
	LocationPrivateInput() interface{}
	LocationsPublic() *[]*string
	SetLocationsPublic(val *[]*string)
	LocationsPublicInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Period() *string
	SetPeriod(val *string)
	PeriodInMinutes() *float64
	PeriodInput() *string
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
	RuntimeType() *string
	SetRuntimeType(val *string)
	RuntimeTypeInput() *string
	RuntimeTypeVersion() *string
	SetRuntimeTypeVersion(val *string)
	RuntimeTypeVersionInput() *string
	Script() *string
	SetScript(val *string)
	ScriptInput() *string
	ScriptLanguage() *string
	SetScriptLanguage(val *string)
	ScriptLanguageInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	Tag() SyntheticsScriptMonitorTagList
	TagInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UseUnsupportedLegacyRuntime() interface{}
	SetUseUnsupportedLegacyRuntime(val interface{})
	UseUnsupportedLegacyRuntimeInput() interface{}
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
	PutLocationPrivate(value interface{})
	PutTag(value interface{})
	ResetAccountId()
	ResetDeviceOrientation()
	ResetDeviceType()
	ResetEnableScreenshotOnFailureAndScript()
	ResetId()
	ResetLocationPrivate()
	ResetLocationsPublic()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRuntimeType()
	ResetRuntimeTypeVersion()
	ResetScript()
	ResetScriptLanguage()
	ResetTag()
	ResetUseUnsupportedLegacyRuntime()
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

// The jsii proxy struct for SyntheticsScriptMonitor
type jsiiProxy_SyntheticsScriptMonitor struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SyntheticsScriptMonitor) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) DeviceOrientation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceOrientation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) DeviceOrientationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceOrientationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) DeviceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) DeviceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) EnableScreenshotOnFailureAndScript() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableScreenshotOnFailureAndScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) EnableScreenshotOnFailureAndScriptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableScreenshotOnFailureAndScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) Guid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) LocationPrivate() SyntheticsScriptMonitorLocationPrivateList {
	var returns SyntheticsScriptMonitorLocationPrivateList
	_jsii_.Get(
		j,
		"locationPrivate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) LocationPrivateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locationPrivateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) LocationsPublic() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsPublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) LocationsPublicInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsPublicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) Period() *string {
	var returns *string
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) PeriodInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) PeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) RuntimeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) RuntimeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) RuntimeTypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeTypeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) RuntimeTypeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeTypeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) Script() *string {
	var returns *string
	_jsii_.Get(
		j,
		"script",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) ScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) ScriptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) ScriptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) Tag() SyntheticsScriptMonitorTagList {
	var returns SyntheticsScriptMonitorTagList
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) TagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) UseUnsupportedLegacyRuntime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useUnsupportedLegacyRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsScriptMonitor) UseUnsupportedLegacyRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useUnsupportedLegacyRuntimeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.44.0/docs/resources/synthetics_script_monitor newrelic_synthetics_script_monitor} Resource.
func NewSyntheticsScriptMonitor(scope constructs.Construct, id *string, config *SyntheticsScriptMonitorConfig) SyntheticsScriptMonitor {
	_init_.Initialize()

	if err := validateNewSyntheticsScriptMonitorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsScriptMonitor{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.syntheticsScriptMonitor.SyntheticsScriptMonitor",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.44.0/docs/resources/synthetics_script_monitor newrelic_synthetics_script_monitor} Resource.
func NewSyntheticsScriptMonitor_Override(s SyntheticsScriptMonitor, scope constructs.Construct, id *string, config *SyntheticsScriptMonitorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.syntheticsScriptMonitor.SyntheticsScriptMonitor",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetAccountId(val *float64) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetDeviceOrientation(val *string) {
	if err := j.validateSetDeviceOrientationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceOrientation",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetDeviceType(val *string) {
	if err := j.validateSetDeviceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceType",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetEnableScreenshotOnFailureAndScript(val interface{}) {
	if err := j.validateSetEnableScreenshotOnFailureAndScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableScreenshotOnFailureAndScript",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetLocationsPublic(val *[]*string) {
	if err := j.validateSetLocationsPublicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationsPublic",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetPeriod(val *string) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetRuntimeType(val *string) {
	if err := j.validateSetRuntimeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeType",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetRuntimeTypeVersion(val *string) {
	if err := j.validateSetRuntimeTypeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeTypeVersion",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetScript(val *string) {
	if err := j.validateSetScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"script",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetScriptLanguage(val *string) {
	if err := j.validateSetScriptLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptLanguage",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_SyntheticsScriptMonitor)SetUseUnsupportedLegacyRuntime(val interface{}) {
	if err := j.validateSetUseUnsupportedLegacyRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useUnsupportedLegacyRuntime",
		val,
	)
}

// Generates CDKTF code for importing a SyntheticsScriptMonitor resource upon running "cdktf plan <stack-name>".
func SyntheticsScriptMonitor_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSyntheticsScriptMonitor_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.syntheticsScriptMonitor.SyntheticsScriptMonitor",
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
func SyntheticsScriptMonitor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsScriptMonitor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.syntheticsScriptMonitor.SyntheticsScriptMonitor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsScriptMonitor_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsScriptMonitor_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.syntheticsScriptMonitor.SyntheticsScriptMonitor",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsScriptMonitor_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsScriptMonitor_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.syntheticsScriptMonitor.SyntheticsScriptMonitor",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SyntheticsScriptMonitor_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.syntheticsScriptMonitor.SyntheticsScriptMonitor",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SyntheticsScriptMonitor) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SyntheticsScriptMonitor) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsScriptMonitor) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SyntheticsScriptMonitor) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SyntheticsScriptMonitor) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SyntheticsScriptMonitor) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SyntheticsScriptMonitor) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SyntheticsScriptMonitor) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SyntheticsScriptMonitor) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SyntheticsScriptMonitor) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsScriptMonitor) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) PutLocationPrivate(value interface{}) {
	if err := s.validatePutLocationPrivateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putLocationPrivate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) PutTag(value interface{}) {
	if err := s.validatePutTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTag",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ResetAccountId() {
	_jsii_.InvokeVoid(
		s,
		"resetAccountId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ResetDeviceOrientation() {
	_jsii_.InvokeVoid(
		s,
		"resetDeviceOrientation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ResetDeviceType() {
	_jsii_.InvokeVoid(
		s,
		"resetDeviceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ResetEnableScreenshotOnFailureAndScript() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableScreenshotOnFailureAndScript",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ResetLocationPrivate() {
	_jsii_.InvokeVoid(
		s,
		"resetLocationPrivate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ResetLocationsPublic() {
	_jsii_.InvokeVoid(
		s,
		"resetLocationsPublic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ResetRuntimeType() {
	_jsii_.InvokeVoid(
		s,
		"resetRuntimeType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ResetRuntimeTypeVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetRuntimeTypeVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ResetScript() {
	_jsii_.InvokeVoid(
		s,
		"resetScript",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ResetScriptLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetScriptLanguage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ResetTag() {
	_jsii_.InvokeVoid(
		s,
		"resetTag",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ResetUseUnsupportedLegacyRuntime() {
	_jsii_.InvokeVoid(
		s,
		"resetUseUnsupportedLegacyRuntime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsScriptMonitor) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsScriptMonitor) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsScriptMonitor) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

