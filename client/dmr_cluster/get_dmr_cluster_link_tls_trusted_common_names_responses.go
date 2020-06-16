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

// GetDmrClusterLinkTLSTrustedCommonNamesReader is a Reader for the GetDmrClusterLinkTLSTrustedCommonNames structure.
type GetDmrClusterLinkTLSTrustedCommonNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDmrClusterLinkTLSTrustedCommonNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDmrClusterLinkTLSTrustedCommonNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDmrClusterLinkTLSTrustedCommonNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDmrClusterLinkTLSTrustedCommonNamesOK creates a GetDmrClusterLinkTLSTrustedCommonNamesOK with default headers values
func NewGetDmrClusterLinkTLSTrustedCommonNamesOK() *GetDmrClusterLinkTLSTrustedCommonNamesOK {
	return &GetDmrClusterLinkTLSTrustedCommonNamesOK{}
}

/*GetDmrClusterLinkTLSTrustedCommonNamesOK handles this case with default header values.

The list of Trusted Common Name objects' attributes, and the request metadata.
*/
type GetDmrClusterLinkTLSTrustedCommonNamesOK struct {
	Payload *models.DmrClusterLinkTLSTrustedCommonNamesResponse
}

func (o *GetDmrClusterLinkTLSTrustedCommonNamesOK) Error() string {
	return fmt.Sprintf("[GET /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames][%d] getDmrClusterLinkTlsTrustedCommonNamesOK  %+v", 200, o.Payload)
}

func (o *GetDmrClusterLinkTLSTrustedCommonNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DmrClusterLinkTLSTrustedCommonNamesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDmrClusterLinkTLSTrustedCommonNamesDefault creates a GetDmrClusterLinkTLSTrustedCommonNamesDefault with default headers values
func NewGetDmrClusterLinkTLSTrustedCommonNamesDefault(code int) *GetDmrClusterLinkTLSTrustedCommonNamesDefault {
	return &GetDmrClusterLinkTLSTrustedCommonNamesDefault{
		_statusCode: code,
	}
}

/*GetDmrClusterLinkTLSTrustedCommonNamesDefault handles this case with default header values.

The error response.
*/
type GetDmrClusterLinkTLSTrustedCommonNamesDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get dmr cluster link Tls trusted common names default response
func (o *GetDmrClusterLinkTLSTrustedCommonNamesDefault) Code() int {
	return o._statusCode
}

func (o *GetDmrClusterLinkTLSTrustedCommonNamesDefault) Error() string {
	return fmt.Sprintf("[GET /dmrClusters/{dmrClusterName}/links/{remoteNodeName}/tlsTrustedCommonNames][%d] getDmrClusterLinkTlsTrustedCommonNames default  %+v", o._statusCode, o.Payload)
}

func (o *GetDmrClusterLinkTLSTrustedCommonNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
