// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-newrelic-go/newrelic/v2/jsii"

	"github.com/hashicorp/cdktf-provider-newrelic-go/newrelic/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OneDashboardPageWidgetStackedBarOutputReference interface {
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
	NrqlQuery() OneDashboardPageWidgetStackedBarNrqlQueryList
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
	ResetHeight()
	ResetIgnoreTimeRange()
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

// The jsii proxy struct for OneDashboardPageWidgetStackedBarOutputReference
type jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) Column() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) ColumnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) HeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) IgnoreTimeRange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTimeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) IgnoreTimeRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTimeRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) NrqlQuery() OneDashboardPageWidgetStackedBarNrqlQueryList {
	var returns OneDashboardPageWidgetStackedBarNrqlQueryList
	_jsii_.Get(
		j,
		"nrqlQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) NrqlQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nrqlQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) Row() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"row",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) RowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) WidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widthInput",
		&returns,
	)
	return returns
}


func NewOneDashboardPageWidgetStackedBarOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OneDashboardPageWidgetStackedBarOutputReference {
	_init_.Initialize()

	j := jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.OneDashboardPageWidgetStackedBarOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOneDashboardPageWidgetStackedBarOutputReference_Override(o OneDashboardPageWidgetStackedBarOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.OneDashboardPageWidgetStackedBarOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) SetColumn(val *float64) {
	_jsii_.Set(
		j,
		"column",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) SetHeight(val *float64) {
	_jsii_.Set(
		j,
		"height",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) SetIgnoreTimeRange(val interface{}) {
	_jsii_.Set(
		j,
		"ignoreTimeRange",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) SetRow(val *float64) {
	_jsii_.Set(
		j,
		"row",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) SetTitle(val *string) {
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) SetWidth(val *float64) {
	_jsii_.Set(
		j,
		"width",
		val,
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) PutNrqlQuery(value interface{}) {
	_jsii_.InvokeVoid(
		o,
		"putNrqlQuery",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) ResetHeight() {
	_jsii_.InvokeVoid(
		o,
		"resetHeight",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) ResetIgnoreTimeRange() {
	_jsii_.InvokeVoid(
		o,
		"resetIgnoreTimeRange",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) ResetWidth() {
	_jsii_.InvokeVoid(
		o,
		"resetWidth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetStackedBarOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
