// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DmrClusterLinkTLSTrustedCommonName dmr cluster link Tls trusted common name
//
// swagger:model DmrClusterLinkTlsTrustedCommonName
type DmrClusterLinkTLSTrustedCommonName struct {

	// The name of the Cluster. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.
	DmrClusterName string `json:"dmrClusterName,omitempty"`

	// The name of the node at the remote end of the Link. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.
	RemoteNodeName string `json:"remoteNodeName,omitempty"`

	// The expected trusted common name of the remote certificate. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.
	TLSTrustedCommonName string `json:"tlsTrustedCommonName,omitempty"`
}

// Validate validates this dmr cluster link Tls trusted common name
func (m *DmrClusterLinkTLSTrustedCommonName) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dmr cluster link Tls trusted common name based on context it is used
func (m *DmrClusterLinkTLSTrustedCommonName) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DmrClusterLinkTLSTrustedCommonName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DmrClusterLinkTLSTrustedCommonName) UnmarshalBinary(b []byte) error {
	var res DmrClusterLinkTLSTrustedCommonName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}