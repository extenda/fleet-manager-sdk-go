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

// FleetHardwareProfile fleet hardware profile
// swagger:model FleetHardwareProfile
type FleetHardwareProfile struct {

	// devices
	// Required: true
	Devices []*FleetHardwareProfileDevicesItems0 `json:"devices"`

	// id
	// Required: true
	ID *string `json:"id"`

	// jpos entries Xml Url
	JposEntriesXMLURL string `json:"jposEntriesXmlUrl,omitempty"`

	// jpos paths properties Url
	JposPathsPropertiesURL string `json:"jposPathsPropertiesUrl,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this fleet hardware profile
func (m *FleetHardwareProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FleetHardwareProfile) validateDevices(formats strfmt.Registry) error {

	if err := validate.Required("devices", "body", m.Devices); err != nil {
		return err
	}

	for i := 0; i < len(m.Devices); i++ {
		if swag.IsZero(m.Devices[i]) { // not required
			continue
		}

		if m.Devices[i] != nil {
			if err := m.Devices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FleetHardwareProfile) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *FleetHardwareProfile) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FleetHardwareProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FleetHardwareProfile) UnmarshalBinary(b []byte) error {
	var res FleetHardwareProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FleetHardwareProfileDevicesItems0 fleet hardware profile devices items0
// swagger:model FleetHardwareProfileDevicesItems0
type FleetHardwareProfileDevicesItems0 struct {

	// layer
	// Required: true
	Layer *string `json:"layer"`

	// manufacturer
	// Required: true
	Manufacturer *FleetHardwareProfileDevicesItems0Manufacturer `json:"manufacturer"`

	// model
	// Required: true
	Model *FleetHardwareProfileDevicesItems0Model `json:"model"`

	// port
	// Required: true
	Port *string `json:"port"`
}

// Validate validates this fleet hardware profile devices items0
func (m *FleetHardwareProfileDevicesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLayer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManufacturer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FleetHardwareProfileDevicesItems0) validateLayer(formats strfmt.Registry) error {

	if err := validate.Required("layer", "body", m.Layer); err != nil {
		return err
	}

	return nil
}

func (m *FleetHardwareProfileDevicesItems0) validateManufacturer(formats strfmt.Registry) error {

	if err := validate.Required("manufacturer", "body", m.Manufacturer); err != nil {
		return err
	}

	if m.Manufacturer != nil {
		if err := m.Manufacturer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manufacturer")
			}
			return err
		}
	}

	return nil
}

func (m *FleetHardwareProfileDevicesItems0) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	if m.Model != nil {
		if err := m.Model.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("model")
			}
			return err
		}
	}

	return nil
}

func (m *FleetHardwareProfileDevicesItems0) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FleetHardwareProfileDevicesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FleetHardwareProfileDevicesItems0) UnmarshalBinary(b []byte) error {
	var res FleetHardwareProfileDevicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FleetHardwareProfileDevicesItems0Manufacturer fleet hardware profile devices items0 manufacturer
// swagger:model FleetHardwareProfileDevicesItems0Manufacturer
type FleetHardwareProfileDevicesItems0Manufacturer struct {

	// href
	// Required: true
	Href *string `json:"href"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this fleet hardware profile devices items0 manufacturer
func (m *FleetHardwareProfileDevicesItems0Manufacturer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FleetHardwareProfileDevicesItems0Manufacturer) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("manufacturer"+"."+"href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

func (m *FleetHardwareProfileDevicesItems0Manufacturer) validateName(formats strfmt.Registry) error {

	if err := validate.Required("manufacturer"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FleetHardwareProfileDevicesItems0Manufacturer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FleetHardwareProfileDevicesItems0Manufacturer) UnmarshalBinary(b []byte) error {
	var res FleetHardwareProfileDevicesItems0Manufacturer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FleetHardwareProfileDevicesItems0Model fleet hardware profile devices items0 model
// swagger:model FleetHardwareProfileDevicesItems0Model
type FleetHardwareProfileDevicesItems0Model struct {

	// href
	// Required: true
	Href *string `json:"href"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this fleet hardware profile devices items0 model
func (m *FleetHardwareProfileDevicesItems0Model) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHref(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FleetHardwareProfileDevicesItems0Model) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("model"+"."+"href", "body", m.Href); err != nil {
		return err
	}

	return nil
}

func (m *FleetHardwareProfileDevicesItems0Model) validateName(formats strfmt.Registry) error {

	if err := validate.Required("model"+"."+"name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FleetHardwareProfileDevicesItems0Model) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FleetHardwareProfileDevicesItems0Model) UnmarshalBinary(b []byte) error {
	var res FleetHardwareProfileDevicesItems0Model
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
