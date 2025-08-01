// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/onedashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OneDashboardPageWidgetLineOutputReference interface {
	cdktf.ComplexObject
	Colors() OneDashboardPageWidgetLineColorsList
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
	DataFormat() OneDashboardPageWidgetLineDataFormatList
	DataFormatInput() interface{}
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
	InitialSorting() OneDashboardPageWidgetLineInitialSortingOutputReference
	InitialSortingInput() *OneDashboardPageWidgetLineInitialSorting
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsLabelVisible() interface{}
	SetIsLabelVisible(val interface{})
	IsLabelVisibleInput() interface{}
	LegendEnabled() interface{}
	SetLegendEnabled(val interface{})
	LegendEnabledInput() interface{}
	NrqlQuery() OneDashboardPageWidgetLineNrqlQueryList
	NrqlQueryInput() interface{}
	NullValues() OneDashboardPageWidgetLineNullValuesList
	NullValuesInput() interface{}
	RefreshRate() *string
	SetRefreshRate(val *string)
	RefreshRateInput() *string
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
	Threshold() OneDashboardPageWidgetLineThresholdList
	ThresholdInput() interface{}
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	Tooltip() OneDashboardPageWidgetLineTooltipOutputReference
	TooltipInput() *OneDashboardPageWidgetLineTooltip
	Units() OneDashboardPageWidgetLineUnitsList
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
	YAxisLeftZero() interface{}
	SetYAxisLeftZero(val interface{})
	YAxisLeftZeroInput() interface{}
	YAxisRight() OneDashboardPageWidgetLineYAxisRightOutputReference
	YAxisRightInput() *OneDashboardPageWidgetLineYAxisRight
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
	PutDataFormat(value interface{})
	PutInitialSorting(value *OneDashboardPageWidgetLineInitialSorting)
	PutNrqlQuery(value interface{})
	PutNullValues(value interface{})
	PutThreshold(value interface{})
	PutTooltip(value *OneDashboardPageWidgetLineTooltip)
	PutUnits(value interface{})
	PutYAxisRight(value *OneDashboardPageWidgetLineYAxisRight)
	ResetColors()
	ResetDataFormat()
	ResetFacetShowOtherSeries()
	ResetHeight()
	ResetIgnoreTimeRange()
	ResetInitialSorting()
	ResetIsLabelVisible()
	ResetLegendEnabled()
	ResetNullValues()
	ResetRefreshRate()
	ResetThreshold()
	ResetTooltip()
	ResetUnits()
	ResetWidth()
	ResetYAxisLeftMax()
	ResetYAxisLeftMin()
	ResetYAxisLeftZero()
	ResetYAxisRight()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OneDashboardPageWidgetLineOutputReference
type jsiiProxy_OneDashboardPageWidgetLineOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) Colors() OneDashboardPageWidgetLineColorsList {
	var returns OneDashboardPageWidgetLineColorsList
	_jsii_.Get(
		j,
		"colors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ColorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"colorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) Column() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ColumnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) DataFormat() OneDashboardPageWidgetLineDataFormatList {
	var returns OneDashboardPageWidgetLineDataFormatList
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) DataFormatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) FacetShowOtherSeries() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facetShowOtherSeries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) FacetShowOtherSeriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facetShowOtherSeriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) HeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) IgnoreTimeRange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTimeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) IgnoreTimeRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTimeRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) InitialSorting() OneDashboardPageWidgetLineInitialSortingOutputReference {
	var returns OneDashboardPageWidgetLineInitialSortingOutputReference
	_jsii_.Get(
		j,
		"initialSorting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) InitialSortingInput() *OneDashboardPageWidgetLineInitialSorting {
	var returns *OneDashboardPageWidgetLineInitialSorting
	_jsii_.Get(
		j,
		"initialSortingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) IsLabelVisible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLabelVisible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) IsLabelVisibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLabelVisibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) LegendEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"legendEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) LegendEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"legendEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) NrqlQuery() OneDashboardPageWidgetLineNrqlQueryList {
	var returns OneDashboardPageWidgetLineNrqlQueryList
	_jsii_.Get(
		j,
		"nrqlQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) NrqlQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nrqlQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) NullValues() OneDashboardPageWidgetLineNullValuesList {
	var returns OneDashboardPageWidgetLineNullValuesList
	_jsii_.Get(
		j,
		"nullValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) NullValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nullValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) RefreshRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) RefreshRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) Row() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"row",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) RowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) Threshold() OneDashboardPageWidgetLineThresholdList {
	var returns OneDashboardPageWidgetLineThresholdList
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ThresholdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) Tooltip() OneDashboardPageWidgetLineTooltipOutputReference {
	var returns OneDashboardPageWidgetLineTooltipOutputReference
	_jsii_.Get(
		j,
		"tooltip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) TooltipInput() *OneDashboardPageWidgetLineTooltip {
	var returns *OneDashboardPageWidgetLineTooltip
	_jsii_.Get(
		j,
		"tooltipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) Units() OneDashboardPageWidgetLineUnitsList {
	var returns OneDashboardPageWidgetLineUnitsList
	_jsii_.Get(
		j,
		"units",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) UnitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) WidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) YAxisLeftMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisLeftMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) YAxisLeftMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisLeftMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) YAxisLeftMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisLeftMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) YAxisLeftMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisLeftMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) YAxisLeftZero() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"yAxisLeftZero",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) YAxisLeftZeroInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"yAxisLeftZeroInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) YAxisRight() OneDashboardPageWidgetLineYAxisRightOutputReference {
	var returns OneDashboardPageWidgetLineYAxisRightOutputReference
	_jsii_.Get(
		j,
		"yAxisRight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference) YAxisRightInput() *OneDashboardPageWidgetLineYAxisRight {
	var returns *OneDashboardPageWidgetLineYAxisRight
	_jsii_.Get(
		j,
		"yAxisRightInput",
		&returns,
	)
	return returns
}


func NewOneDashboardPageWidgetLineOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OneDashboardPageWidgetLineOutputReference {
	_init_.Initialize()

	if err := validateNewOneDashboardPageWidgetLineOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OneDashboardPageWidgetLineOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.oneDashboard.OneDashboardPageWidgetLineOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOneDashboardPageWidgetLineOutputReference_Override(o OneDashboardPageWidgetLineOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.oneDashboard.OneDashboardPageWidgetLineOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetColumn(val *float64) {
	if err := j.validateSetColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"column",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetFacetShowOtherSeries(val interface{}) {
	if err := j.validateSetFacetShowOtherSeriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"facetShowOtherSeries",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetHeight(val *float64) {
	if err := j.validateSetHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"height",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetIgnoreTimeRange(val interface{}) {
	if err := j.validateSetIgnoreTimeRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreTimeRange",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetIsLabelVisible(val interface{}) {
	if err := j.validateSetIsLabelVisibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isLabelVisible",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetLegendEnabled(val interface{}) {
	if err := j.validateSetLegendEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"legendEnabled",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetRefreshRate(val *string) {
	if err := j.validateSetRefreshRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshRate",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetRow(val *float64) {
	if err := j.validateSetRowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"row",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetWidth(val *float64) {
	if err := j.validateSetWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"width",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetYAxisLeftMax(val *float64) {
	if err := j.validateSetYAxisLeftMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yAxisLeftMax",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetYAxisLeftMin(val *float64) {
	if err := j.validateSetYAxisLeftMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yAxisLeftMin",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineOutputReference)SetYAxisLeftZero(val interface{}) {
	if err := j.validateSetYAxisLeftZeroParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yAxisLeftZero",
		val,
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) PutColors(value interface{}) {
	if err := o.validatePutColorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putColors",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) PutDataFormat(value interface{}) {
	if err := o.validatePutDataFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDataFormat",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) PutInitialSorting(value *OneDashboardPageWidgetLineInitialSorting) {
	if err := o.validatePutInitialSortingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putInitialSorting",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) PutNrqlQuery(value interface{}) {
	if err := o.validatePutNrqlQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNrqlQuery",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) PutNullValues(value interface{}) {
	if err := o.validatePutNullValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putNullValues",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) PutThreshold(value interface{}) {
	if err := o.validatePutThresholdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putThreshold",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) PutTooltip(value *OneDashboardPageWidgetLineTooltip) {
	if err := o.validatePutTooltipParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTooltip",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) PutUnits(value interface{}) {
	if err := o.validatePutUnitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putUnits",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) PutYAxisRight(value *OneDashboardPageWidgetLineYAxisRight) {
	if err := o.validatePutYAxisRightParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putYAxisRight",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetColors() {
	_jsii_.InvokeVoid(
		o,
		"resetColors",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetDataFormat() {
	_jsii_.InvokeVoid(
		o,
		"resetDataFormat",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetFacetShowOtherSeries() {
	_jsii_.InvokeVoid(
		o,
		"resetFacetShowOtherSeries",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetHeight() {
	_jsii_.InvokeVoid(
		o,
		"resetHeight",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetIgnoreTimeRange() {
	_jsii_.InvokeVoid(
		o,
		"resetIgnoreTimeRange",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetInitialSorting() {
	_jsii_.InvokeVoid(
		o,
		"resetInitialSorting",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetIsLabelVisible() {
	_jsii_.InvokeVoid(
		o,
		"resetIsLabelVisible",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetLegendEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetLegendEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetNullValues() {
	_jsii_.InvokeVoid(
		o,
		"resetNullValues",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetRefreshRate() {
	_jsii_.InvokeVoid(
		o,
		"resetRefreshRate",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetThreshold() {
	_jsii_.InvokeVoid(
		o,
		"resetThreshold",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetTooltip() {
	_jsii_.InvokeVoid(
		o,
		"resetTooltip",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetUnits() {
	_jsii_.InvokeVoid(
		o,
		"resetUnits",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetWidth() {
	_jsii_.InvokeVoid(
		o,
		"resetWidth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetYAxisLeftMax() {
	_jsii_.InvokeVoid(
		o,
		"resetYAxisLeftMax",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetYAxisLeftMin() {
	_jsii_.InvokeVoid(
		o,
		"resetYAxisLeftMin",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetYAxisLeftZero() {
	_jsii_.InvokeVoid(
		o,
		"resetYAxisLeftZero",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ResetYAxisRight() {
	_jsii_.InvokeVoid(
		o,
		"resetYAxisRight",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

