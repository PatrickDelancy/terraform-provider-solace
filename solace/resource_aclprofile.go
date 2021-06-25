package solace

import (
	"fmt"
	"log"
	"strings"

	"github.com/PatrickDelancy/semp-client/client/all"
	"github.com/PatrickDelancy/semp-client/models"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceACLProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceACLProfileCreate,
		Read:   resourceACLProfileRead,
		Update: resourceACLProfileUpdate,
		Delete: resourceACLProfileDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the ACL profile. Used as a unique identifier.",
				Required:    true,
				ForceNew:    true,
			},
			// Each ACL must belong to a VPN, but optionally we use the provider set default,
			// and bail if neither is set. Thus the parameter is optional.
			"msg_vpn": {
				Type:        schema.TypeString,
				Description: "The name of the MSG VPN. If unset the provider default is used.",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"client_connection_default_action": {
				Type:         schema.TypeString,
				Description:  "The default action when a Client connects to the Message VPN. Must be one of \"allow\" or \"disallow\"",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"allow", "disallow"}, false),
			},
			"publish_topic_default_action": {
				Type:         schema.TypeString,
				Description:  "The default action to take when a Client publishes to a Topic in the Message VPN. Must be one of \"allow\" or \"disallow\"",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"allow", "disallow"}, false),
			},
			"subscribe_topic_default_action": {
				Type:         schema.TypeString,
				Description:  "The default action to take when a Client subscribes to a Topic in the Message VPN. Must be one of \"allow\" or \"disallow\"",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"allow", "disallow"}, false),
			},
			"subscribe_sharename_default_action": {
				Type:         schema.TypeString,
				Description:  "The default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. Must be one of \"allow\" or \"disallow\"",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"allow", "disallow"}, false),
			},
		},
		Importer: &schema.ResourceImporter{
			State: resourceACLProfileImport,
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
	// Only set these if they're actually set (not their default value)
	if v, ok := d.GetOk("client_connection_default_action"); ok {
		acl.ClientConnectDefaultAction = v.(string)
	}
	if v, ok := d.GetOk("publish_topic_default_action"); ok {
		acl.PublishTopicDefaultAction = v.(string)
	}
	if v, ok := d.GetOk("subscribe_topic_default_action"); ok {
		acl.SubscribeTopicDefaultAction = v.(string)
	}
	if v, ok := d.GetOk("subscribe_sharename_default_action"); ok {
		acl.SubscribeShareNameDefaultAction = v.(string)
	}

	params := all.NewCreateMsgVpnACLProfileParams()
	params.MsgVpnName = vpn
	params.Body = &acl

	resp, err := client.All.CreateMsgVpnACLProfile(params, auth)
	if err != nil {
		sempErr := err.(*all.CreateMsgVpnACLProfileDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create ACL profile %q on vpn %q: %v", name, vpn, formatError(sempErr))
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
	params := all.NewGetMsgVpnACLProfileParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.ACLProfileName = d.Id()
	params.MsgVpnName = vpn

	resp, err := client.All.GetMsgVpnACLProfile(params, auth)
	if err != nil {
		log.Printf("[WARN] No ACL profile found: %s", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", resp.Payload.Data.ACLProfileName)
	d.Set("msg_vpn", resp.Payload.Data.MsgVpnName)
	d.Set("client_connection_default_action", resp.Payload.Data.ClientConnectDefaultAction)
	d.Set("publish_topic_default_action", resp.Payload.Data.PublishTopicDefaultAction)
	d.Set("subscribe_topic_default_action", resp.Payload.Data.SubscribeTopicDefaultAction)
	d.Set("subscribe_sharename_default_action", resp.Payload.Data.SubscribeShareNameDefaultAction)

	return nil
}

func resourceACLProfileUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Updating ACL Profile %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewUpdateMsgVpnACLProfileParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.ACLProfileName = d.Id()
	params.MsgVpnName = vpn
	acl := models.MsgVpnACLProfile{}
	acl.MsgVpnName = vpn

	// Only include changed values; anything we don't specify does not get updated
	if d.HasChange("client_connection_default_action") {
		acl.ClientConnectDefaultAction = d.Get("client_connection_default_action").(string)
	}
	if d.HasChange("publish_topic_default_action") {
		acl.PublishTopicDefaultAction = d.Get("publish_topic_default_action").(string)
	}
	if d.HasChange("subscribe_topic_default_action") {
		acl.SubscribeTopicDefaultAction = d.Get("subscribe_topic_default_action").(string)
	}
	if d.HasChange("subscribe_sharename_default_action") {
		acl.SubscribeShareNameDefaultAction = d.Get("subscribe_sharename_default_action").(string)
	}
	params.Body = &acl

	_, err = client.All.UpdateMsgVpnACLProfile(params, auth)
	if err != nil {
		sempErr := err.(*all.UpdateMsgVpnACLProfileDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to update ACL %q: %v", params.ACLProfileName, formatError(sempErr))
	}

	return resourceACLProfileRead(d, m)
}

func resourceACLProfileDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewDeleteMsgVpnACLProfileParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}
	params.ACLProfileName = d.Id()
	params.MsgVpnName = vpn

	_, err = client.All.DeleteMsgVpnACLProfile(params, auth)
	if err != nil {
		sempErr := err.(*all.DeleteMsgVpnACLProfileDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to delete ACL %q: %v", params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceACLProfileImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "/")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected MSG-VPN/ACL-PROFILE", d.Id())
	}
	vpn := idParts[0]
	acl := idParts[1]
	d.Set("msg_vpn", vpn)

	d.SetId(acl)
	return []*schema.ResourceData{d}, nil
}
