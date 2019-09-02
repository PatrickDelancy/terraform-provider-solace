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

// DeleteMsgVpnACLProfilePublishExceptionReader is a Reader for the DeleteMsgVpnACLProfilePublishException structure.
type DeleteMsgVpnACLProfilePublishExceptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnACLProfilePublishExceptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMsgVpnACLProfilePublishExceptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMsgVpnACLProfilePublishExceptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnACLProfilePublishExceptionOK creates a DeleteMsgVpnACLProfilePublishExceptionOK with default headers values
func NewDeleteMsgVpnACLProfilePublishExceptionOK() *DeleteMsgVpnACLProfilePublishExceptionOK {
	return &DeleteMsgVpnACLProfilePublishExceptionOK{}
}

/*DeleteMsgVpnACLProfilePublishExceptionOK handles this case with default header values.

The request metadata.
*/
type DeleteMsgVpnACLProfilePublishExceptionOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnACLProfilePublishExceptionOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic}][%d] deleteMsgVpnAclProfilePublishExceptionOK  %+v", 200, o.Payload)
}

func (o *DeleteMsgVpnACLProfilePublishExceptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnACLProfilePublishExceptionDefault creates a DeleteMsgVpnACLProfilePublishExceptionDefault with default headers values
func NewDeleteMsgVpnACLProfilePublishExceptionDefault(code int) *DeleteMsgVpnACLProfilePublishExceptionDefault {
	return &DeleteMsgVpnACLProfilePublishExceptionDefault{
		_statusCode: code,
	}
}

/*DeleteMsgVpnACLProfilePublishExceptionDefault handles this case with default header values.

Error response
*/
type DeleteMsgVpnACLProfilePublishExceptionDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn Acl profile publish exception default response
func (o *DeleteMsgVpnACLProfilePublishExceptionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnACLProfilePublishExceptionDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions/{topicSyntax},{publishExceptionTopic}][%d] deleteMsgVpnAclProfilePublishException default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMsgVpnACLProfilePublishExceptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
