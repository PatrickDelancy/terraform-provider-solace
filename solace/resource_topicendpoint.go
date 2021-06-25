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

func resourceTopicEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceTopicEndpointCreate,
		Read:   resourceTopicEndpointRead,
		Update: resourceTopicEndpointUpdate,
		Delete: resourceTopicEndpointDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the Topic Endpoint. Used as a unique identifier.",
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
				Type:         schema.TypeString,
				Description:  "The access type for delivering messages to consumer flows bound to the Topic Endpoint. The default value is \"exclusive\". Must be one of \"exclusive\" or \"non-exclusive\"",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"exclusive", "non-exclusive"}, false),
			},
			"dead_msg_queue": {
				Type:        schema.TypeString,
				Description: "The name of the Dead Message Queue (DMQ) used by the Topic Endpoint. The default value is \"#DEAD_MSG_QUEUE\".",
				Optional:    true,
			},
			"egress_enabled": {
				Type:        schema.TypeBool,
				Description: "Enable or disable the transmission of messages from the Topic Endpoint. The default value is false.",
				Optional:    true,
			},
			"ingess_enabled": {
				Type:        schema.TypeBool,
				Description: "Enable or disable the reception of messages to the Topic Endpoint. The default value is false.",
				Optional:    true,
			},
			"redelivery_enabled": {
				Type:        schema.TypeBool,
				Description: "Enable or disable message redelivery. When enabled, the number of redelivery attempts is controlled by maxRedeliveryCount. When disabled, the message will never be delivered from the topic-endpoint more than once. The default value is true. Available since 2.18.",
				Optional:    true,
			},
			"max_redelivery_count": {
				Type:        schema.TypeInt,
				Description: "The maximum number of times the Topic Endpoint will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. The default value is 0.",
				Optional:    true,
			},
			"owner": {
				Type:        schema.TypeString,
				Description: "The Client Username that owns the TopicEndpoint and has permission equivalent to \"delete\". The default value is \"\".",
				Optional:    true,
			},
			"permission": {
				Type:         schema.TypeString,
				Description:  "The permission level for all consumers of the TopicEndpoint, excluding the owner. The default value is \"no-access\". Must be one of \"no-access\", \"read-only\", \"consume\", \"modify-topic\" or \"delete\"",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"no-access", "read-only", "consume", "modify-topic", "delete"}, false),
			},
		},
		Importer: &schema.ResourceImporter{
			State: resourceTopicEndpointImport,
		},
	}
}

func resourceTopicEndpointCreate(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Creating topic endpoint ...")

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

	topicEndpoint := models.MsgVpnTopicEndpoint{
		TopicEndpointName: name,
		MsgVpnName:        vpn,
	}
	// Only set these if they're actually set (not their default value)
	if v, ok := d.GetOk("access_type"); ok {
		topicEndpoint.AccessType = v.(string)
	}
	if v, ok := d.GetOk("dead_msg_queue"); ok {
		topicEndpoint.DeadMsgQueue = v.(string)
	}
	if v, ok := d.GetOk("egress_enabled"); ok {
		topicEndpoint.EgressEnabled = v.(bool)
	}
	if v, ok := d.GetOk("ingess_enabled"); ok {
		topicEndpoint.IngressEnabled = v.(bool)
	}
	if v, ok := d.GetOk("redelivery_enabled"); ok {
		topicEndpoint.RedeliveryEnabled = v.(bool)
	}
	if v, ok := d.GetOk("max_redelivery_count"); ok {
		topicEndpoint.MaxRedeliveryCount = int64(v.(int))
	}
	if v, ok := d.GetOk("owner"); ok {
		topicEndpoint.Owner = v.(string)
	}
	if v, ok := d.GetOk("permission"); ok {
		topicEndpoint.Permission = v.(string)
	}

	params := all.NewCreateMsgVpnTopicEndpointParams()
	params.MsgVpnName = vpn
	params.Body = &topicEndpoint

	resp, err := client.All.CreateMsgVpnTopicEndpoint(params, auth)
	if err != nil {
		sempErr := err.(*all.CreateMsgVpnTopicEndpointDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create topic endpoint %q on vpn %q: %v", name, vpn, formatError(sempErr))
	}
	d.SetId(resp.Payload.Data.TopicEndpointName)

	log.Printf("[DEBUG] Finished creating topic endpoint %q on VPN %q", name, vpn)
	return resourceTopicEndpointRead(d, m)
}

func resourceTopicEndpointRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Reading topic endpoint %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewGetMsgVpnTopicEndpointParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.TopicEndpointName = d.Id()
	params.MsgVpnName = vpn

	resp, err := client.All.GetMsgVpnTopicEndpoint(params, auth)
	if err != nil {
		log.Printf("[WARN] No topic endpoint found: %s", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", resp.Payload.Data.TopicEndpointName)
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

func resourceTopicEndpointUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Updating topic endpoint %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewUpdateMsgVpnTopicEndpointParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.TopicEndpointName = d.Id()
	params.MsgVpnName = vpn
	queue := models.MsgVpnTopicEndpoint{}
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

	_, err = client.All.UpdateMsgVpnTopicEndpoint(params, auth)
	if err != nil {
		sempErr := err.(*all.UpdateMsgVpnTopicEndpointDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to update topic endpoint %q: %v", params.TopicEndpointName, formatError(sempErr))
	}

	return resourceTopicEndpointRead(d, m)
}

func resourceTopicEndpointDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewDeleteMsgVpnTopicEndpointParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}
	params.TopicEndpointName = d.Id()
	params.MsgVpnName = vpn

	_, err = client.All.DeleteMsgVpnTopicEndpoint(params, auth)
	if err != nil {
		sempErr := err.(*all.DeleteMsgVpnTopicEndpointDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to delete topic endpoint %q: %v", params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceTopicEndpointImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "/")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected MSG-VPN/TOPIC-ENDPOINT-NAME", d.Id())
	}
	vpn := idParts[0]
	name := idParts[1]
	d.Set("msg_vpn", vpn)

	d.SetId(name)
	return []*schema.ResourceData{d}, nil
}
