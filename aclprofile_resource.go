package main

import (
	"fmt"
	"log"

	"github.com/ExalDraen/semp-client/client/operations"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceACLProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceACLProfileCreate,
		Read:   resourceACLProfileRead,
		Update: resourceACLProfileUpdate,
		Delete: resourceACLProfileDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the ACL profile. Used as a unique identifier.",
				Required:    true,
			},
			// Each ACL must belong to a VPN, but optionally we use the provider set default,
			// and bail if neither is set. Thus the parameter is optional
			"msg_vpn": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the MSG VPN. If unset the provider default is used.",
				Optional:    true,
			},
		},
	}
}

func resourceACLProfileCreate(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Creating msg vpn ...")

	// Get our Solace client
	// ca := m.(ClientAndAuth)
	// client := ca.Client
	// auth := ca.Auth

	// // Extract config data from resource data and prepare new VPN object
	// name := d.Get("name").(string)

	return resourceACLProfileRead(d, m)
}

func resourceACLProfileRead(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Reading ACL Profile ...")
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	getParams := operations.NewGetMsgVpnACLProfileParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	getParams.ACLProfileName = d.Id()
	getParams.MsgVpnName = vpn

	resp, err := client.Operations.GetMsgVpnACLProfile(getParams, auth)
	if err != nil {
		log.Printf("[WARN] No ACL profile found: %s", d.Id())
		d.SetId("")
		return nil
	}
	fmt.Printf("%#v\n", resp.Payload.Data)
	d.Set("name", resp.Payload.Data.ACLProfileName)
	d.Set("msg_vpn", resp.Payload.Data.MsgVpnName)
	return nil
}

func resourceACLProfileUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceACLProfileRead(d, m)
}

func resourceACLProfileDelete(d *schema.ResourceData, m interface{}) error {

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}
