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

// MsgVpnACLProfilePublishException msg vpn Acl profile publish exception
//
// swagger:model MsgVpnAclProfilePublishException
type MsgVpnACLProfilePublishException struct {

	// The name of the ACL Profile. Deprecated since 2.14. Replaced by publishTopicExceptions.
	ACLProfileName string `json:"aclProfileName,omitempty"`

	// The name of the Message VPN. Deprecated since 2.14. Replaced by publishTopicExceptions.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The topic for the exception to the default action taken. May include wildcard characters. Deprecated since 2.14. Replaced by publishTopicExceptions.
	PublishExceptionTopic string `json:"publishExceptionTopic,omitempty"`

	// The syntax of the topic for the exception to the default action taken. The allowed values and their meaning are:
	//
	// <pre>
	// "smf" - Topic uses SMF syntax.
	// "mqtt" - Topic uses MQTT syntax.
	// </pre>
	//  Deprecated since 2.14. Replaced by publishTopicExceptions.
	// Enum: [smf mqtt]
	TopicSyntax string `json:"topicSyntax,omitempty"`
}

// Validate validates this msg vpn Acl profile publish exception
func (m *MsgVpnACLProfilePublishException) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTopicSyntax(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var msgVpnAclProfilePublishExceptionTypeTopicSyntaxPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["smf","mqtt"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnAclProfilePublishExceptionTypeTopicSyntaxPropEnum = append(msgVpnAclProfilePublishExceptionTypeTopicSyntaxPropEnum, v)
	}
}

const (

	// MsgVpnACLProfilePublishExceptionTopicSyntaxSmf captures enum value "smf"
	MsgVpnACLProfilePublishExceptionTopicSyntaxSmf string = "smf"

	// MsgVpnACLProfilePublishExceptionTopicSyntaxMqtt captures enum value "mqtt"
	MsgVpnACLProfilePublishExceptionTopicSyntaxMqtt string = "mqtt"
)

// prop value enum
func (m *MsgVpnACLProfilePublishException) validateTopicSyntaxEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, msgVpnAclProfilePublishExceptionTypeTopicSyntaxPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnACLProfilePublishException) validateTopicSyntax(formats strfmt.Registry) error {
	if swag.IsZero(m.TopicSyntax) { // not required
		return nil
	}

	// value enum
	if err := m.validateTopicSyntaxEnum("topicSyntax", "body", m.TopicSyntax); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this msg vpn Acl profile publish exception based on context it is used
func (m *MsgVpnACLProfilePublishException) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnACLProfilePublishException) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnACLProfilePublishException) UnmarshalBinary(b []byte) error {
	var res MsgVpnACLProfilePublishException
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}