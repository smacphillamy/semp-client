// Code generated by go-swagger; DO NOT EDIT.

package dmr_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/smacphillamy/semp-client/models"
)

// ReplaceDmrClusterReader is a Reader for the ReplaceDmrCluster structure.
type ReplaceDmrClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceDmrClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceDmrClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewReplaceDmrClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReplaceDmrClusterOK creates a ReplaceDmrClusterOK with default headers values
func NewReplaceDmrClusterOK() *ReplaceDmrClusterOK {
	return &ReplaceDmrClusterOK{}
}

/*ReplaceDmrClusterOK handles this case with default header values.

The Cluster object's attributes after being replaced, and the request metadata.
*/
type ReplaceDmrClusterOK struct {
	Payload *models.DmrClusterResponse
}

func (o *ReplaceDmrClusterOK) Error() string {
	return fmt.Sprintf("[PUT /dmrClusters/{dmrClusterName}][%d] replaceDmrClusterOK  %+v", 200, o.Payload)
}

func (o *ReplaceDmrClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DmrClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceDmrClusterDefault creates a ReplaceDmrClusterDefault with default headers values
func NewReplaceDmrClusterDefault(code int) *ReplaceDmrClusterDefault {
	return &ReplaceDmrClusterDefault{
		_statusCode: code,
	}
}

/*ReplaceDmrClusterDefault handles this case with default header values.

The error response.
*/
type ReplaceDmrClusterDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the replace dmr cluster default response
func (o *ReplaceDmrClusterDefault) Code() int {
	return o._statusCode
}

func (o *ReplaceDmrClusterDefault) Error() string {
	return fmt.Sprintf("[PUT /dmrClusters/{dmrClusterName}][%d] replaceDmrCluster default  %+v", o._statusCode, o.Payload)
}

func (o *ReplaceDmrClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
