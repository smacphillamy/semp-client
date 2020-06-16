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

// NewUpdateMsgVpnJndiQueueParams creates a new UpdateMsgVpnJndiQueueParams object
// with the default values initialized.
func NewUpdateMsgVpnJndiQueueParams() *UpdateMsgVpnJndiQueueParams {
	var ()
	return &UpdateMsgVpnJndiQueueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateMsgVpnJndiQueueParamsWithTimeout creates a new UpdateMsgVpnJndiQueueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateMsgVpnJndiQueueParamsWithTimeout(timeout time.Duration) *UpdateMsgVpnJndiQueueParams {
	var ()
	return &UpdateMsgVpnJndiQueueParams{

		timeout: timeout,
	}
}

// NewUpdateMsgVpnJndiQueueParamsWithContext creates a new UpdateMsgVpnJndiQueueParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateMsgVpnJndiQueueParamsWithContext(ctx context.Context) *UpdateMsgVpnJndiQueueParams {
	var ()
	return &UpdateMsgVpnJndiQueueParams{

		Context: ctx,
	}
}

// NewUpdateMsgVpnJndiQueueParamsWithHTTPClient creates a new UpdateMsgVpnJndiQueueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateMsgVpnJndiQueueParamsWithHTTPClient(client *http.Client) *UpdateMsgVpnJndiQueueParams {
	var ()
	return &UpdateMsgVpnJndiQueueParams{
		HTTPClient: client,
	}
}

/*UpdateMsgVpnJndiQueueParams contains all the parameters to send to the API endpoint
for the update msg vpn jndi queue operation typically these are written to a http.Request
*/
type UpdateMsgVpnJndiQueueParams struct {

	/*Body
	  The JNDI Queue object's attributes.

	*/
	Body *models.MsgVpnJndiQueue
	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string
	/*QueueName
	  The JNDI name of the JMS Queue.

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

// WithTimeout adds the timeout to the update msg vpn jndi queue params
func (o *UpdateMsgVpnJndiQueueParams) WithTimeout(timeout time.Duration) *UpdateMsgVpnJndiQueueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update msg vpn jndi queue params
func (o *UpdateMsgVpnJndiQueueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update msg vpn jndi queue params
func (o *UpdateMsgVpnJndiQueueParams) WithContext(ctx context.Context) *UpdateMsgVpnJndiQueueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update msg vpn jndi queue params
func (o *UpdateMsgVpnJndiQueueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update msg vpn jndi queue params
func (o *UpdateMsgVpnJndiQueueParams) WithHTTPClient(client *http.Client) *UpdateMsgVpnJndiQueueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update msg vpn jndi queue params
func (o *UpdateMsgVpnJndiQueueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update msg vpn jndi queue params
func (o *UpdateMsgVpnJndiQueueParams) WithBody(body *models.MsgVpnJndiQueue) *UpdateMsgVpnJndiQueueParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update msg vpn jndi queue params
func (o *UpdateMsgVpnJndiQueueParams) SetBody(body *models.MsgVpnJndiQueue) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the update msg vpn jndi queue params
func (o *UpdateMsgVpnJndiQueueParams) WithMsgVpnName(msgVpnName string) *UpdateMsgVpnJndiQueueParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the update msg vpn jndi queue params
func (o *UpdateMsgVpnJndiQueueParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithQueueName adds the queueName to the update msg vpn jndi queue params
func (o *UpdateMsgVpnJndiQueueParams) WithQueueName(queueName string) *UpdateMsgVpnJndiQueueParams {
	o.SetQueueName(queueName)
	return o
}

// SetQueueName adds the queueName to the update msg vpn jndi queue params
func (o *UpdateMsgVpnJndiQueueParams) SetQueueName(queueName string) {
	o.QueueName = queueName
}

// WithSelect adds the selectVar to the update msg vpn jndi queue params
func (o *UpdateMsgVpnJndiQueueParams) WithSelect(selectVar []string) *UpdateMsgVpnJndiQueueParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the update msg vpn jndi queue params
func (o *UpdateMsgVpnJndiQueueParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateMsgVpnJndiQueueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
