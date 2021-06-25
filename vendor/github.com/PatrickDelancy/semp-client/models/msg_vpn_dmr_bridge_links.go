// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnDmrBridgeLinks msg vpn dmr bridge links
//
// swagger:model MsgVpnDmrBridgeLinks
type MsgVpnDmrBridgeLinks struct {

	// The URI of this DMR Bridge object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn dmr bridge links
func (m *MsgVpnDmrBridgeLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn dmr bridge links based on context it is used
func (m *MsgVpnDmrBridgeLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnDmrBridgeLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnDmrBridgeLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnDmrBridgeLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
