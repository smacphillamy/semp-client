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

// GetAboutUserMsgVpnsReader is a Reader for the GetAboutUserMsgVpns structure.
type GetAboutUserMsgVpnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAboutUserMsgVpnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAboutUserMsgVpnsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAboutUserMsgVpnsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAboutUserMsgVpnsOK creates a GetAboutUserMsgVpnsOK with default headers values
func NewGetAboutUserMsgVpnsOK() *GetAboutUserMsgVpnsOK {
	return &GetAboutUserMsgVpnsOK{}
}

/*GetAboutUserMsgVpnsOK handles this case with default header values.

The list of User Message VPN objects' attributes, and the request metadata.
*/
type GetAboutUserMsgVpnsOK struct {
	Payload *models.AboutUserMsgVpnsResponse
}

func (o *GetAboutUserMsgVpnsOK) Error() string {
	return fmt.Sprintf("[GET /about/user/msgVpns][%d] getAboutUserMsgVpnsOK  %+v", 200, o.Payload)
}

func (o *GetAboutUserMsgVpnsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AboutUserMsgVpnsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAboutUserMsgVpnsDefault creates a GetAboutUserMsgVpnsDefault with default headers values
func NewGetAboutUserMsgVpnsDefault(code int) *GetAboutUserMsgVpnsDefault {
	return &GetAboutUserMsgVpnsDefault{
		_statusCode: code,
	}
}

/*GetAboutUserMsgVpnsDefault handles this case with default header values.

The error response.
*/
type GetAboutUserMsgVpnsDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get about user msg vpns default response
func (o *GetAboutUserMsgVpnsDefault) Code() int {
	return o._statusCode
}

func (o *GetAboutUserMsgVpnsDefault) Error() string {
	return fmt.Sprintf("[GET /about/user/msgVpns][%d] getAboutUserMsgVpns default  %+v", o._statusCode, o.Payload)
}

func (o *GetAboutUserMsgVpnsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
