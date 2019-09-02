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

// NewUpdateMsgVpnBridgeRemoteMsgVpnParams creates a new UpdateMsgVpnBridgeRemoteMsgVpnParams object
// with the default values initialized.
func NewUpdateMsgVpnBridgeRemoteMsgVpnParams() *UpdateMsgVpnBridgeRemoteMsgVpnParams {
	var ()
	return &UpdateMsgVpnBridgeRemoteMsgVpnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMsgVpnBridgeRemoteMsgVpnParamsWithTimeout creates a new UpdateMsgVpnBridgeRemoteMsgVpnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateMsgVpnBridgeRemoteMsgVpnParamsWithTimeout(timeout time.Duration) *UpdateMsgVpnBridgeRemoteMsgVpnParams {
	var ()
	return &UpdateMsgVpnBridgeRemoteMsgVpnParams{

		timeout: timeout,
	}
}

// NewUpdateMsgVpnBridgeRemoteMsgVpnParamsWithContext creates a new UpdateMsgVpnBridgeRemoteMsgVpnParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateMsgVpnBridgeRemoteMsgVpnParamsWithContext(ctx context.Context) *UpdateMsgVpnBridgeRemoteMsgVpnParams {
	var ()
	return &UpdateMsgVpnBridgeRemoteMsgVpnParams{

		Context: ctx,
	}
}

// NewUpdateMsgVpnBridgeRemoteMsgVpnParamsWithHTTPClient creates a new UpdateMsgVpnBridgeRemoteMsgVpnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateMsgVpnBridgeRemoteMsgVpnParamsWithHTTPClient(client *http.Client) *UpdateMsgVpnBridgeRemoteMsgVpnParams {
	var ()
	return &UpdateMsgVpnBridgeRemoteMsgVpnParams{
		HTTPClient: client,
	}
}

/*UpdateMsgVpnBridgeRemoteMsgVpnParams contains all the parameters to send to the API endpoint
for the update msg vpn bridge remote msg vpn operation typically these are written to a http.Request
*/
type UpdateMsgVpnBridgeRemoteMsgVpnParams struct {

	/*Body
	  The Remote Message VPN object's attributes.

	*/
	Body *models.MsgVpnBridgeRemoteMsgVpn
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
	/*RemoteMsgVpnInterface
	  The remoteMsgVpnInterface of the Remote Message VPN.

	*/
	RemoteMsgVpnInterface string
	/*RemoteMsgVpnLocation
	  The remoteMsgVpnLocation of the Remote Message VPN.

	*/
	RemoteMsgVpnLocation string
	/*RemoteMsgVpnName
	  The remoteMsgVpnName of the Remote Message VPN.

	*/
	RemoteMsgVpnName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See [Select](#select "Description of the syntax of the `select` parameter").

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) WithTimeout(timeout time.Duration) *UpdateMsgVpnBridgeRemoteMsgVpnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) WithContext(ctx context.Context) *UpdateMsgVpnBridgeRemoteMsgVpnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) WithHTTPClient(client *http.Client) *UpdateMsgVpnBridgeRemoteMsgVpnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) WithBody(body *models.MsgVpnBridgeRemoteMsgVpn) *UpdateMsgVpnBridgeRemoteMsgVpnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) SetBody(body *models.MsgVpnBridgeRemoteMsgVpn) {
	o.Body = body
}

// WithBridgeName adds the bridgeName to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) WithBridgeName(bridgeName string) *UpdateMsgVpnBridgeRemoteMsgVpnParams {
	o.SetBridgeName(bridgeName)
	return o
}

// SetBridgeName adds the bridgeName to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) SetBridgeName(bridgeName string) {
	o.BridgeName = bridgeName
}

// WithBridgeVirtualRouter adds the bridgeVirtualRouter to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) WithBridgeVirtualRouter(bridgeVirtualRouter string) *UpdateMsgVpnBridgeRemoteMsgVpnParams {
	o.SetBridgeVirtualRouter(bridgeVirtualRouter)
	return o
}

// SetBridgeVirtualRouter adds the bridgeVirtualRouter to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) SetBridgeVirtualRouter(bridgeVirtualRouter string) {
	o.BridgeVirtualRouter = bridgeVirtualRouter
}

// WithMsgVpnName adds the msgVpnName to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) WithMsgVpnName(msgVpnName string) *UpdateMsgVpnBridgeRemoteMsgVpnParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithRemoteMsgVpnInterface adds the remoteMsgVpnInterface to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) WithRemoteMsgVpnInterface(remoteMsgVpnInterface string) *UpdateMsgVpnBridgeRemoteMsgVpnParams {
	o.SetRemoteMsgVpnInterface(remoteMsgVpnInterface)
	return o
}

// SetRemoteMsgVpnInterface adds the remoteMsgVpnInterface to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) SetRemoteMsgVpnInterface(remoteMsgVpnInterface string) {
	o.RemoteMsgVpnInterface = remoteMsgVpnInterface
}

// WithRemoteMsgVpnLocation adds the remoteMsgVpnLocation to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) WithRemoteMsgVpnLocation(remoteMsgVpnLocation string) *UpdateMsgVpnBridgeRemoteMsgVpnParams {
	o.SetRemoteMsgVpnLocation(remoteMsgVpnLocation)
	return o
}

// SetRemoteMsgVpnLocation adds the remoteMsgVpnLocation to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) SetRemoteMsgVpnLocation(remoteMsgVpnLocation string) {
	o.RemoteMsgVpnLocation = remoteMsgVpnLocation
}

// WithRemoteMsgVpnName adds the remoteMsgVpnName to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) WithRemoteMsgVpnName(remoteMsgVpnName string) *UpdateMsgVpnBridgeRemoteMsgVpnParams {
	o.SetRemoteMsgVpnName(remoteMsgVpnName)
	return o
}

// SetRemoteMsgVpnName adds the remoteMsgVpnName to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) SetRemoteMsgVpnName(remoteMsgVpnName string) {
	o.RemoteMsgVpnName = remoteMsgVpnName
}

// WithSelect adds the selectVar to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) WithSelect(selectVar []string) *UpdateMsgVpnBridgeRemoteMsgVpnParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the update msg vpn bridge remote msg vpn params
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMsgVpnBridgeRemoteMsgVpnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param remoteMsgVpnInterface
	if err := r.SetPathParam("remoteMsgVpnInterface", o.RemoteMsgVpnInterface); err != nil {
		return err
	}

	// path param remoteMsgVpnLocation
	if err := r.SetPathParam("remoteMsgVpnLocation", o.RemoteMsgVpnLocation); err != nil {
		return err
	}

	// path param remoteMsgVpnName
	if err := r.SetPathParam("remoteMsgVpnName", o.RemoteMsgVpnName); err != nil {
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
