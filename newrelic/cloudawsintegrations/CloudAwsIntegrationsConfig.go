package cloudawsintegrations

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudAwsIntegrationsConfig struct {
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
	// The ID of the linked AWS account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#linked_account_id CloudAwsIntegrations#linked_account_id}
	LinkedAccountId *float64 `field:"required" json:"linkedAccountId" yaml:"linkedAccountId"`
	// The ID of the account in New Relic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#account_id CloudAwsIntegrations#account_id}
	AccountId *float64 `field:"optional" json:"accountId" yaml:"accountId"`
	// alb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#alb CloudAwsIntegrations#alb}
	Alb *CloudAwsIntegrationsAlb `field:"optional" json:"alb" yaml:"alb"`
	// api_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#api_gateway CloudAwsIntegrations#api_gateway}
	ApiGateway *CloudAwsIntegrationsApiGateway `field:"optional" json:"apiGateway" yaml:"apiGateway"`
	// auto_scaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#auto_scaling CloudAwsIntegrations#auto_scaling}
	AutoScaling *CloudAwsIntegrationsAutoScaling `field:"optional" json:"autoScaling" yaml:"autoScaling"`
	// aws_app_sync block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_app_sync CloudAwsIntegrations#aws_app_sync}
	AwsAppSync *CloudAwsIntegrationsAwsAppSync `field:"optional" json:"awsAppSync" yaml:"awsAppSync"`
	// aws_athena block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_athena CloudAwsIntegrations#aws_athena}
	AwsAthena *CloudAwsIntegrationsAwsAthena `field:"optional" json:"awsAthena" yaml:"awsAthena"`
	// aws_cognito block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_cognito CloudAwsIntegrations#aws_cognito}
	AwsCognito *CloudAwsIntegrationsAwsCognito `field:"optional" json:"awsCognito" yaml:"awsCognito"`
	// aws_connect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_connect CloudAwsIntegrations#aws_connect}
	AwsConnect *CloudAwsIntegrationsAwsConnect `field:"optional" json:"awsConnect" yaml:"awsConnect"`
	// aws_direct_connect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_direct_connect CloudAwsIntegrations#aws_direct_connect}
	AwsDirectConnect *CloudAwsIntegrationsAwsDirectConnect `field:"optional" json:"awsDirectConnect" yaml:"awsDirectConnect"`
	// aws_fsx block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_fsx CloudAwsIntegrations#aws_fsx}
	AwsFsx *CloudAwsIntegrationsAwsFsx `field:"optional" json:"awsFsx" yaml:"awsFsx"`
	// aws_glue block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_glue CloudAwsIntegrations#aws_glue}
	AwsGlue *CloudAwsIntegrationsAwsGlue `field:"optional" json:"awsGlue" yaml:"awsGlue"`
	// aws_kinesis_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_kinesis_analytics CloudAwsIntegrations#aws_kinesis_analytics}
	AwsKinesisAnalytics *CloudAwsIntegrationsAwsKinesisAnalytics `field:"optional" json:"awsKinesisAnalytics" yaml:"awsKinesisAnalytics"`
	// aws_media_convert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_media_convert CloudAwsIntegrations#aws_media_convert}
	AwsMediaConvert *CloudAwsIntegrationsAwsMediaConvert `field:"optional" json:"awsMediaConvert" yaml:"awsMediaConvert"`
	// aws_media_package_vod block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_media_package_vod CloudAwsIntegrations#aws_media_package_vod}
	AwsMediaPackageVod *CloudAwsIntegrationsAwsMediaPackageVod `field:"optional" json:"awsMediaPackageVod" yaml:"awsMediaPackageVod"`
	// aws_mq block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_mq CloudAwsIntegrations#aws_mq}
	AwsMq *CloudAwsIntegrationsAwsMq `field:"optional" json:"awsMq" yaml:"awsMq"`
	// aws_msk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_msk CloudAwsIntegrations#aws_msk}
	AwsMsk *CloudAwsIntegrationsAwsMsk `field:"optional" json:"awsMsk" yaml:"awsMsk"`
	// aws_neptune block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_neptune CloudAwsIntegrations#aws_neptune}
	AwsNeptune *CloudAwsIntegrationsAwsNeptune `field:"optional" json:"awsNeptune" yaml:"awsNeptune"`
	// aws_qldb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_qldb CloudAwsIntegrations#aws_qldb}
	AwsQldb *CloudAwsIntegrationsAwsQldb `field:"optional" json:"awsQldb" yaml:"awsQldb"`
	// aws_route53resolver block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_route53resolver CloudAwsIntegrations#aws_route53resolver}
	AwsRoute53Resolver *CloudAwsIntegrationsAwsRoute53Resolver `field:"optional" json:"awsRoute53Resolver" yaml:"awsRoute53Resolver"`
	// aws_states block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_states CloudAwsIntegrations#aws_states}
	AwsStates *CloudAwsIntegrationsAwsStates `field:"optional" json:"awsStates" yaml:"awsStates"`
	// aws_transit_gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_transit_gateway CloudAwsIntegrations#aws_transit_gateway}
	AwsTransitGateway *CloudAwsIntegrationsAwsTransitGateway `field:"optional" json:"awsTransitGateway" yaml:"awsTransitGateway"`
	// aws_waf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_waf CloudAwsIntegrations#aws_waf}
	AwsWaf *CloudAwsIntegrationsAwsWaf `field:"optional" json:"awsWaf" yaml:"awsWaf"`
	// aws_wafv2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#aws_wafv2 CloudAwsIntegrations#aws_wafv2}
	AwsWafv2 *CloudAwsIntegrationsAwsWafv2 `field:"optional" json:"awsWafv2" yaml:"awsWafv2"`
	// billing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#billing CloudAwsIntegrations#billing}
	Billing *CloudAwsIntegrationsBilling `field:"optional" json:"billing" yaml:"billing"`
	// cloudfront block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#cloudfront CloudAwsIntegrations#cloudfront}
	Cloudfront *CloudAwsIntegrationsCloudfront `field:"optional" json:"cloudfront" yaml:"cloudfront"`
	// cloudtrail block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#cloudtrail CloudAwsIntegrations#cloudtrail}
	Cloudtrail *CloudAwsIntegrationsCloudtrail `field:"optional" json:"cloudtrail" yaml:"cloudtrail"`
	// doc_db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#doc_db CloudAwsIntegrations#doc_db}
	DocDb *CloudAwsIntegrationsDocDb `field:"optional" json:"docDb" yaml:"docDb"`
	// dynamodb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#dynamodb CloudAwsIntegrations#dynamodb}
	Dynamodb *CloudAwsIntegrationsDynamodb `field:"optional" json:"dynamodb" yaml:"dynamodb"`
	// ebs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#ebs CloudAwsIntegrations#ebs}
	Ebs *CloudAwsIntegrationsEbs `field:"optional" json:"ebs" yaml:"ebs"`
	// ec2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#ec2 CloudAwsIntegrations#ec2}
	Ec2 *CloudAwsIntegrationsEc2 `field:"optional" json:"ec2" yaml:"ec2"`
	// ecs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#ecs CloudAwsIntegrations#ecs}
	Ecs *CloudAwsIntegrationsEcs `field:"optional" json:"ecs" yaml:"ecs"`
	// efs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#efs CloudAwsIntegrations#efs}
	Efs *CloudAwsIntegrationsEfs `field:"optional" json:"efs" yaml:"efs"`
	// elasticache block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#elasticache CloudAwsIntegrations#elasticache}
	Elasticache *CloudAwsIntegrationsElasticache `field:"optional" json:"elasticache" yaml:"elasticache"`
	// elasticbeanstalk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#elasticbeanstalk CloudAwsIntegrations#elasticbeanstalk}
	Elasticbeanstalk *CloudAwsIntegrationsElasticbeanstalk `field:"optional" json:"elasticbeanstalk" yaml:"elasticbeanstalk"`
	// elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#elasticsearch CloudAwsIntegrations#elasticsearch}
	Elasticsearch *CloudAwsIntegrationsElasticsearch `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// elb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#elb CloudAwsIntegrations#elb}
	Elb *CloudAwsIntegrationsElb `field:"optional" json:"elb" yaml:"elb"`
	// emr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#emr CloudAwsIntegrations#emr}
	Emr *CloudAwsIntegrationsEmr `field:"optional" json:"emr" yaml:"emr"`
	// health block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#health CloudAwsIntegrations#health}
	Health *CloudAwsIntegrationsHealth `field:"optional" json:"health" yaml:"health"`
	// iam block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#iam CloudAwsIntegrations#iam}
	Iam *CloudAwsIntegrationsIam `field:"optional" json:"iam" yaml:"iam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#id CloudAwsIntegrations#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// iot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#iot CloudAwsIntegrations#iot}
	Iot *CloudAwsIntegrationsIot `field:"optional" json:"iot" yaml:"iot"`
	// kinesis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#kinesis CloudAwsIntegrations#kinesis}
	Kinesis *CloudAwsIntegrationsKinesis `field:"optional" json:"kinesis" yaml:"kinesis"`
	// kinesis_firehose block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#kinesis_firehose CloudAwsIntegrations#kinesis_firehose}
	KinesisFirehose *CloudAwsIntegrationsKinesisFirehose `field:"optional" json:"kinesisFirehose" yaml:"kinesisFirehose"`
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#lambda CloudAwsIntegrations#lambda}
	Lambda *CloudAwsIntegrationsLambda `field:"optional" json:"lambda" yaml:"lambda"`
	// rds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#rds CloudAwsIntegrations#rds}
	Rds *CloudAwsIntegrationsRds `field:"optional" json:"rds" yaml:"rds"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#redshift CloudAwsIntegrations#redshift}
	Redshift *CloudAwsIntegrationsRedshift `field:"optional" json:"redshift" yaml:"redshift"`
	// route53 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#route53 CloudAwsIntegrations#route53}
	Route53 *CloudAwsIntegrationsRoute53 `field:"optional" json:"route53" yaml:"route53"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#s3 CloudAwsIntegrations#s3}
	S3 *CloudAwsIntegrationsS3 `field:"optional" json:"s3" yaml:"s3"`
	// ses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#ses CloudAwsIntegrations#ses}
	Ses *CloudAwsIntegrationsSes `field:"optional" json:"ses" yaml:"ses"`
	// sns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#sns CloudAwsIntegrations#sns}
	Sns *CloudAwsIntegrationsSns `field:"optional" json:"sns" yaml:"sns"`
	// sqs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#sqs CloudAwsIntegrations#sqs}
	Sqs *CloudAwsIntegrationsSqs `field:"optional" json:"sqs" yaml:"sqs"`
	// trusted_advisor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#trusted_advisor CloudAwsIntegrations#trusted_advisor}
	TrustedAdvisor *CloudAwsIntegrationsTrustedAdvisor `field:"optional" json:"trustedAdvisor" yaml:"trustedAdvisor"`
	// vpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#vpc CloudAwsIntegrations#vpc}
	Vpc *CloudAwsIntegrationsVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// x_ray block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/newrelic/newrelic/3.25.0/docs/resources/cloud_aws_integrations#x_ray CloudAwsIntegrations#x_ray}
	XRay *CloudAwsIntegrationsXRay `field:"optional" json:"xRay" yaml:"xRay"`
}

