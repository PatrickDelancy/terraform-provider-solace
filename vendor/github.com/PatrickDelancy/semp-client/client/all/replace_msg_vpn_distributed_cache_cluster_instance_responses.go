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

// ReplaceMsgVpnDistributedCacheClusterInstanceReader is a Reader for the ReplaceMsgVpnDistributedCacheClusterInstance structure.
type ReplaceMsgVpnDistributedCacheClusterInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceMsgVpnDistributedCacheClusterInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceMsgVpnDistributedCacheClusterInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReplaceMsgVpnDistributedCacheClusterInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceMsgVpnDistributedCacheClusterInstanceOK creates a ReplaceMsgVpnDistributedCacheClusterInstanceOK with default headers values
func NewReplaceMsgVpnDistributedCacheClusterInstanceOK() *ReplaceMsgVpnDistributedCacheClusterInstanceOK {
	return &ReplaceMsgVpnDistributedCacheClusterInstanceOK{}
}

/* ReplaceMsgVpnDistributedCacheClusterInstanceOK describes a response with status code 200, with default header values.

The Cache Instance object's attributes after being replaced, and the request metadata.
*/
type ReplaceMsgVpnDistributedCacheClusterInstanceOK struct {
	Payload *models.MsgVpnDistributedCacheClusterInstanceResponse
}

func (o *ReplaceMsgVpnDistributedCacheClusterInstanceOK) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}][%d] replaceMsgVpnDistributedCacheClusterInstanceOK  %+v", 200, o.Payload)
}
func (o *ReplaceMsgVpnDistributedCacheClusterInstanceOK) GetPayload() *models.MsgVpnDistributedCacheClusterInstanceResponse {
	return o.Payload
}

func (o *ReplaceMsgVpnDistributedCacheClusterInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnDistributedCacheClusterInstanceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceMsgVpnDistributedCacheClusterInstanceDefault creates a ReplaceMsgVpnDistributedCacheClusterInstanceDefault with default headers values
func NewReplaceMsgVpnDistributedCacheClusterInstanceDefault(code int) *ReplaceMsgVpnDistributedCacheClusterInstanceDefault {
	return &ReplaceMsgVpnDistributedCacheClusterInstanceDefault{
		_statusCode: code,
	}
}

/* ReplaceMsgVpnDistributedCacheClusterInstanceDefault describes a response with status code -1, with default header values.

The error response.
*/
type ReplaceMsgVpnDistributedCacheClusterInstanceDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the replace msg vpn distributed cache cluster instance default response
func (o *ReplaceMsgVpnDistributedCacheClusterInstanceDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceMsgVpnDistributedCacheClusterInstanceDefault) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/instances/{instanceName}][%d] replaceMsgVpnDistributedCacheClusterInstance default  %+v", o._statusCode, o.Payload)
}
func (o *ReplaceMsgVpnDistributedCacheClusterInstanceDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *ReplaceMsgVpnDistributedCacheClusterInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}