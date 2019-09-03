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

	models "github.com/ExalDraen/semp-client/models"
)

// NewCreateMsgVpnACLProfileParams creates a new CreateMsgVpnACLProfileParams object
// with the default values initialized.
func NewCreateMsgVpnACLProfileParams() *CreateMsgVpnACLProfileParams {
	var ()
	return &CreateMsgVpnACLProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnACLProfileParamsWithTimeout creates a new CreateMsgVpnACLProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMsgVpnACLProfileParamsWithTimeout(timeout time.Duration) *CreateMsgVpnACLProfileParams {
	var ()
	return &CreateMsgVpnACLProfileParams{

		timeout: timeout,
	}
}

// NewCreateMsgVpnACLProfileParamsWithContext creates a new CreateMsgVpnACLProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMsgVpnACLProfileParamsWithContext(ctx context.Context) *CreateMsgVpnACLProfileParams {
	var ()
	return &CreateMsgVpnACLProfileParams{

		Context: ctx,
	}
}

// NewCreateMsgVpnACLProfileParamsWithHTTPClient creates a new CreateMsgVpnACLProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMsgVpnACLProfileParamsWithHTTPClient(client *http.Client) *CreateMsgVpnACLProfileParams {
	var ()
	return &CreateMsgVpnACLProfileParams{
		HTTPClient: client,
	}
}

/*CreateMsgVpnACLProfileParams contains all the parameters to send to the API endpoint
for the create msg vpn Acl profile operation typically these are written to a http.Request
*/
type CreateMsgVpnACLProfileParams struct {

	/*Body
	  The ACL Profile object's attributes.

	*/
	Body *models.MsgVpnACLProfile
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

// WithTimeout adds the timeout to the create msg vpn Acl profile params
func (o *CreateMsgVpnACLProfileParams) WithTimeout(timeout time.Duration) *CreateMsgVpnACLProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn Acl profile params
func (o *CreateMsgVpnACLProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn Acl profile params
func (o *CreateMsgVpnACLProfileParams) WithContext(ctx context.Context) *CreateMsgVpnACLProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn Acl profile params
func (o *CreateMsgVpnACLProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn Acl profile params
func (o *CreateMsgVpnACLProfileParams) WithHTTPClient(client *http.Client) *CreateMsgVpnACLProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn Acl profile params
func (o *CreateMsgVpnACLProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create msg vpn Acl profile params
func (o *CreateMsgVpnACLProfileParams) WithBody(body *models.MsgVpnACLProfile) *CreateMsgVpnACLProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn Acl profile params
func (o *CreateMsgVpnACLProfileParams) SetBody(body *models.MsgVpnACLProfile) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the create msg vpn Acl profile params
func (o *CreateMsgVpnACLProfileParams) WithMsgVpnName(msgVpnName string) *CreateMsgVpnACLProfileParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the create msg vpn Acl profile params
func (o *CreateMsgVpnACLProfileParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the create msg vpn Acl profile params
func (o *CreateMsgVpnACLProfileParams) WithSelect(selectVar []string) *CreateMsgVpnACLProfileParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn Acl profile params
func (o *CreateMsgVpnACLProfileParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnACLProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
