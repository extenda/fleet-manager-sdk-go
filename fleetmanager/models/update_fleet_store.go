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

// UpdateFleetStore update fleet store
// swagger:model UpdateFleetStore
type UpdateFleetStore struct {

	// city
	// Required: true
	City *string `json:"city"`

	// contact person
	// Required: true
	ContactPerson *FleetStoreContactPerson `json:"contactPerson"`

	// name
	// Required: true
	Name *string `json:"name"`

	// post code
	// Required: true
	PostCode *string `json:"postCode"`

	// street
	// Required: true
	Street *string `json:"street"`
}

// Validate validates this update fleet store
func (m *UpdateFleetStore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactPerson(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateFleetStore) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	return nil
}

func (m *UpdateFleetStore) validateContactPerson(formats strfmt.Registry) error {

	if err := validate.Required("contactPerson", "body", m.ContactPerson); err != nil {
		return err
	}

	if m.ContactPerson != nil {
		if err := m.ContactPerson.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("contactPerson")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateFleetStore) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *UpdateFleetStore) validatePostCode(formats strfmt.Registry) error {

	if err := validate.Required("postCode", "body", m.PostCode); err != nil {
		return err
	}

	return nil
}

func (m *UpdateFleetStore) validateStreet(formats strfmt.Registry) error {

	if err := validate.Required("street", "body", m.Street); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateFleetStore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateFleetStore) UnmarshalBinary(b []byte) error {
	var res UpdateFleetStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
