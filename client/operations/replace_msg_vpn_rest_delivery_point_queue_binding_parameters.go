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

// NewReplaceMsgVpnRestDeliveryPointQueueBindingParams creates a new ReplaceMsgVpnRestDeliveryPointQueueBindingParams object
// with the default values initialized.
func NewReplaceMsgVpnRestDeliveryPointQueueBindingParams() *ReplaceMsgVpnRestDeliveryPointQueueBindingParams {
	var ()
	return &ReplaceMsgVpnRestDeliveryPointQueueBindingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMsgVpnRestDeliveryPointQueueBindingParamsWithTimeout creates a new ReplaceMsgVpnRestDeliveryPointQueueBindingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceMsgVpnRestDeliveryPointQueueBindingParamsWithTimeout(timeout time.Duration) *ReplaceMsgVpnRestDeliveryPointQueueBindingParams {
	var ()
	return &ReplaceMsgVpnRestDeliveryPointQueueBindingParams{

		timeout: timeout,
	}
}

// NewReplaceMsgVpnRestDeliveryPointQueueBindingParamsWithContext creates a new ReplaceMsgVpnRestDeliveryPointQueueBindingParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceMsgVpnRestDeliveryPointQueueBindingParamsWithContext(ctx context.Context) *ReplaceMsgVpnRestDeliveryPointQueueBindingParams {
	var ()
	return &ReplaceMsgVpnRestDeliveryPointQueueBindingParams{

		Context: ctx,
	}
}

// NewReplaceMsgVpnRestDeliveryPointQueueBindingParamsWithHTTPClient creates a new ReplaceMsgVpnRestDeliveryPointQueueBindingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceMsgVpnRestDeliveryPointQueueBindingParamsWithHTTPClient(client *http.Client) *ReplaceMsgVpnRestDeliveryPointQueueBindingParams {
	var ()
	return &ReplaceMsgVpnRestDeliveryPointQueueBindingParams{
		HTTPClient: client,
	}
}

/*ReplaceMsgVpnRestDeliveryPointQueueBindingParams contains all the parameters to send to the API endpoint
for the replace msg vpn rest delivery point queue binding operation typically these are written to a http.Request
*/
type ReplaceMsgVpnRestDeliveryPointQueueBindingParams struct {

	/*Body
	  The Queue Binding object's attributes.

	*/
	Body *models.MsgVpnRestDeliveryPointQueueBinding
	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string
	/*QueueBindingName
	  The name of a queue in the Message VPN.

	*/
	QueueBindingName string
	/*RestDeliveryPointName
	  The name of the REST Delivery Point.

	*/
	RestDeliveryPointName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) WithTimeout(timeout time.Duration) *ReplaceMsgVpnRestDeliveryPointQueueBindingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) WithContext(ctx context.Context) *ReplaceMsgVpnRestDeliveryPointQueueBindingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) WithHTTPClient(client *http.Client) *ReplaceMsgVpnRestDeliveryPointQueueBindingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) WithBody(body *models.MsgVpnRestDeliveryPointQueueBinding) *ReplaceMsgVpnRestDeliveryPointQueueBindingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) SetBody(body *models.MsgVpnRestDeliveryPointQueueBinding) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) WithMsgVpnName(msgVpnName string) *ReplaceMsgVpnRestDeliveryPointQueueBindingParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithQueueBindingName adds the queueBindingName to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) WithQueueBindingName(queueBindingName string) *ReplaceMsgVpnRestDeliveryPointQueueBindingParams {
	o.SetQueueBindingName(queueBindingName)
	return o
}

// SetQueueBindingName adds the queueBindingName to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) SetQueueBindingName(queueBindingName string) {
	o.QueueBindingName = queueBindingName
}

// WithRestDeliveryPointName adds the restDeliveryPointName to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) WithRestDeliveryPointName(restDeliveryPointName string) *ReplaceMsgVpnRestDeliveryPointQueueBindingParams {
	o.SetRestDeliveryPointName(restDeliveryPointName)
	return o
}

// SetRestDeliveryPointName adds the restDeliveryPointName to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) SetRestDeliveryPointName(restDeliveryPointName string) {
	o.RestDeliveryPointName = restDeliveryPointName
}

// WithSelect adds the selectVar to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) WithSelect(selectVar []string) *ReplaceMsgVpnRestDeliveryPointQueueBindingParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace msg vpn rest delivery point queue binding params
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceMsgVpnRestDeliveryPointQueueBindingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param queueBindingName
	if err := r.SetPathParam("queueBindingName", o.QueueBindingName); err != nil {
		return err
	}

	// path param restDeliveryPointName
	if err := r.SetPathParam("restDeliveryPointName", o.RestDeliveryPointName); err != nil {
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
