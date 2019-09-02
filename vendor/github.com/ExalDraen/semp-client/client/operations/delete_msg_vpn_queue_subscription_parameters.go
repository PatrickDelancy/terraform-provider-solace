// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteMsgVpnQueueSubscriptionParams creates a new DeleteMsgVpnQueueSubscriptionParams object
// with the default values initialized.
func NewDeleteMsgVpnQueueSubscriptionParams() *DeleteMsgVpnQueueSubscriptionParams {
	var ()
	return &DeleteMsgVpnQueueSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnQueueSubscriptionParamsWithTimeout creates a new DeleteMsgVpnQueueSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMsgVpnQueueSubscriptionParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnQueueSubscriptionParams {
	var ()
	return &DeleteMsgVpnQueueSubscriptionParams{

		timeout: timeout,
	}
}

// NewDeleteMsgVpnQueueSubscriptionParamsWithContext creates a new DeleteMsgVpnQueueSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMsgVpnQueueSubscriptionParamsWithContext(ctx context.Context) *DeleteMsgVpnQueueSubscriptionParams {
	var ()
	return &DeleteMsgVpnQueueSubscriptionParams{

		Context: ctx,
	}
}

// NewDeleteMsgVpnQueueSubscriptionParamsWithHTTPClient creates a new DeleteMsgVpnQueueSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMsgVpnQueueSubscriptionParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnQueueSubscriptionParams {
	var ()
	return &DeleteMsgVpnQueueSubscriptionParams{
		HTTPClient: client,
	}
}

/*DeleteMsgVpnQueueSubscriptionParams contains all the parameters to send to the API endpoint
for the delete msg vpn queue subscription operation typically these are written to a http.Request
*/
type DeleteMsgVpnQueueSubscriptionParams struct {

	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*QueueName
	  The queueName of the Queue.

	*/
	QueueName string
	/*SubscriptionTopic
	  The subscriptionTopic of the Queue Subscription.

	*/
	SubscriptionTopic string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete msg vpn queue subscription params
func (o *DeleteMsgVpnQueueSubscriptionParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnQueueSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn queue subscription params
func (o *DeleteMsgVpnQueueSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn queue subscription params
func (o *DeleteMsgVpnQueueSubscriptionParams) WithContext(ctx context.Context) *DeleteMsgVpnQueueSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn queue subscription params
func (o *DeleteMsgVpnQueueSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn queue subscription params
func (o *DeleteMsgVpnQueueSubscriptionParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnQueueSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn queue subscription params
func (o *DeleteMsgVpnQueueSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn queue subscription params
func (o *DeleteMsgVpnQueueSubscriptionParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnQueueSubscriptionParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn queue subscription params
func (o *DeleteMsgVpnQueueSubscriptionParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithQueueName adds the queueName to the delete msg vpn queue subscription params
func (o *DeleteMsgVpnQueueSubscriptionParams) WithQueueName(queueName string) *DeleteMsgVpnQueueSubscriptionParams {
	o.SetQueueName(queueName)
	return o
}

// SetQueueName adds the queueName to the delete msg vpn queue subscription params
func (o *DeleteMsgVpnQueueSubscriptionParams) SetQueueName(queueName string) {
	o.QueueName = queueName
}

// WithSubscriptionTopic adds the subscriptionTopic to the delete msg vpn queue subscription params
func (o *DeleteMsgVpnQueueSubscriptionParams) WithSubscriptionTopic(subscriptionTopic string) *DeleteMsgVpnQueueSubscriptionParams {
	o.SetSubscriptionTopic(subscriptionTopic)
	return o
}

// SetSubscriptionTopic adds the subscriptionTopic to the delete msg vpn queue subscription params
func (o *DeleteMsgVpnQueueSubscriptionParams) SetSubscriptionTopic(subscriptionTopic string) {
	o.SubscriptionTopic = subscriptionTopic
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnQueueSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	// path param queueName
	if err := r.SetPathParam("queueName", o.QueueName); err != nil {
		return err
	}

	// path param subscriptionTopic
	if err := r.SetPathParam("subscriptionTopic", o.SubscriptionTopic); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
