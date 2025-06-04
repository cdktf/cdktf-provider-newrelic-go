// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package nrqlalertcondition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/nrqlalertcondition/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NrqlAlertConditionCriticalPredictionOutputReference interface {
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
	InternalValue() *NrqlAlertConditionCriticalPrediction
	SetInternalValue(val *NrqlAlertConditionCriticalPrediction)
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

// The jsii proxy struct for NrqlAlertConditionCriticalPredictionOutputReference
type jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) InternalValue() *NrqlAlertConditionCriticalPrediction {
	var returns *NrqlAlertConditionCriticalPrediction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) PredictBy() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"predictBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) PredictByInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"predictByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) PreferPredictionViolation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preferPredictionViolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) PreferPredictionViolationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preferPredictionViolationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNrqlAlertConditionCriticalPredictionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NrqlAlertConditionCriticalPredictionOutputReference {
	_init_.Initialize()

	if err := validateNewNrqlAlertConditionCriticalPredictionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionCriticalPredictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNrqlAlertConditionCriticalPredictionOutputReference_Override(n NrqlAlertConditionCriticalPredictionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.nrqlAlertCondition.NrqlAlertConditionCriticalPredictionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference)SetInternalValue(val *NrqlAlertConditionCriticalPrediction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference)SetPredictBy(val *float64) {
	if err := j.validateSetPredictByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictBy",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference)SetPreferPredictionViolation(val interface{}) {
	if err := j.validateSetPreferPredictionViolationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferPredictionViolation",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) ResetPredictBy() {
	_jsii_.InvokeVoid(
		n,
		"resetPredictBy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) ResetPreferPredictionViolation() {
	_jsii_.InvokeVoid(
		n,
		"resetPreferPredictionViolation",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NrqlAlertConditionCriticalPredictionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

