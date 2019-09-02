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

// GetMsgVpnRestDeliveryPointReader is a Reader for the GetMsgVpnRestDeliveryPoint structure.
type GetMsgVpnRestDeliveryPointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnRestDeliveryPointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMsgVpnRestDeliveryPointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMsgVpnRestDeliveryPointDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnRestDeliveryPointOK creates a GetMsgVpnRestDeliveryPointOK with default headers values
func NewGetMsgVpnRestDeliveryPointOK() *GetMsgVpnRestDeliveryPointOK {
	return &GetMsgVpnRestDeliveryPointOK{}
}

/*GetMsgVpnRestDeliveryPointOK handles this case with default header values.

The REST Delivery Point object's attributes, and the request metadata.
*/
type GetMsgVpnRestDeliveryPointOK struct {
	Payload *models.MsgVpnRestDeliveryPointResponse
}

func (o *GetMsgVpnRestDeliveryPointOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}][%d] getMsgVpnRestDeliveryPointOK  %+v", 200, o.Payload)
}

func (o *GetMsgVpnRestDeliveryPointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnRestDeliveryPointResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnRestDeliveryPointDefault creates a GetMsgVpnRestDeliveryPointDefault with default headers values
func NewGetMsgVpnRestDeliveryPointDefault(code int) *GetMsgVpnRestDeliveryPointDefault {
	return &GetMsgVpnRestDeliveryPointDefault{
		_statusCode: code,
	}
}

/*GetMsgVpnRestDeliveryPointDefault handles this case with default header values.

Error response
*/
type GetMsgVpnRestDeliveryPointDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn rest delivery point default response
func (o *GetMsgVpnRestDeliveryPointDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnRestDeliveryPointDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}][%d] getMsgVpnRestDeliveryPoint default  %+v", o._statusCode, o.Payload)
}

func (o *GetMsgVpnRestDeliveryPointDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}