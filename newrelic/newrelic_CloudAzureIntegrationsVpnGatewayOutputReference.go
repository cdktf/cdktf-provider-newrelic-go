// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-newrelic-go/newrelic/v2/jsii"

	"github.com/hashicorp/cdktf-provider-newrelic-go/newrelic/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudAzureIntegrationsVpnGatewayOutputReference interface {
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
	InternalValue() *CloudAzureIntegrationsVpnGateway
	SetInternalValue(val *CloudAzureIntegrationsVpnGateway)
	MetricsPollingInterval() *float64
	SetMetricsPollingInterval(val *float64)
	MetricsPollingIntervalInput() *float64
	ResourceGroups() *[]*string
	SetResourceGroups(val *[]*string)
	ResourceGroupsInput() *[]*string
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
	ResetMetricsPollingInterval()
	ResetResourceGroups()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudAzureIntegrationsVpnGatewayOutputReference
type jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) InternalValue() *CloudAzureIntegrationsVpnGateway {
	var returns *CloudAzureIntegrationsVpnGateway
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) MetricsPollingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsPollingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) MetricsPollingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricsPollingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) ResourceGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) ResourceGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudAzureIntegrationsVpnGatewayOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudAzureIntegrationsVpnGatewayOutputReference {
	_init_.Initialize()

	j := jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.CloudAzureIntegrationsVpnGatewayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudAzureIntegrationsVpnGatewayOutputReference_Override(c CloudAzureIntegrationsVpnGatewayOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.CloudAzureIntegrationsVpnGatewayOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) SetInternalValue(val *CloudAzureIntegrationsVpnGateway) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) SetMetricsPollingInterval(val *float64) {
	_jsii_.Set(
		j,
		"metricsPollingInterval",
		val,
	)
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) SetResourceGroups(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceGroups",
		val,
	)
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) ResetMetricsPollingInterval() {
	_jsii_.InvokeVoid(
		c,
		"resetMetricsPollingInterval",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) ResetResourceGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrationsVpnGatewayOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

