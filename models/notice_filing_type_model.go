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

// NoticeFilingTypeModel Tax Notice FilingType Model
// Example: {"activeFlag":true,"description":"Electronic Return ","id":1,"sortOrder":5}
//
// swagger:model NoticeFilingTypeModel
type NoticeFilingTypeModel struct {

	// A flag if the type is active
	// Example: true
	ActiveFlag bool `json:"activeFlag,omitempty"`

	// The description name of this tax authority type.
	// Example: Electronic Return
	// Required: true
	Description *string `json:"description"`

	// The unique ID number of this tax notice customer type.
	// Example: 1
	// Required: true
	ID *int32 `json:"id"`

	// sort order of the types
	// Example: 5
	SortOrder int32 `json:"sortOrder,omitempty"`
}

// Validate validates this notice filing type model
func (m *NoticeFilingTypeModel) Validate(formats strfmt.Registry) error {
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

func (m *NoticeFilingTypeModel) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *NoticeFilingTypeModel) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this notice filing type model based on context it is used
func (m *NoticeFilingTypeModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NoticeFilingTypeModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NoticeFilingTypeModel) UnmarshalBinary(b []byte) error {
	var res NoticeFilingTypeModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}