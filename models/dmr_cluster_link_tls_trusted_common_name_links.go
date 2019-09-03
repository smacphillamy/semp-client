// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DmrClusterLinkTLSTrustedCommonNameLinks dmr cluster link Tls trusted common name links
// swagger:model DmrClusterLinkTlsTrustedCommonNameLinks
type DmrClusterLinkTLSTrustedCommonNameLinks struct {

	// The URI of this Trusted Common Name object.
	URI string `json:"uri,omitempty"`
}

// Validate validates this dmr cluster link Tls trusted common name links
func (m *DmrClusterLinkTLSTrustedCommonNameLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DmrClusterLinkTLSTrustedCommonNameLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DmrClusterLinkTLSTrustedCommonNameLinks) UnmarshalBinary(b []byte) error {
	var res DmrClusterLinkTLSTrustedCommonNameLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
