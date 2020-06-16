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

// NewCreateMsgVpnDmrBridgeParams creates a new CreateMsgVpnDmrBridgeParams object
// with the default values initialized.
func NewCreateMsgVpnDmrBridgeParams() *CreateMsgVpnDmrBridgeParams {
	var ()
	return &CreateMsgVpnDmrBridgeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnDmrBridgeParamsWithTimeout creates a new CreateMsgVpnDmrBridgeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMsgVpnDmrBridgeParamsWithTimeout(timeout time.Duration) *CreateMsgVpnDmrBridgeParams {
	var ()
	return &CreateMsgVpnDmrBridgeParams{

		timeout: timeout,
	}
}

// NewCreateMsgVpnDmrBridgeParamsWithContext creates a new CreateMsgVpnDmrBridgeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMsgVpnDmrBridgeParamsWithContext(ctx context.Context) *CreateMsgVpnDmrBridgeParams {
	var ()
	return &CreateMsgVpnDmrBridgeParams{

		Context: ctx,
	}
}

// NewCreateMsgVpnDmrBridgeParamsWithHTTPClient creates a new CreateMsgVpnDmrBridgeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMsgVpnDmrBridgeParamsWithHTTPClient(client *http.Client) *CreateMsgVpnDmrBridgeParams {
	var ()
	return &CreateMsgVpnDmrBridgeParams{
		HTTPClient: client,
	}
}

/*CreateMsgVpnDmrBridgeParams contains all the parameters to send to the API endpoint
for the create msg vpn dmr bridge operation typically these are written to a http.Request
*/
type CreateMsgVpnDmrBridgeParams struct {

	/*Body
	  The DMR Bridge object's attributes.

	*/
	Body *models.MsgVpnDmrBridge
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

// WithTimeout adds the timeout to the create msg vpn dmr bridge params
func (o *CreateMsgVpnDmrBridgeParams) WithTimeout(timeout time.Duration) *CreateMsgVpnDmrBridgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn dmr bridge params
func (o *CreateMsgVpnDmrBridgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn dmr bridge params
func (o *CreateMsgVpnDmrBridgeParams) WithContext(ctx context.Context) *CreateMsgVpnDmrBridgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn dmr bridge params
func (o *CreateMsgVpnDmrBridgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn dmr bridge params
func (o *CreateMsgVpnDmrBridgeParams) WithHTTPClient(client *http.Client) *CreateMsgVpnDmrBridgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn dmr bridge params
func (o *CreateMsgVpnDmrBridgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create msg vpn dmr bridge params
func (o *CreateMsgVpnDmrBridgeParams) WithBody(body *models.MsgVpnDmrBridge) *CreateMsgVpnDmrBridgeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn dmr bridge params
func (o *CreateMsgVpnDmrBridgeParams) SetBody(body *models.MsgVpnDmrBridge) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the create msg vpn dmr bridge params
func (o *CreateMsgVpnDmrBridgeParams) WithMsgVpnName(msgVpnName string) *CreateMsgVpnDmrBridgeParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the create msg vpn dmr bridge params
func (o *CreateMsgVpnDmrBridgeParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the create msg vpn dmr bridge params
func (o *CreateMsgVpnDmrBridgeParams) WithSelect(selectVar []string) *CreateMsgVpnDmrBridgeParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn dmr bridge params
func (o *CreateMsgVpnDmrBridgeParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnDmrBridgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
