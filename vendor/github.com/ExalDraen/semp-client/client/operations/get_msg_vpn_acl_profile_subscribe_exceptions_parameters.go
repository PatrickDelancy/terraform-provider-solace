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
)

// NewGetMsgVpnACLProfileSubscribeExceptionsParams creates a new GetMsgVpnACLProfileSubscribeExceptionsParams object
// with the default values initialized.
func NewGetMsgVpnACLProfileSubscribeExceptionsParams() *GetMsgVpnACLProfileSubscribeExceptionsParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnACLProfileSubscribeExceptionsParams{
		Count: &countDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnACLProfileSubscribeExceptionsParamsWithTimeout creates a new GetMsgVpnACLProfileSubscribeExceptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMsgVpnACLProfileSubscribeExceptionsParamsWithTimeout(timeout time.Duration) *GetMsgVpnACLProfileSubscribeExceptionsParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnACLProfileSubscribeExceptionsParams{
		Count: &countDefault,

		timeout: timeout,
	}
}

// NewGetMsgVpnACLProfileSubscribeExceptionsParamsWithContext creates a new GetMsgVpnACLProfileSubscribeExceptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMsgVpnACLProfileSubscribeExceptionsParamsWithContext(ctx context.Context) *GetMsgVpnACLProfileSubscribeExceptionsParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnACLProfileSubscribeExceptionsParams{
		Count: &countDefault,

		Context: ctx,
	}
}

// NewGetMsgVpnACLProfileSubscribeExceptionsParamsWithHTTPClient creates a new GetMsgVpnACLProfileSubscribeExceptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMsgVpnACLProfileSubscribeExceptionsParamsWithHTTPClient(client *http.Client) *GetMsgVpnACLProfileSubscribeExceptionsParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnACLProfileSubscribeExceptionsParams{
		Count:      &countDefault,
		HTTPClient: client,
	}
}

/*GetMsgVpnACLProfileSubscribeExceptionsParams contains all the parameters to send to the API endpoint
for the get msg vpn Acl profile subscribe exceptions operation typically these are written to a http.Request
*/
type GetMsgVpnACLProfileSubscribeExceptionsParams struct {

	/*ACLProfileName
	  The aclProfileName of the ACL Profile.

	*/
	ACLProfileName string
	/*Count
	  Limit the count of objects in the response. See [Count](#count "Description of the syntax of the `count` parameter").

	*/
	Count *int64
	/*Cursor
	  The cursor, or position, for the next page of objects. See [Cursor](#cursor "Description of the syntax of the `cursor` parameter").

	*/
	Cursor *string
	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See [Select](#select "Description of the syntax of the `select` parameter").

	*/
	Select []string
	/*Where
	  Include in the response only objects where certain conditions are true. See [Where](#where "Description of the syntax of the `where` parameter").

	*/
	Where []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) WithTimeout(timeout time.Duration) *GetMsgVpnACLProfileSubscribeExceptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) WithContext(ctx context.Context) *GetMsgVpnACLProfileSubscribeExceptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) WithHTTPClient(client *http.Client) *GetMsgVpnACLProfileSubscribeExceptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithACLProfileName adds the aCLProfileName to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) WithACLProfileName(aCLProfileName string) *GetMsgVpnACLProfileSubscribeExceptionsParams {
	o.SetACLProfileName(aCLProfileName)
	return o
}

// SetACLProfileName adds the aclProfileName to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) SetACLProfileName(aCLProfileName string) {
	o.ACLProfileName = aCLProfileName
}

// WithCount adds the count to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) WithCount(count *int64) *GetMsgVpnACLProfileSubscribeExceptionsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) SetCount(count *int64) {
	o.Count = count
}

// WithCursor adds the cursor to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) WithCursor(cursor *string) *GetMsgVpnACLProfileSubscribeExceptionsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnACLProfileSubscribeExceptionsParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) WithSelect(selectVar []string) *GetMsgVpnACLProfileSubscribeExceptionsParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithWhere adds the where to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) WithWhere(where []string) *GetMsgVpnACLProfileSubscribeExceptionsParams {
	o.SetWhere(where)
	return o
}

// SetWhere adds the where to the get msg vpn Acl profile subscribe exceptions params
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) SetWhere(where []string) {
	o.Where = where
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnACLProfileSubscribeExceptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param aclProfileName
	if err := r.SetPathParam("aclProfileName", o.ACLProfileName); err != nil {
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

	valuesSelect := o.Select

	joinedSelect := swag.JoinByFormat(valuesSelect, "csv")
	// query array param select
	if err := r.SetQueryParam("select", joinedSelect...); err != nil {
		return err
	}

	valuesWhere := o.Where

	joinedWhere := swag.JoinByFormat(valuesWhere, "csv")
	// query array param where
	if err := r.SetQueryParam("where", joinedWhere...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
