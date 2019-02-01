package solace

import (
	"fmt"
	"log"
	"strings"

	"github.com/ExalDraen/semp-client/models"

	"github.com/ExalDraen/semp-client/client/operations"
	"github.com/hashicorp/terraform/helper/schema"
) 

func resourceACLClientConnException() *schema.Resource {
	return &schema.Resource{
		Create: resourceACLClientConnExceptionCreate,
		Read:   resourceACLClientConnExceptionRead,
		Delete: resourceACLClientConnExceptionDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The IP Address/Netmask of the Client Connect Exception in the Classless Inter-Domain Routing (CIDR) form. Used as an identifier.",
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
			State: resourceACLClientConnExceptionImport,
		},
	}
}

func resourceACLClientConnExceptionCreate(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Creating ACL Profile client exception...")

	// Get our Solace client
	c := m.(*Config)
	client := c.Client
	auth := c.Auth

	// Extract config data from resource data and prepare new VPN object
	address := d.Get("address").(string)
	acl := d.Get("acl").(string)
	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	ClientConnExc := models.MsgVpnACLProfileClientConnectException{
		ClientConnectExceptionAddress: address,
		ACLProfileName:                acl,
		MsgVpnName:                    vpn,
	}

	params := operations.NewCreateMsgVpnACLProfileClientConnectExceptionParams()
	params.MsgVpnName = vpn
	params.ACLProfileName = acl
	params.Body = &ClientConnExc

	resp, err := client.Operations.CreateMsgVpnACLProfileClientConnectException(params, auth)
	if err != nil {
		sempErr := err.(*operations.CreateMsgVpnACLProfileClientConnectExceptionDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create ACL profile client connection exception %q for ACL %q on VPN %q: %q", ClientConnExc, acl, vpn, formatError(sempErr))
	}
	d.SetId(resp.Payload.Data.ClientConnectExceptionAddress)

	log.Printf("[DEBUG] Finished creating Client connection exception %q on ACL %q on VPN %q", ClientConnExc, acl, vpn)
	return resourceACLClientConnExceptionRead(d, m)
}

func resourceACLClientConnExceptionRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Reading ACL Profile client exception %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := operations.NewGetMsgVpnACLProfileClientConnectExceptionParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.ClientConnectExceptionAddress = d.Id()
	params.ACLProfileName = d.Get("acl").(string)
	params.MsgVpnName = vpn

	resp, err := client.Operations.GetMsgVpnACLProfileClientConnectException(params, auth)
	if err != nil {
		log.Printf("[WARN] No ACL profile client exception found found: %q", d.Id())
		d.SetId("")
		return nil
	}
	log.Printf("%#v\n", resp.Payload.Data)
	d.Set("address", resp.Payload.Data.ClientConnectExceptionAddress)
	d.Set("acl", resp.Payload.Data.ACLProfileName)
	d.Set("msg_vpn", resp.Payload.Data.MsgVpnName)

	return nil
} 

func resourceACLClientConnExceptionDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Deleting ACL Profile client exception %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := operations.NewDeleteMsgVpnACLProfileClientConnectExceptionParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}
	params.ClientConnectExceptionAddress = d.Id()
	params.ACLProfileName = d.Get("acl").(string)
	params.MsgVpnName = vpn

	_, err = client.Operations.DeleteMsgVpnACLProfileClientConnectException(params, auth)
	if err != nil {
		sempErr := err.(*operations.DeleteMsgVpnACLProfileClientConnectExceptionDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to ACL profile client connection exception %q for ACL %q on VPN %q: %q",
			params.ClientConnectExceptionAddress, params.ACLProfileName, params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceACLClientConnExceptionImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "|")
	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected MSG-VPN|ACL-PROFILE|CLIENT-CONN-EXC-ADDR", d.Id())
	}
	vpn := idParts[0]
	acl := idParts[1]
	ClientConnExc := idParts[2]
	d.Set("msg_vpn", vpn)
	d.Set("acl", acl)
	d.SetId(ClientConnExc)
	return []*schema.ResourceData{d}, nil
}
