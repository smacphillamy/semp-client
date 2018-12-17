// Code generated by go-swagger; DO NOT EDIT.

package msg_vpn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ExalDraen/semp-client/models"
)

// DeleteMsgVpnReader is a Reader for the DeleteMsgVpn structure.
type DeleteMsgVpnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMsgVpnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMsgVpnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnOK creates a DeleteMsgVpnOK with default headers values
func NewDeleteMsgVpnOK() *DeleteMsgVpnOK {
	return &DeleteMsgVpnOK{}
}

/*DeleteMsgVpnOK handles this case with default header values.

The request metadata.
*/
type DeleteMsgVpnOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}][%d] deleteMsgVpnOK  %+v", 200, o.Payload)
}

func (o *DeleteMsgVpnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnDefault creates a DeleteMsgVpnDefault with default headers values
func NewDeleteMsgVpnDefault(code int) *DeleteMsgVpnDefault {
	return &DeleteMsgVpnDefault{
		_statusCode: code,
	}
}

/*DeleteMsgVpnDefault handles this case with default header values.

Error response
*/
type DeleteMsgVpnDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn default response
func (o *DeleteMsgVpnDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}][%d] deleteMsgVpn default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMsgVpnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}