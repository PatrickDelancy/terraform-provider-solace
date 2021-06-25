package solace

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider is the main entry point for all resources defined by this terraform plugin
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SOLACE_USER", nil),
				Description: descriptions["user"],
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SOLACE_PASSWORD", nil),
				Description: descriptions["password"],
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SOLACE_HOST", ""),
				Description: descriptions["host"],
			},
			"base_path": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SOLACE_BASE_PATH", "/SEMP/v2/config"),
				Description: descriptions["base_path"],
			},
			"msg_vpn": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SOLACE_MSG_VPN", ""),
				Description: descriptions["msg_vpn"],
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"solace_msgvpn":                                 resourceMsgVpn(),
			"solace_queue":                                  resourceQueue(),
			"solace_queue_subscription":                     resourceQueueSubscription(),
			"solace_queuetemplate":                          resourceQueueTemplate(),
			"solace_topicendpoint":                          resourceTopicEndpoint(),
			"solace_topicendpointtemplate":                  resourceTopicEndpointTemplate(),
			"solace_aclprofile":                             resourceACLProfile(),
			"solace_aclprofile_clientconnexception":         resourceACLClientConnException(),
			"solace_aclprofile_publishexception":            resourceACLPublishException(),
			"solace_aclprofile_subscribeexception":          resourceACLSubscribeException(),
			"solace_aclprofile_subscribesharenameexception": resourceACLSubscribeShareNameException(),
			"solace_clientprofile":                          resourceClientProfile(),
			"solace_clientusername":                         resourceClientUsername(),
		},

		ConfigureFunc: providerConfigure,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"user":      "The Solace management user",
		"password":  "The password for the Solace management user",
		"host":      "URL of the Solace host ",
		"base_path": "The Solace SEMP v2 base API path. Usually something like '/SEMP/v2/config'",
		"msg_vpn":   "The default Solace MSG VPN to use for resources.",
	}
}

// providerConfigure fully configures a solace provider and returns
// a config struct ready to be used by all resources and data sources.
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		User:     d.Get("user").(string),
		Password: d.Get("password").(string),
		BasePath: d.Get("base_path").(string),
		Host:     d.Get("host").(string),
		MsgVPN:   d.Get("msg_vpn").(string),
	}

	if err := config.loadAndValidate(); err != nil {
		return nil, err
	}
	return &config, nil
}
