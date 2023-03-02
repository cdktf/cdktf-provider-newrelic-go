//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NewrelicProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (n *jsiiProxy_NewrelicProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateNewrelicProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewrelicProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateNewrelicProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_NewrelicProvider) validateSetInsecureSkipVerifyParameters(val interface{}) error {
	return nil
}

func validateNewNewrelicProviderParameters(scope constructs.Construct, id *string, config *NewrelicProviderConfig) error {
	return nil
}

