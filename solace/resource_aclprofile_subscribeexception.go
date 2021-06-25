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

func resourceACLSubscribeException() *schema.Resource {
	return &schema.Resource{
		Create: resourceACLSubscribeExceptionCreate,
		Read:   resourceACLSubscribeExceptionRead,
		Delete: resourceACLSubscribeExceptionDelete,

		Schema: map[string]*schema.Schema{
			"topic": {
				Type:        schema.TypeString,
				Description: "The name of the Topic for the Exception to the default action taken. May include wildcard characters.",
				Required:    true,
				ForceNew:    true,
			},
			"topic_syntax": {
				Type:        schema.TypeString,
				Description: "The syntax of the Topic for the Exception to the default action taken. The allowed values are \"smf\" and \"mqtt\"",
				Required:    true,
				ForceNew:    true,
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

	SubExc := models.MsgVpnACLProfileSubscribeTopicException{
		SubscribeTopicExceptionSyntax: syntax,
		SubscribeTopicException: topic,
		ACLProfileName:          acl,
		MsgVpnName:              vpn,
	}

	params := all.NewCreateMsgVpnACLProfileSubscribeTopicExceptionParams()
	params.MsgVpnName = vpn
	params.ACLProfileName = acl
	params.Body = &SubExc

	resp, err := client.All.CreateMsgVpnACLProfileSubscribeTopicException(params, auth)
	if err != nil {
		sempErr := err.(*all.CreateMsgVpnACLProfileSubscribeExceptionDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create ACL Profile subscribe exception %q for ACL %q on VPN %q: %q", SubExc, acl, vpn, formatError(sempErr))
	}
	d.SetId(resp.Payload.Data.SubscribeTopicException)

	log.Printf("[DEBUG] Finished creating ACL Profile subscribe exception %q on ACL %q on VPN %q", SubExc, acl, vpn)
	return resourceACLSubscribeExceptionRead(d, m)
}

func resourceACLSubscribeExceptionRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Reading ACL Profile subscribe exception %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewGetMsgVpnACLProfileSubscribeTopicExceptionParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.SubscribeTopicException = d.Id()
	params.SubscribeTopicExceptionSyntax = d.Get("topic_syntax").(string)
	params.ACLProfileName = d.Get("acl").(string)
	params.MsgVpnName = vpn

	resp, err := client.All.GetMsgVpnACLProfileSubscribeTopicException(params, auth)
	if err != nil {
		log.Printf("[WARN] No ACL Profile subscribe exception found found: %q", d.Id())
		d.SetId("")
		return nil
	}
	log.Printf("%#v\n", resp.Payload.Data)
	d.Set("topic", resp.Payload.Data.SubscribeTopicException)
	d.Set("topic_syntax", resp.Payload.Data.SubscribeTopicExceptionSyntax)
	d.Set("acl", resp.Payload.Data.ACLProfileName)
	d.Set("msg_vpn", resp.Payload.Data.MsgVpnName)

	return nil
}

func resourceACLSubscribeExceptionDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Deleting ACL Profile subscribe exception %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewDeleteMsgVpnACLProfileSubscribeTopicExceptionParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}
	params.SubscribeTopicException = d.Id()
	params.SubscribeTopicExceptionSyntax = d.Get("topic_syntax").(string)
	params.ACLProfileName = d.Get("acl").(string)
	params.MsgVpnName = vpn

	_, err = client.All.DeleteMsgVpnACLProfileSubscribeTopicException(params, auth)
	if err != nil {
		sempErr := err.(*all.DeleteMsgVpnACLProfileSubscribeExceptionDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to ACL Profile subscribe exception exception %q for ACL %q on VPN %q: %q",
			params.SubscribeTopicExceptionSyntax, params.ACLProfileName, params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceACLSubscribeExceptionImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
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
