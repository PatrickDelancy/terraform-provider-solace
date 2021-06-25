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

// GetMsgVpnDmrBridgesReader is a Reader for the GetMsgVpnDmrBridges structure.
type GetMsgVpnDmrBridgesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnDmrBridgesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMsgVpnDmrBridgesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMsgVpnDmrBridgesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnDmrBridgesOK creates a GetMsgVpnDmrBridgesOK with default headers values
func NewGetMsgVpnDmrBridgesOK() *GetMsgVpnDmrBridgesOK {
	return &GetMsgVpnDmrBridgesOK{}
}

/* GetMsgVpnDmrBridgesOK describes a response with status code 200, with default header values.

The list of DMR Bridge objects' attributes, and the request metadata.
*/
type GetMsgVpnDmrBridgesOK struct {
	Payload *models.MsgVpnDmrBridgesResponse
}

func (o *GetMsgVpnDmrBridgesOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/dmrBridges][%d] getMsgVpnDmrBridgesOK  %+v", 200, o.Payload)
}
func (o *GetMsgVpnDmrBridgesOK) GetPayload() *models.MsgVpnDmrBridgesResponse {
	return o.Payload
}

func (o *GetMsgVpnDmrBridgesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnDmrBridgesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnDmrBridgesDefault creates a GetMsgVpnDmrBridgesDefault with default headers values
func NewGetMsgVpnDmrBridgesDefault(code int) *GetMsgVpnDmrBridgesDefault {
	return &GetMsgVpnDmrBridgesDefault{
		_statusCode: code,
	}
}

/* GetMsgVpnDmrBridgesDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetMsgVpnDmrBridgesDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn dmr bridges default response
func (o *GetMsgVpnDmrBridgesDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnDmrBridgesDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/dmrBridges][%d] getMsgVpnDmrBridges default  %+v", o._statusCode, o.Payload)
}
func (o *GetMsgVpnDmrBridgesDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetMsgVpnDmrBridgesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
