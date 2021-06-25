// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DmrClusterLinks dmr cluster links
//
// swagger:model DmrClusterLinks
type DmrClusterLinks struct {

	// The URI of this Cluster's collection of Link objects.
	LinksURI string `json:"linksUri,omitempty"`

	// The URI of this Cluster object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this dmr cluster links
func (m *DmrClusterLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dmr cluster links based on context it is used
func (m *DmrClusterLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DmrClusterLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DmrClusterLinks) UnmarshalBinary(b []byte) error {
	var res DmrClusterLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
