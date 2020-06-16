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

// UpdateMsgVpnDmrBridgeReader is a Reader for the UpdateMsgVpnDmrBridge structure.
type UpdateMsgVpnDmrBridgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMsgVpnDmrBridgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateMsgVpnDmrBridgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateMsgVpnDmrBridgeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMsgVpnDmrBridgeOK creates a UpdateMsgVpnDmrBridgeOK with default headers values
func NewUpdateMsgVpnDmrBridgeOK() *UpdateMsgVpnDmrBridgeOK {
	return &UpdateMsgVpnDmrBridgeOK{}
}

/*UpdateMsgVpnDmrBridgeOK handles this case with default header values.

The DMR Bridge object's attributes after being updated, and the request metadata.
*/
type UpdateMsgVpnDmrBridgeOK struct {
	Payload *models.MsgVpnDmrBridgeResponse
}

func (o *UpdateMsgVpnDmrBridgeOK) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName}][%d] updateMsgVpnDmrBridgeOK  %+v", 200, o.Payload)
}

func (o *UpdateMsgVpnDmrBridgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnDmrBridgeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMsgVpnDmrBridgeDefault creates a UpdateMsgVpnDmrBridgeDefault with default headers values
func NewUpdateMsgVpnDmrBridgeDefault(code int) *UpdateMsgVpnDmrBridgeDefault {
	return &UpdateMsgVpnDmrBridgeDefault{
		_statusCode: code,
	}
}

/*UpdateMsgVpnDmrBridgeDefault handles this case with default header values.

The error response.
*/
type UpdateMsgVpnDmrBridgeDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the update msg vpn dmr bridge default response
func (o *UpdateMsgVpnDmrBridgeDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMsgVpnDmrBridgeDefault) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName}][%d] updateMsgVpnDmrBridge default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMsgVpnDmrBridgeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
