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

// NewCreateDmrClusterLinkTLSTrustedCommonNameParams creates a new CreateDmrClusterLinkTLSTrustedCommonNameParams object
// with the default values initialized.
func NewCreateDmrClusterLinkTLSTrustedCommonNameParams() *CreateDmrClusterLinkTLSTrustedCommonNameParams {
	var ()
	return &CreateDmrClusterLinkTLSTrustedCommonNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDmrClusterLinkTLSTrustedCommonNameParamsWithTimeout creates a new CreateDmrClusterLinkTLSTrustedCommonNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDmrClusterLinkTLSTrustedCommonNameParamsWithTimeout(timeout time.Duration) *CreateDmrClusterLinkTLSTrustedCommonNameParams {
	var ()
	return &CreateDmrClusterLinkTLSTrustedCommonNameParams{

		timeout: timeout,
	}
}

// NewCreateDmrClusterLinkTLSTrustedCommonNameParamsWithContext creates a new CreateDmrClusterLinkTLSTrustedCommonNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDmrClusterLinkTLSTrustedCommonNameParamsWithContext(ctx context.Context) *CreateDmrClusterLinkTLSTrustedCommonNameParams {
	var ()
	return &CreateDmrClusterLinkTLSTrustedCommonNameParams{

		Context: ctx,
	}
}

// NewCreateDmrClusterLinkTLSTrustedCommonNameParamsWithHTTPClient creates a new CreateDmrClusterLinkTLSTrustedCommonNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDmrClusterLinkTLSTrustedCommonNameParamsWithHTTPClient(client *http.Client) *CreateDmrClusterLinkTLSTrustedCommonNameParams {
	var ()
	return &CreateDmrClusterLinkTLSTrustedCommonNameParams{
		HTTPClient: client,
	}
}

/*CreateDmrClusterLinkTLSTrustedCommonNameParams contains all the parameters to send to the API endpoint
for the create dmr cluster link Tls trusted common name operation typically these are written to a http.Request
*/
type CreateDmrClusterLinkTLSTrustedCommonNameParams struct {

	/*Body
	  The Trusted Common Name object's attributes.

	*/
	Body *models.DmrClusterLinkTLSTrustedCommonName
	/*DmrClusterName
	  The name of the Cluster.

	*/
	DmrClusterName string
	/*RemoteNodeName
	  The name of the node at the remote end of the Link.

	*/
	RemoteNodeName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create dmr cluster link Tls trusted common name params
func (o *CreateDmrClusterLinkTLSTrustedCommonNameParams) WithTimeout(timeout time.Duration) *CreateDmrClusterLinkTLSTrustedCommonNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create dmr cluster link Tls trusted common name params
func (o *CreateDmrClusterLinkTLSTrustedCommonNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create dmr cluster link Tls trusted common name params
func (o *CreateDmrClusterLinkTLSTrustedCommonNameParams) WithContext(ctx context.Context) *CreateDmrClusterLinkTLSTrustedCommonNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create dmr cluster link Tls trusted common name params
func (o *CreateDmrClusterLinkTLSTrustedCommonNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create dmr cluster link Tls trusted common name params
func (o *CreateDmrClusterLinkTLSTrustedCommonNameParams) WithHTTPClient(client *http.Client) *CreateDmrClusterLinkTLSTrustedCommonNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create dmr cluster link Tls trusted common name params
func (o *CreateDmrClusterLinkTLSTrustedCommonNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create dmr cluster link Tls trusted common name params
func (o *CreateDmrClusterLinkTLSTrustedCommonNameParams) WithBody(body *models.DmrClusterLinkTLSTrustedCommonName) *CreateDmrClusterLinkTLSTrustedCommonNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create dmr cluster link Tls trusted common name params
func (o *CreateDmrClusterLinkTLSTrustedCommonNameParams) SetBody(body *models.DmrClusterLinkTLSTrustedCommonName) {
	o.Body = body
}

// WithDmrClusterName adds the dmrClusterName to the create dmr cluster link Tls trusted common name params
func (o *CreateDmrClusterLinkTLSTrustedCommonNameParams) WithDmrClusterName(dmrClusterName string) *CreateDmrClusterLinkTLSTrustedCommonNameParams {
	o.SetDmrClusterName(dmrClusterName)
	return o
}

// SetDmrClusterName adds the dmrClusterName to the create dmr cluster link Tls trusted common name params
func (o *CreateDmrClusterLinkTLSTrustedCommonNameParams) SetDmrClusterName(dmrClusterName string) {
	o.DmrClusterName = dmrClusterName
}

// WithRemoteNodeName adds the remoteNodeName to the create dmr cluster link Tls trusted common name params
func (o *CreateDmrClusterLinkTLSTrustedCommonNameParams) WithRemoteNodeName(remoteNodeName string) *CreateDmrClusterLinkTLSTrustedCommonNameParams {
	o.SetRemoteNodeName(remoteNodeName)
	return o
}

// SetRemoteNodeName adds the remoteNodeName to the create dmr cluster link Tls trusted common name params
func (o *CreateDmrClusterLinkTLSTrustedCommonNameParams) SetRemoteNodeName(remoteNodeName string) {
	o.RemoteNodeName = remoteNodeName
}

// WithSelect adds the selectVar to the create dmr cluster link Tls trusted common name params
func (o *CreateDmrClusterLinkTLSTrustedCommonNameParams) WithSelect(selectVar []string) *CreateDmrClusterLinkTLSTrustedCommonNameParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the create dmr cluster link Tls trusted common name params
func (o *CreateDmrClusterLinkTLSTrustedCommonNameParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDmrClusterLinkTLSTrustedCommonNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param remoteNodeName
	if err := r.SetPathParam("remoteNodeName", o.RemoteNodeName); err != nil {
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
