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

// DeleteMsgVpnDistributedCacheClusterTopicReader is a Reader for the DeleteMsgVpnDistributedCacheClusterTopic structure.
type DeleteMsgVpnDistributedCacheClusterTopicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnDistributedCacheClusterTopicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMsgVpnDistributedCacheClusterTopicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMsgVpnDistributedCacheClusterTopicDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnDistributedCacheClusterTopicOK creates a DeleteMsgVpnDistributedCacheClusterTopicOK with default headers values
func NewDeleteMsgVpnDistributedCacheClusterTopicOK() *DeleteMsgVpnDistributedCacheClusterTopicOK {
	return &DeleteMsgVpnDistributedCacheClusterTopicOK{}
}

/*DeleteMsgVpnDistributedCacheClusterTopicOK handles this case with default header values.

The request metadata.
*/
type DeleteMsgVpnDistributedCacheClusterTopicOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnDistributedCacheClusterTopicOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics/{topic}][%d] deleteMsgVpnDistributedCacheClusterTopicOK  %+v", 200, o.Payload)
}

func (o *DeleteMsgVpnDistributedCacheClusterTopicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnDistributedCacheClusterTopicDefault creates a DeleteMsgVpnDistributedCacheClusterTopicDefault with default headers values
func NewDeleteMsgVpnDistributedCacheClusterTopicDefault(code int) *DeleteMsgVpnDistributedCacheClusterTopicDefault {
	return &DeleteMsgVpnDistributedCacheClusterTopicDefault{
		_statusCode: code,
	}
}

/*DeleteMsgVpnDistributedCacheClusterTopicDefault handles this case with default header values.

The error response.
*/
type DeleteMsgVpnDistributedCacheClusterTopicDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn distributed cache cluster topic default response
func (o *DeleteMsgVpnDistributedCacheClusterTopicDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnDistributedCacheClusterTopicDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/topics/{topic}][%d] deleteMsgVpnDistributedCacheClusterTopic default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMsgVpnDistributedCacheClusterTopicDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}