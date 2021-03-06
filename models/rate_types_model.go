// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RateTypesModel Rate types Model
// Example: {"description":"Alcohol","id":503,"rateType":"Alcohol"}
//
// swagger:model RateTypesModel
type RateTypesModel struct {

	// The description of this rate type.
	// Example: Alcohol
	Description string `json:"description,omitempty"`

	// The unique ID number of this rate type.
	// Example: 503
	ID int32 `json:"id,omitempty"`

	// The name of this rateType
	// Example: Alcohol
	RateType string `json:"rateType,omitempty"`
}

// Validate validates this rate types model
func (m *RateTypesModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rate types model based on context it is used
func (m *RateTypesModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RateTypesModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RateTypesModel) UnmarshalBinary(b []byte) error {
	var res RateTypesModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
