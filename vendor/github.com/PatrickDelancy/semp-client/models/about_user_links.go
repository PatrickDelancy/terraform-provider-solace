// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AboutUserLinks about user links
//
// swagger:model AboutUserLinks
type AboutUserLinks struct {

	// The URI of this User's collection of User Message VPN objects.
	MsgVpnsURI string `json:"msgVpnsUri,omitempty"`

	// The URI of this User object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this about user links
func (m *AboutUserLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this about user links based on context it is used
func (m *AboutUserLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AboutUserLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AboutUserLinks) UnmarshalBinary(b []byte) error {
	var res AboutUserLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}