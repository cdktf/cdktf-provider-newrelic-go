package obfuscationrule


type ObfuscationRuleAction struct {
	// Attribute names for action. An empty list applies the action to all the attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/obfuscation_rule#attribute ObfuscationRule#attribute}
	Attribute *[]*string `field:"required" json:"attribute" yaml:"attribute"`
	// Expression Id for action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/obfuscation_rule#expression_id ObfuscationRule#expression_id}
	ExpressionId *string `field:"required" json:"expressionId" yaml:"expressionId"`
	// Obfuscation method to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.24.2/docs/resources/obfuscation_rule#method ObfuscationRule#method}
	Method *string `field:"required" json:"method" yaml:"method"`
}

