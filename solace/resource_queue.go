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

func resourceQueue() *schema.Resource {
	return &schema.Resource{
		Create: resourceQueueCreate,
		Read:   resourceQueueRead,
		Update: resourceQueueUpdate,
		Delete: resourceQueueDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the Queue. Used as a unique identifier.",
				Required:    true,
				ForceNew:    true,
			},
			"msg_vpn": {
				Type:        schema.TypeString,
				Description: "The name of the MSG VPN. If unset the provider default is used.",
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
			},
			"access_type": {
				Type: schema.TypeString,
				Description: "The access type for delivering messages to consumer flows bound to the Queue. The default value is 'exclusive'. Must be one of \"exclusive\" or \"non-exclusive\"",
				Optional:    true,
				ValidateFunc: validation.StringInSlice([]string{"exclusive", "non-exclusive"}, false),
			},
			"dead_msg_queue": {
				Type:        schema.TypeString,
				Description: "The name of the Dead Message Queue (DMQ) used by the Queue. The default value is \"#DEAD_MSG_QUEUE\".",
				Optional:    true,
			},
			"egress_enabled": {
				Type:        schema.TypeBool,
				Description: "Enable or disable the transmission of messages from the Queue. The default value is false.",
				Optional:    true,
			},
			"ingess_enabled": {
				Type:        schema.TypeBool,
				Description: "Enable or disable the reception of messages to the Queue. The default value is false.",
				Optional:    true,
			},
			"redelivery_enabled": {
				Type:        schema.TypeBool,
				Description: "Enable or disable message redelivery. When enabled, the number of redelivery attempts is controlled by maxRedeliveryCount. When disabled, the message will never be delivered from the queue more than once. The default value is true. Available since 2.18.",
				Optional:    true,
			},
			"max_redelivery_count": {
				Type:        schema.TypeInt,
				Description: "The maximum number of times the Queue will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. The default value is 0.",
				Optional:    true,
			},
			"owner": {
				Type:        schema.TypeString,
				Description: "The Client Username that owns the Queue and has permission equivalent to \"delete\". The default value is \"\".",
				Optional:    true,
			},
			"permission": {
				Type: schema.TypeString,
				Description: "The permission level for all consumers of the Queue, excluding the owner. The default value is \"no-access\". Must be one of \"no-access\", \"read-only\", \"consume\", \"modify-topic\" or \"delete\"",
				Optional:    true,
				ValidateFunc: validation.StringInSlice([]string{"no-access", "read-only", "consume", "modify-topic", "delete"}, false),
			},
		},
		Importer: &schema.ResourceImporter{
			State: resourceQueueImport,
		},
	}
}

func resourceQueueCreate(d *schema.ResourceData, m interface{}) error {
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

	queue := models.MsgVpnQueue{
		QueueName: name,
		MsgVpnName:     vpn,
	}
	// Only set these if they're actually set (not their default value)
	if v, ok := d.GetOk("access_type"); ok {
		queue.AccessType = v.(string)
	}
	if v, ok := d.GetOk("dead_msg_queue"); ok {
		queue.DeadMsgQueue = v.(string)
	}
	if v, ok := d.GetOk("egress_enabled"); ok {
		queue.EgressEnabled = v.(bool)
	}
	if v, ok := d.GetOk("ingess_enabled"); ok {
		queue.IngressEnabled = v.(bool)
	}
	if v, ok := d.GetOk("redelivery_enabled"); ok {
		queue.RedeliveryEnabled = v.(bool)
	}
	if v, ok := d.GetOk("max_redelivery_count"); ok {
		queue.MaxRedeliveryCount = int64(v.(int))
	}
	if v, ok := d.GetOk("owner"); ok {
		queue.Owner = v.(string)
	}
	if v, ok := d.GetOk("permission"); ok {
		queue.Permission = v.(string)
	}

	params := all.NewCreateMsgVpnQueueParams()
	params.MsgVpnName = vpn
	params.Body = &queue

	resp, err := client.All.CreateMsgVpnQueue(params, auth)
	if err != nil {
		sempErr := err.(*all.CreateMsgVpnQueueDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create queue %q on vpn %q: %v", name, vpn, formatError(sempErr))
	}
	d.SetId(resp.Payload.Data.QueueName)

	log.Printf("[DEBUG] Finished creating queue %q on VPN %q", name, vpn)
	return resourceQueueRead(d, m)
}

func resourceQueueRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Reading ACL Profile %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewGetMsgVpnQueueParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.QueueName = d.Id()
	params.MsgVpnName = vpn

	resp, err := client.All.GetMsgVpnQueue(params, auth)
	if err != nil {
		log.Printf("[WARN] No ACL profile found: %s", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", resp.Payload.Data.QueueName)
	d.Set("msg_vpn", resp.Payload.Data.MsgVpnName)
	d.Set("access_type", resp.Payload.Data.AccessType)
	d.Set("dead_msg_queue", resp.Payload.Data.DeadMsgQueue)
	d.Set("egress_enabled", resp.Payload.Data.EgressEnabled)
	d.Set("ingess_enabled", resp.Payload.Data.IngressEnabled)
	d.Set("redelivery_enabled", resp.Payload.Data.RedeliveryEnabled)
	d.Set("max_redelivery_count", resp.Payload.Data.MaxRedeliveryCount)
	d.Set("owner", resp.Payload.Data.Owner)
	d.Set("permission", resp.Payload.Data.Permission)

	return nil
}

func resourceQueueUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Updating ACL Profile %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewUpdateMsgVpnQueueParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.QueueName = d.Id()
	params.MsgVpnName = vpn
	queue := models.MsgVpnQueue{}
	queue.MsgVpnName = vpn

	// Only include changed values; anything we don't specify does not get updated
	if d.HasChange("access_type") {
		queue.AccessType = d.Get("access_type").(string)
	}
	if d.HasChange("dead_msg_queue") {
		queue.DeadMsgQueue = d.Get("dead_msg_queue").(string)
	}
	if d.HasChange("egress_enabled") {
		queue.EgressEnabled = d.Get("egress_enabled").(bool)
	}
	if d.HasChange("ingess_enabled") {
		queue.IngressEnabled = d.Get("ingess_enabled").(bool)
	}
	if d.HasChange("redelivery_enabled") {
		queue.RedeliveryEnabled = d.Get("redelivery_enabled").(bool)
	}
	if d.HasChange("max_redelivery_count") {
		queue.MaxRedeliveryCount = int64(d.Get("max_redelivery_count").(int))
	}
	if d.HasChange("owner") {
		queue.Owner = d.Get("owner").(string)
	}
	if d.HasChange("permission") {
		queue.Permission = d.Get("permission").(string)
	}
	params.Body = &queue

	_, err = client.All.UpdateMsgVpnQueue(params, auth)
	if err != nil {
		sempErr := err.(*all.UpdateMsgVpnQueueDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to update queue %q: %v", params.QueueName, formatError(sempErr))
	}

	return resourceQueueRead(d, m)
}

func resourceQueueDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewDeleteMsgVpnQueueParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}
	params.QueueName = d.Id()
	params.MsgVpnName = vpn

	_, err = client.All.DeleteMsgVpnQueue(params, auth)
	if err != nil {
		sempErr := err.(*all.DeleteMsgVpnQueueDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to delete queue %q: %v", params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceQueueImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "/")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected MSG-VPN/QUEUE-NAME", d.Id())
	}
	vpn := idParts[0]
	acl := idParts[1]
	d.Set("msg_vpn", vpn)

	d.SetId(acl)
	return []*schema.ResourceData{d}, nil
}
