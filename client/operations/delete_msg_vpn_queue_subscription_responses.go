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

// DeleteMsgVpnQueueSubscriptionReader is a Reader for the DeleteMsgVpnQueueSubscription structure.
type DeleteMsgVpnQueueSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnQueueSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMsgVpnQueueSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMsgVpnQueueSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnQueueSubscriptionOK creates a DeleteMsgVpnQueueSubscriptionOK with default headers values
func NewDeleteMsgVpnQueueSubscriptionOK() *DeleteMsgVpnQueueSubscriptionOK {
	return &DeleteMsgVpnQueueSubscriptionOK{}
}

/*DeleteMsgVpnQueueSubscriptionOK handles this case with default header values.

The request metadata.
*/
type DeleteMsgVpnQueueSubscriptionOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnQueueSubscriptionOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic}][%d] deleteMsgVpnQueueSubscriptionOK  %+v", 200, o.Payload)
}

func (o *DeleteMsgVpnQueueSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnQueueSubscriptionDefault creates a DeleteMsgVpnQueueSubscriptionDefault with default headers values
func NewDeleteMsgVpnQueueSubscriptionDefault(code int) *DeleteMsgVpnQueueSubscriptionDefault {
	return &DeleteMsgVpnQueueSubscriptionDefault{
		_statusCode: code,
	}
}

/*DeleteMsgVpnQueueSubscriptionDefault handles this case with default header values.

The error response.
*/
type DeleteMsgVpnQueueSubscriptionDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn queue subscription default response
func (o *DeleteMsgVpnQueueSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnQueueSubscriptionDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/queues/{queueName}/subscriptions/{subscriptionTopic}][%d] deleteMsgVpnQueueSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMsgVpnQueueSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
