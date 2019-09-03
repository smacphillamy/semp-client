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

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteMsgVpnBridgeRemoteMsgVpnParams creates a new DeleteMsgVpnBridgeRemoteMsgVpnParams object
// with the default values initialized.
func NewDeleteMsgVpnBridgeRemoteMsgVpnParams() *DeleteMsgVpnBridgeRemoteMsgVpnParams {
	var ()
	return &DeleteMsgVpnBridgeRemoteMsgVpnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnBridgeRemoteMsgVpnParamsWithTimeout creates a new DeleteMsgVpnBridgeRemoteMsgVpnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMsgVpnBridgeRemoteMsgVpnParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnBridgeRemoteMsgVpnParams {
	var ()
	return &DeleteMsgVpnBridgeRemoteMsgVpnParams{

		timeout: timeout,
	}
}

// NewDeleteMsgVpnBridgeRemoteMsgVpnParamsWithContext creates a new DeleteMsgVpnBridgeRemoteMsgVpnParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMsgVpnBridgeRemoteMsgVpnParamsWithContext(ctx context.Context) *DeleteMsgVpnBridgeRemoteMsgVpnParams {
	var ()
	return &DeleteMsgVpnBridgeRemoteMsgVpnParams{

		Context: ctx,
	}
}

// NewDeleteMsgVpnBridgeRemoteMsgVpnParamsWithHTTPClient creates a new DeleteMsgVpnBridgeRemoteMsgVpnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMsgVpnBridgeRemoteMsgVpnParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnBridgeRemoteMsgVpnParams {
	var ()
	return &DeleteMsgVpnBridgeRemoteMsgVpnParams{
		HTTPClient: client,
	}
}

/*DeleteMsgVpnBridgeRemoteMsgVpnParams contains all the parameters to send to the API endpoint
for the delete msg vpn bridge remote msg vpn operation typically these are written to a http.Request
*/
type DeleteMsgVpnBridgeRemoteMsgVpnParams struct {

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
	/*RemoteMsgVpnInterface
	  The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, `remoteMsgVpnLocation` must not be a virtual router name.

	*/
	RemoteMsgVpnInterface string
	/*RemoteMsgVpnLocation
	  The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with "v:").

	*/
	RemoteMsgVpnLocation string
	/*RemoteMsgVpnName
	  The name of the remote Message VPN.

	*/
	RemoteMsgVpnName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnBridgeRemoteMsgVpnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) WithContext(ctx context.Context) *DeleteMsgVpnBridgeRemoteMsgVpnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnBridgeRemoteMsgVpnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBridgeName adds the bridgeName to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) WithBridgeName(bridgeName string) *DeleteMsgVpnBridgeRemoteMsgVpnParams {
	o.SetBridgeName(bridgeName)
	return o
}

// SetBridgeName adds the bridgeName to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) SetBridgeName(bridgeName string) {
	o.BridgeName = bridgeName
}

// WithBridgeVirtualRouter adds the bridgeVirtualRouter to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) WithBridgeVirtualRouter(bridgeVirtualRouter string) *DeleteMsgVpnBridgeRemoteMsgVpnParams {
	o.SetBridgeVirtualRouter(bridgeVirtualRouter)
	return o
}

// SetBridgeVirtualRouter adds the bridgeVirtualRouter to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) SetBridgeVirtualRouter(bridgeVirtualRouter string) {
	o.BridgeVirtualRouter = bridgeVirtualRouter
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnBridgeRemoteMsgVpnParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithRemoteMsgVpnInterface adds the remoteMsgVpnInterface to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) WithRemoteMsgVpnInterface(remoteMsgVpnInterface string) *DeleteMsgVpnBridgeRemoteMsgVpnParams {
	o.SetRemoteMsgVpnInterface(remoteMsgVpnInterface)
	return o
}

// SetRemoteMsgVpnInterface adds the remoteMsgVpnInterface to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) SetRemoteMsgVpnInterface(remoteMsgVpnInterface string) {
	o.RemoteMsgVpnInterface = remoteMsgVpnInterface
}

// WithRemoteMsgVpnLocation adds the remoteMsgVpnLocation to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) WithRemoteMsgVpnLocation(remoteMsgVpnLocation string) *DeleteMsgVpnBridgeRemoteMsgVpnParams {
	o.SetRemoteMsgVpnLocation(remoteMsgVpnLocation)
	return o
}

// SetRemoteMsgVpnLocation adds the remoteMsgVpnLocation to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) SetRemoteMsgVpnLocation(remoteMsgVpnLocation string) {
	o.RemoteMsgVpnLocation = remoteMsgVpnLocation
}

// WithRemoteMsgVpnName adds the remoteMsgVpnName to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) WithRemoteMsgVpnName(remoteMsgVpnName string) *DeleteMsgVpnBridgeRemoteMsgVpnParams {
	o.SetRemoteMsgVpnName(remoteMsgVpnName)
	return o
}

// SetRemoteMsgVpnName adds the remoteMsgVpnName to the delete msg vpn bridge remote msg vpn params
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) SetRemoteMsgVpnName(remoteMsgVpnName string) {
	o.RemoteMsgVpnName = remoteMsgVpnName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnBridgeRemoteMsgVpnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param remoteMsgVpnInterface
	if err := r.SetPathParam("remoteMsgVpnInterface", o.RemoteMsgVpnInterface); err != nil {
		return err
	}

	// path param remoteMsgVpnLocation
	if err := r.SetPathParam("remoteMsgVpnLocation", o.RemoteMsgVpnLocation); err != nil {
		return err
	}

	// path param remoteMsgVpnName
	if err := r.SetPathParam("remoteMsgVpnName", o.RemoteMsgVpnName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
