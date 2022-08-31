//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OneDashboardPageWidgetLogTableList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OneDashboardPageWidgetLogTableList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OneDashboardPageWidgetLogTableList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOneDashboardPageWidgetLogTableListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
