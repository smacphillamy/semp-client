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

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteDmrClusterLinkTLSTrustedCommonNameParams creates a new DeleteDmrClusterLinkTLSTrustedCommonNameParams object
// with the default values initialized.
func NewDeleteDmrClusterLinkTLSTrustedCommonNameParams() *DeleteDmrClusterLinkTLSTrustedCommonNameParams {
	var ()
	return &DeleteDmrClusterLinkTLSTrustedCommonNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDmrClusterLinkTLSTrustedCommonNameParamsWithTimeout creates a new DeleteDmrClusterLinkTLSTrustedCommonNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDmrClusterLinkTLSTrustedCommonNameParamsWithTimeout(timeout time.Duration) *DeleteDmrClusterLinkTLSTrustedCommonNameParams {
	var ()
	return &DeleteDmrClusterLinkTLSTrustedCommonNameParams{

		timeout: timeout,
	}
}

// NewDeleteDmrClusterLinkTLSTrustedCommonNameParamsWithContext creates a new DeleteDmrClusterLinkTLSTrustedCommonNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDmrClusterLinkTLSTrustedCommonNameParamsWithContext(ctx context.Context) *DeleteDmrClusterLinkTLSTrustedCommonNameParams {
	var ()
	return &DeleteDmrClusterLinkTLSTrustedCommonNameParams{

		Context: ctx,
	}
}

// NewDeleteDmrClusterLinkTLSTrustedCommonNameParamsWithHTTPClient creates a new DeleteDmrClusterLinkTLSTrustedCommonNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDmrClusterLinkTLSTrustedCommonNameParamsWithHTTPClient(client *http.Client) *DeleteDmrClusterLinkTLSTrustedCommonNameParams {
	var ()
	return &DeleteDmrClusterLinkTLSTrustedCommonNameParams{
		HTTPClient: client,
	}
}

/*DeleteDmrClusterLinkTLSTrustedCommonNameParams contains all the parameters to send to the API endpoint
for the delete dmr cluster link Tls trusted common name operation typically these are written to a http.Request
*/
type DeleteDmrClusterLinkTLSTrustedCommonNameParams struct {

	/*DmrClusterName
	  The name of the Cluster.

	*/
	DmrClusterName string
	/*RemoteNodeName
	  The name of the node at the remote end of the Link.

	*/
	RemoteNodeName string
	/*TLSTrustedCommonName
	  The expected trusted common name of the remote certificate.

	*/
	TLSTrustedCommonName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete dmr cluster link Tls trusted common name params
func (o *DeleteDmrClusterLinkTLSTrustedCommonNameParams) WithTimeout(timeout time.Duration) *DeleteDmrClusterLinkTLSTrustedCommonNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete dmr cluster link Tls trusted common name params
func (o *DeleteDmrClusterLinkTLSTrustedCommonNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete dmr cluster link Tls trusted common name params
func (o *DeleteDmrClusterLinkTLSTrustedCommonNameParams) WithContext(ctx context.Context) *DeleteDmrClusterLinkTLSTrustedCommonNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete dmr cluster link Tls trusted common name params
func (o *DeleteDmrClusterLinkTLSTrustedCommonNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete dmr cluster link Tls trusted common name params
func (o *DeleteDmrClusterLinkTLSTrustedCommonNameParams) WithHTTPClient(client *http.Client) *DeleteDmrClusterLinkTLSTrustedCommonNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete dmr cluster link Tls trusted common name params
func (o *DeleteDmrClusterLinkTLSTrustedCommonNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDmrClusterName adds the dmrClusterName to the delete dmr cluster link Tls trusted common name params
func (o *DeleteDmrClusterLinkTLSTrustedCommonNameParams) WithDmrClusterName(dmrClusterName string) *DeleteDmrClusterLinkTLSTrustedCommonNameParams {
	o.SetDmrClusterName(dmrClusterName)
	return o
}

// SetDmrClusterName adds the dmrClusterName to the delete dmr cluster link Tls trusted common name params
func (o *DeleteDmrClusterLinkTLSTrustedCommonNameParams) SetDmrClusterName(dmrClusterName string) {
	o.DmrClusterName = dmrClusterName
}

// WithRemoteNodeName adds the remoteNodeName to the delete dmr cluster link Tls trusted common name params
func (o *DeleteDmrClusterLinkTLSTrustedCommonNameParams) WithRemoteNodeName(remoteNodeName string) *DeleteDmrClusterLinkTLSTrustedCommonNameParams {
	o.SetRemoteNodeName(remoteNodeName)
	return o
}

// SetRemoteNodeName adds the remoteNodeName to the delete dmr cluster link Tls trusted common name params
func (o *DeleteDmrClusterLinkTLSTrustedCommonNameParams) SetRemoteNodeName(remoteNodeName string) {
	o.RemoteNodeName = remoteNodeName
}

// WithTLSTrustedCommonName adds the tLSTrustedCommonName to the delete dmr cluster link Tls trusted common name params
func (o *DeleteDmrClusterLinkTLSTrustedCommonNameParams) WithTLSTrustedCommonName(tLSTrustedCommonName string) *DeleteDmrClusterLinkTLSTrustedCommonNameParams {
	o.SetTLSTrustedCommonName(tLSTrustedCommonName)
	return o
}

// SetTLSTrustedCommonName adds the tlsTrustedCommonName to the delete dmr cluster link Tls trusted common name params
func (o *DeleteDmrClusterLinkTLSTrustedCommonNameParams) SetTLSTrustedCommonName(tLSTrustedCommonName string) {
	o.TLSTrustedCommonName = tLSTrustedCommonName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDmrClusterLinkTLSTrustedCommonNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dmrClusterName
	if err := r.SetPathParam("dmrClusterName", o.DmrClusterName); err != nil {
		return err
	}

	// path param remoteNodeName
	if err := r.SetPathParam("remoteNodeName", o.RemoteNodeName); err != nil {
		return err
	}

	// path param tlsTrustedCommonName
	if err := r.SetPathParam("tlsTrustedCommonName", o.TLSTrustedCommonName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
