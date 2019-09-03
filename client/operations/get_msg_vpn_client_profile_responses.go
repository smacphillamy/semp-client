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

// GetMsgVpnClientProfileReader is a Reader for the GetMsgVpnClientProfile structure.
type GetMsgVpnClientProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnClientProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMsgVpnClientProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMsgVpnClientProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnClientProfileOK creates a GetMsgVpnClientProfileOK with default headers values
func NewGetMsgVpnClientProfileOK() *GetMsgVpnClientProfileOK {
	return &GetMsgVpnClientProfileOK{}
}

/*GetMsgVpnClientProfileOK handles this case with default header values.

The Client Profile object's attributes, and the request metadata.
*/
type GetMsgVpnClientProfileOK struct {
	Payload *models.MsgVpnClientProfileResponse
}

func (o *GetMsgVpnClientProfileOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName}][%d] getMsgVpnClientProfileOK  %+v", 200, o.Payload)
}

func (o *GetMsgVpnClientProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnClientProfileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnClientProfileDefault creates a GetMsgVpnClientProfileDefault with default headers values
func NewGetMsgVpnClientProfileDefault(code int) *GetMsgVpnClientProfileDefault {
	return &GetMsgVpnClientProfileDefault{
		_statusCode: code,
	}
}

/*GetMsgVpnClientProfileDefault handles this case with default header values.

The error response.
*/
type GetMsgVpnClientProfileDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn client profile default response
func (o *GetMsgVpnClientProfileDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnClientProfileDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/clientProfiles/{clientProfileName}][%d] getMsgVpnClientProfile default  %+v", o._statusCode, o.Payload)
}

func (o *GetMsgVpnClientProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
