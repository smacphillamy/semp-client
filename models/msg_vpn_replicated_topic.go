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

// MsgVpnReplicatedTopic msg vpn replicated topic
// swagger:model MsgVpnReplicatedTopic
type MsgVpnReplicatedTopic struct {

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// Topic for applying replication. Published messages matching this topic will be replicated to the standby site.
	ReplicatedTopic string `json:"replicatedTopic,omitempty"`

	// Choose the replication-mode for the Replicated Topic. The default value is `"async"`. The allowed values and their meaning are:
	//
	// <pre>
	// "sync" - Synchronous replication-mode. Published messages are acknowledged when they are spooled on the standby site.
	// "async" - Asynchronous replication-mode. Published messages are acknowledged when they are spooled locally.
	// </pre>
	//  Available since 2.1.
	// Enum: [sync async]
	ReplicationMode string `json:"replicationMode,omitempty"`
}

// Validate validates this msg vpn replicated topic
func (m *MsgVpnReplicatedTopic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReplicationMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var msgVpnReplicatedTopicTypeReplicationModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sync","async"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnReplicatedTopicTypeReplicationModePropEnum = append(msgVpnReplicatedTopicTypeReplicationModePropEnum, v)
	}
}

const (

	// MsgVpnReplicatedTopicReplicationModeSync captures enum value "sync"
	MsgVpnReplicatedTopicReplicationModeSync string = "sync"

	// MsgVpnReplicatedTopicReplicationModeAsync captures enum value "async"
	MsgVpnReplicatedTopicReplicationModeAsync string = "async"
)

// prop value enum
func (m *MsgVpnReplicatedTopic) validateReplicationModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, msgVpnReplicatedTopicTypeReplicationModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnReplicatedTopic) validateReplicationMode(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplicationMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateReplicationModeEnum("replicationMode", "body", m.ReplicationMode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnReplicatedTopic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnReplicatedTopic) UnmarshalBinary(b []byte) error {
	var res MsgVpnReplicatedTopic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}