//go:build no_runtime_type_checking

package workload

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkloadStatusConfigAutomaticRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkloadStatusConfigAutomaticRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkloadStatusConfigAutomaticRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

