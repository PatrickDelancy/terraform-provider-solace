// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnACLProfileSubscribeShareNameExceptionLinks msg vpn Acl profile subscribe share name exception links
//
// swagger:model MsgVpnAclProfileSubscribeShareNameExceptionLinks
type MsgVpnACLProfileSubscribeShareNameExceptionLinks struct {

	// The URI of this Subscribe Share Name Exception object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn Acl profile subscribe share name exception links
func (m *MsgVpnACLProfileSubscribeShareNameExceptionLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn Acl profile subscribe share name exception links based on context it is used
func (m *MsgVpnACLProfileSubscribeShareNameExceptionLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnACLProfileSubscribeShareNameExceptionLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnACLProfileSubscribeShareNameExceptionLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnACLProfileSubscribeShareNameExceptionLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
