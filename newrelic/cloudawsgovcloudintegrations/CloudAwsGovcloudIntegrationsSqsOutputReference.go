package cloudawsgovcloudintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v7/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v7/cloudawsgovcloudintegrations/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudAwsGovcloudIntegrationsSqsOutputReference interface {
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
	FetchTags() interface{}
	SetFetchTags(val interface{})
	FetchTagsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *CloudAwsGovcloudIntegrationsSqs
	SetInternalValue(val *CloudAwsGovcloudIntegrationsSqs)
	MetricsPollingInterval() *float64
	SetMetricsPollingInterval(val *float64)
	MetricsPollingIntervalInput() *float64
	QueuePrefixes() *[]*string
	SetQueuePrefixes(val *[]*string)
	QueuePrefixesInput() *[]*string
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
	ResetAwsRegions()
	ResetFetchExtendedInventory()
	ResetFetchTags()
	ResetMetricsPollingInterval()
	ResetQueuePrefixes()
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

// The jsii proxy struct for CloudAwsGovcloudIntegrationsSqsOutputReference
type jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) AwsRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"awsRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) AwsRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"awsRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) FetchExtendedInventory() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchExtendedInventory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) FetchExtendedInventoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchExtendedInventoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) FetchTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) FetchTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fetchTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) InternalValue() *CloudAwsGovcloudIntegrationsSqs {
	var returns *CloudAwsGovcloudIntegrationsSqs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) MetricsPollingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsPollingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) MetricsPollingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsPollingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) QueuePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queuePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) QueuePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queuePrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) TagKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) TagKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) TagValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) TagValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudAwsGovcloudIntegrationsSqsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudAwsGovcloudIntegrationsSqsOutputReference {
	_init_.Initialize()

	if err := validateNewCloudAwsGovcloudIntegrationsSqsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.cloudAwsGovcloudIntegrations.CloudAwsGovcloudIntegrationsSqsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudAwsGovcloudIntegrationsSqsOutputReference_Override(c CloudAwsGovcloudIntegrationsSqsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.cloudAwsGovcloudIntegrations.CloudAwsGovcloudIntegrationsSqsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference)SetAwsRegions(val *[]*string) {
	if err := j.validateSetAwsRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsRegions",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference)SetFetchExtendedInventory(val interface{}) {
	if err := j.validateSetFetchExtendedInventoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fetchExtendedInventory",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference)SetFetchTags(val interface{}) {
	if err := j.validateSetFetchTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fetchTags",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference)SetInternalValue(val *CloudAwsGovcloudIntegrationsSqs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference)SetMetricsPollingInterval(val *float64) {
	if err := j.validateSetMetricsPollingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsPollingInterval",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference)SetQueuePrefixes(val *[]*string) {
	if err := j.validateSetQueuePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queuePrefixes",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference)SetTagKey(val *string) {
	if err := j.validateSetTagKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagKey",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference)SetTagValue(val *string) {
	if err := j.validateSetTagValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagValue",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) ResetAwsRegions() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsRegions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) ResetFetchExtendedInventory() {
	_jsii_.InvokeVoid(
		c,
		"resetFetchExtendedInventory",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) ResetFetchTags() {
	_jsii_.InvokeVoid(
		c,
		"resetFetchTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) ResetMetricsPollingInterval() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricsPollingInterval",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) ResetQueuePrefixes() {
	_jsii_.InvokeVoid(
		c,
		"resetQueuePrefixes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) ResetTagKey() {
	_jsii_.InvokeVoid(
		c,
		"resetTagKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) ResetTagValue() {
	_jsii_.InvokeVoid(
		c,
		"resetTagValue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrationsSqsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

