// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/smacphillamy/semp-client/models"
)

// NewCreateMsgVpnACLProfileClientConnectExceptionParams creates a new CreateMsgVpnACLProfileClientConnectExceptionParams object
// with the default values initialized.
func NewCreateMsgVpnACLProfileClientConnectExceptionParams() *CreateMsgVpnACLProfileClientConnectExceptionParams {
	var ()
	return &CreateMsgVpnACLProfileClientConnectExceptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnACLProfileClientConnectExceptionParamsWithTimeout creates a new CreateMsgVpnACLProfileClientConnectExceptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMsgVpnACLProfileClientConnectExceptionParamsWithTimeout(timeout time.Duration) *CreateMsgVpnACLProfileClientConnectExceptionParams {
	var ()
	return &CreateMsgVpnACLProfileClientConnectExceptionParams{

		timeout: timeout,
	}
}

// NewCreateMsgVpnACLProfileClientConnectExceptionParamsWithContext creates a new CreateMsgVpnACLProfileClientConnectExceptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMsgVpnACLProfileClientConnectExceptionParamsWithContext(ctx context.Context) *CreateMsgVpnACLProfileClientConnectExceptionParams {
	var ()
	return &CreateMsgVpnACLProfileClientConnectExceptionParams{

		Context: ctx,
	}
}

// NewCreateMsgVpnACLProfileClientConnectExceptionParamsWithHTTPClient creates a new CreateMsgVpnACLProfileClientConnectExceptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMsgVpnACLProfileClientConnectExceptionParamsWithHTTPClient(client *http.Client) *CreateMsgVpnACLProfileClientConnectExceptionParams {
	var ()
	return &CreateMsgVpnACLProfileClientConnectExceptionParams{
		HTTPClient: client,
	}
}

/*CreateMsgVpnACLProfileClientConnectExceptionParams contains all the parameters to send to the API endpoint
for the create msg vpn Acl profile client connect exception operation typically these are written to a http.Request
*/
type CreateMsgVpnACLProfileClientConnectExceptionParams struct {

	/*ACLProfileName
	  The name of the ACL Profile.

	*/
	ACLProfileName string
	/*Body
	  The Client Connect Exception object's attributes.

	*/
	Body *models.MsgVpnACLProfileClientConnectException
	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create msg vpn Acl profile client connect exception params
func (o *CreateMsgVpnACLProfileClientConnectExceptionParams) WithTimeout(timeout time.Duration) *CreateMsgVpnACLProfileClientConnectExceptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn Acl profile client connect exception params
func (o *CreateMsgVpnACLProfileClientConnectExceptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn Acl profile client connect exception params
func (o *CreateMsgVpnACLProfileClientConnectExceptionParams) WithContext(ctx context.Context) *CreateMsgVpnACLProfileClientConnectExceptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn Acl profile client connect exception params
func (o *CreateMsgVpnACLProfileClientConnectExceptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn Acl profile client connect exception params
func (o *CreateMsgVpnACLProfileClientConnectExceptionParams) WithHTTPClient(client *http.Client) *CreateMsgVpnACLProfileClientConnectExceptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn Acl profile client connect exception params
func (o *CreateMsgVpnACLProfileClientConnectExceptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithACLProfileName adds the aCLProfileName to the create msg vpn Acl profile client connect exception params
func (o *CreateMsgVpnACLProfileClientConnectExceptionParams) WithACLProfileName(aCLProfileName string) *CreateMsgVpnACLProfileClientConnectExceptionParams {
	o.SetACLProfileName(aCLProfileName)
	return o
}

// SetACLProfileName adds the aclProfileName to the create msg vpn Acl profile client connect exception params
func (o *CreateMsgVpnACLProfileClientConnectExceptionParams) SetACLProfileName(aCLProfileName string) {
	o.ACLProfileName = aCLProfileName
}

// WithBody adds the body to the create msg vpn Acl profile client connect exception params
func (o *CreateMsgVpnACLProfileClientConnectExceptionParams) WithBody(body *models.MsgVpnACLProfileClientConnectException) *CreateMsgVpnACLProfileClientConnectExceptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn Acl profile client connect exception params
func (o *CreateMsgVpnACLProfileClientConnectExceptionParams) SetBody(body *models.MsgVpnACLProfileClientConnectException) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the create msg vpn Acl profile client connect exception params
func (o *CreateMsgVpnACLProfileClientConnectExceptionParams) WithMsgVpnName(msgVpnName string) *CreateMsgVpnACLProfileClientConnectExceptionParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the create msg vpn Acl profile client connect exception params
func (o *CreateMsgVpnACLProfileClientConnectExceptionParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the create msg vpn Acl profile client connect exception params
func (o *CreateMsgVpnACLProfileClientConnectExceptionParams) WithSelect(selectVar []string) *CreateMsgVpnACLProfileClientConnectExceptionParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn Acl profile client connect exception params
func (o *CreateMsgVpnACLProfileClientConnectExceptionParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnACLProfileClientConnectExceptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param aclProfileName
	if err := r.SetPathParam("aclProfileName", o.ACLProfileName); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	valuesSelect := o.Select

	joinedSelect := swag.JoinByFormat(valuesSelect, "csv")
	// query array param select
	if err := r.SetQueryParam("select", joinedSelect...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
