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

// NewReplaceMsgVpnDmrBridgeParams creates a new ReplaceMsgVpnDmrBridgeParams object
// with the default values initialized.
func NewReplaceMsgVpnDmrBridgeParams() *ReplaceMsgVpnDmrBridgeParams {
	var ()
	return &ReplaceMsgVpnDmrBridgeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMsgVpnDmrBridgeParamsWithTimeout creates a new ReplaceMsgVpnDmrBridgeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceMsgVpnDmrBridgeParamsWithTimeout(timeout time.Duration) *ReplaceMsgVpnDmrBridgeParams {
	var ()
	return &ReplaceMsgVpnDmrBridgeParams{

		timeout: timeout,
	}
}

// NewReplaceMsgVpnDmrBridgeParamsWithContext creates a new ReplaceMsgVpnDmrBridgeParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceMsgVpnDmrBridgeParamsWithContext(ctx context.Context) *ReplaceMsgVpnDmrBridgeParams {
	var ()
	return &ReplaceMsgVpnDmrBridgeParams{

		Context: ctx,
	}
}

// NewReplaceMsgVpnDmrBridgeParamsWithHTTPClient creates a new ReplaceMsgVpnDmrBridgeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceMsgVpnDmrBridgeParamsWithHTTPClient(client *http.Client) *ReplaceMsgVpnDmrBridgeParams {
	var ()
	return &ReplaceMsgVpnDmrBridgeParams{
		HTTPClient: client,
	}
}

/*ReplaceMsgVpnDmrBridgeParams contains all the parameters to send to the API endpoint
for the replace msg vpn dmr bridge operation typically these are written to a http.Request
*/
type ReplaceMsgVpnDmrBridgeParams struct {

	/*Body
	  The DMR Bridge object's attributes.

	*/
	Body *models.MsgVpnDmrBridge
	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string
	/*RemoteNodeName
	  The name of the node at the remote end of the DMR Bridge.

	*/
	RemoteNodeName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace msg vpn dmr bridge params
func (o *ReplaceMsgVpnDmrBridgeParams) WithTimeout(timeout time.Duration) *ReplaceMsgVpnDmrBridgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace msg vpn dmr bridge params
func (o *ReplaceMsgVpnDmrBridgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace msg vpn dmr bridge params
func (o *ReplaceMsgVpnDmrBridgeParams) WithContext(ctx context.Context) *ReplaceMsgVpnDmrBridgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace msg vpn dmr bridge params
func (o *ReplaceMsgVpnDmrBridgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace msg vpn dmr bridge params
func (o *ReplaceMsgVpnDmrBridgeParams) WithHTTPClient(client *http.Client) *ReplaceMsgVpnDmrBridgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace msg vpn dmr bridge params
func (o *ReplaceMsgVpnDmrBridgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace msg vpn dmr bridge params
func (o *ReplaceMsgVpnDmrBridgeParams) WithBody(body *models.MsgVpnDmrBridge) *ReplaceMsgVpnDmrBridgeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace msg vpn dmr bridge params
func (o *ReplaceMsgVpnDmrBridgeParams) SetBody(body *models.MsgVpnDmrBridge) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the replace msg vpn dmr bridge params
func (o *ReplaceMsgVpnDmrBridgeParams) WithMsgVpnName(msgVpnName string) *ReplaceMsgVpnDmrBridgeParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the replace msg vpn dmr bridge params
func (o *ReplaceMsgVpnDmrBridgeParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithRemoteNodeName adds the remoteNodeName to the replace msg vpn dmr bridge params
func (o *ReplaceMsgVpnDmrBridgeParams) WithRemoteNodeName(remoteNodeName string) *ReplaceMsgVpnDmrBridgeParams {
	o.SetRemoteNodeName(remoteNodeName)
	return o
}

// SetRemoteNodeName adds the remoteNodeName to the replace msg vpn dmr bridge params
func (o *ReplaceMsgVpnDmrBridgeParams) SetRemoteNodeName(remoteNodeName string) {
	o.RemoteNodeName = remoteNodeName
}

// WithSelect adds the selectVar to the replace msg vpn dmr bridge params
func (o *ReplaceMsgVpnDmrBridgeParams) WithSelect(selectVar []string) *ReplaceMsgVpnDmrBridgeParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace msg vpn dmr bridge params
func (o *ReplaceMsgVpnDmrBridgeParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceMsgVpnDmrBridgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param remoteNodeName
	if err := r.SetPathParam("remoteNodeName", o.RemoteNodeName); err != nil {
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
