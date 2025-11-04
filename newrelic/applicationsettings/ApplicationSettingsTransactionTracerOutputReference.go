// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/applicationsettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationSettingsTransactionTracerOutputReference interface {
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
	ExplainQueryPlans() ApplicationSettingsTransactionTracerExplainQueryPlansList
	ExplainQueryPlansInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Sql() ApplicationSettingsTransactionTracerSqlOutputReference
	SqlInput() *ApplicationSettingsTransactionTracerSql
	StackTraceThresholdValue() *float64
	SetStackTraceThresholdValue(val *float64)
	StackTraceThresholdValueInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransactionThresholdType() *string
	SetTransactionThresholdType(val *string)
	TransactionThresholdTypeInput() *string
	TransactionThresholdValue() *float64
	SetTransactionThresholdValue(val *float64)
	TransactionThresholdValueInput() *float64
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutExplainQueryPlans(value interface{})
	PutSql(value *ApplicationSettingsTransactionTracerSql)
	ResetExplainQueryPlans()
	ResetSql()
	ResetStackTraceThresholdValue()
	ResetTransactionThresholdType()
	ResetTransactionThresholdValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationSettingsTransactionTracerOutputReference
type jsiiProxy_ApplicationSettingsTransactionTracerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) ExplainQueryPlans() ApplicationSettingsTransactionTracerExplainQueryPlansList {
	var returns ApplicationSettingsTransactionTracerExplainQueryPlansList
	_jsii_.Get(
		j,
		"explainQueryPlans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) ExplainQueryPlansInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"explainQueryPlansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) Sql() ApplicationSettingsTransactionTracerSqlOutputReference {
	var returns ApplicationSettingsTransactionTracerSqlOutputReference
	_jsii_.Get(
		j,
		"sql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) SqlInput() *ApplicationSettingsTransactionTracerSql {
	var returns *ApplicationSettingsTransactionTracerSql
	_jsii_.Get(
		j,
		"sqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) StackTraceThresholdValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stackTraceThresholdValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) StackTraceThresholdValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stackTraceThresholdValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) TransactionThresholdType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transactionThresholdType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) TransactionThresholdTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transactionThresholdTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) TransactionThresholdValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transactionThresholdValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) TransactionThresholdValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transactionThresholdValueInput",
		&returns,
	)
	return returns
}


func NewApplicationSettingsTransactionTracerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApplicationSettingsTransactionTracerOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationSettingsTransactionTracerOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationSettingsTransactionTracerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsTransactionTracerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApplicationSettingsTransactionTracerOutputReference_Override(a ApplicationSettingsTransactionTracerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsTransactionTracerOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference)SetStackTraceThresholdValue(val *float64) {
	if err := j.validateSetStackTraceThresholdValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackTraceThresholdValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference)SetTransactionThresholdType(val *string) {
	if err := j.validateSetTransactionThresholdTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionThresholdType",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference)SetTransactionThresholdValue(val *float64) {
	if err := j.validateSetTransactionThresholdValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionThresholdValue",
		val,
	)
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) PutExplainQueryPlans(value interface{}) {
	if err := a.validatePutExplainQueryPlansParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExplainQueryPlans",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) PutSql(value *ApplicationSettingsTransactionTracerSql) {
	if err := a.validatePutSqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSql",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) ResetExplainQueryPlans() {
	_jsii_.InvokeVoid(
		a,
		"resetExplainQueryPlans",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) ResetSql() {
	_jsii_.InvokeVoid(
		a,
		"resetSql",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) ResetStackTraceThresholdValue() {
	_jsii_.InvokeVoid(
		a,
		"resetStackTraceThresholdValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) ResetTransactionThresholdType() {
	_jsii_.InvokeVoid(
		a,
		"resetTransactionThresholdType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) ResetTransactionThresholdValue() {
	_jsii_.InvokeVoid(
		a,
		"resetTransactionThresholdValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsTransactionTracerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

