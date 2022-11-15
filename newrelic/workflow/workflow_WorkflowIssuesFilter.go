package workflow


type WorkflowIssuesFilter struct {
	// (Required) Filter's name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#name Workflow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// (Required) The type of the filter. One of: (FILTER, VIEW).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#type Workflow#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// predicate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/newrelic/r/workflow#predicate Workflow#predicate}
	Predicate interface{} `field:"optional" json:"predicate" yaml:"predicate"`
}

