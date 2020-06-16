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

// UpdateMsgVpnBridgeRemoteMsgVpnReader is a Reader for the UpdateMsgVpnBridgeRemoteMsgVpn structure.
type UpdateMsgVpnBridgeRemoteMsgVpnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMsgVpnBridgeRemoteMsgVpnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateMsgVpnBridgeRemoteMsgVpnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateMsgVpnBridgeRemoteMsgVpnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMsgVpnBridgeRemoteMsgVpnOK creates a UpdateMsgVpnBridgeRemoteMsgVpnOK with default headers values
func NewUpdateMsgVpnBridgeRemoteMsgVpnOK() *UpdateMsgVpnBridgeRemoteMsgVpnOK {
	return &UpdateMsgVpnBridgeRemoteMsgVpnOK{}
}

/*UpdateMsgVpnBridgeRemoteMsgVpnOK handles this case with default header values.

The Remote Message VPN object's attributes after being updated, and the request metadata.
*/
type UpdateMsgVpnBridgeRemoteMsgVpnOK struct {
	Payload *models.MsgVpnBridgeRemoteMsgVpnResponse
}

func (o *UpdateMsgVpnBridgeRemoteMsgVpnOK) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface}][%d] updateMsgVpnBridgeRemoteMsgVpnOK  %+v", 200, o.Payload)
}

func (o *UpdateMsgVpnBridgeRemoteMsgVpnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnBridgeRemoteMsgVpnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMsgVpnBridgeRemoteMsgVpnDefault creates a UpdateMsgVpnBridgeRemoteMsgVpnDefault with default headers values
func NewUpdateMsgVpnBridgeRemoteMsgVpnDefault(code int) *UpdateMsgVpnBridgeRemoteMsgVpnDefault {
	return &UpdateMsgVpnBridgeRemoteMsgVpnDefault{
		_statusCode: code,
	}
}

/*UpdateMsgVpnBridgeRemoteMsgVpnDefault handles this case with default header values.

The error response.
*/
type UpdateMsgVpnBridgeRemoteMsgVpnDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the update msg vpn bridge remote msg vpn default response
func (o *UpdateMsgVpnBridgeRemoteMsgVpnDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMsgVpnBridgeRemoteMsgVpnDefault) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface}][%d] updateMsgVpnBridgeRemoteMsgVpn default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMsgVpnBridgeRemoteMsgVpnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
