package main

import (
	"fmt"
	"log"

	"github.com/ExalDraen/semp-client/client/msg_vpn"

	"github.com/ExalDraen/semp-client/models"

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
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceACLProfileCreate(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Creating ACL Profile ...")

	// Get our Solace client
	c := m.(*Config)
	client := c.Client
	auth := c.Auth

	// Extract config data from resource data and prepare new VPN object
	name := d.Get("name").(string)
	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	acl := models.MsgVpnACLProfile{
		ACLProfileName: name,
		MsgVpnName:     vpn,
	}

	params := operations.NewCreateMsgVpnACLProfileParams()
	params.Body = &acl

	resp, err := client.Operations.CreateMsgVpnACLProfile(params, auth)
	if err != nil {
		sempErr := err.(*operations.CreateMsgVpnACLProfileDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create message VPN %q on vpn %q: %v", name, vpn, formatError(sempErr))
	}
	d.SetId(resp.Payload.Data.ACLProfileName)

	log.Printf("[DEBUG] Finished creating ACL %q on VPN %q", name, vpn)
	return resourceACLProfileRead(d, m)
}

func resourceACLProfileRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Reading ACL Profile %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := operations.NewGetMsgVpnACLProfileParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.ACLProfileName = d.Id()
	params.MsgVpnName = vpn

	resp, err := client.Operations.GetMsgVpnACLProfile(params, auth)
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
	log.Printf("[DEBUG] Updating ACL Profile %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := operations.NewUpdateMsgVpnACLProfileParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.ACLProfileName = d.Id()
	params.MsgVpnName = vpn

	_, err = client.Operations.UpdateMsgVpnACLProfile(params, auth)
	if err != nil {
		sempErr := err.(*msg_vpn.UpdateMsgVpnDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to update ACL %q: %v", params.ACLProfileName, formatError(sempErr))
	}

	return resourceACLProfileRead(d, m)
}

func resourceACLProfileDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := operations.NewDeleteMsgVpnACLProfileParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}
	params.ACLProfileName = d.Id()
	params.MsgVpnName = vpn

	_, err = client.Operations.DeleteMsgVpnACLProfile(params, auth)
	if err != nil {
		sempErr := err.(*operations.DeleteMsgVpnACLProfileDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to delete ACL %q: %v", params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}
