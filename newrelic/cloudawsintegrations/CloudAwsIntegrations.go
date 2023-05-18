package cloudawsintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v7/cloudawsintegrations/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations newrelic_cloud_aws_integrations}.
type CloudAwsIntegrations interface {
	cdktf.TerraformResource
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
	Alb() CloudAwsIntegrationsAlbOutputReference
	AlbInput() *CloudAwsIntegrationsAlb
	ApiGateway() CloudAwsIntegrationsApiGatewayOutputReference
	ApiGatewayInput() *CloudAwsIntegrationsApiGateway
	AutoScaling() CloudAwsIntegrationsAutoScalingOutputReference
	AutoScalingInput() *CloudAwsIntegrationsAutoScaling
	AwsAppSync() CloudAwsIntegrationsAwsAppSyncOutputReference
	AwsAppSyncInput() *CloudAwsIntegrationsAwsAppSync
	AwsAthena() CloudAwsIntegrationsAwsAthenaOutputReference
	AwsAthenaInput() *CloudAwsIntegrationsAwsAthena
	AwsCognito() CloudAwsIntegrationsAwsCognitoOutputReference
	AwsCognitoInput() *CloudAwsIntegrationsAwsCognito
	AwsConnect() CloudAwsIntegrationsAwsConnectOutputReference
	AwsConnectInput() *CloudAwsIntegrationsAwsConnect
	AwsDirectConnect() CloudAwsIntegrationsAwsDirectConnectOutputReference
	AwsDirectConnectInput() *CloudAwsIntegrationsAwsDirectConnect
	AwsFsx() CloudAwsIntegrationsAwsFsxOutputReference
	AwsFsxInput() *CloudAwsIntegrationsAwsFsx
	Billing() CloudAwsIntegrationsBillingOutputReference
	BillingInput() *CloudAwsIntegrationsBilling
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Cloudtrail() CloudAwsIntegrationsCloudtrailOutputReference
	CloudtrailInput() *CloudAwsIntegrationsCloudtrail
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DocDb() CloudAwsIntegrationsDocDbOutputReference
	DocDbInput() *CloudAwsIntegrationsDocDb
	Ebs() CloudAwsIntegrationsEbsOutputReference
	EbsInput() *CloudAwsIntegrationsEbs
	Elasticache() CloudAwsIntegrationsElasticacheOutputReference
	ElasticacheInput() *CloudAwsIntegrationsElasticache
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Health() CloudAwsIntegrationsHealthOutputReference
	HealthInput() *CloudAwsIntegrationsHealth
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinkedAccountId() *float64
	SetLinkedAccountId(val *float64)
	LinkedAccountIdInput() *float64
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	S3() CloudAwsIntegrationsS3OutputReference
	S3Input() *CloudAwsIntegrationsS3
	Sqs() CloudAwsIntegrationsSqsOutputReference
	SqsInput() *CloudAwsIntegrationsSqs
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrustedAdvisor() CloudAwsIntegrationsTrustedAdvisorOutputReference
	TrustedAdvisorInput() *CloudAwsIntegrationsTrustedAdvisor
	Vpc() CloudAwsIntegrationsVpcOutputReference
	VpcInput() *CloudAwsIntegrationsVpc
	XRay() CloudAwsIntegrationsXRayOutputReference
	XRayInput() *CloudAwsIntegrationsXRay
	// Experimental.
	AddOverride(path *string, value interface{})
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAlb(value *CloudAwsIntegrationsAlb)
	PutApiGateway(value *CloudAwsIntegrationsApiGateway)
	PutAutoScaling(value *CloudAwsIntegrationsAutoScaling)
	PutAwsAppSync(value *CloudAwsIntegrationsAwsAppSync)
	PutAwsAthena(value *CloudAwsIntegrationsAwsAthena)
	PutAwsCognito(value *CloudAwsIntegrationsAwsCognito)
	PutAwsConnect(value *CloudAwsIntegrationsAwsConnect)
	PutAwsDirectConnect(value *CloudAwsIntegrationsAwsDirectConnect)
	PutAwsFsx(value *CloudAwsIntegrationsAwsFsx)
	PutBilling(value *CloudAwsIntegrationsBilling)
	PutCloudtrail(value *CloudAwsIntegrationsCloudtrail)
	PutDocDb(value *CloudAwsIntegrationsDocDb)
	PutEbs(value *CloudAwsIntegrationsEbs)
	PutElasticache(value *CloudAwsIntegrationsElasticache)
	PutHealth(value *CloudAwsIntegrationsHealth)
	PutS3(value *CloudAwsIntegrationsS3)
	PutSqs(value *CloudAwsIntegrationsSqs)
	PutTrustedAdvisor(value *CloudAwsIntegrationsTrustedAdvisor)
	PutVpc(value *CloudAwsIntegrationsVpc)
	PutXRay(value *CloudAwsIntegrationsXRay)
	ResetAccountId()
	ResetAlb()
	ResetApiGateway()
	ResetAutoScaling()
	ResetAwsAppSync()
	ResetAwsAthena()
	ResetAwsCognito()
	ResetAwsConnect()
	ResetAwsDirectConnect()
	ResetAwsFsx()
	ResetBilling()
	ResetCloudtrail()
	ResetDocDb()
	ResetEbs()
	ResetElasticache()
	ResetHealth()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetS3()
	ResetSqs()
	ResetTrustedAdvisor()
	ResetVpc()
	ResetXRay()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudAwsIntegrations
type jsiiProxy_CloudAwsIntegrations struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudAwsIntegrations) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Alb() CloudAwsIntegrationsAlbOutputReference {
	var returns CloudAwsIntegrationsAlbOutputReference
	_jsii_.Get(
		j,
		"alb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AlbInput() *CloudAwsIntegrationsAlb {
	var returns *CloudAwsIntegrationsAlb
	_jsii_.Get(
		j,
		"albInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) ApiGateway() CloudAwsIntegrationsApiGatewayOutputReference {
	var returns CloudAwsIntegrationsApiGatewayOutputReference
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) ApiGatewayInput() *CloudAwsIntegrationsApiGateway {
	var returns *CloudAwsIntegrationsApiGateway
	_jsii_.Get(
		j,
		"apiGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AutoScaling() CloudAwsIntegrationsAutoScalingOutputReference {
	var returns CloudAwsIntegrationsAutoScalingOutputReference
	_jsii_.Get(
		j,
		"autoScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AutoScalingInput() *CloudAwsIntegrationsAutoScaling {
	var returns *CloudAwsIntegrationsAutoScaling
	_jsii_.Get(
		j,
		"autoScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsAppSync() CloudAwsIntegrationsAwsAppSyncOutputReference {
	var returns CloudAwsIntegrationsAwsAppSyncOutputReference
	_jsii_.Get(
		j,
		"awsAppSync",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsAppSyncInput() *CloudAwsIntegrationsAwsAppSync {
	var returns *CloudAwsIntegrationsAwsAppSync
	_jsii_.Get(
		j,
		"awsAppSyncInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsAthena() CloudAwsIntegrationsAwsAthenaOutputReference {
	var returns CloudAwsIntegrationsAwsAthenaOutputReference
	_jsii_.Get(
		j,
		"awsAthena",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsAthenaInput() *CloudAwsIntegrationsAwsAthena {
	var returns *CloudAwsIntegrationsAwsAthena
	_jsii_.Get(
		j,
		"awsAthenaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsCognito() CloudAwsIntegrationsAwsCognitoOutputReference {
	var returns CloudAwsIntegrationsAwsCognitoOutputReference
	_jsii_.Get(
		j,
		"awsCognito",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsCognitoInput() *CloudAwsIntegrationsAwsCognito {
	var returns *CloudAwsIntegrationsAwsCognito
	_jsii_.Get(
		j,
		"awsCognitoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsConnect() CloudAwsIntegrationsAwsConnectOutputReference {
	var returns CloudAwsIntegrationsAwsConnectOutputReference
	_jsii_.Get(
		j,
		"awsConnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsConnectInput() *CloudAwsIntegrationsAwsConnect {
	var returns *CloudAwsIntegrationsAwsConnect
	_jsii_.Get(
		j,
		"awsConnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsDirectConnect() CloudAwsIntegrationsAwsDirectConnectOutputReference {
	var returns CloudAwsIntegrationsAwsDirectConnectOutputReference
	_jsii_.Get(
		j,
		"awsDirectConnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsDirectConnectInput() *CloudAwsIntegrationsAwsDirectConnect {
	var returns *CloudAwsIntegrationsAwsDirectConnect
	_jsii_.Get(
		j,
		"awsDirectConnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsFsx() CloudAwsIntegrationsAwsFsxOutputReference {
	var returns CloudAwsIntegrationsAwsFsxOutputReference
	_jsii_.Get(
		j,
		"awsFsx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsFsxInput() *CloudAwsIntegrationsAwsFsx {
	var returns *CloudAwsIntegrationsAwsFsx
	_jsii_.Get(
		j,
		"awsFsxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Billing() CloudAwsIntegrationsBillingOutputReference {
	var returns CloudAwsIntegrationsBillingOutputReference
	_jsii_.Get(
		j,
		"billing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) BillingInput() *CloudAwsIntegrationsBilling {
	var returns *CloudAwsIntegrationsBilling
	_jsii_.Get(
		j,
		"billingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Cloudtrail() CloudAwsIntegrationsCloudtrailOutputReference {
	var returns CloudAwsIntegrationsCloudtrailOutputReference
	_jsii_.Get(
		j,
		"cloudtrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) CloudtrailInput() *CloudAwsIntegrationsCloudtrail {
	var returns *CloudAwsIntegrationsCloudtrail
	_jsii_.Get(
		j,
		"cloudtrailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) DocDb() CloudAwsIntegrationsDocDbOutputReference {
	var returns CloudAwsIntegrationsDocDbOutputReference
	_jsii_.Get(
		j,
		"docDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) DocDbInput() *CloudAwsIntegrationsDocDb {
	var returns *CloudAwsIntegrationsDocDb
	_jsii_.Get(
		j,
		"docDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Ebs() CloudAwsIntegrationsEbsOutputReference {
	var returns CloudAwsIntegrationsEbsOutputReference
	_jsii_.Get(
		j,
		"ebs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) EbsInput() *CloudAwsIntegrationsEbs {
	var returns *CloudAwsIntegrationsEbs
	_jsii_.Get(
		j,
		"ebsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Elasticache() CloudAwsIntegrationsElasticacheOutputReference {
	var returns CloudAwsIntegrationsElasticacheOutputReference
	_jsii_.Get(
		j,
		"elasticache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) ElasticacheInput() *CloudAwsIntegrationsElasticache {
	var returns *CloudAwsIntegrationsElasticache
	_jsii_.Get(
		j,
		"elasticacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Health() CloudAwsIntegrationsHealthOutputReference {
	var returns CloudAwsIntegrationsHealthOutputReference
	_jsii_.Get(
		j,
		"health",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) HealthInput() *CloudAwsIntegrationsHealth {
	var returns *CloudAwsIntegrationsHealth
	_jsii_.Get(
		j,
		"healthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) LinkedAccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"linkedAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) LinkedAccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"linkedAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) S3() CloudAwsIntegrationsS3OutputReference {
	var returns CloudAwsIntegrationsS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) S3Input() *CloudAwsIntegrationsS3 {
	var returns *CloudAwsIntegrationsS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Sqs() CloudAwsIntegrationsSqsOutputReference {
	var returns CloudAwsIntegrationsSqsOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) SqsInput() *CloudAwsIntegrationsSqs {
	var returns *CloudAwsIntegrationsSqs
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) TrustedAdvisor() CloudAwsIntegrationsTrustedAdvisorOutputReference {
	var returns CloudAwsIntegrationsTrustedAdvisorOutputReference
	_jsii_.Get(
		j,
		"trustedAdvisor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) TrustedAdvisorInput() *CloudAwsIntegrationsTrustedAdvisor {
	var returns *CloudAwsIntegrationsTrustedAdvisor
	_jsii_.Get(
		j,
		"trustedAdvisorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Vpc() CloudAwsIntegrationsVpcOutputReference {
	var returns CloudAwsIntegrationsVpcOutputReference
	_jsii_.Get(
		j,
		"vpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) VpcInput() *CloudAwsIntegrationsVpc {
	var returns *CloudAwsIntegrationsVpc
	_jsii_.Get(
		j,
		"vpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) XRay() CloudAwsIntegrationsXRayOutputReference {
	var returns CloudAwsIntegrationsXRayOutputReference
	_jsii_.Get(
		j,
		"xRay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) XRayInput() *CloudAwsIntegrationsXRay {
	var returns *CloudAwsIntegrationsXRay
	_jsii_.Get(
		j,
		"xRayInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations newrelic_cloud_aws_integrations} Resource.
func NewCloudAwsIntegrations(scope constructs.Construct, id *string, config *CloudAwsIntegrationsConfig) CloudAwsIntegrations {
	_init_.Initialize()

	if err := validateNewCloudAwsIntegrationsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudAwsIntegrations{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.cloudAwsIntegrations.CloudAwsIntegrations",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.23.0/docs/resources/cloud_aws_integrations newrelic_cloud_aws_integrations} Resource.
func NewCloudAwsIntegrations_Override(c CloudAwsIntegrations, scope constructs.Construct, id *string, config *CloudAwsIntegrationsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.cloudAwsIntegrations.CloudAwsIntegrations",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudAwsIntegrations)SetAccountId(val *float64) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrations)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrations)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrations)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrations)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrations)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrations)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrations)SetLinkedAccountId(val *float64) {
	if err := j.validateSetLinkedAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkedAccountId",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrations)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudAwsIntegrations)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CloudAwsIntegrations_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAwsIntegrations_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudAwsIntegrations.CloudAwsIntegrations",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudAwsIntegrations_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAwsIntegrations_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudAwsIntegrations.CloudAwsIntegrations",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudAwsIntegrations_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAwsIntegrations_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudAwsIntegrations.CloudAwsIntegrations",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudAwsIntegrations_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.cloudAwsIntegrations.CloudAwsIntegrations",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudAwsIntegrations) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudAwsIntegrations) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudAwsIntegrations) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudAwsIntegrations) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudAwsIntegrations) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudAwsIntegrations) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudAwsIntegrations) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudAwsIntegrations) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudAwsIntegrations) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudAwsIntegrations) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrations) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAlb(value *CloudAwsIntegrationsAlb) {
	if err := c.validatePutAlbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAlb",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutApiGateway(value *CloudAwsIntegrationsApiGateway) {
	if err := c.validatePutApiGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putApiGateway",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAutoScaling(value *CloudAwsIntegrationsAutoScaling) {
	if err := c.validatePutAutoScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutoScaling",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsAppSync(value *CloudAwsIntegrationsAwsAppSync) {
	if err := c.validatePutAwsAppSyncParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsAppSync",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsAthena(value *CloudAwsIntegrationsAwsAthena) {
	if err := c.validatePutAwsAthenaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsAthena",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsCognito(value *CloudAwsIntegrationsAwsCognito) {
	if err := c.validatePutAwsCognitoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsCognito",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsConnect(value *CloudAwsIntegrationsAwsConnect) {
	if err := c.validatePutAwsConnectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsConnect",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsDirectConnect(value *CloudAwsIntegrationsAwsDirectConnect) {
	if err := c.validatePutAwsDirectConnectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsDirectConnect",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsFsx(value *CloudAwsIntegrationsAwsFsx) {
	if err := c.validatePutAwsFsxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsFsx",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutBilling(value *CloudAwsIntegrationsBilling) {
	if err := c.validatePutBillingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBilling",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutCloudtrail(value *CloudAwsIntegrationsCloudtrail) {
	if err := c.validatePutCloudtrailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCloudtrail",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutDocDb(value *CloudAwsIntegrationsDocDb) {
	if err := c.validatePutDocDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDocDb",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutEbs(value *CloudAwsIntegrationsEbs) {
	if err := c.validatePutEbsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEbs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutElasticache(value *CloudAwsIntegrationsElasticache) {
	if err := c.validatePutElasticacheParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putElasticache",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutHealth(value *CloudAwsIntegrationsHealth) {
	if err := c.validatePutHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHealth",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutS3(value *CloudAwsIntegrationsS3) {
	if err := c.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putS3",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutSqs(value *CloudAwsIntegrationsSqs) {
	if err := c.validatePutSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSqs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutTrustedAdvisor(value *CloudAwsIntegrationsTrustedAdvisor) {
	if err := c.validatePutTrustedAdvisorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTrustedAdvisor",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutVpc(value *CloudAwsIntegrationsVpc) {
	if err := c.validatePutVpcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVpc",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutXRay(value *CloudAwsIntegrationsXRay) {
	if err := c.validatePutXRayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putXRay",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAccountId() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAlb() {
	_jsii_.InvokeVoid(
		c,
		"resetAlb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetApiGateway() {
	_jsii_.InvokeVoid(
		c,
		"resetApiGateway",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAutoScaling() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoScaling",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsAppSync() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsAppSync",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsAthena() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsAthena",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsCognito() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsCognito",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsConnect() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsConnect",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsDirectConnect() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsDirectConnect",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsFsx() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsFsx",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetBilling() {
	_jsii_.InvokeVoid(
		c,
		"resetBilling",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetCloudtrail() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudtrail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetDocDb() {
	_jsii_.InvokeVoid(
		c,
		"resetDocDb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetEbs() {
	_jsii_.InvokeVoid(
		c,
		"resetEbs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetElasticache() {
	_jsii_.InvokeVoid(
		c,
		"resetElasticache",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetHealth() {
	_jsii_.InvokeVoid(
		c,
		"resetHealth",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetS3() {
	_jsii_.InvokeVoid(
		c,
		"resetS3",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetSqs() {
	_jsii_.InvokeVoid(
		c,
		"resetSqs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetTrustedAdvisor() {
	_jsii_.InvokeVoid(
		c,
		"resetTrustedAdvisor",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetVpc() {
	_jsii_.InvokeVoid(
		c,
		"resetVpc",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetXRay() {
	_jsii_.InvokeVoid(
		c,
		"resetXRay",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrations) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrations) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrations) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

