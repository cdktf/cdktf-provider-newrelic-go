// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-newrelic-go/newrelic/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-newrelic-go/newrelic/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_gcp_integrations newrelic_cloud_gcp_integrations}.
type CloudGcpIntegrations interface {
	cdktf.TerraformResource
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
	AppEngine() CloudGcpIntegrationsAppEngineOutputReference
	AppEngineInput() *CloudGcpIntegrationsAppEngine
	BigQuery() CloudGcpIntegrationsBigQueryOutputReference
	BigQueryInput() *CloudGcpIntegrationsBigQuery
	BigTable() CloudGcpIntegrationsBigTableOutputReference
	BigTableInput() *CloudGcpIntegrationsBigTable
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Composer() CloudGcpIntegrationsComposerOutputReference
	ComposerInput() *CloudGcpIntegrationsComposer
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DataFlow() CloudGcpIntegrationsDataFlowOutputReference
	DataFlowInput() *CloudGcpIntegrationsDataFlow
	DataProc() CloudGcpIntegrationsDataProcOutputReference
	DataProcInput() *CloudGcpIntegrationsDataProc
	DataStore() CloudGcpIntegrationsDataStoreOutputReference
	DataStoreInput() *CloudGcpIntegrationsDataStore
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FireBaseDatabase() CloudGcpIntegrationsFireBaseDatabaseOutputReference
	FireBaseDatabaseInput() *CloudGcpIntegrationsFireBaseDatabase
	FireBaseHosting() CloudGcpIntegrationsFireBaseHostingOutputReference
	FireBaseHostingInput() *CloudGcpIntegrationsFireBaseHosting
	FireBaseStorage() CloudGcpIntegrationsFireBaseStorageOutputReference
	FireBaseStorageInput() *CloudGcpIntegrationsFireBaseStorage
	FireStore() CloudGcpIntegrationsFireStoreOutputReference
	FireStoreInput() *CloudGcpIntegrationsFireStore
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Functions() CloudGcpIntegrationsFunctionsOutputReference
	FunctionsInput() *CloudGcpIntegrationsFunctions
	Id() *string
	SetId(val *string)
	IdInput() *string
	Interconnect() CloudGcpIntegrationsInterconnectOutputReference
	InterconnectInput() *CloudGcpIntegrationsInterconnect
	Kubernetes() CloudGcpIntegrationsKubernetesOutputReference
	KubernetesInput() *CloudGcpIntegrationsKubernetes
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinkedAccountId() *float64
	SetLinkedAccountId(val *float64)
	LinkedAccountIdInput() *float64
	LoadBalancing() CloudGcpIntegrationsLoadBalancingOutputReference
	LoadBalancingInput() *CloudGcpIntegrationsLoadBalancing
	MemCache() CloudGcpIntegrationsMemCacheOutputReference
	MemCacheInput() *CloudGcpIntegrationsMemCache
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
	PubSub() CloudGcpIntegrationsPubSubOutputReference
	PubSubInput() *CloudGcpIntegrationsPubSub
	// Experimental.
	RawOverrides() interface{}
	Redis() CloudGcpIntegrationsRedisOutputReference
	RedisInput() *CloudGcpIntegrationsRedis
	Router() CloudGcpIntegrationsRouterOutputReference
	RouterInput() *CloudGcpIntegrationsRouter
	Run() CloudGcpIntegrationsRunOutputReference
	RunInput() *CloudGcpIntegrationsRun
	Spanner() CloudGcpIntegrationsSpannerOutputReference
	SpannerInput() *CloudGcpIntegrationsSpanner
	Sql() CloudGcpIntegrationsSqlOutputReference
	SqlInput() *CloudGcpIntegrationsSql
	Storage() CloudGcpIntegrationsStorageOutputReference
	StorageInput() *CloudGcpIntegrationsStorage
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VirtualMachines() CloudGcpIntegrationsVirtualMachinesOutputReference
	VirtualMachinesInput() *CloudGcpIntegrationsVirtualMachines
	VpcAccess() CloudGcpIntegrationsVpcAccessOutputReference
	VpcAccessInput() *CloudGcpIntegrationsVpcAccess
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
	PutAppEngine(value *CloudGcpIntegrationsAppEngine)
	PutBigQuery(value *CloudGcpIntegrationsBigQuery)
	PutBigTable(value *CloudGcpIntegrationsBigTable)
	PutComposer(value *CloudGcpIntegrationsComposer)
	PutDataFlow(value *CloudGcpIntegrationsDataFlow)
	PutDataProc(value *CloudGcpIntegrationsDataProc)
	PutDataStore(value *CloudGcpIntegrationsDataStore)
	PutFireBaseDatabase(value *CloudGcpIntegrationsFireBaseDatabase)
	PutFireBaseHosting(value *CloudGcpIntegrationsFireBaseHosting)
	PutFireBaseStorage(value *CloudGcpIntegrationsFireBaseStorage)
	PutFireStore(value *CloudGcpIntegrationsFireStore)
	PutFunctions(value *CloudGcpIntegrationsFunctions)
	PutInterconnect(value *CloudGcpIntegrationsInterconnect)
	PutKubernetes(value *CloudGcpIntegrationsKubernetes)
	PutLoadBalancing(value *CloudGcpIntegrationsLoadBalancing)
	PutMemCache(value *CloudGcpIntegrationsMemCache)
	PutPubSub(value *CloudGcpIntegrationsPubSub)
	PutRedis(value *CloudGcpIntegrationsRedis)
	PutRouter(value *CloudGcpIntegrationsRouter)
	PutRun(value *CloudGcpIntegrationsRun)
	PutSpanner(value *CloudGcpIntegrationsSpanner)
	PutSql(value *CloudGcpIntegrationsSql)
	PutStorage(value *CloudGcpIntegrationsStorage)
	PutVirtualMachines(value *CloudGcpIntegrationsVirtualMachines)
	PutVpcAccess(value *CloudGcpIntegrationsVpcAccess)
	ResetAccountId()
	ResetAppEngine()
	ResetBigQuery()
	ResetBigTable()
	ResetComposer()
	ResetDataFlow()
	ResetDataProc()
	ResetDataStore()
	ResetFireBaseDatabase()
	ResetFireBaseHosting()
	ResetFireBaseStorage()
	ResetFireStore()
	ResetFunctions()
	ResetId()
	ResetInterconnect()
	ResetKubernetes()
	ResetLoadBalancing()
	ResetMemCache()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPubSub()
	ResetRedis()
	ResetRouter()
	ResetRun()
	ResetSpanner()
	ResetSql()
	ResetStorage()
	ResetVirtualMachines()
	ResetVpcAccess()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for CloudGcpIntegrations
type jsiiProxy_CloudGcpIntegrations struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudGcpIntegrations) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) AppEngine() CloudGcpIntegrationsAppEngineOutputReference {
	var returns CloudGcpIntegrationsAppEngineOutputReference
	_jsii_.Get(
		j,
		"appEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) AppEngineInput() *CloudGcpIntegrationsAppEngine {
	var returns *CloudGcpIntegrationsAppEngine
	_jsii_.Get(
		j,
		"appEngineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) BigQuery() CloudGcpIntegrationsBigQueryOutputReference {
	var returns CloudGcpIntegrationsBigQueryOutputReference
	_jsii_.Get(
		j,
		"bigQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) BigQueryInput() *CloudGcpIntegrationsBigQuery {
	var returns *CloudGcpIntegrationsBigQuery
	_jsii_.Get(
		j,
		"bigQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) BigTable() CloudGcpIntegrationsBigTableOutputReference {
	var returns CloudGcpIntegrationsBigTableOutputReference
	_jsii_.Get(
		j,
		"bigTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) BigTableInput() *CloudGcpIntegrationsBigTable {
	var returns *CloudGcpIntegrationsBigTable
	_jsii_.Get(
		j,
		"bigTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Composer() CloudGcpIntegrationsComposerOutputReference {
	var returns CloudGcpIntegrationsComposerOutputReference
	_jsii_.Get(
		j,
		"composer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) ComposerInput() *CloudGcpIntegrationsComposer {
	var returns *CloudGcpIntegrationsComposer
	_jsii_.Get(
		j,
		"composerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) DataFlow() CloudGcpIntegrationsDataFlowOutputReference {
	var returns CloudGcpIntegrationsDataFlowOutputReference
	_jsii_.Get(
		j,
		"dataFlow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) DataFlowInput() *CloudGcpIntegrationsDataFlow {
	var returns *CloudGcpIntegrationsDataFlow
	_jsii_.Get(
		j,
		"dataFlowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) DataProc() CloudGcpIntegrationsDataProcOutputReference {
	var returns CloudGcpIntegrationsDataProcOutputReference
	_jsii_.Get(
		j,
		"dataProc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) DataProcInput() *CloudGcpIntegrationsDataProc {
	var returns *CloudGcpIntegrationsDataProc
	_jsii_.Get(
		j,
		"dataProcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) DataStore() CloudGcpIntegrationsDataStoreOutputReference {
	var returns CloudGcpIntegrationsDataStoreOutputReference
	_jsii_.Get(
		j,
		"dataStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) DataStoreInput() *CloudGcpIntegrationsDataStore {
	var returns *CloudGcpIntegrationsDataStore
	_jsii_.Get(
		j,
		"dataStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) FireBaseDatabase() CloudGcpIntegrationsFireBaseDatabaseOutputReference {
	var returns CloudGcpIntegrationsFireBaseDatabaseOutputReference
	_jsii_.Get(
		j,
		"fireBaseDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) FireBaseDatabaseInput() *CloudGcpIntegrationsFireBaseDatabase {
	var returns *CloudGcpIntegrationsFireBaseDatabase
	_jsii_.Get(
		j,
		"fireBaseDatabaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) FireBaseHosting() CloudGcpIntegrationsFireBaseHostingOutputReference {
	var returns CloudGcpIntegrationsFireBaseHostingOutputReference
	_jsii_.Get(
		j,
		"fireBaseHosting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) FireBaseHostingInput() *CloudGcpIntegrationsFireBaseHosting {
	var returns *CloudGcpIntegrationsFireBaseHosting
	_jsii_.Get(
		j,
		"fireBaseHostingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) FireBaseStorage() CloudGcpIntegrationsFireBaseStorageOutputReference {
	var returns CloudGcpIntegrationsFireBaseStorageOutputReference
	_jsii_.Get(
		j,
		"fireBaseStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) FireBaseStorageInput() *CloudGcpIntegrationsFireBaseStorage {
	var returns *CloudGcpIntegrationsFireBaseStorage
	_jsii_.Get(
		j,
		"fireBaseStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) FireStore() CloudGcpIntegrationsFireStoreOutputReference {
	var returns CloudGcpIntegrationsFireStoreOutputReference
	_jsii_.Get(
		j,
		"fireStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) FireStoreInput() *CloudGcpIntegrationsFireStore {
	var returns *CloudGcpIntegrationsFireStore
	_jsii_.Get(
		j,
		"fireStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Functions() CloudGcpIntegrationsFunctionsOutputReference {
	var returns CloudGcpIntegrationsFunctionsOutputReference
	_jsii_.Get(
		j,
		"functions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) FunctionsInput() *CloudGcpIntegrationsFunctions {
	var returns *CloudGcpIntegrationsFunctions
	_jsii_.Get(
		j,
		"functionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Interconnect() CloudGcpIntegrationsInterconnectOutputReference {
	var returns CloudGcpIntegrationsInterconnectOutputReference
	_jsii_.Get(
		j,
		"interconnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) InterconnectInput() *CloudGcpIntegrationsInterconnect {
	var returns *CloudGcpIntegrationsInterconnect
	_jsii_.Get(
		j,
		"interconnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Kubernetes() CloudGcpIntegrationsKubernetesOutputReference {
	var returns CloudGcpIntegrationsKubernetesOutputReference
	_jsii_.Get(
		j,
		"kubernetes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) KubernetesInput() *CloudGcpIntegrationsKubernetes {
	var returns *CloudGcpIntegrationsKubernetes
	_jsii_.Get(
		j,
		"kubernetesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) LinkedAccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"linkedAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) LinkedAccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"linkedAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) LoadBalancing() CloudGcpIntegrationsLoadBalancingOutputReference {
	var returns CloudGcpIntegrationsLoadBalancingOutputReference
	_jsii_.Get(
		j,
		"loadBalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) LoadBalancingInput() *CloudGcpIntegrationsLoadBalancing {
	var returns *CloudGcpIntegrationsLoadBalancing
	_jsii_.Get(
		j,
		"loadBalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) MemCache() CloudGcpIntegrationsMemCacheOutputReference {
	var returns CloudGcpIntegrationsMemCacheOutputReference
	_jsii_.Get(
		j,
		"memCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) MemCacheInput() *CloudGcpIntegrationsMemCache {
	var returns *CloudGcpIntegrationsMemCache
	_jsii_.Get(
		j,
		"memCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) PubSub() CloudGcpIntegrationsPubSubOutputReference {
	var returns CloudGcpIntegrationsPubSubOutputReference
	_jsii_.Get(
		j,
		"pubSub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) PubSubInput() *CloudGcpIntegrationsPubSub {
	var returns *CloudGcpIntegrationsPubSub
	_jsii_.Get(
		j,
		"pubSubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Redis() CloudGcpIntegrationsRedisOutputReference {
	var returns CloudGcpIntegrationsRedisOutputReference
	_jsii_.Get(
		j,
		"redis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) RedisInput() *CloudGcpIntegrationsRedis {
	var returns *CloudGcpIntegrationsRedis
	_jsii_.Get(
		j,
		"redisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Router() CloudGcpIntegrationsRouterOutputReference {
	var returns CloudGcpIntegrationsRouterOutputReference
	_jsii_.Get(
		j,
		"router",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) RouterInput() *CloudGcpIntegrationsRouter {
	var returns *CloudGcpIntegrationsRouter
	_jsii_.Get(
		j,
		"routerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Run() CloudGcpIntegrationsRunOutputReference {
	var returns CloudGcpIntegrationsRunOutputReference
	_jsii_.Get(
		j,
		"run",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) RunInput() *CloudGcpIntegrationsRun {
	var returns *CloudGcpIntegrationsRun
	_jsii_.Get(
		j,
		"runInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Spanner() CloudGcpIntegrationsSpannerOutputReference {
	var returns CloudGcpIntegrationsSpannerOutputReference
	_jsii_.Get(
		j,
		"spanner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) SpannerInput() *CloudGcpIntegrationsSpanner {
	var returns *CloudGcpIntegrationsSpanner
	_jsii_.Get(
		j,
		"spannerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Sql() CloudGcpIntegrationsSqlOutputReference {
	var returns CloudGcpIntegrationsSqlOutputReference
	_jsii_.Get(
		j,
		"sql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) SqlInput() *CloudGcpIntegrationsSql {
	var returns *CloudGcpIntegrationsSql
	_jsii_.Get(
		j,
		"sqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) Storage() CloudGcpIntegrationsStorageOutputReference {
	var returns CloudGcpIntegrationsStorageOutputReference
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) StorageInput() *CloudGcpIntegrationsStorage {
	var returns *CloudGcpIntegrationsStorage
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) VirtualMachines() CloudGcpIntegrationsVirtualMachinesOutputReference {
	var returns CloudGcpIntegrationsVirtualMachinesOutputReference
	_jsii_.Get(
		j,
		"virtualMachines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) VirtualMachinesInput() *CloudGcpIntegrationsVirtualMachines {
	var returns *CloudGcpIntegrationsVirtualMachines
	_jsii_.Get(
		j,
		"virtualMachinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) VpcAccess() CloudGcpIntegrationsVpcAccessOutputReference {
	var returns CloudGcpIntegrationsVpcAccessOutputReference
	_jsii_.Get(
		j,
		"vpcAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudGcpIntegrations) VpcAccessInput() *CloudGcpIntegrationsVpcAccess {
	var returns *CloudGcpIntegrationsVpcAccess
	_jsii_.Get(
		j,
		"vpcAccessInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_gcp_integrations newrelic_cloud_gcp_integrations} Resource.
func NewCloudGcpIntegrations(scope constructs.Construct, id *string, config *CloudGcpIntegrationsConfig) CloudGcpIntegrations {
	_init_.Initialize()

	if err := validateNewCloudGcpIntegrationsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudGcpIntegrations{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.CloudGcpIntegrations",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/newrelic/r/cloud_gcp_integrations newrelic_cloud_gcp_integrations} Resource.
func NewCloudGcpIntegrations_Override(c CloudGcpIntegrations, scope constructs.Construct, id *string, config *CloudGcpIntegrationsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.CloudGcpIntegrations",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudGcpIntegrations)SetAccountId(val *float64) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_CloudGcpIntegrations)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudGcpIntegrations)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudGcpIntegrations)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudGcpIntegrations)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudGcpIntegrations)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudGcpIntegrations)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudGcpIntegrations)SetLinkedAccountId(val *float64) {
	if err := j.validateSetLinkedAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkedAccountId",
		val,
	)
}

func (j *jsiiProxy_CloudGcpIntegrations)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudGcpIntegrations)SetProvisioners(val *[]interface{}) {
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
func CloudGcpIntegrations_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudGcpIntegrations_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.CloudGcpIntegrations",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudGcpIntegrations_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.CloudGcpIntegrations",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudGcpIntegrations) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudGcpIntegrations) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudGcpIntegrations) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudGcpIntegrations) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudGcpIntegrations) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudGcpIntegrations) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudGcpIntegrations) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudGcpIntegrations) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudGcpIntegrations) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudGcpIntegrations) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudGcpIntegrations) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutAppEngine(value *CloudGcpIntegrationsAppEngine) {
	if err := c.validatePutAppEngineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAppEngine",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutBigQuery(value *CloudGcpIntegrationsBigQuery) {
	if err := c.validatePutBigQueryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBigQuery",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutBigTable(value *CloudGcpIntegrationsBigTable) {
	if err := c.validatePutBigTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBigTable",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutComposer(value *CloudGcpIntegrationsComposer) {
	if err := c.validatePutComposerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putComposer",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutDataFlow(value *CloudGcpIntegrationsDataFlow) {
	if err := c.validatePutDataFlowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataFlow",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutDataProc(value *CloudGcpIntegrationsDataProc) {
	if err := c.validatePutDataProcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataProc",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutDataStore(value *CloudGcpIntegrationsDataStore) {
	if err := c.validatePutDataStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataStore",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutFireBaseDatabase(value *CloudGcpIntegrationsFireBaseDatabase) {
	if err := c.validatePutFireBaseDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFireBaseDatabase",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutFireBaseHosting(value *CloudGcpIntegrationsFireBaseHosting) {
	if err := c.validatePutFireBaseHostingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFireBaseHosting",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutFireBaseStorage(value *CloudGcpIntegrationsFireBaseStorage) {
	if err := c.validatePutFireBaseStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFireBaseStorage",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutFireStore(value *CloudGcpIntegrationsFireStore) {
	if err := c.validatePutFireStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFireStore",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutFunctions(value *CloudGcpIntegrationsFunctions) {
	if err := c.validatePutFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFunctions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutInterconnect(value *CloudGcpIntegrationsInterconnect) {
	if err := c.validatePutInterconnectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInterconnect",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutKubernetes(value *CloudGcpIntegrationsKubernetes) {
	if err := c.validatePutKubernetesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKubernetes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutLoadBalancing(value *CloudGcpIntegrationsLoadBalancing) {
	if err := c.validatePutLoadBalancingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLoadBalancing",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutMemCache(value *CloudGcpIntegrationsMemCache) {
	if err := c.validatePutMemCacheParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMemCache",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutPubSub(value *CloudGcpIntegrationsPubSub) {
	if err := c.validatePutPubSubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPubSub",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutRedis(value *CloudGcpIntegrationsRedis) {
	if err := c.validatePutRedisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRedis",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutRouter(value *CloudGcpIntegrationsRouter) {
	if err := c.validatePutRouterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRouter",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutRun(value *CloudGcpIntegrationsRun) {
	if err := c.validatePutRunParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRun",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutSpanner(value *CloudGcpIntegrationsSpanner) {
	if err := c.validatePutSpannerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSpanner",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutSql(value *CloudGcpIntegrationsSql) {
	if err := c.validatePutSqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSql",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutStorage(value *CloudGcpIntegrationsStorage) {
	if err := c.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStorage",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutVirtualMachines(value *CloudGcpIntegrationsVirtualMachines) {
	if err := c.validatePutVirtualMachinesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVirtualMachines",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) PutVpcAccess(value *CloudGcpIntegrationsVpcAccess) {
	if err := c.validatePutVpcAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVpcAccess",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetAccountId() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetAppEngine() {
	_jsii_.InvokeVoid(
		c,
		"resetAppEngine",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetBigQuery() {
	_jsii_.InvokeVoid(
		c,
		"resetBigQuery",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetBigTable() {
	_jsii_.InvokeVoid(
		c,
		"resetBigTable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetComposer() {
	_jsii_.InvokeVoid(
		c,
		"resetComposer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetDataFlow() {
	_jsii_.InvokeVoid(
		c,
		"resetDataFlow",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetDataProc() {
	_jsii_.InvokeVoid(
		c,
		"resetDataProc",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetDataStore() {
	_jsii_.InvokeVoid(
		c,
		"resetDataStore",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetFireBaseDatabase() {
	_jsii_.InvokeVoid(
		c,
		"resetFireBaseDatabase",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetFireBaseHosting() {
	_jsii_.InvokeVoid(
		c,
		"resetFireBaseHosting",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetFireBaseStorage() {
	_jsii_.InvokeVoid(
		c,
		"resetFireBaseStorage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetFireStore() {
	_jsii_.InvokeVoid(
		c,
		"resetFireStore",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetFunctions() {
	_jsii_.InvokeVoid(
		c,
		"resetFunctions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetInterconnect() {
	_jsii_.InvokeVoid(
		c,
		"resetInterconnect",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetKubernetes() {
	_jsii_.InvokeVoid(
		c,
		"resetKubernetes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetLoadBalancing() {
	_jsii_.InvokeVoid(
		c,
		"resetLoadBalancing",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetMemCache() {
	_jsii_.InvokeVoid(
		c,
		"resetMemCache",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetPubSub() {
	_jsii_.InvokeVoid(
		c,
		"resetPubSub",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetRedis() {
	_jsii_.InvokeVoid(
		c,
		"resetRedis",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetRouter() {
	_jsii_.InvokeVoid(
		c,
		"resetRouter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetRun() {
	_jsii_.InvokeVoid(
		c,
		"resetRun",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetSpanner() {
	_jsii_.InvokeVoid(
		c,
		"resetSpanner",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetSql() {
	_jsii_.InvokeVoid(
		c,
		"resetSql",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetStorage() {
	_jsii_.InvokeVoid(
		c,
		"resetStorage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetVirtualMachines() {
	_jsii_.InvokeVoid(
		c,
		"resetVirtualMachines",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) ResetVpcAccess() {
	_jsii_.InvokeVoid(
		c,
		"resetVpcAccess",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudGcpIntegrations) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpIntegrations) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpIntegrations) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudGcpIntegrations) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

