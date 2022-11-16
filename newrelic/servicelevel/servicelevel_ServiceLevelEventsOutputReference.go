package servicelevel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v5/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v5/servicelevel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceLevelEventsOutputReference interface {
	cdktf.ComplexObject
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
	BadEvents() ServiceLevelEventsBadEventsOutputReference
	BadEventsInput() *ServiceLevelEventsBadEvents
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
	GoodEvents() ServiceLevelEventsGoodEventsOutputReference
	GoodEventsInput() *ServiceLevelEventsGoodEvents
	InternalValue() *ServiceLevelEvents
	SetInternalValue(val *ServiceLevelEvents)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValidEvents() ServiceLevelEventsValidEventsOutputReference
	ValidEventsInput() *ServiceLevelEventsValidEvents
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
	PutBadEvents(value *ServiceLevelEventsBadEvents)
	PutGoodEvents(value *ServiceLevelEventsGoodEvents)
	PutValidEvents(value *ServiceLevelEventsValidEvents)
	ResetBadEvents()
	ResetGoodEvents()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceLevelEventsOutputReference
type jsiiProxy_ServiceLevelEventsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference) BadEvents() ServiceLevelEventsBadEventsOutputReference {
	var returns ServiceLevelEventsBadEventsOutputReference
	_jsii_.Get(
		j,
		"badEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference) BadEventsInput() *ServiceLevelEventsBadEvents {
	var returns *ServiceLevelEventsBadEvents
	_jsii_.Get(
		j,
		"badEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference) GoodEvents() ServiceLevelEventsGoodEventsOutputReference {
	var returns ServiceLevelEventsGoodEventsOutputReference
	_jsii_.Get(
		j,
		"goodEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference) GoodEventsInput() *ServiceLevelEventsGoodEvents {
	var returns *ServiceLevelEventsGoodEvents
	_jsii_.Get(
		j,
		"goodEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference) InternalValue() *ServiceLevelEvents {
	var returns *ServiceLevelEvents
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference) ValidEvents() ServiceLevelEventsValidEventsOutputReference {
	var returns ServiceLevelEventsValidEventsOutputReference
	_jsii_.Get(
		j,
		"validEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference) ValidEventsInput() *ServiceLevelEventsValidEvents {
	var returns *ServiceLevelEventsValidEvents
	_jsii_.Get(
		j,
		"validEventsInput",
		&returns,
	)
	return returns
}


func NewServiceLevelEventsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ServiceLevelEventsOutputReference {
	_init_.Initialize()

	if err := validateNewServiceLevelEventsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceLevelEventsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.serviceLevel.ServiceLevelEventsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServiceLevelEventsOutputReference_Override(s ServiceLevelEventsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.serviceLevel.ServiceLevelEventsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference)SetAccountId(val *float64) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference)SetInternalValue(val *ServiceLevelEvents) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceLevelEventsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) PutBadEvents(value *ServiceLevelEventsBadEvents) {
	if err := s.validatePutBadEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBadEvents",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) PutGoodEvents(value *ServiceLevelEventsGoodEvents) {
	if err := s.validatePutGoodEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGoodEvents",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) PutValidEvents(value *ServiceLevelEventsValidEvents) {
	if err := s.validatePutValidEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putValidEvents",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) ResetBadEvents() {
	_jsii_.InvokeVoid(
		s,
		"resetBadEvents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) ResetGoodEvents() {
	_jsii_.InvokeVoid(
		s,
		"resetGoodEvents",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceLevelEventsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

