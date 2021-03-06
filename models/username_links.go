// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UsernameLinks username links
// swagger:model UsernameLinks
type UsernameLinks struct {

	// The URI of this Username's collection of Message VPN Access Level Exception objects.
	MsgVpnAccessLevelExceptionsURI string `json:"msgVpnAccessLevelExceptionsUri,omitempty"`

	// The URI of this Username object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this username links
func (m *UsernameLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UsernameLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsernameLinks) UnmarshalBinary(b []byte) error {
	var res UsernameLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
