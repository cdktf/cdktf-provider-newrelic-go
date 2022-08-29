// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-newrelic-go/newrelic/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-newrelic-go/newrelic/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition newrelic_nrql_alert_condition}.
type NrqlAlertCondition interface {
	cdktf.TerraformResource
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
	AggregationDelay() *string
	SetAggregationDelay(val *string)
	AggregationDelayInput() *string
	AggregationMethod() *string
	SetAggregationMethod(val *string)
	AggregationMethodInput() *string
	AggregationTimer() *string
	SetAggregationTimer(val *string)
	AggregationTimerInput() *string
	AggregationWindow() *float64
	SetAggregationWindow(val *float64)
	AggregationWindowInput() *float64
	BaselineDirection() *string
	SetBaselineDirection(val *string)
	BaselineDirectionInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloseViolationsOnExpiration() interface{}
	SetCloseViolationsOnExpiration(val interface{})
	CloseViolationsOnExpirationInput() interface{}
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
	Critical() NrqlAlertConditionCriticalOutputReference
	CriticalInput() *NrqlAlertConditionCritical
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
	ExpirationDuration() *float64
	SetExpirationDuration(val *float64)
	ExpirationDurationInput() *float64
	FillOption() *string
	SetFillOption(val *string)
	FillOptionInput() *string
	FillValue() *float64
	SetFillValue(val *float64)
	FillValueInput() *float64
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
	Nrql() NrqlAlertConditionNrqlOutputReference
	NrqlInput() *NrqlAlertConditionNrql
	OpenViolationOnExpiration() interface{}
	SetOpenViolationOnExpiration(val interface{})
	OpenViolationOnExpirationInput() interface{}
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
	SlideBy() *float64
	SetSlideBy(val *float64)
	SlideByInput() *float64
	Term() NrqlAlertConditionTermList
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
	ValueFunction() *string
	SetValueFunction(val *string)
	ValueFunctionInput() *string
	ViolationTimeLimit() *string
	SetViolationTimeLimit(val *string)
	ViolationTimeLimitInput() *string
	ViolationTimeLimitSeconds() *float64
	SetViolationTimeLimitSeconds(val *float64)
	ViolationTimeLimitSecondsInput() *float64
	Warning() NrqlAlertConditionWarningOutputReference
	WarningInput() *NrqlAlertConditionWarning
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
	PutCritical(value *NrqlAlertConditionCritical)
	PutNrql(value *NrqlAlertConditionNrql)
	PutTerm(value interface{})
	PutWarning(value *NrqlAlertConditionWarning)
	ResetAccountId()
	ResetAggregationDelay()
	ResetAggregationMethod()
	ResetAggregationTimer()
	ResetAggregationWindow()
	ResetBaselineDirection()
	ResetCloseViolationsOnExpiration()
	ResetCritical()
	ResetDescription()
	ResetEnabled()
	ResetExpirationDuration()
	ResetFillOption()
	ResetFillValue()
	ResetId()
	ResetOpenViolationOnExpiration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRunbookUrl()
	ResetSlideBy()
	ResetTerm()
	ResetType()
	ResetValueFunction()
	ResetViolationTimeLimit()
	ResetViolationTimeLimitSeconds()
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

// The jsii proxy struct for NrqlAlertCondition
type jsiiProxy_NrqlAlertCondition struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NrqlAlertCondition) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) AggregationDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) AggregationDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) AggregationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) AggregationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) AggregationTimer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationTimer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) AggregationTimerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aggregationTimerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) AggregationWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"aggregationWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) AggregationWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"aggregationWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) BaselineDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baselineDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) BaselineDirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baselineDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) CloseViolationsOnExpiration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"closeViolationsOnExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) CloseViolationsOnExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"closeViolationsOnExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Critical() NrqlAlertConditionCriticalOutputReference {
	var returns NrqlAlertConditionCriticalOutputReference
	_jsii_.Get(
		j,
		"critical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) CriticalInput() *NrqlAlertConditionCritical {
	var returns *NrqlAlertConditionCritical
	_jsii_.Get(
		j,
		"criticalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) EntityGuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityGuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) ExpirationDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) ExpirationDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expirationDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) FillOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fillOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) FillOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fillOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) FillValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fillValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) FillValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fillValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Nrql() NrqlAlertConditionNrqlOutputReference {
	var returns NrqlAlertConditionNrqlOutputReference
	_jsii_.Get(
		j,
		"nrql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) NrqlInput() *NrqlAlertConditionNrql {
	var returns *NrqlAlertConditionNrql
	_jsii_.Get(
		j,
		"nrqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) OpenViolationOnExpiration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openViolationOnExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) OpenViolationOnExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"openViolationOnExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) PolicyId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) PolicyIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"policyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) RunbookUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runbookUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) RunbookUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runbookUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) SlideBy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slideBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) SlideByInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slideByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Term() NrqlAlertConditionTermList {
	var returns NrqlAlertConditionTermList
	_jsii_.Get(
		j,
		"term",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) TermInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"termInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) ValueFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) ValueFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) ViolationTimeLimit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"violationTimeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) ViolationTimeLimitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"violationTimeLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) ViolationTimeLimitSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"violationTimeLimitSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) ViolationTimeLimitSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"violationTimeLimitSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) Warning() NrqlAlertConditionWarningOutputReference {
	var returns NrqlAlertConditionWarningOutputReference
	_jsii_.Get(
		j,
		"warning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertCondition) WarningInput() *NrqlAlertConditionWarning {
	var returns *NrqlAlertConditionWarning
	_jsii_.Get(
		j,
		"warningInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition newrelic_nrql_alert_condition} Resource.
func NewNrqlAlertCondition(scope constructs.Construct, id *string, config *NrqlAlertConditionConfig) NrqlAlertCondition {
	_init_.Initialize()

	j := jsiiProxy_NrqlAlertCondition{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.NrqlAlertCondition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/newrelic/r/nrql_alert_condition newrelic_nrql_alert_condition} Resource.
func NewNrqlAlertCondition_Override(n NrqlAlertCondition, scope constructs.Construct, id *string, config *NrqlAlertConditionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.NrqlAlertCondition",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetAccountId(val *float64) {
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetAggregationDelay(val *string) {
	_jsii_.Set(
		j,
		"aggregationDelay",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetAggregationMethod(val *string) {
	_jsii_.Set(
		j,
		"aggregationMethod",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetAggregationTimer(val *string) {
	_jsii_.Set(
		j,
		"aggregationTimer",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetAggregationWindow(val *float64) {
	_jsii_.Set(
		j,
		"aggregationWindow",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetBaselineDirection(val *string) {
	_jsii_.Set(
		j,
		"baselineDirection",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetCloseViolationsOnExpiration(val interface{}) {
	_jsii_.Set(
		j,
		"closeViolationsOnExpiration",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetExpirationDuration(val *float64) {
	_jsii_.Set(
		j,
		"expirationDuration",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetFillOption(val *string) {
	_jsii_.Set(
		j,
		"fillOption",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetFillValue(val *float64) {
	_jsii_.Set(
		j,
		"fillValue",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetOpenViolationOnExpiration(val interface{}) {
	_jsii_.Set(
		j,
		"openViolationOnExpiration",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetPolicyId(val *float64) {
	_jsii_.Set(
		j,
		"policyId",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetRunbookUrl(val *string) {
	_jsii_.Set(
		j,
		"runbookUrl",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetSlideBy(val *float64) {
	_jsii_.Set(
		j,
		"slideBy",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetValueFunction(val *string) {
	_jsii_.Set(
		j,
		"valueFunction",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetViolationTimeLimit(val *string) {
	_jsii_.Set(
		j,
		"violationTimeLimit",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertCondition) SetViolationTimeLimitSeconds(val *float64) {
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
func NrqlAlertCondition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.NrqlAlertCondition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NrqlAlertCondition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.NrqlAlertCondition",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NrqlAlertCondition) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NrqlAlertCondition) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertCondition) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertCondition) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertCondition) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertCondition) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertCondition) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertCondition) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertCondition) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertCondition) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertCondition) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertCondition) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NrqlAlertCondition) PutCritical(value *NrqlAlertConditionCritical) {
	_jsii_.InvokeVoid(
		n,
		"putCritical",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NrqlAlertCondition) PutNrql(value *NrqlAlertConditionNrql) {
	_jsii_.InvokeVoid(
		n,
		"putNrql",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NrqlAlertCondition) PutTerm(value interface{}) {
	_jsii_.InvokeVoid(
		n,
		"putTerm",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NrqlAlertCondition) PutWarning(value *NrqlAlertConditionWarning) {
	_jsii_.InvokeVoid(
		n,
		"putWarning",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetAccountId() {
	_jsii_.InvokeVoid(
		n,
		"resetAccountId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetAggregationDelay() {
	_jsii_.InvokeVoid(
		n,
		"resetAggregationDelay",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetAggregationMethod() {
	_jsii_.InvokeVoid(
		n,
		"resetAggregationMethod",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetAggregationTimer() {
	_jsii_.InvokeVoid(
		n,
		"resetAggregationTimer",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetAggregationWindow() {
	_jsii_.InvokeVoid(
		n,
		"resetAggregationWindow",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetBaselineDirection() {
	_jsii_.InvokeVoid(
		n,
		"resetBaselineDirection",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetCloseViolationsOnExpiration() {
	_jsii_.InvokeVoid(
		n,
		"resetCloseViolationsOnExpiration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetCritical() {
	_jsii_.InvokeVoid(
		n,
		"resetCritical",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetExpirationDuration() {
	_jsii_.InvokeVoid(
		n,
		"resetExpirationDuration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetFillOption() {
	_jsii_.InvokeVoid(
		n,
		"resetFillOption",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetFillValue() {
	_jsii_.InvokeVoid(
		n,
		"resetFillValue",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetOpenViolationOnExpiration() {
	_jsii_.InvokeVoid(
		n,
		"resetOpenViolationOnExpiration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetRunbookUrl() {
	_jsii_.InvokeVoid(
		n,
		"resetRunbookUrl",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetSlideBy() {
	_jsii_.InvokeVoid(
		n,
		"resetSlideBy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetTerm() {
	_jsii_.InvokeVoid(
		n,
		"resetTerm",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetType() {
	_jsii_.InvokeVoid(
		n,
		"resetType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetValueFunction() {
	_jsii_.InvokeVoid(
		n,
		"resetValueFunction",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetViolationTimeLimit() {
	_jsii_.InvokeVoid(
		n,
		"resetViolationTimeLimit",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetViolationTimeLimitSeconds() {
	_jsii_.InvokeVoid(
		n,
		"resetViolationTimeLimitSeconds",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) ResetWarning() {
	_jsii_.InvokeVoid(
		n,
		"resetWarning",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertCondition) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertCondition) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertCondition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertCondition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
