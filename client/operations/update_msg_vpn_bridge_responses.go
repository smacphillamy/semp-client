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

// UpdateMsgVpnBridgeReader is a Reader for the UpdateMsgVpnBridge structure.
type UpdateMsgVpnBridgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMsgVpnBridgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateMsgVpnBridgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateMsgVpnBridgeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMsgVpnBridgeOK creates a UpdateMsgVpnBridgeOK with default headers values
func NewUpdateMsgVpnBridgeOK() *UpdateMsgVpnBridgeOK {
	return &UpdateMsgVpnBridgeOK{}
}

/*UpdateMsgVpnBridgeOK handles this case with default header values.

The Bridge object's attributes after being updated, and the request metadata.
*/
type UpdateMsgVpnBridgeOK struct {
	Payload *models.MsgVpnBridgeResponse
}

func (o *UpdateMsgVpnBridgeOK) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}][%d] updateMsgVpnBridgeOK  %+v", 200, o.Payload)
}

func (o *UpdateMsgVpnBridgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnBridgeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMsgVpnBridgeDefault creates a UpdateMsgVpnBridgeDefault with default headers values
func NewUpdateMsgVpnBridgeDefault(code int) *UpdateMsgVpnBridgeDefault {
	return &UpdateMsgVpnBridgeDefault{
		_statusCode: code,
	}
}

/*UpdateMsgVpnBridgeDefault handles this case with default header values.

The error response.
*/
type UpdateMsgVpnBridgeDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the update msg vpn bridge default response
func (o *UpdateMsgVpnBridgeDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMsgVpnBridgeDefault) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}][%d] updateMsgVpnBridge default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMsgVpnBridgeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
