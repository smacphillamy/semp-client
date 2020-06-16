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

// GetMsgVpnAuthorizationGroupReader is a Reader for the GetMsgVpnAuthorizationGroup structure.
type GetMsgVpnAuthorizationGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMsgVpnAuthorizationGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMsgVpnAuthorizationGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMsgVpnAuthorizationGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMsgVpnAuthorizationGroupOK creates a GetMsgVpnAuthorizationGroupOK with default headers values
func NewGetMsgVpnAuthorizationGroupOK() *GetMsgVpnAuthorizationGroupOK {
	return &GetMsgVpnAuthorizationGroupOK{}
}

/*GetMsgVpnAuthorizationGroupOK handles this case with default header values.

The LDAP Authorization Group object's attributes, and the request metadata.
*/
type GetMsgVpnAuthorizationGroupOK struct {
	Payload *models.MsgVpnAuthorizationGroupResponse
}

func (o *GetMsgVpnAuthorizationGroupOK) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName}][%d] getMsgVpnAuthorizationGroupOK  %+v", 200, o.Payload)
}

func (o *GetMsgVpnAuthorizationGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnAuthorizationGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMsgVpnAuthorizationGroupDefault creates a GetMsgVpnAuthorizationGroupDefault with default headers values
func NewGetMsgVpnAuthorizationGroupDefault(code int) *GetMsgVpnAuthorizationGroupDefault {
	return &GetMsgVpnAuthorizationGroupDefault{
		_statusCode: code,
	}
}

/*GetMsgVpnAuthorizationGroupDefault handles this case with default header values.

The error response.
*/
type GetMsgVpnAuthorizationGroupDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the get msg vpn authorization group default response
func (o *GetMsgVpnAuthorizationGroupDefault) Code() int {
	return o._statusCode
}

func (o *GetMsgVpnAuthorizationGroupDefault) Error() string {
	return fmt.Sprintf("[GET /msgVpns/{msgVpnName}/authorizationGroups/{authorizationGroupName}][%d] getMsgVpnAuthorizationGroup default  %+v", o._statusCode, o.Payload)
}

func (o *GetMsgVpnAuthorizationGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
