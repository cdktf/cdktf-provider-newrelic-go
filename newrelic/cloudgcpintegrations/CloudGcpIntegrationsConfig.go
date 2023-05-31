package cloudgcpintegrations

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudGcpIntegrationsConfig struct {
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
	// Id of the linked gcp account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#linked_account_id CloudGcpIntegrations#linked_account_id}
	LinkedAccountId *float64 `field:"required" json:"linkedAccountId" yaml:"linkedAccountId"`
	// ID of the newrelic account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#account_id CloudGcpIntegrations#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// alloy_db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#alloy_db CloudGcpIntegrations#alloy_db}
	AlloyDb *CloudGcpIntegrationsAlloyDb `field:"optional" json:"alloyDb" yaml:"alloyDb"`
	// app_engine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#app_engine CloudGcpIntegrations#app_engine}
	AppEngine *CloudGcpIntegrationsAppEngine `field:"optional" json:"appEngine" yaml:"appEngine"`
	// big_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#big_query CloudGcpIntegrations#big_query}
	BigQuery *CloudGcpIntegrationsBigQuery `field:"optional" json:"bigQuery" yaml:"bigQuery"`
	// big_table block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#big_table CloudGcpIntegrations#big_table}
	BigTable *CloudGcpIntegrationsBigTable `field:"optional" json:"bigTable" yaml:"bigTable"`
	// composer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#composer CloudGcpIntegrations#composer}
	Composer *CloudGcpIntegrationsComposer `field:"optional" json:"composer" yaml:"composer"`
	// data_flow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#data_flow CloudGcpIntegrations#data_flow}
	DataFlow *CloudGcpIntegrationsDataFlow `field:"optional" json:"dataFlow" yaml:"dataFlow"`
	// data_proc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#data_proc CloudGcpIntegrations#data_proc}
	DataProc *CloudGcpIntegrationsDataProc `field:"optional" json:"dataProc" yaml:"dataProc"`
	// data_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#data_store CloudGcpIntegrations#data_store}
	DataStore *CloudGcpIntegrationsDataStore `field:"optional" json:"dataStore" yaml:"dataStore"`
	// fire_base_database block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#fire_base_database CloudGcpIntegrations#fire_base_database}
	FireBaseDatabase *CloudGcpIntegrationsFireBaseDatabase `field:"optional" json:"fireBaseDatabase" yaml:"fireBaseDatabase"`
	// fire_base_hosting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#fire_base_hosting CloudGcpIntegrations#fire_base_hosting}
	FireBaseHosting *CloudGcpIntegrationsFireBaseHosting `field:"optional" json:"fireBaseHosting" yaml:"fireBaseHosting"`
	// fire_base_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#fire_base_storage CloudGcpIntegrations#fire_base_storage}
	FireBaseStorage *CloudGcpIntegrationsFireBaseStorage `field:"optional" json:"fireBaseStorage" yaml:"fireBaseStorage"`
	// fire_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#fire_store CloudGcpIntegrations#fire_store}
	FireStore *CloudGcpIntegrationsFireStore `field:"optional" json:"fireStore" yaml:"fireStore"`
	// functions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#functions CloudGcpIntegrations#functions}
	Functions *CloudGcpIntegrationsFunctions `field:"optional" json:"functions" yaml:"functions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#id CloudGcpIntegrations#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// interconnect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#interconnect CloudGcpIntegrations#interconnect}
	Interconnect *CloudGcpIntegrationsInterconnect `field:"optional" json:"interconnect" yaml:"interconnect"`
	// kubernetes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#kubernetes CloudGcpIntegrations#kubernetes}
	Kubernetes *CloudGcpIntegrationsKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// load_balancing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#load_balancing CloudGcpIntegrations#load_balancing}
	LoadBalancing *CloudGcpIntegrationsLoadBalancing `field:"optional" json:"loadBalancing" yaml:"loadBalancing"`
	// mem_cache block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#mem_cache CloudGcpIntegrations#mem_cache}
	MemCache *CloudGcpIntegrationsMemCache `field:"optional" json:"memCache" yaml:"memCache"`
	// pub_sub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#pub_sub CloudGcpIntegrations#pub_sub}
	PubSub *CloudGcpIntegrationsPubSub `field:"optional" json:"pubSub" yaml:"pubSub"`
	// redis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#redis CloudGcpIntegrations#redis}
	Redis *CloudGcpIntegrationsRedis `field:"optional" json:"redis" yaml:"redis"`
	// router block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#router CloudGcpIntegrations#router}
	Router *CloudGcpIntegrationsRouter `field:"optional" json:"router" yaml:"router"`
	// run block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#run CloudGcpIntegrations#run}
	Run *CloudGcpIntegrationsRun `field:"optional" json:"run" yaml:"run"`
	// spanner block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#spanner CloudGcpIntegrations#spanner}
	Spanner *CloudGcpIntegrationsSpanner `field:"optional" json:"spanner" yaml:"spanner"`
	// sql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#sql CloudGcpIntegrations#sql}
	Sql *CloudGcpIntegrationsSql `field:"optional" json:"sql" yaml:"sql"`
	// storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#storage CloudGcpIntegrations#storage}
	Storage *CloudGcpIntegrationsStorage `field:"optional" json:"storage" yaml:"storage"`
	// virtual_machines block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#virtual_machines CloudGcpIntegrations#virtual_machines}
	VirtualMachines *CloudGcpIntegrationsVirtualMachines `field:"optional" json:"virtualMachines" yaml:"virtualMachines"`
	// vpc_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.0/docs/resources/cloud_gcp_integrations#vpc_access CloudGcpIntegrations#vpc_access}
	VpcAccess *CloudGcpIntegrationsVpcAccess `field:"optional" json:"vpcAccess" yaml:"vpcAccess"`
}

