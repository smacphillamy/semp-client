// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UsernameMsgVpnAccessLevelExceptionResponse username msg vpn access level exception response
// swagger:model UsernameMsgVpnAccessLevelExceptionResponse
type UsernameMsgVpnAccessLevelExceptionResponse struct {

	// data
	Data *UsernameMsgVpnAccessLevelException `json:"data,omitempty"`

	// links
	Links *UsernameMsgVpnAccessLevelExceptionLinks `json:"links,omitempty"`

	// meta
	// Required: true
	Meta *SempMeta `json:"meta"`
}

// Validate validates this username msg vpn access level exception response
func (m *UsernameMsgVpnAccessLevelExceptionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UsernameMsgVpnAccessLevelExceptionResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *UsernameMsgVpnAccessLevelExceptionResponse) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *UsernameMsgVpnAccessLevelExceptionResponse) validateMeta(formats strfmt.Registry) error {

	if err := validate.Required("meta", "body", m.Meta); err != nil {
		return err
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UsernameMsgVpnAccessLevelExceptionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsernameMsgVpnAccessLevelExceptionResponse) UnmarshalBinary(b []byte) error {
	var res UsernameMsgVpnAccessLevelExceptionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
