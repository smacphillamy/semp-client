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
)

// NewGetMsgVpnClientUsernameParams creates a new GetMsgVpnClientUsernameParams object
// with the default values initialized.
func NewGetMsgVpnClientUsernameParams() *GetMsgVpnClientUsernameParams {
	var ()
	return &GetMsgVpnClientUsernameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnClientUsernameParamsWithTimeout creates a new GetMsgVpnClientUsernameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMsgVpnClientUsernameParamsWithTimeout(timeout time.Duration) *GetMsgVpnClientUsernameParams {
	var ()
	return &GetMsgVpnClientUsernameParams{

		timeout: timeout,
	}
}

// NewGetMsgVpnClientUsernameParamsWithContext creates a new GetMsgVpnClientUsernameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMsgVpnClientUsernameParamsWithContext(ctx context.Context) *GetMsgVpnClientUsernameParams {
	var ()
	return &GetMsgVpnClientUsernameParams{

		Context: ctx,
	}
}

// NewGetMsgVpnClientUsernameParamsWithHTTPClient creates a new GetMsgVpnClientUsernameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMsgVpnClientUsernameParamsWithHTTPClient(client *http.Client) *GetMsgVpnClientUsernameParams {
	var ()
	return &GetMsgVpnClientUsernameParams{
		HTTPClient: client,
	}
}

/*GetMsgVpnClientUsernameParams contains all the parameters to send to the API endpoint
for the get msg vpn client username operation typically these are written to a http.Request
*/
type GetMsgVpnClientUsernameParams struct {

	/*ClientUsername
	  The clientUsername of the Client Username.

	*/
	ClientUsername string
	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See [Select](#select "Description of the syntax of the `select` parameter").

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get msg vpn client username params
func (o *GetMsgVpnClientUsernameParams) WithTimeout(timeout time.Duration) *GetMsgVpnClientUsernameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn client username params
func (o *GetMsgVpnClientUsernameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn client username params
func (o *GetMsgVpnClientUsernameParams) WithContext(ctx context.Context) *GetMsgVpnClientUsernameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn client username params
func (o *GetMsgVpnClientUsernameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn client username params
func (o *GetMsgVpnClientUsernameParams) WithHTTPClient(client *http.Client) *GetMsgVpnClientUsernameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn client username params
func (o *GetMsgVpnClientUsernameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientUsername adds the clientUsername to the get msg vpn client username params
func (o *GetMsgVpnClientUsernameParams) WithClientUsername(clientUsername string) *GetMsgVpnClientUsernameParams {
	o.SetClientUsername(clientUsername)
	return o
}

// SetClientUsername adds the clientUsername to the get msg vpn client username params
func (o *GetMsgVpnClientUsernameParams) SetClientUsername(clientUsername string) {
	o.ClientUsername = clientUsername
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn client username params
func (o *GetMsgVpnClientUsernameParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnClientUsernameParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn client username params
func (o *GetMsgVpnClientUsernameParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the get msg vpn client username params
func (o *GetMsgVpnClientUsernameParams) WithSelect(selectVar []string) *GetMsgVpnClientUsernameParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn client username params
func (o *GetMsgVpnClientUsernameParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnClientUsernameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clientUsername
	if err := r.SetPathParam("clientUsername", o.ClientUsername); err != nil {
		return err
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