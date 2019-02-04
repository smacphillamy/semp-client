// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MsgVpnRestDeliveryPointRestConsumer msg vpn rest delivery point rest consumer
// swagger:model MsgVpnRestDeliveryPointRestConsumer
type MsgVpnRestDeliveryPointRestConsumer struct {

	// The PEM formatted content for the client certificate that the REST Consumer will present to the REST host. It must consist of a private key and between one and three certificates comprising the certificate trust chain. The default value is `""`. Available since 2.9.
	AuthenticationClientCertContent string `json:"authenticationClientCertContent,omitempty"`

	// The password for the client certificate that the REST Consumer will present to the REST host. The default value is `""`. Available since 2.9.
	AuthenticationClientCertPassword string `json:"authenticationClientCertPassword,omitempty"`

	// The password that the REST Consumer will use to login to the REST host. The default value is `""`.
	AuthenticationHTTPBasicPassword string `json:"authenticationHttpBasicPassword,omitempty"`

	// The username that the REST Consumer will use to login to the REST host. Normally a username is only configured when basic authentication is selected for the REST Consumer. The default value is `""`.
	AuthenticationHTTPBasicUsername string `json:"authenticationHttpBasicUsername,omitempty"`

	// The authentication scheme used by the REST Consumer to login to the REST host. The default value is `"none"`. The allowed values and their meaning are:
	//
	// <pre>
	// "none" - Login with no authentication. This may be useful for anonymous connections or when a REST Consumer does not require authentication.
	// "http-basic" - Login with a username and optional password according to HTTP Basic authentication as per RFC2616.
	// "client-certificate" - Login with a client TLS certificate as per RFC5246. Client certificate authentication is only available on TLS connections.
	// </pre>
	//
	// Enum: [none http-basic client-certificate]
	AuthenticationScheme string `json:"authenticationScheme,omitempty"`

	// Enable or disable the REST Consumer. When disabled, no connections are initiated or messages delivered to this particular REST Consumer. The default value is `false`.
	Enabled bool `json:"enabled,omitempty"`

	// The interface that will be used for all outgoing connections associated with the REST Consumer. When unspecified, an interface is automatically chosen. The default value is `""`.
	LocalInterface string `json:"localInterface,omitempty"`

	// The maximum amount of time (in seconds) to wait for an HTTP POST response from the REST Consumer. Once this time is exceeded, the TCP connection is reset. If the POST request is for a direct message, then the message is discarded. If for a persistent message, then message redelivery is attempted on another available outgoing connection for the REST Delivery Point. The default value is `30`.
	MaxPostWaitTime int32 `json:"maxPostWaitTime,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The total number of concurrent TCP connections open to the REST Consumer. Multiple connections to a single REST Consumer increase throughput via concurrency. The default value is `3`.
	OutgoingConnectionCount int32 `json:"outgoingConnectionCount,omitempty"`

	// The IP address or DNS name to which the router is to connect to deliver messages for this REST Consumer. If the REST Consumer is enabled while the host value is not configured then the REST Consumer has an operational Down state due to the empty host configuration until a usable host value is configured. The default value is `""`.
	RemoteHost string `json:"remoteHost,omitempty"`

	// The port associated with the host of the REST Consumer. The port can only be changed when the REST Consumer is disabled. The default value is `8080`.
	RemotePort int64 `json:"remotePort,omitempty"`

	// The name of the REST Consumer.
	RestConsumerName string `json:"restConsumerName,omitempty"`

	// The name of the REST Delivery Point.
	RestDeliveryPointName string `json:"restDeliveryPointName,omitempty"`

	// The number of seconds that must pass before retrying the remote REST Consumer connection. The default value is `3`.
	RetryDelay int32 `json:"retryDelay,omitempty"`

	// The colon-separated list of cipher-suites the REST Consumer uses in its encrypted connection. All supported suites are included by default, from most-secure to least-secure. The REST Consumer should choose the first suite from this list that it supports. The cipher-suite list can only be changed when the REST Consumer is disabled. The default value is `"default"`.
	TLSCipherSuiteList string `json:"tlsCipherSuiteList,omitempty"`

	// Enable or disable TLS for the REST Consumer. This may only be done when the REST Consumer is disabled. The default value is `false`.
	TLSEnabled bool `json:"tlsEnabled,omitempty"`
}

// Validate validates this msg vpn rest delivery point rest consumer
func (m *MsgVpnRestDeliveryPointRestConsumer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticationScheme(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var msgVpnRestDeliveryPointRestConsumerTypeAuthenticationSchemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","http-basic","client-certificate"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnRestDeliveryPointRestConsumerTypeAuthenticationSchemePropEnum = append(msgVpnRestDeliveryPointRestConsumerTypeAuthenticationSchemePropEnum, v)
	}
}

const (

	// MsgVpnRestDeliveryPointRestConsumerAuthenticationSchemeNone captures enum value "none"
	MsgVpnRestDeliveryPointRestConsumerAuthenticationSchemeNone string = "none"

	// MsgVpnRestDeliveryPointRestConsumerAuthenticationSchemeHTTPBasic captures enum value "http-basic"
	MsgVpnRestDeliveryPointRestConsumerAuthenticationSchemeHTTPBasic string = "http-basic"

	// MsgVpnRestDeliveryPointRestConsumerAuthenticationSchemeClientCertificate captures enum value "client-certificate"
	MsgVpnRestDeliveryPointRestConsumerAuthenticationSchemeClientCertificate string = "client-certificate"
)

// prop value enum
func (m *MsgVpnRestDeliveryPointRestConsumer) validateAuthenticationSchemeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, msgVpnRestDeliveryPointRestConsumerTypeAuthenticationSchemePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnRestDeliveryPointRestConsumer) validateAuthenticationScheme(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthenticationScheme) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationSchemeEnum("authenticationScheme", "body", m.AuthenticationScheme); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnRestDeliveryPointRestConsumer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnRestDeliveryPointRestConsumer) UnmarshalBinary(b []byte) error {
	var res MsgVpnRestDeliveryPointRestConsumer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
