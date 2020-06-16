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

// GetMsgVpnDmrBridgeReader is a Reader for the GetMsgVpnDmrBridge structure.
type GetMsgVpnDmrBridgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnDmrBridgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMsgVpnDmrBridgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMsgVpnDmrBridgeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnDmrBridgeOK creates a GetMsgVpnDmrBridgeOK with default headers values
func NewGetMsgVpnDmrBridgeOK() *GetMsgVpnDmrBridgeOK {
	return &GetMsgVpnDmrBridgeOK{}
}

/*GetMsgVpnDmrBridgeOK handles this case with default header values.

The DMR Bridge object's attributes, and the request metadata.
*/
type GetMsgVpnDmrBridgeOK struct {
	Payload *models.MsgVpnDmrBridgeResponse
}

func (o *GetMsgVpnDmrBridgeOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName}][%d] getMsgVpnDmrBridgeOK  %+v", 200, o.Payload)
}

func (o *GetMsgVpnDmrBridgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnDmrBridgeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnDmrBridgeDefault creates a GetMsgVpnDmrBridgeDefault with default headers values
func NewGetMsgVpnDmrBridgeDefault(code int) *GetMsgVpnDmrBridgeDefault {
	return &GetMsgVpnDmrBridgeDefault{
		_statusCode: code,
	}
}

/*GetMsgVpnDmrBridgeDefault handles this case with default header values.

The error response.
*/
type GetMsgVpnDmrBridgeDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn dmr bridge default response
func (o *GetMsgVpnDmrBridgeDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnDmrBridgeDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/dmrBridges/{remoteNodeName}][%d] getMsgVpnDmrBridge default  %+v", o._statusCode, o.Payload)
}

func (o *GetMsgVpnDmrBridgeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
