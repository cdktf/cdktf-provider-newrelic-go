//go:build no_runtime_type_checking

package syntheticsstepmonitor

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SyntheticsStepMonitorStepsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsStepMonitorStepsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SyntheticsStepMonitorStepsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SyntheticsStepMonitorStepsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SyntheticsStepMonitorStepsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SyntheticsStepMonitorStepsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSyntheticsStepMonitorStepsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
