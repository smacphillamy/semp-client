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

// UpdateMsgVpnClientUsernameReader is a Reader for the UpdateMsgVpnClientUsername structure.
type UpdateMsgVpnClientUsernameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMsgVpnClientUsernameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateMsgVpnClientUsernameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateMsgVpnClientUsernameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMsgVpnClientUsernameOK creates a UpdateMsgVpnClientUsernameOK with default headers values
func NewUpdateMsgVpnClientUsernameOK() *UpdateMsgVpnClientUsernameOK {
	return &UpdateMsgVpnClientUsernameOK{}
}

/*UpdateMsgVpnClientUsernameOK handles this case with default header values.

The Client Username object's attributes after being updated, and the request metadata.
*/
type UpdateMsgVpnClientUsernameOK struct {
	Payload *models.MsgVpnClientUsernameResponse
}

func (o *UpdateMsgVpnClientUsernameOK) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/clientUsernames/{clientUsername}][%d] updateMsgVpnClientUsernameOK  %+v", 200, o.Payload)
}

func (o *UpdateMsgVpnClientUsernameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnClientUsernameResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMsgVpnClientUsernameDefault creates a UpdateMsgVpnClientUsernameDefault with default headers values
func NewUpdateMsgVpnClientUsernameDefault(code int) *UpdateMsgVpnClientUsernameDefault {
	return &UpdateMsgVpnClientUsernameDefault{
		_statusCode: code,
	}
}

/*UpdateMsgVpnClientUsernameDefault handles this case with default header values.

The error response.
*/
type UpdateMsgVpnClientUsernameDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the update msg vpn client username default response
func (o *UpdateMsgVpnClientUsernameDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMsgVpnClientUsernameDefault) Error() string {
	return fmt.Sprintf("[PATCH /msgVpns/{msgVpnName}/clientUsernames/{clientUsername}][%d] updateMsgVpnClientUsername default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateMsgVpnClientUsernameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
