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

// NewReplaceMsgVpnRestDeliveryPointRestConsumerParams creates a new ReplaceMsgVpnRestDeliveryPointRestConsumerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReplaceMsgVpnRestDeliveryPointRestConsumerParams() *ReplaceMsgVpnRestDeliveryPointRestConsumerParams {
	return &ReplaceMsgVpnRestDeliveryPointRestConsumerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMsgVpnRestDeliveryPointRestConsumerParamsWithTimeout creates a new ReplaceMsgVpnRestDeliveryPointRestConsumerParams object
// with the ability to set a timeout on a request.
func NewReplaceMsgVpnRestDeliveryPointRestConsumerParamsWithTimeout(timeout time.Duration) *ReplaceMsgVpnRestDeliveryPointRestConsumerParams {
	return &ReplaceMsgVpnRestDeliveryPointRestConsumerParams{
		timeout: timeout,
	}
}

// NewReplaceMsgVpnRestDeliveryPointRestConsumerParamsWithContext creates a new ReplaceMsgVpnRestDeliveryPointRestConsumerParams object
// with the ability to set a context for a request.
func NewReplaceMsgVpnRestDeliveryPointRestConsumerParamsWithContext(ctx context.Context) *ReplaceMsgVpnRestDeliveryPointRestConsumerParams {
	return &ReplaceMsgVpnRestDeliveryPointRestConsumerParams{
		Context: ctx,
	}
}

// NewReplaceMsgVpnRestDeliveryPointRestConsumerParamsWithHTTPClient creates a new ReplaceMsgVpnRestDeliveryPointRestConsumerParams object
// with the ability to set a custom HTTPClient for a request.
func NewReplaceMsgVpnRestDeliveryPointRestConsumerParamsWithHTTPClient(client *http.Client) *ReplaceMsgVpnRestDeliveryPointRestConsumerParams {
	return &ReplaceMsgVpnRestDeliveryPointRestConsumerParams{
		HTTPClient: client,
	}
}

/* ReplaceMsgVpnRestDeliveryPointRestConsumerParams contains all the parameters to send to the API endpoint
   for the replace msg vpn rest delivery point rest consumer operation.

   Typically these are written to a http.Request.
*/
type ReplaceMsgVpnRestDeliveryPointRestConsumerParams struct {

	/* Body.

	   The REST Consumer object's attributes.
	*/
	Body *models.MsgVpnRestDeliveryPointRestConsumer

	/* MsgVpnName.

	   The name of the Message VPN.
	*/
	MsgVpnName string

	/* OpaquePassword.

	   Accept opaque attributes in the request or return opaque attributes in the response, encrypted with the specified password. See that documentation for the `opaquePassword` parameter.
	*/
	OpaquePassword *string

	/* RestConsumerName.

	   The name of the REST Consumer.
	*/
	RestConsumerName string

	/* RestDeliveryPointName.

	   The name of the REST Delivery Point.
	*/
	RestDeliveryPointName string

	/* Select.

	   Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.
	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the replace msg vpn rest delivery point rest consumer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) WithDefaults() *ReplaceMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the replace msg vpn rest delivery point rest consumer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) WithTimeout(timeout time.Duration) *ReplaceMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) WithContext(ctx context.Context) *ReplaceMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) WithHTTPClient(client *http.Client) *ReplaceMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) WithBody(body *models.MsgVpnRestDeliveryPointRestConsumer) *ReplaceMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) SetBody(body *models.MsgVpnRestDeliveryPointRestConsumer) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) WithMsgVpnName(msgVpnName string) *ReplaceMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) WithOpaquePassword(opaquePassword *string) *ReplaceMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithRestConsumerName adds the restConsumerName to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) WithRestConsumerName(restConsumerName string) *ReplaceMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetRestConsumerName(restConsumerName)
	return o
}

// SetRestConsumerName adds the restConsumerName to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) SetRestConsumerName(restConsumerName string) {
	o.RestConsumerName = restConsumerName
}

// WithRestDeliveryPointName adds the restDeliveryPointName to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) WithRestDeliveryPointName(restDeliveryPointName string) *ReplaceMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetRestDeliveryPointName(restDeliveryPointName)
	return o
}

// SetRestDeliveryPointName adds the restDeliveryPointName to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) SetRestDeliveryPointName(restDeliveryPointName string) {
	o.RestDeliveryPointName = restDeliveryPointName
}

// WithSelect adds the selectVar to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) WithSelect(selectVar []string) *ReplaceMsgVpnRestDeliveryPointRestConsumerParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace msg vpn rest delivery point rest consumer params
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param restConsumerName
	if err := r.SetPathParam("restConsumerName", o.RestConsumerName); err != nil {
		return err
	}

	// path param restDeliveryPointName
	if err := r.SetPathParam("restDeliveryPointName", o.RestDeliveryPointName); err != nil {
		return err
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

// bindParamReplaceMsgVpnRestDeliveryPointRestConsumer binds the parameter select
func (o *ReplaceMsgVpnRestDeliveryPointRestConsumerParams) bindParamSelect(formats strfmt.Registry) []string {
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