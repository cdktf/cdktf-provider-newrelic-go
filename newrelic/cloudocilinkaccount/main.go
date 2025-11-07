// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudocilinkaccount

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.cloudOciLinkAccount.CloudOciLinkAccount",
		reflect.TypeOf((*CloudOciLinkAccount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "compartmentOcid", GoGetter: "CompartmentOcid"},
			_jsii_.MemberProperty{JsiiProperty: "compartmentOcidInput", GoGetter: "CompartmentOcidInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "ingestVaultOcid", GoGetter: "IngestVaultOcid"},
			_jsii_.MemberProperty{JsiiProperty: "ingestVaultOcidInput", GoGetter: "IngestVaultOcidInput"},
			_jsii_.MemberProperty{JsiiProperty: "instrumentationType", GoGetter: "InstrumentationType"},
			_jsii_.MemberProperty{JsiiProperty: "instrumentationTypeInput", GoGetter: "InstrumentationTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "loggingStackOcid", GoGetter: "LoggingStackOcid"},
			_jsii_.MemberProperty{JsiiProperty: "loggingStackOcidInput", GoGetter: "LoggingStackOcidInput"},
			_jsii_.MemberProperty{JsiiProperty: "metricStackOcid", GoGetter: "MetricStackOcid"},
			_jsii_.MemberProperty{JsiiProperty: "metricStackOcidInput", GoGetter: "MetricStackOcidInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ociClientId", GoGetter: "OciClientId"},
			_jsii_.MemberProperty{JsiiProperty: "ociClientIdInput", GoGetter: "OciClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "ociClientSecret", GoGetter: "OciClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "ociClientSecretInput", GoGetter: "OciClientSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "ociDomainUrl", GoGetter: "OciDomainUrl"},
			_jsii_.MemberProperty{JsiiProperty: "ociDomainUrlInput", GoGetter: "OciDomainUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "ociHomeRegion", GoGetter: "OciHomeRegion"},
			_jsii_.MemberProperty{JsiiProperty: "ociHomeRegionInput", GoGetter: "OciHomeRegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "ociRegion", GoGetter: "OciRegion"},
			_jsii_.MemberProperty{JsiiProperty: "ociRegionInput", GoGetter: "OciRegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountId", GoMethod: "ResetAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIngestVaultOcid", GoMethod: "ResetIngestVaultOcid"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstrumentationType", GoMethod: "ResetInstrumentationType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingStackOcid", GoMethod: "ResetLoggingStackOcid"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetricStackOcid", GoMethod: "ResetMetricStackOcid"},
			_jsii_.MemberMethod{JsiiMethod: "resetOciRegion", GoMethod: "ResetOciRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserVaultOcid", GoMethod: "ResetUserVaultOcid"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tenantId", GoGetter: "TenantId"},
			_jsii_.MemberProperty{JsiiProperty: "tenantIdInput", GoGetter: "TenantIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userVaultOcid", GoGetter: "UserVaultOcid"},
			_jsii_.MemberProperty{JsiiProperty: "userVaultOcidInput", GoGetter: "UserVaultOcidInput"},
		},
		func() interface{} {
			j := jsiiProxy_CloudOciLinkAccount{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-newrelic.cloudOciLinkAccount.CloudOciLinkAccountConfig",
		reflect.TypeOf((*CloudOciLinkAccountConfig)(nil)).Elem(),
	)
}
