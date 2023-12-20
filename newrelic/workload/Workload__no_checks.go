// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package workload

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_Workload) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (w *jsiiProxy_Workload) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (w *jsiiProxy_Workload) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Workload) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Workload) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Workload) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Workload) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Workload) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Workload) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Workload) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Workload) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Workload) validateImportFromParameters(id *string) error {
	return nil
}

func (w *jsiiProxy_Workload) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_Workload) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (w *jsiiProxy_Workload) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (w *jsiiProxy_Workload) validateMoveToIdParameters(id *string) error {
	return nil
}

func (w *jsiiProxy_Workload) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (w *jsiiProxy_Workload) validatePutEntitySearchQueryParameters(value interface{}) error {
	return nil
}

func (w *jsiiProxy_Workload) validatePutStatusConfigAutomaticParameters(value *WorkloadStatusConfigAutomatic) error {
	return nil
}

func (w *jsiiProxy_Workload) validatePutStatusConfigStaticParameters(value *WorkloadStatusConfigStatic) error {
	return nil
}

func validateWorkload_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateWorkload_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWorkload_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateWorkload_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Workload) validateSetAccountIdParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_Workload) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Workload) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Workload) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Workload) validateSetEntityGuidsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Workload) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Workload) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Workload) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Workload) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Workload) validateSetScopeAccountIdsParameters(val *[]*float64) error {
	return nil
}

func validateNewWorkloadParameters(scope constructs.Construct, id *string, config *WorkloadConfig) error {
	return nil
}

