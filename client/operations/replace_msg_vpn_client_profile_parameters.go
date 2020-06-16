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

// NewReplaceMsgVpnClientProfileParams creates a new ReplaceMsgVpnClientProfileParams object
// with the default values initialized.
func NewReplaceMsgVpnClientProfileParams() *ReplaceMsgVpnClientProfileParams {
	var ()
	return &ReplaceMsgVpnClientProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMsgVpnClientProfileParamsWithTimeout creates a new ReplaceMsgVpnClientProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceMsgVpnClientProfileParamsWithTimeout(timeout time.Duration) *ReplaceMsgVpnClientProfileParams {
	var ()
	return &ReplaceMsgVpnClientProfileParams{

		timeout: timeout,
	}
}

// NewReplaceMsgVpnClientProfileParamsWithContext creates a new ReplaceMsgVpnClientProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceMsgVpnClientProfileParamsWithContext(ctx context.Context) *ReplaceMsgVpnClientProfileParams {
	var ()
	return &ReplaceMsgVpnClientProfileParams{

		Context: ctx,
	}
}

// NewReplaceMsgVpnClientProfileParamsWithHTTPClient creates a new ReplaceMsgVpnClientProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceMsgVpnClientProfileParamsWithHTTPClient(client *http.Client) *ReplaceMsgVpnClientProfileParams {
	var ()
	return &ReplaceMsgVpnClientProfileParams{
		HTTPClient: client,
	}
}

/*ReplaceMsgVpnClientProfileParams contains all the parameters to send to the API endpoint
for the replace msg vpn client profile operation typically these are written to a http.Request
*/
type ReplaceMsgVpnClientProfileParams struct {

	/*Body
	  The Client Profile object's attributes.

	*/
	Body *models.MsgVpnClientProfile
	/*ClientProfileName
	  The name of the Client Profile.

	*/
	ClientProfileName string
	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) WithTimeout(timeout time.Duration) *ReplaceMsgVpnClientProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) WithContext(ctx context.Context) *ReplaceMsgVpnClientProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) WithHTTPClient(client *http.Client) *ReplaceMsgVpnClientProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) WithBody(body *models.MsgVpnClientProfile) *ReplaceMsgVpnClientProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) SetBody(body *models.MsgVpnClientProfile) {
	o.Body = body
}

// WithClientProfileName adds the clientProfileName to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) WithClientProfileName(clientProfileName string) *ReplaceMsgVpnClientProfileParams {
	o.SetClientProfileName(clientProfileName)
	return o
}

// SetClientProfileName adds the clientProfileName to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) SetClientProfileName(clientProfileName string) {
	o.ClientProfileName = clientProfileName
}

// WithMsgVpnName adds the msgVpnName to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) WithMsgVpnName(msgVpnName string) *ReplaceMsgVpnClientProfileParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) WithSelect(selectVar []string) *ReplaceMsgVpnClientProfileParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace msg vpn client profile params
func (o *ReplaceMsgVpnClientProfileParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceMsgVpnClientProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param clientProfileName
	if err := r.SetPathParam("clientProfileName", o.ClientProfileName); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
