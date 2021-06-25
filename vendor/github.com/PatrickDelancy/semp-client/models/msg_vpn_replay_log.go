// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnReplayLog msg vpn replay log
//
// swagger:model MsgVpnReplayLog
type MsgVpnReplayLog struct {

	// Enable or disable the transmission of messages from the Replay Log. The default value is `false`.
	EgressEnabled bool `json:"egressEnabled,omitempty"`

	// Enable or disable the reception of messages to the Replay Log. The default value is `false`.
	IngressEnabled bool `json:"ingressEnabled,omitempty"`

	// The maximum spool usage allowed by the Replay Log, in megabytes (MB). If this limit is exceeded, old messages will be trimmed. The default value is `0`.
	MaxSpoolUsage int64 `json:"maxSpoolUsage,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The name of the Replay Log.
	ReplayLogName string `json:"replayLogName,omitempty"`
}

// Validate validates this msg vpn replay log
func (m *MsgVpnReplayLog) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn replay log based on context it is used
func (m *MsgVpnReplayLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnReplayLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnReplayLog) UnmarshalBinary(b []byte) error {
	var res MsgVpnReplayLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}