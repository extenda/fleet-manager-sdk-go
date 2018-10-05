// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DriverPackages driver packages
// swagger:model DriverPackages
type DriverPackages struct {

	// next token
	NextToken string `json:"nextToken,omitempty"`

	// packages
	// Required: true
	Packages []*DriverPackagesPackagesItems0 `json:"packages"`
}

// Validate validates this driver packages
func (m *DriverPackages) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DriverPackages) validatePackages(formats strfmt.Registry) error {

	if err := validate.Required("packages", "body", m.Packages); err != nil {
		return err
	}

	for i := 0; i < len(m.Packages); i++ {
		if swag.IsZero(m.Packages[i]) { // not required
			continue
		}

		if m.Packages[i] != nil {
			if err := m.Packages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DriverPackages) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DriverPackages) UnmarshalBinary(b []byte) error {
	var res DriverPackages
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DriverPackagesPackagesItems0 driver packages packages items0
// swagger:model DriverPackagesPackagesItems0
type DriverPackagesPackagesItems0 struct {

	// href
	// Required: true
	Href *string `json:"href"`
}

// Validate validates this driver packages packages items0
func (m *DriverPackagesPackagesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DriverPackagesPackagesItems0) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DriverPackagesPackagesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DriverPackagesPackagesItems0) UnmarshalBinary(b []byte) error {
	var res DriverPackagesPackagesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}