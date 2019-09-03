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

// NewGetMsgVpnJndiConnectionFactoryParams creates a new GetMsgVpnJndiConnectionFactoryParams object
// with the default values initialized.
func NewGetMsgVpnJndiConnectionFactoryParams() *GetMsgVpnJndiConnectionFactoryParams {
	var ()
	return &GetMsgVpnJndiConnectionFactoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnJndiConnectionFactoryParamsWithTimeout creates a new GetMsgVpnJndiConnectionFactoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMsgVpnJndiConnectionFactoryParamsWithTimeout(timeout time.Duration) *GetMsgVpnJndiConnectionFactoryParams {
	var ()
	return &GetMsgVpnJndiConnectionFactoryParams{

		timeout: timeout,
	}
}

// NewGetMsgVpnJndiConnectionFactoryParamsWithContext creates a new GetMsgVpnJndiConnectionFactoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMsgVpnJndiConnectionFactoryParamsWithContext(ctx context.Context) *GetMsgVpnJndiConnectionFactoryParams {
	var ()
	return &GetMsgVpnJndiConnectionFactoryParams{

		Context: ctx,
	}
}

// NewGetMsgVpnJndiConnectionFactoryParamsWithHTTPClient creates a new GetMsgVpnJndiConnectionFactoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMsgVpnJndiConnectionFactoryParamsWithHTTPClient(client *http.Client) *GetMsgVpnJndiConnectionFactoryParams {
	var ()
	return &GetMsgVpnJndiConnectionFactoryParams{
		HTTPClient: client,
	}
}

/*GetMsgVpnJndiConnectionFactoryParams contains all the parameters to send to the API endpoint
for the get msg vpn jndi connection factory operation typically these are written to a http.Request
*/
type GetMsgVpnJndiConnectionFactoryParams struct {

	/*ConnectionFactoryName
	  The name of the JMS Connection Factory.

	*/
	ConnectionFactoryName string
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

// WithTimeout adds the timeout to the get msg vpn jndi connection factory params
func (o *GetMsgVpnJndiConnectionFactoryParams) WithTimeout(timeout time.Duration) *GetMsgVpnJndiConnectionFactoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn jndi connection factory params
func (o *GetMsgVpnJndiConnectionFactoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn jndi connection factory params
func (o *GetMsgVpnJndiConnectionFactoryParams) WithContext(ctx context.Context) *GetMsgVpnJndiConnectionFactoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn jndi connection factory params
func (o *GetMsgVpnJndiConnectionFactoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn jndi connection factory params
func (o *GetMsgVpnJndiConnectionFactoryParams) WithHTTPClient(client *http.Client) *GetMsgVpnJndiConnectionFactoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn jndi connection factory params
func (o *GetMsgVpnJndiConnectionFactoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionFactoryName adds the connectionFactoryName to the get msg vpn jndi connection factory params
func (o *GetMsgVpnJndiConnectionFactoryParams) WithConnectionFactoryName(connectionFactoryName string) *GetMsgVpnJndiConnectionFactoryParams {
	o.SetConnectionFactoryName(connectionFactoryName)
	return o
}

// SetConnectionFactoryName adds the connectionFactoryName to the get msg vpn jndi connection factory params
func (o *GetMsgVpnJndiConnectionFactoryParams) SetConnectionFactoryName(connectionFactoryName string) {
	o.ConnectionFactoryName = connectionFactoryName
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn jndi connection factory params
func (o *GetMsgVpnJndiConnectionFactoryParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnJndiConnectionFactoryParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn jndi connection factory params
func (o *GetMsgVpnJndiConnectionFactoryParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the get msg vpn jndi connection factory params
func (o *GetMsgVpnJndiConnectionFactoryParams) WithSelect(selectVar []string) *GetMsgVpnJndiConnectionFactoryParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn jndi connection factory params
func (o *GetMsgVpnJndiConnectionFactoryParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnJndiConnectionFactoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connectionFactoryName
	if err := r.SetPathParam("connectionFactoryName", o.ConnectionFactoryName); err != nil {
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
