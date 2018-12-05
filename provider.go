package main

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
				DefaultFunc: schema.EnvDefaultFunc("SOLACE_BASE_PATH", ""),
				Description: descriptions["base_path"],
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"solace_msgvpn": resourceMsgVpn(),
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
	}
}

// providerConfigure returns a configured Solace ClientAndAuth
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		User:     d.Get("user").(string),
		Password: d.Get("password").(string),
		BasePath: d.Get("base_path").(string),
		Host:     d.Get("host").(string),
	}

	return config.Client()
}
