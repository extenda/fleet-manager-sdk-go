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

// S3PresignedPost s3 presigned post
// swagger:model S3PresignedPost
type S3PresignedPost struct {

	// fields
	// Required: true
	Fields *S3PresignedPostFields `json:"fields"`

	// url
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this s3 presigned post
func (m *S3PresignedPost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3PresignedPost) validateFields(formats strfmt.Registry) error {

	if err := validate.Required("fields", "body", m.Fields); err != nil {
		return err
	}

	if m.Fields != nil {
		if err := m.Fields.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fields")
			}
			return err
		}
	}

	return nil
}

func (m *S3PresignedPost) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3PresignedPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3PresignedPost) UnmarshalBinary(b []byte) error {
	var res S3PresignedPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// S3PresignedPostFields s3 presigned post fields
// swagger:model S3PresignedPostFields
type S3PresignedPostFields struct {

	// policy
	Policy string `json:"Policy,omitempty"`

	// x amz algorithm
	XAmzAlgorithm string `json:"X-Amz-Algorithm,omitempty"`

	// x amz credential
	XAmzCredential string `json:"X-Amz-Credential,omitempty"`

	// x amz date
	XAmzDate string `json:"X-Amz-Date,omitempty"`

	// x amz security token
	XAmzSecurityToken string `json:"X-Amz-Security-Token,omitempty"`

	// x amz signature
	XAmzSignature string `json:"X-Amz-Signature,omitempty"`

	// acl
	ACL string `json:"acl,omitempty"`

	// bucket
	// Required: true
	Bucket *string `json:"bucket"`

	// key
	// Required: true
	Key *string `json:"key"`
}

// Validate validates this s3 presigned post fields
func (m *S3PresignedPostFields) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3PresignedPostFields) validateBucket(formats strfmt.Registry) error {

	if err := validate.Required("fields"+"."+"bucket", "body", m.Bucket); err != nil {
		return err
	}

	return nil
}

func (m *S3PresignedPostFields) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("fields"+"."+"key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *S3PresignedPostFields) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3PresignedPostFields) UnmarshalBinary(b []byte) error {
	var res S3PresignedPostFields
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
