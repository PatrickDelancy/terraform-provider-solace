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

// GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixReader is a Reader for the GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix structure.
type GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK creates a GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK with default headers values
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK() *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK {
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK{}
}

/* GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK describes a response with status code 200, with default header values.

The Topic Prefix object's attributes, and the request metadata.
*/
type GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK struct {
	Payload *models.MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse
}

func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix}][%d] getMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK  %+v", 200, o.Payload)
}
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK) GetPayload() *models.MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse {
	return o.Payload
}

func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault creates a GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault with default headers values
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault(code int) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault {
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault{
		_statusCode: code,
	}
}

/* GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault describes a response with status code -1, with default header values.

The error response.
*/
type GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn distributed cache cluster global caching home cluster topic prefix default response
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix}][%d] getMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix default  %+v", o._statusCode, o.Payload)
}
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault) GetPayload() *models.SempMetaOnlyResponse {
	return o.Payload
}

func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
