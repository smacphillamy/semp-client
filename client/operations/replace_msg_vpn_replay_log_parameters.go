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

// NewReplaceMsgVpnReplayLogParams creates a new ReplaceMsgVpnReplayLogParams object
// with the default values initialized.
func NewReplaceMsgVpnReplayLogParams() *ReplaceMsgVpnReplayLogParams {
	var ()
	return &ReplaceMsgVpnReplayLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMsgVpnReplayLogParamsWithTimeout creates a new ReplaceMsgVpnReplayLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceMsgVpnReplayLogParamsWithTimeout(timeout time.Duration) *ReplaceMsgVpnReplayLogParams {
	var ()
	return &ReplaceMsgVpnReplayLogParams{

		timeout: timeout,
	}
}

// NewReplaceMsgVpnReplayLogParamsWithContext creates a new ReplaceMsgVpnReplayLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceMsgVpnReplayLogParamsWithContext(ctx context.Context) *ReplaceMsgVpnReplayLogParams {
	var ()
	return &ReplaceMsgVpnReplayLogParams{

		Context: ctx,
	}
}

// NewReplaceMsgVpnReplayLogParamsWithHTTPClient creates a new ReplaceMsgVpnReplayLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceMsgVpnReplayLogParamsWithHTTPClient(client *http.Client) *ReplaceMsgVpnReplayLogParams {
	var ()
	return &ReplaceMsgVpnReplayLogParams{
		HTTPClient: client,
	}
}

/*ReplaceMsgVpnReplayLogParams contains all the parameters to send to the API endpoint
for the replace msg vpn replay log operation typically these are written to a http.Request
*/
type ReplaceMsgVpnReplayLogParams struct {

	/*Body
	  The Replay Log object's attributes.

	*/
	Body *models.MsgVpnReplayLog
	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string
	/*ReplayLogName
	  The name of the Replay Log.

	*/
	ReplayLogName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace msg vpn replay log params
func (o *ReplaceMsgVpnReplayLogParams) WithTimeout(timeout time.Duration) *ReplaceMsgVpnReplayLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace msg vpn replay log params
func (o *ReplaceMsgVpnReplayLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace msg vpn replay log params
func (o *ReplaceMsgVpnReplayLogParams) WithContext(ctx context.Context) *ReplaceMsgVpnReplayLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace msg vpn replay log params
func (o *ReplaceMsgVpnReplayLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace msg vpn replay log params
func (o *ReplaceMsgVpnReplayLogParams) WithHTTPClient(client *http.Client) *ReplaceMsgVpnReplayLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace msg vpn replay log params
func (o *ReplaceMsgVpnReplayLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace msg vpn replay log params
func (o *ReplaceMsgVpnReplayLogParams) WithBody(body *models.MsgVpnReplayLog) *ReplaceMsgVpnReplayLogParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace msg vpn replay log params
func (o *ReplaceMsgVpnReplayLogParams) SetBody(body *models.MsgVpnReplayLog) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the replace msg vpn replay log params
func (o *ReplaceMsgVpnReplayLogParams) WithMsgVpnName(msgVpnName string) *ReplaceMsgVpnReplayLogParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the replace msg vpn replay log params
func (o *ReplaceMsgVpnReplayLogParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithReplayLogName adds the replayLogName to the replace msg vpn replay log params
func (o *ReplaceMsgVpnReplayLogParams) WithReplayLogName(replayLogName string) *ReplaceMsgVpnReplayLogParams {
	o.SetReplayLogName(replayLogName)
	return o
}

// SetReplayLogName adds the replayLogName to the replace msg vpn replay log params
func (o *ReplaceMsgVpnReplayLogParams) SetReplayLogName(replayLogName string) {
	o.ReplayLogName = replayLogName
}

// WithSelect adds the selectVar to the replace msg vpn replay log params
func (o *ReplaceMsgVpnReplayLogParams) WithSelect(selectVar []string) *ReplaceMsgVpnReplayLogParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace msg vpn replay log params
func (o *ReplaceMsgVpnReplayLogParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceMsgVpnReplayLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
