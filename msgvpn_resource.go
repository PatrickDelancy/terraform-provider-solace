package main

import (
	"fmt"
	"log"

	"github.com/ExalDraen/semp-client/models"

	"github.com/ExalDraen/semp-client/client/msg_vpn"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMsgVpn() *schema.Resource {
	return &schema.Resource{
		Create: resourceMsgVpnCreate,
		Read:   resourceMsgVpnRead,
		Update: resourceMsgVpnUpdate,
		Delete: resourceMsgVpnDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceMsgVpnCreate(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Creating msg vpn ...")

	ca := m.(ClientAndAuth)
	client := ca.Client
	auth := ca.Auth
	name := d.Get("name").(string)

	// Prepare new VPN object
	params := msg_vpn.NewCreateMsgVpnParams()
	newMsgVpn := models.MsgVpn{MsgVpnName: name}
	params.Body = &newMsgVpn

	log.Printf("[DEBUG] create msg vpn %q", name)

	resp, err := client.MsgVpn.CreateMsgVpn(params, auth)
	if err != nil {
		log.Printf("[ERROR] Failed to create VPN %q", name)
		return err
	}
	d.SetId(resp.Payload.Data.MsgVpnName)
	return resourceMsgVpnRead(d, m)
}

func resourceMsgVpnRead(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Reading msg vpn ...")
	ca := m.(ClientAndAuth)
	client := ca.Client
	auth := ca.Auth
	getParams := msg_vpn.NewGetMsgVpnParams()
	getParams.MsgVpnName = d.Id()

	resp, err := client.MsgVpn.GetMsgVpn(getParams, auth)
	if err != nil {
		log.Printf("[WARN] No Server found: %s", d.Id())
		d.SetId("")
		return nil
	}
	fmt.Printf("%#v\n", resp.Payload.Data)
	d.Set("name", resp.Payload.Data.MsgVpnName)
	return nil
}

func resourceMsgVpnUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceMsgVpnRead(d, m)
}

func resourceMsgVpnDelete(d *schema.ResourceData, m interface{}) error {
	ca := m.(ClientAndAuth)
	client := ca.Client
	auth := ca.Auth
	params := msg_vpn.NewDeleteMsgVpnParams()
	params.MsgVpnName = d.Id()

	_, err := client.MsgVpn.DeleteMsgVpn(params, auth)
	if err != nil {
		return fmt.Errorf("[ERROR] Unable to delete project %q: %s", params.MsgVpnName, err)
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}
