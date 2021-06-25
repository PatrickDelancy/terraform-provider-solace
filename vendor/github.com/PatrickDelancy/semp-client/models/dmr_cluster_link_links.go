// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DmrClusterLinkLinks dmr cluster link links
//
// swagger:model DmrClusterLinkLinks
type DmrClusterLinkLinks struct {

	// The URI of this Link's collection of Remote Address objects.
	RemoteAddressesURI string `json:"remoteAddressesUri,omitempty"`

	// The URI of this Link's collection of Trusted Common Name objects. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation.
	TLSTrustedCommonNamesURI string `json:"tlsTrustedCommonNamesUri,omitempty"`

	// The URI of this Link object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this dmr cluster link links
func (m *DmrClusterLinkLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dmr cluster link links based on context it is used
func (m *DmrClusterLinkLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DmrClusterLinkLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DmrClusterLinkLinks) UnmarshalBinary(b []byte) error {
	var res DmrClusterLinkLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}