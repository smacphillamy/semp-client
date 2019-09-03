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

// GetMsgVpnJndiTopicsReader is a Reader for the GetMsgVpnJndiTopics structure.
type GetMsgVpnJndiTopicsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnJndiTopicsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMsgVpnJndiTopicsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMsgVpnJndiTopicsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnJndiTopicsOK creates a GetMsgVpnJndiTopicsOK with default headers values
func NewGetMsgVpnJndiTopicsOK() *GetMsgVpnJndiTopicsOK {
	return &GetMsgVpnJndiTopicsOK{}
}

/*GetMsgVpnJndiTopicsOK handles this case with default header values.

The list of JNDI Topic objects' attributes, and the request metadata.
*/
type GetMsgVpnJndiTopicsOK struct {
	Payload *models.MsgVpnJndiTopicsResponse
}

func (o *GetMsgVpnJndiTopicsOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/jndiTopics][%d] getMsgVpnJndiTopicsOK  %+v", 200, o.Payload)
}

func (o *GetMsgVpnJndiTopicsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnJndiTopicsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnJndiTopicsDefault creates a GetMsgVpnJndiTopicsDefault with default headers values
func NewGetMsgVpnJndiTopicsDefault(code int) *GetMsgVpnJndiTopicsDefault {
	return &GetMsgVpnJndiTopicsDefault{
		_statusCode: code,
	}
}

/*GetMsgVpnJndiTopicsDefault handles this case with default header values.

The error response.
*/
type GetMsgVpnJndiTopicsDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn jndi topics default response
func (o *GetMsgVpnJndiTopicsDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnJndiTopicsDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/jndiTopics][%d] getMsgVpnJndiTopics default  %+v", o._statusCode, o.Payload)
}

func (o *GetMsgVpnJndiTopicsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
