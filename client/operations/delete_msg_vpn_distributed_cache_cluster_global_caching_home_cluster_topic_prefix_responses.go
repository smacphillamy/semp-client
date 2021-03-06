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

// DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixReader is a Reader for the DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix structure.
type DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK creates a DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK with default headers values
func NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK() *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK {
	return &DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK{}
}

/*DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK handles this case with default header values.

The request metadata.
*/
type DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix}][%d] deleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK  %+v", 200, o.Payload)
}

func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault creates a DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault with default headers values
func NewDeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault(code int) *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault {
	return &DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault{
		_statusCode: code,
	}
}

/*DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault handles this case with default header values.

The error response.
*/
type DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn distributed cache cluster global caching home cluster topic prefix default response
func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters/{homeClusterName}/topicPrefixes/{topicPrefix}][%d] deleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefix default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
