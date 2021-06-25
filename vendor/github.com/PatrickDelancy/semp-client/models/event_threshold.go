// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EventThreshold event threshold
//
// swagger:model EventThreshold
type EventThreshold struct {

	// The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.
	ClearPercent int64 `json:"clearPercent,omitempty"`

	// The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.
	ClearValue int64 `json:"clearValue,omitempty"`

	// The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.
	SetPercent int64 `json:"setPercent,omitempty"`

	// The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.
	SetValue int64 `json:"setValue,omitempty"`
}

// Validate validates this event threshold
func (m *EventThreshold) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this event threshold based on context it is used
func (m *EventThreshold) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventThreshold) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventThreshold) UnmarshalBinary(b []byte) error {
	var res EventThreshold
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}