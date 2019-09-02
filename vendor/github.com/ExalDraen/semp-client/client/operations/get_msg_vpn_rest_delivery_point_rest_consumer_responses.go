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

// GetMsgVpnRestDeliveryPointRestConsumerReader is a Reader for the GetMsgVpnRestDeliveryPointRestConsumer structure.
type GetMsgVpnRestDeliveryPointRestConsumerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnRestDeliveryPointRestConsumerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMsgVpnRestDeliveryPointRestConsumerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMsgVpnRestDeliveryPointRestConsumerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnRestDeliveryPointRestConsumerOK creates a GetMsgVpnRestDeliveryPointRestConsumerOK with default headers values
func NewGetMsgVpnRestDeliveryPointRestConsumerOK() *GetMsgVpnRestDeliveryPointRestConsumerOK {
	return &GetMsgVpnRestDeliveryPointRestConsumerOK{}
}

/*GetMsgVpnRestDeliveryPointRestConsumerOK handles this case with default header values.

The REST Consumer object's attributes, and the request metadata.
*/
type GetMsgVpnRestDeliveryPointRestConsumerOK struct {
	Payload *models.MsgVpnRestDeliveryPointRestConsumerResponse
}

func (o *GetMsgVpnRestDeliveryPointRestConsumerOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}][%d] getMsgVpnRestDeliveryPointRestConsumerOK  %+v", 200, o.Payload)
}

func (o *GetMsgVpnRestDeliveryPointRestConsumerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnRestDeliveryPointRestConsumerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnRestDeliveryPointRestConsumerDefault creates a GetMsgVpnRestDeliveryPointRestConsumerDefault with default headers values
func NewGetMsgVpnRestDeliveryPointRestConsumerDefault(code int) *GetMsgVpnRestDeliveryPointRestConsumerDefault {
	return &GetMsgVpnRestDeliveryPointRestConsumerDefault{
		_statusCode: code,
	}
}

/*GetMsgVpnRestDeliveryPointRestConsumerDefault handles this case with default header values.

Error response
*/
type GetMsgVpnRestDeliveryPointRestConsumerDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn rest delivery point rest consumer default response
func (o *GetMsgVpnRestDeliveryPointRestConsumerDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnRestDeliveryPointRestConsumerDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/restConsumers/{restConsumerName}][%d] getMsgVpnRestDeliveryPointRestConsumer default  %+v", o._statusCode, o.Payload)
}

func (o *GetMsgVpnRestDeliveryPointRestConsumerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
