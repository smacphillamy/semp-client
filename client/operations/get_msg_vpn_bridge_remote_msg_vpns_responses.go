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

// GetMsgVpnBridgeRemoteMsgVpnsReader is a Reader for the GetMsgVpnBridgeRemoteMsgVpns structure.
type GetMsgVpnBridgeRemoteMsgVpnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnBridgeRemoteMsgVpnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMsgVpnBridgeRemoteMsgVpnsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMsgVpnBridgeRemoteMsgVpnsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnBridgeRemoteMsgVpnsOK creates a GetMsgVpnBridgeRemoteMsgVpnsOK with default headers values
func NewGetMsgVpnBridgeRemoteMsgVpnsOK() *GetMsgVpnBridgeRemoteMsgVpnsOK {
	return &GetMsgVpnBridgeRemoteMsgVpnsOK{}
}

/*GetMsgVpnBridgeRemoteMsgVpnsOK handles this case with default header values.

The list of Remote Message VPN objects' attributes, and the request metadata.
*/
type GetMsgVpnBridgeRemoteMsgVpnsOK struct {
	Payload *models.MsgVpnBridgeRemoteMsgVpnsResponse
}

func (o *GetMsgVpnBridgeRemoteMsgVpnsOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns][%d] getMsgVpnBridgeRemoteMsgVpnsOK  %+v", 200, o.Payload)
}

func (o *GetMsgVpnBridgeRemoteMsgVpnsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnBridgeRemoteMsgVpnsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnBridgeRemoteMsgVpnsDefault creates a GetMsgVpnBridgeRemoteMsgVpnsDefault with default headers values
func NewGetMsgVpnBridgeRemoteMsgVpnsDefault(code int) *GetMsgVpnBridgeRemoteMsgVpnsDefault {
	return &GetMsgVpnBridgeRemoteMsgVpnsDefault{
		_statusCode: code,
	}
}

/*GetMsgVpnBridgeRemoteMsgVpnsDefault handles this case with default header values.

The error response.
*/
type GetMsgVpnBridgeRemoteMsgVpnsDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn bridge remote msg vpns default response
func (o *GetMsgVpnBridgeRemoteMsgVpnsDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnBridgeRemoteMsgVpnsDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/remoteMsgVpns][%d] getMsgVpnBridgeRemoteMsgVpns default  %+v", o._statusCode, o.Payload)
}

func (o *GetMsgVpnBridgeRemoteMsgVpnsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
