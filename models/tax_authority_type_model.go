// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TaxAuthorityTypeModel Tax Authority Type Model
// Example: {"description":"County","id":0,"taxAuthorityGroup":"LocalCounty"}
//
// swagger:model TaxAuthorityTypeModel
type TaxAuthorityTypeModel struct {

	// The description name of this tax authority type.
	// Example: County
	// Required: true
	Description *string `json:"description"`

	// The unique ID number of this tax Authority customer type.
	// Example: 0
	// Required: true
	ID *int32 `json:"id"`

	// Tax Authority Group
	// Example: LocalCounty
	TaxAuthorityGroup string `json:"taxAuthorityGroup,omitempty"`
}

// Validate validates this tax authority type model
func (m *TaxAuthorityTypeModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxAuthorityTypeModel) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *TaxAuthorityTypeModel) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this tax authority type model based on context it is used
func (m *TaxAuthorityTypeModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaxAuthorityTypeModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaxAuthorityTypeModel) UnmarshalBinary(b []byte) error {
	var res TaxAuthorityTypeModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
