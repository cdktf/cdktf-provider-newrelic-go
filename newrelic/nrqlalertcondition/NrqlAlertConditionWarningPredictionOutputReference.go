// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nrqlalertcondition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/nrqlalertcondition/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NrqlAlertConditionWarningPredictionOutputReference interface {
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
	InternalValue() *NrqlAlertConditionWarningPrediction
	SetInternalValue(val *NrqlAlertConditionWarningPrediction)
	PredictBy() *float64
	SetPredictBy(val *float64)
	PredictByInput() *float64
	PreferPredictionViolation() interface{}
	SetPreferPredictionViolation(val interface{})
	PreferPredictionViolationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetPredictBy()
	ResetPreferPredictionViolation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NrqlAlertConditionWarningPredictionOutputReference
type jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) InternalValue() *NrqlAlertConditionWarningPrediction {
	var returns *NrqlAlertConditionWarningPrediction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) PredictBy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"predictBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) PredictByInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"predictByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) PreferPredictionViolation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preferPredictionViolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) PreferPredictionViolationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preferPredictionViolationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNrqlAlertConditionWarningPredictionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NrqlAlertConditionWarningPredictionOutputReference {
	_init_.Initialize()

	if err := validateNewNrqlAlertConditionWarningPredictionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionWarningPredictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNrqlAlertConditionWarningPredictionOutputReference_Override(n NrqlAlertConditionWarningPredictionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionWarningPredictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference)SetInternalValue(val *NrqlAlertConditionWarningPrediction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference)SetPredictBy(val *float64) {
	if err := j.validateSetPredictByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictBy",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference)SetPreferPredictionViolation(val interface{}) {
	if err := j.validateSetPreferPredictionViolationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferPredictionViolation",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) ResetPredictBy() {
	_jsii_.InvokeVoid(
		n,
		"resetPredictBy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) ResetPreferPredictionViolation() {
	_jsii_.InvokeVoid(
		n,
		"resetPreferPredictionViolation",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NrqlAlertConditionWarningPredictionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

