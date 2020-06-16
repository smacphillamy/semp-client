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

// NewCreateMsgVpnMqttSessionParams creates a new CreateMsgVpnMqttSessionParams object
// with the default values initialized.
func NewCreateMsgVpnMqttSessionParams() *CreateMsgVpnMqttSessionParams {
	var ()
	return &CreateMsgVpnMqttSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnMqttSessionParamsWithTimeout creates a new CreateMsgVpnMqttSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMsgVpnMqttSessionParamsWithTimeout(timeout time.Duration) *CreateMsgVpnMqttSessionParams {
	var ()
	return &CreateMsgVpnMqttSessionParams{

		timeout: timeout,
	}
}

// NewCreateMsgVpnMqttSessionParamsWithContext creates a new CreateMsgVpnMqttSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMsgVpnMqttSessionParamsWithContext(ctx context.Context) *CreateMsgVpnMqttSessionParams {
	var ()
	return &CreateMsgVpnMqttSessionParams{

		Context: ctx,
	}
}

// NewCreateMsgVpnMqttSessionParamsWithHTTPClient creates a new CreateMsgVpnMqttSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMsgVpnMqttSessionParamsWithHTTPClient(client *http.Client) *CreateMsgVpnMqttSessionParams {
	var ()
	return &CreateMsgVpnMqttSessionParams{
		HTTPClient: client,
	}
}

/*CreateMsgVpnMqttSessionParams contains all the parameters to send to the API endpoint
for the create msg vpn mqtt session operation typically these are written to a http.Request
*/
type CreateMsgVpnMqttSessionParams struct {

	/*Body
	  The MQTT Session object's attributes.

	*/
	Body *models.MsgVpnMqttSession
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

// WithTimeout adds the timeout to the create msg vpn mqtt session params
func (o *CreateMsgVpnMqttSessionParams) WithTimeout(timeout time.Duration) *CreateMsgVpnMqttSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn mqtt session params
func (o *CreateMsgVpnMqttSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn mqtt session params
func (o *CreateMsgVpnMqttSessionParams) WithContext(ctx context.Context) *CreateMsgVpnMqttSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn mqtt session params
func (o *CreateMsgVpnMqttSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn mqtt session params
func (o *CreateMsgVpnMqttSessionParams) WithHTTPClient(client *http.Client) *CreateMsgVpnMqttSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn mqtt session params
func (o *CreateMsgVpnMqttSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create msg vpn mqtt session params
func (o *CreateMsgVpnMqttSessionParams) WithBody(body *models.MsgVpnMqttSession) *CreateMsgVpnMqttSessionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn mqtt session params
func (o *CreateMsgVpnMqttSessionParams) SetBody(body *models.MsgVpnMqttSession) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the create msg vpn mqtt session params
func (o *CreateMsgVpnMqttSessionParams) WithMsgVpnName(msgVpnName string) *CreateMsgVpnMqttSessionParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the create msg vpn mqtt session params
func (o *CreateMsgVpnMqttSessionParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the create msg vpn mqtt session params
func (o *CreateMsgVpnMqttSessionParams) WithSelect(selectVar []string) *CreateMsgVpnMqttSessionParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn mqtt session params
func (o *CreateMsgVpnMqttSessionParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnMqttSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
