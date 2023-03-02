package alertmutingrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v6/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v6/alertmutingrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlertMutingRuleScheduleOutputReference interface {
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
	EndRepeat() *string
	SetEndRepeat(val *string)
	EndRepeatInput() *string
	EndTime() *string
	SetEndTime(val *string)
	EndTimeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *AlertMutingRuleSchedule
	SetInternalValue(val *AlertMutingRuleSchedule)
	Repeat() *string
	SetRepeat(val *string)
	RepeatCount() *float64
	SetRepeatCount(val *float64)
	RepeatCountInput() *float64
	RepeatInput() *string
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
	WeeklyRepeatDays() *[]*string
	SetWeeklyRepeatDays(val *[]*string)
	WeeklyRepeatDaysInput() *[]*string
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
	ResetEndRepeat()
	ResetEndTime()
	ResetRepeat()
	ResetRepeatCount()
	ResetStartTime()
	ResetWeeklyRepeatDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlertMutingRuleScheduleOutputReference
type jsiiProxy_AlertMutingRuleScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) EndRepeat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endRepeat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) EndRepeatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endRepeatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) EndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) InternalValue() *AlertMutingRuleSchedule {
	var returns *AlertMutingRuleSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) Repeat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repeat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) RepeatCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"repeatCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) RepeatCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"repeatCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) RepeatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repeatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) WeeklyRepeatDays() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeklyRepeatDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference) WeeklyRepeatDaysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"weeklyRepeatDaysInput",
		&returns,
	)
	return returns
}


func NewAlertMutingRuleScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlertMutingRuleScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewAlertMutingRuleScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertMutingRuleScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.alertMutingRule.AlertMutingRuleScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlertMutingRuleScheduleOutputReference_Override(a AlertMutingRuleScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.alertMutingRule.AlertMutingRuleScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference)SetEndRepeat(val *string) {
	if err := j.validateSetEndRepeatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endRepeat",
		val,
	)
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference)SetEndTime(val *string) {
	if err := j.validateSetEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference)SetInternalValue(val *AlertMutingRuleSchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference)SetRepeat(val *string) {
	if err := j.validateSetRepeatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repeat",
		val,
	)
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference)SetRepeatCount(val *float64) {
	if err := j.validateSetRepeatCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repeatCount",
		val,
	)
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (j *jsiiProxy_AlertMutingRuleScheduleOutputReference)SetWeeklyRepeatDays(val *[]*string) {
	if err := j.validateSetWeeklyRepeatDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeklyRepeatDays",
		val,
	)
}

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) ResetEndRepeat() {
	_jsii_.InvokeVoid(
		a,
		"resetEndRepeat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) ResetEndTime() {
	_jsii_.InvokeVoid(
		a,
		"resetEndTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) ResetRepeat() {
	_jsii_.InvokeVoid(
		a,
		"resetRepeat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) ResetRepeatCount() {
	_jsii_.InvokeVoid(
		a,
		"resetRepeatCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		a,
		"resetStartTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) ResetWeeklyRepeatDays() {
	_jsii_.InvokeVoid(
		a,
		"resetWeeklyRepeatDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AlertMutingRuleScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

