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

// NewGetMsgVpnBridgeTLSTrustedCommonNamesParams creates a new GetMsgVpnBridgeTLSTrustedCommonNamesParams object
// with the default values initialized.
func NewGetMsgVpnBridgeTLSTrustedCommonNamesParams() *GetMsgVpnBridgeTLSTrustedCommonNamesParams {
	var ()
	return &GetMsgVpnBridgeTLSTrustedCommonNamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnBridgeTLSTrustedCommonNamesParamsWithTimeout creates a new GetMsgVpnBridgeTLSTrustedCommonNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMsgVpnBridgeTLSTrustedCommonNamesParamsWithTimeout(timeout time.Duration) *GetMsgVpnBridgeTLSTrustedCommonNamesParams {
	var ()
	return &GetMsgVpnBridgeTLSTrustedCommonNamesParams{

		timeout: timeout,
	}
}

// NewGetMsgVpnBridgeTLSTrustedCommonNamesParamsWithContext creates a new GetMsgVpnBridgeTLSTrustedCommonNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMsgVpnBridgeTLSTrustedCommonNamesParamsWithContext(ctx context.Context) *GetMsgVpnBridgeTLSTrustedCommonNamesParams {
	var ()
	return &GetMsgVpnBridgeTLSTrustedCommonNamesParams{

		Context: ctx,
	}
}

// NewGetMsgVpnBridgeTLSTrustedCommonNamesParamsWithHTTPClient creates a new GetMsgVpnBridgeTLSTrustedCommonNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMsgVpnBridgeTLSTrustedCommonNamesParamsWithHTTPClient(client *http.Client) *GetMsgVpnBridgeTLSTrustedCommonNamesParams {
	var ()
	return &GetMsgVpnBridgeTLSTrustedCommonNamesParams{
		HTTPClient: client,
	}
}

/*GetMsgVpnBridgeTLSTrustedCommonNamesParams contains all the parameters to send to the API endpoint
for the get msg vpn bridge Tls trusted common names operation typically these are written to a http.Request
*/
type GetMsgVpnBridgeTLSTrustedCommonNamesParams struct {

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
	/*Where
	  Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter.

	*/
	Where []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) WithTimeout(timeout time.Duration) *GetMsgVpnBridgeTLSTrustedCommonNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) WithContext(ctx context.Context) *GetMsgVpnBridgeTLSTrustedCommonNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) WithHTTPClient(client *http.Client) *GetMsgVpnBridgeTLSTrustedCommonNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBridgeName adds the bridgeName to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) WithBridgeName(bridgeName string) *GetMsgVpnBridgeTLSTrustedCommonNamesParams {
	o.SetBridgeName(bridgeName)
	return o
}

// SetBridgeName adds the bridgeName to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) SetBridgeName(bridgeName string) {
	o.BridgeName = bridgeName
}

// WithBridgeVirtualRouter adds the bridgeVirtualRouter to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) WithBridgeVirtualRouter(bridgeVirtualRouter string) *GetMsgVpnBridgeTLSTrustedCommonNamesParams {
	o.SetBridgeVirtualRouter(bridgeVirtualRouter)
	return o
}

// SetBridgeVirtualRouter adds the bridgeVirtualRouter to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) SetBridgeVirtualRouter(bridgeVirtualRouter string) {
	o.BridgeVirtualRouter = bridgeVirtualRouter
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnBridgeTLSTrustedCommonNamesParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) WithSelect(selectVar []string) *GetMsgVpnBridgeTLSTrustedCommonNamesParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithWhere adds the where to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) WithWhere(where []string) *GetMsgVpnBridgeTLSTrustedCommonNamesParams {
	o.SetWhere(where)
	return o
}

// SetWhere adds the where to the get msg vpn bridge Tls trusted common names params
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) SetWhere(where []string) {
	o.Where = where
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnBridgeTLSTrustedCommonNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	valuesWhere := o.Where

	joinedWhere := swag.JoinByFormat(valuesWhere, "csv")
	// query array param where
	if err := r.SetQueryParam("where", joinedWhere...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
