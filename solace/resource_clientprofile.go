package solace

import (
	"fmt"
	"log"
	"strings"

	"github.com/ExalDraen/semp-client/models"

	"github.com/ExalDraen/semp-client/client/operations"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceClientProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceClientProfileCreate,
		Read:   resourceClientProfileRead,
		Update: resourceClientProfileUpdate,
		Delete: resourceClientProfileDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the Client profile. Used as a unique identifier.",
				Required:    true,
				ForceNew:    true,
			},
			// Each Client profile must belong to a VPN, but optionally we use the provider set default,
			// and bail if neither is set. Thus the parameter is optional.
			"msg_vpn": {
				Type:        schema.TypeString,
				Description: "The name of the MSG VPN. If unset the provider default is used.",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"allow_cut_through_forwarding": {
				Type:        schema.TypeBool,
				Description: "Allow or deny clients to bind to topic endpoints or queues with the cut-through delivery mode. Changing this setting does not affect existing client connections",
				Optional:    true,
			},
			"allow_guaranteed_endpoint_create": {
				Type:        schema.TypeBool,
				Description: "Allow or deny clients to create topic endponts or queues. Changing this setting does not affect existing client connections.",
				Optional:    true,
			},
			"allow_guaranteed_msg_receive": {
				Type:        schema.TypeBool,
				Description: "Allow or deny clients to receive guaranteed messages. Changing this setting does not affect existing client connections",
				Optional:    true,
			},
			"allow_guaranteed_msg_send": {
				Type:        schema.TypeBool,
				Description: "Allow or deny clients to send guaranteed messages. Changing this setting does not affect existing client connections.",
				Optional:    true,
			},
			"allow_transacted_sessions": {
				Type:        schema.TypeBool,
				Description: "Allow or deny clients to establish transacted sessions. Changing this setting does not affect existing client connections.",
				Optional:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: resourceClientProfileImport,
		},
	}
}

func resourceClientProfileCreate(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Creating Client profile ...")

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

	acl := models.MsgVpnClientProfile{
		ClientProfileName: name,
		MsgVpnName:        vpn,
	}
	// Only set these if they're actually set (not their default value)
	if v, ok := d.GetOk("allow_cut_through_forwarding"); ok {
		val := v.(bool)
		acl.AllowCutThroughForwardingEnabled = &val
	}
	if v, ok := d.GetOk("allow_guaranteed_endpoint_create"); ok {
		val := v.(bool)
		acl.AllowGuaranteedEndpointCreateEnabled = &val
	}
	if v, ok := d.GetOk("allow_guaranteed_msg_receive"); ok {
		val := v.(bool)
		acl.AllowGuaranteedMsgReceiveEnabled = &val
	}
	if v, ok := d.GetOk("allow_guaranteed_msg_send"); ok {
		val := v.(bool)
		acl.AllowGuaranteedMsgSendEnabled = &val
	}
	if v, ok := d.GetOk("allow_transacted_sessions"); ok {
		val := v.(bool)
		acl.AllowTransactedSessionsEnabled = &val
	}

	params := operations.NewCreateMsgVpnClientProfileParams()
	params.MsgVpnName = vpn
	params.Body = &acl

	resp, err := client.Operations.CreateMsgVpnClientProfile(params, auth)
	if err != nil {
		sempErr := err.(*operations.CreateMsgVpnClientProfileDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create Client profile %q on vpn %q: %v", name, vpn, formatError(sempErr))
	}
	d.SetId(resp.Payload.Data.ClientProfileName)

	log.Printf("[DEBUG] Finished creating ACL %q on VPN %q", name, vpn)
	return resourceClientProfileRead(d, m)
}

func resourceClientProfileRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Reading Client profile %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := operations.NewGetMsgVpnClientProfileParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.ClientProfileName = d.Id()
	params.MsgVpnName = vpn

	resp, err := client.Operations.GetMsgVpnClientProfile(params, auth)
	if err != nil {
		log.Printf("[WARN] No Client profile found: %s", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", resp.Payload.Data.ClientProfileName)
	d.Set("msg_vpn", resp.Payload.Data.MsgVpnName)
	d.Set("allow_cut_through_forwarding", resp.Payload.Data.AllowCutThroughForwardingEnabled)
	d.Set("allow_guaranteed_endpoint_create", resp.Payload.Data.AllowGuaranteedEndpointCreateEnabled)
	d.Set("allow_guaranteed_msg_receive", resp.Payload.Data.AllowGuaranteedMsgReceiveEnabled)
	d.Set("allow_guaranteed_msg_send", resp.Payload.Data.AllowGuaranteedMsgSendEnabled)
	d.Set("allow_transacted_sessions", resp.Payload.Data.AllowTransactedSessionsEnabled)

	return nil
}

func resourceClientProfileUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Updating Client profile %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := operations.NewUpdateMsgVpnClientProfileParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.ClientProfileName = d.Id()
	params.MsgVpnName = vpn
	acl := models.MsgVpnClientProfile{}
	acl.MsgVpnName = vpn

	// Only include changed values; anything we don't specify does not get updated
	if d.HasChange("allow_cut_through_forwarding") {
		val := d.Get("allow_cut_through_forwarding").(bool)
		acl.AllowCutThroughForwardingEnabled = &val
	}
	if d.HasChange("allow_guaranteed_endpoint_create") {
		val := d.Get("allow_guaranteed_endpoint_create").(bool)
		acl.AllowGuaranteedEndpointCreateEnabled = &val
	}
	if d.HasChange("allow_guaranteed_msg_receive") {
		val := d.Get("allow_guaranteed_msg_receive").(bool)
		acl.AllowGuaranteedMsgReceiveEnabled = &val
	}
	if d.HasChange("allow_guaranteed_msg_send") {
		val := d.Get("allow_guaranteed_msg_send").(bool)
		acl.AllowGuaranteedMsgSendEnabled = &val
	}
	if d.HasChange("allow_transacted_sessions") {
		val := d.Get("allow_transacted_sessions").(bool)
		acl.AllowTransactedSessionsEnabled = &val
	}
	params.Body = &acl

	_, err = client.Operations.UpdateMsgVpnClientProfile(params, auth)
	if err != nil {
		sempErr := err.(*operations.UpdateMsgVpnClientProfileDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to update Client profile %q: %v", params.ClientProfileName, formatError(sempErr))
	}

	return resourceClientProfileRead(d, m)
}

func resourceClientProfileDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := operations.NewDeleteMsgVpnClientProfileParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}
	params.ClientProfileName = d.Id()
	params.MsgVpnName = vpn

	_, err = client.Operations.DeleteMsgVpnClientProfile(params, auth)
	if err != nil {
		sempErr := err.(*operations.DeleteMsgVpnClientProfileDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to delete Client profile %q: %v", params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceClientProfileImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "/")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected MSG-VPN/CLIENT-PROFILE", d.Id())
	}
	vpn := idParts[0]
	acl := idParts[1]
	d.Set("msg_vpn", vpn)

	d.SetId(acl)
	return []*schema.ResourceData{d}, nil
}
