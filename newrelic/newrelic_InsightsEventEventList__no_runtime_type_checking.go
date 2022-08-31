//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_InsightsEventEventList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_InsightsEventEventList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_InsightsEventEventList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_InsightsEventEventList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_InsightsEventEventList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_InsightsEventEventList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewInsightsEventEventListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

