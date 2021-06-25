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

// DeleteMsgVpnDistributedCacheClusterReader is a Reader for the DeleteMsgVpnDistributedCacheCluster structure.
type DeleteMsgVpnDistributedCacheClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnDistributedCacheClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMsgVpnDistributedCacheClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteMsgVpnDistributedCacheClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnDistributedCacheClusterOK creates a DeleteMsgVpnDistributedCacheClusterOK with default headers values
func NewDeleteMsgVpnDistributedCacheClusterOK() *DeleteMsgVpnDistributedCacheClusterOK {
	return &DeleteMsgVpnDistributedCacheClusterOK{}
}

/* DeleteMsgVpnDistributedCacheClusterOK describes a response with status code 200, with default header values.

The request metadata.
*/
type DeleteMsgVpnDistributedCacheClusterOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnDistributedCacheClusterOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}][%d] deleteMsgVpnDistributedCacheClusterOK  %+v", 200, o.Payload)
}
func (o *DeleteMsgVpnDistributedCacheClusterOK) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteMsgVpnDistributedCacheClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnDistributedCacheClusterDefault creates a DeleteMsgVpnDistributedCacheClusterDefault with default headers values
func NewDeleteMsgVpnDistributedCacheClusterDefault(code int) *DeleteMsgVpnDistributedCacheClusterDefault {
	return &DeleteMsgVpnDistributedCacheClusterDefault{
		_statusCode: code,
	}
}

/* DeleteMsgVpnDistributedCacheClusterDefault describes a response with status code -1, with default header values.

The error response.
*/
type DeleteMsgVpnDistributedCacheClusterDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn distributed cache cluster default response
func (o *DeleteMsgVpnDistributedCacheClusterDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnDistributedCacheClusterDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}][%d] deleteMsgVpnDistributedCacheCluster default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteMsgVpnDistributedCacheClusterDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *DeleteMsgVpnDistributedCacheClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
