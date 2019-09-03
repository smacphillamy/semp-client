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

// GetMsgVpnBridgesReader is a Reader for the GetMsgVpnBridges structure.
type GetMsgVpnBridgesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnBridgesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMsgVpnBridgesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMsgVpnBridgesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnBridgesOK creates a GetMsgVpnBridgesOK with default headers values
func NewGetMsgVpnBridgesOK() *GetMsgVpnBridgesOK {
	return &GetMsgVpnBridgesOK{}
}

/*GetMsgVpnBridgesOK handles this case with default header values.

The list of Bridge objects' attributes, and the request metadata.
*/
type GetMsgVpnBridgesOK struct {
	Payload *models.MsgVpnBridgesResponse
}

func (o *GetMsgVpnBridgesOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/bridges][%d] getMsgVpnBridgesOK  %+v", 200, o.Payload)
}

func (o *GetMsgVpnBridgesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnBridgesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnBridgesDefault creates a GetMsgVpnBridgesDefault with default headers values
func NewGetMsgVpnBridgesDefault(code int) *GetMsgVpnBridgesDefault {
	return &GetMsgVpnBridgesDefault{
		_statusCode: code,
	}
}

/*GetMsgVpnBridgesDefault handles this case with default header values.

The error response.
*/
type GetMsgVpnBridgesDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn bridges default response
func (o *GetMsgVpnBridgesDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnBridgesDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/bridges][%d] getMsgVpnBridges default  %+v", o._statusCode, o.Payload)
}

func (o *GetMsgVpnBridgesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
