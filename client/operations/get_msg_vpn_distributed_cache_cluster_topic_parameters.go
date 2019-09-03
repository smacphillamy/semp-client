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

// NewGetMsgVpnDistributedCacheClusterTopicParams creates a new GetMsgVpnDistributedCacheClusterTopicParams object
// with the default values initialized.
func NewGetMsgVpnDistributedCacheClusterTopicParams() *GetMsgVpnDistributedCacheClusterTopicParams {
	var ()
	return &GetMsgVpnDistributedCacheClusterTopicParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnDistributedCacheClusterTopicParamsWithTimeout creates a new GetMsgVpnDistributedCacheClusterTopicParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMsgVpnDistributedCacheClusterTopicParamsWithTimeout(timeout time.Duration) *GetMsgVpnDistributedCacheClusterTopicParams {
	var ()
	return &GetMsgVpnDistributedCacheClusterTopicParams{

		timeout: timeout,
	}
}

// NewGetMsgVpnDistributedCacheClusterTopicParamsWithContext creates a new GetMsgVpnDistributedCacheClusterTopicParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMsgVpnDistributedCacheClusterTopicParamsWithContext(ctx context.Context) *GetMsgVpnDistributedCacheClusterTopicParams {
	var ()
	return &GetMsgVpnDistributedCacheClusterTopicParams{

		Context: ctx,
	}
}

// NewGetMsgVpnDistributedCacheClusterTopicParamsWithHTTPClient creates a new GetMsgVpnDistributedCacheClusterTopicParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMsgVpnDistributedCacheClusterTopicParamsWithHTTPClient(client *http.Client) *GetMsgVpnDistributedCacheClusterTopicParams {
	var ()
	return &GetMsgVpnDistributedCacheClusterTopicParams{
		HTTPClient: client,
	}
}

/*GetMsgVpnDistributedCacheClusterTopicParams contains all the parameters to send to the API endpoint
for the get msg vpn distributed cache cluster topic operation typically these are written to a http.Request
*/
type GetMsgVpnDistributedCacheClusterTopicParams struct {

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
	/*Topic
	  The value of the Topic in the form a/b/c.

	*/
	Topic string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) WithTimeout(timeout time.Duration) *GetMsgVpnDistributedCacheClusterTopicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) WithContext(ctx context.Context) *GetMsgVpnDistributedCacheClusterTopicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) WithHTTPClient(client *http.Client) *GetMsgVpnDistributedCacheClusterTopicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCacheName adds the cacheName to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) WithCacheName(cacheName string) *GetMsgVpnDistributedCacheClusterTopicParams {
	o.SetCacheName(cacheName)
	return o
}

// SetCacheName adds the cacheName to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) SetCacheName(cacheName string) {
	o.CacheName = cacheName
}

// WithClusterName adds the clusterName to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) WithClusterName(clusterName string) *GetMsgVpnDistributedCacheClusterTopicParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnDistributedCacheClusterTopicParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) WithSelect(selectVar []string) *GetMsgVpnDistributedCacheClusterTopicParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithTopic adds the topic to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) WithTopic(topic string) *GetMsgVpnDistributedCacheClusterTopicParams {
	o.SetTopic(topic)
	return o
}

// SetTopic adds the topic to the get msg vpn distributed cache cluster topic params
func (o *GetMsgVpnDistributedCacheClusterTopicParams) SetTopic(topic string) {
	o.Topic = topic
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnDistributedCacheClusterTopicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param topic
	if err := r.SetPathParam("topic", o.Topic); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
