// Code generated by go-swagger; DO NOT EDIT.

package all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/PatrickDelancy/semp-client/models"
)

// NewCreateMsgVpnBridgeParams creates a new CreateMsgVpnBridgeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateMsgVpnBridgeParams() *CreateMsgVpnBridgeParams {
	return &CreateMsgVpnBridgeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnBridgeParamsWithTimeout creates a new CreateMsgVpnBridgeParams object
// with the ability to set a timeout on a request.
func NewCreateMsgVpnBridgeParamsWithTimeout(timeout time.Duration) *CreateMsgVpnBridgeParams {
	return &CreateMsgVpnBridgeParams{
		timeout: timeout,
	}
}

// NewCreateMsgVpnBridgeParamsWithContext creates a new CreateMsgVpnBridgeParams object
// with the ability to set a context for a request.
func NewCreateMsgVpnBridgeParamsWithContext(ctx context.Context) *CreateMsgVpnBridgeParams {
	return &CreateMsgVpnBridgeParams{
		Context: ctx,
	}
}

// NewCreateMsgVpnBridgeParamsWithHTTPClient creates a new CreateMsgVpnBridgeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateMsgVpnBridgeParamsWithHTTPClient(client *http.Client) *CreateMsgVpnBridgeParams {
	return &CreateMsgVpnBridgeParams{
		HTTPClient: client,
	}
}

/* CreateMsgVpnBridgeParams contains all the parameters to send to the API endpoint
   for the create msg vpn bridge operation.

   Typically these are written to a http.Request.
*/
type CreateMsgVpnBridgeParams struct {

	/* Body.

	   The Bridge object's attributes.
	*/
	Body *models.MsgVpnBridge

	/* MsgVpnName.

	   The name of the Message VPN.
	*/
	MsgVpnName string

	/* OpaquePassword.

	   Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter.
	*/
	OpaquePassword *string

	/* Select.

	   Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.
	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create msg vpn bridge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMsgVpnBridgeParams) WithDefaults() *CreateMsgVpnBridgeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create msg vpn bridge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMsgVpnBridgeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create msg vpn bridge params
func (o *CreateMsgVpnBridgeParams) WithTimeout(timeout time.Duration) *CreateMsgVpnBridgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn bridge params
func (o *CreateMsgVpnBridgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn bridge params
func (o *CreateMsgVpnBridgeParams) WithContext(ctx context.Context) *CreateMsgVpnBridgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn bridge params
func (o *CreateMsgVpnBridgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn bridge params
func (o *CreateMsgVpnBridgeParams) WithHTTPClient(client *http.Client) *CreateMsgVpnBridgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn bridge params
func (o *CreateMsgVpnBridgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create msg vpn bridge params
func (o *CreateMsgVpnBridgeParams) WithBody(body *models.MsgVpnBridge) *CreateMsgVpnBridgeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn bridge params
func (o *CreateMsgVpnBridgeParams) SetBody(body *models.MsgVpnBridge) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the create msg vpn bridge params
func (o *CreateMsgVpnBridgeParams) WithMsgVpnName(msgVpnName string) *CreateMsgVpnBridgeParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the create msg vpn bridge params
func (o *CreateMsgVpnBridgeParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the create msg vpn bridge params
func (o *CreateMsgVpnBridgeParams) WithOpaquePassword(opaquePassword *string) *CreateMsgVpnBridgeParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the create msg vpn bridge params
func (o *CreateMsgVpnBridgeParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the create msg vpn bridge params
func (o *CreateMsgVpnBridgeParams) WithSelect(selectVar []string) *CreateMsgVpnBridgeParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn bridge params
func (o *CreateMsgVpnBridgeParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnBridgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	if o.OpaquePassword != nil {

		// query param opaquePassword
		var qrOpaquePassword string

		if o.OpaquePassword != nil {
			qrOpaquePassword = *o.OpaquePassword
		}
		qOpaquePassword := qrOpaquePassword
		if qOpaquePassword != "" {

			if err := r.SetQueryParam("opaquePassword", qOpaquePassword); err != nil {
				return err
			}
		}
	}

	if o.Select != nil {

		// binding items for select
		joinedSelect := o.bindParamSelect(reg)

		// query array param select
		if err := r.SetQueryParam("select", joinedSelect...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCreateMsgVpnBridge binds the parameter select
func (o *CreateMsgVpnBridgeParams) bindParamSelect(formats strfmt.Registry) []string {
	selectIR := o.Select

	var selectIC []string
	for _, selectIIR := range selectIR { // explode []string

		selectIIV := selectIIR // string as string
		selectIC = append(selectIC, selectIIV)
	}

	// items.CollectionFormat: "csv"
	selectIS := swag.JoinByFormat(selectIC, "csv")

	return selectIS
}