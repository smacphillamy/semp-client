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

// DeleteMsgVpnACLProfileReader is a Reader for the DeleteMsgVpnACLProfile structure.
type DeleteMsgVpnACLProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMsgVpnACLProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMsgVpnACLProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMsgVpnACLProfileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMsgVpnACLProfileOK creates a DeleteMsgVpnACLProfileOK with default headers values
func NewDeleteMsgVpnACLProfileOK() *DeleteMsgVpnACLProfileOK {
	return &DeleteMsgVpnACLProfileOK{}
}

/*DeleteMsgVpnACLProfileOK handles this case with default header values.

The request metadata.
*/
type DeleteMsgVpnACLProfileOK struct {
	Payload *models.SempMetaOnlyResponse
}

func (o *DeleteMsgVpnACLProfileOK) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}][%d] deleteMsgVpnAclProfileOK  %+v", 200, o.Payload)
}

func (o *DeleteMsgVpnACLProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMsgVpnACLProfileDefault creates a DeleteMsgVpnACLProfileDefault with default headers values
func NewDeleteMsgVpnACLProfileDefault(code int) *DeleteMsgVpnACLProfileDefault {
	return &DeleteMsgVpnACLProfileDefault{
		_statusCode: code,
	}
}

/*DeleteMsgVpnACLProfileDefault handles this case with default header values.

The error response.
*/
type DeleteMsgVpnACLProfileDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the delete msg vpn Acl profile default response
func (o *DeleteMsgVpnACLProfileDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMsgVpnACLProfileDefault) Error() string {
	return fmt.Sprintf("[DELETE /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}][%d] deleteMsgVpnAclProfile default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMsgVpnACLProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
