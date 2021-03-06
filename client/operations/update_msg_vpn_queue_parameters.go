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

// NewUpdateMsgVpnQueueParams creates a new UpdateMsgVpnQueueParams object
// with the default values initialized.
func NewUpdateMsgVpnQueueParams() *UpdateMsgVpnQueueParams {
	var ()
	return &UpdateMsgVpnQueueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMsgVpnQueueParamsWithTimeout creates a new UpdateMsgVpnQueueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateMsgVpnQueueParamsWithTimeout(timeout time.Duration) *UpdateMsgVpnQueueParams {
	var ()
	return &UpdateMsgVpnQueueParams{

		timeout: timeout,
	}
}

// NewUpdateMsgVpnQueueParamsWithContext creates a new UpdateMsgVpnQueueParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateMsgVpnQueueParamsWithContext(ctx context.Context) *UpdateMsgVpnQueueParams {
	var ()
	return &UpdateMsgVpnQueueParams{

		Context: ctx,
	}
}

// NewUpdateMsgVpnQueueParamsWithHTTPClient creates a new UpdateMsgVpnQueueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateMsgVpnQueueParamsWithHTTPClient(client *http.Client) *UpdateMsgVpnQueueParams {
	var ()
	return &UpdateMsgVpnQueueParams{
		HTTPClient: client,
	}
}

/*UpdateMsgVpnQueueParams contains all the parameters to send to the API endpoint
for the update msg vpn queue operation typically these are written to a http.Request
*/
type UpdateMsgVpnQueueParams struct {

	/*Body
	  The Queue object's attributes.

	*/
	Body *models.MsgVpnQueue
	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string
	/*QueueName
	  The name of the Queue.

	*/
	QueueName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update msg vpn queue params
func (o *UpdateMsgVpnQueueParams) WithTimeout(timeout time.Duration) *UpdateMsgVpnQueueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update msg vpn queue params
func (o *UpdateMsgVpnQueueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update msg vpn queue params
func (o *UpdateMsgVpnQueueParams) WithContext(ctx context.Context) *UpdateMsgVpnQueueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update msg vpn queue params
func (o *UpdateMsgVpnQueueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update msg vpn queue params
func (o *UpdateMsgVpnQueueParams) WithHTTPClient(client *http.Client) *UpdateMsgVpnQueueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update msg vpn queue params
func (o *UpdateMsgVpnQueueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update msg vpn queue params
func (o *UpdateMsgVpnQueueParams) WithBody(body *models.MsgVpnQueue) *UpdateMsgVpnQueueParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update msg vpn queue params
func (o *UpdateMsgVpnQueueParams) SetBody(body *models.MsgVpnQueue) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the update msg vpn queue params
func (o *UpdateMsgVpnQueueParams) WithMsgVpnName(msgVpnName string) *UpdateMsgVpnQueueParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the update msg vpn queue params
func (o *UpdateMsgVpnQueueParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithQueueName adds the queueName to the update msg vpn queue params
func (o *UpdateMsgVpnQueueParams) WithQueueName(queueName string) *UpdateMsgVpnQueueParams {
	o.SetQueueName(queueName)
	return o
}

// SetQueueName adds the queueName to the update msg vpn queue params
func (o *UpdateMsgVpnQueueParams) SetQueueName(queueName string) {
	o.QueueName = queueName
}

// WithSelect adds the selectVar to the update msg vpn queue params
func (o *UpdateMsgVpnQueueParams) WithSelect(selectVar []string) *UpdateMsgVpnQueueParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the update msg vpn queue params
func (o *UpdateMsgVpnQueueParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMsgVpnQueueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param queueName
	if err := r.SetPathParam("queueName", o.QueueName); err != nil {
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
