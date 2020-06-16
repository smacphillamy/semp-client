// Code generated by go-swagger; DO NOT EDIT.

package dmr_cluster

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

// NewUpdateDmrClusterParams creates a new UpdateDmrClusterParams object
// with the default values initialized.
func NewUpdateDmrClusterParams() *UpdateDmrClusterParams {
	var ()
	return &UpdateDmrClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDmrClusterParamsWithTimeout creates a new UpdateDmrClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDmrClusterParamsWithTimeout(timeout time.Duration) *UpdateDmrClusterParams {
	var ()
	return &UpdateDmrClusterParams{

		timeout: timeout,
	}
}

// NewUpdateDmrClusterParamsWithContext creates a new UpdateDmrClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDmrClusterParamsWithContext(ctx context.Context) *UpdateDmrClusterParams {
	var ()
	return &UpdateDmrClusterParams{

		Context: ctx,
	}
}

// NewUpdateDmrClusterParamsWithHTTPClient creates a new UpdateDmrClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDmrClusterParamsWithHTTPClient(client *http.Client) *UpdateDmrClusterParams {
	var ()
	return &UpdateDmrClusterParams{
		HTTPClient: client,
	}
}

/*UpdateDmrClusterParams contains all the parameters to send to the API endpoint
for the update dmr cluster operation typically these are written to a http.Request
*/
type UpdateDmrClusterParams struct {

	/*Body
	  The Cluster object's attributes.

	*/
	Body *models.DmrCluster
	/*DmrClusterName
	  The name of the Cluster.

	*/
	DmrClusterName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update dmr cluster params
func (o *UpdateDmrClusterParams) WithTimeout(timeout time.Duration) *UpdateDmrClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update dmr cluster params
func (o *UpdateDmrClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update dmr cluster params
func (o *UpdateDmrClusterParams) WithContext(ctx context.Context) *UpdateDmrClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update dmr cluster params
func (o *UpdateDmrClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update dmr cluster params
func (o *UpdateDmrClusterParams) WithHTTPClient(client *http.Client) *UpdateDmrClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update dmr cluster params
func (o *UpdateDmrClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update dmr cluster params
func (o *UpdateDmrClusterParams) WithBody(body *models.DmrCluster) *UpdateDmrClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update dmr cluster params
func (o *UpdateDmrClusterParams) SetBody(body *models.DmrCluster) {
	o.Body = body
}

// WithDmrClusterName adds the dmrClusterName to the update dmr cluster params
func (o *UpdateDmrClusterParams) WithDmrClusterName(dmrClusterName string) *UpdateDmrClusterParams {
	o.SetDmrClusterName(dmrClusterName)
	return o
}

// SetDmrClusterName adds the dmrClusterName to the update dmr cluster params
func (o *UpdateDmrClusterParams) SetDmrClusterName(dmrClusterName string) {
	o.DmrClusterName = dmrClusterName
}

// WithSelect adds the selectVar to the update dmr cluster params
func (o *UpdateDmrClusterParams) WithSelect(selectVar []string) *UpdateDmrClusterParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the update dmr cluster params
func (o *UpdateDmrClusterParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDmrClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param dmrClusterName
	if err := r.SetPathParam("dmrClusterName", o.DmrClusterName); err != nil {
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
