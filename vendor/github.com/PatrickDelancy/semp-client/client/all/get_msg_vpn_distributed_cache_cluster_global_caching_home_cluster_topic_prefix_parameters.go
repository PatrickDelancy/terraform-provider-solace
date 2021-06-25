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

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams creates a new GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams() *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParamsWithTimeout creates a new GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams object
// with the ability to set a timeout on a request.
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParamsWithTimeout(timeout time.Duration) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams{
		timeout: timeout,
	}
}

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParamsWithContext creates a new GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams object
// with the ability to set a context for a request.
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParamsWithContext(ctx context.Context) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams{
		Context: ctx,
	}
}

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParamsWithHTTPClient creates a new GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParamsWithHTTPClient(client *http.Client) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams{
		HTTPClient: client,
	}
}

/* GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams contains all the parameters to send to the API endpoint
   for the get msg vpn distributed cache cluster global caching home cluster topic prefix operation.

   Typically these are written to a http.Request.
*/
type GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams struct {

	/* CacheName.

	   The name of the Distributed Cache.
	*/
	CacheName string

	/* ClusterName.

	   The name of the Cache Cluster.
	*/
	ClusterName string

	/* HomeClusterName.

	   The name of the remote Home Cache Cluster.
	*/
	HomeClusterName string

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

	/* TopicPrefix.

	   A topic prefix for global topics available from the remote Home Cache Cluster. A wildcard (/>) is implied at the end of the prefix.
	*/
	TopicPrefix string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get msg vpn distributed cache cluster global caching home cluster topic prefix params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithDefaults() *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get msg vpn distributed cache cluster global caching home cluster topic prefix params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithTimeout(timeout time.Duration) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithContext(ctx context.Context) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithHTTPClient(client *http.Client) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCacheName adds the cacheName to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithCacheName(cacheName string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithClusterName adds the clusterName to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithClusterName(clusterName string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithHomeClusterName adds the homeClusterName to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithHomeClusterName(homeClusterName string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetHomeClusterName(homeClusterName)
	return o
}

// SetHomeClusterName adds the homeClusterName to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetHomeClusterName(homeClusterName string) {
	o.HomeClusterName = homeClusterName
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithOpaquePassword adds the opaquePassword to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithOpaquePassword(opaquePassword *string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithSelect(selectVar []string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithTopicPrefix adds the topicPrefix to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WithTopicPrefix(topicPrefix string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams {
	o.SetTopicPrefix(topicPrefix)
	return o
}

// SetTopicPrefix adds the topicPrefix to the get msg vpn distributed cache cluster global caching home cluster topic prefix params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) SetTopicPrefix(topicPrefix string) {
	o.TopicPrefix = topicPrefix
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cacheName
	if err := r.SetPathParam("cacheName", o.CacheName); err != nil {
		return err
	}

	// path param clusterName
	if err := r.SetPathParam("clusterName", o.ClusterName); err != nil {
		return err
	}

	// path param homeClusterName
	if err := r.SetPathParam("homeClusterName", o.HomeClusterName); err != nil {
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

	// path param topicPrefix
	if err := r.SetPathParam("topicPrefix", o.TopicPrefix); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix binds the parameter select
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixParams) bindParamSelect(formats strfmt.Registry) []string {
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