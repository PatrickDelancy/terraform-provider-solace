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

func resourceACLPublishException() *schema.Resource {
	return &schema.Resource{
		Create: resourceACLPublishExceptionCreate,
		Read:   resourceACLPublishExceptionRead,
		Delete: resourceACLPublishExceptionDelete,

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

	PubExc := models.MsgVpnACLProfilePublishTopicException{
		PublishTopicExceptionSyntax:	syntax,
		PublishTopicException:			topic,
		ACLProfileName:        			acl,
		MsgVpnName:            			vpn,
	}

	params := all.NewCreateMsgVpnACLProfilePublishTopicExceptionParams()
	params.MsgVpnName = vpn
	params.ACLProfileName = acl
	params.Body = &PubExc

	resp, err := client.All.CreateMsgVpnACLProfilePublishTopicException(params, auth)
	if err != nil {
		sempErr := err.(*all.CreateMsgVpnACLProfilePublishTopicExceptionDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create ACL profile publish exception %q for ACL %q on VPN %q: %q", PubExc, acl, vpn, formatError(sempErr))
	}
	d.SetId(resp.Payload.Data.PublishTopicException)

	log.Printf("[DEBUG] Finished creating ACL profile publish exception %q on ACL %q on VPN %q", PubExc, acl, vpn)
	return resourceACLPublishExceptionRead(d, m)
}

func resourceACLPublishExceptionRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Reading ACL Profile publish exception %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewGetMsgVpnACLProfilePublishTopicExceptionParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.PublishTopicException = d.Id()
	params.PublishTopicExceptionSyntax = d.Get("topic_syntax").(string)
	params.ACLProfileName = d.Get("acl").(string)
	params.MsgVpnName = vpn

	resp, err := client.All.GetMsgVpnACLProfilePublishTopicException(params, auth)
	if err != nil {
		log.Printf("[WARN] No ACL Profile publish exception found found: %q", d.Id())
		d.SetId("")
		return nil
	}
	log.Printf("%#v\n", resp.Payload.Data)
	d.Set("topic", resp.Payload.Data.PublishTopicException)
	d.Set("topic_syntax", resp.Payload.Data.PublishTopicExceptionSyntax)
	d.Set("acl", resp.Payload.Data.ACLProfileName)
	d.Set("msg_vpn", resp.Payload.Data.MsgVpnName)

	return nil
}

func resourceACLPublishExceptionDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Deleting ACL Profile publish exception %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewDeleteMsgVpnACLProfilePublishTopicExceptionParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}
	params.PublishTopicException = d.Id()
	params.PublishTopicExceptionSyntax = d.Get("topic_syntax").(string)
	params.ACLProfileName = d.Get("acl").(string)
	params.MsgVpnName = vpn

	_, err = client.All.DeleteMsgVpnACLProfilePublishTopicException(params, auth)
	if err != nil {
		sempErr := err.(*all.DeleteMsgVpnACLProfilePublishTopicExceptionDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to ACL profile publish exception exception %q for ACL %q on VPN %q: %q",
			params.PublishTopicExceptionSyntax, params.ACLProfileName, params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceACLPublishExceptionImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "|")
	if len(idParts) != 4 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" || idParts[3] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected MSG-VPN|ACL-PROFILE|TOPIC-SYNTAX|PUB-EXC-TOPIC", d.Id())
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
