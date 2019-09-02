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

// CreateMsgVpnACLProfileSubscribeExceptionReader is a Reader for the CreateMsgVpnACLProfileSubscribeException structure.
type CreateMsgVpnACLProfileSubscribeExceptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnACLProfileSubscribeExceptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateMsgVpnACLProfileSubscribeExceptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateMsgVpnACLProfileSubscribeExceptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnACLProfileSubscribeExceptionOK creates a CreateMsgVpnACLProfileSubscribeExceptionOK with default headers values
func NewCreateMsgVpnACLProfileSubscribeExceptionOK() *CreateMsgVpnACLProfileSubscribeExceptionOK {
	return &CreateMsgVpnACLProfileSubscribeExceptionOK{}
}

/*CreateMsgVpnACLProfileSubscribeExceptionOK handles this case with default header values.

The Subscribe Topic Exception object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnACLProfileSubscribeExceptionOK struct {
	Payload *models.MsgVpnACLProfileSubscribeExceptionResponse
}

func (o *CreateMsgVpnACLProfileSubscribeExceptionOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions][%d] createMsgVpnAclProfileSubscribeExceptionOK  %+v", 200, o.Payload)
}

func (o *CreateMsgVpnACLProfileSubscribeExceptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnACLProfileSubscribeExceptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnACLProfileSubscribeExceptionDefault creates a CreateMsgVpnACLProfileSubscribeExceptionDefault with default headers values
func NewCreateMsgVpnACLProfileSubscribeExceptionDefault(code int) *CreateMsgVpnACLProfileSubscribeExceptionDefault {
	return &CreateMsgVpnACLProfileSubscribeExceptionDefault{
		_statusCode: code,
	}
}

/*CreateMsgVpnACLProfileSubscribeExceptionDefault handles this case with default header values.

Error response
*/
type CreateMsgVpnACLProfileSubscribeExceptionDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn Acl profile subscribe exception default response
func (o *CreateMsgVpnACLProfileSubscribeExceptionDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnACLProfileSubscribeExceptionDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/subscribeExceptions][%d] createMsgVpnAclProfileSubscribeException default  %+v", o._statusCode, o.Payload)
}

func (o *CreateMsgVpnACLProfileSubscribeExceptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
