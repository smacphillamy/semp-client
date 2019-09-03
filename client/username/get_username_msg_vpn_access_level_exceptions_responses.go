// Code generated by go-swagger; DO NOT EDIT.

package username

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ExalDraen/semp-client/models"
)

// GetUsernameMsgVpnAccessLevelExceptionsReader is a Reader for the GetUsernameMsgVpnAccessLevelExceptions structure.
type GetUsernameMsgVpnAccessLevelExceptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsernameMsgVpnAccessLevelExceptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsernameMsgVpnAccessLevelExceptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetUsernameMsgVpnAccessLevelExceptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUsernameMsgVpnAccessLevelExceptionsOK creates a GetUsernameMsgVpnAccessLevelExceptionsOK with default headers values
func NewGetUsernameMsgVpnAccessLevelExceptionsOK() *GetUsernameMsgVpnAccessLevelExceptionsOK {
	return &GetUsernameMsgVpnAccessLevelExceptionsOK{}
}

/*GetUsernameMsgVpnAccessLevelExceptionsOK handles this case with default header values.

The list of Message VPN Access Level Exception objects' attributes, and the request metadata.
*/
type GetUsernameMsgVpnAccessLevelExceptionsOK struct {
	Payload *models.UsernameMsgVpnAccessLevelExceptionsResponse
}

func (o *GetUsernameMsgVpnAccessLevelExceptionsOK) Error() string {
	return fmt.Sprintf("[GET /usernames/{userName}/msgVpnAccessLevelExceptions][%d] getUsernameMsgVpnAccessLevelExceptionsOK  %+v", 200, o.Payload)
}

func (o *GetUsernameMsgVpnAccessLevelExceptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UsernameMsgVpnAccessLevelExceptionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsernameMsgVpnAccessLevelExceptionsDefault creates a GetUsernameMsgVpnAccessLevelExceptionsDefault with default headers values
func NewGetUsernameMsgVpnAccessLevelExceptionsDefault(code int) *GetUsernameMsgVpnAccessLevelExceptionsDefault {
	return &GetUsernameMsgVpnAccessLevelExceptionsDefault{
		_statusCode: code,
	}
}

/*GetUsernameMsgVpnAccessLevelExceptionsDefault handles this case with default header values.

The error response.
*/
type GetUsernameMsgVpnAccessLevelExceptionsDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get username msg vpn access level exceptions default response
func (o *GetUsernameMsgVpnAccessLevelExceptionsDefault) Code() int {
	return o._statusCode
}

func (o *GetUsernameMsgVpnAccessLevelExceptionsDefault) Error() string {
	return fmt.Sprintf("[GET /usernames/{userName}/msgVpnAccessLevelExceptions][%d] getUsernameMsgVpnAccessLevelExceptions default  %+v", o._statusCode, o.Payload)
}

func (o *GetUsernameMsgVpnAccessLevelExceptionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
