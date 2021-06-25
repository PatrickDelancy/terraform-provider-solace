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

func resourceACLSubscribeShareNameException() *schema.Resource {
	return &schema.Resource{
		Create: resourceACLSubscribeShareNameExceptionCreate,
		Read:   resourceACLSubscribeShareNameExceptionRead,
		Delete: resourceACLSubscribeShareNameExceptionDelete,

		Schema: map[string]*schema.Schema{
			"topic": {
				Type:        schema.TypeString,
				Description: "The subscribe share name exception to the default action taken. May include wildcard characters.",
				Required:    true,
				ForceNew:    true,
			},
			"topic_syntax": {
				Type:         schema.TypeString,
				Description:  "The syntax of the subscribe share name for the exception to the default action taken. The allowed values are \"smf\" and \"mqtt\"",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"smf", "mqtt"}, false),
			},
			"acl": {
				Type:        schema.TypeString,
				Description: "The name of the ACL Profile.",
				Required:    true,
				ForceNew:    true,
			},
			// Each ACL client conn exception must belong to a VPN, but optionally we use the provider set default,
			// and bail if neither is set. Thus the parameter is optional.
			"msg_vpn": {
				Type:        schema.TypeString,
				Description: "The name of the MSG VPN. If unset the provider default is used.",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: resourceACLSubscribeShareNameExceptionImport,
		},
	}
}

func resourceACLSubscribeShareNameExceptionCreate(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Creating ACL Profile subscribe share name exception...")

	// Get our Solace client
	c := m.(*Config)
	client := c.Client
	auth := c.Auth

	// Extract config data from resource data and prepare new VPN object
	topic := d.Get("topic").(string)
	syntax := d.Get("topic_syntax").(string)
	acl := d.Get("acl").(string)
	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	SubExc := models.MsgVpnACLProfileSubscribeShareNameException{
		SubscribeShareNameExceptionSyntax: syntax,
		SubscribeShareNameException:       topic,
		ACLProfileName:                    acl,
		MsgVpnName:                        vpn,
	}

	params := all.NewCreateMsgVpnACLProfileSubscribeShareNameExceptionParams()
	params.MsgVpnName = vpn
	params.ACLProfileName = acl
	params.Body = &SubExc

	resp, err := client.All.CreateMsgVpnACLProfileSubscribeShareNameException(params, auth)
	if err != nil {
		sempErr := err.(*all.CreateMsgVpnACLProfileSubscribeShareNameExceptionDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create ACL Profile subscribe share name exception %q for ACL %q on VPN %q: %q", SubExc, acl, vpn, formatError(sempErr))
	}
	d.SetId(resp.Payload.Data.SubscribeShareNameException)

	log.Printf("[DEBUG] Finished creating ACL Profile subscribe share name exception %q on ACL %q on VPN %q", SubExc, acl, vpn)
	return resourceACLSubscribeShareNameExceptionRead(d, m)
}

func resourceACLSubscribeShareNameExceptionRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Reading ACL Profile subscribe share name exception %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewGetMsgVpnACLProfileSubscribeShareNameExceptionParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.SubscribeShareNameException = d.Id()
	params.SubscribeShareNameExceptionSyntax = d.Get("topic_syntax").(string)
	params.ACLProfileName = d.Get("acl").(string)
	params.MsgVpnName = vpn

	resp, err := client.All.GetMsgVpnACLProfileSubscribeShareNameException(params, auth)
	if err != nil {
		log.Printf("[WARN] No ACL Profile subscribe share name exception found found: %q", d.Id())
		d.SetId("")
		return nil
	}
	log.Printf("%#v\n", resp.Payload.Data)
	d.Set("topic", resp.Payload.Data.SubscribeShareNameException)
	d.Set("topic_syntax", resp.Payload.Data.SubscribeShareNameExceptionSyntax)
	d.Set("acl", resp.Payload.Data.ACLProfileName)
	d.Set("msg_vpn", resp.Payload.Data.MsgVpnName)

	return nil
}

func resourceACLSubscribeShareNameExceptionDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Deleting ACL Profile subscribe share name exception %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewDeleteMsgVpnACLProfileSubscribeShareNameExceptionParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}
	params.SubscribeShareNameException = d.Id()
	params.SubscribeShareNameExceptionSyntax = d.Get("topic_syntax").(string)
	params.ACLProfileName = d.Get("acl").(string)
	params.MsgVpnName = vpn

	_, err = client.All.DeleteMsgVpnACLProfileSubscribeShareNameException(params, auth)
	if err != nil {
		sempErr := err.(*all.DeleteMsgVpnACLProfileSubscribeShareNameExceptionDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to ACL Profile subscribe share name exception exception %q for ACL %q on VPN %q: %q",
			params.SubscribeShareNameExceptionSyntax, params.ACLProfileName, params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceACLSubscribeShareNameExceptionImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "|")
	if len(idParts) != 4 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" || idParts[3] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected MSG-VPN|ACL-PROFILE|TOPIC-SYNTAX|SUB-EXC-TOPIC", d.Id())
	}
	vpn := idParts[0]
	acl := idParts[1]
	syntax := idParts[2]
	SubExc := idParts[3]
	d.Set("msg_vpn", vpn)
	d.Set("acl", acl)
	d.Set("topic_syntax", syntax)
	d.SetId(SubExc)
	return []*schema.ResourceData{d}, nil
}
