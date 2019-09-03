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

// ReplaceUsernameMsgVpnAccessLevelExceptionReader is a Reader for the ReplaceUsernameMsgVpnAccessLevelException structure.
type ReplaceUsernameMsgVpnAccessLevelExceptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceUsernameMsgVpnAccessLevelExceptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceUsernameMsgVpnAccessLevelExceptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewReplaceUsernameMsgVpnAccessLevelExceptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceUsernameMsgVpnAccessLevelExceptionOK creates a ReplaceUsernameMsgVpnAccessLevelExceptionOK with default headers values
func NewReplaceUsernameMsgVpnAccessLevelExceptionOK() *ReplaceUsernameMsgVpnAccessLevelExceptionOK {
	return &ReplaceUsernameMsgVpnAccessLevelExceptionOK{}
}

/*ReplaceUsernameMsgVpnAccessLevelExceptionOK handles this case with default header values.

The Message VPN Access Level Exception object's attributes after being replaced, and the request metadata.
*/
type ReplaceUsernameMsgVpnAccessLevelExceptionOK struct {
	Payload *models.UsernameMsgVpnAccessLevelExceptionResponse
}

func (o *ReplaceUsernameMsgVpnAccessLevelExceptionOK) Error() string {
	return fmt.Sprintf("[PUT /usernames/{userName}/msgVpnAccessLevelExceptions/{msgVpnName}][%d] replaceUsernameMsgVpnAccessLevelExceptionOK  %+v", 200, o.Payload)
}

func (o *ReplaceUsernameMsgVpnAccessLevelExceptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UsernameMsgVpnAccessLevelExceptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceUsernameMsgVpnAccessLevelExceptionDefault creates a ReplaceUsernameMsgVpnAccessLevelExceptionDefault with default headers values
func NewReplaceUsernameMsgVpnAccessLevelExceptionDefault(code int) *ReplaceUsernameMsgVpnAccessLevelExceptionDefault {
	return &ReplaceUsernameMsgVpnAccessLevelExceptionDefault{
		_statusCode: code,
	}
}

/*ReplaceUsernameMsgVpnAccessLevelExceptionDefault handles this case with default header values.

The error response.
*/
type ReplaceUsernameMsgVpnAccessLevelExceptionDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the replace username msg vpn access level exception default response
func (o *ReplaceUsernameMsgVpnAccessLevelExceptionDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceUsernameMsgVpnAccessLevelExceptionDefault) Error() string {
	return fmt.Sprintf("[PUT /usernames/{userName}/msgVpnAccessLevelExceptions/{msgVpnName}][%d] replaceUsernameMsgVpnAccessLevelException default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceUsernameMsgVpnAccessLevelExceptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}