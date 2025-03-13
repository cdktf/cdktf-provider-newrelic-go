// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudazureintegrations

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudAzureIntegrationsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the linked Azure account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#linked_account_id CloudAzureIntegrations#linked_account_id}
	LinkedAccountId *float64 `field:"required" json:"linkedAccountId" yaml:"linkedAccountId"`
	// The ID of the account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#account_id CloudAzureIntegrations#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// api_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#api_management CloudAzureIntegrations#api_management}
	ApiManagement *CloudAzureIntegrationsApiManagement `field:"optional" json:"apiManagement" yaml:"apiManagement"`
	// app_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#app_gateway CloudAzureIntegrations#app_gateway}
	AppGateway *CloudAzureIntegrationsAppGateway `field:"optional" json:"appGateway" yaml:"appGateway"`
	// app_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#app_service CloudAzureIntegrations#app_service}
	AppService *CloudAzureIntegrationsAppService `field:"optional" json:"appService" yaml:"appService"`
	// containers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#containers CloudAzureIntegrations#containers}
	Containers *CloudAzureIntegrationsContainers `field:"optional" json:"containers" yaml:"containers"`
	// cosmos_db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#cosmos_db CloudAzureIntegrations#cosmos_db}
	CosmosDb *CloudAzureIntegrationsCosmosDb `field:"optional" json:"cosmosDb" yaml:"cosmosDb"`
	// cost_management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#cost_management CloudAzureIntegrations#cost_management}
	CostManagement *CloudAzureIntegrationsCostManagement `field:"optional" json:"costManagement" yaml:"costManagement"`
	// data_factory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#data_factory CloudAzureIntegrations#data_factory}
	DataFactory *CloudAzureIntegrationsDataFactory `field:"optional" json:"dataFactory" yaml:"dataFactory"`
	// event_hub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#event_hub CloudAzureIntegrations#event_hub}
	EventHub *CloudAzureIntegrationsEventHub `field:"optional" json:"eventHub" yaml:"eventHub"`
	// express_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#express_route CloudAzureIntegrations#express_route}
	ExpressRoute *CloudAzureIntegrationsExpressRoute `field:"optional" json:"expressRoute" yaml:"expressRoute"`
	// firewalls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#firewalls CloudAzureIntegrations#firewalls}
	Firewalls *CloudAzureIntegrationsFirewalls `field:"optional" json:"firewalls" yaml:"firewalls"`
	// front_door block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#front_door CloudAzureIntegrations#front_door}
	FrontDoor *CloudAzureIntegrationsFrontDoor `field:"optional" json:"frontDoor" yaml:"frontDoor"`
	// functions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#functions CloudAzureIntegrations#functions}
	Functions *CloudAzureIntegrationsFunctions `field:"optional" json:"functions" yaml:"functions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#id CloudAzureIntegrations#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// key_vault block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#key_vault CloudAzureIntegrations#key_vault}
	KeyVault *CloudAzureIntegrationsKeyVault `field:"optional" json:"keyVault" yaml:"keyVault"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#load_balancer CloudAzureIntegrations#load_balancer}
	LoadBalancer *CloudAzureIntegrationsLoadBalancer `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// logic_apps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#logic_apps CloudAzureIntegrations#logic_apps}
	LogicApps *CloudAzureIntegrationsLogicApps `field:"optional" json:"logicApps" yaml:"logicApps"`
	// machine_learning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#machine_learning CloudAzureIntegrations#machine_learning}
	MachineLearning *CloudAzureIntegrationsMachineLearning `field:"optional" json:"machineLearning" yaml:"machineLearning"`
	// maria_db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#maria_db CloudAzureIntegrations#maria_db}
	MariaDb *CloudAzureIntegrationsMariaDb `field:"optional" json:"mariaDb" yaml:"mariaDb"`
	// monitor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#monitor CloudAzureIntegrations#monitor}
	Monitor *CloudAzureIntegrationsMonitor `field:"optional" json:"monitor" yaml:"monitor"`
	// mysql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#mysql CloudAzureIntegrations#mysql}
	Mysql *CloudAzureIntegrationsMysql `field:"optional" json:"mysql" yaml:"mysql"`
	// mysql_flexible block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#mysql_flexible CloudAzureIntegrations#mysql_flexible}
	MysqlFlexible *CloudAzureIntegrationsMysqlFlexible `field:"optional" json:"mysqlFlexible" yaml:"mysqlFlexible"`
	// postgresql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#postgresql CloudAzureIntegrations#postgresql}
	Postgresql *CloudAzureIntegrationsPostgresql `field:"optional" json:"postgresql" yaml:"postgresql"`
	// postgresql_flexible block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#postgresql_flexible CloudAzureIntegrations#postgresql_flexible}
	PostgresqlFlexible *CloudAzureIntegrationsPostgresqlFlexible `field:"optional" json:"postgresqlFlexible" yaml:"postgresqlFlexible"`
	// power_bi_dedicated block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#power_bi_dedicated CloudAzureIntegrations#power_bi_dedicated}
	PowerBiDedicated *CloudAzureIntegrationsPowerBiDedicated `field:"optional" json:"powerBiDedicated" yaml:"powerBiDedicated"`
	// redis_cache block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#redis_cache CloudAzureIntegrations#redis_cache}
	RedisCache *CloudAzureIntegrationsRedisCache `field:"optional" json:"redisCache" yaml:"redisCache"`
	// service_bus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#service_bus CloudAzureIntegrations#service_bus}
	ServiceBus *CloudAzureIntegrationsServiceBus `field:"optional" json:"serviceBus" yaml:"serviceBus"`
	// sql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#sql CloudAzureIntegrations#sql}
	Sql *CloudAzureIntegrationsSql `field:"optional" json:"sql" yaml:"sql"`
	// sql_managed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#sql_managed CloudAzureIntegrations#sql_managed}
	SqlManaged *CloudAzureIntegrationsSqlManaged `field:"optional" json:"sqlManaged" yaml:"sqlManaged"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#storage CloudAzureIntegrations#storage}
	Storage *CloudAzureIntegrationsStorage `field:"optional" json:"storage" yaml:"storage"`
	// virtual_machine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#virtual_machine CloudAzureIntegrations#virtual_machine}
	VirtualMachine *CloudAzureIntegrationsVirtualMachine `field:"optional" json:"virtualMachine" yaml:"virtualMachine"`
	// virtual_networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#virtual_networks CloudAzureIntegrations#virtual_networks}
	VirtualNetworks *CloudAzureIntegrationsVirtualNetworks `field:"optional" json:"virtualNetworks" yaml:"virtualNetworks"`
	// vms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#vms CloudAzureIntegrations#vms}
	Vms *CloudAzureIntegrationsVms `field:"optional" json:"vms" yaml:"vms"`
	// vpn_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.58.1/docs/resources/cloud_azure_integrations#vpn_gateway CloudAzureIntegrations#vpn_gateway}
	VpnGateway *CloudAzureIntegrationsVpnGateway `field:"optional" json:"vpnGateway" yaml:"vpnGateway"`
}

