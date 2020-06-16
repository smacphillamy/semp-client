// Code generated by go-swagger; DO NOT EDIT.

package msg_vpn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/smacphillamy/semp-client/models"
)

// CreateMsgVpnSequencedTopicReader is a Reader for the CreateMsgVpnSequencedTopic structure.
type CreateMsgVpnSequencedTopicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnSequencedTopicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateMsgVpnSequencedTopicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateMsgVpnSequencedTopicDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnSequencedTopicOK creates a CreateMsgVpnSequencedTopicOK with default headers values
func NewCreateMsgVpnSequencedTopicOK() *CreateMsgVpnSequencedTopicOK {
	return &CreateMsgVpnSequencedTopicOK{}
}

/*CreateMsgVpnSequencedTopicOK handles this case with default header values.

The Sequenced Topic object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnSequencedTopicOK struct {
	Payload *models.MsgVpnSequencedTopicResponse
}

func (o *CreateMsgVpnSequencedTopicOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/sequencedTopics][%d] createMsgVpnSequencedTopicOK  %+v", 200, o.Payload)
}

func (o *CreateMsgVpnSequencedTopicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnSequencedTopicResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnSequencedTopicDefault creates a CreateMsgVpnSequencedTopicDefault with default headers values
func NewCreateMsgVpnSequencedTopicDefault(code int) *CreateMsgVpnSequencedTopicDefault {
	return &CreateMsgVpnSequencedTopicDefault{
		_statusCode: code,
	}
}

/*CreateMsgVpnSequencedTopicDefault handles this case with default header values.

The error response.
*/
type CreateMsgVpnSequencedTopicDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn sequenced topic default response
func (o *CreateMsgVpnSequencedTopicDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnSequencedTopicDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/sequencedTopics][%d] createMsgVpnSequencedTopic default  %+v", o._statusCode, o.Payload)
}

func (o *CreateMsgVpnSequencedTopicDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
