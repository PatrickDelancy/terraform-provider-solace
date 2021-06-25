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
)

// NewGetMsgVpnDistributedCacheClustersParams creates a new GetMsgVpnDistributedCacheClustersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMsgVpnDistributedCacheClustersParams() *GetMsgVpnDistributedCacheClustersParams {
	return &GetMsgVpnDistributedCacheClustersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnDistributedCacheClustersParamsWithTimeout creates a new GetMsgVpnDistributedCacheClustersParams object
// with the ability to set a timeout on a request.
func NewGetMsgVpnDistributedCacheClustersParamsWithTimeout(timeout time.Duration) *GetMsgVpnDistributedCacheClustersParams {
	return &GetMsgVpnDistributedCacheClustersParams{
		timeout: timeout,
	}
}

// NewGetMsgVpnDistributedCacheClustersParamsWithContext creates a new GetMsgVpnDistributedCacheClustersParams object
// with the ability to set a context for a request.
func NewGetMsgVpnDistributedCacheClustersParamsWithContext(ctx context.Context) *GetMsgVpnDistributedCacheClustersParams {
	return &GetMsgVpnDistributedCacheClustersParams{
		Context: ctx,
	}
}

// NewGetMsgVpnDistributedCacheClustersParamsWithHTTPClient creates a new GetMsgVpnDistributedCacheClustersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMsgVpnDistributedCacheClustersParamsWithHTTPClient(client *http.Client) *GetMsgVpnDistributedCacheClustersParams {
	return &GetMsgVpnDistributedCacheClustersParams{
		HTTPClient: client,
	}
}

/* GetMsgVpnDistributedCacheClustersParams contains all the parameters to send to the API endpoint
   for the get msg vpn distributed cache clusters operation.

   Typically these are written to a http.Request.
*/
type GetMsgVpnDistributedCacheClustersParams struct {

	/* CacheName.

	   The name of the Distributed Cache.
	*/
	CacheName string

	/* Count.

	   Limit the count of objects in the response. See the documentation for the `count` parameter.

	   Default: 10
	*/
	Count *int64

	/* Cursor.

	   The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter.
	*/
	Cursor *string

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

	/* Where.

	   Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter.
	*/
	Where []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get msg vpn distributed cache clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMsgVpnDistributedCacheClustersParams) WithDefaults() *GetMsgVpnDistributedCacheClustersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get msg vpn distributed cache clusters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMsgVpnDistributedCacheClustersParams) SetDefaults() {
	var (
		countDefault = int64(10)
	)

	val := GetMsgVpnDistributedCacheClustersParams{
		Count: &countDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) WithTimeout(timeout time.Duration) *GetMsgVpnDistributedCacheClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) WithContext(ctx context.Context) *GetMsgVpnDistributedCacheClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) WithHTTPClient(client *http.Client) *GetMsgVpnDistributedCacheClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCacheName adds the cacheName to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) WithCacheName(cacheName string) *GetMsgVpnDistributedCacheClustersParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithCount adds the count to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) WithCount(count *int64) *GetMsgVpnDistributedCacheClustersParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) SetCount(count *int64) {
	o.Count = count
}

// WithCursor adds the cursor to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) WithCursor(cursor *string) *GetMsgVpnDistributedCacheClustersParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnDistributedCacheClustersParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) WithOpaquePassword(opaquePassword *string) *GetMsgVpnDistributedCacheClustersParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) WithSelect(selectVar []string) *GetMsgVpnDistributedCacheClustersParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithWhere adds the where to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) WithWhere(where []string) *GetMsgVpnDistributedCacheClustersParams {
	o.SetWhere(where)
	return o
}

// SetWhere adds the where to the get msg vpn distributed cache clusters params
func (o *GetMsgVpnDistributedCacheClustersParams) SetWhere(where []string) {
	o.Where = where
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnDistributedCacheClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cacheName
	if err := r.SetPathParam("cacheName", o.CacheName); err != nil {
		return err
	}

	if o.Count != nil {

		// query param count
		var qrCount int64

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}
	}

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string

		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {

			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
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

	if o.Where != nil {

		// binding items for where
		joinedWhere := o.bindParamWhere(reg)

		// query array param where
		if err := r.SetQueryParam("where", joinedWhere...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetMsgVpnDistributedCacheClusters binds the parameter select
func (o *GetMsgVpnDistributedCacheClustersParams) bindParamSelect(formats strfmt.Registry) []string {
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

// bindParamGetMsgVpnDistributedCacheClusters binds the parameter where
func (o *GetMsgVpnDistributedCacheClustersParams) bindParamWhere(formats strfmt.Registry) []string {
	whereIR := o.Where

	var whereIC []string
	for _, whereIIR := range whereIR { // explode []string

		whereIIV := whereIIR // string as string
		whereIC = append(whereIC, whereIIV)
	}

	// items.CollectionFormat: "csv"
	whereIS := swag.JoinByFormat(whereIC, "csv")

	return whereIS
}