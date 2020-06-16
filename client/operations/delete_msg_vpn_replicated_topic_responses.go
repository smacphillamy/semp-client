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

// DeleteMsgVpnReplicatedTopicReader is a Reader for the DeleteMsgVpnReplicatedTopic structure.
type DeleteMsgVpnReplicatedTopicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnReplicatedTopicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMsgVpnReplicatedTopicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMsgVpnReplicatedTopicDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnReplicatedTopicOK creates a DeleteMsgVpnReplicatedTopicOK with default headers values
func NewDeleteMsgVpnReplicatedTopicOK() *DeleteMsgVpnReplicatedTopicOK {
	return &DeleteMsgVpnReplicatedTopicOK{}
}

/*DeleteMsgVpnReplicatedTopicOK handles this case with default header values.

The request metadata.
*/
type DeleteMsgVpnReplicatedTopicOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnReplicatedTopicOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic}][%d] deleteMsgVpnReplicatedTopicOK  %+v", 200, o.Payload)
}

func (o *DeleteMsgVpnReplicatedTopicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnReplicatedTopicDefault creates a DeleteMsgVpnReplicatedTopicDefault with default headers values
func NewDeleteMsgVpnReplicatedTopicDefault(code int) *DeleteMsgVpnReplicatedTopicDefault {
	return &DeleteMsgVpnReplicatedTopicDefault{
		_statusCode: code,
	}
}

/*DeleteMsgVpnReplicatedTopicDefault handles this case with default header values.

The error response.
*/
type DeleteMsgVpnReplicatedTopicDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn replicated topic default response
func (o *DeleteMsgVpnReplicatedTopicDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnReplicatedTopicDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/replicatedTopics/{replicatedTopic}][%d] deleteMsgVpnReplicatedTopic default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMsgVpnReplicatedTopicDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
