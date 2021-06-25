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

// NewUpdateMsgVpnClientUsernameParams creates a new UpdateMsgVpnClientUsernameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMsgVpnClientUsernameParams() *UpdateMsgVpnClientUsernameParams {
	return &UpdateMsgVpnClientUsernameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMsgVpnClientUsernameParamsWithTimeout creates a new UpdateMsgVpnClientUsernameParams object
// with the ability to set a timeout on a request.
func NewUpdateMsgVpnClientUsernameParamsWithTimeout(timeout time.Duration) *UpdateMsgVpnClientUsernameParams {
	return &UpdateMsgVpnClientUsernameParams{
		timeout: timeout,
	}
}

// NewUpdateMsgVpnClientUsernameParamsWithContext creates a new UpdateMsgVpnClientUsernameParams object
// with the ability to set a context for a request.
func NewUpdateMsgVpnClientUsernameParamsWithContext(ctx context.Context) *UpdateMsgVpnClientUsernameParams {
	return &UpdateMsgVpnClientUsernameParams{
		Context: ctx,
	}
}

// NewUpdateMsgVpnClientUsernameParamsWithHTTPClient creates a new UpdateMsgVpnClientUsernameParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMsgVpnClientUsernameParamsWithHTTPClient(client *http.Client) *UpdateMsgVpnClientUsernameParams {
	return &UpdateMsgVpnClientUsernameParams{
		HTTPClient: client,
	}
}

/* UpdateMsgVpnClientUsernameParams contains all the parameters to send to the API endpoint
   for the update msg vpn client username operation.

   Typically these are written to a http.Request.
*/
type UpdateMsgVpnClientUsernameParams struct {

	/* Body.

	   The Client Username object's attributes.
	*/
	Body *models.MsgVpnClientUsername

	/* ClientUsername.

	   The name of the Client Username.
	*/
	ClientUsername string

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

// WithDefaults hydrates default values in the update msg vpn client username params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMsgVpnClientUsernameParams) WithDefaults() *UpdateMsgVpnClientUsernameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update msg vpn client username params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMsgVpnClientUsernameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) WithTimeout(timeout time.Duration) *UpdateMsgVpnClientUsernameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) WithContext(ctx context.Context) *UpdateMsgVpnClientUsernameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) WithHTTPClient(client *http.Client) *UpdateMsgVpnClientUsernameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) WithBody(body *models.MsgVpnClientUsername) *UpdateMsgVpnClientUsernameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) SetBody(body *models.MsgVpnClientUsername) {
	o.Body = body
}

// WithClientUsername adds the clientUsername to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) WithClientUsername(clientUsername string) *UpdateMsgVpnClientUsernameParams {
	o.SetClientUsername(clientUsername)
	return o
}

// SetClientUsername adds the clientUsername to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) SetClientUsername(clientUsername string) {
	o.ClientUsername = clientUsername
}

// WithMsgVpnName adds the msgVpnName to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) WithMsgVpnName(msgVpnName string) *UpdateMsgVpnClientUsernameParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) WithOpaquePassword(opaquePassword *string) *UpdateMsgVpnClientUsernameParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) WithSelect(selectVar []string) *UpdateMsgVpnClientUsernameParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the update msg vpn client username params
func (o *UpdateMsgVpnClientUsernameParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMsgVpnClientUsernameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param clientUsername
	if err := r.SetPathParam("clientUsername", o.ClientUsername); err != nil {
		return err
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

// bindParamUpdateMsgVpnClientUsername binds the parameter select
func (o *UpdateMsgVpnClientUsernameParams) bindParamSelect(formats strfmt.Registry) []string {
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