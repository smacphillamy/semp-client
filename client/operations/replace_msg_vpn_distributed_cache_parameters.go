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

// NewReplaceMsgVpnDistributedCacheParams creates a new ReplaceMsgVpnDistributedCacheParams object
// with the default values initialized.
func NewReplaceMsgVpnDistributedCacheParams() *ReplaceMsgVpnDistributedCacheParams {
	var ()
	return &ReplaceMsgVpnDistributedCacheParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMsgVpnDistributedCacheParamsWithTimeout creates a new ReplaceMsgVpnDistributedCacheParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceMsgVpnDistributedCacheParamsWithTimeout(timeout time.Duration) *ReplaceMsgVpnDistributedCacheParams {
	var ()
	return &ReplaceMsgVpnDistributedCacheParams{

		timeout: timeout,
	}
}

// NewReplaceMsgVpnDistributedCacheParamsWithContext creates a new ReplaceMsgVpnDistributedCacheParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceMsgVpnDistributedCacheParamsWithContext(ctx context.Context) *ReplaceMsgVpnDistributedCacheParams {
	var ()
	return &ReplaceMsgVpnDistributedCacheParams{

		Context: ctx,
	}
}

// NewReplaceMsgVpnDistributedCacheParamsWithHTTPClient creates a new ReplaceMsgVpnDistributedCacheParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceMsgVpnDistributedCacheParamsWithHTTPClient(client *http.Client) *ReplaceMsgVpnDistributedCacheParams {
	var ()
	return &ReplaceMsgVpnDistributedCacheParams{
		HTTPClient: client,
	}
}

/*ReplaceMsgVpnDistributedCacheParams contains all the parameters to send to the API endpoint
for the replace msg vpn distributed cache operation typically these are written to a http.Request
*/
type ReplaceMsgVpnDistributedCacheParams struct {

	/*Body
	  The Distributed Cache object's attributes.

	*/
	Body *models.MsgVpnDistributedCache
	/*CacheName
	  The name of the Distributed Cache.

	*/
	CacheName string
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

// WithTimeout adds the timeout to the replace msg vpn distributed cache params
func (o *ReplaceMsgVpnDistributedCacheParams) WithTimeout(timeout time.Duration) *ReplaceMsgVpnDistributedCacheParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace msg vpn distributed cache params
func (o *ReplaceMsgVpnDistributedCacheParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace msg vpn distributed cache params
func (o *ReplaceMsgVpnDistributedCacheParams) WithContext(ctx context.Context) *ReplaceMsgVpnDistributedCacheParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace msg vpn distributed cache params
func (o *ReplaceMsgVpnDistributedCacheParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace msg vpn distributed cache params
func (o *ReplaceMsgVpnDistributedCacheParams) WithHTTPClient(client *http.Client) *ReplaceMsgVpnDistributedCacheParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace msg vpn distributed cache params
func (o *ReplaceMsgVpnDistributedCacheParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace msg vpn distributed cache params
func (o *ReplaceMsgVpnDistributedCacheParams) WithBody(body *models.MsgVpnDistributedCache) *ReplaceMsgVpnDistributedCacheParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace msg vpn distributed cache params
func (o *ReplaceMsgVpnDistributedCacheParams) SetBody(body *models.MsgVpnDistributedCache) {
	o.Body = body
}

// WithCacheName adds the cacheName to the replace msg vpn distributed cache params
func (o *ReplaceMsgVpnDistributedCacheParams) WithCacheName(cacheName string) *ReplaceMsgVpnDistributedCacheParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the replace msg vpn distributed cache params
func (o *ReplaceMsgVpnDistributedCacheParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithMsgVpnName adds the msgVpnName to the replace msg vpn distributed cache params
func (o *ReplaceMsgVpnDistributedCacheParams) WithMsgVpnName(msgVpnName string) *ReplaceMsgVpnDistributedCacheParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the replace msg vpn distributed cache params
func (o *ReplaceMsgVpnDistributedCacheParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the replace msg vpn distributed cache params
func (o *ReplaceMsgVpnDistributedCacheParams) WithSelect(selectVar []string) *ReplaceMsgVpnDistributedCacheParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace msg vpn distributed cache params
func (o *ReplaceMsgVpnDistributedCacheParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceMsgVpnDistributedCacheParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cacheName
	if err := r.SetPathParam("cacheName", o.CacheName); err != nil {
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
