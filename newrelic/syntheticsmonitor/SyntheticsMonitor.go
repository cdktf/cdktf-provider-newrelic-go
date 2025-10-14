// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticsmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/syntheticsmonitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/synthetics_monitor newrelic_synthetics_monitor}.
type SyntheticsMonitor interface {
	cdktf.TerraformResource
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
	Browsers() *[]*string
	SetBrowsers(val *[]*string)
	BrowsersInput() *[]*string
	BypassHeadRequest() interface{}
	SetBypassHeadRequest(val interface{})
	BypassHeadRequestInput() interface{}
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
	CustomHeader() SyntheticsMonitorCustomHeaderList
	CustomHeaderInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeviceOrientation() *string
	SetDeviceOrientation(val *string)
	DeviceOrientationInput() *string
	Devices() *[]*string
	SetDevices(val *[]*string)
	DevicesInput() *[]*string
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocationsPrivate() *[]*string
	SetLocationsPrivate(val *[]*string)
	LocationsPrivateInput() *[]*string
	LocationsPublic() *[]*string
	SetLocationsPublic(val *[]*string)
	LocationsPublicInput() *[]*string
	MonitorId() *string
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
	ScriptLanguage() *string
	SetScriptLanguage(val *string)
	ScriptLanguageInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	Tag() SyntheticsMonitorTagList
	TagInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TreatRedirectAsFailure() interface{}
	SetTreatRedirectAsFailure(val interface{})
	TreatRedirectAsFailureInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Uri() *string
	SetUri(val *string)
	UriInput() *string
	UseUnsupportedLegacyRuntime() interface{}
	SetUseUnsupportedLegacyRuntime(val interface{})
	UseUnsupportedLegacyRuntimeInput() interface{}
	ValidationString() *string
	SetValidationString(val *string)
	ValidationStringInput() *string
	VerifySsl() interface{}
	SetVerifySsl(val interface{})
	VerifySslInput() interface{}
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
	PutCustomHeader(value interface{})
	PutTag(value interface{})
	ResetAccountId()
	ResetBrowsers()
	ResetBypassHeadRequest()
	ResetCustomHeader()
	ResetDeviceOrientation()
	ResetDevices()
	ResetDeviceType()
	ResetEnableScreenshotOnFailureAndScript()
	ResetId()
	ResetLocationsPrivate()
	ResetLocationsPublic()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeriod()
	ResetRuntimeType()
	ResetRuntimeTypeVersion()
	ResetScriptLanguage()
	ResetTag()
	ResetTreatRedirectAsFailure()
	ResetUri()
	ResetUseUnsupportedLegacyRuntime()
	ResetValidationString()
	ResetVerifySsl()
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

// The jsii proxy struct for SyntheticsMonitor
type jsiiProxy_SyntheticsMonitor struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SyntheticsMonitor) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Browsers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"browsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) BrowsersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"browsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) BypassHeadRequest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassHeadRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) BypassHeadRequestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassHeadRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) CustomHeader() SyntheticsMonitorCustomHeaderList {
	var returns SyntheticsMonitorCustomHeaderList
	_jsii_.Get(
		j,
		"customHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) CustomHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) DeviceOrientation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceOrientation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) DeviceOrientationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceOrientationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Devices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"devices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) DevicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"devicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) DeviceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) DeviceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) EnableScreenshotOnFailureAndScript() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableScreenshotOnFailureAndScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) EnableScreenshotOnFailureAndScriptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableScreenshotOnFailureAndScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) LocationsPrivate() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsPrivate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) LocationsPrivateInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsPrivateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) LocationsPublic() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsPublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) LocationsPublicInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsPublicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) MonitorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Period() *string {
	var returns *string
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) PeriodInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) PeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) RuntimeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) RuntimeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) RuntimeTypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeTypeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) RuntimeTypeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeTypeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) ScriptLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) ScriptLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Tag() SyntheticsMonitorTagList {
	var returns SyntheticsMonitorTagList
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) TagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) TreatRedirectAsFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"treatRedirectAsFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) TreatRedirectAsFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"treatRedirectAsFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) UseUnsupportedLegacyRuntime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useUnsupportedLegacyRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) UseUnsupportedLegacyRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useUnsupportedLegacyRuntimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) ValidationString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) ValidationStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validationStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) VerifySsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifySsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsMonitor) VerifySslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifySslInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/synthetics_monitor newrelic_synthetics_monitor} Resource.
func NewSyntheticsMonitor(scope constructs.Construct, id *string, config *SyntheticsMonitorConfig) SyntheticsMonitor {
	_init_.Initialize()

	if err := validateNewSyntheticsMonitorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsMonitor{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.syntheticsMonitor.SyntheticsMonitor",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/synthetics_monitor newrelic_synthetics_monitor} Resource.
func NewSyntheticsMonitor_Override(s SyntheticsMonitor, scope constructs.Construct, id *string, config *SyntheticsMonitorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.syntheticsMonitor.SyntheticsMonitor",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetAccountId(val *float64) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetBrowsers(val *[]*string) {
	if err := j.validateSetBrowsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"browsers",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetBypassHeadRequest(val interface{}) {
	if err := j.validateSetBypassHeadRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bypassHeadRequest",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetDeviceOrientation(val *string) {
	if err := j.validateSetDeviceOrientationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceOrientation",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetDevices(val *[]*string) {
	if err := j.validateSetDevicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"devices",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetDeviceType(val *string) {
	if err := j.validateSetDeviceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceType",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetEnableScreenshotOnFailureAndScript(val interface{}) {
	if err := j.validateSetEnableScreenshotOnFailureAndScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableScreenshotOnFailureAndScript",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetLocationsPrivate(val *[]*string) {
	if err := j.validateSetLocationsPrivateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationsPrivate",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetLocationsPublic(val *[]*string) {
	if err := j.validateSetLocationsPublicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationsPublic",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetPeriod(val *string) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetRuntimeType(val *string) {
	if err := j.validateSetRuntimeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeType",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetRuntimeTypeVersion(val *string) {
	if err := j.validateSetRuntimeTypeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeTypeVersion",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetScriptLanguage(val *string) {
	if err := j.validateSetScriptLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptLanguage",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetTreatRedirectAsFailure(val interface{}) {
	if err := j.validateSetTreatRedirectAsFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"treatRedirectAsFailure",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetUseUnsupportedLegacyRuntime(val interface{}) {
	if err := j.validateSetUseUnsupportedLegacyRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useUnsupportedLegacyRuntime",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetValidationString(val *string) {
	if err := j.validateSetValidationStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationString",
		val,
	)
}

func (j *jsiiProxy_SyntheticsMonitor)SetVerifySsl(val interface{}) {
	if err := j.validateSetVerifySslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifySsl",
		val,
	)
}

// Generates CDKTF code for importing a SyntheticsMonitor resource upon running "cdktf plan <stack-name>".
func SyntheticsMonitor_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSyntheticsMonitor_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.syntheticsMonitor.SyntheticsMonitor",
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
func SyntheticsMonitor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsMonitor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.syntheticsMonitor.SyntheticsMonitor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsMonitor_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsMonitor_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.syntheticsMonitor.SyntheticsMonitor",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsMonitor_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsMonitor_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.syntheticsMonitor.SyntheticsMonitor",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SyntheticsMonitor_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.syntheticsMonitor.SyntheticsMonitor",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SyntheticsMonitor) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SyntheticsMonitor) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SyntheticsMonitor) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SyntheticsMonitor) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsMonitor) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SyntheticsMonitor) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SyntheticsMonitor) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SyntheticsMonitor) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SyntheticsMonitor) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SyntheticsMonitor) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SyntheticsMonitor) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SyntheticsMonitor) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMonitor) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SyntheticsMonitor) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsMonitor) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsMonitor) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SyntheticsMonitor) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsMonitor) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SyntheticsMonitor) PutCustomHeader(value interface{}) {
	if err := s.validatePutCustomHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomHeader",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsMonitor) PutTag(value interface{}) {
	if err := s.validatePutTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTag",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetAccountId() {
	_jsii_.InvokeVoid(
		s,
		"resetAccountId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetBrowsers() {
	_jsii_.InvokeVoid(
		s,
		"resetBrowsers",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetBypassHeadRequest() {
	_jsii_.InvokeVoid(
		s,
		"resetBypassHeadRequest",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetCustomHeader() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomHeader",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetDeviceOrientation() {
	_jsii_.InvokeVoid(
		s,
		"resetDeviceOrientation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetDevices() {
	_jsii_.InvokeVoid(
		s,
		"resetDevices",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetDeviceType() {
	_jsii_.InvokeVoid(
		s,
		"resetDeviceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetEnableScreenshotOnFailureAndScript() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableScreenshotOnFailureAndScript",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetLocationsPrivate() {
	_jsii_.InvokeVoid(
		s,
		"resetLocationsPrivate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetLocationsPublic() {
	_jsii_.InvokeVoid(
		s,
		"resetLocationsPublic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetPeriod() {
	_jsii_.InvokeVoid(
		s,
		"resetPeriod",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetRuntimeType() {
	_jsii_.InvokeVoid(
		s,
		"resetRuntimeType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetRuntimeTypeVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetRuntimeTypeVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetScriptLanguage() {
	_jsii_.InvokeVoid(
		s,
		"resetScriptLanguage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetTag() {
	_jsii_.InvokeVoid(
		s,
		"resetTag",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetTreatRedirectAsFailure() {
	_jsii_.InvokeVoid(
		s,
		"resetTreatRedirectAsFailure",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetUri() {
	_jsii_.InvokeVoid(
		s,
		"resetUri",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetUseUnsupportedLegacyRuntime() {
	_jsii_.InvokeVoid(
		s,
		"resetUseUnsupportedLegacyRuntime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetValidationString() {
	_jsii_.InvokeVoid(
		s,
		"resetValidationString",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) ResetVerifySsl() {
	_jsii_.InvokeVoid(
		s,
		"resetVerifySsl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsMonitor) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMonitor) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMonitor) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMonitor) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMonitor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsMonitor) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

