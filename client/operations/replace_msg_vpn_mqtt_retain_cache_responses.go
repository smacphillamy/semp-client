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

// ReplaceMsgVpnMqttRetainCacheReader is a Reader for the ReplaceMsgVpnMqttRetainCache structure.
type ReplaceMsgVpnMqttRetainCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceMsgVpnMqttRetainCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceMsgVpnMqttRetainCacheOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewReplaceMsgVpnMqttRetainCacheDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceMsgVpnMqttRetainCacheOK creates a ReplaceMsgVpnMqttRetainCacheOK with default headers values
func NewReplaceMsgVpnMqttRetainCacheOK() *ReplaceMsgVpnMqttRetainCacheOK {
	return &ReplaceMsgVpnMqttRetainCacheOK{}
}

/*ReplaceMsgVpnMqttRetainCacheOK handles this case with default header values.

The MQTT Retain Cache object's attributes after being replaced, and the request metadata.
*/
type ReplaceMsgVpnMqttRetainCacheOK struct {
	Payload *models.MsgVpnMqttRetainCacheResponse
}

func (o *ReplaceMsgVpnMqttRetainCacheOK) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName}][%d] replaceMsgVpnMqttRetainCacheOK  %+v", 200, o.Payload)
}

func (o *ReplaceMsgVpnMqttRetainCacheOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnMqttRetainCacheResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceMsgVpnMqttRetainCacheDefault creates a ReplaceMsgVpnMqttRetainCacheDefault with default headers values
func NewReplaceMsgVpnMqttRetainCacheDefault(code int) *ReplaceMsgVpnMqttRetainCacheDefault {
	return &ReplaceMsgVpnMqttRetainCacheDefault{
		_statusCode: code,
	}
}

/*ReplaceMsgVpnMqttRetainCacheDefault handles this case with default header values.

The error response.
*/
type ReplaceMsgVpnMqttRetainCacheDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the replace msg vpn mqtt retain cache default response
func (o *ReplaceMsgVpnMqttRetainCacheDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceMsgVpnMqttRetainCacheDefault) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}/mqttRetainCaches/{cacheName}][%d] replaceMsgVpnMqttRetainCache default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceMsgVpnMqttRetainCacheDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
