// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationSettingsConfig struct {
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
	// The response time threshold value for Apdex score calculation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.1/docs/resources/application_settings#app_apdex_threshold ApplicationSettings#app_apdex_threshold}
	AppApdexThreshold *float64 `field:"optional" json:"appApdexThreshold" yaml:"appApdexThreshold"`
	// Dummy field to support backward compatibility of previous version.should be removed with next major version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.1/docs/resources/application_settings#enable_real_user_monitoring ApplicationSettings#enable_real_user_monitoring}
	EnableRealUserMonitoring interface{} `field:"optional" json:"enableRealUserMonitoring" yaml:"enableRealUserMonitoring"`
	// Samples and reports the slowest database queries in your traces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.1/docs/resources/application_settings#enable_slow_sql ApplicationSettings#enable_slow_sql}
	EnableSlowSql interface{} `field:"optional" json:"enableSlowSql" yaml:"enableSlowSql"`
	// Enable or disable the thread profiler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.1/docs/resources/application_settings#enable_thread_profiler ApplicationSettings#enable_thread_profiler}
	EnableThreadProfiler interface{} `field:"optional" json:"enableThreadProfiler" yaml:"enableThreadProfiler"`
	// Dummy field to support backward compatibility of previous version.should be removed with next major version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.1/docs/resources/application_settings#end_user_apdex_threshold ApplicationSettings#end_user_apdex_threshold}
	EndUserApdexThreshold *float64 `field:"optional" json:"endUserApdexThreshold" yaml:"endUserApdexThreshold"`
	// error_collector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.1/docs/resources/application_settings#error_collector ApplicationSettings#error_collector}
	ErrorCollector interface{} `field:"optional" json:"errorCollector" yaml:"errorCollector"`
	// The GUID of the application in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.1/docs/resources/application_settings#guid ApplicationSettings#guid}
	Guid *string `field:"optional" json:"guid" yaml:"guid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.1/docs/resources/application_settings#id ApplicationSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the application in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.1/docs/resources/application_settings#name ApplicationSettings#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of tracer to use, either 'CROSS_APPLICATION_TRACER', 'DISTRIBUTED_TRACING', 'NONE', or 'OPT_OUT'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.1/docs/resources/application_settings#tracer_type ApplicationSettings#tracer_type}
	TracerType *string `field:"optional" json:"tracerType" yaml:"tracerType"`
	// transaction_tracer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.1/docs/resources/application_settings#transaction_tracer ApplicationSettings#transaction_tracer}
	TransactionTracer interface{} `field:"optional" json:"transactionTracer" yaml:"transactionTracer"`
	// Enable or disable server side monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.72.1/docs/resources/application_settings#use_server_side_config ApplicationSettings#use_server_side_config}
	UseServerSideConfig interface{} `field:"optional" json:"useServerSideConfig" yaml:"useServerSideConfig"`
}

