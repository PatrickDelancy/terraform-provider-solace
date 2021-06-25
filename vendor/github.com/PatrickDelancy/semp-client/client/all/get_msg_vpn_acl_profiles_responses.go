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

// GetMsgVpnACLProfilesReader is a Reader for the GetMsgVpnACLProfiles structure.
type GetMsgVpnACLProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnACLProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMsgVpnACLProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMsgVpnACLProfilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnACLProfilesOK creates a GetMsgVpnACLProfilesOK with default headers values
func NewGetMsgVpnACLProfilesOK() *GetMsgVpnACLProfilesOK {
	return &GetMsgVpnACLProfilesOK{}
}

/* GetMsgVpnACLProfilesOK describes a response with status code 200, with default header values.

The list of ACL Profile objects' attributes, and the request metadata.
*/
type GetMsgVpnACLProfilesOK struct {
	Payload *models.MsgVpnACLProfilesResponse
}

func (o *GetMsgVpnACLProfilesOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/aclProfiles][%d] getMsgVpnAclProfilesOK  %+v", 200, o.Payload)
}
func (o *GetMsgVpnACLProfilesOK) GetPayload() *models.MsgVpnACLProfilesResponse {
	return o.Payload
}

func (o *GetMsgVpnACLProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnACLProfilesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnACLProfilesDefault creates a GetMsgVpnACLProfilesDefault with default headers values
func NewGetMsgVpnACLProfilesDefault(code int) *GetMsgVpnACLProfilesDefault {
	return &GetMsgVpnACLProfilesDefault{
		_statusCode: code,
	}
}

/* GetMsgVpnACLProfilesDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetMsgVpnACLProfilesDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn Acl profiles default response
func (o *GetMsgVpnACLProfilesDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnACLProfilesDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/aclProfiles][%d] getMsgVpnAclProfiles default  %+v", o._statusCode, o.Payload)
}
func (o *GetMsgVpnACLProfilesDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetMsgVpnACLProfilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}