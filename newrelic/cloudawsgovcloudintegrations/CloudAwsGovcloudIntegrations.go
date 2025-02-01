// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudawsgovcloudintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v12/cloudawsgovcloudintegrations/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/cloud_aws_govcloud_integrations newrelic_cloud_aws_govcloud_integrations}.
type CloudAwsGovcloudIntegrations interface {
	cdktf.TerraformResource
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
	Alb() CloudAwsGovcloudIntegrationsAlbOutputReference
	AlbInput() *CloudAwsGovcloudIntegrationsAlb
	ApiGateway() CloudAwsGovcloudIntegrationsApiGatewayOutputReference
	ApiGatewayInput() *CloudAwsGovcloudIntegrationsApiGateway
	AutoScaling() CloudAwsGovcloudIntegrationsAutoScalingOutputReference
	AutoScalingInput() *CloudAwsGovcloudIntegrationsAutoScaling
	AwsDirectConnect() CloudAwsGovcloudIntegrationsAwsDirectConnectOutputReference
	AwsDirectConnectInput() *CloudAwsGovcloudIntegrationsAwsDirectConnect
	AwsStates() CloudAwsGovcloudIntegrationsAwsStatesOutputReference
	AwsStatesInput() *CloudAwsGovcloudIntegrationsAwsStates
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Cloudtrail() CloudAwsGovcloudIntegrationsCloudtrailOutputReference
	CloudtrailInput() *CloudAwsGovcloudIntegrationsCloudtrail
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
	DynamoDb() CloudAwsGovcloudIntegrationsDynamoDbOutputReference
	DynamoDbInput() *CloudAwsGovcloudIntegrationsDynamoDb
	Ebs() CloudAwsGovcloudIntegrationsEbsOutputReference
	EbsInput() *CloudAwsGovcloudIntegrationsEbs
	Ec2() CloudAwsGovcloudIntegrationsEc2OutputReference
	Ec2Input() *CloudAwsGovcloudIntegrationsEc2
	ElasticSearch() CloudAwsGovcloudIntegrationsElasticSearchOutputReference
	ElasticSearchInput() *CloudAwsGovcloudIntegrationsElasticSearch
	Elb() CloudAwsGovcloudIntegrationsElbOutputReference
	ElbInput() *CloudAwsGovcloudIntegrationsElb
	Emr() CloudAwsGovcloudIntegrationsEmrOutputReference
	EmrInput() *CloudAwsGovcloudIntegrationsEmr
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Iam() CloudAwsGovcloudIntegrationsIamOutputReference
	IamInput() *CloudAwsGovcloudIntegrationsIam
	Id() *string
	SetId(val *string)
	IdInput() *string
	Lambda() CloudAwsGovcloudIntegrationsLambdaOutputReference
	LambdaInput() *CloudAwsGovcloudIntegrationsLambda
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
	Rds() CloudAwsGovcloudIntegrationsRdsOutputReference
	RdsInput() *CloudAwsGovcloudIntegrationsRds
	RedShift() CloudAwsGovcloudIntegrationsRedShiftOutputReference
	RedShiftInput() *CloudAwsGovcloudIntegrationsRedShift
	Route53() CloudAwsGovcloudIntegrationsRoute53OutputReference
	Route53Input() *CloudAwsGovcloudIntegrationsRoute53
	S3() CloudAwsGovcloudIntegrationsS3OutputReference
	S3Input() *CloudAwsGovcloudIntegrationsS3
	Sns() CloudAwsGovcloudIntegrationsSnsOutputReference
	SnsInput() *CloudAwsGovcloudIntegrationsSns
	Sqs() CloudAwsGovcloudIntegrationsSqsOutputReference
	SqsInput() *CloudAwsGovcloudIntegrationsSqs
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutAlb(value *CloudAwsGovcloudIntegrationsAlb)
	PutApiGateway(value *CloudAwsGovcloudIntegrationsApiGateway)
	PutAutoScaling(value *CloudAwsGovcloudIntegrationsAutoScaling)
	PutAwsDirectConnect(value *CloudAwsGovcloudIntegrationsAwsDirectConnect)
	PutAwsStates(value *CloudAwsGovcloudIntegrationsAwsStates)
	PutCloudtrail(value *CloudAwsGovcloudIntegrationsCloudtrail)
	PutDynamoDb(value *CloudAwsGovcloudIntegrationsDynamoDb)
	PutEbs(value *CloudAwsGovcloudIntegrationsEbs)
	PutEc2(value *CloudAwsGovcloudIntegrationsEc2)
	PutElasticSearch(value *CloudAwsGovcloudIntegrationsElasticSearch)
	PutElb(value *CloudAwsGovcloudIntegrationsElb)
	PutEmr(value *CloudAwsGovcloudIntegrationsEmr)
	PutIam(value *CloudAwsGovcloudIntegrationsIam)
	PutLambda(value *CloudAwsGovcloudIntegrationsLambda)
	PutRds(value *CloudAwsGovcloudIntegrationsRds)
	PutRedShift(value *CloudAwsGovcloudIntegrationsRedShift)
	PutRoute53(value *CloudAwsGovcloudIntegrationsRoute53)
	PutS3(value *CloudAwsGovcloudIntegrationsS3)
	PutSns(value *CloudAwsGovcloudIntegrationsSns)
	PutSqs(value *CloudAwsGovcloudIntegrationsSqs)
	ResetAccountId()
	ResetAlb()
	ResetApiGateway()
	ResetAutoScaling()
	ResetAwsDirectConnect()
	ResetAwsStates()
	ResetCloudtrail()
	ResetDynamoDb()
	ResetEbs()
	ResetEc2()
	ResetElasticSearch()
	ResetElb()
	ResetEmr()
	ResetIam()
	ResetId()
	ResetLambda()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRds()
	ResetRedShift()
	ResetRoute53()
	ResetS3()
	ResetSns()
	ResetSqs()
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

// The jsii proxy struct for CloudAwsGovcloudIntegrations
type jsiiProxy_CloudAwsGovcloudIntegrations struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Alb() CloudAwsGovcloudIntegrationsAlbOutputReference {
	var returns CloudAwsGovcloudIntegrationsAlbOutputReference
	_jsii_.Get(
		j,
		"alb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) AlbInput() *CloudAwsGovcloudIntegrationsAlb {
	var returns *CloudAwsGovcloudIntegrationsAlb
	_jsii_.Get(
		j,
		"albInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) ApiGateway() CloudAwsGovcloudIntegrationsApiGatewayOutputReference {
	var returns CloudAwsGovcloudIntegrationsApiGatewayOutputReference
	_jsii_.Get(
		j,
		"apiGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) ApiGatewayInput() *CloudAwsGovcloudIntegrationsApiGateway {
	var returns *CloudAwsGovcloudIntegrationsApiGateway
	_jsii_.Get(
		j,
		"apiGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) AutoScaling() CloudAwsGovcloudIntegrationsAutoScalingOutputReference {
	var returns CloudAwsGovcloudIntegrationsAutoScalingOutputReference
	_jsii_.Get(
		j,
		"autoScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) AutoScalingInput() *CloudAwsGovcloudIntegrationsAutoScaling {
	var returns *CloudAwsGovcloudIntegrationsAutoScaling
	_jsii_.Get(
		j,
		"autoScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) AwsDirectConnect() CloudAwsGovcloudIntegrationsAwsDirectConnectOutputReference {
	var returns CloudAwsGovcloudIntegrationsAwsDirectConnectOutputReference
	_jsii_.Get(
		j,
		"awsDirectConnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) AwsDirectConnectInput() *CloudAwsGovcloudIntegrationsAwsDirectConnect {
	var returns *CloudAwsGovcloudIntegrationsAwsDirectConnect
	_jsii_.Get(
		j,
		"awsDirectConnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) AwsStates() CloudAwsGovcloudIntegrationsAwsStatesOutputReference {
	var returns CloudAwsGovcloudIntegrationsAwsStatesOutputReference
	_jsii_.Get(
		j,
		"awsStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) AwsStatesInput() *CloudAwsGovcloudIntegrationsAwsStates {
	var returns *CloudAwsGovcloudIntegrationsAwsStates
	_jsii_.Get(
		j,
		"awsStatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Cloudtrail() CloudAwsGovcloudIntegrationsCloudtrailOutputReference {
	var returns CloudAwsGovcloudIntegrationsCloudtrailOutputReference
	_jsii_.Get(
		j,
		"cloudtrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) CloudtrailInput() *CloudAwsGovcloudIntegrationsCloudtrail {
	var returns *CloudAwsGovcloudIntegrationsCloudtrail
	_jsii_.Get(
		j,
		"cloudtrailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) DynamoDb() CloudAwsGovcloudIntegrationsDynamoDbOutputReference {
	var returns CloudAwsGovcloudIntegrationsDynamoDbOutputReference
	_jsii_.Get(
		j,
		"dynamoDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) DynamoDbInput() *CloudAwsGovcloudIntegrationsDynamoDb {
	var returns *CloudAwsGovcloudIntegrationsDynamoDb
	_jsii_.Get(
		j,
		"dynamoDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Ebs() CloudAwsGovcloudIntegrationsEbsOutputReference {
	var returns CloudAwsGovcloudIntegrationsEbsOutputReference
	_jsii_.Get(
		j,
		"ebs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) EbsInput() *CloudAwsGovcloudIntegrationsEbs {
	var returns *CloudAwsGovcloudIntegrationsEbs
	_jsii_.Get(
		j,
		"ebsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Ec2() CloudAwsGovcloudIntegrationsEc2OutputReference {
	var returns CloudAwsGovcloudIntegrationsEc2OutputReference
	_jsii_.Get(
		j,
		"ec2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Ec2Input() *CloudAwsGovcloudIntegrationsEc2 {
	var returns *CloudAwsGovcloudIntegrationsEc2
	_jsii_.Get(
		j,
		"ec2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) ElasticSearch() CloudAwsGovcloudIntegrationsElasticSearchOutputReference {
	var returns CloudAwsGovcloudIntegrationsElasticSearchOutputReference
	_jsii_.Get(
		j,
		"elasticSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) ElasticSearchInput() *CloudAwsGovcloudIntegrationsElasticSearch {
	var returns *CloudAwsGovcloudIntegrationsElasticSearch
	_jsii_.Get(
		j,
		"elasticSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Elb() CloudAwsGovcloudIntegrationsElbOutputReference {
	var returns CloudAwsGovcloudIntegrationsElbOutputReference
	_jsii_.Get(
		j,
		"elb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) ElbInput() *CloudAwsGovcloudIntegrationsElb {
	var returns *CloudAwsGovcloudIntegrationsElb
	_jsii_.Get(
		j,
		"elbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Emr() CloudAwsGovcloudIntegrationsEmrOutputReference {
	var returns CloudAwsGovcloudIntegrationsEmrOutputReference
	_jsii_.Get(
		j,
		"emr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) EmrInput() *CloudAwsGovcloudIntegrationsEmr {
	var returns *CloudAwsGovcloudIntegrationsEmr
	_jsii_.Get(
		j,
		"emrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Iam() CloudAwsGovcloudIntegrationsIamOutputReference {
	var returns CloudAwsGovcloudIntegrationsIamOutputReference
	_jsii_.Get(
		j,
		"iam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) IamInput() *CloudAwsGovcloudIntegrationsIam {
	var returns *CloudAwsGovcloudIntegrationsIam
	_jsii_.Get(
		j,
		"iamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Lambda() CloudAwsGovcloudIntegrationsLambdaOutputReference {
	var returns CloudAwsGovcloudIntegrationsLambdaOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) LambdaInput() *CloudAwsGovcloudIntegrationsLambda {
	var returns *CloudAwsGovcloudIntegrationsLambda
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) LinkedAccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"linkedAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) LinkedAccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"linkedAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Rds() CloudAwsGovcloudIntegrationsRdsOutputReference {
	var returns CloudAwsGovcloudIntegrationsRdsOutputReference
	_jsii_.Get(
		j,
		"rds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) RdsInput() *CloudAwsGovcloudIntegrationsRds {
	var returns *CloudAwsGovcloudIntegrationsRds
	_jsii_.Get(
		j,
		"rdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) RedShift() CloudAwsGovcloudIntegrationsRedShiftOutputReference {
	var returns CloudAwsGovcloudIntegrationsRedShiftOutputReference
	_jsii_.Get(
		j,
		"redShift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) RedShiftInput() *CloudAwsGovcloudIntegrationsRedShift {
	var returns *CloudAwsGovcloudIntegrationsRedShift
	_jsii_.Get(
		j,
		"redShiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Route53() CloudAwsGovcloudIntegrationsRoute53OutputReference {
	var returns CloudAwsGovcloudIntegrationsRoute53OutputReference
	_jsii_.Get(
		j,
		"route53",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Route53Input() *CloudAwsGovcloudIntegrationsRoute53 {
	var returns *CloudAwsGovcloudIntegrationsRoute53
	_jsii_.Get(
		j,
		"route53Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) S3() CloudAwsGovcloudIntegrationsS3OutputReference {
	var returns CloudAwsGovcloudIntegrationsS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) S3Input() *CloudAwsGovcloudIntegrationsS3 {
	var returns *CloudAwsGovcloudIntegrationsS3
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Sns() CloudAwsGovcloudIntegrationsSnsOutputReference {
	var returns CloudAwsGovcloudIntegrationsSnsOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) SnsInput() *CloudAwsGovcloudIntegrationsSns {
	var returns *CloudAwsGovcloudIntegrationsSns
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) Sqs() CloudAwsGovcloudIntegrationsSqsOutputReference {
	var returns CloudAwsGovcloudIntegrationsSqsOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) SqsInput() *CloudAwsGovcloudIntegrationsSqs {
	var returns *CloudAwsGovcloudIntegrationsSqs
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/cloud_aws_govcloud_integrations newrelic_cloud_aws_govcloud_integrations} Resource.
func NewCloudAwsGovcloudIntegrations(scope constructs.Construct, id *string, config *CloudAwsGovcloudIntegrationsConfig) CloudAwsGovcloudIntegrations {
	_init_.Initialize()

	if err := validateNewCloudAwsGovcloudIntegrationsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudAwsGovcloudIntegrations{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.cloudAwsGovcloudIntegrations.CloudAwsGovcloudIntegrations",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.54.1/docs/resources/cloud_aws_govcloud_integrations newrelic_cloud_aws_govcloud_integrations} Resource.
func NewCloudAwsGovcloudIntegrations_Override(c CloudAwsGovcloudIntegrations, scope constructs.Construct, id *string, config *CloudAwsGovcloudIntegrationsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.cloudAwsGovcloudIntegrations.CloudAwsGovcloudIntegrations",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations)SetAccountId(val *float64) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations)SetLinkedAccountId(val *float64) {
	if err := j.validateSetLinkedAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkedAccountId",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudAwsGovcloudIntegrations)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a CloudAwsGovcloudIntegrations resource upon running "cdktf plan <stack-name>".
func CloudAwsGovcloudIntegrations_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudAwsGovcloudIntegrations_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudAwsGovcloudIntegrations.CloudAwsGovcloudIntegrations",
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
func CloudAwsGovcloudIntegrations_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAwsGovcloudIntegrations_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudAwsGovcloudIntegrations.CloudAwsGovcloudIntegrations",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudAwsGovcloudIntegrations_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAwsGovcloudIntegrations_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudAwsGovcloudIntegrations.CloudAwsGovcloudIntegrations",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudAwsGovcloudIntegrations_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAwsGovcloudIntegrations_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudAwsGovcloudIntegrations.CloudAwsGovcloudIntegrations",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudAwsGovcloudIntegrations_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.cloudAwsGovcloudIntegrations.CloudAwsGovcloudIntegrations",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutAlb(value *CloudAwsGovcloudIntegrationsAlb) {
	if err := c.validatePutAlbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAlb",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutApiGateway(value *CloudAwsGovcloudIntegrationsApiGateway) {
	if err := c.validatePutApiGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putApiGateway",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutAutoScaling(value *CloudAwsGovcloudIntegrationsAutoScaling) {
	if err := c.validatePutAutoScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutoScaling",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutAwsDirectConnect(value *CloudAwsGovcloudIntegrationsAwsDirectConnect) {
	if err := c.validatePutAwsDirectConnectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsDirectConnect",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutAwsStates(value *CloudAwsGovcloudIntegrationsAwsStates) {
	if err := c.validatePutAwsStatesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAwsStates",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutCloudtrail(value *CloudAwsGovcloudIntegrationsCloudtrail) {
	if err := c.validatePutCloudtrailParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCloudtrail",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutDynamoDb(value *CloudAwsGovcloudIntegrationsDynamoDb) {
	if err := c.validatePutDynamoDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDynamoDb",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutEbs(value *CloudAwsGovcloudIntegrationsEbs) {
	if err := c.validatePutEbsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEbs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutEc2(value *CloudAwsGovcloudIntegrationsEc2) {
	if err := c.validatePutEc2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEc2",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutElasticSearch(value *CloudAwsGovcloudIntegrationsElasticSearch) {
	if err := c.validatePutElasticSearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putElasticSearch",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutElb(value *CloudAwsGovcloudIntegrationsElb) {
	if err := c.validatePutElbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putElb",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutEmr(value *CloudAwsGovcloudIntegrationsEmr) {
	if err := c.validatePutEmrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEmr",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutIam(value *CloudAwsGovcloudIntegrationsIam) {
	if err := c.validatePutIamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIam",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutLambda(value *CloudAwsGovcloudIntegrationsLambda) {
	if err := c.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLambda",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutRds(value *CloudAwsGovcloudIntegrationsRds) {
	if err := c.validatePutRdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRds",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutRedShift(value *CloudAwsGovcloudIntegrationsRedShift) {
	if err := c.validatePutRedShiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRedShift",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutRoute53(value *CloudAwsGovcloudIntegrationsRoute53) {
	if err := c.validatePutRoute53Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRoute53",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutS3(value *CloudAwsGovcloudIntegrationsS3) {
	if err := c.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putS3",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutSns(value *CloudAwsGovcloudIntegrationsSns) {
	if err := c.validatePutSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSns",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) PutSqs(value *CloudAwsGovcloudIntegrationsSqs) {
	if err := c.validatePutSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSqs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetAccountId() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetAlb() {
	_jsii_.InvokeVoid(
		c,
		"resetAlb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetApiGateway() {
	_jsii_.InvokeVoid(
		c,
		"resetApiGateway",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetAutoScaling() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoScaling",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetAwsDirectConnect() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsDirectConnect",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetAwsStates() {
	_jsii_.InvokeVoid(
		c,
		"resetAwsStates",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetCloudtrail() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudtrail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetDynamoDb() {
	_jsii_.InvokeVoid(
		c,
		"resetDynamoDb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetEbs() {
	_jsii_.InvokeVoid(
		c,
		"resetEbs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetEc2() {
	_jsii_.InvokeVoid(
		c,
		"resetEc2",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetElasticSearch() {
	_jsii_.InvokeVoid(
		c,
		"resetElasticSearch",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetElb() {
	_jsii_.InvokeVoid(
		c,
		"resetElb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetEmr() {
	_jsii_.InvokeVoid(
		c,
		"resetEmr",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetIam() {
	_jsii_.InvokeVoid(
		c,
		"resetIam",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetLambda() {
	_jsii_.InvokeVoid(
		c,
		"resetLambda",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetRds() {
	_jsii_.InvokeVoid(
		c,
		"resetRds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetRedShift() {
	_jsii_.InvokeVoid(
		c,
		"resetRedShift",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetRoute53() {
	_jsii_.InvokeVoid(
		c,
		"resetRoute53",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetS3() {
	_jsii_.InvokeVoid(
		c,
		"resetS3",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetSns() {
	_jsii_.InvokeVoid(
		c,
		"resetSns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ResetSqs() {
	_jsii_.InvokeVoid(
		c,
		"resetSqs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAwsGovcloudIntegrations) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

