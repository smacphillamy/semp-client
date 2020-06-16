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

// CreateMsgVpnACLProfilePublishExceptionReader is a Reader for the CreateMsgVpnACLProfilePublishException structure.
type CreateMsgVpnACLProfilePublishExceptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMsgVpnACLProfilePublishExceptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateMsgVpnACLProfilePublishExceptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateMsgVpnACLProfilePublishExceptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateMsgVpnACLProfilePublishExceptionOK creates a CreateMsgVpnACLProfilePublishExceptionOK with default headers values
func NewCreateMsgVpnACLProfilePublishExceptionOK() *CreateMsgVpnACLProfilePublishExceptionOK {
	return &CreateMsgVpnACLProfilePublishExceptionOK{}
}

/*CreateMsgVpnACLProfilePublishExceptionOK handles this case with default header values.

The Publish Topic Exception object's attributes after being created, and the request metadata.
*/
type CreateMsgVpnACLProfilePublishExceptionOK struct {
	Payload *models.MsgVpnACLProfilePublishExceptionResponse
}

func (o *CreateMsgVpnACLProfilePublishExceptionOK) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions][%d] createMsgVpnAclProfilePublishExceptionOK  %+v", 200, o.Payload)
}

func (o *CreateMsgVpnACLProfilePublishExceptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsgVpnACLProfilePublishExceptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMsgVpnACLProfilePublishExceptionDefault creates a CreateMsgVpnACLProfilePublishExceptionDefault with default headers values
func NewCreateMsgVpnACLProfilePublishExceptionDefault(code int) *CreateMsgVpnACLProfilePublishExceptionDefault {
	return &CreateMsgVpnACLProfilePublishExceptionDefault{
		_statusCode: code,
	}
}

/*CreateMsgVpnACLProfilePublishExceptionDefault handles this case with default header values.

The error response.
*/
type CreateMsgVpnACLProfilePublishExceptionDefault struct {
	_statusCode int

	Payload *models.SempMetaOnlyResponse
}

// Code gets the status code for the create msg vpn Acl profile publish exception default response
func (o *CreateMsgVpnACLProfilePublishExceptionDefault) Code() int {
	return o._statusCode
}

func (o *CreateMsgVpnACLProfilePublishExceptionDefault) Error() string {
	return fmt.Sprintf("[POST /msgVpns/{msgVpnName}/aclProfiles/{aclProfileName}/publishExceptions][%d] createMsgVpnAclProfilePublishException default  %+v", o._statusCode, o.Payload)
}

func (o *CreateMsgVpnACLProfilePublishExceptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SempMetaOnlyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
