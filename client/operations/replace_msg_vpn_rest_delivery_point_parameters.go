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

// NewReplaceMsgVpnRestDeliveryPointParams creates a new ReplaceMsgVpnRestDeliveryPointParams object
// with the default values initialized.
func NewReplaceMsgVpnRestDeliveryPointParams() *ReplaceMsgVpnRestDeliveryPointParams {
	var ()
	return &ReplaceMsgVpnRestDeliveryPointParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMsgVpnRestDeliveryPointParamsWithTimeout creates a new ReplaceMsgVpnRestDeliveryPointParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceMsgVpnRestDeliveryPointParamsWithTimeout(timeout time.Duration) *ReplaceMsgVpnRestDeliveryPointParams {
	var ()
	return &ReplaceMsgVpnRestDeliveryPointParams{

		timeout: timeout,
	}
}

// NewReplaceMsgVpnRestDeliveryPointParamsWithContext creates a new ReplaceMsgVpnRestDeliveryPointParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceMsgVpnRestDeliveryPointParamsWithContext(ctx context.Context) *ReplaceMsgVpnRestDeliveryPointParams {
	var ()
	return &ReplaceMsgVpnRestDeliveryPointParams{

		Context: ctx,
	}
}

// NewReplaceMsgVpnRestDeliveryPointParamsWithHTTPClient creates a new ReplaceMsgVpnRestDeliveryPointParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceMsgVpnRestDeliveryPointParamsWithHTTPClient(client *http.Client) *ReplaceMsgVpnRestDeliveryPointParams {
	var ()
	return &ReplaceMsgVpnRestDeliveryPointParams{
		HTTPClient: client,
	}
}

/*ReplaceMsgVpnRestDeliveryPointParams contains all the parameters to send to the API endpoint
for the replace msg vpn rest delivery point operation typically these are written to a http.Request
*/
type ReplaceMsgVpnRestDeliveryPointParams struct {

	/*Body
	  The REST Delivery Point object's attributes.

	*/
	Body *models.MsgVpnRestDeliveryPoint
	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string
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

// WithTimeout adds the timeout to the replace msg vpn rest delivery point params
func (o *ReplaceMsgVpnRestDeliveryPointParams) WithTimeout(timeout time.Duration) *ReplaceMsgVpnRestDeliveryPointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace msg vpn rest delivery point params
func (o *ReplaceMsgVpnRestDeliveryPointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace msg vpn rest delivery point params
func (o *ReplaceMsgVpnRestDeliveryPointParams) WithContext(ctx context.Context) *ReplaceMsgVpnRestDeliveryPointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace msg vpn rest delivery point params
func (o *ReplaceMsgVpnRestDeliveryPointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace msg vpn rest delivery point params
func (o *ReplaceMsgVpnRestDeliveryPointParams) WithHTTPClient(client *http.Client) *ReplaceMsgVpnRestDeliveryPointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace msg vpn rest delivery point params
func (o *ReplaceMsgVpnRestDeliveryPointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace msg vpn rest delivery point params
func (o *ReplaceMsgVpnRestDeliveryPointParams) WithBody(body *models.MsgVpnRestDeliveryPoint) *ReplaceMsgVpnRestDeliveryPointParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace msg vpn rest delivery point params
func (o *ReplaceMsgVpnRestDeliveryPointParams) SetBody(body *models.MsgVpnRestDeliveryPoint) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the replace msg vpn rest delivery point params
func (o *ReplaceMsgVpnRestDeliveryPointParams) WithMsgVpnName(msgVpnName string) *ReplaceMsgVpnRestDeliveryPointParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the replace msg vpn rest delivery point params
func (o *ReplaceMsgVpnRestDeliveryPointParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithRestDeliveryPointName adds the restDeliveryPointName to the replace msg vpn rest delivery point params
func (o *ReplaceMsgVpnRestDeliveryPointParams) WithRestDeliveryPointName(restDeliveryPointName string) *ReplaceMsgVpnRestDeliveryPointParams {
	o.SetRestDeliveryPointName(restDeliveryPointName)
	return o
}

// SetRestDeliveryPointName adds the restDeliveryPointName to the replace msg vpn rest delivery point params
func (o *ReplaceMsgVpnRestDeliveryPointParams) SetRestDeliveryPointName(restDeliveryPointName string) {
	o.RestDeliveryPointName = restDeliveryPointName
}

// WithSelect adds the selectVar to the replace msg vpn rest delivery point params
func (o *ReplaceMsgVpnRestDeliveryPointParams) WithSelect(selectVar []string) *ReplaceMsgVpnRestDeliveryPointParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace msg vpn rest delivery point params
func (o *ReplaceMsgVpnRestDeliveryPointParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceMsgVpnRestDeliveryPointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
