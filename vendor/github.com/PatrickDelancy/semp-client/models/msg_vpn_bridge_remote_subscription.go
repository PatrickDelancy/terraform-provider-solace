// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MsgVpnBridgeRemoteSubscription msg vpn bridge remote subscription
//
// swagger:model MsgVpnBridgeRemoteSubscription
type MsgVpnBridgeRemoteSubscription struct {

	// The name of the Bridge.
	BridgeName string `json:"bridgeName,omitempty"`

	// The virtual router of the Bridge. The allowed values and their meaning are:
	//
	// <pre>
	// "primary" - The Bridge is used for the primary virtual router.
	// "backup" - The Bridge is used for the backup virtual router.
	// "auto" - The Bridge is automatically assigned a virtual router at creation, depending on the broker's active-standby role.
	// </pre>
	//
	// Enum: [primary backup auto]
	BridgeVirtualRouter string `json:"bridgeVirtualRouter,omitempty"`

	// Enable or disable deliver-always for the Bridge remote subscription topic instead of a deliver-to-one remote priority. A given topic for the Bridge may be deliver-to-one or deliver-always but not both.
	DeliverAlwaysEnabled bool `json:"deliverAlwaysEnabled,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The topic of the Bridge remote subscription.
	RemoteSubscriptionTopic string `json:"remoteSubscriptionTopic,omitempty"`
}

// Validate validates this msg vpn bridge remote subscription
func (m *MsgVpnBridgeRemoteSubscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBridgeVirtualRouter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var msgVpnBridgeRemoteSubscriptionTypeBridgeVirtualRouterPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["primary","backup","auto"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnBridgeRemoteSubscriptionTypeBridgeVirtualRouterPropEnum = append(msgVpnBridgeRemoteSubscriptionTypeBridgeVirtualRouterPropEnum, v)
	}
}

const (

	// MsgVpnBridgeRemoteSubscriptionBridgeVirtualRouterPrimary captures enum value "primary"
	MsgVpnBridgeRemoteSubscriptionBridgeVirtualRouterPrimary string = "primary"

	// MsgVpnBridgeRemoteSubscriptionBridgeVirtualRouterBackup captures enum value "backup"
	MsgVpnBridgeRemoteSubscriptionBridgeVirtualRouterBackup string = "backup"

	// MsgVpnBridgeRemoteSubscriptionBridgeVirtualRouterAuto captures enum value "auto"
	MsgVpnBridgeRemoteSubscriptionBridgeVirtualRouterAuto string = "auto"
)

// prop value enum
func (m *MsgVpnBridgeRemoteSubscription) validateBridgeVirtualRouterEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, msgVpnBridgeRemoteSubscriptionTypeBridgeVirtualRouterPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnBridgeRemoteSubscription) validateBridgeVirtualRouter(formats strfmt.Registry) error {
	if swag.IsZero(m.BridgeVirtualRouter) { // not required
		return nil
	}

	// value enum
	if err := m.validateBridgeVirtualRouterEnum("bridgeVirtualRouter", "body", m.BridgeVirtualRouter); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this msg vpn bridge remote subscription based on context it is used
func (m *MsgVpnBridgeRemoteSubscription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnBridgeRemoteSubscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnBridgeRemoteSubscription) UnmarshalBinary(b []byte) error {
	var res MsgVpnBridgeRemoteSubscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}