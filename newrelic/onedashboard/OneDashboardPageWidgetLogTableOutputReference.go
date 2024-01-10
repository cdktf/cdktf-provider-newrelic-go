// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/onedashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OneDashboardPageWidgetLogTableOutputReference interface {
	cdktf.ComplexObject
	Colors() OneDashboardPageWidgetLogTableColorsList
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
	NrqlQuery() OneDashboardPageWidgetLogTableNrqlQueryList
	NrqlQueryInput() interface{}
	NullValues() OneDashboardPageWidgetLogTableNullValuesList
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
	Units() OneDashboardPageWidgetLogTableUnitsList
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

// The jsii proxy struct for OneDashboardPageWidgetLogTableOutputReference
type jsiiProxy_OneDashboardPageWidgetLogTableOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) Colors() OneDashboardPageWidgetLogTableColorsList {
	var returns OneDashboardPageWidgetLogTableColorsList
	_jsii_.Get(
		j,
		"colors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ColorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"colorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) Column() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ColumnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) FacetShowOtherSeries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facetShowOtherSeries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) FacetShowOtherSeriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facetShowOtherSeriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) HeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) IgnoreTimeRange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTimeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) IgnoreTimeRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTimeRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) LegendEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"legendEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) LegendEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"legendEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) NrqlQuery() OneDashboardPageWidgetLogTableNrqlQueryList {
	var returns OneDashboardPageWidgetLogTableNrqlQueryList
	_jsii_.Get(
		j,
		"nrqlQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) NrqlQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nrqlQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) NullValues() OneDashboardPageWidgetLogTableNullValuesList {
	var returns OneDashboardPageWidgetLogTableNullValuesList
	_jsii_.Get(
		j,
		"nullValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) NullValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nullValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) Row() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"row",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) RowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) Units() OneDashboardPageWidgetLogTableUnitsList {
	var returns OneDashboardPageWidgetLogTableUnitsList
	_jsii_.Get(
		j,
		"units",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) UnitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) WidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) YAxisLeftMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisLeftMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) YAxisLeftMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisLeftMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) YAxisLeftMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisLeftMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) YAxisLeftMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisLeftMinInput",
		&returns,
	)
	return returns
}


func NewOneDashboardPageWidgetLogTableOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OneDashboardPageWidgetLogTableOutputReference {
	_init_.Initialize()

	if err := validateNewOneDashboardPageWidgetLogTableOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OneDashboardPageWidgetLogTableOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.oneDashboard.OneDashboardPageWidgetLogTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOneDashboardPageWidgetLogTableOutputReference_Override(o OneDashboardPageWidgetLogTableOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.oneDashboard.OneDashboardPageWidgetLogTableOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference)SetColumn(val *float64) {
	if err := j.validateSetColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"column",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference)SetFacetShowOtherSeries(val interface{}) {
	if err := j.validateSetFacetShowOtherSeriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"facetShowOtherSeries",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference)SetHeight(val *float64) {
	if err := j.validateSetHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"height",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference)SetIgnoreTimeRange(val interface{}) {
	if err := j.validateSetIgnoreTimeRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreTimeRange",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference)SetLegendEnabled(val interface{}) {
	if err := j.validateSetLegendEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendEnabled",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference)SetRow(val *float64) {
	if err := j.validateSetRowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"row",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference)SetWidth(val *float64) {
	if err := j.validateSetWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"width",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference)SetYAxisLeftMax(val *float64) {
	if err := j.validateSetYAxisLeftMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yAxisLeftMax",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference)SetYAxisLeftMin(val *float64) {
	if err := j.validateSetYAxisLeftMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yAxisLeftMin",
		val,
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) PutColors(value interface{}) {
	if err := o.validatePutColorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putColors",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) PutNrqlQuery(value interface{}) {
	if err := o.validatePutNrqlQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNrqlQuery",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) PutNullValues(value interface{}) {
	if err := o.validatePutNullValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNullValues",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) PutUnits(value interface{}) {
	if err := o.validatePutUnitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putUnits",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ResetColors() {
	_jsii_.InvokeVoid(
		o,
		"resetColors",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ResetFacetShowOtherSeries() {
	_jsii_.InvokeVoid(
		o,
		"resetFacetShowOtherSeries",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ResetHeight() {
	_jsii_.InvokeVoid(
		o,
		"resetHeight",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ResetIgnoreTimeRange() {
	_jsii_.InvokeVoid(
		o,
		"resetIgnoreTimeRange",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ResetLegendEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetLegendEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ResetNullValues() {
	_jsii_.InvokeVoid(
		o,
		"resetNullValues",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ResetUnits() {
	_jsii_.InvokeVoid(
		o,
		"resetUnits",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ResetWidth() {
	_jsii_.InvokeVoid(
		o,
		"resetWidth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ResetYAxisLeftMax() {
	_jsii_.InvokeVoid(
		o,
		"resetYAxisLeftMax",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ResetYAxisLeftMin() {
	_jsii_.InvokeVoid(
		o,
		"resetYAxisLeftMin",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OneDashboardPageWidgetLogTableOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

