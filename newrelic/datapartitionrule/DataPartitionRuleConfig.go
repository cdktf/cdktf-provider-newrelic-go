package datapartitionrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataPartitionRuleConfig struct {
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
	// Whether or not this data partition rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/data_partition_rule#enabled DataPartitionRule#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The NRQL to match events for this data partition rule.
	//
	// Logs matching this criteria will be routed to the specified data partition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/data_partition_rule#nrql DataPartitionRule#nrql}
	Nrql *string `field:"required" json:"nrql" yaml:"nrql"`
	// The retention policy of the data partition data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/data_partition_rule#retention_policy DataPartitionRule#retention_policy}
	RetentionPolicy *string `field:"required" json:"retentionPolicy" yaml:"retentionPolicy"`
	// The name of the data partition where logs will be allocated once the rule is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/data_partition_rule#target_data_partition DataPartitionRule#target_data_partition}
	TargetDataPartition *string `field:"required" json:"targetDataPartition" yaml:"targetDataPartition"`
	// The account id associated with the data partition rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/data_partition_rule#account_id DataPartitionRule#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// The description of the data partition rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/data_partition_rule#description DataPartitionRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/data_partition_rule#id DataPartitionRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.20.2/docs/resources/data_partition_rule#timeouts DataPartitionRule#timeouts}
	Timeouts *DataPartitionRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

