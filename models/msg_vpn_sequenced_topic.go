// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MsgVpnSequencedTopic msg vpn sequenced topic
// swagger:model MsgVpnSequencedTopic
type MsgVpnSequencedTopic struct {

	// The name of the Message VPN.
	MsgVpnName string `json:"msgVpnName,omitempty"`

	// Topic for applying sequence numbers.
	SequencedTopic string `json:"sequencedTopic,omitempty"`
}

// Validate validates this msg vpn sequenced topic
func (m *MsgVpnSequencedTopic) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnSequencedTopic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnSequencedTopic) UnmarshalBinary(b []byte) error {
	var res MsgVpnSequencedTopic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}