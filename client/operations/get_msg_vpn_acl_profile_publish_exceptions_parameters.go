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

// NewGetMsgVpnACLProfilePublishExceptionsParams creates a new GetMsgVpnACLProfilePublishExceptionsParams object
// with the default values initialized.
func NewGetMsgVpnACLProfilePublishExceptionsParams() *GetMsgVpnACLProfilePublishExceptionsParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnACLProfilePublishExceptionsParams{
		Count: &countDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMsgVpnACLProfilePublishExceptionsParamsWithTimeout creates a new GetMsgVpnACLProfilePublishExceptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMsgVpnACLProfilePublishExceptionsParamsWithTimeout(timeout time.Duration) *GetMsgVpnACLProfilePublishExceptionsParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnACLProfilePublishExceptionsParams{
		Count: &countDefault,

		timeout: timeout,
	}
}

// NewGetMsgVpnACLProfilePublishExceptionsParamsWithContext creates a new GetMsgVpnACLProfilePublishExceptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMsgVpnACLProfilePublishExceptionsParamsWithContext(ctx context.Context) *GetMsgVpnACLProfilePublishExceptionsParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnACLProfilePublishExceptionsParams{
		Count: &countDefault,

		Context: ctx,
	}
}

// NewGetMsgVpnACLProfilePublishExceptionsParamsWithHTTPClient creates a new GetMsgVpnACLProfilePublishExceptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMsgVpnACLProfilePublishExceptionsParamsWithHTTPClient(client *http.Client) *GetMsgVpnACLProfilePublishExceptionsParams {
	var (
		countDefault = int64(10)
	)
	return &GetMsgVpnACLProfilePublishExceptionsParams{
		Count:      &countDefault,
		HTTPClient: client,
	}
}

/*GetMsgVpnACLProfilePublishExceptionsParams contains all the parameters to send to the API endpoint
for the get msg vpn Acl profile publish exceptions operation typically these are written to a http.Request
*/
type GetMsgVpnACLProfilePublishExceptionsParams struct {

	/*ACLProfileName
	  The name of the ACL Profile.

	*/
	ACLProfileName string
	/*Count
	  Limit the count of objects in the response. See the documentation for the `count` parameter.

	*/
	Count *int64
	/*Cursor
	  The cursor, or position, for the next page of objects. See the documentation for the `cursor` parameter.

	*/
	Cursor *string
	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string
	/*Where
	  Include in the response only objects where certain conditions are true. See the the documentation for the `where` parameter.

	*/
	Where []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) WithTimeout(timeout time.Duration) *GetMsgVpnACLProfilePublishExceptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) WithContext(ctx context.Context) *GetMsgVpnACLProfilePublishExceptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) WithHTTPClient(client *http.Client) *GetMsgVpnACLProfilePublishExceptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithACLProfileName adds the aCLProfileName to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) WithACLProfileName(aCLProfileName string) *GetMsgVpnACLProfilePublishExceptionsParams {
	o.SetACLProfileName(aCLProfileName)
	return o
}

// SetACLProfileName adds the aclProfileName to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) SetACLProfileName(aCLProfileName string) {
	o.ACLProfileName = aCLProfileName
}

// WithCount adds the count to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) WithCount(count *int64) *GetMsgVpnACLProfilePublishExceptionsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) SetCount(count *int64) {
	o.Count = count
}

// WithCursor adds the cursor to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) WithCursor(cursor *string) *GetMsgVpnACLProfilePublishExceptionsParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithMsgVpnName adds the msgVpnName to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) WithMsgVpnName(msgVpnName string) *GetMsgVpnACLProfilePublishExceptionsParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) WithSelect(selectVar []string) *GetMsgVpnACLProfilePublishExceptionsParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithWhere adds the where to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) WithWhere(where []string) *GetMsgVpnACLProfilePublishExceptionsParams {
	o.SetWhere(where)
	return o
}

// SetWhere adds the where to the get msg vpn Acl profile publish exceptions params
func (o *GetMsgVpnACLProfilePublishExceptionsParams) SetWhere(where []string) {
	o.Where = where
}

// WriteToRequest writes these params to a swagger request
func (o *GetMsgVpnACLProfilePublishExceptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param aclProfileName
	if err := r.SetPathParam("aclProfileName", o.ACLProfileName); err != nil {
		return err
	}

	if o.Count != nil {

		// query param count
		var qrCount int64
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string
		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {
			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
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

	valuesWhere := o.Where

	joinedWhere := swag.JoinByFormat(valuesWhere, "csv")
	// query array param where
	if err := r.SetQueryParam("where", joinedWhere...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
