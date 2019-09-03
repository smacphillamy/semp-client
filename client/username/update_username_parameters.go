// Code generated by go-swagger; DO NOT EDIT.

package username

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

// NewUpdateUsernameParams creates a new UpdateUsernameParams object
// with the default values initialized.
func NewUpdateUsernameParams() *UpdateUsernameParams {
	var ()
	return &UpdateUsernameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUsernameParamsWithTimeout creates a new UpdateUsernameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUsernameParamsWithTimeout(timeout time.Duration) *UpdateUsernameParams {
	var ()
	return &UpdateUsernameParams{

		timeout: timeout,
	}
}

// NewUpdateUsernameParamsWithContext creates a new UpdateUsernameParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUsernameParamsWithContext(ctx context.Context) *UpdateUsernameParams {
	var ()
	return &UpdateUsernameParams{

		Context: ctx,
	}
}

// NewUpdateUsernameParamsWithHTTPClient creates a new UpdateUsernameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUsernameParamsWithHTTPClient(client *http.Client) *UpdateUsernameParams {
	var ()
	return &UpdateUsernameParams{
		HTTPClient: client,
	}
}

/*UpdateUsernameParams contains all the parameters to send to the API endpoint
for the update username operation typically these are written to a http.Request
*/
type UpdateUsernameParams struct {

	/*Body
	  The Username object's attributes.

	*/
	Body *models.Username
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string
	/*UserName
	  Username.

	*/
	UserName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update username params
func (o *UpdateUsernameParams) WithTimeout(timeout time.Duration) *UpdateUsernameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update username params
func (o *UpdateUsernameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update username params
func (o *UpdateUsernameParams) WithContext(ctx context.Context) *UpdateUsernameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update username params
func (o *UpdateUsernameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update username params
func (o *UpdateUsernameParams) WithHTTPClient(client *http.Client) *UpdateUsernameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update username params
func (o *UpdateUsernameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update username params
func (o *UpdateUsernameParams) WithBody(body *models.Username) *UpdateUsernameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update username params
func (o *UpdateUsernameParams) SetBody(body *models.Username) {
	o.Body = body
}

// WithSelect adds the selectVar to the update username params
func (o *UpdateUsernameParams) WithSelect(selectVar []string) *UpdateUsernameParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the update username params
func (o *UpdateUsernameParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithUserName adds the userName to the update username params
func (o *UpdateUsernameParams) WithUserName(userName string) *UpdateUsernameParams {
	o.SetUserName(userName)
	return o
}

// SetUserName adds the userName to the update username params
func (o *UpdateUsernameParams) SetUserName(userName string) {
	o.UserName = userName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUsernameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	valuesSelect := o.Select

	joinedSelect := swag.JoinByFormat(valuesSelect, "csv")
	// query array param select
	if err := r.SetQueryParam("select", joinedSelect...); err != nil {
		return err
	}

	// path param userName
	if err := r.SetPathParam("userName", o.UserName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
