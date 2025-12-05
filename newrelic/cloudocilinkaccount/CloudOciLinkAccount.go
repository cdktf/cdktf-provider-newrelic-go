// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudocilinkaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/cloudocilinkaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.3/docs/resources/cloud_oci_link_account newrelic_cloud_oci_link_account}.
type CloudOciLinkAccount interface {
	cdktf.TerraformResource
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CompartmentOcid() *string
	SetCompartmentOcid(val *string)
	CompartmentOcidInput() *string
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
	IngestVaultOcid() *string
	SetIngestVaultOcid(val *string)
	IngestVaultOcidInput() *string
	InstrumentationType() *string
	SetInstrumentationType(val *string)
	InstrumentationTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoggingStackOcid() *string
	SetLoggingStackOcid(val *string)
	LoggingStackOcidInput() *string
	MetricStackOcid() *string
	SetMetricStackOcid(val *string)
	MetricStackOcidInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OciClientId() *string
	SetOciClientId(val *string)
	OciClientIdInput() *string
	OciClientSecret() *string
	SetOciClientSecret(val *string)
	OciClientSecretInput() *string
	OciDomainUrl() *string
	SetOciDomainUrl(val *string)
	OciDomainUrlInput() *string
	OciHomeRegion() *string
	SetOciHomeRegion(val *string)
	OciHomeRegionInput() *string
	OciRegion() *string
	SetOciRegion(val *string)
	OciRegionInput() *string
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
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserVaultOcid() *string
	SetUserVaultOcid(val *string)
	UserVaultOcidInput() *string
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
	ResetAccountId()
	ResetId()
	ResetIngestVaultOcid()
	ResetInstrumentationType()
	ResetLoggingStackOcid()
	ResetMetricStackOcid()
	ResetOciRegion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetUserVaultOcid()
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

// The jsii proxy struct for CloudOciLinkAccount
type jsiiProxy_CloudOciLinkAccount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudOciLinkAccount) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) CompartmentOcid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compartmentOcid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) CompartmentOcidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compartmentOcidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) IngestVaultOcid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingestVaultOcid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) IngestVaultOcidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingestVaultOcidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) InstrumentationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instrumentationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) InstrumentationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instrumentationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) LoggingStackOcid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingStackOcid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) LoggingStackOcidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingStackOcidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) MetricStackOcid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricStackOcid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) MetricStackOcidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricStackOcidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) OciClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) OciClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) OciClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) OciClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociClientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) OciDomainUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociDomainUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) OciDomainUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociDomainUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) OciHomeRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociHomeRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) OciHomeRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociHomeRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) OciRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) OciRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) UserVaultOcid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userVaultOcid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudOciLinkAccount) UserVaultOcidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userVaultOcidInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.3/docs/resources/cloud_oci_link_account newrelic_cloud_oci_link_account} Resource.
func NewCloudOciLinkAccount(scope constructs.Construct, id *string, config *CloudOciLinkAccountConfig) CloudOciLinkAccount {
	_init_.Initialize()

	if err := validateNewCloudOciLinkAccountParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudOciLinkAccount{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.cloudOciLinkAccount.CloudOciLinkAccount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.3/docs/resources/cloud_oci_link_account newrelic_cloud_oci_link_account} Resource.
func NewCloudOciLinkAccount_Override(c CloudOciLinkAccount, scope constructs.Construct, id *string, config *CloudOciLinkAccountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.cloudOciLinkAccount.CloudOciLinkAccount",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetAccountId(val *float64) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetCompartmentOcid(val *string) {
	if err := j.validateSetCompartmentOcidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compartmentOcid",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetIngestVaultOcid(val *string) {
	if err := j.validateSetIngestVaultOcidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingestVaultOcid",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetInstrumentationType(val *string) {
	if err := j.validateSetInstrumentationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instrumentationType",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetLoggingStackOcid(val *string) {
	if err := j.validateSetLoggingStackOcidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingStackOcid",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetMetricStackOcid(val *string) {
	if err := j.validateSetMetricStackOcidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricStackOcid",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetOciClientId(val *string) {
	if err := j.validateSetOciClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ociClientId",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetOciClientSecret(val *string) {
	if err := j.validateSetOciClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ociClientSecret",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetOciDomainUrl(val *string) {
	if err := j.validateSetOciDomainUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ociDomainUrl",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetOciHomeRegion(val *string) {
	if err := j.validateSetOciHomeRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ociHomeRegion",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetOciRegion(val *string) {
	if err := j.validateSetOciRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ociRegion",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_CloudOciLinkAccount)SetUserVaultOcid(val *string) {
	if err := j.validateSetUserVaultOcidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userVaultOcid",
		val,
	)
}

// Generates CDKTF code for importing a CloudOciLinkAccount resource upon running "cdktf plan <stack-name>".
func CloudOciLinkAccount_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudOciLinkAccount_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudOciLinkAccount.CloudOciLinkAccount",
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
func CloudOciLinkAccount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudOciLinkAccount_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudOciLinkAccount.CloudOciLinkAccount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudOciLinkAccount_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudOciLinkAccount_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudOciLinkAccount.CloudOciLinkAccount",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudOciLinkAccount_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudOciLinkAccount_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudOciLinkAccount.CloudOciLinkAccount",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudOciLinkAccount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.cloudOciLinkAccount.CloudOciLinkAccount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) ResetAccountId() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) ResetIngestVaultOcid() {
	_jsii_.InvokeVoid(
		c,
		"resetIngestVaultOcid",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) ResetInstrumentationType() {
	_jsii_.InvokeVoid(
		c,
		"resetInstrumentationType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) ResetLoggingStackOcid() {
	_jsii_.InvokeVoid(
		c,
		"resetLoggingStackOcid",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) ResetMetricStackOcid() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricStackOcid",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) ResetOciRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetOciRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) ResetUserVaultOcid() {
	_jsii_.InvokeVoid(
		c,
		"resetUserVaultOcid",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudOciLinkAccount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudOciLinkAccount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

