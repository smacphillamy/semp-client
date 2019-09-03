// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MsgVpnClientProfileLinks msg vpn client profile links
// swagger:model MsgVpnClientProfileLinks
type MsgVpnClientProfileLinks struct {

	// The URI of this Client Profile object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this msg vpn client profile links
func (m *MsgVpnClientProfileLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MsgVpnClientProfileLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsgVpnClientProfileLinks) UnmarshalBinary(b []byte) error {
	var res MsgVpnClientProfileLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
