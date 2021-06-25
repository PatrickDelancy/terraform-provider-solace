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

// NewUpdateDomainCertAuthorityParams creates a new UpdateDomainCertAuthorityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDomainCertAuthorityParams() *UpdateDomainCertAuthorityParams {
	return &UpdateDomainCertAuthorityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDomainCertAuthorityParamsWithTimeout creates a new UpdateDomainCertAuthorityParams object
// with the ability to set a timeout on a request.
func NewUpdateDomainCertAuthorityParamsWithTimeout(timeout time.Duration) *UpdateDomainCertAuthorityParams {
	return &UpdateDomainCertAuthorityParams{
		timeout: timeout,
	}
}

// NewUpdateDomainCertAuthorityParamsWithContext creates a new UpdateDomainCertAuthorityParams object
// with the ability to set a context for a request.
func NewUpdateDomainCertAuthorityParamsWithContext(ctx context.Context) *UpdateDomainCertAuthorityParams {
	return &UpdateDomainCertAuthorityParams{
		Context: ctx,
	}
}

// NewUpdateDomainCertAuthorityParamsWithHTTPClient creates a new UpdateDomainCertAuthorityParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDomainCertAuthorityParamsWithHTTPClient(client *http.Client) *UpdateDomainCertAuthorityParams {
	return &UpdateDomainCertAuthorityParams{
		HTTPClient: client,
	}
}

/* UpdateDomainCertAuthorityParams contains all the parameters to send to the API endpoint
   for the update domain cert authority operation.

   Typically these are written to a http.Request.
*/
type UpdateDomainCertAuthorityParams struct {

	/* Body.

	   The Domain Certificate Authority object's attributes.
	*/
	Body *models.DomainCertAuthority

	/* CertAuthorityName.

	   The name of the Certificate Authority.
	*/
	CertAuthorityName string

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

// WithDefaults hydrates default values in the update domain cert authority params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDomainCertAuthorityParams) WithDefaults() *UpdateDomainCertAuthorityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update domain cert authority params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDomainCertAuthorityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update domain cert authority params
func (o *UpdateDomainCertAuthorityParams) WithTimeout(timeout time.Duration) *UpdateDomainCertAuthorityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update domain cert authority params
func (o *UpdateDomainCertAuthorityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update domain cert authority params
func (o *UpdateDomainCertAuthorityParams) WithContext(ctx context.Context) *UpdateDomainCertAuthorityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update domain cert authority params
func (o *UpdateDomainCertAuthorityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update domain cert authority params
func (o *UpdateDomainCertAuthorityParams) WithHTTPClient(client *http.Client) *UpdateDomainCertAuthorityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update domain cert authority params
func (o *UpdateDomainCertAuthorityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update domain cert authority params
func (o *UpdateDomainCertAuthorityParams) WithBody(body *models.DomainCertAuthority) *UpdateDomainCertAuthorityParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update domain cert authority params
func (o *UpdateDomainCertAuthorityParams) SetBody(body *models.DomainCertAuthority) {
	o.Body = body
}

// WithCertAuthorityName adds the certAuthorityName to the update domain cert authority params
func (o *UpdateDomainCertAuthorityParams) WithCertAuthorityName(certAuthorityName string) *UpdateDomainCertAuthorityParams {
	o.SetCertAuthorityName(certAuthorityName)
	return o
}

// SetCertAuthorityName adds the certAuthorityName to the update domain cert authority params
func (o *UpdateDomainCertAuthorityParams) SetCertAuthorityName(certAuthorityName string) {
	o.CertAuthorityName = certAuthorityName
}

// WithOpaquePassword adds the opaquePassword to the update domain cert authority params
func (o *UpdateDomainCertAuthorityParams) WithOpaquePassword(opaquePassword *string) *UpdateDomainCertAuthorityParams {
	o.SetOpaquePassword(opaquePassword)
	return o
}

// SetOpaquePassword adds the opaquePassword to the update domain cert authority params
func (o *UpdateDomainCertAuthorityParams) SetOpaquePassword(opaquePassword *string) {
	o.OpaquePassword = opaquePassword
}

// WithSelect adds the selectVar to the update domain cert authority params
func (o *UpdateDomainCertAuthorityParams) WithSelect(selectVar []string) *UpdateDomainCertAuthorityParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the update domain cert authority params
func (o *UpdateDomainCertAuthorityParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDomainCertAuthorityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param certAuthorityName
	if err := r.SetPathParam("certAuthorityName", o.CertAuthorityName); err != nil {
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

// bindParamUpdateDomainCertAuthority binds the parameter select
func (o *UpdateDomainCertAuthorityParams) bindParamSelect(formats strfmt.Registry) []string {
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