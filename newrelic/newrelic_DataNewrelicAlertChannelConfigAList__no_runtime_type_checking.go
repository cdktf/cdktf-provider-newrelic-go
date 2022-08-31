//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt newrelic Provider for Terraform CDK (cdktf)
package newrelic

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataNewrelicAlertChannelConfigAList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataNewrelicAlertChannelConfigAList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataNewrelicAlertChannelConfigAList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataNewrelicAlertChannelConfigAList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataNewrelicAlertChannelConfigAList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataNewrelicAlertChannelConfigAListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

