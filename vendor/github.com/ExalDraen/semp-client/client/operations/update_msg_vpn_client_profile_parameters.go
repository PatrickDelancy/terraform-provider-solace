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

// NewUpdateMsgVpnClientProfileParams creates a new UpdateMsgVpnClientProfileParams object
// with the default values initialized.
func NewUpdateMsgVpnClientProfileParams() *UpdateMsgVpnClientProfileParams {
	var ()
	return &UpdateMsgVpnClientProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMsgVpnClientProfileParamsWithTimeout creates a new UpdateMsgVpnClientProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateMsgVpnClientProfileParamsWithTimeout(timeout time.Duration) *UpdateMsgVpnClientProfileParams {
	var ()
	return &UpdateMsgVpnClientProfileParams{

		timeout: timeout,
	}
}

// NewUpdateMsgVpnClientProfileParamsWithContext creates a new UpdateMsgVpnClientProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateMsgVpnClientProfileParamsWithContext(ctx context.Context) *UpdateMsgVpnClientProfileParams {
	var ()
	return &UpdateMsgVpnClientProfileParams{

		Context: ctx,
	}
}

// NewUpdateMsgVpnClientProfileParamsWithHTTPClient creates a new UpdateMsgVpnClientProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateMsgVpnClientProfileParamsWithHTTPClient(client *http.Client) *UpdateMsgVpnClientProfileParams {
	var ()
	return &UpdateMsgVpnClientProfileParams{
		HTTPClient: client,
	}
}

/*UpdateMsgVpnClientProfileParams contains all the parameters to send to the API endpoint
for the update msg vpn client profile operation typically these are written to a http.Request
*/
type UpdateMsgVpnClientProfileParams struct {

	/*Body
	  The Client Profile object's attributes.

	*/
	Body *models.MsgVpnClientProfile
	/*ClientProfileName
	  The clientProfileName of the Client Profile.

	*/
	ClientProfileName string
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

// WithTimeout adds the timeout to the update msg vpn client profile params
func (o *UpdateMsgVpnClientProfileParams) WithTimeout(timeout time.Duration) *UpdateMsgVpnClientProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update msg vpn client profile params
func (o *UpdateMsgVpnClientProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update msg vpn client profile params
func (o *UpdateMsgVpnClientProfileParams) WithContext(ctx context.Context) *UpdateMsgVpnClientProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update msg vpn client profile params
func (o *UpdateMsgVpnClientProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update msg vpn client profile params
func (o *UpdateMsgVpnClientProfileParams) WithHTTPClient(client *http.Client) *UpdateMsgVpnClientProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update msg vpn client profile params
func (o *UpdateMsgVpnClientProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update msg vpn client profile params
func (o *UpdateMsgVpnClientProfileParams) WithBody(body *models.MsgVpnClientProfile) *UpdateMsgVpnClientProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update msg vpn client profile params
func (o *UpdateMsgVpnClientProfileParams) SetBody(body *models.MsgVpnClientProfile) {
	o.Body = body
}

// WithClientProfileName adds the clientProfileName to the update msg vpn client profile params
func (o *UpdateMsgVpnClientProfileParams) WithClientProfileName(clientProfileName string) *UpdateMsgVpnClientProfileParams {
	o.SetClientProfileName(clientProfileName)
	return o
}

// SetClientProfileName adds the clientProfileName to the update msg vpn client profile params
func (o *UpdateMsgVpnClientProfileParams) SetClientProfileName(clientProfileName string) {
	o.ClientProfileName = clientProfileName
}

// WithMsgVpnName adds the msgVpnName to the update msg vpn client profile params
func (o *UpdateMsgVpnClientProfileParams) WithMsgVpnName(msgVpnName string) *UpdateMsgVpnClientProfileParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the update msg vpn client profile params
func (o *UpdateMsgVpnClientProfileParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the update msg vpn client profile params
func (o *UpdateMsgVpnClientProfileParams) WithSelect(selectVar []string) *UpdateMsgVpnClientProfileParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the update msg vpn client profile params
func (o *UpdateMsgVpnClientProfileParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMsgVpnClientProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param clientProfileName
	if err := r.SetPathParam("clientProfileName", o.ClientProfileName); err != nil {
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
