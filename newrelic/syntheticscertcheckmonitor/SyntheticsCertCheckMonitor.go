// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syntheticscertcheckmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/syntheticscertcheckmonitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.0/docs/resources/synthetics_cert_check_monitor newrelic_synthetics_cert_check_monitor}.
type SyntheticsCertCheckMonitor interface {
	cdktf.TerraformResource
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateExpiration() *float64
	SetCertificateExpiration(val *float64)
	CertificateExpirationInput() *float64
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
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
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
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	Tag() SyntheticsCertCheckMonitorTagList
	TagInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutTag(value interface{})
	ResetAccountId()
	ResetId()
	ResetLocationsPrivate()
	ResetLocationsPublic()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRuntimeType()
	ResetRuntimeTypeVersion()
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

// The jsii proxy struct for SyntheticsCertCheckMonitor
type jsiiProxy_SyntheticsCertCheckMonitor struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) CertificateExpiration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"certificateExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) CertificateExpirationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"certificateExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) LocationsPrivate() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsPrivate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) LocationsPrivateInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsPrivateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) LocationsPublic() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsPublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) LocationsPublicInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsPublicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) MonitorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) Period() *string {
	var returns *string
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) PeriodInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) PeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) RuntimeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) RuntimeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) RuntimeTypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeTypeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) RuntimeTypeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeTypeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) Tag() SyntheticsCertCheckMonitorTagList {
	var returns SyntheticsCertCheckMonitorTagList
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) TagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) UseUnsupportedLegacyRuntime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useUnsupportedLegacyRuntime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor) UseUnsupportedLegacyRuntimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useUnsupportedLegacyRuntimeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.0/docs/resources/synthetics_cert_check_monitor newrelic_synthetics_cert_check_monitor} Resource.
func NewSyntheticsCertCheckMonitor(scope constructs.Construct, id *string, config *SyntheticsCertCheckMonitorConfig) SyntheticsCertCheckMonitor {
	_init_.Initialize()

	if err := validateNewSyntheticsCertCheckMonitorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsCertCheckMonitor{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.syntheticsCertCheckMonitor.SyntheticsCertCheckMonitor",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.0/docs/resources/synthetics_cert_check_monitor newrelic_synthetics_cert_check_monitor} Resource.
func NewSyntheticsCertCheckMonitor_Override(s SyntheticsCertCheckMonitor, scope constructs.Construct, id *string, config *SyntheticsCertCheckMonitorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.syntheticsCertCheckMonitor.SyntheticsCertCheckMonitor",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetAccountId(val *float64) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetCertificateExpiration(val *float64) {
	if err := j.validateSetCertificateExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateExpiration",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetLocationsPrivate(val *[]*string) {
	if err := j.validateSetLocationsPrivateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationsPrivate",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetLocationsPublic(val *[]*string) {
	if err := j.validateSetLocationsPublicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationsPublic",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetPeriod(val *string) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetRuntimeType(val *string) {
	if err := j.validateSetRuntimeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeType",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetRuntimeTypeVersion(val *string) {
	if err := j.validateSetRuntimeTypeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeTypeVersion",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_SyntheticsCertCheckMonitor)SetUseUnsupportedLegacyRuntime(val interface{}) {
	if err := j.validateSetUseUnsupportedLegacyRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useUnsupportedLegacyRuntime",
		val,
	)
}

// Generates CDKTF code for importing a SyntheticsCertCheckMonitor resource upon running "cdktf plan <stack-name>".
func SyntheticsCertCheckMonitor_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSyntheticsCertCheckMonitor_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.syntheticsCertCheckMonitor.SyntheticsCertCheckMonitor",
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
func SyntheticsCertCheckMonitor_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsCertCheckMonitor_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.syntheticsCertCheckMonitor.SyntheticsCertCheckMonitor",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsCertCheckMonitor_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsCertCheckMonitor_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.syntheticsCertCheckMonitor.SyntheticsCertCheckMonitor",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SyntheticsCertCheckMonitor_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSyntheticsCertCheckMonitor_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.syntheticsCertCheckMonitor.SyntheticsCertCheckMonitor",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SyntheticsCertCheckMonitor_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.syntheticsCertCheckMonitor.SyntheticsCertCheckMonitor",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SyntheticsCertCheckMonitor) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsCertCheckMonitor) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SyntheticsCertCheckMonitor) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SyntheticsCertCheckMonitor) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SyntheticsCertCheckMonitor) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SyntheticsCertCheckMonitor) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SyntheticsCertCheckMonitor) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SyntheticsCertCheckMonitor) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SyntheticsCertCheckMonitor) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SyntheticsCertCheckMonitor) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) PutTag(value interface{}) {
	if err := s.validatePutTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTag",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) ResetAccountId() {
	_jsii_.InvokeVoid(
		s,
		"resetAccountId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) ResetLocationsPrivate() {
	_jsii_.InvokeVoid(
		s,
		"resetLocationsPrivate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) ResetLocationsPublic() {
	_jsii_.InvokeVoid(
		s,
		"resetLocationsPublic",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) ResetRuntimeType() {
	_jsii_.InvokeVoid(
		s,
		"resetRuntimeType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) ResetRuntimeTypeVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetRuntimeTypeVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) ResetTag() {
	_jsii_.InvokeVoid(
		s,
		"resetTag",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) ResetUseUnsupportedLegacyRuntime() {
	_jsii_.InvokeVoid(
		s,
		"resetUseUnsupportedLegacyRuntime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCertCheckMonitor) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

