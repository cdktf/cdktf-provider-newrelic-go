// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-newrelic-go/newrelic/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-newrelic-go/newrelic/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/newrelic/r/workflow newrelic_workflow}.
type Workflow interface {
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DestinationConfiguration() WorkflowDestinationConfigurationList
	DestinationConfigurationInput() interface{}
	DestinationsEnabled() interface{}
	SetDestinationsEnabled(val interface{})
	DestinationsEnabledInput() interface{}
	Enrichments() WorkflowEnrichmentsOutputReference
	EnrichmentsEnabled() interface{}
	SetEnrichmentsEnabled(val interface{})
	EnrichmentsEnabledInput() interface{}
	EnrichmentsInput() *WorkflowEnrichments
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
	IssuesFilter() WorkflowIssuesFilterOutputReference
	IssuesFilterInput() *WorkflowIssuesFilter
	LastRun() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MutingRulesHandling() *string
	SetMutingRulesHandling(val *string)
	MutingRulesHandlingInput() *string
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
	WorkflowEnabled() interface{}
	SetWorkflowEnabled(val interface{})
	WorkflowEnabledInput() interface{}
	WorkflowId() *string
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
	PutDestinationConfiguration(value interface{})
	PutEnrichments(value *WorkflowEnrichments)
	PutIssuesFilter(value *WorkflowIssuesFilter)
	ResetAccountId()
	ResetDestinationsEnabled()
	ResetEnrichments()
	ResetEnrichmentsEnabled()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetWorkflowEnabled()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Workflow
type jsiiProxy_Workflow struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Workflow) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) DestinationConfiguration() WorkflowDestinationConfigurationList {
	var returns WorkflowDestinationConfigurationList
	_jsii_.Get(
		j,
		"destinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) DestinationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) DestinationsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) DestinationsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Enrichments() WorkflowEnrichmentsOutputReference {
	var returns WorkflowEnrichmentsOutputReference
	_jsii_.Get(
		j,
		"enrichments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) EnrichmentsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enrichmentsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) EnrichmentsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enrichmentsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) EnrichmentsInput() *WorkflowEnrichments {
	var returns *WorkflowEnrichments
	_jsii_.Get(
		j,
		"enrichmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) IssuesFilter() WorkflowIssuesFilterOutputReference {
	var returns WorkflowIssuesFilterOutputReference
	_jsii_.Get(
		j,
		"issuesFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) IssuesFilterInput() *WorkflowIssuesFilter {
	var returns *WorkflowIssuesFilter
	_jsii_.Get(
		j,
		"issuesFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) LastRun() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) MutingRulesHandling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mutingRulesHandling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) MutingRulesHandlingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mutingRulesHandlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) WorkflowEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) WorkflowEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Workflow) WorkflowId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowId",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/newrelic/r/workflow newrelic_workflow} Resource.
func NewWorkflow(scope constructs.Construct, id *string, config *WorkflowConfig) Workflow {
	_init_.Initialize()

	j := jsiiProxy_Workflow{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.Workflow",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/newrelic/r/workflow newrelic_workflow} Resource.
func NewWorkflow_Override(w Workflow, scope constructs.Construct, id *string, config *WorkflowConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.Workflow",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_Workflow) SetAccountId(val *float64) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetDestinationsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"destinationsEnabled",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetEnrichmentsEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enrichmentsEnabled",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetMutingRulesHandling(val *string) {
	_jsii_.Set(
		j,
		"mutingRulesHandling",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Workflow) SetWorkflowEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"workflowEnabled",
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
func Workflow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.Workflow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Workflow_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.Workflow",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_Workflow) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_Workflow) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_Workflow) PutDestinationConfiguration(value interface{}) {
	_jsii_.InvokeVoid(
		w,
		"putDestinationConfiguration",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Workflow) PutEnrichments(value *WorkflowEnrichments) {
	_jsii_.InvokeVoid(
		w,
		"putEnrichments",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Workflow) PutIssuesFilter(value *WorkflowIssuesFilter) {
	_jsii_.InvokeVoid(
		w,
		"putIssuesFilter",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Workflow) ResetAccountId() {
	_jsii_.InvokeVoid(
		w,
		"resetAccountId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workflow) ResetDestinationsEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetDestinationsEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workflow) ResetEnrichments() {
	_jsii_.InvokeVoid(
		w,
		"resetEnrichments",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workflow) ResetEnrichmentsEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetEnrichmentsEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workflow) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workflow) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workflow) ResetWorkflowEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetWorkflowEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Workflow) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Workflow) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

