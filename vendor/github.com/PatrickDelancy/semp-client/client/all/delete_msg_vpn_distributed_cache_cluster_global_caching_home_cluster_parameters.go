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
)

// NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams creates a new DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams() *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	return &DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithTimeout creates a new DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams object
// with the ability to set a timeout on a request.
func NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	return &DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams{
		timeout: timeout,
	}
}

// NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithContext creates a new DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams object
// with the ability to set a context for a request.
func NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithContext(ctx context.Context) *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	return &DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams{
		Context: ctx,
	}
}

// NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithHTTPClient creates a new DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	return &DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams{
		HTTPClient: client,
	}
}

/* DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams contains all the parameters to send to the API endpoint
   for the delete msg vpn distributed cache cluster global caching home cluster operation.

   Typically these are written to a http.Request.
*/
type DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete msg vpn distributed cache cluster global caching home cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithDefaults() *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete msg vpn distributed cache cluster global caching home cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete msg vpn distributed cache cluster global caching home cluster params
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn distributed cache cluster global caching home cluster params
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn distributed cache cluster global caching home cluster params
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithContext(ctx context.Context) *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn distributed cache cluster global caching home cluster params
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn distributed cache cluster global caching home cluster params
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn distributed cache cluster global caching home cluster params
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCacheName adds the cacheName to the delete msg vpn distributed cache cluster global caching home cluster params
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithCacheName(cacheName string) *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the delete msg vpn distributed cache cluster global caching home cluster params
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithClusterName adds the clusterName to the delete msg vpn distributed cache cluster global caching home cluster params
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithClusterName(clusterName string) *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the delete msg vpn distributed cache cluster global caching home cluster params
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithHomeClusterName adds the homeClusterName to the delete msg vpn distributed cache cluster global caching home cluster params
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithHomeClusterName(homeClusterName string) *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetHomeClusterName(homeClusterName)
	return o
}

// SetHomeClusterName adds the homeClusterName to the delete msg vpn distributed cache cluster global caching home cluster params
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetHomeClusterName(homeClusterName string) {
	o.HomeClusterName = homeClusterName
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn distributed cache cluster global caching home cluster params
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn distributed cache cluster global caching home cluster params
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
