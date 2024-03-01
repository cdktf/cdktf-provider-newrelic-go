// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package group

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_Group) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (g *jsiiProxy_Group) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_Group) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Group) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Group) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Group) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Group) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Group) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Group) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Group) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Group) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Group) validateImportFromParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_Group) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_Group) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_Group) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (g *jsiiProxy_Group) validateMoveToIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_Group) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateGroup_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGroup_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateGroup_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Group) validateSetAuthenticationDomainIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Group) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Group) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Group) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Group) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Group) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Group) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Group) validateSetUserIdsParameters(val *[]*string) error {
	return nil
}

func validateNewGroupParameters(scope constructs.Construct, id *string, config *GroupConfig) error {
	return nil
}

