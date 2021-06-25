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

func resourceTopicEndpointTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceTopicEndpointTemplateCreate,
		Read:   resourceTopicEndpointTemplateRead,
		Update: resourceTopicEndpointTemplateUpdate,
		Delete: resourceTopicEndpointTemplateDelete,

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
			"topic_name_filter": {
				Type:        schema.TypeString,
				Description: "A wildcardable pattern used to determine which Topic Endpoints use settings from this Template. Two different wildcards are supported: * and >. Similar to topic filters or subscription patterns, a > matches anything (but only when used at the end), and a * matches zero or more characters but never a slash (/). A > is only a wildcard when used at the end, after a /. A * is only allowed at the end, after a slash (/). The default value is \"\".",
				Required:    true,
				ForceNew:    true,
			},
			"access_type": {
				Type:         schema.TypeString,
				Description:  "The access type for delivering messages to consumer flows bound to the Topic Endpoint. The default value is \"exclusive\". Must be one of \"exclusive\" or \"non-exclusive\"",
				Default:      "exclusive",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"exclusive", "non-exclusive"}, false),
			},
			"dead_msg_queue": {
				Type:        schema.TypeString,
				Description: "The name of the Dead Message Queue (DMQ) used by the Topic Endpoint. The default value is \"#DEAD_MSG_QUEUE\".",
				Default:     "#DEAD_MSG_QUEUE",
				Optional:    true,
			},
			"redelivery_enabled": {
				Type:        schema.TypeBool,
				Description: "Enable or disable message redelivery. When enabled, the number of redelivery attempts is controlled by maxRedeliveryCount. When disabled, the message will never be delivered from the topic-endpoint more than once. The default value is true. Available since 2.18.",
				Default:     true,
				Optional:    true,
			},
			"max_redelivery_count": {
				Type:        schema.TypeInt,
				Description: "The maximum number of times the Topic Endpoint will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. The default value is 0.",
				Default:     0,
				Optional:    true,
			},
			"permission": {
				Type:         schema.TypeString,
				Description:  "The permission level for all consumers of the TopicEndpointTemplate, excluding the owner. The default value is \"no-access\". Must be one of \"no-access\", \"read-only\", \"consume\", \"modify-topic\" or \"delete\"",
				Default:      "no-access",
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"no-access", "read-only", "consume", "modify-topic", "delete"}, false),
			},
		},
		Importer: &schema.ResourceImporter{
			State: resourceTopicEndpointTemplateImport,
		},
	}
}

func resourceTopicEndpointTemplateCreate(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Creating topic endpoint template ...")

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

	topicEndpoint := models.MsgVpnTopicEndpointTemplate{
		TopicEndpointTemplateName: name,
		MsgVpnName:                vpn,
	}
	// Only set these if they're actually set (not their default value)
	if v, ok := d.GetOk("access_type"); ok {
		topicEndpoint.AccessType = v.(string)
	}
	if v, ok := d.GetOk("topic_name_filter"); ok {
		topicEndpoint.TopicEndpointNameFilter = v.(string)
	}
	if v, ok := d.GetOk("dead_msg_queue"); ok {
		topicEndpoint.DeadMsgQueue = v.(string)
	}
	if v, ok := d.GetOk("redelivery_enabled"); ok {
		topicEndpoint.RedeliveryEnabled = v.(bool)
	}
	if v, ok := d.GetOk("max_redelivery_count"); ok {
		topicEndpoint.MaxRedeliveryCount = int64(v.(int))
	}
	if v, ok := d.GetOk("permission"); ok {
		topicEndpoint.Permission = v.(string)
	}

	params := all.NewCreateMsgVpnTopicEndpointTemplateParams()
	params.MsgVpnName = vpn
	params.Body = &topicEndpoint

	resp, err := client.All.CreateMsgVpnTopicEndpointTemplate(params, auth)
	if err != nil {
		sempErr := err.(*all.CreateMsgVpnTopicEndpointTemplateDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create topic endpoint template %q on vpn %q: %v", name, vpn, formatError(sempErr))
	}
	d.SetId(resp.Payload.Data.TopicEndpointTemplateName)

	log.Printf("[DEBUG] Finished creating topic endpoint template %q on VPN %q", name, vpn)
	return resourceTopicEndpointTemplateRead(d, m)
}

func resourceTopicEndpointTemplateRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Reading topic endpoint template %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewGetMsgVpnTopicEndpointTemplateParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.TopicEndpointTemplateName = d.Id()
	params.MsgVpnName = vpn

	resp, err := client.All.GetMsgVpnTopicEndpointTemplate(params, auth)
	if err != nil {
		log.Printf("[WARN] No topic endpoint template found: %s", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", resp.Payload.Data.TopicEndpointTemplateName)
	d.Set("msg_vpn", resp.Payload.Data.MsgVpnName)
	d.Set("topic_name_filter", resp.Payload.Data.TopicEndpointNameFilter)
	d.Set("access_type", resp.Payload.Data.AccessType)
	d.Set("dead_msg_queue", resp.Payload.Data.DeadMsgQueue)
	d.Set("redelivery_enabled", resp.Payload.Data.RedeliveryEnabled)
	d.Set("max_redelivery_count", resp.Payload.Data.MaxRedeliveryCount)
	d.Set("permission", resp.Payload.Data.Permission)

	return nil
}

func resourceTopicEndpointTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Updating topic endpoint template %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewUpdateMsgVpnTopicEndpointTemplateParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	params.TopicEndpointTemplateName = d.Id()
	params.MsgVpnName = vpn
	queue := models.MsgVpnTopicEndpointTemplate{}
	queue.MsgVpnName = vpn

	// Only include changed values; anything we don't specify does not get updated
	if d.HasChange("topic_name_filter") {
		queue.TopicEndpointNameFilter = d.Get("topic_name_filter").(string)
	}
	if d.HasChange("access_type") {
		queue.AccessType = d.Get("access_type").(string)
	}
	if d.HasChange("dead_msg_queue") {
		queue.DeadMsgQueue = d.Get("dead_msg_queue").(string)
	}
	if d.HasChange("redelivery_enabled") {
		queue.RedeliveryEnabled = d.Get("redelivery_enabled").(bool)
	}
	if d.HasChange("max_redelivery_count") {
		queue.MaxRedeliveryCount = int64(d.Get("max_redelivery_count").(int))
	}
	if d.HasChange("permission") {
		queue.Permission = d.Get("permission").(string)
	}
	params.Body = &queue

	_, err = client.All.UpdateMsgVpnTopicEndpointTemplate(params, auth)
	if err != nil {
		sempErr := err.(*all.UpdateMsgVpnTopicEndpointTemplateDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to update topic endpoint template %q: %v", params.TopicEndpointTemplateName, formatError(sempErr))
	}

	return resourceTopicEndpointTemplateRead(d, m)
}

func resourceTopicEndpointTemplateDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewDeleteMsgVpnTopicEndpointTemplateParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}
	params.TopicEndpointTemplateName = d.Id()
	params.MsgVpnName = vpn

	_, err = client.All.DeleteMsgVpnTopicEndpointTemplate(params, auth)
	if err != nil {
		sempErr := err.(*all.DeleteMsgVpnTopicEndpointTemplateDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to delete topic endpoint template %q: %v", params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceTopicEndpointTemplateImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "/")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected MSG-VPN/TOPIC-ENDPOINT-TEMPLATE-NAME", d.Id())
	}
	vpn := idParts[0]
	name := idParts[1]
	d.Set("msg_vpn", vpn)

	d.SetId(name)
	return []*schema.ResourceData{d}, nil
}
