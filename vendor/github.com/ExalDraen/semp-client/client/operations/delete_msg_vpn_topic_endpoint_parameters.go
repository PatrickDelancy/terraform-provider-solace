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

// NewDeleteMsgVpnTopicEndpointParams creates a new DeleteMsgVpnTopicEndpointParams object
// with the default values initialized.
func NewDeleteMsgVpnTopicEndpointParams() *DeleteMsgVpnTopicEndpointParams {
	var ()
	return &DeleteMsgVpnTopicEndpointParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnTopicEndpointParamsWithTimeout creates a new DeleteMsgVpnTopicEndpointParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMsgVpnTopicEndpointParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnTopicEndpointParams {
	var ()
	return &DeleteMsgVpnTopicEndpointParams{

		timeout: timeout,
	}
}

// NewDeleteMsgVpnTopicEndpointParamsWithContext creates a new DeleteMsgVpnTopicEndpointParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMsgVpnTopicEndpointParamsWithContext(ctx context.Context) *DeleteMsgVpnTopicEndpointParams {
	var ()
	return &DeleteMsgVpnTopicEndpointParams{

		Context: ctx,
	}
}

// NewDeleteMsgVpnTopicEndpointParamsWithHTTPClient creates a new DeleteMsgVpnTopicEndpointParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMsgVpnTopicEndpointParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnTopicEndpointParams {
	var ()
	return &DeleteMsgVpnTopicEndpointParams{
		HTTPClient: client,
	}
}

/*DeleteMsgVpnTopicEndpointParams contains all the parameters to send to the API endpoint
for the delete msg vpn topic endpoint operation typically these are written to a http.Request
*/
type DeleteMsgVpnTopicEndpointParams struct {

	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*TopicEndpointName
	  The topicEndpointName of the Topic Endpoint.

	*/
	TopicEndpointName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete msg vpn topic endpoint params
func (o *DeleteMsgVpnTopicEndpointParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnTopicEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn topic endpoint params
func (o *DeleteMsgVpnTopicEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn topic endpoint params
func (o *DeleteMsgVpnTopicEndpointParams) WithContext(ctx context.Context) *DeleteMsgVpnTopicEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn topic endpoint params
func (o *DeleteMsgVpnTopicEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn topic endpoint params
func (o *DeleteMsgVpnTopicEndpointParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnTopicEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn topic endpoint params
func (o *DeleteMsgVpnTopicEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn topic endpoint params
func (o *DeleteMsgVpnTopicEndpointParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnTopicEndpointParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn topic endpoint params
func (o *DeleteMsgVpnTopicEndpointParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithTopicEndpointName adds the topicEndpointName to the delete msg vpn topic endpoint params
func (o *DeleteMsgVpnTopicEndpointParams) WithTopicEndpointName(topicEndpointName string) *DeleteMsgVpnTopicEndpointParams {
	o.SetTopicEndpointName(topicEndpointName)
	return o
}

// SetTopicEndpointName adds the topicEndpointName to the delete msg vpn topic endpoint params
func (o *DeleteMsgVpnTopicEndpointParams) SetTopicEndpointName(topicEndpointName string) {
	o.TopicEndpointName = topicEndpointName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnTopicEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	// path param topicEndpointName
	if err := r.SetPathParam("topicEndpointName", o.TopicEndpointName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
