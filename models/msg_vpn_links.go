// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MsgVpnLinks msg vpn links
// swagger:model MsgVpnLinks
type MsgVpnLinks struct {

	// The URI of this Message VPN's collection of ACL Profile objects.
	ACLProfilesURI string `json:"aclProfilesUri,omitempty"`

	// The URI of this Message VPN's collection of LDAP Authorization Group objects.
	AuthorizationGroupsURI string `json:"authorizationGroupsUri,omitempty"`

	// The URI of this Message VPN's collection of Bridge objects.
	BridgesURI string `json:"bridgesUri,omitempty"`

	// The URI of this Message VPN's collection of Client Profile objects.
	ClientProfilesURI string `json:"clientProfilesUri,omitempty"`

	// The URI of this Message VPN's collection of Client Username objects.
	ClientUsernamesURI string `json:"clientUsernamesUri,omitempty"`

	// The URI of this Message VPN's collection of Distributed Cache objects. Available since 2.11.
	DistributedCachesURI string `json:"distributedCachesUri,omitempty"`

	// The URI of this Message VPN's collection of DMR Bridge objects. Available since 2.11.
	DmrBridgesURI string `json:"dmrBridgesUri,omitempty"`

	// The URI of this Message VPN's collection of JNDI Connection Factory objects. Available since 2.2.
	JndiConnectionFactoriesURI string `json:"jndiConnectionFactoriesUri,omitempty"`

	// The URI of this Message VPN's collection of JNDI Queue objects. Available since 2.2.
	JndiQueuesURI string `json:"jndiQueuesUri,omitempty"`

	// The URI of this Message VPN's collection of JNDI Topic objects. Available since 2.2.
	JndiTopicsURI string `json:"jndiTopicsUri,omitempty"`

	// The URI of this Message VPN's collection of MQTT Retain Cache objects. Available since 2.11.
	MqttRetainCachesURI string `json:"mqttRetainCachesUri,omitempty"`

	// The URI of this Message VPN's collection of MQTT Session objects. Available since 2.1.
	MqttSessionsURI string `json:"mqttSessionsUri,omitempty"`

	// The URI of this Message VPN's collection of Queue objects.
	QueuesURI string `json:"queuesUri,omitempty"`

	// The URI of this Message VPN's collection of Replay Log objects. Available since 2.10.
	ReplayLogsURI string `json:"replayLogsUri,omitempty"`

	// The URI of this Message VPN's collection of Replicated Topic objects. Available since 2.9.
	ReplicatedTopicsURI string `json:"replicatedTopicsUri,omitempty"`

	// The URI of this Message VPN's collection of REST Delivery Point objects.
	RestDeliveryPointsURI string `json:"restDeliveryPointsUri,omitempty"`

	// The URI of this Message VPN's collection of Sequenced Topic objects.
	SequencedTopicsURI string `json:"sequencedTopicsUri,omitempty"`

	// The URI of this Message VPN's collection of Topic Endpoint objects. Available since 2.1.
	TopicEndpointsURI string `json:"topicEndpointsUri,omitempty"`

	// The URI of this Message VPN object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn links
func (m *MsgVpnLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
