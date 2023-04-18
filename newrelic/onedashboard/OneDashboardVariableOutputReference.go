package onedashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v7/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v7/onedashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OneDashboardVariableOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultValues() *[]*string
	SetDefaultValues(val *[]*string)
	DefaultValuesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsMultiSelection() interface{}
	SetIsMultiSelection(val interface{})
	IsMultiSelectionInput() interface{}
	Item() OneDashboardVariableItemList
	ItemInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NrqlQuery() OneDashboardVariableNrqlQueryOutputReference
	NrqlQueryInput() *OneDashboardVariableNrqlQuery
	ReplacementStrategy() *string
	SetReplacementStrategy(val *string)
	ReplacementStrategyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutItem(value interface{})
	PutNrqlQuery(value *OneDashboardVariableNrqlQuery)
	ResetDefaultValues()
	ResetIsMultiSelection()
	ResetItem()
	ResetNrqlQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OneDashboardVariableOutputReference
type jsiiProxy_OneDashboardVariableOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) DefaultValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) DefaultValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) IsMultiSelection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMultiSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) IsMultiSelectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMultiSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) Item() OneDashboardVariableItemList {
	var returns OneDashboardVariableItemList
	_jsii_.Get(
		j,
		"item",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) ItemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"itemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) NrqlQuery() OneDashboardVariableNrqlQueryOutputReference {
	var returns OneDashboardVariableNrqlQueryOutputReference
	_jsii_.Get(
		j,
		"nrqlQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) NrqlQueryInput() *OneDashboardVariableNrqlQuery {
	var returns *OneDashboardVariableNrqlQuery
	_jsii_.Get(
		j,
		"nrqlQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) ReplacementStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacementStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) ReplacementStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacementStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardVariableOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewOneDashboardVariableOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OneDashboardVariableOutputReference {
	_init_.Initialize()

	if err := validateNewOneDashboardVariableOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OneDashboardVariableOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.oneDashboard.OneDashboardVariableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOneDashboardVariableOutputReference_Override(o OneDashboardVariableOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.oneDashboard.OneDashboardVariableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OneDashboardVariableOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OneDashboardVariableOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OneDashboardVariableOutputReference)SetDefaultValues(val *[]*string) {
	if err := j.validateSetDefaultValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultValues",
		val,
	)
}

func (j *jsiiProxy_OneDashboardVariableOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OneDashboardVariableOutputReference)SetIsMultiSelection(val interface{}) {
	if err := j.validateSetIsMultiSelectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMultiSelection",
		val,
	)
}

func (j *jsiiProxy_OneDashboardVariableOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OneDashboardVariableOutputReference)SetReplacementStrategy(val *string) {
	if err := j.validateSetReplacementStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replacementStrategy",
		val,
	)
}

func (j *jsiiProxy_OneDashboardVariableOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OneDashboardVariableOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OneDashboardVariableOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_OneDashboardVariableOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) PutItem(value interface{}) {
	if err := o.validatePutItemParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putItem",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) PutNrqlQuery(value *OneDashboardVariableNrqlQuery) {
	if err := o.validatePutNrqlQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNrqlQuery",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) ResetDefaultValues() {
	_jsii_.InvokeVoid(
		o,
		"resetDefaultValues",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) ResetIsMultiSelection() {
	_jsii_.InvokeVoid(
		o,
		"resetIsMultiSelection",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) ResetItem() {
	_jsii_.InvokeVoid(
		o,
		"resetItem",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) ResetNrqlQuery() {
	_jsii_.InvokeVoid(
		o,
		"resetNrqlQuery",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardVariableOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

