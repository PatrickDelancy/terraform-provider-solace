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
				Type:        schema.TypeString,
				Description: "The name of the MSG VPN. Used as an identifier.",
				Required:    true,
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Whether or not the MSG VPN should be enabled.",
				Optional:    true,
				Default:     true,
			},
			"max_spool_usage": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The maximum Message Spool usage by the Message VPN, in megabytes. The default value is 0.",
				Optional:    true,
			},
			"replication_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Enable or disable the Replication feature for the Message VPN. The default value is false.",
				Optional:    true,
				Default:     false,
			},
		},
	}
}

func resourceMsgVpnCreate(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Creating msg vpn ...")

	// Get our Solace client
	ca := m.(ClientAndAuth)
	client := ca.Client
	auth := ca.Auth

	// Extract config data from resource data and prepare new VPN object
	name := d.Get("name").(string)
	newMsgVpn := models.MsgVpn{
		MsgVpnName:         name,
		Enabled:            d.Get("enabled").(bool),
		ReplicationEnabled: d.Get("replication_enabled").(bool),
	}
	// Values that do not have defaults shouldn't be provided if not configured in Terraform
	if v, ok := d.GetOk("max_spool_usage"); ok {
		newMsgVpn.MaxMsgSpoolUsage = int64(v.(int))
	}

	params := msg_vpn.NewCreateMsgVpnParams()
	params.Body = &newMsgVpn

	log.Printf("[DEBUG] create msg vpn %v", name)

	resp, err := client.MsgVpn.CreateMsgVpn(params, auth)
	if err != nil {
		sempErr := err.(*msg_vpn.CreateMsgVpnDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create message VPN %q: %v", name, formatError(sempErr))
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
	d.Set("enabled", resp.Payload.Data.Enabled)
	d.Set("max_spool_usage", resp.Payload.Data.MaxMsgSpoolUsage)
	d.Set("replication_enabled", resp.Payload.Data.ReplicationEnabled)
	return nil
}

func resourceMsgVpnUpdate(d *schema.ResourceData, m interface{}) error {
	ca := m.(ClientAndAuth)
	client := ca.Client
	auth := ca.Auth

	params := msg_vpn.NewUpdateMsgVpnParams()
	params.MsgVpnName = d.Id()
	newMsgVpn := models.MsgVpn{}

	// Only include changed values; anything we don't specify does not get updated
	if d.HasChange("enabled") {
		newMsgVpn.Enabled = d.Get("enabled").(bool)
	}
	if d.HasChange("max_spool_usage") {
		newMsgVpn.MaxMsgSpoolUsage = int64(d.Get("max_spool_usage").(int))
	}
	if d.HasChange("replication_enabled") {
		newMsgVpn.ReplicationEnabled = d.Get("replication_enabled").(bool)
	}
	params.Body = &newMsgVpn

	log.Printf("[DEBUG] update solace msg VPN group %s", d.Id())
	log.Printf("[TRACE] msg VPN params are %+v", params)
	log.Printf("[TRACE] msg VPN Body: %+v", params.Body)
	_, err := client.MsgVpn.UpdateMsgVpn(params, auth)
	if err != nil {
		sempErr := err.(*msg_vpn.UpdateMsgVpnDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to update message VPN %q: %v", params.MsgVpnName, formatError(sempErr))
	}

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
		sempErr := err.(*msg_vpn.DeleteMsgVpnDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to delete message VPN %q: %v", params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}
