// Code generated by go-swagger; DO NOT EDIT.

package all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PatrickDelancy/semp-client/models"
)

// GetMsgVpnMqttSessionSubscriptionReader is a Reader for the GetMsgVpnMqttSessionSubscription structure.
type GetMsgVpnMqttSessionSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnMqttSessionSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMsgVpnMqttSessionSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMsgVpnMqttSessionSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnMqttSessionSubscriptionOK creates a GetMsgVpnMqttSessionSubscriptionOK with default headers values
func NewGetMsgVpnMqttSessionSubscriptionOK() *GetMsgVpnMqttSessionSubscriptionOK {
	return &GetMsgVpnMqttSessionSubscriptionOK{}
}

/* GetMsgVpnMqttSessionSubscriptionOK describes a response with status code 200, with default header values.

The Subscription object's attributes, and the request metadata.
*/
type GetMsgVpnMqttSessionSubscriptionOK struct {
	Payload *models.MsgVpnMqttSessionSubscriptionResponse
}

func (o *GetMsgVpnMqttSessionSubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic}][%d] getMsgVpnMqttSessionSubscriptionOK  %+v", 200, o.Payload)
}
func (o *GetMsgVpnMqttSessionSubscriptionOK) GetPayload() *models.MsgVpnMqttSessionSubscriptionResponse {
	return o.Payload
}

func (o *GetMsgVpnMqttSessionSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnMqttSessionSubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnMqttSessionSubscriptionDefault creates a GetMsgVpnMqttSessionSubscriptionDefault with default headers values
func NewGetMsgVpnMqttSessionSubscriptionDefault(code int) *GetMsgVpnMqttSessionSubscriptionDefault {
	return &GetMsgVpnMqttSessionSubscriptionDefault{
		_statusCode: code,
	}
}

/* GetMsgVpnMqttSessionSubscriptionDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetMsgVpnMqttSessionSubscriptionDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn mqtt session subscription default response
func (o *GetMsgVpnMqttSessionSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnMqttSessionSubscriptionDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/mqttSessions/{mqttSessionClientId},{mqttSessionVirtualRouter}/subscriptions/{subscriptionTopic}][%d] getMsgVpnMqttSessionSubscription default  %+v", o._statusCode, o.Payload)
}
func (o *GetMsgVpnMqttSessionSubscriptionDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetMsgVpnMqttSessionSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}