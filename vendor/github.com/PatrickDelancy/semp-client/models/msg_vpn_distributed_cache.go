// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MsgVpnDistributedCache msg vpn distributed cache
//
// swagger:model MsgVpnDistributedCache
type MsgVpnDistributedCache struct {

	// The name of the Distributed Cache.
	CacheName string `json:"cacheName,omitempty"`

	// Enable or disable the Distributed Cache. The default value is `false`.
	Enabled bool `json:"enabled,omitempty"`

	// The heartbeat interval, in seconds, used by the Cache Instances to monitor connectivity with the message broker. The default value is `10`.
	Heartbeat int64 `json:"heartbeat,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The scheduled delete message day(s), specified as "daily" or a comma-separated list of days. Days must be specified as "Sun", "Mon", "Tue", "Wed", "Thu", "Fri", or "Sat", with no spaces, and in sorted order from Sunday to Saturday. The default value is `""`.
	ScheduledDeleteMsgDayList string `json:"scheduledDeleteMsgDayList,omitempty"`

	// The scheduled delete message time(s), specified as "hourly" or a comma-separated list of 24-hour times in the form hh:mm, or h:mm. There must be no spaces, and times must be in sorted order from 0:00 to 23:59. The default value is `""`.
	ScheduledDeleteMsgTimeList string `json:"scheduledDeleteMsgTimeList,omitempty"`
}

// Validate validates this msg vpn distributed cache
func (m *MsgVpnDistributedCache) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this msg vpn distributed cache based on context it is used
func (m *MsgVpnDistributedCache) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnDistributedCache) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnDistributedCache) UnmarshalBinary(b []byte) error {
	var res MsgVpnDistributedCache
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}