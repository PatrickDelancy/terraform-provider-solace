package solace

import (
	"fmt"
	"log"

	"github.com/PatrickDelancy/semp-client/client/all"
	"github.com/PatrickDelancy/semp-client/models"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMsgVpn() *schema.Resource {
	return &schema.Resource{
		Create: resourceMsgVpnCreate,
		Read:   resourceMsgVpnRead,
		Update: resourceMsgVpnUpdate,
		Delete: resourceMsgVpnDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the MSG VPN. Used as an identifier.",
				Required:    true,
				ForceNew:    true,
			},
			"authentication_basic_enabled": {
				Type:        schema.TypeBool,
				Description: "Enable or disable Basic Authentication for clients connecting to the Message VPN. The default value is true.",
				Optional:    true,
				Default:     true,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Description: "Whether or not the MSG VPN should be enabled.",
				Optional:    true,
				Default:     true,
			},
			"max_connection_count": {
				Type:        schema.TypeInt,
				Description: "The maximum number of client connections that can be simultaneously connected to the Message VPN. This value may be higher than supported by the hardware. The default is 100.",
				Optional:    true,
				Default:     100,
			},
			"max_egress_flow_count": {
				Type:        schema.TypeInt,
				Description: "The maximum number of egress flows that can be created in the Message VPN. The default value is 1000.",
				Optional:    true,
				Default:     1000,
			},
			"max_endpoint_count": {
				Type:        schema.TypeInt,
				Description: "The maximum number of Queues and Topic Endpoints that can be created in the Message VPN. The default value is 1000.",
				Optional:    true,
				Default:     1000,
			},
			"max_ingress_flow_count": {
				Type:        schema.TypeInt,
				Description: "The maximum number of ingress flows that can be created in the Message VPN. The default value is 1000.",
				Optional:    true,
				Default:     1000,
			},
			"max_spool_usage": {
				Type:        schema.TypeInt,
				Description: "The maximum Message Spool usage by the Message VPN, in megabytes. The default value is 0.",
				Optional:    true,
				Default:     0,
			},
			"max_subscription_count": {
				Type:        schema.TypeInt,
				Description: "The maximum number of local client subscriptions (both primary and backup) that can be added to the Message VPN. The default is 500000.",
				Optional:    true,
				Default:     500000,
			},
			"max_transacted_session_count": {
				Type:        schema.TypeInt,
				Description: "The maximum number of transacted sessions for the Message VPN. The default varies by platform. The default is 1000.",
				Optional:    true,
				Default:     1000,
			},
			"max_transaction_count": {
				Type:        schema.TypeInt,
				Description: "The maximum number of transactions for the Message VPN. The default varies by platform. The default is 5000.",
				Optional:    true,
				Default:     5000,
			},
			"replication_enabled": {
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

	vpn := models.MsgVpn{
		MsgVpnName: name,
	}
	// Only set these if they're actually set (not their default value)
	if v, ok := d.GetOk("authentication_basic_enabled"); ok {
		val := v.(bool)
		vpn.AuthenticationBasicEnabled = val
	}
	if v, ok := d.GetOk("enabled"); ok {
		val := v.(bool)
		vpn.Enabled = val
	}
	if v, ok := d.GetOk("replication_enabled"); ok {
		val := v.(bool)
		vpn.ReplicationEnabled = val
	}
	if v, ok := d.GetOk("max_connection_count"); ok {
		val := int64(v.(int))
		vpn.MaxConnectionCount = val
	}
	if v, ok := d.GetOk("max_egress_flow_count"); ok {
		val := int64(v.(int))
		vpn.MaxEgressFlowCount = val
	}
	if v, ok := d.GetOk("max_endpoint_count"); ok {
		val := int64(v.(int))
		vpn.MaxEndpointCount = val
	}
	if v, ok := d.GetOk("max_ingress_flow_count"); ok {
		val := int64(v.(int))
		vpn.MaxIngressFlowCount = val
	}
	if v, ok := d.GetOk("max_spool_usage"); ok {
		val := int64(v.(int))
		vpn.MaxMsgSpoolUsage = val
	}
	if v, ok := d.GetOk("max_subscription_count"); ok {
		val := int64(v.(int))
		vpn.MaxSubscriptionCount = val
	}
	if v, ok := d.GetOk("max_transacted_session_count"); ok {
		val := int64(v.(int))
		vpn.MaxTransactedSessionCount = val
	}
	if v, ok := d.GetOk("max_transaction_count"); ok {
		val := int64(v.(int))
		vpn.MaxTransactionCount = val
	}

	params := all.NewCreateMsgVpnParams()
	params.Body = &vpn

	resp, err := client.All.CreateMsgVpn(params, auth)
	if err != nil {
		sempErr := err.(*all.CreateMsgVpnDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create message VPN %q: %v", name, formatError(sempErr))
	}
	d.SetId(resp.Payload.Data.MsgVpnName)

	log.Printf("[DEBUG] Finished creating message VPN %q", name)
	return resourceMsgVpnRead(d, m)
}

func resourceMsgVpnRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Reading msg vpn %q...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	getParams := all.NewGetMsgVpnParams()
	getParams.MsgVpnName = d.Id()

	resp, err := client.All.GetMsgVpn(getParams, auth)
	if err != nil {
		log.Printf("[WARN] No Server found: %s", d.Id())
		d.SetId("")
		return nil
	}

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
	log.Printf("[DEBUG] Updateing msg VPN %v", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth

	params := all.NewUpdateMsgVpnParams()
	params.MsgVpnName = d.Id()
	newMsgVpn := models.MsgVpn{}

	// Only include changed values; anything we don't specify does not get updated
	if d.HasChange("authentication_basic_enabled") {
		val := d.Get("authentication_basic_enabled").(bool)
		newMsgVpn.AuthenticationBasicEnabled = val
	}
	if d.HasChange("enabled") {
		val := d.Get("enabled").(bool)
		newMsgVpn.Enabled = val
	}
	if d.HasChange("max_connection_count") {
		val := int64(d.Get("max_connection_count").(int))
		newMsgVpn.MaxConnectionCount = val
	}
	if d.HasChange("max_egress_flow_count") {
		val := int64(d.Get("max_egress_flow_count").(int))
		newMsgVpn.MaxEgressFlowCount = val
	}
	if d.HasChange("max_endpoint_count") {
		val := int64(d.Get("max_endpoint_count").(int))
		newMsgVpn.MaxEndpointCount = val
	}
	if d.HasChange("max_ingress_flow_count") {
		val := int64(d.Get("max_ingress_flow_count").(int))
		newMsgVpn.MaxIngressFlowCount = val
	}
	if d.HasChange("max_spool_usage") {
		val := int64(d.Get("max_spool_usage").(int))
		newMsgVpn.MaxMsgSpoolUsage = val
	}
	if d.HasChange("max_subscription_count") {
		val := int64(d.Get("max_subscription_count").(int))
		newMsgVpn.MaxSubscriptionCount = val
	}
	if d.HasChange("max_transacted_session_count") {
		val := int64(d.Get("max_transacted_session_count").(int))
		newMsgVpn.MaxTransactedSessionCount = val
	}
	if d.HasChange("max_transaction_count") {
		val := int64(d.Get("max_transaction_count").(int))
		newMsgVpn.MaxTransactionCount = val
	}
	if d.HasChange("replication_enabled") {
		val := d.Get("replication_enabled").(bool)
		newMsgVpn.ReplicationEnabled = val
	}
	params.Body = &newMsgVpn

	log.Printf("[TRACE] msg VPN params are %+v", params)
	log.Printf("[TRACE] msg VPN Body: %+v", params.Body)
	_, err := client.All.UpdateMsgVpn(params, auth)
	if err != nil {
		sempErr := err.(*all.UpdateMsgVpnDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to update message VPN %q: %v", params.MsgVpnName, formatError(sempErr))
	}

	return resourceMsgVpnRead(d, m)
}

func resourceMsgVpnDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewDeleteMsgVpnParams()
	params.MsgVpnName = d.Id()

	_, err := client.All.DeleteMsgVpn(params, auth)
	if err != nil {
		sempErr := err.(*all.DeleteMsgVpnDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to delete message VPN %q: %v", params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}
