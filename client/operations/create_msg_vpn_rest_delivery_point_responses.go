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

// CreateMsgVpnRestDeliveryPointReader is a Reader for the CreateMsgVpnRestDeliveryPoint structure.
type CreateMsgVpnRestDeliveryPointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnRestDeliveryPointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateMsgVpnRestDeliveryPointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateMsgVpnRestDeliveryPointDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnRestDeliveryPointOK creates a CreateMsgVpnRestDeliveryPointOK with default headers values
func NewCreateMsgVpnRestDeliveryPointOK() *CreateMsgVpnRestDeliveryPointOK {
	return &CreateMsgVpnRestDeliveryPointOK{}
}

/*CreateMsgVpnRestDeliveryPointOK handles this case with default header values.

The REST Delivery Point object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnRestDeliveryPointOK struct {
	Payload *models.MsgVpnRestDeliveryPointResponse
}

func (o *CreateMsgVpnRestDeliveryPointOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/restDeliveryPoints][%d] createMsgVpnRestDeliveryPointOK  %+v", 200, o.Payload)
}

func (o *CreateMsgVpnRestDeliveryPointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnRestDeliveryPointResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnRestDeliveryPointDefault creates a CreateMsgVpnRestDeliveryPointDefault with default headers values
func NewCreateMsgVpnRestDeliveryPointDefault(code int) *CreateMsgVpnRestDeliveryPointDefault {
	return &CreateMsgVpnRestDeliveryPointDefault{
		_statusCode: code,
	}
}

/*CreateMsgVpnRestDeliveryPointDefault handles this case with default header values.

The error response.
*/
type CreateMsgVpnRestDeliveryPointDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn rest delivery point default response
func (o *CreateMsgVpnRestDeliveryPointDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnRestDeliveryPointDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/restDeliveryPoints][%d] createMsgVpnRestDeliveryPoint default  %+v", o._statusCode, o.Payload)
}

func (o *CreateMsgVpnRestDeliveryPointDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
