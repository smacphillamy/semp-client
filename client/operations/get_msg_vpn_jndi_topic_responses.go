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

// GetMsgVpnJndiTopicReader is a Reader for the GetMsgVpnJndiTopic structure.
type GetMsgVpnJndiTopicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnJndiTopicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMsgVpnJndiTopicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMsgVpnJndiTopicDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnJndiTopicOK creates a GetMsgVpnJndiTopicOK with default headers values
func NewGetMsgVpnJndiTopicOK() *GetMsgVpnJndiTopicOK {
	return &GetMsgVpnJndiTopicOK{}
}

/*GetMsgVpnJndiTopicOK handles this case with default header values.

The JNDI Topic object's attributes, and the request metadata.
*/
type GetMsgVpnJndiTopicOK struct {
	Payload *models.MsgVpnJndiTopicResponse
}

func (o *GetMsgVpnJndiTopicOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/jndiTopics/{topicName}][%d] getMsgVpnJndiTopicOK  %+v", 200, o.Payload)
}

func (o *GetMsgVpnJndiTopicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnJndiTopicResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnJndiTopicDefault creates a GetMsgVpnJndiTopicDefault with default headers values
func NewGetMsgVpnJndiTopicDefault(code int) *GetMsgVpnJndiTopicDefault {
	return &GetMsgVpnJndiTopicDefault{
		_statusCode: code,
	}
}

/*GetMsgVpnJndiTopicDefault handles this case with default header values.

Error response
*/
type GetMsgVpnJndiTopicDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn jndi topic default response
func (o *GetMsgVpnJndiTopicDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnJndiTopicDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/jndiTopics/{topicName}][%d] getMsgVpnJndiTopic default  %+v", o._statusCode, o.Payload)
}

func (o *GetMsgVpnJndiTopicDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}