package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ExalDraen/semp-client/models"

	"github.com/ExalDraen/semp-client/client/operations"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceACLPublishException() *schema.Resource {
	return &schema.Resource{
		Create: resourceACLPublishExceptionCreate,
		Read:   resourceACLPublishExceptionRead,
		Delete: resourceACLPublishExceptionDelete,

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
			State: resourceACLPublishExceptionImport,
		},
	}
}

func resourceACLPublishExceptionCreate(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Creating ACL Profile publish exception...")

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

	PubExc := models.MsgVpnACLProfilePublishException{
		TopicSyntax:           syntax,
		PublishExceptionTopic: topic,
		ACLProfileName:        acl,
		MsgVpnName:            vpn,
	}

	params := operations.NewCreateMsgVpnACLProfilePublishExceptionParams()
	params.MsgVpnName = vpn
	params.ACLProfileName = acl
	params.Body = &PubExc

	resp, err := client.Operations.CreateMsgVpnACLProfilePublishException(params, auth)
	if err != nil {
		sempErr := err.(*operations.CreateMsgVpnACLProfilePublishExceptionDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create ACL profile publish exception %q for ACL %q on VPN %q: %q", PubExc, acl, vpn, formatError(sempErr))
	}
	d.SetId(resp.Payload.Data.PublishExceptionTopic)

	log.Printf("[DEBUG] Finished creating ACL profile publish exception %q on ACL %q on VPN %q", PubExc, acl, vpn)
	return resourceACLPublishExceptionRead(d, m)
}

func resourceACLPublishExceptionRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Reading ACL Profile publish exception %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := operations.NewGetMsgVpnACLProfilePublishExceptionParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.PublishExceptionTopic = d.Id()
	params.TopicSyntax = d.Get("topic_syntax").(string)
	params.ACLProfileName = d.Get("acl").(string)
	params.MsgVpnName = vpn

	resp, err := client.Operations.GetMsgVpnACLProfilePublishException(params, auth)
	if err != nil {
		log.Printf("[WARN] No ACL Profile publish exception found found: %q", d.Id())
		d.SetId("")
		return nil
	}
	log.Printf("%#v\n", resp.Payload.Data)
	d.Set("topic", resp.Payload.Data.PublishExceptionTopic)
	d.Set("topic_syntax", resp.Payload.Data.TopicSyntax)
	d.Set("acl", resp.Payload.Data.ACLProfileName)
	d.Set("msg_vpn", resp.Payload.Data.MsgVpnName)

	return nil
}

func resourceACLPublishExceptionDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Deleting ACL Profile publish exception %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := operations.NewDeleteMsgVpnACLProfilePublishExceptionParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}
	params.PublishExceptionTopic = d.Id()
	params.TopicSyntax = d.Get("topic_syntax").(string)
	params.ACLProfileName = d.Get("acl").(string)
	params.MsgVpnName = vpn

	_, err = client.Operations.DeleteMsgVpnACLProfilePublishException(params, auth)
	if err != nil {
		sempErr := err.(*operations.DeleteMsgVpnACLProfilePublishExceptionDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to ACL profile publish exception exception %q for ACL %q on VPN %q: %q",
			params.TopicSyntax, params.ACLProfileName, params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceACLPublishExceptionImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "/")
	if len(idParts) != 4 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" || idParts[3] == "" {
		return nil, fmt.Errorf("Unexpected format of ID (%q), expected MSG-VPN/ACL-PROFILE/TOPIC-SYNTAX/PUB-EXC-TOPIC", d.Id())
	}
	vpn := idParts[0]
	acl := idParts[1]
	syntax := idParts[2]
	PubExc := idParts[3]
	d.Set("msg_vpn", vpn)
	d.Set("acl", acl)
	d.Set("topic_syntax", syntax)
	d.SetId(PubExc)
	return []*schema.ResourceData{d}, nil
}
