// Code generated by go-swagger; DO NOT EDIT.

package about

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

// NewGetAboutUserMsgVpnParams creates a new GetAboutUserMsgVpnParams object
// with the default values initialized.
func NewGetAboutUserMsgVpnParams() *GetAboutUserMsgVpnParams {
	var ()
	return &GetAboutUserMsgVpnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAboutUserMsgVpnParamsWithTimeout creates a new GetAboutUserMsgVpnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAboutUserMsgVpnParamsWithTimeout(timeout time.Duration) *GetAboutUserMsgVpnParams {
	var ()
	return &GetAboutUserMsgVpnParams{

		timeout: timeout,
	}
}

// NewGetAboutUserMsgVpnParamsWithContext creates a new GetAboutUserMsgVpnParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAboutUserMsgVpnParamsWithContext(ctx context.Context) *GetAboutUserMsgVpnParams {
	var ()
	return &GetAboutUserMsgVpnParams{

		Context: ctx,
	}
}

// NewGetAboutUserMsgVpnParamsWithHTTPClient creates a new GetAboutUserMsgVpnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAboutUserMsgVpnParamsWithHTTPClient(client *http.Client) *GetAboutUserMsgVpnParams {
	var ()
	return &GetAboutUserMsgVpnParams{
		HTTPClient: client,
	}
}

/*GetAboutUserMsgVpnParams contains all the parameters to send to the API endpoint
for the get about user msg vpn operation typically these are written to a http.Request
*/
type GetAboutUserMsgVpnParams struct {

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

// WithTimeout adds the timeout to the get about user msg vpn params
func (o *GetAboutUserMsgVpnParams) WithTimeout(timeout time.Duration) *GetAboutUserMsgVpnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get about user msg vpn params
func (o *GetAboutUserMsgVpnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get about user msg vpn params
func (o *GetAboutUserMsgVpnParams) WithContext(ctx context.Context) *GetAboutUserMsgVpnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get about user msg vpn params
func (o *GetAboutUserMsgVpnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get about user msg vpn params
func (o *GetAboutUserMsgVpnParams) WithHTTPClient(client *http.Client) *GetAboutUserMsgVpnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get about user msg vpn params
func (o *GetAboutUserMsgVpnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMsgVpnName adds the msgVpnName to the get about user msg vpn params
func (o *GetAboutUserMsgVpnParams) WithMsgVpnName(msgVpnName string) *GetAboutUserMsgVpnParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get about user msg vpn params
func (o *GetAboutUserMsgVpnParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the get about user msg vpn params
func (o *GetAboutUserMsgVpnParams) WithSelect(selectVar []string) *GetAboutUserMsgVpnParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get about user msg vpn params
func (o *GetAboutUserMsgVpnParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetAboutUserMsgVpnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
