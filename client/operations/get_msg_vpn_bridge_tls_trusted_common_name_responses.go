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

// GetMsgVpnBridgeTLSTrustedCommonNameReader is a Reader for the GetMsgVpnBridgeTLSTrustedCommonName structure.
type GetMsgVpnBridgeTLSTrustedCommonNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnBridgeTLSTrustedCommonNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMsgVpnBridgeTLSTrustedCommonNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMsgVpnBridgeTLSTrustedCommonNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnBridgeTLSTrustedCommonNameOK creates a GetMsgVpnBridgeTLSTrustedCommonNameOK with default headers values
func NewGetMsgVpnBridgeTLSTrustedCommonNameOK() *GetMsgVpnBridgeTLSTrustedCommonNameOK {
	return &GetMsgVpnBridgeTLSTrustedCommonNameOK{}
}

/*GetMsgVpnBridgeTLSTrustedCommonNameOK handles this case with default header values.

The Trusted Common Name object's attributes, and the request metadata.
*/
type GetMsgVpnBridgeTLSTrustedCommonNameOK struct {
	Payload *models.MsgVpnBridgeTLSTrustedCommonNameResponse
}

func (o *GetMsgVpnBridgeTLSTrustedCommonNameOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName}][%d] getMsgVpnBridgeTlsTrustedCommonNameOK  %+v", 200, o.Payload)
}

func (o *GetMsgVpnBridgeTLSTrustedCommonNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnBridgeTLSTrustedCommonNameResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnBridgeTLSTrustedCommonNameDefault creates a GetMsgVpnBridgeTLSTrustedCommonNameDefault with default headers values
func NewGetMsgVpnBridgeTLSTrustedCommonNameDefault(code int) *GetMsgVpnBridgeTLSTrustedCommonNameDefault {
	return &GetMsgVpnBridgeTLSTrustedCommonNameDefault{
		_statusCode: code,
	}
}

/*GetMsgVpnBridgeTLSTrustedCommonNameDefault handles this case with default header values.

The error response.
*/
type GetMsgVpnBridgeTLSTrustedCommonNameDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn bridge Tls trusted common name default response
func (o *GetMsgVpnBridgeTLSTrustedCommonNameDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnBridgeTLSTrustedCommonNameDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/bridges/{bridgeName},{bridgeVirtualRouter}/tlsTrustedCommonNames/{tlsTrustedCommonName}][%d] getMsgVpnBridgeTlsTrustedCommonName default  %+v", o._statusCode, o.Payload)
}

func (o *GetMsgVpnBridgeTLSTrustedCommonNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
