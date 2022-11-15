package dashboard

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DashboardConfig struct {
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
	// The title of the dashboard.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#title Dashboard#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Determines who can edit the dashboard in an account.
	//
	// Valid values are all, editable_by_all, editable_by_owner, or read_only. Defaults to editable_by_all.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#editable Dashboard#editable}
	Editable *string `field:"optional" json:"editable" yaml:"editable"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#filter Dashboard#filter}
	Filter *DashboardFilter `field:"optional" json:"filter" yaml:"filter"`
	// New Relic One supports a 3 column grid or a 12 column grid.
	//
	// New Relic Insights supports a 3 column grid.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#grid_column_count Dashboard#grid_column_count}
	GridColumnCount *float64 `field:"optional" json:"gridColumnCount" yaml:"gridColumnCount"`
	// The icon for the dashboard.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#icon Dashboard#icon}
	Icon *string `field:"optional" json:"icon" yaml:"icon"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#id Dashboard#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Determines who can see the dashboard in an account. Valid values are all or owner. Defaults to all.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#visibility Dashboard#visibility}
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// widget block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/dashboard#widget Dashboard#widget}
	Widget interface{} `field:"optional" json:"widget" yaml:"widget"`
}

