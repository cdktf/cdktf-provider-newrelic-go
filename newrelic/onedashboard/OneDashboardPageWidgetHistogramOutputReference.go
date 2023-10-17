// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v11/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v11/onedashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OneDashboardPageWidgetHistogramOutputReference interface {
	cdktf.ComplexObject
	Colors() OneDashboardPageWidgetHistogramColorsList
	ColorsInput() interface{}
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
	FacetShowOtherSeries() interface{}
	SetFacetShowOtherSeries(val interface{})
	FacetShowOtherSeriesInput() interface{}
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
	LegendEnabled() interface{}
	SetLegendEnabled(val interface{})
	LegendEnabledInput() interface{}
	NrqlQuery() OneDashboardPageWidgetHistogramNrqlQueryList
	NrqlQueryInput() interface{}
	NullValues() OneDashboardPageWidgetHistogramNullValuesList
	NullValuesInput() interface{}
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
	Units() OneDashboardPageWidgetHistogramUnitsList
	UnitsInput() interface{}
	Width() *float64
	SetWidth(val *float64)
	WidthInput() *float64
	YAxisLeftMax() *float64
	SetYAxisLeftMax(val *float64)
	YAxisLeftMaxInput() *float64
	YAxisLeftMin() *float64
	SetYAxisLeftMin(val *float64)
	YAxisLeftMinInput() *float64
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
	PutColors(value interface{})
	PutNrqlQuery(value interface{})
	PutNullValues(value interface{})
	PutUnits(value interface{})
	ResetColors()
	ResetFacetShowOtherSeries()
	ResetHeight()
	ResetIgnoreTimeRange()
	ResetLegendEnabled()
	ResetNullValues()
	ResetUnits()
	ResetWidth()
	ResetYAxisLeftMax()
	ResetYAxisLeftMin()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OneDashboardPageWidgetHistogramOutputReference
type jsiiProxy_OneDashboardPageWidgetHistogramOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) Colors() OneDashboardPageWidgetHistogramColorsList {
	var returns OneDashboardPageWidgetHistogramColorsList
	_jsii_.Get(
		j,
		"colors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ColorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"colorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) Column() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ColumnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) FacetShowOtherSeries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facetShowOtherSeries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) FacetShowOtherSeriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facetShowOtherSeriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) HeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) IgnoreTimeRange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTimeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) IgnoreTimeRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTimeRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) LegendEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"legendEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) LegendEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"legendEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) NrqlQuery() OneDashboardPageWidgetHistogramNrqlQueryList {
	var returns OneDashboardPageWidgetHistogramNrqlQueryList
	_jsii_.Get(
		j,
		"nrqlQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) NrqlQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nrqlQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) NullValues() OneDashboardPageWidgetHistogramNullValuesList {
	var returns OneDashboardPageWidgetHistogramNullValuesList
	_jsii_.Get(
		j,
		"nullValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) NullValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nullValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) Row() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"row",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) RowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) Units() OneDashboardPageWidgetHistogramUnitsList {
	var returns OneDashboardPageWidgetHistogramUnitsList
	_jsii_.Get(
		j,
		"units",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) UnitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) WidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) YAxisLeftMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisLeftMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) YAxisLeftMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisLeftMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) YAxisLeftMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisLeftMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) YAxisLeftMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisLeftMinInput",
		&returns,
	)
	return returns
}


func NewOneDashboardPageWidgetHistogramOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OneDashboardPageWidgetHistogramOutputReference {
	_init_.Initialize()

	if err := validateNewOneDashboardPageWidgetHistogramOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OneDashboardPageWidgetHistogramOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.oneDashboard.OneDashboardPageWidgetHistogramOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOneDashboardPageWidgetHistogramOutputReference_Override(o OneDashboardPageWidgetHistogramOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.oneDashboard.OneDashboardPageWidgetHistogramOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference)SetColumn(val *float64) {
	if err := j.validateSetColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"column",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference)SetFacetShowOtherSeries(val interface{}) {
	if err := j.validateSetFacetShowOtherSeriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"facetShowOtherSeries",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference)SetHeight(val *float64) {
	if err := j.validateSetHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"height",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference)SetIgnoreTimeRange(val interface{}) {
	if err := j.validateSetIgnoreTimeRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreTimeRange",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference)SetLegendEnabled(val interface{}) {
	if err := j.validateSetLegendEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendEnabled",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference)SetRow(val *float64) {
	if err := j.validateSetRowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"row",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference)SetWidth(val *float64) {
	if err := j.validateSetWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"width",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference)SetYAxisLeftMax(val *float64) {
	if err := j.validateSetYAxisLeftMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yAxisLeftMax",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference)SetYAxisLeftMin(val *float64) {
	if err := j.validateSetYAxisLeftMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yAxisLeftMin",
		val,
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) PutColors(value interface{}) {
	if err := o.validatePutColorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putColors",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) PutNrqlQuery(value interface{}) {
	if err := o.validatePutNrqlQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNrqlQuery",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) PutNullValues(value interface{}) {
	if err := o.validatePutNullValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNullValues",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) PutUnits(value interface{}) {
	if err := o.validatePutUnitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putUnits",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ResetColors() {
	_jsii_.InvokeVoid(
		o,
		"resetColors",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ResetFacetShowOtherSeries() {
	_jsii_.InvokeVoid(
		o,
		"resetFacetShowOtherSeries",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ResetHeight() {
	_jsii_.InvokeVoid(
		o,
		"resetHeight",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ResetIgnoreTimeRange() {
	_jsii_.InvokeVoid(
		o,
		"resetIgnoreTimeRange",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ResetLegendEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetLegendEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ResetNullValues() {
	_jsii_.InvokeVoid(
		o,
		"resetNullValues",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ResetUnits() {
	_jsii_.InvokeVoid(
		o,
		"resetUnits",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ResetWidth() {
	_jsii_.InvokeVoid(
		o,
		"resetWidth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ResetYAxisLeftMax() {
	_jsii_.InvokeVoid(
		o,
		"resetYAxisLeftMax",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ResetYAxisLeftMin() {
	_jsii_.InvokeVoid(
		o,
		"resetYAxisLeftMin",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OneDashboardPageWidgetHistogramOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

