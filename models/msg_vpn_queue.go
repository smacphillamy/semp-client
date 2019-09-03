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

// MsgVpnQueue msg vpn queue
// swagger:model MsgVpnQueue
type MsgVpnQueue struct {

	// The access type for delivering messages to consumer flows bound to the Queue. The default value is `"exclusive"`. The allowed values and their meaning are:
	//
	// <pre>
	// "exclusive" - Exclusive delivery of messages to the first bound consumer flow.
	// "non-exclusive" - Non-exclusive delivery of messages to all bound consumer flows in a round-robin fashion.
	// </pre>
	//
	// Enum: [exclusive non-exclusive]
	AccessType string `json:"accessType,omitempty"`

	// Enable or disable the propagation of consumer acknowledgements (ACKs) received on the active replication Message VPN to the standby replication Message VPN. The default value is `true`.
	ConsumerAckPropagationEnabled bool `json:"consumerAckPropagationEnabled,omitempty"`

	// The name of the Dead Message Queue (DMQ) used by the Queue. The default value is `"#DEAD_MSG_QUEUE"`. Available since 2.2.
	DeadMsgQueue string `json:"deadMsgQueue,omitempty"`

	// Enable or disable the transmission of messages from the Queue. The default value is `false`.
	EgressEnabled bool `json:"egressEnabled,omitempty"`

	// event bind count threshold
	EventBindCountThreshold *EventThreshold `json:"eventBindCountThreshold,omitempty"`

	// event msg spool usage threshold
	EventMsgSpoolUsageThreshold *EventThreshold `json:"eventMsgSpoolUsageThreshold,omitempty"`

	// event reject low priority msg limit threshold
	EventRejectLowPriorityMsgLimitThreshold *EventThreshold `json:"eventRejectLowPriorityMsgLimitThreshold,omitempty"`

	// Enable or disable the reception of messages to the Queue. The default value is `false`.
	IngressEnabled bool `json:"ingressEnabled,omitempty"`

	// The maximum number of consumer flows that can bind to the Queue. The default value is `1000`.
	MaxBindCount int64 `json:"maxBindCount,omitempty"`

	// The maximum number of messages delivered but not acknowledged per flow for the Queue. The default is the max value supported by the platform.
	MaxDeliveredUnackedMsgsPerFlow int64 `json:"maxDeliveredUnackedMsgsPerFlow,omitempty"`

	// The maximum message size allowed in the Queue, in bytes (B). The default value is `10000000`.
	MaxMsgSize int32 `json:"maxMsgSize,omitempty"`

	// The maximum message spool usage allowed by the Queue, in megabytes (MB). A value of 0 only allows spooling of the last message received and disables quota checking. The default varies by platform.
	MaxMsgSpoolUsage int64 `json:"maxMsgSpoolUsage,omitempty"`

	// The maximum number of times the Queue will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. The default value is `0`.
	MaxRedeliveryCount int64 `json:"maxRedeliveryCount,omitempty"`

	// The maximum time in seconds a message can stay in the Queue when `respectTtlEnabled` is `"true"`. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the `maxTtl` configured for the Queue, is exceeded. A value of 0 disables expiry. The default value is `0`.
	MaxTTL int64 `json:"maxTtl,omitempty"`

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// The Client Username that owns the Queue and has permission equivalent to `"delete"`. The default value is `""`.
	Owner string `json:"owner,omitempty"`

	// The permission level for all consumers of the Queue, excluding the owner. The default value is `"no-access"`. The allowed values and their meaning are:
	//
	// <pre>
	// "no-access" - Disallows all access.
	// "read-only" - Read-only access to the messages.
	// "consume" - Consume (read and remove) messages.
	// "modify-topic" - Consume messages or modify the topic/selector.
	// "delete" - Consume messages, modify the topic/selector or delete the Client created endpoint altogether.
	// </pre>
	//
	// Enum: [no-access read-only consume modify-topic delete]
	Permission string `json:"permission,omitempty"`

	// The name of the Queue.
	QueueName string `json:"queueName,omitempty"`

	// Enable or disable the checking of low priority messages against the `rejectLowPriorityMsgLimit`. This may only be enabled if `rejectMsgToSenderOnDiscardBehavior` does not have a value of `"never"`. The default value is `false`.
	RejectLowPriorityMsgEnabled bool `json:"rejectLowPriorityMsgEnabled,omitempty"`

	// The number of messages of any priority in the Queue above which low priority messages are not admitted but higher priority messages are allowed. The default value is `0`.
	RejectLowPriorityMsgLimit int64 `json:"rejectLowPriorityMsgLimit,omitempty"`

	// Determines when to return negative acknowledgements (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. The default value is `"when-queue-enabled"`. The allowed values and their meaning are:
	//
	// <pre>
	// "always" - Always return a negative acknowledgment (NACK) to the sending client on message discard.
	// "when-queue-enabled" - Only return a negative acknowledgment (NACK) to the sending client on message discard when the Queue is enabled.
	// "never" - Never return a negative acknowledgment (NACK) to the sending client on message discard.
	// </pre>
	//  Available since 2.1.
	// Enum: [always when-queue-enabled never]
	RejectMsgToSenderOnDiscardBehavior string `json:"rejectMsgToSenderOnDiscardBehavior,omitempty"`

	// Enable or disable the respecting of message priority. When enabled, messages contained in the Queue are delivered in priority order, from 9 (highest) to 0 (lowest). MQTT queues do not support enabling message priority. The default value is `false`. Available since 2.8.
	RespectMsgPriorityEnabled bool `json:"respectMsgPriorityEnabled,omitempty"`

	// Enable or disable the respecting of the time-to-live (TTL) for messages in the Queue. When enabled, expired messages are discarded or moved to the DMQ. The default value is `false`.
	RespectTTLEnabled bool `json:"respectTtlEnabled,omitempty"`
}

// Validate validates this msg vpn queue
func (m *MsgVpnQueue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventBindCountThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventMsgSpoolUsageThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventRejectLowPriorityMsgLimitThreshold(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRejectMsgToSenderOnDiscardBehavior(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var msgVpnQueueTypeAccessTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["exclusive","non-exclusive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnQueueTypeAccessTypePropEnum = append(msgVpnQueueTypeAccessTypePropEnum, v)
	}
}

const (

	// MsgVpnQueueAccessTypeExclusive captures enum value "exclusive"
	MsgVpnQueueAccessTypeExclusive string = "exclusive"

	// MsgVpnQueueAccessTypeNonExclusive captures enum value "non-exclusive"
	MsgVpnQueueAccessTypeNonExclusive string = "non-exclusive"
)

// prop value enum
func (m *MsgVpnQueue) validateAccessTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, msgVpnQueueTypeAccessTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnQueue) validateAccessType(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessType) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessTypeEnum("accessType", "body", m.AccessType); err != nil {
		return err
	}

	return nil
}

func (m *MsgVpnQueue) validateEventBindCountThreshold(formats strfmt.Registry) error {

	if swag.IsZero(m.EventBindCountThreshold) { // not required
		return nil
	}

	if m.EventBindCountThreshold != nil {
		if err := m.EventBindCountThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventBindCountThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnQueue) validateEventMsgSpoolUsageThreshold(formats strfmt.Registry) error {

	if swag.IsZero(m.EventMsgSpoolUsageThreshold) { // not required
		return nil
	}

	if m.EventMsgSpoolUsageThreshold != nil {
		if err := m.EventMsgSpoolUsageThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventMsgSpoolUsageThreshold")
			}
			return err
		}
	}

	return nil
}

func (m *MsgVpnQueue) validateEventRejectLowPriorityMsgLimitThreshold(formats strfmt.Registry) error {

	if swag.IsZero(m.EventRejectLowPriorityMsgLimitThreshold) { // not required
		return nil
	}

	if m.EventRejectLowPriorityMsgLimitThreshold != nil {
		if err := m.EventRejectLowPriorityMsgLimitThreshold.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventRejectLowPriorityMsgLimitThreshold")
			}
			return err
		}
	}

	return nil
}

var msgVpnQueueTypePermissionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["no-access","read-only","consume","modify-topic","delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnQueueTypePermissionPropEnum = append(msgVpnQueueTypePermissionPropEnum, v)
	}
}

const (

	// MsgVpnQueuePermissionNoAccess captures enum value "no-access"
	MsgVpnQueuePermissionNoAccess string = "no-access"

	// MsgVpnQueuePermissionReadOnly captures enum value "read-only"
	MsgVpnQueuePermissionReadOnly string = "read-only"

	// MsgVpnQueuePermissionConsume captures enum value "consume"
	MsgVpnQueuePermissionConsume string = "consume"

	// MsgVpnQueuePermissionModifyTopic captures enum value "modify-topic"
	MsgVpnQueuePermissionModifyTopic string = "modify-topic"

	// MsgVpnQueuePermissionDelete captures enum value "delete"
	MsgVpnQueuePermissionDelete string = "delete"
)

// prop value enum
func (m *MsgVpnQueue) validatePermissionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, msgVpnQueueTypePermissionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnQueue) validatePermission(formats strfmt.Registry) error {

	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	// value enum
	if err := m.validatePermissionEnum("permission", "body", m.Permission); err != nil {
		return err
	}

	return nil
}

var msgVpnQueueTypeRejectMsgToSenderOnDiscardBehaviorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["always","when-queue-enabled","never"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		msgVpnQueueTypeRejectMsgToSenderOnDiscardBehaviorPropEnum = append(msgVpnQueueTypeRejectMsgToSenderOnDiscardBehaviorPropEnum, v)
	}
}

const (

	// MsgVpnQueueRejectMsgToSenderOnDiscardBehaviorAlways captures enum value "always"
	MsgVpnQueueRejectMsgToSenderOnDiscardBehaviorAlways string = "always"

	// MsgVpnQueueRejectMsgToSenderOnDiscardBehaviorWhenQueueEnabled captures enum value "when-queue-enabled"
	MsgVpnQueueRejectMsgToSenderOnDiscardBehaviorWhenQueueEnabled string = "when-queue-enabled"

	// MsgVpnQueueRejectMsgToSenderOnDiscardBehaviorNever captures enum value "never"
	MsgVpnQueueRejectMsgToSenderOnDiscardBehaviorNever string = "never"
)

// prop value enum
func (m *MsgVpnQueue) validateRejectMsgToSenderOnDiscardBehaviorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, msgVpnQueueTypeRejectMsgToSenderOnDiscardBehaviorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MsgVpnQueue) validateRejectMsgToSenderOnDiscardBehavior(formats strfmt.Registry) error {

	if swag.IsZero(m.RejectMsgToSenderOnDiscardBehavior) { // not required
		return nil
	}

	// value enum
	if err := m.validateRejectMsgToSenderOnDiscardBehaviorEnum("rejectMsgToSenderOnDiscardBehavior", "body", m.RejectMsgToSenderOnDiscardBehavior); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnQueue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnQueue) UnmarshalBinary(b []byte) error {
	var res MsgVpnQueue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
