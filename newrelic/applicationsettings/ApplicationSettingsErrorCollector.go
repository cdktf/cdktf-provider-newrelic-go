// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationsettings


type ApplicationSettingsErrorCollector struct {
	// A list of error classes that are expected and should not trigger alerts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs/resources/application_settings#expected_error_classes ApplicationSettings#expected_error_classes}
	ExpectedErrorClasses *[]*string `field:"optional" json:"expectedErrorClasses" yaml:"expectedErrorClasses"`
	// A list of error codes that are expected and should not trigger alerts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs/resources/application_settings#expected_error_codes ApplicationSettings#expected_error_codes}
	ExpectedErrorCodes *[]*string `field:"optional" json:"expectedErrorCodes" yaml:"expectedErrorCodes"`
	// A list of error classes that should be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs/resources/application_settings#ignored_error_classes ApplicationSettings#ignored_error_classes}
	IgnoredErrorClasses *[]*string `field:"optional" json:"ignoredErrorClasses" yaml:"ignoredErrorClasses"`
	// A list of error codes that should be ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.61.3/docs/resources/application_settings#ignored_error_codes ApplicationSettings#ignored_error_codes}
	IgnoredErrorCodes *[]*string `field:"optional" json:"ignoredErrorCodes" yaml:"ignoredErrorCodes"`
}

