//go:build no_runtime_type_checking

package onedashboard

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OneDashboardVariableList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OneDashboardVariableList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OneDashboardVariableList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OneDashboardVariableList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OneDashboardVariableList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OneDashboardVariableList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOneDashboardVariableListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

