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

// NewGetMsgVpnTopicEndpointParams creates a new GetMsgVpnTopicEndpointParams object
// with the default values initialized.
func NewGetMsgVpnTopicEndpointParams() *GetMsgVpnTopicEndpointParams {
	var ()
	return &GetMsgVpnTopicEndpointParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnTopicEndpointParamsWithTimeout creates a new GetMsgVpnTopicEndpointParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMsgVpnTopicEndpointParamsWithTimeout(timeout time.Duration) *GetMsgVpnTopicEndpointParams {
	var ()
	return &GetMsgVpnTopicEndpointParams{

		timeout: timeout,
	}
}

// NewGetMsgVpnTopicEndpointParamsWithContext creates a new GetMsgVpnTopicEndpointParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMsgVpnTopicEndpointParamsWithContext(ctx context.Context) *GetMsgVpnTopicEndpointParams {
	var ()
	return &GetMsgVpnTopicEndpointParams{

		Context: ctx,
	}
}

// NewGetMsgVpnTopicEndpointParamsWithHTTPClient creates a new GetMsgVpnTopicEndpointParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMsgVpnTopicEndpointParamsWithHTTPClient(client *http.Client) *GetMsgVpnTopicEndpointParams {
	var ()
	return &GetMsgVpnTopicEndpointParams{
		HTTPClient: client,
	}
}

/*GetMsgVpnTopicEndpointParams contains all the parameters to send to the API endpoint
for the get msg vpn topic endpoint operation typically these are written to a http.Request
*/
type GetMsgVpnTopicEndpointParams struct {

	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See [Select](#select "Description of the syntax of the `select` parameter").

	*/
	Select []string
	/*TopicEndpointName
	  The topicEndpointName of the Topic Endpoint.

	*/
	TopicEndpointName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get msg vpn topic endpoint params
func (o *GetMsgVpnTopicEndpointParams) WithTimeout(timeout time.Duration) *GetMsgVpnTopicEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn topic endpoint params
func (o *GetMsgVpnTopicEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn topic endpoint params
func (o *GetMsgVpnTopicEndpointParams) WithContext(ctx context.Context) *GetMsgVpnTopicEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn topic endpoint params
func (o *GetMsgVpnTopicEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn topic endpoint params
func (o *GetMsgVpnTopicEndpointParams) WithHTTPClient(client *http.Client) *GetMsgVpnTopicEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn topic endpoint params
func (o *GetMsgVpnTopicEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn topic endpoint params
func (o *GetMsgVpnTopicEndpointParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnTopicEndpointParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn topic endpoint params
func (o *GetMsgVpnTopicEndpointParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the get msg vpn topic endpoint params
func (o *GetMsgVpnTopicEndpointParams) WithSelect(selectVar []string) *GetMsgVpnTopicEndpointParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn topic endpoint params
func (o *GetMsgVpnTopicEndpointParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithTopicEndpointName adds the topicEndpointName to the get msg vpn topic endpoint params
func (o *GetMsgVpnTopicEndpointParams) WithTopicEndpointName(topicEndpointName string) *GetMsgVpnTopicEndpointParams {
	o.SetTopicEndpointName(topicEndpointName)
	return o
}

// SetTopicEndpointName adds the topicEndpointName to the get msg vpn topic endpoint params
func (o *GetMsgVpnTopicEndpointParams) SetTopicEndpointName(topicEndpointName string) {
	o.TopicEndpointName = topicEndpointName
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnTopicEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param topicEndpointName
	if err := r.SetPathParam("topicEndpointName", o.TopicEndpointName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}