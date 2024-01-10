// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alertchannel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/jsii"

	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/alertchannel/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlertChannelConfigAOutputReference interface {
	cdktf.ComplexObject
	ApiKey() *string
	SetApiKey(val *string)
	ApiKeyInput() *string
	AuthPassword() *string
	SetAuthPassword(val *string)
	AuthPasswordInput() *string
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
	AuthUsername() *string
	SetAuthUsername(val *string)
	AuthUsernameInput() *string
	BaseUrl() *string
	SetBaseUrl(val *string)
	BaseUrlInput() *string
	Channel() *string
	SetChannel(val *string)
	ChannelInput() *string
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
	Headers() *map[string]*string
	SetHeaders(val *map[string]*string)
	HeadersInput() *map[string]*string
	HeadersString() *string
	SetHeadersString(val *string)
	HeadersStringInput() *string
	IncludeJsonAttachment() *string
	SetIncludeJsonAttachment(val *string)
	IncludeJsonAttachmentInput() *string
	InternalValue() *AlertChannelConfigA
	SetInternalValue(val *AlertChannelConfigA)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	Payload() *map[string]*string
	SetPayload(val *map[string]*string)
	PayloadInput() *map[string]*string
	PayloadString() *string
	SetPayloadString(val *string)
	PayloadStringInput() *string
	PayloadType() *string
	SetPayloadType(val *string)
	PayloadTypeInput() *string
	Recipients() *string
	SetRecipients(val *string)
	RecipientsInput() *string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RouteKey() *string
	SetRouteKey(val *string)
	RouteKeyInput() *string
	ServiceKey() *string
	SetServiceKey(val *string)
	ServiceKeyInput() *string
	Tags() *string
	SetTags(val *string)
	TagsInput() *string
	Teams() *string
	SetTeams(val *string)
	TeamsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	UserId() *string
	SetUserId(val *string)
	UserIdInput() *string
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
	ResetApiKey()
	ResetAuthPassword()
	ResetAuthType()
	ResetAuthUsername()
	ResetBaseUrl()
	ResetChannel()
	ResetHeaders()
	ResetHeadersString()
	ResetIncludeJsonAttachment()
	ResetKey()
	ResetPayload()
	ResetPayloadString()
	ResetPayloadType()
	ResetRecipients()
	ResetRegion()
	ResetRouteKey()
	ResetServiceKey()
	ResetTags()
	ResetTeams()
	ResetUrl()
	ResetUserId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlertChannelConfigAOutputReference
type jsiiProxy_AlertChannelConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) ApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) AuthPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) AuthPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) AuthUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) AuthUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) BaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) BaseUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) Channel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) ChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) Headers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) HeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) HeadersString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headersString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) HeadersStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headersStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) IncludeJsonAttachment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeJsonAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) IncludeJsonAttachmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeJsonAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) InternalValue() *AlertChannelConfigA {
	var returns *AlertChannelConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) Payload() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"payload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) PayloadInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"payloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) PayloadString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) PayloadStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) PayloadType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) PayloadTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) Recipients() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) RecipientsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) RouteKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) RouteKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) ServiceKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) ServiceKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) Tags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) TagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) Teams() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) TeamsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference) UserIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userIdInput",
		&returns,
	)
	return returns
}


func NewAlertChannelConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlertChannelConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewAlertChannelConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertChannelConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.alertChannel.AlertChannelConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlertChannelConfigAOutputReference_Override(a AlertChannelConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.alertChannel.AlertChannelConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetApiKey(val *string) {
	if err := j.validateSetApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiKey",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetAuthPassword(val *string) {
	if err := j.validateSetAuthPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authPassword",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetAuthUsername(val *string) {
	if err := j.validateSetAuthUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authUsername",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetBaseUrl(val *string) {
	if err := j.validateSetBaseUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseUrl",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetChannel(val *string) {
	if err := j.validateSetChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channel",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetHeaders(val *map[string]*string) {
	if err := j.validateSetHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetHeadersString(val *string) {
	if err := j.validateSetHeadersStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headersString",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetIncludeJsonAttachment(val *string) {
	if err := j.validateSetIncludeJsonAttachmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeJsonAttachment",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetInternalValue(val *AlertChannelConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetPayload(val *map[string]*string) {
	if err := j.validateSetPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payload",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetPayloadString(val *string) {
	if err := j.validateSetPayloadStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payloadString",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetPayloadType(val *string) {
	if err := j.validateSetPayloadTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payloadType",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetRecipients(val *string) {
	if err := j.validateSetRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recipients",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetRouteKey(val *string) {
	if err := j.validateSetRouteKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeKey",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetServiceKey(val *string) {
	if err := j.validateSetServiceKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceKey",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetTags(val *string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetTeams(val *string) {
	if err := j.validateSetTeamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teams",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_AlertChannelConfigAOutputReference)SetUserId(val *string) {
	if err := j.validateSetUserIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userId",
		val,
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AlertChannelConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertChannelConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AlertChannelConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AlertChannelConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AlertChannelConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AlertChannelConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AlertChannelConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AlertChannelConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AlertChannelConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetApiKey() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetAuthPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetAuthType() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetAuthUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetBaseUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetBaseUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetChannel() {
	_jsii_.InvokeVoid(
		a,
		"resetChannel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetHeadersString() {
	_jsii_.InvokeVoid(
		a,
		"resetHeadersString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetIncludeJsonAttachment() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeJsonAttachment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		a,
		"resetKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetPayload() {
	_jsii_.InvokeVoid(
		a,
		"resetPayload",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetPayloadString() {
	_jsii_.InvokeVoid(
		a,
		"resetPayloadString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetPayloadType() {
	_jsii_.InvokeVoid(
		a,
		"resetPayloadType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetRecipients() {
	_jsii_.InvokeVoid(
		a,
		"resetRecipients",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetRouteKey() {
	_jsii_.InvokeVoid(
		a,
		"resetRouteKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetServiceKey() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetTeams() {
	_jsii_.InvokeVoid(
		a,
		"resetTeams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ResetUserId() {
	_jsii_.InvokeVoid(
		a,
		"resetUserId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlertChannelConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AlertChannelConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

