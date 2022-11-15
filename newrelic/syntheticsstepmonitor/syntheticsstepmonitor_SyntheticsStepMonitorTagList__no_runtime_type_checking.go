//go:build no_runtime_type_checking

package syntheticsstepmonitor

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SyntheticsStepMonitorTagList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SyntheticsStepMonitorTagList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SyntheticsStepMonitorTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SyntheticsStepMonitorTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SyntheticsStepMonitorTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SyntheticsStepMonitorTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSyntheticsStepMonitorTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

