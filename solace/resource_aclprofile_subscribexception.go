package solace

import (
	"fmt"
	"log"
	"strings"

	"github.com/ExalDraen/semp-client/models"

	"github.com/ExalDraen/semp-client/client/operations"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceACLSubscribeException() *schema.Resource {
	return &schema.Resource{
		Create: resourceACLSubscribeExceptionCreate,
		Read:   resourceACLSubscribeExceptionRead,
		Delete: resourceACLSubscribeExceptionDelete,

		Schema: map[string]*schema.Schema{
			"topic": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the Topic for the Exception to the default action taken. May include wildcard characters.",
				Required:    true,
				ForceNew:    true,
			},
			"topic_syntax": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The syntax of the Topic for the Exception to the default action taken. The allowed values are \"smf\" and \"mqtt\"",
				Required:    true,
				ForceNew:    true,
			},
			"acl": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the ACL Profile.",
				Required:    true,
				ForceNew:    true,
			},
			// Each ACL client conn exception must belong to a VPN, but optionally we use the provider set default,
			// and bail if neither is set. Thus the parameter is optional.
			"msg_vpn": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The name of the MSG VPN. If unset the provider default is used.",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: resourceACLSubscribeExceptionImport,
		},
	}
}

func resourceACLSubscribeExceptionCreate(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Creating ACL Profile subscribe exception...")

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

	SubExc := models.MsgVpnACLProfileSubscribeException{
		TopicSyntax:             syntax,
		SubscribeExceptionTopic: topic,
		ACLProfileName:          acl,
		MsgVpnName:              vpn,
	}

	params := operations.NewCreateMsgVpnACLProfileSubscribeExceptionParams()
	params.MsgVpnName = vpn
	params.ACLProfileName = acl
	params.Body = &SubExc

	resp, err := client.Operations.CreateMsgVpnACLProfileSubscribeException(params, auth)
	if err != nil {
		sempErr := err.(*operations.CreateMsgVpnACLProfileSubscribeExceptionDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create ACL Profile subscribe exception %q for ACL %q on VPN %q: %q", SubExc, acl, vpn, formatError(sempErr))
	}
	d.SetId(resp.Payload.Data.SubscribeExceptionTopic)

	log.Printf("[DEBUG] Finished creating ACL Profile subscribe exception %q on ACL %q on VPN %q", SubExc, acl, vpn)
	return resourceACLSubscribeExceptionRead(d, m)
}

func resourceACLSubscribeExceptionRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Reading ACL Profile subscribe exception %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := operations.NewGetMsgVpnACLProfileSubscribeExceptionParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.SubscribeExceptionTopic = d.Id()
	params.TopicSyntax = d.Get("topic_syntax").(string)
	params.ACLProfileName = d.Get("acl").(string)
	params.MsgVpnName = vpn

	resp, err := client.Operations.GetMsgVpnACLProfileSubscribeException(params, auth)
	if err != nil {
		log.Printf("[WARN] No ACL Profile subscribe exception found found: %q", d.Id())
		d.SetId("")
		return nil
	}
	log.Printf("%#v\n", resp.Payload.Data)
	d.Set("topic", resp.Payload.Data.SubscribeExceptionTopic)
	d.Set("topic_syntax", resp.Payload.Data.TopicSyntax)
	d.Set("acl", resp.Payload.Data.ACLProfileName)
	d.Set("msg_vpn", resp.Payload.Data.MsgVpnName)

	return nil
}

func resourceACLSubscribeExceptionDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Deleting ACL Profile subscribe exception %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := operations.NewDeleteMsgVpnACLProfileSubscribeExceptionParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}
	params.SubscribeExceptionTopic = d.Id()
	params.TopicSyntax = d.Get("topic_syntax").(string)
	params.ACLProfileName = d.Get("acl").(string)
	params.MsgVpnName = vpn

	_, err = client.Operations.DeleteMsgVpnACLProfileSubscribeException(params, auth)
	if err != nil {
		sempErr := err.(*operations.DeleteMsgVpnACLProfileSubscribeExceptionDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to ACL Profile subscribe exception exception %q for ACL %q on VPN %q: %q",
			params.TopicSyntax, params.ACLProfileName, params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceACLSubscribeExceptionImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "|")
	if len(idParts) != 4 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" || idParts[3] == "" {
		return nil, fmt.Errorf("Unexpected format of ID (%q), expected MSG-VPN|ACL-PROFILE|TOPIC-SYNTAX|SUB-EXC-TOPIC", d.Id())
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
