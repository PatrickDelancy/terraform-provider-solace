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

// CreateMsgVpnACLProfilePublishTopicExceptionReader is a Reader for the CreateMsgVpnACLProfilePublishTopicException structure.
type CreateMsgVpnACLProfilePublishTopicExceptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnACLProfilePublishTopicExceptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateMsgVpnACLProfilePublishTopicExceptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateMsgVpnACLProfilePublishTopicExceptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnACLProfilePublishTopicExceptionOK creates a CreateMsgVpnACLProfilePublishTopicExceptionOK with default headers values
func NewCreateMsgVpnACLProfilePublishTopicExceptionOK() *CreateMsgVpnACLProfilePublishTopicExceptionOK {
	return &CreateMsgVpnACLProfilePublishTopicExceptionOK{}
}

/* CreateMsgVpnACLProfilePublishTopicExceptionOK describes a response with status code 200, with default header values.

The Publish Topic Exception object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnACLProfilePublishTopicExceptionOK struct {
	Payload *models.MsgVpnACLProfilePublishTopicExceptionResponse
}

func (o *CreateMsgVpnACLProfilePublishTopicExceptionOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions][%d] createMsgVpnAclProfilePublishTopicExceptionOK  %+v", 200, o.Payload)
}
func (o *CreateMsgVpnACLProfilePublishTopicExceptionOK) GetPayload() *models.MsgVpnACLProfilePublishTopicExceptionResponse {
	return o.Payload
}

func (o *CreateMsgVpnACLProfilePublishTopicExceptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnACLProfilePublishTopicExceptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnACLProfilePublishTopicExceptionDefault creates a CreateMsgVpnACLProfilePublishTopicExceptionDefault with default headers values
func NewCreateMsgVpnACLProfilePublishTopicExceptionDefault(code int) *CreateMsgVpnACLProfilePublishTopicExceptionDefault {
	return &CreateMsgVpnACLProfilePublishTopicExceptionDefault{
		_statusCode: code,
	}
}

/* CreateMsgVpnACLProfilePublishTopicExceptionDefault describes a response with status code -1, with default header values.

The error response.
*/
type CreateMsgVpnACLProfilePublishTopicExceptionDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn Acl profile publish topic exception default response
func (o *CreateMsgVpnACLProfilePublishTopicExceptionDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnACLProfilePublishTopicExceptionDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishTopicExceptions][%d] createMsgVpnAclProfilePublishTopicException default  %+v", o._statusCode, o.Payload)
}
func (o *CreateMsgVpnACLProfilePublishTopicExceptionDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *CreateMsgVpnACLProfilePublishTopicExceptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}