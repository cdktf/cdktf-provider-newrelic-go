// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/applicationsettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationSettingsErrorCollectorOutputReference interface {
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
	ExpectedErrorClasses() *[]*string
	SetExpectedErrorClasses(val *[]*string)
	ExpectedErrorClassesInput() *[]*string
	ExpectedErrorCodes() *[]*string
	SetExpectedErrorCodes(val *[]*string)
	ExpectedErrorCodesInput() *[]*string
	// Experimental.
	Fqn() *string
	IgnoredErrorClasses() *[]*string
	SetIgnoredErrorClasses(val *[]*string)
	IgnoredErrorClassesInput() *[]*string
	IgnoredErrorCodes() *[]*string
	SetIgnoredErrorCodes(val *[]*string)
	IgnoredErrorCodesInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetExpectedErrorClasses()
	ResetExpectedErrorCodes()
	ResetIgnoredErrorClasses()
	ResetIgnoredErrorCodes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationSettingsErrorCollectorOutputReference
type jsiiProxy_ApplicationSettingsErrorCollectorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) ExpectedErrorClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"expectedErrorClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) ExpectedErrorClassesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"expectedErrorClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) ExpectedErrorCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"expectedErrorCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) ExpectedErrorCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"expectedErrorCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) IgnoredErrorClasses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoredErrorClasses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) IgnoredErrorClassesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoredErrorClassesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) IgnoredErrorCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoredErrorCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) IgnoredErrorCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoredErrorCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationSettingsErrorCollectorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApplicationSettingsErrorCollectorOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationSettingsErrorCollectorOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationSettingsErrorCollectorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsErrorCollectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApplicationSettingsErrorCollectorOutputReference_Override(a ApplicationSettingsErrorCollectorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.applicationSettings.ApplicationSettingsErrorCollectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference)SetExpectedErrorClasses(val *[]*string) {
	if err := j.validateSetExpectedErrorClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedErrorClasses",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference)SetExpectedErrorCodes(val *[]*string) {
	if err := j.validateSetExpectedErrorCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedErrorCodes",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference)SetIgnoredErrorClasses(val *[]*string) {
	if err := j.validateSetIgnoredErrorClassesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoredErrorClasses",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference)SetIgnoredErrorCodes(val *[]*string) {
	if err := j.validateSetIgnoredErrorCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoredErrorCodes",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) ResetExpectedErrorClasses() {
	_jsii_.InvokeVoid(
		a,
		"resetExpectedErrorClasses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) ResetExpectedErrorCodes() {
	_jsii_.InvokeVoid(
		a,
		"resetExpectedErrorCodes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) ResetIgnoredErrorClasses() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoredErrorClasses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) ResetIgnoredErrorCodes() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoredErrorCodes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationSettingsErrorCollectorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

