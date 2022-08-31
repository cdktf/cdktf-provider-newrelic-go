// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-newrelic-go/newrelic/v2/jsii"

	"github.com/hashicorp/cdktf-provider-newrelic-go/newrelic/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OneDashboardPageWidgetBarOutputReference interface {
	cdktf.ComplexObject
	Column() *float64
	SetColumn(val *float64)
	ColumnInput() *float64
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
	FilterCurrentDashboard() interface{}
	SetFilterCurrentDashboard(val interface{})
	FilterCurrentDashboardInput() interface{}
	// Experimental.
	Fqn() *string
	Height() *float64
	SetHeight(val *float64)
	HeightInput() *float64
	Id() *string
	IgnoreTimeRange() interface{}
	SetIgnoreTimeRange(val interface{})
	IgnoreTimeRangeInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LinkedEntityGuids() *[]*string
	SetLinkedEntityGuids(val *[]*string)
	LinkedEntityGuidsInput() *[]*string
	NrqlQuery() OneDashboardPageWidgetBarNrqlQueryList
	NrqlQueryInput() interface{}
	Row() *float64
	SetRow(val *float64)
	RowInput() *float64
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
	Width() *float64
	SetWidth(val *float64)
	WidthInput() *float64
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
	PutNrqlQuery(value interface{})
	ResetFilterCurrentDashboard()
	ResetHeight()
	ResetIgnoreTimeRange()
	ResetLinkedEntityGuids()
	ResetWidth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OneDashboardPageWidgetBarOutputReference
type jsiiProxy_OneDashboardPageWidgetBarOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) Column() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) ColumnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) FilterCurrentDashboard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterCurrentDashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) FilterCurrentDashboardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterCurrentDashboardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) HeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) IgnoreTimeRange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTimeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) IgnoreTimeRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTimeRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) LinkedEntityGuids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"linkedEntityGuids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) LinkedEntityGuidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"linkedEntityGuidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) NrqlQuery() OneDashboardPageWidgetBarNrqlQueryList {
	var returns OneDashboardPageWidgetBarNrqlQueryList
	_jsii_.Get(
		j,
		"nrqlQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) NrqlQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nrqlQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) Row() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"row",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) RowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference) WidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widthInput",
		&returns,
	)
	return returns
}


func NewOneDashboardPageWidgetBarOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OneDashboardPageWidgetBarOutputReference {
	_init_.Initialize()

	if err := validateNewOneDashboardPageWidgetBarOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OneDashboardPageWidgetBarOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.OneDashboardPageWidgetBarOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOneDashboardPageWidgetBarOutputReference_Override(o OneDashboardPageWidgetBarOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.OneDashboardPageWidgetBarOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference)SetColumn(val *float64) {
	if err := j.validateSetColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"column",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference)SetFilterCurrentDashboard(val interface{}) {
	if err := j.validateSetFilterCurrentDashboardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filterCurrentDashboard",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference)SetHeight(val *float64) {
	if err := j.validateSetHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"height",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference)SetIgnoreTimeRange(val interface{}) {
	if err := j.validateSetIgnoreTimeRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreTimeRange",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference)SetLinkedEntityGuids(val *[]*string) {
	if err := j.validateSetLinkedEntityGuidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkedEntityGuids",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference)SetRow(val *float64) {
	if err := j.validateSetRowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"row",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetBarOutputReference)SetWidth(val *float64) {
	if err := j.validateSetWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"width",
		val,
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) PutNrqlQuery(value interface{}) {
	if err := o.validatePutNrqlQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNrqlQuery",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) ResetFilterCurrentDashboard() {
	_jsii_.InvokeVoid(
		o,
		"resetFilterCurrentDashboard",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) ResetHeight() {
	_jsii_.InvokeVoid(
		o,
		"resetHeight",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) ResetIgnoreTimeRange() {
	_jsii_.InvokeVoid(
		o,
		"resetIgnoreTimeRange",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) ResetLinkedEntityGuids() {
	_jsii_.InvokeVoid(
		o,
		"resetLinkedEntityGuids",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) ResetWidth() {
	_jsii_.InvokeVoid(
		o,
		"resetWidth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OneDashboardPageWidgetBarOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

