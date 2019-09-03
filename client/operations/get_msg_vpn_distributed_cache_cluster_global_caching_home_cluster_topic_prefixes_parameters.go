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

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams creates a new GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams object
// with the default values initialized.
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams() *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams{
		Count: &countDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParamsWithTimeout creates a new GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParamsWithTimeout(timeout time.Duration) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams{
		Count: &countDefault,

		timeout: timeout,
	}
}

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParamsWithContext creates a new GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParamsWithContext(ctx context.Context) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams{
		Count: &countDefault,

		Context: ctx,
	}
}

// NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParamsWithHTTPClient creates a new GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParamsWithHTTPClient(client *http.Client) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams{
		Count:      &countDefault,
		HTTPClient: client,
	}
}

/*GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams contains all the parameters to send to the API endpoint
for the get msg vpn distributed cache cluster global caching home cluster topic prefixes operation typically these are written to a http.Request
*/
type GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams struct {

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
	/*HomeClusterName
	  The name of the remote Home Cache Cluster.

	*/
	HomeClusterName string
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

// WithTimeout adds the timeout to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) WithTimeout(timeout time.Duration) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) WithContext(ctx context.Context) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) WithHTTPClient(client *http.Client) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCacheName adds the cacheName to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) WithCacheName(cacheName string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithClusterName adds the clusterName to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) WithClusterName(clusterName string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithCount adds the count to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) WithCount(count *int64) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) SetCount(count *int64) {
	o.Count = count
}

// WithCursor adds the cursor to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) WithCursor(cursor *string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithHomeClusterName adds the homeClusterName to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) WithHomeClusterName(homeClusterName string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams {
	o.SetHomeClusterName(homeClusterName)
	return o
}

// SetHomeClusterName adds the homeClusterName to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) SetHomeClusterName(homeClusterName string) {
	o.HomeClusterName = homeClusterName
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) WithSelect(selectVar []string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithWhere adds the where to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) WithWhere(where []string) *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams {
	o.SetWhere(where)
	return o
}

// SetWhere adds the where to the get msg vpn distributed cache cluster global caching home cluster topic prefixes params
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) SetWhere(where []string) {
	o.Where = where
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnDistributedCacheClusterGlobalCachingHomeClusterTopicPrefixesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param homeClusterName
	if err := r.SetPathParam("homeClusterName", o.HomeClusterName); err != nil {
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