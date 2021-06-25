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

// NewUpdateMsgVpnDistributedCacheClusterInstanceParams creates a new UpdateMsgVpnDistributedCacheClusterInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateMsgVpnDistributedCacheClusterInstanceParams() *UpdateMsgVpnDistributedCacheClusterInstanceParams {
	return &UpdateMsgVpnDistributedCacheClusterInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMsgVpnDistributedCacheClusterInstanceParamsWithTimeout creates a new UpdateMsgVpnDistributedCacheClusterInstanceParams object
// with the ability to set a timeout on a request.
func NewUpdateMsgVpnDistributedCacheClusterInstanceParamsWithTimeout(timeout time.Duration) *UpdateMsgVpnDistributedCacheClusterInstanceParams {
	return &UpdateMsgVpnDistributedCacheClusterInstanceParams{
		timeout: timeout,
	}
}

// NewUpdateMsgVpnDistributedCacheClusterInstanceParamsWithContext creates a new UpdateMsgVpnDistributedCacheClusterInstanceParams object
// with the ability to set a context for a request.
func NewUpdateMsgVpnDistributedCacheClusterInstanceParamsWithContext(ctx context.Context) *UpdateMsgVpnDistributedCacheClusterInstanceParams {
	return &UpdateMsgVpnDistributedCacheClusterInstanceParams{
		Context: ctx,
	}
}

// NewUpdateMsgVpnDistributedCacheClusterInstanceParamsWithHTTPClient creates a new UpdateMsgVpnDistributedCacheClusterInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateMsgVpnDistributedCacheClusterInstanceParamsWithHTTPClient(client *http.Client) *UpdateMsgVpnDistributedCacheClusterInstanceParams {
	return &UpdateMsgVpnDistributedCacheClusterInstanceParams{
		HTTPClient: client,
	}
}

/* UpdateMsgVpnDistributedCacheClusterInstanceParams contains all the parameters to send to the API endpoint
   for the update msg vpn distributed cache cluster instance operation.

   Typically these are written to a http.Request.
*/
type UpdateMsgVpnDistributedCacheClusterInstanceParams struct {

	/* Body.

	   The Cache Instance object's attributes.
	*/
	Body *models.MsgVpnDistributedCacheClusterInstance

	/* CacheName.

	   The name of the Distributed Cache.
	*/
	CacheName string

	/* ClusterName.

	   The name of the Cache Cluster.
	*/
	ClusterName string

	/* InstanceName.

	   The name of the Cache Instance.
	*/
	InstanceName string

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

// WithDefaults hydrates default values in the update msg vpn distributed cache cluster instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) WithDefaults() *UpdateMsgVpnDistributedCacheClusterInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update msg vpn distributed cache cluster instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) WithTimeout(timeout time.Duration) *UpdateMsgVpnDistributedCacheClusterInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) WithContext(ctx context.Context) *UpdateMsgVpnDistributedCacheClusterInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) WithHTTPClient(client *http.Client) *UpdateMsgVpnDistributedCacheClusterInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) WithBody(body *models.MsgVpnDistributedCacheClusterInstance) *UpdateMsgVpnDistributedCacheClusterInstanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) SetBody(body *models.MsgVpnDistributedCacheClusterInstance) {
	o.Body = body
}

// WithCacheName adds the cacheName to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) WithCacheName(cacheName string) *UpdateMsgVpnDistributedCacheClusterInstanceParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithClusterName adds the clusterName to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) WithClusterName(clusterName string) *UpdateMsgVpnDistributedCacheClusterInstanceParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithInstanceName adds the instanceName to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) WithInstanceName(instanceName string) *UpdateMsgVpnDistributedCacheClusterInstanceParams {
	o.SetInstanceName(instanceName)
	return o
}

// SetInstanceName adds the instanceName to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) SetInstanceName(instanceName string) {
	o.InstanceName = instanceName
}

// WithMsgVpnName adds the msgVpnName to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) WithMsgVpnName(msgVpnName string) *UpdateMsgVpnDistributedCacheClusterInstanceParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) WithOpaquePassword(opaquePassword *string) *UpdateMsgVpnDistributedCacheClusterInstanceParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) WithSelect(selectVar []string) *UpdateMsgVpnDistributedCacheClusterInstanceParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the update msg vpn distributed cache cluster instance params
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cacheName
	if err := r.SetPathParam("cacheName", o.CacheName); err != nil {
		return err
	}

	// path param clusterName
	if err := r.SetPathParam("clusterName", o.ClusterName); err != nil {
		return err
	}

	// path param instanceName
	if err := r.SetPathParam("instanceName", o.InstanceName); err != nil {
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

// bindParamUpdateMsgVpnDistributedCacheClusterInstance binds the parameter select
func (o *UpdateMsgVpnDistributedCacheClusterInstanceParams) bindParamSelect(formats strfmt.Registry) []string {
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