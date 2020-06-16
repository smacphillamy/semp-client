// Code generated by go-swagger; DO NOT EDIT.

package msg_vpn

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

// NewCreateMsgVpnParams creates a new CreateMsgVpnParams object
// with the default values initialized.
func NewCreateMsgVpnParams() *CreateMsgVpnParams {
	var ()
	return &CreateMsgVpnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnParamsWithTimeout creates a new CreateMsgVpnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMsgVpnParamsWithTimeout(timeout time.Duration) *CreateMsgVpnParams {
	var ()
	return &CreateMsgVpnParams{

		timeout: timeout,
	}
}

// NewCreateMsgVpnParamsWithContext creates a new CreateMsgVpnParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMsgVpnParamsWithContext(ctx context.Context) *CreateMsgVpnParams {
	var ()
	return &CreateMsgVpnParams{

		Context: ctx,
	}
}

// NewCreateMsgVpnParamsWithHTTPClient creates a new CreateMsgVpnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMsgVpnParamsWithHTTPClient(client *http.Client) *CreateMsgVpnParams {
	var ()
	return &CreateMsgVpnParams{
		HTTPClient: client,
	}
}

/*CreateMsgVpnParams contains all the parameters to send to the API endpoint
for the create msg vpn operation typically these are written to a http.Request
*/
type CreateMsgVpnParams struct {

	/*Body
	  The Message VPN object's attributes.

	*/
	Body *models.MsgVpn
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create msg vpn params
func (o *CreateMsgVpnParams) WithTimeout(timeout time.Duration) *CreateMsgVpnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn params
func (o *CreateMsgVpnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn params
func (o *CreateMsgVpnParams) WithContext(ctx context.Context) *CreateMsgVpnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn params
func (o *CreateMsgVpnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn params
func (o *CreateMsgVpnParams) WithHTTPClient(client *http.Client) *CreateMsgVpnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn params
func (o *CreateMsgVpnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create msg vpn params
func (o *CreateMsgVpnParams) WithBody(body *models.MsgVpn) *CreateMsgVpnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn params
func (o *CreateMsgVpnParams) SetBody(body *models.MsgVpn) {
	o.Body = body
}

// WithSelect adds the selectVar to the create msg vpn params
func (o *CreateMsgVpnParams) WithSelect(selectVar []string) *CreateMsgVpnParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn params
func (o *CreateMsgVpnParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
