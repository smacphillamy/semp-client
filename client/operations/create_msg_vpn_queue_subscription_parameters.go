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

// NewCreateMsgVpnQueueSubscriptionParams creates a new CreateMsgVpnQueueSubscriptionParams object
// with the default values initialized.
func NewCreateMsgVpnQueueSubscriptionParams() *CreateMsgVpnQueueSubscriptionParams {
	var ()
	return &CreateMsgVpnQueueSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMsgVpnQueueSubscriptionParamsWithTimeout creates a new CreateMsgVpnQueueSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMsgVpnQueueSubscriptionParamsWithTimeout(timeout time.Duration) *CreateMsgVpnQueueSubscriptionParams {
	var ()
	return &CreateMsgVpnQueueSubscriptionParams{

		timeout: timeout,
	}
}

// NewCreateMsgVpnQueueSubscriptionParamsWithContext creates a new CreateMsgVpnQueueSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMsgVpnQueueSubscriptionParamsWithContext(ctx context.Context) *CreateMsgVpnQueueSubscriptionParams {
	var ()
	return &CreateMsgVpnQueueSubscriptionParams{

		Context: ctx,
	}
}

// NewCreateMsgVpnQueueSubscriptionParamsWithHTTPClient creates a new CreateMsgVpnQueueSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMsgVpnQueueSubscriptionParamsWithHTTPClient(client *http.Client) *CreateMsgVpnQueueSubscriptionParams {
	var ()
	return &CreateMsgVpnQueueSubscriptionParams{
		HTTPClient: client,
	}
}

/*CreateMsgVpnQueueSubscriptionParams contains all the parameters to send to the API endpoint
for the create msg vpn queue subscription operation typically these are written to a http.Request
*/
type CreateMsgVpnQueueSubscriptionParams struct {

	/*Body
	  The Queue Subscription object's attributes.

	*/
	Body *models.MsgVpnQueueSubscription
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

// WithTimeout adds the timeout to the create msg vpn queue subscription params
func (o *CreateMsgVpnQueueSubscriptionParams) WithTimeout(timeout time.Duration) *CreateMsgVpnQueueSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create msg vpn queue subscription params
func (o *CreateMsgVpnQueueSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create msg vpn queue subscription params
func (o *CreateMsgVpnQueueSubscriptionParams) WithContext(ctx context.Context) *CreateMsgVpnQueueSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create msg vpn queue subscription params
func (o *CreateMsgVpnQueueSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create msg vpn queue subscription params
func (o *CreateMsgVpnQueueSubscriptionParams) WithHTTPClient(client *http.Client) *CreateMsgVpnQueueSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create msg vpn queue subscription params
func (o *CreateMsgVpnQueueSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create msg vpn queue subscription params
func (o *CreateMsgVpnQueueSubscriptionParams) WithBody(body *models.MsgVpnQueueSubscription) *CreateMsgVpnQueueSubscriptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create msg vpn queue subscription params
func (o *CreateMsgVpnQueueSubscriptionParams) SetBody(body *models.MsgVpnQueueSubscription) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the create msg vpn queue subscription params
func (o *CreateMsgVpnQueueSubscriptionParams) WithMsgVpnName(msgVpnName string) *CreateMsgVpnQueueSubscriptionParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the create msg vpn queue subscription params
func (o *CreateMsgVpnQueueSubscriptionParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithQueueName adds the queueName to the create msg vpn queue subscription params
func (o *CreateMsgVpnQueueSubscriptionParams) WithQueueName(queueName string) *CreateMsgVpnQueueSubscriptionParams {
	o.SetQueueName(queueName)
	return o
}

// SetQueueName adds the queueName to the create msg vpn queue subscription params
func (o *CreateMsgVpnQueueSubscriptionParams) SetQueueName(queueName string) {
	o.QueueName = queueName
}

// WithSelect adds the selectVar to the create msg vpn queue subscription params
func (o *CreateMsgVpnQueueSubscriptionParams) WithSelect(selectVar []string) *CreateMsgVpnQueueSubscriptionParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create msg vpn queue subscription params
func (o *CreateMsgVpnQueueSubscriptionParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMsgVpnQueueSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
