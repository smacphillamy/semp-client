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

// GetMsgVpnJndiConnectionFactoryReader is a Reader for the GetMsgVpnJndiConnectionFactory structure.
type GetMsgVpnJndiConnectionFactoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnJndiConnectionFactoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMsgVpnJndiConnectionFactoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMsgVpnJndiConnectionFactoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnJndiConnectionFactoryOK creates a GetMsgVpnJndiConnectionFactoryOK with default headers values
func NewGetMsgVpnJndiConnectionFactoryOK() *GetMsgVpnJndiConnectionFactoryOK {
	return &GetMsgVpnJndiConnectionFactoryOK{}
}

/*GetMsgVpnJndiConnectionFactoryOK handles this case with default header values.

The JNDI Connection Factory object's attributes, and the request metadata.
*/
type GetMsgVpnJndiConnectionFactoryOK struct {
	Payload *models.MsgVpnJndiConnectionFactoryResponse
}

func (o *GetMsgVpnJndiConnectionFactoryOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName}][%d] getMsgVpnJndiConnectionFactoryOK  %+v", 200, o.Payload)
}

func (o *GetMsgVpnJndiConnectionFactoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnJndiConnectionFactoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnJndiConnectionFactoryDefault creates a GetMsgVpnJndiConnectionFactoryDefault with default headers values
func NewGetMsgVpnJndiConnectionFactoryDefault(code int) *GetMsgVpnJndiConnectionFactoryDefault {
	return &GetMsgVpnJndiConnectionFactoryDefault{
		_statusCode: code,
	}
}

/*GetMsgVpnJndiConnectionFactoryDefault handles this case with default header values.

The error response.
*/
type GetMsgVpnJndiConnectionFactoryDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn jndi connection factory default response
func (o *GetMsgVpnJndiConnectionFactoryDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnJndiConnectionFactoryDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName}][%d] getMsgVpnJndiConnectionFactory default  %+v", o._statusCode, o.Payload)
}

func (o *GetMsgVpnJndiConnectionFactoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
