// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package onedashboard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/onedashboard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OneDashboardPageWidgetLineYAxisRightOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *OneDashboardPageWidgetLineYAxisRight
	SetInternalValue(val *OneDashboardPageWidgetLineYAxisRight)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	YAxisRightMax() *float64
	SetYAxisRightMax(val *float64)
	YAxisRightMaxInput() *float64
	YAxisRightMin() *float64
	SetYAxisRightMin(val *float64)
	YAxisRightMinInput() *float64
	YAxisRightSeries() *[]*string
	SetYAxisRightSeries(val *[]*string)
	YAxisRightSeriesInput() *[]*string
	YAxisRightZero() interface{}
	SetYAxisRightZero(val interface{})
	YAxisRightZeroInput() interface{}
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
	ResetYAxisRightMax()
	ResetYAxisRightMin()
	ResetYAxisRightSeries()
	ResetYAxisRightZero()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OneDashboardPageWidgetLineYAxisRightOutputReference
type jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) InternalValue() *OneDashboardPageWidgetLineYAxisRight {
	var returns *OneDashboardPageWidgetLineYAxisRight
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) YAxisRightMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisRightMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) YAxisRightMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisRightMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) YAxisRightMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisRightMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) YAxisRightMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yAxisRightMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) YAxisRightSeries() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"yAxisRightSeries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) YAxisRightSeriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"yAxisRightSeriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) YAxisRightZero() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"yAxisRightZero",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) YAxisRightZeroInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"yAxisRightZeroInput",
		&returns,
	)
	return returns
}


func NewOneDashboardPageWidgetLineYAxisRightOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OneDashboardPageWidgetLineYAxisRightOutputReference {
	_init_.Initialize()

	if err := validateNewOneDashboardPageWidgetLineYAxisRightOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.oneDashboard.OneDashboardPageWidgetLineYAxisRightOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOneDashboardPageWidgetLineYAxisRightOutputReference_Override(o OneDashboardPageWidgetLineYAxisRightOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.oneDashboard.OneDashboardPageWidgetLineYAxisRightOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference)SetInternalValue(val *OneDashboardPageWidgetLineYAxisRight) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference)SetYAxisRightMax(val *float64) {
	if err := j.validateSetYAxisRightMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yAxisRightMax",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference)SetYAxisRightMin(val *float64) {
	if err := j.validateSetYAxisRightMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yAxisRightMin",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference)SetYAxisRightSeries(val *[]*string) {
	if err := j.validateSetYAxisRightSeriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yAxisRightSeries",
		val,
	)
}

func (j *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference)SetYAxisRightZero(val interface{}) {
	if err := j.validateSetYAxisRightZeroParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"yAxisRightZero",
		val,
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) ResetYAxisRightMax() {
	_jsii_.InvokeVoid(
		o,
		"resetYAxisRightMax",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) ResetYAxisRightMin() {
	_jsii_.InvokeVoid(
		o,
		"resetYAxisRightMin",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) ResetYAxisRightSeries() {
	_jsii_.InvokeVoid(
		o,
		"resetYAxisRightSeries",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) ResetYAxisRightZero() {
	_jsii_.InvokeVoid(
		o,
		"resetYAxisRightZero",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OneDashboardPageWidgetLineYAxisRightOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

