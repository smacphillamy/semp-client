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

// CreateMsgVpnReplayLogReader is a Reader for the CreateMsgVpnReplayLog structure.
type CreateMsgVpnReplayLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnReplayLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateMsgVpnReplayLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateMsgVpnReplayLogDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnReplayLogOK creates a CreateMsgVpnReplayLogOK with default headers values
func NewCreateMsgVpnReplayLogOK() *CreateMsgVpnReplayLogOK {
	return &CreateMsgVpnReplayLogOK{}
}

/*CreateMsgVpnReplayLogOK handles this case with default header values.

The Replay Log object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnReplayLogOK struct {
	Payload *models.MsgVpnReplayLogResponse
}

func (o *CreateMsgVpnReplayLogOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/replayLogs][%d] createMsgVpnReplayLogOK  %+v", 200, o.Payload)
}

func (o *CreateMsgVpnReplayLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnReplayLogResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnReplayLogDefault creates a CreateMsgVpnReplayLogDefault with default headers values
func NewCreateMsgVpnReplayLogDefault(code int) *CreateMsgVpnReplayLogDefault {
	return &CreateMsgVpnReplayLogDefault{
		_statusCode: code,
	}
}

/*CreateMsgVpnReplayLogDefault handles this case with default header values.

The error response.
*/
type CreateMsgVpnReplayLogDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn replay log default response
func (o *CreateMsgVpnReplayLogDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnReplayLogDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/replayLogs][%d] createMsgVpnReplayLog default  %+v", o._statusCode, o.Payload)
}

func (o *CreateMsgVpnReplayLogDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
