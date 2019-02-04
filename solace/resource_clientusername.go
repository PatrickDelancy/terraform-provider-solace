package solace

import (
	"fmt"
	"log"
	"strings"

	"github.com/ExalDraen/semp-client/models"

	"github.com/ExalDraen/semp-client/client/operations"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceClientUsername() *schema.Resource {
	return &schema.Resource{
		Create: resourceClientUsernameCreate,
		Read:   resourceClientUsernameRead,
		Update: resourceClientUsernameUpdate,
		Delete: resourceClientUsernameDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the Client. Used as a unique identifier.",
				Required:    true,
				ForceNew:    true,
			},
			// Each Client username must belong to a VPN, but optionally we use the provider set default,
			// and bail if neither is set. Thus the parameter is optional.
			"msg_vpn": {
				Type:        schema.TypeString,
				Description: "The name of the MSG VPN. If unset the provider default is used.",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Description: "Enables or disables the Client Username. When disabled all clients currently connected as the Client Username are disconnected.",
				Optional:    true,
			},
			"acl": {
				Type:        schema.TypeString,
				Description: "The ACL Profile of the Client Username. The default value is \"default\".",
				Optional:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "The password of this Client Username for internal Authentication. The default is to have no password.",
				Optional:    true,
			},
			"profile": {
				Type:        schema.TypeString,
				Description: "The Client Profile of the Client Username. The default value is \"default\".",
				Optional:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: resourceClientUsernameImport,
		},
	}
}

func resourceClientUsernameCreate(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Creating Client username ...")

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

	user := models.MsgVpnClientUsername{
		ClientUsername: name,
		MsgVpnName:     vpn,
	}
	// Only set these if they're actually set (not their default value)
	if v, ok := d.GetOk("enabled"); ok {
		val := v.(bool)
		user.Enabled = &val
	}
	if v, ok := d.GetOk("acl"); ok {
		user.ACLProfileName = v.(string)
	}
	if v, ok := d.GetOk("password"); ok {
		val := v.(string)
		user.Password = &val
	}
	if v, ok := d.GetOk("profile"); ok {
		user.ClientProfileName = v.(string)
	}

	params := operations.NewCreateMsgVpnClientUsernameParams()
	params.MsgVpnName = vpn
	params.Body = &user

	resp, err := client.Operations.CreateMsgVpnClientUsername(params, auth)
	if err != nil {
		sempErr := err.(*operations.CreateMsgVpnClientUsernameDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create Client username %q on vpn %q: %v", name, vpn, formatError(sempErr))
	}
	d.SetId(resp.Payload.Data.ClientUsername)

	log.Printf("[DEBUG] Finished creating Client username %q on VPN %q", name, vpn)
	return resourceClientUsernameRead(d, m)
}

func resourceClientUsernameRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Reading Client username %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := operations.NewGetMsgVpnClientUsernameParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.ClientUsername = d.Id()
	params.MsgVpnName = vpn

	resp, err := client.Operations.GetMsgVpnClientUsername(params, auth)
	if err != nil {
		log.Printf("[WARN] No Client username found: %s", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", resp.Payload.Data.ClientUsername)
	d.Set("msg_vpn", resp.Payload.Data.MsgVpnName)
	d.Set("enabled", resp.Payload.Data.Enabled)
	d.Set("acl", resp.Payload.Data.ACLProfileName)
	d.Set("password", resp.Payload.Data.Password)
	d.Set("profile", resp.Payload.Data.ClientProfileName)

	return nil
}

func resourceClientUsernameUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Updating Client username %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := operations.NewUpdateMsgVpnClientUsernameParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.ClientUsername = d.Id()
	params.MsgVpnName = vpn
	user := models.MsgVpnClientUsername{}
	user.MsgVpnName = vpn

	// Only include changed values; anything we don't specify does not get updated
	if d.HasChange("enabled") {
		val := d.Get("enabled").(bool)
		user.Enabled = &val
	}
	if d.HasChange("acl") {
		user.ACLProfileName = d.Get("acl").(string)
	}
	if d.HasChange("password") {
		val := d.Get("password").(string)
		user.Password = &val
	}
	if d.HasChange("profile") {
		user.ClientProfileName = d.Get("profile").(string)
	}

	params.Body = &user

	_, err = client.Operations.UpdateMsgVpnClientUsername(params, auth)
	if err != nil {
		sempErr := err.(*operations.UpdateMsgVpnClientUsernameDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to update Client username %q: %v", params.ClientUsername, formatError(sempErr))
	}

	return resourceClientUsernameRead(d, m)
}

func resourceClientUsernameDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := operations.NewDeleteMsgVpnClientUsernameParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}
	params.ClientUsername = d.Id()
	params.MsgVpnName = vpn

	_, err = client.Operations.DeleteMsgVpnClientUsername(params, auth)
	if err != nil {
		sempErr := err.(*operations.DeleteMsgVpnClientUsernameDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to delete Client username %q: %v", params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceClientUsernameImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "/")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected MSG-VPN/CLIENT-USERNAME", d.Id())
	}
	vpn := idParts[0]
	user := idParts[1]
	d.Set("msg_vpn", vpn)

	d.SetId(user)
	return []*schema.ResourceData{d}, nil
}
