// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudawsintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v10/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v10/cloudawsintegrations/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudAwsIntegrationsCloudfrontOutputReference interface {
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
	FetchLambdasAtEdge() interface{}
	SetFetchLambdasAtEdge(val interface{})
	FetchLambdasAtEdgeInput() interface{}
	FetchTags() interface{}
	SetFetchTags(val interface{})
	FetchTagsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *CloudAwsIntegrationsCloudfront
	SetInternalValue(val *CloudAwsIntegrationsCloudfront)
	MetricsPollingInterval() *float64
	SetMetricsPollingInterval(val *float64)
	MetricsPollingIntervalInput() *float64
	TagKey() *string
	SetTagKey(val *string)
	TagKeyInput() *string
	TagValue() *string
	SetTagValue(val *string)
	TagValueInput() *string
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
	ResetFetchLambdasAtEdge()
	ResetFetchTags()
	ResetMetricsPollingInterval()
	ResetTagKey()
	ResetTagValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudAwsIntegrationsCloudfrontOutputReference
type jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) FetchLambdasAtEdge() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchLambdasAtEdge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) FetchLambdasAtEdgeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchLambdasAtEdgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) FetchTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) FetchTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) InternalValue() *CloudAwsIntegrationsCloudfront {
	var returns *CloudAwsIntegrationsCloudfront
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) MetricsPollingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsPollingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) MetricsPollingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsPollingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) TagKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) TagKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) TagValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) TagValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudAwsIntegrationsCloudfrontOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudAwsIntegrationsCloudfrontOutputReference {
	_init_.Initialize()

	if err := validateNewCloudAwsIntegrationsCloudfrontOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.cloudAwsIntegrations.CloudAwsIntegrationsCloudfrontOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudAwsIntegrationsCloudfrontOutputReference_Override(c CloudAwsIntegrationsCloudfrontOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.cloudAwsIntegrations.CloudAwsIntegrationsCloudfrontOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference)SetFetchLambdasAtEdge(val interface{}) {
	if err := j.validateSetFetchLambdasAtEdgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fetchLambdasAtEdge",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference)SetFetchTags(val interface{}) {
	if err := j.validateSetFetchTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fetchTags",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference)SetInternalValue(val *CloudAwsIntegrationsCloudfront) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference)SetMetricsPollingInterval(val *float64) {
	if err := j.validateSetMetricsPollingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsPollingInterval",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference)SetTagKey(val *string) {
	if err := j.validateSetTagKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagKey",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference)SetTagValue(val *string) {
	if err := j.validateSetTagValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagValue",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) ResetFetchLambdasAtEdge() {
	_jsii_.InvokeVoid(
		c,
		"resetFetchLambdasAtEdge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) ResetFetchTags() {
	_jsii_.InvokeVoid(
		c,
		"resetFetchTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) ResetMetricsPollingInterval() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricsPollingInterval",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) ResetTagKey() {
	_jsii_.InvokeVoid(
		c,
		"resetTagKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) ResetTagValue() {
	_jsii_.InvokeVoid(
		c,
		"resetTagValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudAwsIntegrationsCloudfrontOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

