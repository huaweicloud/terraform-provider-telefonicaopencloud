package telefonicaopencloud

import (
	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// This is a global MutexKV for use within this plugin.
var osMutexKV = mutexkv.NewMutexKV()

// Provider returns a schema.Provider for TelefonicaOpenCloud.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_ACCESS_KEY", ""),
				Description: descriptions["access_key"],
			},

			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_SECRET_KEY", ""),
				Description: descriptions["secret_key"],
			},

			"auth_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_AUTH_URL", ""),
				Description: descriptions["auth_url"],
			},

			"region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["region"],
				DefaultFunc: schema.EnvDefaultFunc("OS_REGION_NAME", ""),
			},

			"user_name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USERNAME", ""),
				Description: descriptions["user_name"],
			},

			"user_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USER_ID", ""),
				Description: descriptions["user_name"],
			},

			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_ID",
					"OS_PROJECT_ID",
				}, ""),
				Description: descriptions["tenant_id"],
			},

			"tenant_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_TENANT_NAME",
					"OS_PROJECT_NAME",
				}, ""),
				Description: descriptions["tenant_name"],
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("OS_PASSWORD", ""),
				Description: descriptions["password"],
			},

			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_AUTH_TOKEN", ""),
				Description: descriptions["token"],
			},

			"domain_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_USER_DOMAIN_ID",
					"OS_PROJECT_DOMAIN_ID",
					"OS_DOMAIN_ID",
				}, ""),
				Description: descriptions["domain_id"],
			},

			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"OS_USER_DOMAIN_NAME",
					"OS_PROJECT_DOMAIN_NAME",
					"OS_DOMAIN_NAME",
					"OS_DEFAULT_DOMAIN",
				}, ""),
				Description: descriptions["domain_name"],
			},

			"insecure": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_INSECURE", ""),
				Description: descriptions["insecure"],
			},

			"endpoint_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_ENDPOINT_TYPE", ""),
			},

			"cacert_file": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CACERT", ""),
				Description: descriptions["cacert_file"],
			},

			"cert": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CERT", ""),
				Description: descriptions["cert"],
			},

			"key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_KEY", ""),
				Description: descriptions["key"],
			},

			"swauth": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_SWAUTH", ""),
				Description: descriptions["swauth"],
			},

			"use_octavia": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_USE_OCTAVIA", ""),
				Description: descriptions["use_octavia"],
			},

			"cloud": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OS_CLOUD", ""),
				Description: descriptions["cloud"],
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"telefonicaopencloud_dns_zone_v2":            dataSourceDNSZoneV2(),
			"telefonicaopencloud_images_image_v2":        dataSourceImagesImageV2(),
			"telefonicaopencloud_networking_network_v2":  dataSourceNetworkingNetworkV2(),
			"telefonicaopencloud_networking_subnet_v2":   dataSourceNetworkingSubnetV2(),
			"telefonicaopencloud_networking_secgroup_v2": dataSourceNetworkingSecGroupV2(),
			"telefonicaopencloud_rds_flavors_v1":         dataSourceRdsFlavorV1(),
			"telefonicaopencloud_elb_quota":              dataResourceELBQuota(),
			"telefonicaopencloud_s3_bucket_object":       dataSourceS3BucketObject(),
			"telefonicaopencloud_kms_key_v1":             dataSourceKmsKeyV1(),
			"telefonicaopencloud_kms_data_key_v1":        dataSourceKmsDataKeyV1(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"telefonicaopencloud_blockstorage_volume_v2":          resourceBlockStorageVolumeV2(),
			"telefonicaopencloud_compute_instance_v2":             resourceComputeInstanceV2(),
			"telefonicaopencloud_compute_keypair_v2":              resourceComputeKeypairV2(),
			"telefonicaopencloud_compute_secgroup_v2":             resourceComputeSecGroupV2(),
			"telefonicaopencloud_compute_servergroup_v2":          resourceComputeServerGroupV2(),
			"telefonicaopencloud_compute_floatingip_v2":           resourceComputeFloatingIPV2(),
			"telefonicaopencloud_compute_floatingip_associate_v2": resourceComputeFloatingIPAssociateV2(),
			"telefonicaopencloud_compute_volume_attach_v2":        resourceComputeVolumeAttachV2(),
			"telefonicaopencloud_dns_recordset_v2":                resourceDNSRecordSetV2(),
			"telefonicaopencloud_dns_zone_v2":                     resourceDNSZoneV2(),
			"telefonicaopencloud_images_image_v2":                 resourceImagesImageV2(),
			"telefonicaopencloud_kms_key_v1":                      resourceKmsKeyV1(),
			"telefonicaopencloud_elb_loadbalancer":                resourceELBLoadBalancer(),
			"telefonicaopencloud_elb_listener":                    resourceELBListener(),
			"telefonicaopencloud_elb_healthcheck":                 resourceELBHealthCheck(),
			"telefonicaopencloud_elb_backendecs":                  resourceELBBackendECS(),
			"telefonicaopencloud_elb_certificate":                 resourceELBCertificate(),
			"telefonicaopencloud_networking_network_v2":           resourceNetworkingNetworkV2(),
			"telefonicaopencloud_networking_subnet_v2":            resourceNetworkingSubnetV2(),
			"telefonicaopencloud_networking_floatingip_v2":        resourceNetworkingFloatingIPV2(),
			"telefonicaopencloud_networking_port_v2":              resourceNetworkingPortV2(),
			"telefonicaopencloud_networking_router_v2":            resourceNetworkingRouterV2(),
			"telefonicaopencloud_networking_router_interface_v2":  resourceNetworkingRouterInterfaceV2(),
			"telefonicaopencloud_networking_router_route_v2":      resourceNetworkingRouterRouteV2(),
			"telefonicaopencloud_networking_secgroup_v2":          resourceNetworkingSecGroupV2(),
			"telefonicaopencloud_networking_secgroup_rule_v2":     resourceNetworkingSecGroupRuleV2(),
			"telefonicaopencloud_as_group_v1":                     resourceASGroup(),
			"telefonicaopencloud_as_configuration_v1":             resourceASConfiguration(),
			"telefonicaopencloud_as_policy_v1":                    resourceASPolicy(),
			"telefonicaopencloud_vpc_eip_v1":                      resourceVpcEIPV1(),
			"telefonicaopencloud_ces_alarmrule":                   resourceAlarmRule(),
			"telefonicaopencloud_smn_topic_v2":                    resourceTopic(),
			"telefonicaopencloud_smn_subscription_v2":             resourceSubscription(),
			"telefonicaopencloud_rds_instance_v1":                 resourceRdsInstance(),
			"telefonicaopencloud_s3_bucket":                       resourceS3Bucket(),
			"telefonicaopencloud_s3_bucket_policy":                resourceS3BucketPolicy(),
			"telefonicaopencloud_s3_bucket_object":                resourceS3BucketObject(),
			"telefonicaopencloud_mrs_cluster_v1":                  resourceMRSClusterV1(),
			"telefonicaopencloud_mrs_job_v1":                      resourceMRSJobV1(),
		},

		ConfigureFunc: configureProvider,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"auth_url": "The Identity authentication URL.",

		"region": "The TelefonicaOpenCloud region to connect to.",

		"user_name": "Username to login with.",

		"user_id": "User ID to login with.",

		"tenant_id": "The ID of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"tenant_name": "The name of the Tenant (Identity v2) or Project (Identity v3)\n" +
			"to login with.",

		"password": "Password to login with.",

		"token": "Authentication token to use as an alternative to username/password.",

		"domain_id": "The ID of the Domain to scope to (Identity v3).",

		"domain_name": "The name of the Domain to scope to (Identity v3).",

		"insecure": "Trust self-signed certificates.",

		"cacert_file": "A Custom CA certificate.",

		"endpoint_type": "The catalog endpoint type to use.",

		"cert": "A client certificate to authenticate with.",

		"key": "A client private key to authenticate with.",

		"swauth": "Use Swift's authentication system instead of Keystone. Only used for\n" +
			"interaction with Swift.",

		"use_octavia": "If set to `true`, API requests will go the Load Balancer\n" +
			"service (Octavia) instead of the Networking service (Neutron).",

		"cloud": "An entry in a `clouds.yaml` file to use.",
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccessKey:        d.Get("access_key").(string),
		SecretKey:        d.Get("secret_key").(string),
		CACertFile:       d.Get("cacert_file").(string),
		ClientCertFile:   d.Get("cert").(string),
		ClientKeyFile:    d.Get("key").(string),
		Cloud:            d.Get("cloud").(string),
		DomainID:         d.Get("domain_id").(string),
		DomainName:       d.Get("domain_name").(string),
		EndpointType:     d.Get("endpoint_type").(string),
		IdentityEndpoint: d.Get("auth_url").(string),
		Insecure:         d.Get("insecure").(bool),
		Password:         d.Get("password").(string),
		Region:           d.Get("region").(string),
		Swauth:           d.Get("swauth").(bool),
		Token:            d.Get("token").(string),
		TenantID:         d.Get("tenant_id").(string),
		TenantName:       d.Get("tenant_name").(string),
		Username:         d.Get("user_name").(string),
		UserID:           d.Get("user_id").(string),
		useOctavia:       d.Get("use_octavia").(bool),
	}

	if err := config.LoadAndValidate(); err != nil {
		return nil, err
	}

	return &config, nil
}
