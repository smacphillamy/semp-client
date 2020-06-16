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

// GetMsgVpnBridgeRemoteMsgVpnReader is a Reader for the GetMsgVpnBridgeRemoteMsgVpn structure.
type GetMsgVpnBridgeRemoteMsgVpnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnBridgeRemoteMsgVpnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMsgVpnBridgeRemoteMsgVpnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMsgVpnBridgeRemoteMsgVpnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnBridgeRemoteMsgVpnOK creates a GetMsgVpnBridgeRemoteMsgVpnOK with default headers values
func NewGetMsgVpnBridgeRemoteMsgVpnOK() *GetMsgVpnBridgeRemoteMsgVpnOK {
	return &GetMsgVpnBridgeRemoteMsgVpnOK{}
}

/*GetMsgVpnBridgeRemoteMsgVpnOK handles this case with default header values.

The Remote Message VPN object's attributes, and the request metadata.
*/
type GetMsgVpnBridgeRemoteMsgVpnOK struct {
	Payload *models.MsgVpnBridgeRemoteMsgVpnResponse
}

func (o *GetMsgVpnBridgeRemoteMsgVpnOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface}][%d] getMsgVpnBridgeRemoteMsgVpnOK  %+v", 200, o.Payload)
}

func (o *GetMsgVpnBridgeRemoteMsgVpnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnBridgeRemoteMsgVpnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnBridgeRemoteMsgVpnDefault creates a GetMsgVpnBridgeRemoteMsgVpnDefault with default headers values
func NewGetMsgVpnBridgeRemoteMsgVpnDefault(code int) *GetMsgVpnBridgeRemoteMsgVpnDefault {
	return &GetMsgVpnBridgeRemoteMsgVpnDefault{
		_statusCode: code,
	}
}

/*GetMsgVpnBridgeRemoteMsgVpnDefault handles this case with default header values.

The error response.
*/
type GetMsgVpnBridgeRemoteMsgVpnDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn bridge remote msg vpn default response
func (o *GetMsgVpnBridgeRemoteMsgVpnDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnBridgeRemoteMsgVpnDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns/{remoteMsgVpnName},{remoteMsgVpnLocation},{remoteMsgVpnInterface}][%d] getMsgVpnBridgeRemoteMsgVpn default  %+v", o._statusCode, o.Payload)
}

func (o *GetMsgVpnBridgeRemoteMsgVpnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
