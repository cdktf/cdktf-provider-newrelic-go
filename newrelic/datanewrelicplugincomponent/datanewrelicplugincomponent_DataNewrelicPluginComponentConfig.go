package datanewrelicplugincomponent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataNewrelicPluginComponentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The name of the plugin component.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/plugin_component#name DataNewrelicPluginComponent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the plugin instance this component belongs to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/d/plugin_component#plugin_id DataNewrelicPluginComponent#plugin_id}
	PluginId *float64 `field:"required" json:"pluginId" yaml:"pluginId"`
}

