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

// NewCreateMsgVpnDistributedCacheClusterParams creates a new CreateMsgVpnDistributedCacheClusterParams object
// with the default values initialized.
func NewCreateMsgVpnDistributedCacheClusterParams() *CreateMsgVpnDistributedCacheClusterParams {
	var ()
	return &CreateMsgVpnDistributedCacheClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnDistributedCacheClusterParamsWithTimeout creates a new CreateMsgVpnDistributedCacheClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMsgVpnDistributedCacheClusterParamsWithTimeout(timeout time.Duration) *CreateMsgVpnDistributedCacheClusterParams {
	var ()
	return &CreateMsgVpnDistributedCacheClusterParams{

		timeout: timeout,
	}
}

// NewCreateMsgVpnDistributedCacheClusterParamsWithContext creates a new CreateMsgVpnDistributedCacheClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMsgVpnDistributedCacheClusterParamsWithContext(ctx context.Context) *CreateMsgVpnDistributedCacheClusterParams {
	var ()
	return &CreateMsgVpnDistributedCacheClusterParams{

		Context: ctx,
	}
}

// NewCreateMsgVpnDistributedCacheClusterParamsWithHTTPClient creates a new CreateMsgVpnDistributedCacheClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMsgVpnDistributedCacheClusterParamsWithHTTPClient(client *http.Client) *CreateMsgVpnDistributedCacheClusterParams {
	var ()
	return &CreateMsgVpnDistributedCacheClusterParams{
		HTTPClient: client,
	}
}

/*CreateMsgVpnDistributedCacheClusterParams contains all the parameters to send to the API endpoint
for the create msg vpn distributed cache cluster operation typically these are written to a http.Request
*/
type CreateMsgVpnDistributedCacheClusterParams struct {

	/*Body
	  The Cache Cluster object's attributes.

	*/
	Body *models.MsgVpnDistributedCacheCluster
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

// WithTimeout adds the timeout to the create msg vpn distributed cache cluster params
func (o *CreateMsgVpnDistributedCacheClusterParams) WithTimeout(timeout time.Duration) *CreateMsgVpnDistributedCacheClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn distributed cache cluster params
func (o *CreateMsgVpnDistributedCacheClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn distributed cache cluster params
func (o *CreateMsgVpnDistributedCacheClusterParams) WithContext(ctx context.Context) *CreateMsgVpnDistributedCacheClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn distributed cache cluster params
func (o *CreateMsgVpnDistributedCacheClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn distributed cache cluster params
func (o *CreateMsgVpnDistributedCacheClusterParams) WithHTTPClient(client *http.Client) *CreateMsgVpnDistributedCacheClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn distributed cache cluster params
func (o *CreateMsgVpnDistributedCacheClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create msg vpn distributed cache cluster params
func (o *CreateMsgVpnDistributedCacheClusterParams) WithBody(body *models.MsgVpnDistributedCacheCluster) *CreateMsgVpnDistributedCacheClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn distributed cache cluster params
func (o *CreateMsgVpnDistributedCacheClusterParams) SetBody(body *models.MsgVpnDistributedCacheCluster) {
	o.Body = body
}

// WithCacheName adds the cacheName to the create msg vpn distributed cache cluster params
func (o *CreateMsgVpnDistributedCacheClusterParams) WithCacheName(cacheName string) *CreateMsgVpnDistributedCacheClusterParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the create msg vpn distributed cache cluster params
func (o *CreateMsgVpnDistributedCacheClusterParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithMsgVpnName adds the msgVpnName to the create msg vpn distributed cache cluster params
func (o *CreateMsgVpnDistributedCacheClusterParams) WithMsgVpnName(msgVpnName string) *CreateMsgVpnDistributedCacheClusterParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the create msg vpn distributed cache cluster params
func (o *CreateMsgVpnDistributedCacheClusterParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the create msg vpn distributed cache cluster params
func (o *CreateMsgVpnDistributedCacheClusterParams) WithSelect(selectVar []string) *CreateMsgVpnDistributedCacheClusterParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn distributed cache cluster params
func (o *CreateMsgVpnDistributedCacheClusterParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnDistributedCacheClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
