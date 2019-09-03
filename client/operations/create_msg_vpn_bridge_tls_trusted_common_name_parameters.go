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

// NewCreateMsgVpnBridgeTLSTrustedCommonNameParams creates a new CreateMsgVpnBridgeTLSTrustedCommonNameParams object
// with the default values initialized.
func NewCreateMsgVpnBridgeTLSTrustedCommonNameParams() *CreateMsgVpnBridgeTLSTrustedCommonNameParams {
	var ()
	return &CreateMsgVpnBridgeTLSTrustedCommonNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnBridgeTLSTrustedCommonNameParamsWithTimeout creates a new CreateMsgVpnBridgeTLSTrustedCommonNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMsgVpnBridgeTLSTrustedCommonNameParamsWithTimeout(timeout time.Duration) *CreateMsgVpnBridgeTLSTrustedCommonNameParams {
	var ()
	return &CreateMsgVpnBridgeTLSTrustedCommonNameParams{

		timeout: timeout,
	}
}

// NewCreateMsgVpnBridgeTLSTrustedCommonNameParamsWithContext creates a new CreateMsgVpnBridgeTLSTrustedCommonNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMsgVpnBridgeTLSTrustedCommonNameParamsWithContext(ctx context.Context) *CreateMsgVpnBridgeTLSTrustedCommonNameParams {
	var ()
	return &CreateMsgVpnBridgeTLSTrustedCommonNameParams{

		Context: ctx,
	}
}

// NewCreateMsgVpnBridgeTLSTrustedCommonNameParamsWithHTTPClient creates a new CreateMsgVpnBridgeTLSTrustedCommonNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMsgVpnBridgeTLSTrustedCommonNameParamsWithHTTPClient(client *http.Client) *CreateMsgVpnBridgeTLSTrustedCommonNameParams {
	var ()
	return &CreateMsgVpnBridgeTLSTrustedCommonNameParams{
		HTTPClient: client,
	}
}

/*CreateMsgVpnBridgeTLSTrustedCommonNameParams contains all the parameters to send to the API endpoint
for the create msg vpn bridge Tls trusted common name operation typically these are written to a http.Request
*/
type CreateMsgVpnBridgeTLSTrustedCommonNameParams struct {

	/*Body
	  The Trusted Common Name object's attributes.

	*/
	Body *models.MsgVpnBridgeTLSTrustedCommonName
	/*BridgeName
	  The name of the Bridge.

	*/
	BridgeName string
	/*BridgeVirtualRouter
	  The virtual router of the Bridge.

	*/
	BridgeVirtualRouter string
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

// WithTimeout adds the timeout to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) WithTimeout(timeout time.Duration) *CreateMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) WithContext(ctx context.Context) *CreateMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) WithHTTPClient(client *http.Client) *CreateMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) WithBody(body *models.MsgVpnBridgeTLSTrustedCommonName) *CreateMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) SetBody(body *models.MsgVpnBridgeTLSTrustedCommonName) {
	o.Body = body
}

// WithBridgeName adds the bridgeName to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) WithBridgeName(bridgeName string) *CreateMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetBridgeName(bridgeName)
	return o
}

// SetBridgeName adds the bridgeName to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) SetBridgeName(bridgeName string) {
	o.BridgeName = bridgeName
}

// WithBridgeVirtualRouter adds the bridgeVirtualRouter to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) WithBridgeVirtualRouter(bridgeVirtualRouter string) *CreateMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetBridgeVirtualRouter(bridgeVirtualRouter)
	return o
}

// SetBridgeVirtualRouter adds the bridgeVirtualRouter to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) SetBridgeVirtualRouter(bridgeVirtualRouter string) {
	o.BridgeVirtualRouter = bridgeVirtualRouter
}

// WithMsgVpnName adds the msgVpnName to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) WithMsgVpnName(msgVpnName string) *CreateMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) WithSelect(selectVar []string) *CreateMsgVpnBridgeTLSTrustedCommonNameParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn bridge Tls trusted common name params
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnBridgeTLSTrustedCommonNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param bridgeName
	if err := r.SetPathParam("bridgeName", o.BridgeName); err != nil {
		return err
	}

	// path param bridgeVirtualRouter
	if err := r.SetPathParam("bridgeVirtualRouter", o.BridgeVirtualRouter); err != nil {
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
