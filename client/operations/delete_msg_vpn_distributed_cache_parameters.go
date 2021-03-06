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

// NewDeleteMsgVpnDistributedCacheParams creates a new DeleteMsgVpnDistributedCacheParams object
// with the default values initialized.
func NewDeleteMsgVpnDistributedCacheParams() *DeleteMsgVpnDistributedCacheParams {
	var ()
	return &DeleteMsgVpnDistributedCacheParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMsgVpnDistributedCacheParamsWithTimeout creates a new DeleteMsgVpnDistributedCacheParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMsgVpnDistributedCacheParamsWithTimeout(timeout time.Duration) *DeleteMsgVpnDistributedCacheParams {
	var ()
	return &DeleteMsgVpnDistributedCacheParams{

		timeout: timeout,
	}
}

// NewDeleteMsgVpnDistributedCacheParamsWithContext creates a new DeleteMsgVpnDistributedCacheParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMsgVpnDistributedCacheParamsWithContext(ctx context.Context) *DeleteMsgVpnDistributedCacheParams {
	var ()
	return &DeleteMsgVpnDistributedCacheParams{

		Context: ctx,
	}
}

// NewDeleteMsgVpnDistributedCacheParamsWithHTTPClient creates a new DeleteMsgVpnDistributedCacheParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMsgVpnDistributedCacheParamsWithHTTPClient(client *http.Client) *DeleteMsgVpnDistributedCacheParams {
	var ()
	return &DeleteMsgVpnDistributedCacheParams{
		HTTPClient: client,
	}
}

/*DeleteMsgVpnDistributedCacheParams contains all the parameters to send to the API endpoint
for the delete msg vpn distributed cache operation typically these are written to a http.Request
*/
type DeleteMsgVpnDistributedCacheParams struct {

	/*CacheName
	  The name of the Distributed Cache.

	*/
	CacheName string
	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete msg vpn distributed cache params
func (o *DeleteMsgVpnDistributedCacheParams) WithTimeout(timeout time.Duration) *DeleteMsgVpnDistributedCacheParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete msg vpn distributed cache params
func (o *DeleteMsgVpnDistributedCacheParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete msg vpn distributed cache params
func (o *DeleteMsgVpnDistributedCacheParams) WithContext(ctx context.Context) *DeleteMsgVpnDistributedCacheParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete msg vpn distributed cache params
func (o *DeleteMsgVpnDistributedCacheParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete msg vpn distributed cache params
func (o *DeleteMsgVpnDistributedCacheParams) WithHTTPClient(client *http.Client) *DeleteMsgVpnDistributedCacheParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete msg vpn distributed cache params
func (o *DeleteMsgVpnDistributedCacheParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCacheName adds the cacheName to the delete msg vpn distributed cache params
func (o *DeleteMsgVpnDistributedCacheParams) WithCacheName(cacheName string) *DeleteMsgVpnDistributedCacheParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the delete msg vpn distributed cache params
func (o *DeleteMsgVpnDistributedCacheParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithMsgVpnName adds the msgVpnName to the delete msg vpn distributed cache params
func (o *DeleteMsgVpnDistributedCacheParams) WithMsgVpnName(msgVpnName string) *DeleteMsgVpnDistributedCacheParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the delete msg vpn distributed cache params
func (o *DeleteMsgVpnDistributedCacheParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMsgVpnDistributedCacheParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cacheName
	if err := r.SetPathParam("cacheName", o.CacheName); err != nil {
		return err
	}

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
