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

// NewGetMsgVpnReplayLogParams creates a new GetMsgVpnReplayLogParams object
// with the default values initialized.
func NewGetMsgVpnReplayLogParams() *GetMsgVpnReplayLogParams {
	var ()
	return &GetMsgVpnReplayLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnReplayLogParamsWithTimeout creates a new GetMsgVpnReplayLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMsgVpnReplayLogParamsWithTimeout(timeout time.Duration) *GetMsgVpnReplayLogParams {
	var ()
	return &GetMsgVpnReplayLogParams{

		timeout: timeout,
	}
}

// NewGetMsgVpnReplayLogParamsWithContext creates a new GetMsgVpnReplayLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMsgVpnReplayLogParamsWithContext(ctx context.Context) *GetMsgVpnReplayLogParams {
	var ()
	return &GetMsgVpnReplayLogParams{

		Context: ctx,
	}
}

// NewGetMsgVpnReplayLogParamsWithHTTPClient creates a new GetMsgVpnReplayLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMsgVpnReplayLogParamsWithHTTPClient(client *http.Client) *GetMsgVpnReplayLogParams {
	var ()
	return &GetMsgVpnReplayLogParams{
		HTTPClient: client,
	}
}

/*GetMsgVpnReplayLogParams contains all the parameters to send to the API endpoint
for the get msg vpn replay log operation typically these are written to a http.Request
*/
type GetMsgVpnReplayLogParams struct {

	/*MsgVpnName
	  The msgVpnName of the Message VPN.

	*/
	MsgVpnName string
	/*ReplayLogName
	  The replayLogName of the ReplayLog.

	*/
	ReplayLogName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See [Select](#select "Description of the syntax of the `select` parameter").

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get msg vpn replay log params
func (o *GetMsgVpnReplayLogParams) WithTimeout(timeout time.Duration) *GetMsgVpnReplayLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn replay log params
func (o *GetMsgVpnReplayLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn replay log params
func (o *GetMsgVpnReplayLogParams) WithContext(ctx context.Context) *GetMsgVpnReplayLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn replay log params
func (o *GetMsgVpnReplayLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn replay log params
func (o *GetMsgVpnReplayLogParams) WithHTTPClient(client *http.Client) *GetMsgVpnReplayLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn replay log params
func (o *GetMsgVpnReplayLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn replay log params
func (o *GetMsgVpnReplayLogParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnReplayLogParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn replay log params
func (o *GetMsgVpnReplayLogParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithReplayLogName adds the replayLogName to the get msg vpn replay log params
func (o *GetMsgVpnReplayLogParams) WithReplayLogName(replayLogName string) *GetMsgVpnReplayLogParams {
	o.SetReplayLogName(replayLogName)
	return o
}

// SetReplayLogName adds the replayLogName to the get msg vpn replay log params
func (o *GetMsgVpnReplayLogParams) SetReplayLogName(replayLogName string) {
	o.ReplayLogName = replayLogName
}

// WithSelect adds the selectVar to the get msg vpn replay log params
func (o *GetMsgVpnReplayLogParams) WithSelect(selectVar []string) *GetMsgVpnReplayLogParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn replay log params
func (o *GetMsgVpnReplayLogParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnReplayLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param msgVpnName
	if err := r.SetPathParam("msgVpnName", o.MsgVpnName); err != nil {
		return err
	}

	// path param replayLogName
	if err := r.SetPathParam("replayLogName", o.ReplayLogName); err != nil {
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