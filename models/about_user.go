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

// AboutUser about user
// swagger:model AboutUser
type AboutUser struct {

	// The global access level of the User. The allowed values and their meaning are:
	//
	// <pre>
	// "admin" - Full administrative access.
	// "none" - No access.
	// "read-only" - Read only access.
	// "read-write" - Read and write access.
	// </pre>
	//
	// Enum: [admin none read-only read-write]
	GlobalAccessLevel string `json:"globalAccessLevel,omitempty"`
}

// Validate validates this about user
func (m *AboutUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGlobalAccessLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var aboutUserTypeGlobalAccessLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["admin","none","read-only","read-write"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aboutUserTypeGlobalAccessLevelPropEnum = append(aboutUserTypeGlobalAccessLevelPropEnum, v)
	}
}

const (

	// AboutUserGlobalAccessLevelAdmin captures enum value "admin"
	AboutUserGlobalAccessLevelAdmin string = "admin"

	// AboutUserGlobalAccessLevelNone captures enum value "none"
	AboutUserGlobalAccessLevelNone string = "none"

	// AboutUserGlobalAccessLevelReadOnly captures enum value "read-only"
	AboutUserGlobalAccessLevelReadOnly string = "read-only"

	// AboutUserGlobalAccessLevelReadWrite captures enum value "read-write"
	AboutUserGlobalAccessLevelReadWrite string = "read-write"
)

// prop value enum
func (m *AboutUser) validateGlobalAccessLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, aboutUserTypeGlobalAccessLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AboutUser) validateGlobalAccessLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.GlobalAccessLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateGlobalAccessLevelEnum("globalAccessLevel", "body", m.GlobalAccessLevel); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AboutUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AboutUser) UnmarshalBinary(b []byte) error {
	var res AboutUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
