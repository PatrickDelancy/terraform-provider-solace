// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnACLProfilePublishTopicExceptionLinks msg vpn Acl profile publish topic exception links
//
// swagger:model MsgVpnAclProfilePublishTopicExceptionLinks
type MsgVpnACLProfilePublishTopicExceptionLinks struct {

	// The URI of this Publish Topic Exception object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn Acl profile publish topic exception links
func (m *MsgVpnACLProfilePublishTopicExceptionLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn Acl profile publish topic exception links based on context it is used
func (m *MsgVpnACLProfilePublishTopicExceptionLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnACLProfilePublishTopicExceptionLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnACLProfilePublishTopicExceptionLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnACLProfilePublishTopicExceptionLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
