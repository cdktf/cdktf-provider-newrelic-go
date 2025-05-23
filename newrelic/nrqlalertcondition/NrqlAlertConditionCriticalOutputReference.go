// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nrqlalertcondition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/nrqlalertcondition/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NrqlAlertConditionCriticalOutputReference interface {
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
	DisableHealthStatusReporting() interface{}
	SetDisableHealthStatusReporting(val interface{})
	DisableHealthStatusReportingInput() interface{}
	Duration() *float64
	SetDuration(val *float64)
	DurationInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *NrqlAlertConditionCritical
	SetInternalValue(val *NrqlAlertConditionCritical)
	Operator() *string
	SetOperator(val *string)
	OperatorInput() *string
	Prediction() NrqlAlertConditionCriticalPredictionOutputReference
	PredictionInput() *NrqlAlertConditionCriticalPrediction
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Threshold() *float64
	SetThreshold(val *float64)
	ThresholdDuration() *float64
	SetThresholdDuration(val *float64)
	ThresholdDurationInput() *float64
	ThresholdInput() *float64
	ThresholdOccurrences() *string
	SetThresholdOccurrences(val *string)
	ThresholdOccurrencesInput() *string
	TimeFunction() *string
	SetTimeFunction(val *string)
	TimeFunctionInput() *string
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
	PutPrediction(value *NrqlAlertConditionCriticalPrediction)
	ResetDisableHealthStatusReporting()
	ResetDuration()
	ResetOperator()
	ResetPrediction()
	ResetThresholdDuration()
	ResetThresholdOccurrences()
	ResetTimeFunction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NrqlAlertConditionCriticalOutputReference
type jsiiProxy_NrqlAlertConditionCriticalOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) DisableHealthStatusReporting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableHealthStatusReporting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) DisableHealthStatusReportingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableHealthStatusReportingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) Duration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) DurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) InternalValue() *NrqlAlertConditionCritical {
	var returns *NrqlAlertConditionCritical
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) OperatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) Prediction() NrqlAlertConditionCriticalPredictionOutputReference {
	var returns NrqlAlertConditionCriticalPredictionOutputReference
	_jsii_.Get(
		j,
		"prediction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) PredictionInput() *NrqlAlertConditionCriticalPrediction {
	var returns *NrqlAlertConditionCriticalPrediction
	_jsii_.Get(
		j,
		"predictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ThresholdDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ThresholdDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ThresholdOccurrences() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdOccurrences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ThresholdOccurrencesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdOccurrencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) TimeFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference) TimeFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeFunctionInput",
		&returns,
	)
	return returns
}


func NewNrqlAlertConditionCriticalOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NrqlAlertConditionCriticalOutputReference {
	_init_.Initialize()

	if err := validateNewNrqlAlertConditionCriticalOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NrqlAlertConditionCriticalOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionCriticalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNrqlAlertConditionCriticalOutputReference_Override(n NrqlAlertConditionCriticalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionCriticalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference)SetDisableHealthStatusReporting(val interface{}) {
	if err := j.validateSetDisableHealthStatusReportingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableHealthStatusReporting",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference)SetDuration(val *float64) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference)SetInternalValue(val *NrqlAlertConditionCritical) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference)SetThreshold(val *float64) {
	if err := j.validateSetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference)SetThresholdDuration(val *float64) {
	if err := j.validateSetThresholdDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdDuration",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference)SetThresholdOccurrences(val *string) {
	if err := j.validateSetThresholdOccurrencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdOccurrences",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalOutputReference)SetTimeFunction(val *string) {
	if err := j.validateSetTimeFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeFunction",
		val,
	)
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) PutPrediction(value *NrqlAlertConditionCriticalPrediction) {
	if err := n.validatePutPredictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putPrediction",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ResetDisableHealthStatusReporting() {
	_jsii_.InvokeVoid(
		n,
		"resetDisableHealthStatusReporting",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ResetDuration() {
	_jsii_.InvokeVoid(
		n,
		"resetDuration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ResetOperator() {
	_jsii_.InvokeVoid(
		n,
		"resetOperator",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ResetPrediction() {
	_jsii_.InvokeVoid(
		n,
		"resetPrediction",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ResetThresholdDuration() {
	_jsii_.InvokeVoid(
		n,
		"resetThresholdDuration",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ResetThresholdOccurrences() {
	_jsii_.InvokeVoid(
		n,
		"resetThresholdOccurrences",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ResetTimeFunction() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeFunction",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionCriticalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

