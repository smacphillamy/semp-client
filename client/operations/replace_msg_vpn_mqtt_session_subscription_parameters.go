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

// NewReplaceMsgVpnMqttSessionSubscriptionParams creates a new ReplaceMsgVpnMqttSessionSubscriptionParams object
// with the default values initialized.
func NewReplaceMsgVpnMqttSessionSubscriptionParams() *ReplaceMsgVpnMqttSessionSubscriptionParams {
	var ()
	return &ReplaceMsgVpnMqttSessionSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceMsgVpnMqttSessionSubscriptionParamsWithTimeout creates a new ReplaceMsgVpnMqttSessionSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceMsgVpnMqttSessionSubscriptionParamsWithTimeout(timeout time.Duration) *ReplaceMsgVpnMqttSessionSubscriptionParams {
	var ()
	return &ReplaceMsgVpnMqttSessionSubscriptionParams{

		timeout: timeout,
	}
}

// NewReplaceMsgVpnMqttSessionSubscriptionParamsWithContext creates a new ReplaceMsgVpnMqttSessionSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceMsgVpnMqttSessionSubscriptionParamsWithContext(ctx context.Context) *ReplaceMsgVpnMqttSessionSubscriptionParams {
	var ()
	return &ReplaceMsgVpnMqttSessionSubscriptionParams{

		Context: ctx,
	}
}

// NewReplaceMsgVpnMqttSessionSubscriptionParamsWithHTTPClient creates a new ReplaceMsgVpnMqttSessionSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceMsgVpnMqttSessionSubscriptionParamsWithHTTPClient(client *http.Client) *ReplaceMsgVpnMqttSessionSubscriptionParams {
	var ()
	return &ReplaceMsgVpnMqttSessionSubscriptionParams{
		HTTPClient: client,
	}
}

/*ReplaceMsgVpnMqttSessionSubscriptionParams contains all the parameters to send to the API endpoint
for the replace msg vpn mqtt session subscription operation typically these are written to a http.Request
*/
type ReplaceMsgVpnMqttSessionSubscriptionParams struct {

	/*Body
	  The Subscription object's attributes.

	*/
	Body *models.MsgVpnMqttSessionSubscription
	/*MqttSessionClientID
	  The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet.

	*/
	MqttSessionClientID string
	/*MqttSessionVirtualRouter
	  The virtual router of the MQTT Session.

	*/
	MqttSessionVirtualRouter string
	/*MsgVpnName
	  The name of the Message VPN.

	*/
	MsgVpnName string
	/*Select
	  Include in the response only selected attributes of the object, or exclude from the response selected attributes of the object. See the documentation for the `select` parameter.

	*/
	Select []string
	/*SubscriptionTopic
	  The MQTT subscription topic.

	*/
	SubscriptionTopic string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) WithTimeout(timeout time.Duration) *ReplaceMsgVpnMqttSessionSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) WithContext(ctx context.Context) *ReplaceMsgVpnMqttSessionSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) WithHTTPClient(client *http.Client) *ReplaceMsgVpnMqttSessionSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) WithBody(body *models.MsgVpnMqttSessionSubscription) *ReplaceMsgVpnMqttSessionSubscriptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) SetBody(body *models.MsgVpnMqttSessionSubscription) {
	o.Body = body
}

// WithMqttSessionClientID adds the mqttSessionClientID to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) WithMqttSessionClientID(mqttSessionClientID string) *ReplaceMsgVpnMqttSessionSubscriptionParams {
	o.SetMqttSessionClientID(mqttSessionClientID)
	return o
}

// SetMqttSessionClientID adds the mqttSessionClientId to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) SetMqttSessionClientID(mqttSessionClientID string) {
	o.MqttSessionClientID = mqttSessionClientID
}

// WithMqttSessionVirtualRouter adds the mqttSessionVirtualRouter to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) WithMqttSessionVirtualRouter(mqttSessionVirtualRouter string) *ReplaceMsgVpnMqttSessionSubscriptionParams {
	o.SetMqttSessionVirtualRouter(mqttSessionVirtualRouter)
	return o
}

// SetMqttSessionVirtualRouter adds the mqttSessionVirtualRouter to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) SetMqttSessionVirtualRouter(mqttSessionVirtualRouter string) {
	o.MqttSessionVirtualRouter = mqttSessionVirtualRouter
}

// WithMsgVpnName adds the msgVpnName to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) WithMsgVpnName(msgVpnName string) *ReplaceMsgVpnMqttSessionSubscriptionParams {
	o.SetMsgVpnName(msgVpnName)
	return o
}

// SetMsgVpnName adds the msgVpnName to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) SetMsgVpnName(msgVpnName string) {
	o.MsgVpnName = msgVpnName
}

// WithSelect adds the selectVar to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) WithSelect(selectVar []string) *ReplaceMsgVpnMqttSessionSubscriptionParams {
	o.SetSelect(selectVar)
	return o
}

// SetSelect adds the select to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) SetSelect(selectVar []string) {
	o.Select = selectVar
}

// WithSubscriptionTopic adds the subscriptionTopic to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) WithSubscriptionTopic(subscriptionTopic string) *ReplaceMsgVpnMqttSessionSubscriptionParams {
	o.SetSubscriptionTopic(subscriptionTopic)
	return o
}

// SetSubscriptionTopic adds the subscriptionTopic to the replace msg vpn mqtt session subscription params
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) SetSubscriptionTopic(subscriptionTopic string) {
	o.SubscriptionTopic = subscriptionTopic
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceMsgVpnMqttSessionSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param mqttSessionClientId
	if err := r.SetPathParam("mqttSessionClientId", o.MqttSessionClientID); err != nil {
		return err
	}

	// path param mqttSessionVirtualRouter
	if err := r.SetPathParam("mqttSessionVirtualRouter", o.MqttSessionVirtualRouter); err != nil {
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

	// path param subscriptionTopic
	if err := r.SetPathParam("subscriptionTopic", o.SubscriptionTopic); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
