package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-newrelic.provider.NewrelicProvider",
		reflect.TypeOf((*NewrelicProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "adminApiKey", GoGetter: "AdminApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "adminApiKeyInput", GoGetter: "AdminApiKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiKey", GoGetter: "ApiKey"},
			_jsii_.MemberProperty{JsiiProperty: "apiKeyInput", GoGetter: "ApiKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "apiUrl", GoGetter: "ApiUrl"},
			_jsii_.MemberProperty{JsiiProperty: "apiUrlInput", GoGetter: "ApiUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "cacertFile", GoGetter: "CacertFile"},
			_jsii_.MemberProperty{JsiiProperty: "cacertFileInput", GoGetter: "CacertFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureApiUrl", GoGetter: "InfrastructureApiUrl"},
			_jsii_.MemberProperty{JsiiProperty: "infrastructureApiUrlInput", GoGetter: "InfrastructureApiUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "insecureSkipVerify", GoGetter: "InsecureSkipVerify"},
			_jsii_.MemberProperty{JsiiProperty: "insecureSkipVerifyInput", GoGetter: "InsecureSkipVerifyInput"},
			_jsii_.MemberProperty{JsiiProperty: "insightsInsertKey", GoGetter: "InsightsInsertKey"},
			_jsii_.MemberProperty{JsiiProperty: "insightsInsertKeyInput", GoGetter: "InsightsInsertKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "insightsInsertUrl", GoGetter: "InsightsInsertUrl"},
			_jsii_.MemberProperty{JsiiProperty: "insightsInsertUrlInput", GoGetter: "InsightsInsertUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "insightsQueryUrl", GoGetter: "InsightsQueryUrl"},
			_jsii_.MemberProperty{JsiiProperty: "insightsQueryUrlInput", GoGetter: "InsightsQueryUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "nerdgraphApiUrl", GoGetter: "NerdgraphApiUrl"},
			_jsii_.MemberProperty{JsiiProperty: "nerdgraphApiUrlInput", GoGetter: "NerdgraphApiUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdminApiKey", GoMethod: "ResetAdminApiKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiUrl", GoMethod: "ResetApiUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetCacertFile", GoMethod: "ResetCacertFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetInfrastructureApiUrl", GoMethod: "ResetInfrastructureApiUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecureSkipVerify", GoMethod: "ResetInsecureSkipVerify"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsightsInsertKey", GoMethod: "ResetInsightsInsertKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsightsInsertUrl", GoMethod: "ResetInsightsInsertUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsightsQueryUrl", GoMethod: "ResetInsightsQueryUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetNerdgraphApiUrl", GoMethod: "ResetNerdgraphApiUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetSyntheticsApiUrl", GoMethod: "ResetSyntheticsApiUrl"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "syntheticsApiUrl", GoGetter: "SyntheticsApiUrl"},
			_jsii_.MemberProperty{JsiiProperty: "syntheticsApiUrlInput", GoGetter: "SyntheticsApiUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_NewrelicProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-newrelic.provider.NewrelicProviderConfig",
		reflect.TypeOf((*NewrelicProviderConfig)(nil)).Elem(),
	)
}
