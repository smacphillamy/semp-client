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

// NewUpdateMsgVpnDistributedCacheClusterParams creates a new UpdateMsgVpnDistributedCacheClusterParams object
// with the default values initialized.
func NewUpdateMsgVpnDistributedCacheClusterParams() *UpdateMsgVpnDistributedCacheClusterParams {
	var ()
	return &UpdateMsgVpnDistributedCacheClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMsgVpnDistributedCacheClusterParamsWithTimeout creates a new UpdateMsgVpnDistributedCacheClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateMsgVpnDistributedCacheClusterParamsWithTimeout(timeout time.Duration) *UpdateMsgVpnDistributedCacheClusterParams {
	var ()
	return &UpdateMsgVpnDistributedCacheClusterParams{

		timeout: timeout,
	}
}

// NewUpdateMsgVpnDistributedCacheClusterParamsWithContext creates a new UpdateMsgVpnDistributedCacheClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateMsgVpnDistributedCacheClusterParamsWithContext(ctx context.Context) *UpdateMsgVpnDistributedCacheClusterParams {
	var ()
	return &UpdateMsgVpnDistributedCacheClusterParams{

		Context: ctx,
	}
}

// NewUpdateMsgVpnDistributedCacheClusterParamsWithHTTPClient creates a new UpdateMsgVpnDistributedCacheClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateMsgVpnDistributedCacheClusterParamsWithHTTPClient(client *http.Client) *UpdateMsgVpnDistributedCacheClusterParams {
	var ()
	return &UpdateMsgVpnDistributedCacheClusterParams{
		HTTPClient: client,
	}
}

/*UpdateMsgVpnDistributedCacheClusterParams contains all the parameters to send to the API endpoint
for the update msg vpn distributed cache cluster operation typically these are written to a http.Request
*/
type UpdateMsgVpnDistributedCacheClusterParams struct {

	/*Body
	  The Cache Cluster object's attributes.

	*/
	Body *models.MsgVpnDistributedCacheCluster
	/*CacheName
	  The name of the Distributed Cache.

	*/
	CacheName string
	/*ClusterName
	  The name of the Cache Cluster.

	*/
	ClusterName string
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

// WithTimeout adds the timeout to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithTimeout(timeout time.Duration) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithContext(ctx context.Context) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithHTTPClient(client *http.Client) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithBody(body *models.MsgVpnDistributedCacheCluster) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetBody(body *models.MsgVpnDistributedCacheCluster) {
	o.Body = body
}

// WithCacheName adds the cacheName to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithCacheName(cacheName string) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithClusterName adds the clusterName to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithClusterName(clusterName string) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithMsgVpnName adds the msgVpnName to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithMsgVpnName(msgVpnName string) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) WithSelect(selectVar []string) *UpdateMsgVpnDistributedCacheClusterParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the update msg vpn distributed cache cluster params
func (o *UpdateMsgVpnDistributedCacheClusterParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMsgVpnDistributedCacheClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param clusterName
	if err := r.SetPathParam("clusterName", o.ClusterName); err != nil {
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
