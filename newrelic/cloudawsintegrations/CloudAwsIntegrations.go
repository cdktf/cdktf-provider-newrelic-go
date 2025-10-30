// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudawsintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/cloudawsintegrations/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.74.0/docs/resources/cloud_aws_integrations newrelic_cloud_aws_integrations}.
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
	AwsAutoDiscovery() CloudAwsIntegrationsAwsAutoDiscoveryOutputReference
	AwsAutoDiscoveryInput() *CloudAwsIntegrationsAwsAutoDiscovery
	AwsCognito() CloudAwsIntegrationsAwsCognitoOutputReference
	AwsCognitoInput() *CloudAwsIntegrationsAwsCognito
	AwsConnect() CloudAwsIntegrationsAwsConnectOutputReference
	AwsConnectInput() *CloudAwsIntegrationsAwsConnect
	AwsDirectConnect() CloudAwsIntegrationsAwsDirectConnectOutputReference
	AwsDirectConnectInput() *CloudAwsIntegrationsAwsDirectConnect
	AwsFsx() CloudAwsIntegrationsAwsFsxOutputReference
	AwsFsxInput() *CloudAwsIntegrationsAwsFsx
	AwsGlue() CloudAwsIntegrationsAwsGlueOutputReference
	AwsGlueInput() *CloudAwsIntegrationsAwsGlue
	AwsKinesisAnalytics() CloudAwsIntegrationsAwsKinesisAnalyticsOutputReference
	AwsKinesisAnalyticsInput() *CloudAwsIntegrationsAwsKinesisAnalytics
	AwsMediaConvert() CloudAwsIntegrationsAwsMediaConvertOutputReference
	AwsMediaConvertInput() *CloudAwsIntegrationsAwsMediaConvert
	AwsMediaPackageVod() CloudAwsIntegrationsAwsMediaPackageVodOutputReference
	AwsMediaPackageVodInput() *CloudAwsIntegrationsAwsMediaPackageVod
	AwsMq() CloudAwsIntegrationsAwsMqOutputReference
	AwsMqInput() *CloudAwsIntegrationsAwsMq
	AwsMsk() CloudAwsIntegrationsAwsMskOutputReference
	AwsMskInput() *CloudAwsIntegrationsAwsMsk
	AwsNeptune() CloudAwsIntegrationsAwsNeptuneOutputReference
	AwsNeptuneInput() *CloudAwsIntegrationsAwsNeptune
	AwsQldb() CloudAwsIntegrationsAwsQldbOutputReference
	AwsQldbInput() *CloudAwsIntegrationsAwsQldb
	AwsRoute53Resolver() CloudAwsIntegrationsAwsRoute53ResolverOutputReference
	AwsRoute53ResolverInput() *CloudAwsIntegrationsAwsRoute53Resolver
	AwsStates() CloudAwsIntegrationsAwsStatesOutputReference
	AwsStatesInput() *CloudAwsIntegrationsAwsStates
	AwsTransitGateway() CloudAwsIntegrationsAwsTransitGatewayOutputReference
	AwsTransitGatewayInput() *CloudAwsIntegrationsAwsTransitGateway
	AwsWaf() CloudAwsIntegrationsAwsWafOutputReference
	AwsWafInput() *CloudAwsIntegrationsAwsWaf
	AwsWafv2() CloudAwsIntegrationsAwsWafv2OutputReference
	AwsWafv2Input() *CloudAwsIntegrationsAwsWafv2
	Billing() CloudAwsIntegrationsBillingOutputReference
	BillingInput() *CloudAwsIntegrationsBilling
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Cloudfront() CloudAwsIntegrationsCloudfrontOutputReference
	CloudfrontInput() *CloudAwsIntegrationsCloudfront
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
	Dynamodb() CloudAwsIntegrationsDynamodbOutputReference
	DynamodbInput() *CloudAwsIntegrationsDynamodb
	Ebs() CloudAwsIntegrationsEbsOutputReference
	EbsInput() *CloudAwsIntegrationsEbs
	Ec2() CloudAwsIntegrationsEc2OutputReference
	Ec2Input() *CloudAwsIntegrationsEc2
	Ecs() CloudAwsIntegrationsEcsOutputReference
	EcsInput() *CloudAwsIntegrationsEcs
	Efs() CloudAwsIntegrationsEfsOutputReference
	EfsInput() *CloudAwsIntegrationsEfs
	Elasticache() CloudAwsIntegrationsElasticacheOutputReference
	ElasticacheInput() *CloudAwsIntegrationsElasticache
	Elasticbeanstalk() CloudAwsIntegrationsElasticbeanstalkOutputReference
	ElasticbeanstalkInput() *CloudAwsIntegrationsElasticbeanstalk
	Elasticsearch() CloudAwsIntegrationsElasticsearchOutputReference
	ElasticsearchInput() *CloudAwsIntegrationsElasticsearch
	Elb() CloudAwsIntegrationsElbOutputReference
	ElbInput() *CloudAwsIntegrationsElb
	Emr() CloudAwsIntegrationsEmrOutputReference
	EmrInput() *CloudAwsIntegrationsEmr
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
	Iam() CloudAwsIntegrationsIamOutputReference
	IamInput() *CloudAwsIntegrationsIam
	Id() *string
	SetId(val *string)
	IdInput() *string
	Iot() CloudAwsIntegrationsIotOutputReference
	IotInput() *CloudAwsIntegrationsIot
	Kinesis() CloudAwsIntegrationsKinesisOutputReference
	KinesisFirehose() CloudAwsIntegrationsKinesisFirehoseOutputReference
	KinesisFirehoseInput() *CloudAwsIntegrationsKinesisFirehose
	KinesisInput() *CloudAwsIntegrationsKinesis
	Lambda() CloudAwsIntegrationsLambdaOutputReference
	LambdaInput() *CloudAwsIntegrationsLambda
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
	Rds() CloudAwsIntegrationsRdsOutputReference
	RdsInput() *CloudAwsIntegrationsRds
	Redshift() CloudAwsIntegrationsRedshiftOutputReference
	RedshiftInput() *CloudAwsIntegrationsRedshift
	Route53() CloudAwsIntegrationsRoute53OutputReference
	Route53Input() *CloudAwsIntegrationsRoute53
	S3() CloudAwsIntegrationsS3OutputReference
	S3Input() *CloudAwsIntegrationsS3
	Ses() CloudAwsIntegrationsSesOutputReference
	SesInput() *CloudAwsIntegrationsSes
	Sns() CloudAwsIntegrationsSnsOutputReference
	SnsInput() *CloudAwsIntegrationsSns
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
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAlb(value *CloudAwsIntegrationsAlb)
	PutApiGateway(value *CloudAwsIntegrationsApiGateway)
	PutAutoScaling(value *CloudAwsIntegrationsAutoScaling)
	PutAwsAppSync(value *CloudAwsIntegrationsAwsAppSync)
	PutAwsAthena(value *CloudAwsIntegrationsAwsAthena)
	PutAwsAutoDiscovery(value *CloudAwsIntegrationsAwsAutoDiscovery)
	PutAwsCognito(value *CloudAwsIntegrationsAwsCognito)
	PutAwsConnect(value *CloudAwsIntegrationsAwsConnect)
	PutAwsDirectConnect(value *CloudAwsIntegrationsAwsDirectConnect)
	PutAwsFsx(value *CloudAwsIntegrationsAwsFsx)
	PutAwsGlue(value *CloudAwsIntegrationsAwsGlue)
	PutAwsKinesisAnalytics(value *CloudAwsIntegrationsAwsKinesisAnalytics)
	PutAwsMediaConvert(value *CloudAwsIntegrationsAwsMediaConvert)
	PutAwsMediaPackageVod(value *CloudAwsIntegrationsAwsMediaPackageVod)
	PutAwsMq(value *CloudAwsIntegrationsAwsMq)
	PutAwsMsk(value *CloudAwsIntegrationsAwsMsk)
	PutAwsNeptune(value *CloudAwsIntegrationsAwsNeptune)
	PutAwsQldb(value *CloudAwsIntegrationsAwsQldb)
	PutAwsRoute53Resolver(value *CloudAwsIntegrationsAwsRoute53Resolver)
	PutAwsStates(value *CloudAwsIntegrationsAwsStates)
	PutAwsTransitGateway(value *CloudAwsIntegrationsAwsTransitGateway)
	PutAwsWaf(value *CloudAwsIntegrationsAwsWaf)
	PutAwsWafv2(value *CloudAwsIntegrationsAwsWafv2)
	PutBilling(value *CloudAwsIntegrationsBilling)
	PutCloudfront(value *CloudAwsIntegrationsCloudfront)
	PutCloudtrail(value *CloudAwsIntegrationsCloudtrail)
	PutDocDb(value *CloudAwsIntegrationsDocDb)
	PutDynamodb(value *CloudAwsIntegrationsDynamodb)
	PutEbs(value *CloudAwsIntegrationsEbs)
	PutEc2(value *CloudAwsIntegrationsEc2)
	PutEcs(value *CloudAwsIntegrationsEcs)
	PutEfs(value *CloudAwsIntegrationsEfs)
	PutElasticache(value *CloudAwsIntegrationsElasticache)
	PutElasticbeanstalk(value *CloudAwsIntegrationsElasticbeanstalk)
	PutElasticsearch(value *CloudAwsIntegrationsElasticsearch)
	PutElb(value *CloudAwsIntegrationsElb)
	PutEmr(value *CloudAwsIntegrationsEmr)
	PutHealth(value *CloudAwsIntegrationsHealth)
	PutIam(value *CloudAwsIntegrationsIam)
	PutIot(value *CloudAwsIntegrationsIot)
	PutKinesis(value *CloudAwsIntegrationsKinesis)
	PutKinesisFirehose(value *CloudAwsIntegrationsKinesisFirehose)
	PutLambda(value *CloudAwsIntegrationsLambda)
	PutRds(value *CloudAwsIntegrationsRds)
	PutRedshift(value *CloudAwsIntegrationsRedshift)
	PutRoute53(value *CloudAwsIntegrationsRoute53)
	PutS3(value *CloudAwsIntegrationsS3)
	PutSes(value *CloudAwsIntegrationsSes)
	PutSns(value *CloudAwsIntegrationsSns)
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
	ResetAwsAutoDiscovery()
	ResetAwsCognito()
	ResetAwsConnect()
	ResetAwsDirectConnect()
	ResetAwsFsx()
	ResetAwsGlue()
	ResetAwsKinesisAnalytics()
	ResetAwsMediaConvert()
	ResetAwsMediaPackageVod()
	ResetAwsMq()
	ResetAwsMsk()
	ResetAwsNeptune()
	ResetAwsQldb()
	ResetAwsRoute53Resolver()
	ResetAwsStates()
	ResetAwsTransitGateway()
	ResetAwsWaf()
	ResetAwsWafv2()
	ResetBilling()
	ResetCloudfront()
	ResetCloudtrail()
	ResetDocDb()
	ResetDynamodb()
	ResetEbs()
	ResetEc2()
	ResetEcs()
	ResetEfs()
	ResetElasticache()
	ResetElasticbeanstalk()
	ResetElasticsearch()
	ResetElb()
	ResetEmr()
	ResetHealth()
	ResetIam()
	ResetId()
	ResetIot()
	ResetKinesis()
	ResetKinesisFirehose()
	ResetLambda()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRds()
	ResetRedshift()
	ResetRoute53()
	ResetS3()
	ResetSes()
	ResetSns()
	ResetSqs()
	ResetTrustedAdvisor()
	ResetVpc()
	ResetXRay()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
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

func (j *jsiiProxy_CloudAwsIntegrations) AwsAutoDiscovery() CloudAwsIntegrationsAwsAutoDiscoveryOutputReference {
	var returns CloudAwsIntegrationsAwsAutoDiscoveryOutputReference
	_jsii_.Get(
		j,
		"awsAutoDiscovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsAutoDiscoveryInput() *CloudAwsIntegrationsAwsAutoDiscovery {
	var returns *CloudAwsIntegrationsAwsAutoDiscovery
	_jsii_.Get(
		j,
		"awsAutoDiscoveryInput",
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

func (j *jsiiProxy_CloudAwsIntegrations) AwsGlue() CloudAwsIntegrationsAwsGlueOutputReference {
	var returns CloudAwsIntegrationsAwsGlueOutputReference
	_jsii_.Get(
		j,
		"awsGlue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsGlueInput() *CloudAwsIntegrationsAwsGlue {
	var returns *CloudAwsIntegrationsAwsGlue
	_jsii_.Get(
		j,
		"awsGlueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsKinesisAnalytics() CloudAwsIntegrationsAwsKinesisAnalyticsOutputReference {
	var returns CloudAwsIntegrationsAwsKinesisAnalyticsOutputReference
	_jsii_.Get(
		j,
		"awsKinesisAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsKinesisAnalyticsInput() *CloudAwsIntegrationsAwsKinesisAnalytics {
	var returns *CloudAwsIntegrationsAwsKinesisAnalytics
	_jsii_.Get(
		j,
		"awsKinesisAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsMediaConvert() CloudAwsIntegrationsAwsMediaConvertOutputReference {
	var returns CloudAwsIntegrationsAwsMediaConvertOutputReference
	_jsii_.Get(
		j,
		"awsMediaConvert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsMediaConvertInput() *CloudAwsIntegrationsAwsMediaConvert {
	var returns *CloudAwsIntegrationsAwsMediaConvert
	_jsii_.Get(
		j,
		"awsMediaConvertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsMediaPackageVod() CloudAwsIntegrationsAwsMediaPackageVodOutputReference {
	var returns CloudAwsIntegrationsAwsMediaPackageVodOutputReference
	_jsii_.Get(
		j,
		"awsMediaPackageVod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsMediaPackageVodInput() *CloudAwsIntegrationsAwsMediaPackageVod {
	var returns *CloudAwsIntegrationsAwsMediaPackageVod
	_jsii_.Get(
		j,
		"awsMediaPackageVodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsMq() CloudAwsIntegrationsAwsMqOutputReference {
	var returns CloudAwsIntegrationsAwsMqOutputReference
	_jsii_.Get(
		j,
		"awsMq",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsMqInput() *CloudAwsIntegrationsAwsMq {
	var returns *CloudAwsIntegrationsAwsMq
	_jsii_.Get(
		j,
		"awsMqInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsMsk() CloudAwsIntegrationsAwsMskOutputReference {
	var returns CloudAwsIntegrationsAwsMskOutputReference
	_jsii_.Get(
		j,
		"awsMsk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsMskInput() *CloudAwsIntegrationsAwsMsk {
	var returns *CloudAwsIntegrationsAwsMsk
	_jsii_.Get(
		j,
		"awsMskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsNeptune() CloudAwsIntegrationsAwsNeptuneOutputReference {
	var returns CloudAwsIntegrationsAwsNeptuneOutputReference
	_jsii_.Get(
		j,
		"awsNeptune",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsNeptuneInput() *CloudAwsIntegrationsAwsNeptune {
	var returns *CloudAwsIntegrationsAwsNeptune
	_jsii_.Get(
		j,
		"awsNeptuneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsQldb() CloudAwsIntegrationsAwsQldbOutputReference {
	var returns CloudAwsIntegrationsAwsQldbOutputReference
	_jsii_.Get(
		j,
		"awsQldb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsQldbInput() *CloudAwsIntegrationsAwsQldb {
	var returns *CloudAwsIntegrationsAwsQldb
	_jsii_.Get(
		j,
		"awsQldbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsRoute53Resolver() CloudAwsIntegrationsAwsRoute53ResolverOutputReference {
	var returns CloudAwsIntegrationsAwsRoute53ResolverOutputReference
	_jsii_.Get(
		j,
		"awsRoute53Resolver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsRoute53ResolverInput() *CloudAwsIntegrationsAwsRoute53Resolver {
	var returns *CloudAwsIntegrationsAwsRoute53Resolver
	_jsii_.Get(
		j,
		"awsRoute53ResolverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsStates() CloudAwsIntegrationsAwsStatesOutputReference {
	var returns CloudAwsIntegrationsAwsStatesOutputReference
	_jsii_.Get(
		j,
		"awsStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsStatesInput() *CloudAwsIntegrationsAwsStates {
	var returns *CloudAwsIntegrationsAwsStates
	_jsii_.Get(
		j,
		"awsStatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsTransitGateway() CloudAwsIntegrationsAwsTransitGatewayOutputReference {
	var returns CloudAwsIntegrationsAwsTransitGatewayOutputReference
	_jsii_.Get(
		j,
		"awsTransitGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsTransitGatewayInput() *CloudAwsIntegrationsAwsTransitGateway {
	var returns *CloudAwsIntegrationsAwsTransitGateway
	_jsii_.Get(
		j,
		"awsTransitGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsWaf() CloudAwsIntegrationsAwsWafOutputReference {
	var returns CloudAwsIntegrationsAwsWafOutputReference
	_jsii_.Get(
		j,
		"awsWaf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsWafInput() *CloudAwsIntegrationsAwsWaf {
	var returns *CloudAwsIntegrationsAwsWaf
	_jsii_.Get(
		j,
		"awsWafInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsWafv2() CloudAwsIntegrationsAwsWafv2OutputReference {
	var returns CloudAwsIntegrationsAwsWafv2OutputReference
	_jsii_.Get(
		j,
		"awsWafv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) AwsWafv2Input() *CloudAwsIntegrationsAwsWafv2 {
	var returns *CloudAwsIntegrationsAwsWafv2
	_jsii_.Get(
		j,
		"awsWafv2Input",
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

func (j *jsiiProxy_CloudAwsIntegrations) Cloudfront() CloudAwsIntegrationsCloudfrontOutputReference {
	var returns CloudAwsIntegrationsCloudfrontOutputReference
	_jsii_.Get(
		j,
		"cloudfront",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) CloudfrontInput() *CloudAwsIntegrationsCloudfront {
	var returns *CloudAwsIntegrationsCloudfront
	_jsii_.Get(
		j,
		"cloudfrontInput",
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

func (j *jsiiProxy_CloudAwsIntegrations) Dynamodb() CloudAwsIntegrationsDynamodbOutputReference {
	var returns CloudAwsIntegrationsDynamodbOutputReference
	_jsii_.Get(
		j,
		"dynamodb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) DynamodbInput() *CloudAwsIntegrationsDynamodb {
	var returns *CloudAwsIntegrationsDynamodb
	_jsii_.Get(
		j,
		"dynamodbInput",
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

func (j *jsiiProxy_CloudAwsIntegrations) Ec2() CloudAwsIntegrationsEc2OutputReference {
	var returns CloudAwsIntegrationsEc2OutputReference
	_jsii_.Get(
		j,
		"ec2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Ec2Input() *CloudAwsIntegrationsEc2 {
	var returns *CloudAwsIntegrationsEc2
	_jsii_.Get(
		j,
		"ec2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Ecs() CloudAwsIntegrationsEcsOutputReference {
	var returns CloudAwsIntegrationsEcsOutputReference
	_jsii_.Get(
		j,
		"ecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) EcsInput() *CloudAwsIntegrationsEcs {
	var returns *CloudAwsIntegrationsEcs
	_jsii_.Get(
		j,
		"ecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Efs() CloudAwsIntegrationsEfsOutputReference {
	var returns CloudAwsIntegrationsEfsOutputReference
	_jsii_.Get(
		j,
		"efs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) EfsInput() *CloudAwsIntegrationsEfs {
	var returns *CloudAwsIntegrationsEfs
	_jsii_.Get(
		j,
		"efsInput",
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

func (j *jsiiProxy_CloudAwsIntegrations) Elasticbeanstalk() CloudAwsIntegrationsElasticbeanstalkOutputReference {
	var returns CloudAwsIntegrationsElasticbeanstalkOutputReference
	_jsii_.Get(
		j,
		"elasticbeanstalk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) ElasticbeanstalkInput() *CloudAwsIntegrationsElasticbeanstalk {
	var returns *CloudAwsIntegrationsElasticbeanstalk
	_jsii_.Get(
		j,
		"elasticbeanstalkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Elasticsearch() CloudAwsIntegrationsElasticsearchOutputReference {
	var returns CloudAwsIntegrationsElasticsearchOutputReference
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) ElasticsearchInput() *CloudAwsIntegrationsElasticsearch {
	var returns *CloudAwsIntegrationsElasticsearch
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Elb() CloudAwsIntegrationsElbOutputReference {
	var returns CloudAwsIntegrationsElbOutputReference
	_jsii_.Get(
		j,
		"elb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) ElbInput() *CloudAwsIntegrationsElb {
	var returns *CloudAwsIntegrationsElb
	_jsii_.Get(
		j,
		"elbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Emr() CloudAwsIntegrationsEmrOutputReference {
	var returns CloudAwsIntegrationsEmrOutputReference
	_jsii_.Get(
		j,
		"emr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) EmrInput() *CloudAwsIntegrationsEmr {
	var returns *CloudAwsIntegrationsEmr
	_jsii_.Get(
		j,
		"emrInput",
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

func (j *jsiiProxy_CloudAwsIntegrations) Iam() CloudAwsIntegrationsIamOutputReference {
	var returns CloudAwsIntegrationsIamOutputReference
	_jsii_.Get(
		j,
		"iam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) IamInput() *CloudAwsIntegrationsIam {
	var returns *CloudAwsIntegrationsIam
	_jsii_.Get(
		j,
		"iamInput",
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

func (j *jsiiProxy_CloudAwsIntegrations) Iot() CloudAwsIntegrationsIotOutputReference {
	var returns CloudAwsIntegrationsIotOutputReference
	_jsii_.Get(
		j,
		"iot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) IotInput() *CloudAwsIntegrationsIot {
	var returns *CloudAwsIntegrationsIot
	_jsii_.Get(
		j,
		"iotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Kinesis() CloudAwsIntegrationsKinesisOutputReference {
	var returns CloudAwsIntegrationsKinesisOutputReference
	_jsii_.Get(
		j,
		"kinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) KinesisFirehose() CloudAwsIntegrationsKinesisFirehoseOutputReference {
	var returns CloudAwsIntegrationsKinesisFirehoseOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) KinesisFirehoseInput() *CloudAwsIntegrationsKinesisFirehose {
	var returns *CloudAwsIntegrationsKinesisFirehose
	_jsii_.Get(
		j,
		"kinesisFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) KinesisInput() *CloudAwsIntegrationsKinesis {
	var returns *CloudAwsIntegrationsKinesis
	_jsii_.Get(
		j,
		"kinesisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Lambda() CloudAwsIntegrationsLambdaOutputReference {
	var returns CloudAwsIntegrationsLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) LambdaInput() *CloudAwsIntegrationsLambda {
	var returns *CloudAwsIntegrationsLambda
	_jsii_.Get(
		j,
		"lambdaInput",
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

func (j *jsiiProxy_CloudAwsIntegrations) Rds() CloudAwsIntegrationsRdsOutputReference {
	var returns CloudAwsIntegrationsRdsOutputReference
	_jsii_.Get(
		j,
		"rds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) RdsInput() *CloudAwsIntegrationsRds {
	var returns *CloudAwsIntegrationsRds
	_jsii_.Get(
		j,
		"rdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Redshift() CloudAwsIntegrationsRedshiftOutputReference {
	var returns CloudAwsIntegrationsRedshiftOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) RedshiftInput() *CloudAwsIntegrationsRedshift {
	var returns *CloudAwsIntegrationsRedshift
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Route53() CloudAwsIntegrationsRoute53OutputReference {
	var returns CloudAwsIntegrationsRoute53OutputReference
	_jsii_.Get(
		j,
		"route53",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Route53Input() *CloudAwsIntegrationsRoute53 {
	var returns *CloudAwsIntegrationsRoute53
	_jsii_.Get(
		j,
		"route53Input",
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

func (j *jsiiProxy_CloudAwsIntegrations) Ses() CloudAwsIntegrationsSesOutputReference {
	var returns CloudAwsIntegrationsSesOutputReference
	_jsii_.Get(
		j,
		"ses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) SesInput() *CloudAwsIntegrationsSes {
	var returns *CloudAwsIntegrationsSes
	_jsii_.Get(
		j,
		"sesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) Sns() CloudAwsIntegrationsSnsOutputReference {
	var returns CloudAwsIntegrationsSnsOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsIntegrations) SnsInput() *CloudAwsIntegrationsSns {
	var returns *CloudAwsIntegrationsSns
	_jsii_.Get(
		j,
		"snsInput",
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


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.74.0/docs/resources/cloud_aws_integrations newrelic_cloud_aws_integrations} Resource.
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

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.74.0/docs/resources/cloud_aws_integrations newrelic_cloud_aws_integrations} Resource.
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

// Generates CDKTF code for importing a CloudAwsIntegrations resource upon running "cdktf plan <stack-name>".
func CloudAwsIntegrations_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudAwsIntegrations_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudAwsIntegrations.CloudAwsIntegrations",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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

func (c *jsiiProxy_CloudAwsIntegrations) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
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

func (c *jsiiProxy_CloudAwsIntegrations) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrations) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
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

func (c *jsiiProxy_CloudAwsIntegrations) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
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

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsAutoDiscovery(value *CloudAwsIntegrationsAwsAutoDiscovery) {
	if err := c.validatePutAwsAutoDiscoveryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsAutoDiscovery",
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

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsGlue(value *CloudAwsIntegrationsAwsGlue) {
	if err := c.validatePutAwsGlueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsGlue",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsKinesisAnalytics(value *CloudAwsIntegrationsAwsKinesisAnalytics) {
	if err := c.validatePutAwsKinesisAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsKinesisAnalytics",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsMediaConvert(value *CloudAwsIntegrationsAwsMediaConvert) {
	if err := c.validatePutAwsMediaConvertParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsMediaConvert",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsMediaPackageVod(value *CloudAwsIntegrationsAwsMediaPackageVod) {
	if err := c.validatePutAwsMediaPackageVodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsMediaPackageVod",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsMq(value *CloudAwsIntegrationsAwsMq) {
	if err := c.validatePutAwsMqParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsMq",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsMsk(value *CloudAwsIntegrationsAwsMsk) {
	if err := c.validatePutAwsMskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsMsk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsNeptune(value *CloudAwsIntegrationsAwsNeptune) {
	if err := c.validatePutAwsNeptuneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsNeptune",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsQldb(value *CloudAwsIntegrationsAwsQldb) {
	if err := c.validatePutAwsQldbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsQldb",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsRoute53Resolver(value *CloudAwsIntegrationsAwsRoute53Resolver) {
	if err := c.validatePutAwsRoute53ResolverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsRoute53Resolver",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsStates(value *CloudAwsIntegrationsAwsStates) {
	if err := c.validatePutAwsStatesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsStates",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsTransitGateway(value *CloudAwsIntegrationsAwsTransitGateway) {
	if err := c.validatePutAwsTransitGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsTransitGateway",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsWaf(value *CloudAwsIntegrationsAwsWaf) {
	if err := c.validatePutAwsWafParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsWaf",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutAwsWafv2(value *CloudAwsIntegrationsAwsWafv2) {
	if err := c.validatePutAwsWafv2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsWafv2",
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

func (c *jsiiProxy_CloudAwsIntegrations) PutCloudfront(value *CloudAwsIntegrationsCloudfront) {
	if err := c.validatePutCloudfrontParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCloudfront",
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

func (c *jsiiProxy_CloudAwsIntegrations) PutDynamodb(value *CloudAwsIntegrationsDynamodb) {
	if err := c.validatePutDynamodbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDynamodb",
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

func (c *jsiiProxy_CloudAwsIntegrations) PutEc2(value *CloudAwsIntegrationsEc2) {
	if err := c.validatePutEc2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEc2",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutEcs(value *CloudAwsIntegrationsEcs) {
	if err := c.validatePutEcsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEcs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutEfs(value *CloudAwsIntegrationsEfs) {
	if err := c.validatePutEfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEfs",
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

func (c *jsiiProxy_CloudAwsIntegrations) PutElasticbeanstalk(value *CloudAwsIntegrationsElasticbeanstalk) {
	if err := c.validatePutElasticbeanstalkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putElasticbeanstalk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutElasticsearch(value *CloudAwsIntegrationsElasticsearch) {
	if err := c.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutElb(value *CloudAwsIntegrationsElb) {
	if err := c.validatePutElbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putElb",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutEmr(value *CloudAwsIntegrationsEmr) {
	if err := c.validatePutEmrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEmr",
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

func (c *jsiiProxy_CloudAwsIntegrations) PutIam(value *CloudAwsIntegrationsIam) {
	if err := c.validatePutIamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIam",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutIot(value *CloudAwsIntegrationsIot) {
	if err := c.validatePutIotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIot",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutKinesis(value *CloudAwsIntegrationsKinesis) {
	if err := c.validatePutKinesisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKinesis",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutKinesisFirehose(value *CloudAwsIntegrationsKinesisFirehose) {
	if err := c.validatePutKinesisFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKinesisFirehose",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutLambda(value *CloudAwsIntegrationsLambda) {
	if err := c.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLambda",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutRds(value *CloudAwsIntegrationsRds) {
	if err := c.validatePutRdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRds",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutRedshift(value *CloudAwsIntegrationsRedshift) {
	if err := c.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRedshift",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutRoute53(value *CloudAwsIntegrationsRoute53) {
	if err := c.validatePutRoute53Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRoute53",
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

func (c *jsiiProxy_CloudAwsIntegrations) PutSes(value *CloudAwsIntegrationsSes) {
	if err := c.validatePutSesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) PutSns(value *CloudAwsIntegrationsSns) {
	if err := c.validatePutSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSns",
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

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsAutoDiscovery() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsAutoDiscovery",
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

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsGlue() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsGlue",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsKinesisAnalytics() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsKinesisAnalytics",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsMediaConvert() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsMediaConvert",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsMediaPackageVod() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsMediaPackageVod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsMq() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsMq",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsMsk() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsMsk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsNeptune() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsNeptune",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsQldb() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsQldb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsRoute53Resolver() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsRoute53Resolver",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsStates() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsStates",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsTransitGateway() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsTransitGateway",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsWaf() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsWaf",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetAwsWafv2() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsWafv2",
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

func (c *jsiiProxy_CloudAwsIntegrations) ResetCloudfront() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudfront",
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

func (c *jsiiProxy_CloudAwsIntegrations) ResetDynamodb() {
	_jsii_.InvokeVoid(
		c,
		"resetDynamodb",
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

func (c *jsiiProxy_CloudAwsIntegrations) ResetEc2() {
	_jsii_.InvokeVoid(
		c,
		"resetEc2",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetEcs() {
	_jsii_.InvokeVoid(
		c,
		"resetEcs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetEfs() {
	_jsii_.InvokeVoid(
		c,
		"resetEfs",
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

func (c *jsiiProxy_CloudAwsIntegrations) ResetElasticbeanstalk() {
	_jsii_.InvokeVoid(
		c,
		"resetElasticbeanstalk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		c,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetElb() {
	_jsii_.InvokeVoid(
		c,
		"resetElb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetEmr() {
	_jsii_.InvokeVoid(
		c,
		"resetEmr",
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

func (c *jsiiProxy_CloudAwsIntegrations) ResetIam() {
	_jsii_.InvokeVoid(
		c,
		"resetIam",
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

func (c *jsiiProxy_CloudAwsIntegrations) ResetIot() {
	_jsii_.InvokeVoid(
		c,
		"resetIot",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetKinesis() {
	_jsii_.InvokeVoid(
		c,
		"resetKinesis",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetKinesisFirehose() {
	_jsii_.InvokeVoid(
		c,
		"resetKinesisFirehose",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetLambda() {
	_jsii_.InvokeVoid(
		c,
		"resetLambda",
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

func (c *jsiiProxy_CloudAwsIntegrations) ResetRds() {
	_jsii_.InvokeVoid(
		c,
		"resetRds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetRedshift() {
	_jsii_.InvokeVoid(
		c,
		"resetRedshift",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetRoute53() {
	_jsii_.InvokeVoid(
		c,
		"resetRoute53",
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

func (c *jsiiProxy_CloudAwsIntegrations) ResetSes() {
	_jsii_.InvokeVoid(
		c,
		"resetSes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsIntegrations) ResetSns() {
	_jsii_.InvokeVoid(
		c,
		"resetSns",
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

func (c *jsiiProxy_CloudAwsIntegrations) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsIntegrations) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
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

