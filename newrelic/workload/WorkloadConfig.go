package workload

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkloadConfig struct {
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
	// The workload's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/workload#name Workload#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The New Relic account ID where you want to create the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/workload#account_id Workload#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Relevant information about the workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/workload#description Workload#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of entity GUIDs manually assigned to this workload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/workload#entity_guids Workload#entity_guids}
	EntityGuids *[]*string `field:"optional" json:"entityGuids" yaml:"entityGuids"`
	// entity_search_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/workload#entity_search_query Workload#entity_search_query}
	EntitySearchQuery interface{} `field:"optional" json:"entitySearchQuery" yaml:"entitySearchQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/workload#id Workload#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A list of account IDs that will be used to get entities from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/workload#scope_account_ids Workload#scope_account_ids}
	ScopeAccountIds *[]*float64 `field:"optional" json:"scopeAccountIds" yaml:"scopeAccountIds"`
	// status_config_automatic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/workload#status_config_automatic Workload#status_config_automatic}
	StatusConfigAutomatic *WorkloadStatusConfigAutomatic `field:"optional" json:"statusConfigAutomatic" yaml:"statusConfigAutomatic"`
	// status_config_static block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/workload#status_config_static Workload#status_config_static}
	StatusConfigStatic *WorkloadStatusConfigStatic `field:"optional" json:"statusConfigStatic" yaml:"statusConfigStatic"`
}

