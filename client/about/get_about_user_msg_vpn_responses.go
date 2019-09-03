// Code generated by go-swagger; DO NOT EDIT.

package about

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ExalDraen/semp-client/models"
)

// GetAboutUserMsgVpnReader is a Reader for the GetAboutUserMsgVpn structure.
type GetAboutUserMsgVpnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAboutUserMsgVpnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAboutUserMsgVpnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAboutUserMsgVpnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAboutUserMsgVpnOK creates a GetAboutUserMsgVpnOK with default headers values
func NewGetAboutUserMsgVpnOK() *GetAboutUserMsgVpnOK {
	return &GetAboutUserMsgVpnOK{}
}

/*GetAboutUserMsgVpnOK handles this case with default header values.

The User Message VPN object's attributes, and the request metadata.
*/
type GetAboutUserMsgVpnOK struct {
	Payload *models.AboutUserMsgVpnResponse
}

func (o *GetAboutUserMsgVpnOK) Error() string {
	return fmt.Sprintf("[GET /about/user/msgVpns/{msgVpnName}][%d] getAboutUserMsgVpnOK  %+v", 200, o.Payload)
}

func (o *GetAboutUserMsgVpnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AboutUserMsgVpnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAboutUserMsgVpnDefault creates a GetAboutUserMsgVpnDefault with default headers values
func NewGetAboutUserMsgVpnDefault(code int) *GetAboutUserMsgVpnDefault {
	return &GetAboutUserMsgVpnDefault{
		_statusCode: code,
	}
}

/*GetAboutUserMsgVpnDefault handles this case with default header values.

The error response.
*/
type GetAboutUserMsgVpnDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get about user msg vpn default response
func (o *GetAboutUserMsgVpnDefault) Code() int {
	return o._statusCode
}

func (o *GetAboutUserMsgVpnDefault) Error() string {
	return fmt.Sprintf("[GET /about/user/msgVpns/{msgVpnName}][%d] getAboutUserMsgVpn default  %+v", o._statusCode, o.Payload)
}

func (o *GetAboutUserMsgVpnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
