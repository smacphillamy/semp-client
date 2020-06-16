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

// DeleteDmrClusterLinkTLSTrustedCommonNameReader is a Reader for the DeleteDmrClusterLinkTLSTrustedCommonName structure.
type DeleteDmrClusterLinkTLSTrustedCommonNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDmrClusterLinkTLSTrustedCommonNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteDmrClusterLinkTLSTrustedCommonNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteDmrClusterLinkTLSTrustedCommonNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDmrClusterLinkTLSTrustedCommonNameOK creates a DeleteDmrClusterLinkTLSTrustedCommonNameOK with default headers values
func NewDeleteDmrClusterLinkTLSTrustedCommonNameOK() *DeleteDmrClusterLinkTLSTrustedCommonNameOK {
	return &DeleteDmrClusterLinkTLSTrustedCommonNameOK{}
}

/*DeleteDmrClusterLinkTLSTrustedCommonNameOK handles this case with default header values.

The request metadata.
*/
type DeleteDmrClusterLinkTLSTrustedCommonNameOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteDmrClusterLinkTLSTrustedCommonNameOK) Error() string {
	return fmt.Sprintf("[DELETE /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames/{tlsTrustedCommonName}][%d] deleteDmrClusterLinkTlsTrustedCommonNameOK  %+v", 200, o.Payload)
}

func (o *DeleteDmrClusterLinkTLSTrustedCommonNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDmrClusterLinkTLSTrustedCommonNameDefault creates a DeleteDmrClusterLinkTLSTrustedCommonNameDefault with default headers values
func NewDeleteDmrClusterLinkTLSTrustedCommonNameDefault(code int) *DeleteDmrClusterLinkTLSTrustedCommonNameDefault {
	return &DeleteDmrClusterLinkTLSTrustedCommonNameDefault{
		_statusCode: code,
	}
}

/*DeleteDmrClusterLinkTLSTrustedCommonNameDefault handles this case with default header values.

The error response.
*/
type DeleteDmrClusterLinkTLSTrustedCommonNameDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete dmr cluster link Tls trusted common name default response
func (o *DeleteDmrClusterLinkTLSTrustedCommonNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDmrClusterLinkTLSTrustedCommonNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames/{tlsTrustedCommonName}][%d] deleteDmrClusterLinkTlsTrustedCommonName default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDmrClusterLinkTLSTrustedCommonNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
