// Code generated by go-swagger; DO NOT EDIT.

package dmr_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ExalDraen/semp-client/models"
)

// UpdateDmrClusterLinkReader is a Reader for the UpdateDmrClusterLink structure.
type UpdateDmrClusterLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDmrClusterLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateDmrClusterLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateDmrClusterLinkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDmrClusterLinkOK creates a UpdateDmrClusterLinkOK with default headers values
func NewUpdateDmrClusterLinkOK() *UpdateDmrClusterLinkOK {
	return &UpdateDmrClusterLinkOK{}
}

/*UpdateDmrClusterLinkOK handles this case with default header values.

The Link object's attributes after being updated, and the request metadata.
*/
type UpdateDmrClusterLinkOK struct {
	Payload *models.DmrClusterLinkResponse
}

func (o *UpdateDmrClusterLinkOK) Error() string {
	return fmt.Sprintf("[PATCH /dmrClusters/{dmrClusterName}/links/{remoteNodeName}][%d] updateDmrClusterLinkOK  %+v", 200, o.Payload)
}

func (o *UpdateDmrClusterLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DmrClusterLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDmrClusterLinkDefault creates a UpdateDmrClusterLinkDefault with default headers values
func NewUpdateDmrClusterLinkDefault(code int) *UpdateDmrClusterLinkDefault {
	return &UpdateDmrClusterLinkDefault{
		_statusCode: code,
	}
}

/*UpdateDmrClusterLinkDefault handles this case with default header values.

The error response.
*/
type UpdateDmrClusterLinkDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the update dmr cluster link default response
func (o *UpdateDmrClusterLinkDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDmrClusterLinkDefault) Error() string {
	return fmt.Sprintf("[PATCH /dmrClusters/{dmrClusterName}/links/{remoteNodeName}][%d] updateDmrClusterLink default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDmrClusterLinkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
