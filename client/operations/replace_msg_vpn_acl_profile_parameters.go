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

// NewReplaceMsgVpnACLProfileParams creates a new ReplaceMsgVpnACLProfileParams object
// with the default values initialized.
func NewReplaceMsgVpnACLProfileParams() *ReplaceMsgVpnACLProfileParams {
	var ()
	return &ReplaceMsgVpnACLProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMsgVpnACLProfileParamsWithTimeout creates a new ReplaceMsgVpnACLProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceMsgVpnACLProfileParamsWithTimeout(timeout time.Duration) *ReplaceMsgVpnACLProfileParams {
	var ()
	return &ReplaceMsgVpnACLProfileParams{

		timeout: timeout,
	}
}

// NewReplaceMsgVpnACLProfileParamsWithContext creates a new ReplaceMsgVpnACLProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceMsgVpnACLProfileParamsWithContext(ctx context.Context) *ReplaceMsgVpnACLProfileParams {
	var ()
	return &ReplaceMsgVpnACLProfileParams{

		Context: ctx,
	}
}

// NewReplaceMsgVpnACLProfileParamsWithHTTPClient creates a new ReplaceMsgVpnACLProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceMsgVpnACLProfileParamsWithHTTPClient(client *http.Client) *ReplaceMsgVpnACLProfileParams {
	var ()
	return &ReplaceMsgVpnACLProfileParams{
		HTTPClient: client,
	}
}

/*ReplaceMsgVpnACLProfileParams contains all the parameters to send to the API endpoint
for the replace msg vpn Acl profile operation typically these are written to a http.Request
*/
type ReplaceMsgVpnACLProfileParams struct {

	/*ACLProfileName
	  The name of the ACL Profile.

	*/
	ACLProfileName string
	/*Body
	  The ACL Profile object's attributes.

	*/
	Body *models.MsgVpnACLProfile
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

// WithTimeout adds the timeout to the replace msg vpn Acl profile params
func (o *ReplaceMsgVpnACLProfileParams) WithTimeout(timeout time.Duration) *ReplaceMsgVpnACLProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace msg vpn Acl profile params
func (o *ReplaceMsgVpnACLProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace msg vpn Acl profile params
func (o *ReplaceMsgVpnACLProfileParams) WithContext(ctx context.Context) *ReplaceMsgVpnACLProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace msg vpn Acl profile params
func (o *ReplaceMsgVpnACLProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace msg vpn Acl profile params
func (o *ReplaceMsgVpnACLProfileParams) WithHTTPClient(client *http.Client) *ReplaceMsgVpnACLProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace msg vpn Acl profile params
func (o *ReplaceMsgVpnACLProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithACLProfileName adds the aCLProfileName to the replace msg vpn Acl profile params
func (o *ReplaceMsgVpnACLProfileParams) WithACLProfileName(aCLProfileName string) *ReplaceMsgVpnACLProfileParams {
	o.SetACLProfileName(aCLProfileName)
	return o
}

// SetACLProfileName adds the aclProfileName to the replace msg vpn Acl profile params
func (o *ReplaceMsgVpnACLProfileParams) SetACLProfileName(aCLProfileName string) {
	o.ACLProfileName = aCLProfileName
}

// WithBody adds the body to the replace msg vpn Acl profile params
func (o *ReplaceMsgVpnACLProfileParams) WithBody(body *models.MsgVpnACLProfile) *ReplaceMsgVpnACLProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace msg vpn Acl profile params
func (o *ReplaceMsgVpnACLProfileParams) SetBody(body *models.MsgVpnACLProfile) {
	o.Body = body
}

// WithMsgVpnName adds the msgVpnName to the replace msg vpn Acl profile params
func (o *ReplaceMsgVpnACLProfileParams) WithMsgVpnName(msgVpnName string) *ReplaceMsgVpnACLProfileParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the replace msg vpn Acl profile params
func (o *ReplaceMsgVpnACLProfileParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the replace msg vpn Acl profile params
func (o *ReplaceMsgVpnACLProfileParams) WithSelect(selectVar []string) *ReplaceMsgVpnACLProfileParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace msg vpn Acl profile params
func (o *ReplaceMsgVpnACLProfileParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceMsgVpnACLProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param aclProfileName
	if err := r.SetPathParam("aclProfileName", o.ACLProfileName); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
