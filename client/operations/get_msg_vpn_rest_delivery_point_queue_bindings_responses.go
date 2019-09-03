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

// GetMsgVpnRestDeliveryPointQueueBindingsReader is a Reader for the GetMsgVpnRestDeliveryPointQueueBindings structure.
type GetMsgVpnRestDeliveryPointQueueBindingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnRestDeliveryPointQueueBindingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMsgVpnRestDeliveryPointQueueBindingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMsgVpnRestDeliveryPointQueueBindingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnRestDeliveryPointQueueBindingsOK creates a GetMsgVpnRestDeliveryPointQueueBindingsOK with default headers values
func NewGetMsgVpnRestDeliveryPointQueueBindingsOK() *GetMsgVpnRestDeliveryPointQueueBindingsOK {
	return &GetMsgVpnRestDeliveryPointQueueBindingsOK{}
}

/*GetMsgVpnRestDeliveryPointQueueBindingsOK handles this case with default header values.

The list of Queue Binding objects' attributes, and the request metadata.
*/
type GetMsgVpnRestDeliveryPointQueueBindingsOK struct {
	Payload *models.MsgVpnRestDeliveryPointQueueBindingsResponse
}

func (o *GetMsgVpnRestDeliveryPointQueueBindingsOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings][%d] getMsgVpnRestDeliveryPointQueueBindingsOK  %+v", 200, o.Payload)
}

func (o *GetMsgVpnRestDeliveryPointQueueBindingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnRestDeliveryPointQueueBindingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnRestDeliveryPointQueueBindingsDefault creates a GetMsgVpnRestDeliveryPointQueueBindingsDefault with default headers values
func NewGetMsgVpnRestDeliveryPointQueueBindingsDefault(code int) *GetMsgVpnRestDeliveryPointQueueBindingsDefault {
	return &GetMsgVpnRestDeliveryPointQueueBindingsDefault{
		_statusCode: code,
	}
}

/*GetMsgVpnRestDeliveryPointQueueBindingsDefault handles this case with default header values.

The error response.
*/
type GetMsgVpnRestDeliveryPointQueueBindingsDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn rest delivery point queue bindings default response
func (o *GetMsgVpnRestDeliveryPointQueueBindingsDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnRestDeliveryPointQueueBindingsDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}/queueBindings][%d] getMsgVpnRestDeliveryPointQueueBindings default  %+v", o._statusCode, o.Payload)
}

func (o *GetMsgVpnRestDeliveryPointQueueBindingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
