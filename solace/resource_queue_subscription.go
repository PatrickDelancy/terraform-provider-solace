package solace

import (
	"fmt"
	"log"
	"strings"

	"github.com/PatrickDelancy/semp-client/client/all"
	"github.com/PatrickDelancy/semp-client/models"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceQueueSubscription() *schema.Resource {
	return &schema.Resource{
		Create: resourceQueueSubscriptionCreate,
		Read:   resourceQueueSubscriptionRead,
		// Update: resourceQueueSubscriptionUpdate,
		Delete: resourceQueueSubscriptionDelete,

		Schema: map[string]*schema.Schema{
			"queue_name": {
				Type:        schema.TypeString,
				Description: "The name of the Queue. Used as part of a unique identifier.",
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
			"topic": {
				Type: schema.TypeString,
				Description: "The topic of the subscription. Used as part of a unique identifier.",
				Required:    true,
				ForceNew:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: resourceQueueSubscriptionImport,
		},
	}
}

func resourceQueueSubscriptionCreate(d *schema.ResourceData, m interface{}) error {
	log.Print("[DEBUG] Creating Queue Subscription ...")

	// Get our Solace client
	c := m.(*Config)
	client := c.Client
	auth := c.Auth

	// Extract config data from resource data and prepare new VPN object
	name := d.Get("queue_name").(string)
	topic := d.Get("topic").(string)
	
	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

	queue := models.MsgVpnQueueSubscription{
		QueueName:      	name,
		MsgVpnName:     	vpn,
		SubscriptionTopic:	topic,
	}

	params := all.NewCreateMsgVpnQueueSubscriptionParams()
	params.MsgVpnName = vpn
	params.Body = &queue

	resp, err := client.All.CreateMsgVpnQueueSubscription(params, auth)
	if err != nil {
		sempErr := err.(*all.CreateMsgVpnQueueSubscriptionDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to create queue subscription %q|%q on vpn %q: %v", name, topic, vpn, formatError(sempErr))
	}
	d.SetId(resp.Payload.Data.QueueName + "|" + resp.Payload.Data.SubscriptionTopic)

	log.Printf("[DEBUG] Finished creating queue subscription %q|%q on VPN %q", name, topic, vpn)
	return resourceQueueSubscriptionRead(d, m)
}

func resourceQueueSubscriptionRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Reading queue subscription %q ...", d.Id())
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewGetMsgVpnQueueSubscriptionParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}

 	idParts := strings.Split(d.Id(), "|")
	params.QueueName = idParts[0]
	params.SubscriptionTopic = idParts[1]
	params.MsgVpnName = vpn

	resp, err := client.All.GetMsgVpnQueueSubscription(params, auth)
	if err != nil {
		log.Printf("[WARN] No queue subscription found: %s", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("queue_name", resp.Payload.Data.QueueName)
	d.Set("msg_vpn", resp.Payload.Data.MsgVpnName)
	d.Set("topic", resp.Payload.Data.SubscriptionTopic)

	return nil
}

// func resourceQueueSubscriptionUpdate(d *schema.ResourceData, m interface{}) error {
// 	log.Printf("[DEBUG] Updating ACL Profile %q ...", d.Id())
// 	c := m.(*Config)
// 	client := c.Client
// 	auth := c.Auth
// 	params := all.NewUpdateMsgVpnQueueSubscriptionParams()

// 	vpn, err := getMsgVPN(d, c)
// 	if err != nil {
// 		return err
// 	}

// 	params.QueueSubscriptionName = d.Id()
// 	params.MsgVpnName = vpn
// 	queue := models.MsgVpnQueueSubscription{}
// 	queue.MsgVpnName = vpn

// 	// Only include changed values; anything we don't specify does not get updated
// 	if d.HasChange("access_type") {
// 		queue.AccessType = d.Get("access_type").(string)
// 	}
// 	if d.HasChange("dead_msg_queue") {
// 		queue.DeadMsgQueueSubscription = d.Get("dead_msg_queue").(string)
// 	}
// 	if d.HasChange("egress_enabled") {
// 		queue.EgressEnabled = d.Get("egress_enabled").(bool)
// 	}
// 	if d.HasChange("ingess_enabled") {
// 		queue.IngressEnabled = d.Get("ingess_enabled").(bool)
// 	}
// 	if d.HasChange("redelivery_enabled") {
// 		queue.RedeliveryEnabled = d.Get("redelivery_enabled").(bool)
// 	}
// 	if d.HasChange("max_redelivery_count") {
// 		queue.MaxRedeliveryCount = int64(d.Get("max_redelivery_count").(int))
// 	}
// 	if d.HasChange("owner") {
// 		queue.Owner = d.Get("owner").(string)
// 	}
// 	if d.HasChange("permission") {
// 		queue.Permission = d.Get("permission").(string)
// 	}
// 	params.Body = &queue

// 	_, err = client.All.UpdateMsgVpnQueueSubscription(params, auth)
// 	if err != nil {
// 		sempErr := err.(*all.UpdateMsgVpnQueueSubscriptionDefault).Payload.Meta.Error
// 		return fmt.Errorf("[ERROR] Unable to update queue %q: %v", params.QueueSubscriptionName, formatError(sempErr))
// 	}

// 	return resourceQueueSubscriptionRead(d, m)
// }

func resourceQueueSubscriptionDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Config)
	client := c.Client
	auth := c.Auth
	params := all.NewDeleteMsgVpnQueueSubscriptionParams()

	vpn, err := getMsgVPN(d, c)
	if err != nil {
		return err
	}
	idParts := strings.Split(d.Id(), "|")
	params.QueueName = idParts[0]
	params.SubscriptionTopic = idParts[1]
	params.MsgVpnName = vpn

	_, err = client.All.DeleteMsgVpnQueueSubscription(params, auth)
	if err != nil {
		sempErr := err.(*all.DeleteMsgVpnQueueSubscriptionDefault).Payload.Meta.Error
		return fmt.Errorf("[ERROR] Unable to delete queue %q: %v", params.MsgVpnName, formatError(sempErr))
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}

func resourceQueueSubscriptionImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	idParts := strings.Split(d.Id(), "/")
	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		return nil, fmt.Errorf("unexpected format of ID (%q), expected MSG-VPN/QUEUE-NAME/SUBSCRIPTION_TOPIC", d.Id())
	}
	vpn := idParts[0]
	queueName := idParts[1]
	topic := idParts[2]
	d.Set("msg_vpn", vpn)

	d.SetId(queueName + "|" + topic)
	return []*schema.ResourceData{d}, nil
}
