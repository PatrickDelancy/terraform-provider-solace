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
			"authentication_basic_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Enable or disable Basic Authentication for clients connecting to the Message VPN. The default value is true.",
				Optional:    true,
				Default:     true,
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Whether or not the MSG VPN should be enabled.",
				Optional:    true,
				Default:     true,
			},
			"max_connection_count": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The maximum number of client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the hardware. The default is 100.",
				Optional:    true,
				Default:     100,
			},
			"max_egress_flow_count": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The maximum number of egress flows that can be created in the Message VPN. The default value is 1000.",
				Optional:    true,
				Default:     1000,
			},
			"max_endpoint_count": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The maximum number of Queues and Topic Endpoints that can be created in the Message VPN. The default value is 1000.",
				Optional:    true,
				Default:     1000,
			},
			"max_ingress_flow_count": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The maximum number of ingress flows that can be created in the Message VPN. The default value is 1000.",
				Optional:    true,
				Default:     1000,
			},
			"max_spool_usage": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The maximum Message Spool usage by the Message VPN, in megabytes. The default value is 0.",
				Optional:    true,
				Default:     0,
			},
			"max_subscription_count": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The maximum number of local client subscriptions (both primary and backup) that can be added to the Message VPN. The default is 500000.",
				Optional:    true,
				Default:     500000,
			},
			"max_transacted_session_count": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The maximum number of transacted sessions for the Message VPN. The default varies by platform. The default is 1000.",
				Optional:    true,
				Default:     1000,
			},
			"max_transaction_count": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The maximum number of transactions for the Message VPN. The default varies by platform. The default is 5000.",
				Optional:    true,
				Default:     5000,
			},
			"replication_enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Enable or disable the Replication feature for the Message VPN. The default value is false.",
				Optional:    true,
				Default:     false,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceMsgVpnCreate(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Creating msg vpn ...")

	// Get our Solace client
	c := m.(*Config)
	client := c.Client
	auth := c.Auth

	// Extract config data from resource data and prepare new VPN object
	name := d.Get("name").(string)
	newMsgVpn := models.MsgVpn{
		MsgVpnName:                 name,
		AuthenticationBasicEnabled: d.Get("authentication_basic_enabled").(bool),
		Enabled:                    d.Get("enabled").(bool),
		MaxConnectionCount:         int64(d.Get("max_connection_count").(int)),
		MaxEgressFlowCount:         int64(d.Get("max_egress_flow_count").(int)),
		MaxEndpointCount:           int64(d.Get("max_endpoint_count").(int)),
		MaxIngressFlowCount:        int64(d.Get("max_ingress_flow_count").(int)),
		MaxMsgSpoolUsage:           int64(d.Get("max_spool_usage").(int)),
		MaxSubscriptionCount:       int64(d.Get("max_subscription_count").(int)),
		MaxTransactedSessionCount:  int64(d.Get("max_transacted_session_count").(int)),
		MaxTransactionCount:        int64(d.Get("max_transaction_count").(int)),
		ReplicationEnabled:         d.Get("replication_enabled").(bool),
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
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
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
	d.Set("max_connection_count", resp.Payload.Data.MaxConnectionCount)
	d.Set("max_egress_flow_count", resp.Payload.Data.MaxEgressFlowCount)
	d.Set("max_endpoint_count", resp.Payload.Data.MaxEndpointCount)
	d.Set("max_ingress_flow_count", resp.Payload.Data.MaxIngressFlowCount)
	d.Set("max_spool_usage", resp.Payload.Data.MaxMsgSpoolUsage)
	d.Set("max_subscription_count", resp.Payload.Data.MaxSubscriptionCount)
	d.Set("max_transacted_session_count", resp.Payload.Data.MaxTransactedSessionCount)
	d.Set("max_transaction_count", resp.Payload.Data.MaxTransactionCount)
	d.Set("replication_enabled", resp.Payload.Data.ReplicationEnabled)
	d.Set("authentication_basic_enabled", resp.Payload.Data.AuthenticationBasicEnabled)
	return nil
}

func resourceMsgVpnUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*Config)
	client := c.Client
	auth := c.Auth

	params := msg_vpn.NewUpdateMsgVpnParams()
	params.MsgVpnName = d.Id()
	newMsgVpn := models.MsgVpn{}

	// Only include changed values; anything we don't specify does not get updated
	if d.HasChange("authentication_basic_enabled") {
		newMsgVpn.AuthenticationBasicEnabled = d.Get("authentication_basic_enabled").(bool)
	}
	if d.HasChange("enabled") {
		newMsgVpn.Enabled = d.Get("enabled").(bool)
	}
	if d.HasChange("max_connection_count") {
		newMsgVpn.MaxConnectionCount = int64(d.Get("max_connection_count").(int))
	}
	if d.HasChange("max_egress_flow_count") {
		newMsgVpn.MaxEgressFlowCount = int64(d.Get("max_egress_flow_count").(int))
	}
	if d.HasChange("max_endpoint_count") {
		newMsgVpn.MaxEndpointCount = int64(d.Get("max_endpoint_count").(int))
	}
	if d.HasChange("max_ingress_flow_count") {
		newMsgVpn.MaxIngressFlowCount = int64(d.Get("max_ingress_flow_count").(int))
	}
	if d.HasChange("max_spool_usage") {
		newMsgVpn.MaxMsgSpoolUsage = int64(d.Get("max_spool_usage").(int))
	}
	if d.HasChange("max_subscription_count") {
		newMsgVpn.MaxSubscriptionCount = int64(d.Get("max_subscription_count").(int))
	}
	if d.HasChange("max_transacted_session_count") {
		newMsgVpn.MaxTransactedSessionCount = int64(d.Get("max_transacted_session_count").(int))
	}
	if d.HasChange("max_transaction_count") {
		newMsgVpn.MaxTransactionCount = int64(d.Get("max_transaction_count").(int))
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
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
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
