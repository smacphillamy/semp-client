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

// GetDmrClusterLinkReader is a Reader for the GetDmrClusterLink structure.
type GetDmrClusterLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDmrClusterLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDmrClusterLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDmrClusterLinkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDmrClusterLinkOK creates a GetDmrClusterLinkOK with default headers values
func NewGetDmrClusterLinkOK() *GetDmrClusterLinkOK {
	return &GetDmrClusterLinkOK{}
}

/*GetDmrClusterLinkOK handles this case with default header values.

The Link object's attributes, and the request metadata.
*/
type GetDmrClusterLinkOK struct {
	Payload *models.DmrClusterLinkResponse
}

func (o *GetDmrClusterLinkOK) Error() string {
	return fmt.Sprintf("[GET /dmrClusters/{dmrClusterName}/links/{remoteNodeName}][%d] getDmrClusterLinkOK  %+v", 200, o.Payload)
}

func (o *GetDmrClusterLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DmrClusterLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDmrClusterLinkDefault creates a GetDmrClusterLinkDefault with default headers values
func NewGetDmrClusterLinkDefault(code int) *GetDmrClusterLinkDefault {
	return &GetDmrClusterLinkDefault{
		_statusCode: code,
	}
}

/*GetDmrClusterLinkDefault handles this case with default header values.

The error response.
*/
type GetDmrClusterLinkDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get dmr cluster link default response
func (o *GetDmrClusterLinkDefault) Code() int {
	return o._statusCode
}

func (o *GetDmrClusterLinkDefault) Error() string {
	return fmt.Sprintf("[GET /dmrClusters/{dmrClusterName}/links/{remoteNodeName}][%d] getDmrClusterLink default  %+v", o._statusCode, o.Payload)
}

func (o *GetDmrClusterLinkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}