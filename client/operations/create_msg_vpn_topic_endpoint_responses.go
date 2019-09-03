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

// CreateMsgVpnTopicEndpointReader is a Reader for the CreateMsgVpnTopicEndpoint structure.
type CreateMsgVpnTopicEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnTopicEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateMsgVpnTopicEndpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateMsgVpnTopicEndpointDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnTopicEndpointOK creates a CreateMsgVpnTopicEndpointOK with default headers values
func NewCreateMsgVpnTopicEndpointOK() *CreateMsgVpnTopicEndpointOK {
	return &CreateMsgVpnTopicEndpointOK{}
}

/*CreateMsgVpnTopicEndpointOK handles this case with default header values.

The Topic Endpoint object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnTopicEndpointOK struct {
	Payload *models.MsgVpnTopicEndpointResponse
}

func (o *CreateMsgVpnTopicEndpointOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/topicEndpoints][%d] createMsgVpnTopicEndpointOK  %+v", 200, o.Payload)
}

func (o *CreateMsgVpnTopicEndpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnTopicEndpointResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnTopicEndpointDefault creates a CreateMsgVpnTopicEndpointDefault with default headers values
func NewCreateMsgVpnTopicEndpointDefault(code int) *CreateMsgVpnTopicEndpointDefault {
	return &CreateMsgVpnTopicEndpointDefault{
		_statusCode: code,
	}
}

/*CreateMsgVpnTopicEndpointDefault handles this case with default header values.

The error response.
*/
type CreateMsgVpnTopicEndpointDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn topic endpoint default response
func (o *CreateMsgVpnTopicEndpointDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnTopicEndpointDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/topicEndpoints][%d] createMsgVpnTopicEndpoint default  %+v", o._statusCode, o.Payload)
}

func (o *CreateMsgVpnTopicEndpointDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
