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

// RebuildFilingsModel Rebuild a set of filings.
// Example: {"rebuild":true}
//
// swagger:model RebuildFilingsModel
type RebuildFilingsModel struct {

	// Set this value to true in order to rebuild the filings.
	// Example: true
	// Required: true
	Rebuild *bool `json:"rebuild"`
}

// Validate validates this rebuild filings model
func (m *RebuildFilingsModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRebuild(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RebuildFilingsModel) validateRebuild(formats strfmt.Registry) error {

	if err := validate.Required("rebuild", "body", m.Rebuild); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rebuild filings model based on context it is used
func (m *RebuildFilingsModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RebuildFilingsModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RebuildFilingsModel) UnmarshalBinary(b []byte) error {
	var res RebuildFilingsModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
