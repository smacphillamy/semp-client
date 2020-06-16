// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/smacphillamy/semp-client/models"
)

// CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterReader is a Reader for the CreateMsgVpnDistributedCacheClusterGlobalCachingHomeCluster structure.
type CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK creates a CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK with default headers values
func NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK() *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK {
	return &CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK{}
}

/*CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK handles this case with default header values.

The Home Cache Cluster object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK struct {
	Payload *models.MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse
}

func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters][%d] createMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK  %+v", 200, o.Payload)
}

func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnDistributedCacheClusterGlobalCachingHomeClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault creates a CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault with default headers values
func NewCreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault(code int) *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault {
	return &CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault{
		_statusCode: code,
	}
}

/*CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault handles this case with default header values.

The error response.
*/
type CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn distributed cache cluster global caching home cluster default response
func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/distributedCaches/{cacheName}/clusters/{clusterName}/globalCachingHomeClusters][%d] createMsgVpnDistributedCacheClusterGlobalCachingHomeCluster default  %+v", o._statusCode, o.Payload)
}

func (o *CreateMsgVpnDistributedCacheClusterGlobalCachingHomeClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
