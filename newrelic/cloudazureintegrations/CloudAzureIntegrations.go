// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudazureintegrations

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-newrelic-go/newrelic/v13/cloudazureintegrations/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.3/docs/resources/cloud_azure_integrations newrelic_cloud_azure_integrations}.
type CloudAzureIntegrations interface {
	cdktf.TerraformResource
	AccountId() *float64
	SetAccountId(val *float64)
	AccountIdInput() *float64
	ApiManagement() CloudAzureIntegrationsApiManagementOutputReference
	ApiManagementInput() *CloudAzureIntegrationsApiManagement
	AppGateway() CloudAzureIntegrationsAppGatewayOutputReference
	AppGatewayInput() *CloudAzureIntegrationsAppGateway
	AppService() CloudAzureIntegrationsAppServiceOutputReference
	AppServiceInput() *CloudAzureIntegrationsAppService
	AutoDiscovery() CloudAzureIntegrationsAutoDiscoveryOutputReference
	AutoDiscoveryInput() *CloudAzureIntegrationsAutoDiscovery
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Containers() CloudAzureIntegrationsContainersOutputReference
	ContainersInput() *CloudAzureIntegrationsContainers
	CosmosDb() CloudAzureIntegrationsCosmosDbOutputReference
	CosmosDbInput() *CloudAzureIntegrationsCosmosDb
	CostManagement() CloudAzureIntegrationsCostManagementOutputReference
	CostManagementInput() *CloudAzureIntegrationsCostManagement
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DataFactory() CloudAzureIntegrationsDataFactoryOutputReference
	DataFactoryInput() *CloudAzureIntegrationsDataFactory
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EventHub() CloudAzureIntegrationsEventHubOutputReference
	EventHubInput() *CloudAzureIntegrationsEventHub
	ExpressRoute() CloudAzureIntegrationsExpressRouteOutputReference
	ExpressRouteInput() *CloudAzureIntegrationsExpressRoute
	Firewalls() CloudAzureIntegrationsFirewallsOutputReference
	FirewallsInput() *CloudAzureIntegrationsFirewalls
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FrontDoor() CloudAzureIntegrationsFrontDoorOutputReference
	FrontDoorInput() *CloudAzureIntegrationsFrontDoor
	Functions() CloudAzureIntegrationsFunctionsOutputReference
	FunctionsInput() *CloudAzureIntegrationsFunctions
	Id() *string
	SetId(val *string)
	IdInput() *string
	KeyVault() CloudAzureIntegrationsKeyVaultOutputReference
	KeyVaultInput() *CloudAzureIntegrationsKeyVault
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinkedAccountId() *float64
	SetLinkedAccountId(val *float64)
	LinkedAccountIdInput() *float64
	LoadBalancer() CloudAzureIntegrationsLoadBalancerOutputReference
	LoadBalancerInput() *CloudAzureIntegrationsLoadBalancer
	LogicApps() CloudAzureIntegrationsLogicAppsOutputReference
	LogicAppsInput() *CloudAzureIntegrationsLogicApps
	MachineLearning() CloudAzureIntegrationsMachineLearningOutputReference
	MachineLearningInput() *CloudAzureIntegrationsMachineLearning
	MariaDb() CloudAzureIntegrationsMariaDbOutputReference
	MariaDbInput() *CloudAzureIntegrationsMariaDb
	Monitor() CloudAzureIntegrationsMonitorOutputReference
	MonitorInput() *CloudAzureIntegrationsMonitor
	Mysql() CloudAzureIntegrationsMysqlOutputReference
	MysqlFlexible() CloudAzureIntegrationsMysqlFlexibleOutputReference
	MysqlFlexibleInput() *CloudAzureIntegrationsMysqlFlexible
	MysqlInput() *CloudAzureIntegrationsMysql
	// The tree node.
	Node() constructs.Node
	Postgresql() CloudAzureIntegrationsPostgresqlOutputReference
	PostgresqlFlexible() CloudAzureIntegrationsPostgresqlFlexibleOutputReference
	PostgresqlFlexibleInput() *CloudAzureIntegrationsPostgresqlFlexible
	PostgresqlInput() *CloudAzureIntegrationsPostgresql
	PowerBiDedicated() CloudAzureIntegrationsPowerBiDedicatedOutputReference
	PowerBiDedicatedInput() *CloudAzureIntegrationsPowerBiDedicated
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
	RedisCache() CloudAzureIntegrationsRedisCacheOutputReference
	RedisCacheInput() *CloudAzureIntegrationsRedisCache
	ServiceBus() CloudAzureIntegrationsServiceBusOutputReference
	ServiceBusInput() *CloudAzureIntegrationsServiceBus
	Sql() CloudAzureIntegrationsSqlOutputReference
	SqlInput() *CloudAzureIntegrationsSql
	SqlManaged() CloudAzureIntegrationsSqlManagedOutputReference
	SqlManagedInput() *CloudAzureIntegrationsSqlManaged
	Storage() CloudAzureIntegrationsStorageOutputReference
	StorageInput() *CloudAzureIntegrationsStorage
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VirtualMachine() CloudAzureIntegrationsVirtualMachineOutputReference
	VirtualMachineInput() *CloudAzureIntegrationsVirtualMachine
	VirtualNetworks() CloudAzureIntegrationsVirtualNetworksOutputReference
	VirtualNetworksInput() *CloudAzureIntegrationsVirtualNetworks
	Vms() CloudAzureIntegrationsVmsOutputReference
	VmsInput() *CloudAzureIntegrationsVms
	VpnGateway() CloudAzureIntegrationsVpnGatewayOutputReference
	VpnGatewayInput() *CloudAzureIntegrationsVpnGateway
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
	PutApiManagement(value *CloudAzureIntegrationsApiManagement)
	PutAppGateway(value *CloudAzureIntegrationsAppGateway)
	PutAppService(value *CloudAzureIntegrationsAppService)
	PutAutoDiscovery(value *CloudAzureIntegrationsAutoDiscovery)
	PutContainers(value *CloudAzureIntegrationsContainers)
	PutCosmosDb(value *CloudAzureIntegrationsCosmosDb)
	PutCostManagement(value *CloudAzureIntegrationsCostManagement)
	PutDataFactory(value *CloudAzureIntegrationsDataFactory)
	PutEventHub(value *CloudAzureIntegrationsEventHub)
	PutExpressRoute(value *CloudAzureIntegrationsExpressRoute)
	PutFirewalls(value *CloudAzureIntegrationsFirewalls)
	PutFrontDoor(value *CloudAzureIntegrationsFrontDoor)
	PutFunctions(value *CloudAzureIntegrationsFunctions)
	PutKeyVault(value *CloudAzureIntegrationsKeyVault)
	PutLoadBalancer(value *CloudAzureIntegrationsLoadBalancer)
	PutLogicApps(value *CloudAzureIntegrationsLogicApps)
	PutMachineLearning(value *CloudAzureIntegrationsMachineLearning)
	PutMariaDb(value *CloudAzureIntegrationsMariaDb)
	PutMonitor(value *CloudAzureIntegrationsMonitor)
	PutMysql(value *CloudAzureIntegrationsMysql)
	PutMysqlFlexible(value *CloudAzureIntegrationsMysqlFlexible)
	PutPostgresql(value *CloudAzureIntegrationsPostgresql)
	PutPostgresqlFlexible(value *CloudAzureIntegrationsPostgresqlFlexible)
	PutPowerBiDedicated(value *CloudAzureIntegrationsPowerBiDedicated)
	PutRedisCache(value *CloudAzureIntegrationsRedisCache)
	PutServiceBus(value *CloudAzureIntegrationsServiceBus)
	PutSql(value *CloudAzureIntegrationsSql)
	PutSqlManaged(value *CloudAzureIntegrationsSqlManaged)
	PutStorage(value *CloudAzureIntegrationsStorage)
	PutVirtualMachine(value *CloudAzureIntegrationsVirtualMachine)
	PutVirtualNetworks(value *CloudAzureIntegrationsVirtualNetworks)
	PutVms(value *CloudAzureIntegrationsVms)
	PutVpnGateway(value *CloudAzureIntegrationsVpnGateway)
	ResetAccountId()
	ResetApiManagement()
	ResetAppGateway()
	ResetAppService()
	ResetAutoDiscovery()
	ResetContainers()
	ResetCosmosDb()
	ResetCostManagement()
	ResetDataFactory()
	ResetEventHub()
	ResetExpressRoute()
	ResetFirewalls()
	ResetFrontDoor()
	ResetFunctions()
	ResetId()
	ResetKeyVault()
	ResetLoadBalancer()
	ResetLogicApps()
	ResetMachineLearning()
	ResetMariaDb()
	ResetMonitor()
	ResetMysql()
	ResetMysqlFlexible()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPostgresql()
	ResetPostgresqlFlexible()
	ResetPowerBiDedicated()
	ResetRedisCache()
	ResetServiceBus()
	ResetSql()
	ResetSqlManaged()
	ResetStorage()
	ResetVirtualMachine()
	ResetVirtualNetworks()
	ResetVms()
	ResetVpnGateway()
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

// The jsii proxy struct for CloudAzureIntegrations
type jsiiProxy_CloudAzureIntegrations struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudAzureIntegrations) AccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) AccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) ApiManagement() CloudAzureIntegrationsApiManagementOutputReference {
	var returns CloudAzureIntegrationsApiManagementOutputReference
	_jsii_.Get(
		j,
		"apiManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) ApiManagementInput() *CloudAzureIntegrationsApiManagement {
	var returns *CloudAzureIntegrationsApiManagement
	_jsii_.Get(
		j,
		"apiManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) AppGateway() CloudAzureIntegrationsAppGatewayOutputReference {
	var returns CloudAzureIntegrationsAppGatewayOutputReference
	_jsii_.Get(
		j,
		"appGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) AppGatewayInput() *CloudAzureIntegrationsAppGateway {
	var returns *CloudAzureIntegrationsAppGateway
	_jsii_.Get(
		j,
		"appGatewayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) AppService() CloudAzureIntegrationsAppServiceOutputReference {
	var returns CloudAzureIntegrationsAppServiceOutputReference
	_jsii_.Get(
		j,
		"appService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) AppServiceInput() *CloudAzureIntegrationsAppService {
	var returns *CloudAzureIntegrationsAppService
	_jsii_.Get(
		j,
		"appServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) AutoDiscovery() CloudAzureIntegrationsAutoDiscoveryOutputReference {
	var returns CloudAzureIntegrationsAutoDiscoveryOutputReference
	_jsii_.Get(
		j,
		"autoDiscovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) AutoDiscoveryInput() *CloudAzureIntegrationsAutoDiscovery {
	var returns *CloudAzureIntegrationsAutoDiscovery
	_jsii_.Get(
		j,
		"autoDiscoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Containers() CloudAzureIntegrationsContainersOutputReference {
	var returns CloudAzureIntegrationsContainersOutputReference
	_jsii_.Get(
		j,
		"containers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) ContainersInput() *CloudAzureIntegrationsContainers {
	var returns *CloudAzureIntegrationsContainers
	_jsii_.Get(
		j,
		"containersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) CosmosDb() CloudAzureIntegrationsCosmosDbOutputReference {
	var returns CloudAzureIntegrationsCosmosDbOutputReference
	_jsii_.Get(
		j,
		"cosmosDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) CosmosDbInput() *CloudAzureIntegrationsCosmosDb {
	var returns *CloudAzureIntegrationsCosmosDb
	_jsii_.Get(
		j,
		"cosmosDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) CostManagement() CloudAzureIntegrationsCostManagementOutputReference {
	var returns CloudAzureIntegrationsCostManagementOutputReference
	_jsii_.Get(
		j,
		"costManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) CostManagementInput() *CloudAzureIntegrationsCostManagement {
	var returns *CloudAzureIntegrationsCostManagement
	_jsii_.Get(
		j,
		"costManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) DataFactory() CloudAzureIntegrationsDataFactoryOutputReference {
	var returns CloudAzureIntegrationsDataFactoryOutputReference
	_jsii_.Get(
		j,
		"dataFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) DataFactoryInput() *CloudAzureIntegrationsDataFactory {
	var returns *CloudAzureIntegrationsDataFactory
	_jsii_.Get(
		j,
		"dataFactoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) EventHub() CloudAzureIntegrationsEventHubOutputReference {
	var returns CloudAzureIntegrationsEventHubOutputReference
	_jsii_.Get(
		j,
		"eventHub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) EventHubInput() *CloudAzureIntegrationsEventHub {
	var returns *CloudAzureIntegrationsEventHub
	_jsii_.Get(
		j,
		"eventHubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) ExpressRoute() CloudAzureIntegrationsExpressRouteOutputReference {
	var returns CloudAzureIntegrationsExpressRouteOutputReference
	_jsii_.Get(
		j,
		"expressRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) ExpressRouteInput() *CloudAzureIntegrationsExpressRoute {
	var returns *CloudAzureIntegrationsExpressRoute
	_jsii_.Get(
		j,
		"expressRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Firewalls() CloudAzureIntegrationsFirewallsOutputReference {
	var returns CloudAzureIntegrationsFirewallsOutputReference
	_jsii_.Get(
		j,
		"firewalls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) FirewallsInput() *CloudAzureIntegrationsFirewalls {
	var returns *CloudAzureIntegrationsFirewalls
	_jsii_.Get(
		j,
		"firewallsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) FrontDoor() CloudAzureIntegrationsFrontDoorOutputReference {
	var returns CloudAzureIntegrationsFrontDoorOutputReference
	_jsii_.Get(
		j,
		"frontDoor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) FrontDoorInput() *CloudAzureIntegrationsFrontDoor {
	var returns *CloudAzureIntegrationsFrontDoor
	_jsii_.Get(
		j,
		"frontDoorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Functions() CloudAzureIntegrationsFunctionsOutputReference {
	var returns CloudAzureIntegrationsFunctionsOutputReference
	_jsii_.Get(
		j,
		"functions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) FunctionsInput() *CloudAzureIntegrationsFunctions {
	var returns *CloudAzureIntegrationsFunctions
	_jsii_.Get(
		j,
		"functionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) KeyVault() CloudAzureIntegrationsKeyVaultOutputReference {
	var returns CloudAzureIntegrationsKeyVaultOutputReference
	_jsii_.Get(
		j,
		"keyVault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) KeyVaultInput() *CloudAzureIntegrationsKeyVault {
	var returns *CloudAzureIntegrationsKeyVault
	_jsii_.Get(
		j,
		"keyVaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) LinkedAccountId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"linkedAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) LinkedAccountIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"linkedAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) LoadBalancer() CloudAzureIntegrationsLoadBalancerOutputReference {
	var returns CloudAzureIntegrationsLoadBalancerOutputReference
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) LoadBalancerInput() *CloudAzureIntegrationsLoadBalancer {
	var returns *CloudAzureIntegrationsLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) LogicApps() CloudAzureIntegrationsLogicAppsOutputReference {
	var returns CloudAzureIntegrationsLogicAppsOutputReference
	_jsii_.Get(
		j,
		"logicApps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) LogicAppsInput() *CloudAzureIntegrationsLogicApps {
	var returns *CloudAzureIntegrationsLogicApps
	_jsii_.Get(
		j,
		"logicAppsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) MachineLearning() CloudAzureIntegrationsMachineLearningOutputReference {
	var returns CloudAzureIntegrationsMachineLearningOutputReference
	_jsii_.Get(
		j,
		"machineLearning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) MachineLearningInput() *CloudAzureIntegrationsMachineLearning {
	var returns *CloudAzureIntegrationsMachineLearning
	_jsii_.Get(
		j,
		"machineLearningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) MariaDb() CloudAzureIntegrationsMariaDbOutputReference {
	var returns CloudAzureIntegrationsMariaDbOutputReference
	_jsii_.Get(
		j,
		"mariaDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) MariaDbInput() *CloudAzureIntegrationsMariaDb {
	var returns *CloudAzureIntegrationsMariaDb
	_jsii_.Get(
		j,
		"mariaDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Monitor() CloudAzureIntegrationsMonitorOutputReference {
	var returns CloudAzureIntegrationsMonitorOutputReference
	_jsii_.Get(
		j,
		"monitor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) MonitorInput() *CloudAzureIntegrationsMonitor {
	var returns *CloudAzureIntegrationsMonitor
	_jsii_.Get(
		j,
		"monitorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Mysql() CloudAzureIntegrationsMysqlOutputReference {
	var returns CloudAzureIntegrationsMysqlOutputReference
	_jsii_.Get(
		j,
		"mysql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) MysqlFlexible() CloudAzureIntegrationsMysqlFlexibleOutputReference {
	var returns CloudAzureIntegrationsMysqlFlexibleOutputReference
	_jsii_.Get(
		j,
		"mysqlFlexible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) MysqlFlexibleInput() *CloudAzureIntegrationsMysqlFlexible {
	var returns *CloudAzureIntegrationsMysqlFlexible
	_jsii_.Get(
		j,
		"mysqlFlexibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) MysqlInput() *CloudAzureIntegrationsMysql {
	var returns *CloudAzureIntegrationsMysql
	_jsii_.Get(
		j,
		"mysqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Postgresql() CloudAzureIntegrationsPostgresqlOutputReference {
	var returns CloudAzureIntegrationsPostgresqlOutputReference
	_jsii_.Get(
		j,
		"postgresql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) PostgresqlFlexible() CloudAzureIntegrationsPostgresqlFlexibleOutputReference {
	var returns CloudAzureIntegrationsPostgresqlFlexibleOutputReference
	_jsii_.Get(
		j,
		"postgresqlFlexible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) PostgresqlFlexibleInput() *CloudAzureIntegrationsPostgresqlFlexible {
	var returns *CloudAzureIntegrationsPostgresqlFlexible
	_jsii_.Get(
		j,
		"postgresqlFlexibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) PostgresqlInput() *CloudAzureIntegrationsPostgresql {
	var returns *CloudAzureIntegrationsPostgresql
	_jsii_.Get(
		j,
		"postgresqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) PowerBiDedicated() CloudAzureIntegrationsPowerBiDedicatedOutputReference {
	var returns CloudAzureIntegrationsPowerBiDedicatedOutputReference
	_jsii_.Get(
		j,
		"powerBiDedicated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) PowerBiDedicatedInput() *CloudAzureIntegrationsPowerBiDedicated {
	var returns *CloudAzureIntegrationsPowerBiDedicated
	_jsii_.Get(
		j,
		"powerBiDedicatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) RedisCache() CloudAzureIntegrationsRedisCacheOutputReference {
	var returns CloudAzureIntegrationsRedisCacheOutputReference
	_jsii_.Get(
		j,
		"redisCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) RedisCacheInput() *CloudAzureIntegrationsRedisCache {
	var returns *CloudAzureIntegrationsRedisCache
	_jsii_.Get(
		j,
		"redisCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) ServiceBus() CloudAzureIntegrationsServiceBusOutputReference {
	var returns CloudAzureIntegrationsServiceBusOutputReference
	_jsii_.Get(
		j,
		"serviceBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) ServiceBusInput() *CloudAzureIntegrationsServiceBus {
	var returns *CloudAzureIntegrationsServiceBus
	_jsii_.Get(
		j,
		"serviceBusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Sql() CloudAzureIntegrationsSqlOutputReference {
	var returns CloudAzureIntegrationsSqlOutputReference
	_jsii_.Get(
		j,
		"sql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) SqlInput() *CloudAzureIntegrationsSql {
	var returns *CloudAzureIntegrationsSql
	_jsii_.Get(
		j,
		"sqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) SqlManaged() CloudAzureIntegrationsSqlManagedOutputReference {
	var returns CloudAzureIntegrationsSqlManagedOutputReference
	_jsii_.Get(
		j,
		"sqlManaged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) SqlManagedInput() *CloudAzureIntegrationsSqlManaged {
	var returns *CloudAzureIntegrationsSqlManaged
	_jsii_.Get(
		j,
		"sqlManagedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Storage() CloudAzureIntegrationsStorageOutputReference {
	var returns CloudAzureIntegrationsStorageOutputReference
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) StorageInput() *CloudAzureIntegrationsStorage {
	var returns *CloudAzureIntegrationsStorage
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) VirtualMachine() CloudAzureIntegrationsVirtualMachineOutputReference {
	var returns CloudAzureIntegrationsVirtualMachineOutputReference
	_jsii_.Get(
		j,
		"virtualMachine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) VirtualMachineInput() *CloudAzureIntegrationsVirtualMachine {
	var returns *CloudAzureIntegrationsVirtualMachine
	_jsii_.Get(
		j,
		"virtualMachineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) VirtualNetworks() CloudAzureIntegrationsVirtualNetworksOutputReference {
	var returns CloudAzureIntegrationsVirtualNetworksOutputReference
	_jsii_.Get(
		j,
		"virtualNetworks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) VirtualNetworksInput() *CloudAzureIntegrationsVirtualNetworks {
	var returns *CloudAzureIntegrationsVirtualNetworks
	_jsii_.Get(
		j,
		"virtualNetworksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) Vms() CloudAzureIntegrationsVmsOutputReference {
	var returns CloudAzureIntegrationsVmsOutputReference
	_jsii_.Get(
		j,
		"vms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) VmsInput() *CloudAzureIntegrationsVms {
	var returns *CloudAzureIntegrationsVms
	_jsii_.Get(
		j,
		"vmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) VpnGateway() CloudAzureIntegrationsVpnGatewayOutputReference {
	var returns CloudAzureIntegrationsVpnGatewayOutputReference
	_jsii_.Get(
		j,
		"vpnGateway",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudAzureIntegrations) VpnGatewayInput() *CloudAzureIntegrationsVpnGateway {
	var returns *CloudAzureIntegrationsVpnGateway
	_jsii_.Get(
		j,
		"vpnGatewayInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.3/docs/resources/cloud_azure_integrations newrelic_cloud_azure_integrations} Resource.
func NewCloudAzureIntegrations(scope constructs.Construct, id *string, config *CloudAzureIntegrationsConfig) CloudAzureIntegrations {
	_init_.Initialize()

	if err := validateNewCloudAzureIntegrationsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudAzureIntegrations{}

	_jsii_.Create(
		"@cdktf/provider-newrelic.cloudAzureIntegrations.CloudAzureIntegrations",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/newrelic/newrelic/3.76.3/docs/resources/cloud_azure_integrations newrelic_cloud_azure_integrations} Resource.
func NewCloudAzureIntegrations_Override(c CloudAzureIntegrations, scope constructs.Construct, id *string, config *CloudAzureIntegrationsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-newrelic.cloudAzureIntegrations.CloudAzureIntegrations",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudAzureIntegrations)SetAccountId(val *float64) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_CloudAzureIntegrations)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudAzureIntegrations)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudAzureIntegrations)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudAzureIntegrations)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudAzureIntegrations)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudAzureIntegrations)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudAzureIntegrations)SetLinkedAccountId(val *float64) {
	if err := j.validateSetLinkedAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkedAccountId",
		val,
	)
}

func (j *jsiiProxy_CloudAzureIntegrations)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudAzureIntegrations)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a CloudAzureIntegrations resource upon running "cdktf plan <stack-name>".
func CloudAzureIntegrations_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudAzureIntegrations_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudAzureIntegrations.CloudAzureIntegrations",
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
func CloudAzureIntegrations_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAzureIntegrations_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudAzureIntegrations.CloudAzureIntegrations",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudAzureIntegrations_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAzureIntegrations_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudAzureIntegrations.CloudAzureIntegrations",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudAzureIntegrations_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudAzureIntegrations_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-newrelic.cloudAzureIntegrations.CloudAzureIntegrations",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudAzureIntegrations_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-newrelic.cloudAzureIntegrations.CloudAzureIntegrations",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudAzureIntegrations) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudAzureIntegrations) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudAzureIntegrations) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudAzureIntegrations) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudAzureIntegrations) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudAzureIntegrations) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudAzureIntegrations) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudAzureIntegrations) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudAzureIntegrations) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudAzureIntegrations) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrations) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudAzureIntegrations) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutApiManagement(value *CloudAzureIntegrationsApiManagement) {
	if err := c.validatePutApiManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putApiManagement",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutAppGateway(value *CloudAzureIntegrationsAppGateway) {
	if err := c.validatePutAppGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAppGateway",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutAppService(value *CloudAzureIntegrationsAppService) {
	if err := c.validatePutAppServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAppService",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutAutoDiscovery(value *CloudAzureIntegrationsAutoDiscovery) {
	if err := c.validatePutAutoDiscoveryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutoDiscovery",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutContainers(value *CloudAzureIntegrationsContainers) {
	if err := c.validatePutContainersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putContainers",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutCosmosDb(value *CloudAzureIntegrationsCosmosDb) {
	if err := c.validatePutCosmosDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCosmosDb",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutCostManagement(value *CloudAzureIntegrationsCostManagement) {
	if err := c.validatePutCostManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCostManagement",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutDataFactory(value *CloudAzureIntegrationsDataFactory) {
	if err := c.validatePutDataFactoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataFactory",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutEventHub(value *CloudAzureIntegrationsEventHub) {
	if err := c.validatePutEventHubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEventHub",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutExpressRoute(value *CloudAzureIntegrationsExpressRoute) {
	if err := c.validatePutExpressRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExpressRoute",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutFirewalls(value *CloudAzureIntegrationsFirewalls) {
	if err := c.validatePutFirewallsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFirewalls",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutFrontDoor(value *CloudAzureIntegrationsFrontDoor) {
	if err := c.validatePutFrontDoorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFrontDoor",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutFunctions(value *CloudAzureIntegrationsFunctions) {
	if err := c.validatePutFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFunctions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutKeyVault(value *CloudAzureIntegrationsKeyVault) {
	if err := c.validatePutKeyVaultParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKeyVault",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutLoadBalancer(value *CloudAzureIntegrationsLoadBalancer) {
	if err := c.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutLogicApps(value *CloudAzureIntegrationsLogicApps) {
	if err := c.validatePutLogicAppsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLogicApps",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutMachineLearning(value *CloudAzureIntegrationsMachineLearning) {
	if err := c.validatePutMachineLearningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMachineLearning",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutMariaDb(value *CloudAzureIntegrationsMariaDb) {
	if err := c.validatePutMariaDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMariaDb",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutMonitor(value *CloudAzureIntegrationsMonitor) {
	if err := c.validatePutMonitorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMonitor",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutMysql(value *CloudAzureIntegrationsMysql) {
	if err := c.validatePutMysqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMysql",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutMysqlFlexible(value *CloudAzureIntegrationsMysqlFlexible) {
	if err := c.validatePutMysqlFlexibleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMysqlFlexible",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutPostgresql(value *CloudAzureIntegrationsPostgresql) {
	if err := c.validatePutPostgresqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPostgresql",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutPostgresqlFlexible(value *CloudAzureIntegrationsPostgresqlFlexible) {
	if err := c.validatePutPostgresqlFlexibleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPostgresqlFlexible",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutPowerBiDedicated(value *CloudAzureIntegrationsPowerBiDedicated) {
	if err := c.validatePutPowerBiDedicatedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPowerBiDedicated",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutRedisCache(value *CloudAzureIntegrationsRedisCache) {
	if err := c.validatePutRedisCacheParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRedisCache",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutServiceBus(value *CloudAzureIntegrationsServiceBus) {
	if err := c.validatePutServiceBusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServiceBus",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutSql(value *CloudAzureIntegrationsSql) {
	if err := c.validatePutSqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSql",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutSqlManaged(value *CloudAzureIntegrationsSqlManaged) {
	if err := c.validatePutSqlManagedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSqlManaged",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutStorage(value *CloudAzureIntegrationsStorage) {
	if err := c.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStorage",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutVirtualMachine(value *CloudAzureIntegrationsVirtualMachine) {
	if err := c.validatePutVirtualMachineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVirtualMachine",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutVirtualNetworks(value *CloudAzureIntegrationsVirtualNetworks) {
	if err := c.validatePutVirtualNetworksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVirtualNetworks",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutVms(value *CloudAzureIntegrationsVms) {
	if err := c.validatePutVmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVms",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) PutVpnGateway(value *CloudAzureIntegrationsVpnGateway) {
	if err := c.validatePutVpnGatewayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVpnGateway",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetAccountId() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetApiManagement() {
	_jsii_.InvokeVoid(
		c,
		"resetApiManagement",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetAppGateway() {
	_jsii_.InvokeVoid(
		c,
		"resetAppGateway",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetAppService() {
	_jsii_.InvokeVoid(
		c,
		"resetAppService",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetAutoDiscovery() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoDiscovery",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetContainers() {
	_jsii_.InvokeVoid(
		c,
		"resetContainers",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetCosmosDb() {
	_jsii_.InvokeVoid(
		c,
		"resetCosmosDb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetCostManagement() {
	_jsii_.InvokeVoid(
		c,
		"resetCostManagement",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetDataFactory() {
	_jsii_.InvokeVoid(
		c,
		"resetDataFactory",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetEventHub() {
	_jsii_.InvokeVoid(
		c,
		"resetEventHub",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetExpressRoute() {
	_jsii_.InvokeVoid(
		c,
		"resetExpressRoute",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetFirewalls() {
	_jsii_.InvokeVoid(
		c,
		"resetFirewalls",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetFrontDoor() {
	_jsii_.InvokeVoid(
		c,
		"resetFrontDoor",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetFunctions() {
	_jsii_.InvokeVoid(
		c,
		"resetFunctions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetKeyVault() {
	_jsii_.InvokeVoid(
		c,
		"resetKeyVault",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		c,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetLogicApps() {
	_jsii_.InvokeVoid(
		c,
		"resetLogicApps",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetMachineLearning() {
	_jsii_.InvokeVoid(
		c,
		"resetMachineLearning",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetMariaDb() {
	_jsii_.InvokeVoid(
		c,
		"resetMariaDb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetMonitor() {
	_jsii_.InvokeVoid(
		c,
		"resetMonitor",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetMysql() {
	_jsii_.InvokeVoid(
		c,
		"resetMysql",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetMysqlFlexible() {
	_jsii_.InvokeVoid(
		c,
		"resetMysqlFlexible",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetPostgresql() {
	_jsii_.InvokeVoid(
		c,
		"resetPostgresql",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetPostgresqlFlexible() {
	_jsii_.InvokeVoid(
		c,
		"resetPostgresqlFlexible",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetPowerBiDedicated() {
	_jsii_.InvokeVoid(
		c,
		"resetPowerBiDedicated",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetRedisCache() {
	_jsii_.InvokeVoid(
		c,
		"resetRedisCache",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetServiceBus() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceBus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetSql() {
	_jsii_.InvokeVoid(
		c,
		"resetSql",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetSqlManaged() {
	_jsii_.InvokeVoid(
		c,
		"resetSqlManaged",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetStorage() {
	_jsii_.InvokeVoid(
		c,
		"resetStorage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetVirtualMachine() {
	_jsii_.InvokeVoid(
		c,
		"resetVirtualMachine",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetVirtualNetworks() {
	_jsii_.InvokeVoid(
		c,
		"resetVirtualNetworks",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetVms() {
	_jsii_.InvokeVoid(
		c,
		"resetVms",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) ResetVpnGateway() {
	_jsii_.InvokeVoid(
		c,
		"resetVpnGateway",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudAzureIntegrations) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrations) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrations) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrations) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrations) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudAzureIntegrations) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

