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

// ReplaceMsgVpnRestDeliveryPointReader is a Reader for the ReplaceMsgVpnRestDeliveryPoint structure.
type ReplaceMsgVpnRestDeliveryPointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceMsgVpnRestDeliveryPointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceMsgVpnRestDeliveryPointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewReplaceMsgVpnRestDeliveryPointDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceMsgVpnRestDeliveryPointOK creates a ReplaceMsgVpnRestDeliveryPointOK with default headers values
func NewReplaceMsgVpnRestDeliveryPointOK() *ReplaceMsgVpnRestDeliveryPointOK {
	return &ReplaceMsgVpnRestDeliveryPointOK{}
}

/*ReplaceMsgVpnRestDeliveryPointOK handles this case with default header values.

The REST Delivery Point object's attributes after being replaced, and the request metadata.
*/
type ReplaceMsgVpnRestDeliveryPointOK struct {
	Payload *models.MsgVpnRestDeliveryPointResponse
}

func (o *ReplaceMsgVpnRestDeliveryPointOK) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}][%d] replaceMsgVpnRestDeliveryPointOK  %+v", 200, o.Payload)
}

func (o *ReplaceMsgVpnRestDeliveryPointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnRestDeliveryPointResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceMsgVpnRestDeliveryPointDefault creates a ReplaceMsgVpnRestDeliveryPointDefault with default headers values
func NewReplaceMsgVpnRestDeliveryPointDefault(code int) *ReplaceMsgVpnRestDeliveryPointDefault {
	return &ReplaceMsgVpnRestDeliveryPointDefault{
		_statusCode: code,
	}
}

/*ReplaceMsgVpnRestDeliveryPointDefault handles this case with default header values.

The error response.
*/
type ReplaceMsgVpnRestDeliveryPointDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the replace msg vpn rest delivery point default response
func (o *ReplaceMsgVpnRestDeliveryPointDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceMsgVpnRestDeliveryPointDefault) Error() string {
	return fmt.Sprintf("[PUT /msgVpns/{msgVpnName}/restDeliveryPoints/{restDeliveryPointName}][%d] replaceMsgVpnRestDeliveryPoint default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceMsgVpnRestDeliveryPointDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
