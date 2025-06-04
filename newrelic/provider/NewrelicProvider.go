// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.62.0/docs newrelic}.
type NewrelicProvider interface {
	cdktf.TerraformProvider
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
	AdminApiKey() *string
	SetAdminApiKey(val *string)
	AdminApiKeyInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	ApiKey() *string
	SetApiKey(val *string)
	ApiKeyInput() *string
	ApiUrl() *string
	SetApiUrl(val *string)
	ApiUrlInput() *string
	CacertFile() *string
	SetCacertFile(val *string)
	CacertFileInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	InfrastructureApiUrl() *string
	SetInfrastructureApiUrl(val *string)
	InfrastructureApiUrlInput() *string
	InsecureSkipVerify() interface{}
	SetInsecureSkipVerify(val interface{})
	InsecureSkipVerifyInput() interface{}
	InsightsInsertKey() *string
	SetInsightsInsertKey(val *string)
	InsightsInsertKeyInput() *string
	InsightsInsertUrl() *string
	SetInsightsInsertUrl(val *string)
	InsightsInsertUrlInput() *string
	InsightsQueryUrl() *string
	SetInsightsQueryUrl(val *string)
	InsightsQueryUrlInput() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	NerdgraphApiUrl() *string
	SetNerdgraphApiUrl(val *string)
	NerdgraphApiUrlInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SyntheticsApiUrl() *string
	SetSyntheticsApiUrl(val *string)
	SyntheticsApiUrlInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAdminApiKey()
	ResetAlias()
	ResetApiUrl()
	ResetCacertFile()
	ResetInfrastructureApiUrl()
	ResetInsecureSkipVerify()
	ResetInsightsInsertKey()
	ResetInsightsInsertUrl()
	ResetInsightsQueryUrl()
	ResetNerdgraphApiUrl()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSyntheticsApiUrl()
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

// The jsii proxy struct for NewrelicProvider
type jsiiProxy_NewrelicProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_NewrelicProvider) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) AdminApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) AdminApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) ApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) ApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) ApiUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) CacertFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacertFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) CacertFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacertFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) InfrastructureApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureApiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) InfrastructureApiUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureApiUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) InsecureSkipVerify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSkipVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) InsecureSkipVerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureSkipVerifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) InsightsInsertKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insightsInsertKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) InsightsInsertKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insightsInsertKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) InsightsInsertUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insightsInsertUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) InsightsInsertUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insightsInsertUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) InsightsQueryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insightsQueryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) InsightsQueryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insightsQueryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) NerdgraphApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nerdgraphApiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) NerdgraphApiUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nerdgraphApiUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) SyntheticsApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syntheticsApiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) SyntheticsApiUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"syntheticsApiUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewrelicProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.62.0/docs newrelic} Resource.
func NewNewrelicProvider(scope constructs.Construct, id *string, config *NewrelicProviderConfig) NewrelicProvider {
	_init_.Initialize()

	if err := validateNewNewrelicProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NewrelicProvider{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.provider.NewrelicProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.62.0/docs newrelic} Resource.
func NewNewrelicProvider_Override(n NewrelicProvider, scope constructs.Construct, id *string, config *NewrelicProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.provider.NewrelicProvider",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NewrelicProvider)SetAccountId(val *float64) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_NewrelicProvider)SetAdminApiKey(val *string) {
	_jsii_.Set(
		j,
		"adminApiKey",
		val,
	)
}

func (j *jsiiProxy_NewrelicProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_NewrelicProvider)SetApiKey(val *string) {
	_jsii_.Set(
		j,
		"apiKey",
		val,
	)
}

func (j *jsiiProxy_NewrelicProvider)SetApiUrl(val *string) {
	_jsii_.Set(
		j,
		"apiUrl",
		val,
	)
}

func (j *jsiiProxy_NewrelicProvider)SetCacertFile(val *string) {
	_jsii_.Set(
		j,
		"cacertFile",
		val,
	)
}

func (j *jsiiProxy_NewrelicProvider)SetInfrastructureApiUrl(val *string) {
	_jsii_.Set(
		j,
		"infrastructureApiUrl",
		val,
	)
}

func (j *jsiiProxy_NewrelicProvider)SetInsecureSkipVerify(val interface{}) {
	if err := j.validateSetInsecureSkipVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureSkipVerify",
		val,
	)
}

func (j *jsiiProxy_NewrelicProvider)SetInsightsInsertKey(val *string) {
	_jsii_.Set(
		j,
		"insightsInsertKey",
		val,
	)
}

func (j *jsiiProxy_NewrelicProvider)SetInsightsInsertUrl(val *string) {
	_jsii_.Set(
		j,
		"insightsInsertUrl",
		val,
	)
}

func (j *jsiiProxy_NewrelicProvider)SetInsightsQueryUrl(val *string) {
	_jsii_.Set(
		j,
		"insightsQueryUrl",
		val,
	)
}

func (j *jsiiProxy_NewrelicProvider)SetNerdgraphApiUrl(val *string) {
	_jsii_.Set(
		j,
		"nerdgraphApiUrl",
		val,
	)
}

func (j *jsiiProxy_NewrelicProvider)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_NewrelicProvider)SetSyntheticsApiUrl(val *string) {
	_jsii_.Set(
		j,
		"syntheticsApiUrl",
		val,
	)
}

// Generates CDKTF code for importing a NewrelicProvider resource upon running "cdktf plan <stack-name>".
func NewrelicProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNewrelicProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.provider.NewrelicProvider",
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
func NewrelicProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNewrelicProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.provider.NewrelicProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NewrelicProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNewrelicProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.provider.NewrelicProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NewrelicProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNewrelicProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.provider.NewrelicProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NewrelicProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.provider.NewrelicProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NewrelicProvider) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NewrelicProvider) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NewrelicProvider) ResetAdminApiKey() {
	_jsii_.InvokeVoid(
		n,
		"resetAdminApiKey",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewrelicProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		n,
		"resetAlias",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewrelicProvider) ResetApiUrl() {
	_jsii_.InvokeVoid(
		n,
		"resetApiUrl",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewrelicProvider) ResetCacertFile() {
	_jsii_.InvokeVoid(
		n,
		"resetCacertFile",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewrelicProvider) ResetInfrastructureApiUrl() {
	_jsii_.InvokeVoid(
		n,
		"resetInfrastructureApiUrl",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewrelicProvider) ResetInsecureSkipVerify() {
	_jsii_.InvokeVoid(
		n,
		"resetInsecureSkipVerify",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewrelicProvider) ResetInsightsInsertKey() {
	_jsii_.InvokeVoid(
		n,
		"resetInsightsInsertKey",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewrelicProvider) ResetInsightsInsertUrl() {
	_jsii_.InvokeVoid(
		n,
		"resetInsightsInsertUrl",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewrelicProvider) ResetInsightsQueryUrl() {
	_jsii_.InvokeVoid(
		n,
		"resetInsightsQueryUrl",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewrelicProvider) ResetNerdgraphApiUrl() {
	_jsii_.InvokeVoid(
		n,
		"resetNerdgraphApiUrl",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewrelicProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewrelicProvider) ResetRegion() {
	_jsii_.InvokeVoid(
		n,
		"resetRegion",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewrelicProvider) ResetSyntheticsApiUrl() {
	_jsii_.InvokeVoid(
		n,
		"resetSyntheticsApiUrl",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewrelicProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewrelicProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewrelicProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewrelicProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewrelicProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewrelicProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

