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

// GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesReader is a Reader for the GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes structure.
type GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOK creates a GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOK with default headers values
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOK() *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOK {
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOK{}
}

/*GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOK handles this case with default header values.

The list of Topic Prefix objects' attributes, and the request metadata.
*/
type GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOK struct {
	Payload *models.MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse
}

func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes][%d] getMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOK  %+v", 200, o.Payload)
}

func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesDefault creates a GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesDefault with default headers values
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesDefault(code int) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesDefault {
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesDefault{
		_statusCode: code,
	}
}

/*GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesDefault handles this case with default header values.

The error response.
*/
type GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn distributed cache cluster global caching home cluster topic prefixes default response
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes][%d] getMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixes default  %+v", o._statusCode, o.Payload)
}

func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}