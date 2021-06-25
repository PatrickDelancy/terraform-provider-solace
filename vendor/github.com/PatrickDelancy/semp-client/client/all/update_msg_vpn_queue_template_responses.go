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

// UpdateMsgVpnQueueTemplateReader is a Reader for the UpdateMsgVpnQueueTemplate structure.
type UpdateMsgVpnQueueTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMsgVpnQueueTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMsgVpnQueueTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateMsgVpnQueueTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMsgVpnQueueTemplateOK creates a UpdateMsgVpnQueueTemplateOK with default headers values
func NewUpdateMsgVpnQueueTemplateOK() *UpdateMsgVpnQueueTemplateOK {
	return &UpdateMsgVpnQueueTemplateOK{}
}

/* UpdateMsgVpnQueueTemplateOK describes a response with status code 200, with default header values.

The Queue Template object's attributes after being updated, and the request metadata.
*/
type UpdateMsgVpnQueueTemplateOK struct {
	Payload *models.MsgVpnQueueTemplateResponse
}

func (o *UpdateMsgVpnQueueTemplateOK) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName}][%d] updateMsgVpnQueueTemplateOK  %+v", 200, o.Payload)
}
func (o *UpdateMsgVpnQueueTemplateOK) GetPayload() *models.MsgVpnQueueTemplateResponse {
	return o.Payload
}

func (o *UpdateMsgVpnQueueTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnQueueTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMsgVpnQueueTemplateDefault creates a UpdateMsgVpnQueueTemplateDefault with default headers values
func NewUpdateMsgVpnQueueTemplateDefault(code int) *UpdateMsgVpnQueueTemplateDefault {
	return &UpdateMsgVpnQueueTemplateDefault{
		_statusCode: code,
	}
}

/* UpdateMsgVpnQueueTemplateDefault describes a response with status code -1, with default header values.

The error response.
*/
type UpdateMsgVpnQueueTemplateDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the update msg vpn queue template default response
func (o *UpdateMsgVpnQueueTemplateDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMsgVpnQueueTemplateDefault) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/queueTemplates/{queueTemplateName}][%d] updateMsgVpnQueueTemplate default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateMsgVpnQueueTemplateDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *UpdateMsgVpnQueueTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}