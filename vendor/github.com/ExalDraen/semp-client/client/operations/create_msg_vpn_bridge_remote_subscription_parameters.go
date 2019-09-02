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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ExalDraen/semp-client/models"
)

// NewCreateMsgVpnBridgeRemoteSubscriptionParams creates a new CreateMsgVpnBridgeRemoteSubscriptionParams object
// with the default values initialized.
func NewCreateMsgVpnBridgeRemoteSubscriptionParams() *CreateMsgVpnBridgeRemoteSubscriptionParams {
	var ()
	return &CreateMsgVpnBridgeRemoteSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnBridgeRemoteSubscriptionParamsWithTimeout creates a new CreateMsgVpnBridgeRemoteSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMsgVpnBridgeRemoteSubscriptionParamsWithTimeout(timeout time.Duration) *CreateMsgVpnBridgeRemoteSubscriptionParams {
	var ()
	return &CreateMsgVpnBridgeRemoteSubscriptionParams{

		timeout: timeout,
	}
}

// NewCreateMsgVpnBridgeRemoteSubscriptionParamsWithContext creates a new CreateMsgVpnBridgeRemoteSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMsgVpnBridgeRemoteSubscriptionParamsWithContext(ctx context.Context) *CreateMsgVpnBridgeRemoteSubscriptionParams {
	var ()
	return &CreateMsgVpnBridgeRemoteSubscriptionParams{

		Context: ctx,
	}
}

// NewCreateMsgVpnBridgeRemoteSubscriptionParamsWithHTTPClient creates a new CreateMsgVpnBridgeRemoteSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMsgVpnBridgeRemoteSubscriptionParamsWithHTTPClient(client *http.Client) *CreateMsgVpnBridgeRemoteSubscriptionParams {
	var ()
	return &CreateMsgVpnBridgeRemoteSubscriptionParams{
		HTTPClient: client,
	}
}

/*CreateMsgVpnBridgeRemoteSubscriptionParams contains all the parameters to send to the API endpoint
for the create msg vpn bridge remote subscription operation typically these are written to a http.Request
*/
type CreateMsgVpnBridgeRemoteSubscriptionParams struct {

	/*Body
	  The Remote Subscription object's attributes.

	*/
	Body *models.MsgVpnBridgeRemoteSubscription
	/*BridgeName
	  The bridgeName of the Bridge.

	*/
	BridgeName string
	/*BridgeVirtualRouter
	  The bridgeVirtualRouter of the Bridge.

	*/
	BridgeVirtualRouter string
	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See [Select](#select "Description of the syntax of the `select` parameter").

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) WithTimeout(timeout time.Duration) *CreateMsgVpnBridgeRemoteSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) WithContext(ctx context.Context) *CreateMsgVpnBridgeRemoteSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) WithHTTPClient(client *http.Client) *CreateMsgVpnBridgeRemoteSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) WithBody(body *models.MsgVpnBridgeRemoteSubscription) *CreateMsgVpnBridgeRemoteSubscriptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) SetBody(body *models.MsgVpnBridgeRemoteSubscription) {
	o.Body = body
}

// WithBridgeName adds the bridgeName to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) WithBridgeName(bridgeName string) *CreateMsgVpnBridgeRemoteSubscriptionParams {
	o.SetBridgeName(bridgeName)
	return o
}

// SetBridgeName adds the bridgeName to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) SetBridgeName(bridgeName string) {
	o.BridgeName = bridgeName
}

// WithBridgeVirtualRouter adds the bridgeVirtualRouter to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) WithBridgeVirtualRouter(bridgeVirtualRouter string) *CreateMsgVpnBridgeRemoteSubscriptionParams {
	o.SetBridgeVirtualRouter(bridgeVirtualRouter)
	return o
}

// SetBridgeVirtualRouter adds the bridgeVirtualRouter to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) SetBridgeVirtualRouter(bridgeVirtualRouter string) {
	o.BridgeVirtualRouter = bridgeVirtualRouter
}

// WithMsgVpnName adds the msgVpnName to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) WithMsgVpnName(msgVpnName string) *CreateMsgVpnBridgeRemoteSubscriptionParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) WithSelect(selectVar []string) *CreateMsgVpnBridgeRemoteSubscriptionParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn bridge remote subscription params
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnBridgeRemoteSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param bridgeName
	if err := r.SetPathParam("bridgeName", o.BridgeName); err != nil {
		return err
	}

	// path param bridgeVirtualRouter
	if err := r.SetPathParam("bridgeVirtualRouter", o.BridgeVirtualRouter); err != nil {
		return err
	}

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	valuesSelect := o.Select

	joinedSelect := swag.JoinByFormat(valuesSelect, "csv")
	// query array param select
	if err := r.SetQueryParam("select", joinedSelect...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
