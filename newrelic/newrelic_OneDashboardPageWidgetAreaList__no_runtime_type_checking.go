//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OneDashboardPageWidgetAreaList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OneDashboardPageWidgetAreaList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OneDashboardPageWidgetAreaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OneDashboardPageWidgetAreaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OneDashboardPageWidgetAreaList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OneDashboardPageWidgetAreaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOneDashboardPageWidgetAreaListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
