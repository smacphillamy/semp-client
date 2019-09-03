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

// NewGetMsgVpnDistributedCacheClusterInstancesParams creates a new GetMsgVpnDistributedCacheClusterInstancesParams object
// with the default values initialized.
func NewGetMsgVpnDistributedCacheClusterInstancesParams() *GetMsgVpnDistributedCacheClusterInstancesParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnDistributedCacheClusterInstancesParams{
		Count: &countDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnDistributedCacheClusterInstancesParamsWithTimeout creates a new GetMsgVpnDistributedCacheClusterInstancesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMsgVpnDistributedCacheClusterInstancesParamsWithTimeout(timeout time.Duration) *GetMsgVpnDistributedCacheClusterInstancesParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnDistributedCacheClusterInstancesParams{
		Count: &countDefault,

		timeout: timeout,
	}
}

// NewGetMsgVpnDistributedCacheClusterInstancesParamsWithContext creates a new GetMsgVpnDistributedCacheClusterInstancesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMsgVpnDistributedCacheClusterInstancesParamsWithContext(ctx context.Context) *GetMsgVpnDistributedCacheClusterInstancesParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnDistributedCacheClusterInstancesParams{
		Count: &countDefault,

		Context: ctx,
	}
}

// NewGetMsgVpnDistributedCacheClusterInstancesParamsWithHTTPClient creates a new GetMsgVpnDistributedCacheClusterInstancesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMsgVpnDistributedCacheClusterInstancesParamsWithHTTPClient(client *http.Client) *GetMsgVpnDistributedCacheClusterInstancesParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnDistributedCacheClusterInstancesParams{
		Count:      &countDefault,
		HTTPClient: client,
	}
}

/*GetMsgVpnDistributedCacheClusterInstancesParams contains all the parameters to send to the API endpoint
for the get msg vpn distributed cache cluster instances operation typically these are written to a http.Request
*/
type GetMsgVpnDistributedCacheClusterInstancesParams struct {

	/*CacheName
	  The name of the Distributed Cache.

	*/
	CacheName string
	/*ClusterName
	  The name of the Cache Cluster.

	*/
	ClusterName string
	/*Count
	  Limit the count of objects in the response. See the documentation for the `count` parameter.

	*/
	Count *int64
	/*Cursor
	  The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter.

	*/
	Cursor *string
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

// WithTimeout adds the timeout to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) WithTimeout(timeout time.Duration) *GetMsgVpnDistributedCacheClusterInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) WithContext(ctx context.Context) *GetMsgVpnDistributedCacheClusterInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) WithHTTPClient(client *http.Client) *GetMsgVpnDistributedCacheClusterInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCacheName adds the cacheName to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) WithCacheName(cacheName string) *GetMsgVpnDistributedCacheClusterInstancesParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithClusterName adds the clusterName to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) WithClusterName(clusterName string) *GetMsgVpnDistributedCacheClusterInstancesParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithCount adds the count to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) WithCount(count *int64) *GetMsgVpnDistributedCacheClusterInstancesParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) SetCount(count *int64) {
	o.Count = count
}

// WithCursor adds the cursor to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) WithCursor(cursor *string) *GetMsgVpnDistributedCacheClusterInstancesParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnDistributedCacheClusterInstancesParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) WithSelect(selectVar []string) *GetMsgVpnDistributedCacheClusterInstancesParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithWhere adds the where to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) WithWhere(where []string) *GetMsgVpnDistributedCacheClusterInstancesParams {
	o.SetWhere(where)
	return o
}

// SetWhere adds the where to the get msg vpn distributed cache cluster instances params
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) SetWhere(where []string) {
	o.Where = where
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnDistributedCacheClusterInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cacheName
	if err := r.SetPathParam("cacheName", o.CacheName); err != nil {
		return err
	}

	// path param clusterName
	if err := r.SetPathParam("clusterName", o.ClusterName); err != nil {
		return err
	}

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string
		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {
			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}

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
