// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ExalDraen/semp-client/models"
)

// DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameReader is a Reader for the DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonName structure.
type DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameOK creates a DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameOK with default headers values
func NewDeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameOK() *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameOK {
	return &DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameOK{}
}

/*DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameOK handles this case with default header values.

The request metadata.
*/
type DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames/{tlsTrustedCommonName}][%d] deleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonNameOK  %+v", 200, o.Payload)
}

func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameDefault creates a DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameDefault with default headers values
func NewDeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameDefault(code int) *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameDefault {
	return &DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameDefault{
		_statusCode: code,
	}
}

/*DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameDefault handles this case with default header values.

Error response
*/
type DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn rest delivery point rest consumer Tls trusted common name default response
func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}/tlsTrustedCommonNames/{tlsTrustedCommonName}][%d] deleteMsgVpnRestDeliveryPointRestConsumerTlsTrustedCommonName default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMsgVpnRestDeliveryPointRestConsumerTLSTrustedCommonNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
