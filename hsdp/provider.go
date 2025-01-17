package hsdp

import (
	"context"
	"os"

	"github.com/google/fhir/go/jsonformat"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns an instance of the HSDP provider
func Provider(build string) *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions["region"],
			},
			"environment": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "client-test",
				Description: descriptions["environment"],
			},
			"iam_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["iam_url"],
			},
			"idm_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["idm_url"],
			},
			"s3creds_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["s3creds_url"],
			},
			"notification_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["notification_url"],
			},
			"service_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"org_admin_username"},
				RequiredWith:  []string{"service_private_key"},
				Description:   descriptions["service_id"],
			},
			"service_private_key": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				ConflictsWith: []string{"org_admin_password"},
				RequiredWith:  []string{"service_id"},
				Description:   descriptions["service_private_key"],
			},
			"oauth2_client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["oauth2_client_id"],
			},
			"oauth2_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: descriptions["oauth2_password"],
			},
			"org_admin_username": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   descriptions["org_admin_username"],
				RequiredWith:  []string{"org_admin_password"},
				ConflictsWith: []string{"service_id"},
			},
			"org_admin_password": {
				Type:          schema.TypeString,
				Optional:      true,
				Sensitive:     true,
				Description:   descriptions["org_admin_password"],
				RequiredWith:  []string{"org_admin_username"},
				ConflictsWith: []string{"service_private_key"},
			},
			"uaa_username": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  descriptions["uaa_username"],
				RequiredWith: []string{"uaa_password"},
			},
			"uaa_password": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				Description:  descriptions["uaa_password"],
				RequiredWith: []string{"uaa_username"},
			},
			"uaa_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["uaa_url"],
			},
			"shared_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   false,
				Description: descriptions["shared_key"],
			},
			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: descriptions["secret_key"],
			},
			"cartel_host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["cartel_host"],
			},
			"cartel_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: descriptions["cartel_token"],
			},
			"cartel_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: descriptions["cartel_secret"],
			},
			"cartel_no_tls": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions["cartel_no_tls"],
			},
			"cartel_skip_verify": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: descriptions["cartel_skip_verify"],
			},
			"retry_max": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: descriptions["retry_max"],
			},
			"debug_log": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["debug_log"],
			},
			"ai_inference_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"hsdp_iam_org":                          resourceIAMOrg(),
			"hsdp_iam_group":                        resourceIAMGroup(),
			"hsdp_iam_role":                         resourceIAMRole(),
			"hsdp_iam_proposition":                  resourceIAMProposition(),
			"hsdp_iam_application":                  resourceIAMApplication(),
			"hsdp_iam_user":                         resourceIAMUser(),
			"hsdp_iam_client":                       resourceIAMClient(),
			"hsdp_iam_service":                      resourceIAMService(),
			"hsdp_iam_mfa_policy":                   resourceIAMMFAPolicy(),
			"hsdp_iam_password_policy":              resourceIAMPasswordPolicy(),
			"hsdp_iam_email_template":               resourceIAMEmailTemplate(),
			"hsdp_s3creds_policy":                   resourceS3CredsPolicy(),
			"hsdp_container_host":                   resourceContainerHost(),
			"hsdp_container_host_exec":              resourceContainerHostExec(),
			"hsdp_metrics_autoscaler":               resourceMetricsAutoscaler(),
			"hsdp_cdr_org":                          resourceCDROrg(),
			"hsdp_cdr_subscription":                 resourceCDRSubscription(),
			"hsdp_dicom_store_config":               resourceDICOMStoreConfig(),
			"hsdp_dicom_object_store":               resourceDICOMObjectStore(),
			"hsdp_dicom_repository":                 resourceDICOMRepository(),
			"hsdp_pki_tenant":                       resourcePKITenant(),
			"hsdp_pki_cert":                         resourcePKICert(),
			"hsdp_edge_app":                         resourceEdgeApp(),
			"hsdp_edge_config":                      resourceEdgeConfig(),
			"hsdp_edge_custom_cert":                 resourceEdgeCustomCert(),
			"hsdp_edge_sync":                        resourceEdgeSync(),
			"hsdp_function":                         resourceFunction(),
			"hsdp_notification_producer":            resourceNotificationProducer(),
			"hsdp_notification_subscriber":          resourceNotificationSubscriber(),
			"hsdp_notification_topic":               resourceNotificationTopic(),
			"hsdp_notification_subscription":        resourceNotificationSubscription(),
			"hsdp_ai_inference_compute_environment": resourceAIInferenceComputeEnvironment(),
			"hsdp_ai_inference_compute_target":      resourceAIInferenceComputeTarget(),
			"hsdp_ai_inference_model":               resourceAIInferenceModel(),
			"hsdp_ai_inference_job":                 resourceAIInferenceJob(),
			"hsdp_dicom_gateway_config":             resourceDICOMGatewayConfig(),
			"hsdp_cdl_research_study":               resourceCDLResearchStudy(),
			"hsdp_dicom_remote_node":                resourceDICOMRemoteNode(),
			"hsdp_cdl_data_type_definition":         resourceCDLDataTypeDefinition(),
			"hsdp_cdl_label_definition":             resourceCDLLabelDefinition(),
			"hsdp_cdl_export_route":                 resourceCDLExportRoute(),
			"hsdp_ai_workspace_compute_target":      resourceAIWorkspaceComputeTarget(),
			"hsdp_ai_workspace":                     resourceAIWorkspace(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"hsdp_iam_introspect":                    dataSourceIAMIntrospect(),
			"hsdp_iam_user":                          dataSourceUser(),
			"hsdp_iam_service":                       dataSourceService(),
			"hsdp_iam_permissions":                   dataSourceIAMPermissions(),
			"hsdp_iam_org":                           dataSourceIAMOrg(),
			"hsdp_iam_proposition":                   dataSourceIAMProposition(),
			"hsdp_iam_application":                   dataSourceIAMApplication(),
			"hsdp_s3creds_access":                    dataSourceS3CredsAccess(),
			"hsdp_s3creds_policy":                    dataSourceS3CredsPolicy(),
			"hsdp_config":                            dataSourceConfig(),
			"hsdp_container_host_subnet_types":       dataSourceContainerHostSubnetTypes(),
			"hsdp_cdr_fhir_store":                    dataSourceCDRFHIRStore(),
			"hsdp_pki_root":                          dataSourcePKIRoot(),
			"hsdp_pki_policy":                        dataSourcePKIPolicy(),
			"hsdp_edge_device":                       dataSourceEdgeDevice(),
			"hsdp_notification_producers":            dataSourceNotificationProducers(),
			"hsdp_notification_producer":             dataSourceNotificationProducer(),
			"hsdp_notification_topics":               dataSourceNotificationTopics(),
			"hsdp_notification_topic":                dataSourceNotificationTopic(),
			"hsdp_notification_subscription":         dataSourceNotificationSubscription(),
			"hsdp_notification_subscriber":           dataSourceNotificationSubscriber(),
			"hsdp_ai_inference_service_instance":     dataSourceAIInferenceServiceInstance(),
			"hsdp_ai_inference_compute_environments": dataSourceAIInferenceComputeEnvironments(),
			"hsdp_ai_inference_compute_targets":      dataSourceAIInferenceComputeTargets(),
			"hsdp_ai_inference_jobs":                 dataSourceAIInferenceJobs(),
			"hsdp_ai_inference_models":               dataSourceAIInferenceModels(),
			"hsdp_cdl_instance":                      dataSourceCDLInstance(),
			"hsdp_cdl_research_study":                dataSourceCDLResearchStudy(),
			"hsdp_cdl_research_studies":              dataSourceCDLResearchStudies(),
			"hsdp_container_host_instances":          dataSourceContainerHostInstances(),
			"hsdp_cdl_data_type_definitions":         dataSourceCDLDataTypeDefinitions(),
			"hsdp_cdl_data_type_definition":          dataSourceCDLDataTypeDefinition(),
			"hsdp_cdl_label_definition":              dataSourceCDLLabelDefinition(),
			"hsdp_cdl_export_route":                  dataSourceCDLExportRoute(),
			"hsdp_ai_workspace_service_instance":     dataSourceAIWorkspaceServiceInstance(),
			"hsdp_ai_workspace_compute_targets":      dataSourceAIWorkspaceComputeTargets(),
			"hsdp_ai_workspace":                      dataSourceAIWorkspace(),
			"hsdp_iam_group":                         dataSourceIAMGroup(),
			"hsdp_iam_role":                          dataSourceIAMRole(),
		},
		ConfigureContextFunc: providerConfigure(build),
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"region":              "The HSDP region to configure for",
		"environment":         "The HSDP environment to configure for",
		"iam_url":             "The HSDP IAM instance URL",
		"idm_url":             "The HSDP IDM instance URL",
		"s3creds_url":         "The HSDP S3 Credentials instance URL",
		"notification_url":    "The HSDP Notification service base URL to use",
		"oauth2_client_id":    "The OAuth2 client id",
		"oauth2_password":     "The OAuth2 password",
		"service_id":          "The service ID to use as Organization Admin",
		"service_private_key": "The private key of the service ID",
		"org_admin_username":  "The username of the Organization Admin",
		"org_admin_password":  "The password of the Organization Admin",
		"shared_key":          "The shared key",
		"secret_key":          "The secret key",
		"debug_log":           "The log file to write debugging output to",
		"cartel_host":         "The Cartel host",
		"cartel_token":        "The Cartel token key",
		"cartel_secret":       "The Cartel secret key",
		"cartel_no_tls":       "Disable TLS for Cartel",
		"cartel_skip_verify":  "Skip certificate verification",
		"retry_max":           "Maximum number of retries for API requests",
		"uaa_username":        "The username of the Cloudfoundry account to use",
		"uaa_password":        "The password of the Cloudfoundry account to use",
		"uaa_url":             "The URL of the UAA server",
	}
}

func providerConfigure(build string) schema.ConfigureContextFunc {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var diags diag.Diagnostics

		config := &Config{}

		config.BuildVersion = build
		config.Region = d.Get("region").(string)
		config.Environment = d.Get("environment").(string)
		config.IAMURL = d.Get("iam_url").(string)
		config.IDMURL = d.Get("idm_url").(string)
		config.OAuth2ClientID = d.Get("oauth2_client_id").(string)
		config.OAuth2Secret = d.Get("oauth2_password").(string)
		config.ServiceID = d.Get("service_id").(string)
		config.ServicePrivateKey = d.Get("service_private_key").(string)
		config.OrgAdminUsername = d.Get("org_admin_username").(string)
		config.OrgAdminPassword = d.Get("org_admin_password").(string)
		config.SharedKey = d.Get("shared_key").(string)
		config.SecretKey = d.Get("secret_key").(string)
		config.DebugLog = d.Get("debug_log").(string)
		config.S3CredsURL = d.Get("s3creds_url").(string)
		config.CartelHost = d.Get("cartel_host").(string)
		config.CartelToken = d.Get("cartel_token").(string)
		config.CartelSecret = d.Get("cartel_secret").(string)
		config.CartelNoTLS = d.Get("cartel_no_tls").(bool)
		config.CartelSkipVerify = d.Get("cartel_skip_verify").(bool)
		config.RetryMax = d.Get("retry_max").(int)
		config.UAAUsername = d.Get("uaa_username").(string)
		config.UAAPassword = d.Get("uaa_password").(string)
		config.UAAURL = d.Get("uaa_url").(string)
		config.NotificationURL = d.Get("notification_url").(string)
		config.TimeZone = "UTC"
		config.AIInferenceEndpoint = d.Get("ai_inference_endpoint").(string)

		config.setupIAMClient()
		config.setupS3CredsClient()
		config.setupCartelClient()
		config.setupConsoleClient()
		config.setupPKIClient()
		config.setupSTLClient()
		config.setupNotificationClient()

		if config.DebugLog != "" {
			debugFile, err := os.OpenFile(config.DebugLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
			if err != nil {
				config.debugFile = nil
			} else {
				config.debugFile = debugFile
			}
		}

		ma, err := jsonformat.NewMarshaller(false, "", "", jsonformat.STU3)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		config.ma = ma

		um, err := jsonformat.NewUnmarshaller("UTC", jsonformat.STU3)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		config.um = um

		return config, diags
	}
}
