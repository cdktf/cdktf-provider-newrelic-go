// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudawsintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/cloudawsintegrations/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudAwsIntegrationsSnsOutputReference interface {
	cdktf.ComplexObject
	AwsRegions() *[]*string
	SetAwsRegions(val *[]*string)
	AwsRegionsInput() *[]*string
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
	FetchExtendedInventory() interface{}
	SetFetchExtendedInventory(val interface{})
	FetchExtendedInventoryInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *CloudAwsIntegrationsSns
	SetInternalValue(val *CloudAwsIntegrationsSns)
	MetricsPollingInterval() *float64
	SetMetricsPollingInterval(val *float64)
	MetricsPollingIntervalInput() *float64
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
	ResetAwsRegions()
	ResetFetchExtendedInventory()
	ResetMetricsPollingInterval()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudAwsIntegrationsSnsOutputReference
type jsiiProxy_CloudAwsIntegrationsSnsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) AwsRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"awsRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) AwsRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"awsRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) FetchExtendedInventory() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchExtendedInventory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) FetchExtendedInventoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchExtendedInventoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) InternalValue() *CloudAwsIntegrationsSns {
	var returns *CloudAwsIntegrationsSns
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) MetricsPollingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsPollingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) MetricsPollingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsPollingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudAwsIntegrationsSnsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudAwsIntegrationsSnsOutputReference {
	_init_.Initialize()

	if err := validateNewCloudAwsIntegrationsSnsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudAwsIntegrationsSnsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.cloudAwsIntegrations.CloudAwsIntegrationsSnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudAwsIntegrationsSnsOutputReference_Override(c CloudAwsIntegrationsSnsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.cloudAwsIntegrations.CloudAwsIntegrationsSnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference)SetAwsRegions(val *[]*string) {
	if err := j.validateSetAwsRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsRegions",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference)SetFetchExtendedInventory(val interface{}) {
	if err := j.validateSetFetchExtendedInventoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fetchExtendedInventory",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference)SetInternalValue(val *CloudAwsIntegrationsSns) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference)SetMetricsPollingInterval(val *float64) {
	if err := j.validateSetMetricsPollingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsPollingInterval",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsSnsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) ResetAwsRegions() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsRegions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) ResetFetchExtendedInventory() {
	_jsii_.InvokeVoid(
		c,
		"resetFetchExtendedInventory",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) ResetMetricsPollingInterval() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricsPollingInterval",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrationsSnsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

