//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SyntheticsMonitorScriptLocationList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsMonitorScriptLocationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SyntheticsMonitorScriptLocationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SyntheticsMonitorScriptLocationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SyntheticsMonitorScriptLocationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SyntheticsMonitorScriptLocationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSyntheticsMonitorScriptLocationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

