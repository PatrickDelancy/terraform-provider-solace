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

// NewDeleteClientCertAuthorityOcspTLSTrustedCommonNameParams creates a new DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteClientCertAuthorityOcspTLSTrustedCommonNameParams() *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams {
	return &DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClientCertAuthorityOcspTLSTrustedCommonNameParamsWithTimeout creates a new DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams object
// with the ability to set a timeout on a request.
func NewDeleteClientCertAuthorityOcspTLSTrustedCommonNameParamsWithTimeout(timeout time.Duration) *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams {
	return &DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams{
		timeout: timeout,
	}
}

// NewDeleteClientCertAuthorityOcspTLSTrustedCommonNameParamsWithContext creates a new DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams object
// with the ability to set a context for a request.
func NewDeleteClientCertAuthorityOcspTLSTrustedCommonNameParamsWithContext(ctx context.Context) *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams {
	return &DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams{
		Context: ctx,
	}
}

// NewDeleteClientCertAuthorityOcspTLSTrustedCommonNameParamsWithHTTPClient creates a new DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteClientCertAuthorityOcspTLSTrustedCommonNameParamsWithHTTPClient(client *http.Client) *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams {
	return &DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams{
		HTTPClient: client,
	}
}

/* DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams contains all the parameters to send to the API endpoint
   for the delete client cert authority ocsp Tls trusted common name operation.

   Typically these are written to a http.Request.
*/
type DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams struct {

	/* CertAuthorityName.

	   The name of the Certificate Authority.
	*/
	CertAuthorityName string

	/* OcspTLSTrustedCommonName.

	   The expected Trusted Common Name of the OCSP responder remote certificate.
	*/
	OcspTLSTrustedCommonName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete client cert authority ocsp Tls trusted common name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams) WithDefaults() *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete client cert authority ocsp Tls trusted common name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete client cert authority ocsp Tls trusted common name params
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams) WithTimeout(timeout time.Duration) *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete client cert authority ocsp Tls trusted common name params
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete client cert authority ocsp Tls trusted common name params
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams) WithContext(ctx context.Context) *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete client cert authority ocsp Tls trusted common name params
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete client cert authority ocsp Tls trusted common name params
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams) WithHTTPClient(client *http.Client) *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete client cert authority ocsp Tls trusted common name params
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCertAuthorityName adds the certAuthorityName to the delete client cert authority ocsp Tls trusted common name params
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams) WithCertAuthorityName(certAuthorityName string) *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams {
	o.SetCertAuthorityName(certAuthorityName)
	return o
}

// SetCertAuthorityName adds the certAuthorityName to the delete client cert authority ocsp Tls trusted common name params
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams) SetCertAuthorityName(certAuthorityName string) {
	o.CertAuthorityName = certAuthorityName
}

// WithOcspTLSTrustedCommonName adds the ocspTLSTrustedCommonName to the delete client cert authority ocsp Tls trusted common name params
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams) WithOcspTLSTrustedCommonName(ocspTLSTrustedCommonName string) *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams {
	o.SetOcspTLSTrustedCommonName(ocspTLSTrustedCommonName)
	return o
}

// SetOcspTLSTrustedCommonName adds the ocspTlsTrustedCommonName to the delete client cert authority ocsp Tls trusted common name params
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams) SetOcspTLSTrustedCommonName(ocspTLSTrustedCommonName string) {
	o.OcspTLSTrustedCommonName = ocspTLSTrustedCommonName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClientCertAuthorityOcspTLSTrustedCommonNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param certAuthorityName
	if err := r.SetPathParam("certAuthorityName", o.CertAuthorityName); err != nil {
		return err
	}

	// path param ocspTlsTrustedCommonName
	if err := r.SetPathParam("ocspTlsTrustedCommonName", o.OcspTLSTrustedCommonName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}