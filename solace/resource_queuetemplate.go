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

func resourceQueueTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceQueueTemplateCreate,
		Read:   resourceQueueTemplateRead,
		Update: resourceQueueTemplateUpdate,
		Delete: resourceQueueTemplateDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the QueueTemplate. Used as a unique identifier.",
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
			"queue_name_filter": {
				Type:        schema.TypeString,
				Description: "A wildcardable pattern used to determine which Queues use settings from this Template. Two different wildcards are supported: * and >. Similar to topic filters or subscription patterns, a > matches anything (but only when used at the end), and a * matches zero or more characters but never a slash (/). A > is only a wildcard when used at the end, after a /. A * is only allowed at the end, after a slash (/). The default value is \"\".",
				Required:    true,
			},
			"access_type": {
				Type:         schema.TypeString,
				Description:  "The access type for delivering messages to consumer flows bound to the Queue. The default value is 'exclusive'. Must be one of \"exclusive\" or \"non-exclusive\"",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"exclusive", "non-exclusive"}, false),
			},
			"dead_msg_queue": {
				Type:        schema.TypeString,
				Description: "The name of the Dead Message Queue (DMQ). The default value is \"#DEAD_MSG_QUEUE\".",
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
			"permission": {
				Type:         schema.TypeString,
				Description:  "The permission level for all consumers of the Queue, excluding the owner. The default value is \"no-access\". Must be one of \"no-access\", \"read-only\", \"consume\", \"modify-topic\" or \"delete\"",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"no-access", "read-only", "consume", "modify-topic", "delete"}, false),
			},
		},
		Importer: &schema.ResourceImporter{
			State: resourceQueueTemplateImport,
		},
	}
}

func resourceQueueTemplateCreate(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Creating queue template ...")

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

	queueTemplate := models.MsgVpnQueueTemplate{
		QueueTemplateName: name,
		MsgVpnName:        vpn,
	}
	// Only set these if they're actually set (not their default value)
	if v, ok := d.GetOk("queue_name_filter"); ok {
		queueTemplate.QueueNameFilter = v.(string)
	}
	if v, ok := d.GetOk("access_type"); ok {
		queueTemplate.AccessType = v.(string)
	}
	if v, ok := d.GetOk("dead_msg_queue"); ok {
		queueTemplate.DeadMsgQueue = v.(string)
	}
	if v, ok := d.GetOk("redelivery_enabled"); ok {
		queueTemplate.RedeliveryEnabled = v.(bool)
	}
	if v, ok := d.GetOk("max_redelivery_count"); ok {
		queueTemplate.MaxRedeliveryCount = int64(v.(int))
	}
	if v, ok := d.GetOk("permission"); ok {
		queueTemplate.Permission = v.(string)
	}

	params := all.NewCreateMsgVpnQueueTemplateParams()
	params.MsgVpnName = vpn
	params.Body = &queueTemplate

	resp, err := client.All.CreateMsgVpnQueueTemplate(params, auth)
	if err != nil {
		sempErr := err.(*all.CreateMsgVpnQueueTemplateDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create queue template %q on vpn %q: %v", name, vpn, formatError(sempErr))
	}
	d.SetId(resp.Payload.Data.QueueTemplateName)

	log.Printf("[DEBUG] Finished creating queue template %q on VPN %q", name, vpn)
	return resourceQueueTemplateRead(d, m)
}

func resourceQueueTemplateRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Reading queue template %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewGetMsgVpnQueueTemplateParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.QueueTemplateName = d.Id()
	params.MsgVpnName = vpn

	resp, err := client.All.GetMsgVpnQueueTemplate(params, auth)
	if err != nil {
		log.Printf("[WARN] No ACL profile found: %s", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", resp.Payload.Data.QueueTemplateName)
	d.Set("msg_vpn", resp.Payload.Data.MsgVpnName)
	d.Set("queue_name_filter", resp.Payload.Data.QueueNameFilter)
	d.Set("access_type", resp.Payload.Data.AccessType)
	d.Set("dead_msg_queue", resp.Payload.Data.DeadMsgQueue)
	d.Set("redelivery_enabled", resp.Payload.Data.RedeliveryEnabled)
	d.Set("max_redelivery_count", resp.Payload.Data.MaxRedeliveryCount)
	d.Set("permission", resp.Payload.Data.Permission)

	return nil
}

func resourceQueueTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Updating queue template %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewUpdateMsgVpnQueueTemplateParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.QueueTemplateName = d.Id()
	params.MsgVpnName = vpn
	queueTemplate := models.MsgVpnQueueTemplate{}
	queueTemplate.MsgVpnName = vpn

	// Only include changed values; anything we don't specify does not get updated
	if d.HasChange("queue_name_filter") {
		queueTemplate.QueueNameFilter = d.Get("queue_name_filter").(string)
	}
	if d.HasChange("access_type") {
		queueTemplate.AccessType = d.Get("access_type").(string)
	}
	if d.HasChange("dead_msg_queue") {
		queueTemplate.DeadMsgQueue = d.Get("dead_msg_queue").(string)
	}
	if d.HasChange("redelivery_enabled") {
		queueTemplate.RedeliveryEnabled = d.Get("redelivery_enabled").(bool)
	}
	if d.HasChange("max_redelivery_count") {
		queueTemplate.MaxRedeliveryCount = int64(d.Get("max_redelivery_count").(int))
	}
	if d.HasChange("permission") {
		queueTemplate.Permission = d.Get("permission").(string)
	}
	params.Body = &queueTemplate

	_, err = client.All.UpdateMsgVpnQueueTemplate(params, auth)
	if err != nil {
		sempErr := err.(*all.UpdateMsgVpnQueueTemplateDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to update queue template %q: %v", params.QueueTemplateName, formatError(sempErr))
	}

	return resourceQueueTemplateRead(d, m)
}

func resourceQueueTemplateDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewDeleteMsgVpnQueueTemplateParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}
	params.QueueTemplateName = d.Id()
	params.MsgVpnName = vpn

	_, err = client.All.DeleteMsgVpnQueueTemplate(params, auth)
	if err != nil {
		sempErr := err.(*all.DeleteMsgVpnQueueTemplateDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to delete queue template %q: %v", params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceQueueTemplateImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "/")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected MSG-VPN/QUEUE-TEMPLATE-NAME", d.Id())
	}
	vpn := idParts[0]
	name := idParts[1]
	d.Set("msg_vpn", vpn)

	d.SetId(name)
	return []*schema.ResourceData{d}, nil
}
