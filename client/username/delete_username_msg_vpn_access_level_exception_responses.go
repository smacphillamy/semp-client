// Code generated by go-swagger; DO NOT EDIT.

package username

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/smacphillamy/semp-client/models"
)

// DeleteUsernameMsgVpnAccessLevelExceptionReader is a Reader for the DeleteUsernameMsgVpnAccessLevelException structure.
type DeleteUsernameMsgVpnAccessLevelExceptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsernameMsgVpnAccessLevelExceptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteUsernameMsgVpnAccessLevelExceptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteUsernameMsgVpnAccessLevelExceptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteUsernameMsgVpnAccessLevelExceptionOK creates a DeleteUsernameMsgVpnAccessLevelExceptionOK with default headers values
func NewDeleteUsernameMsgVpnAccessLevelExceptionOK() *DeleteUsernameMsgVpnAccessLevelExceptionOK {
	return &DeleteUsernameMsgVpnAccessLevelExceptionOK{}
}

/*DeleteUsernameMsgVpnAccessLevelExceptionOK handles this case with default header values.

The request metadata.
*/
type DeleteUsernameMsgVpnAccessLevelExceptionOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteUsernameMsgVpnAccessLevelExceptionOK) Error() string {
	return fmt.Sprintf("[DELETE /usernames/{userName}/msgVpnAccessLevelExceptions/{msgVpnName}][%d] deleteUsernameMsgVpnAccessLevelExceptionOK  %+v", 200, o.Payload)
}

func (o *DeleteUsernameMsgVpnAccessLevelExceptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUsernameMsgVpnAccessLevelExceptionDefault creates a DeleteUsernameMsgVpnAccessLevelExceptionDefault with default headers values
func NewDeleteUsernameMsgVpnAccessLevelExceptionDefault(code int) *DeleteUsernameMsgVpnAccessLevelExceptionDefault {
	return &DeleteUsernameMsgVpnAccessLevelExceptionDefault{
		_statusCode: code,
	}
}

/*DeleteUsernameMsgVpnAccessLevelExceptionDefault handles this case with default header values.

The error response.
*/
type DeleteUsernameMsgVpnAccessLevelExceptionDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete username msg vpn access level exception default response
func (o *DeleteUsernameMsgVpnAccessLevelExceptionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteUsernameMsgVpnAccessLevelExceptionDefault) Error() string {
	return fmt.Sprintf("[DELETE /usernames/{userName}/msgVpnAccessLevelExceptions/{msgVpnName}][%d] deleteUsernameMsgVpnAccessLevelException default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteUsernameMsgVpnAccessLevelExceptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
