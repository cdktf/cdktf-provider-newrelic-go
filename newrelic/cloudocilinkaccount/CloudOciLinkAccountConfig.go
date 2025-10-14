// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudocilinkaccount

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudOciLinkAccountConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The New Relic compartment OCID in OCI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#compartment_ocid CloudOciLinkAccount#compartment_ocid}
	CompartmentOcid *string `field:"required" json:"compartmentOcid" yaml:"compartmentOcid"`
	// The linked account name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#name CloudOciLinkAccount#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The client ID for OCI WIF.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#oci_client_id CloudOciLinkAccount#oci_client_id}
	OciClientId *string `field:"required" json:"ociClientId" yaml:"ociClientId"`
	// The client secret for OCI WIF.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#oci_client_secret CloudOciLinkAccount#oci_client_secret}
	OciClientSecret *string `field:"required" json:"ociClientSecret" yaml:"ociClientSecret"`
	// The OCI domain URL for WIF.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#oci_domain_url CloudOciLinkAccount#oci_domain_url}
	OciDomainUrl *string `field:"required" json:"ociDomainUrl" yaml:"ociDomainUrl"`
	// The home region of the tenancy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#oci_home_region CloudOciLinkAccount#oci_home_region}
	OciHomeRegion *string `field:"required" json:"ociHomeRegion" yaml:"ociHomeRegion"`
	// The service user name for OCI WIF.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#oci_svc_user_name CloudOciLinkAccount#oci_svc_user_name}
	OciSvcUserName *string `field:"required" json:"ociSvcUserName" yaml:"ociSvcUserName"`
	// The OCI tenant identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#tenant_id CloudOciLinkAccount#tenant_id}
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// The New Relic account ID where you want to link the OCI account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#account_id CloudOciLinkAccount#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#id CloudOciLinkAccount#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The OCI ingest secret OCID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#ingest_vault_ocid CloudOciLinkAccount#ingest_vault_ocid}
	IngestVaultOcid *string `field:"optional" json:"ingestVaultOcid" yaml:"ingestVaultOcid"`
	// Specifies the type of integration, such as metrics, logs, or a combination of logs and metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#instrumentation_type CloudOciLinkAccount#instrumentation_type}
	InstrumentationType *string `field:"optional" json:"instrumentationType" yaml:"instrumentationType"`
	// The Logging stack identifier for the OCI account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#logging_stack_ocid CloudOciLinkAccount#logging_stack_ocid}
	LoggingStackOcid *string `field:"optional" json:"loggingStackOcid" yaml:"loggingStackOcid"`
	// The metric stack identifier for the OCI account. This field is only used for updates, not during initial creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#metric_stack_ocid CloudOciLinkAccount#metric_stack_ocid}
	MetricStackOcid *string `field:"optional" json:"metricStackOcid" yaml:"metricStackOcid"`
	// The OCI region for the account. This field is only used for updates, not during initial creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#oci_region CloudOciLinkAccount#oci_region}
	OciRegion *string `field:"optional" json:"ociRegion" yaml:"ociRegion"`
	// The user secret OCID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.70.6/docs/resources/cloud_oci_link_account#user_vault_ocid CloudOciLinkAccount#user_vault_ocid}
	UserVaultOcid *string `field:"optional" json:"userVaultOcid" yaml:"userVaultOcid"`
}

