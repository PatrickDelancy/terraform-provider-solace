// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClientCertAuthority client cert authority
//
// swagger:model ClientCertAuthority
type ClientCertAuthority struct {

	// The name of the Certificate Authority.
	CertAuthorityName string `json:"certAuthorityName,omitempty"`

	// The PEM formatted content for the trusted root certificate of a client Certificate Authority. The default value is `""`.
	CertContent string `json:"certContent,omitempty"`

	// The scheduled CRL refresh day(s), specified as "daily" or a comma-separated list of days. Days must be specified as "Sun", "Mon", "Tue", "Wed", "Thu", "Fri", or "Sat", with no spaces, and in sorted order from Sunday to Saturday. The default value is `"daily"`.
	CrlDayList string `json:"crlDayList,omitempty"`

	// The scheduled CRL refresh time(s), specified as "hourly" or a comma-separated list of 24-hour times in the form hh:mm, or h:mm. There must be no spaces, and times must be in sorted order from 0:00 to 23:59. The default value is `"3:00"`.
	CrlTimeList string `json:"crlTimeList,omitempty"`

	// The URL for the CRL source. This is a required attribute for CRL to be operational and the URL must be complete with http:// included. The default value is `""`.
	CrlURL string `json:"crlUrl,omitempty"`

	// Enable or disable allowing a non-responder certificate to sign an OCSP response. Typically used with an OCSP override URL in cases where a single certificate is used to sign client certificates and OCSP responses. The default value is `false`.
	OcspNonResponderCertEnabled bool `json:"ocspNonResponderCertEnabled,omitempty"`

	// The OCSP responder URL to use for overriding the one supplied in the client certificate. The URL must be complete with http:// included. The default value is `""`.
	OcspOverrideURL string `json:"ocspOverrideUrl,omitempty"`

	// The timeout in seconds to receive a response from the OCSP responder after sending a request or making the initial connection attempt. The default value is `5`.
	OcspTimeout int64 `json:"ocspTimeout,omitempty"`

	// Enable or disable Certificate Authority revocation checking. The default value is `false`.
	RevocationCheckEnabled bool `json:"revocationCheckEnabled,omitempty"`
}

// Validate validates this client cert authority
func (m *ClientCertAuthority) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this client cert authority based on context it is used
func (m *ClientCertAuthority) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientCertAuthority) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientCertAuthority) UnmarshalBinary(b []byte) error {
	var res ClientCertAuthority
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}