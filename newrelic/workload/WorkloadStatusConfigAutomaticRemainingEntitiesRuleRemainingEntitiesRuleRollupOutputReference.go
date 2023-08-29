// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v10/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v10/workload/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference interface {
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
	GroupBy() *string
	SetGroupBy(val *string)
	GroupByInput() *string
	InternalValue() *WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollup
	SetInternalValue(val *WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollup)
	Strategy() *string
	SetStrategy(val *string)
	StrategyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThresholdType() *string
	SetThresholdType(val *string)
	ThresholdTypeInput() *string
	ThresholdValue() *float64
	SetThresholdValue(val *float64)
	ThresholdValueInput() *float64
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
	ResetThresholdType()
	ResetThresholdValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference
type jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) GroupBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) GroupByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) InternalValue() *WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollup {
	var returns *WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) ThresholdType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) ThresholdTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thresholdTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) ThresholdValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) ThresholdValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdValueInput",
		&returns,
	)
	return returns
}


func NewWorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference {
	_init_.Initialize()

	if err := validateNewWorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.workload.WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference_Override(w WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.workload.WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference)SetGroupBy(val *string) {
	if err := j.validateSetGroupByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupBy",
		val,
	)
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference)SetInternalValue(val *WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollup) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference)SetStrategy(val *string) {
	if err := j.validateSetStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference)SetThresholdType(val *string) {
	if err := j.validateSetThresholdTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdType",
		val,
	)
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference)SetThresholdValue(val *float64) {
	if err := j.validateSetThresholdValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"thresholdValue",
		val,
	)
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) ResetThresholdType() {
	_jsii_.InvokeVoid(
		w,
		"resetThresholdType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) ResetThresholdValue() {
	_jsii_.InvokeVoid(
		w,
		"resetThresholdValue",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRemainingEntitiesRuleRemainingEntitiesRuleRollupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

