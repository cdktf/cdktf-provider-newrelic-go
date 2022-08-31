//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DashboardWidgetCompareWithList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DashboardWidgetCompareWithList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetCompareWithList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetCompareWithList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetCompareWithList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DashboardWidgetCompareWithList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDashboardWidgetCompareWithListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
