//go:build no_runtime_type_checking

package datanewrelicentity

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataNewrelicEntityTagList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataNewrelicEntityTagList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataNewrelicEntityTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataNewrelicEntityTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataNewrelicEntityTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataNewrelicEntityTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataNewrelicEntityTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
