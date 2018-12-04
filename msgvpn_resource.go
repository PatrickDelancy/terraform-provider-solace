package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMsgVpn() *schema.Resource {
	return &schema.Resource{
		Create: resourceMsgVpnCreate,
		Read:   resourceMsgVpnRead,
		Update: resourceMsgVpnUpdate,
		Delete: resourceMsgVpnDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceMsgVpnCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	d.SetId(address)
	return resourceMsgVpnRead(d, m)
}

func resourceMsgVpnRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMsgVpnUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceMsgVpnRead(d, m)
}

func resourceMsgVpnDelete(d *schema.ResourceData, m interface{}) error {
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}
