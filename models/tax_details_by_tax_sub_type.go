// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TaxDetailsByTaxSubType Tax Details by Tax subtype
// Example: {"taxSubType":"S","totalExempt":0.05,"totalNonTaxable":0,"totalTax":0.625,"totalTaxable":100}
//
// swagger:model TaxDetailsByTaxSubType
type TaxDetailsByTaxSubType struct {

	// Tax subtype
	// Example: S
	TaxSubType string `json:"taxSubType,omitempty"`

	// Total exempt by tax type
	// Example: 0.05
	TotalExempt float64 `json:"totalExempt,omitempty"`

	// Total non taxable by tax type
	// Example: 0
	TotalNonTaxable float64 `json:"totalNonTaxable,omitempty"`

	// Total tax by tax type
	// Example: 0.625
	TotalTax float64 `json:"totalTax,omitempty"`

	// Total taxable amount by tax type
	// Example: 100
	TotalTaxable float64 `json:"totalTaxable,omitempty"`
}

// Validate validates this tax details by tax sub type
func (m *TaxDetailsByTaxSubType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tax details by tax sub type based on context it is used
func (m *TaxDetailsByTaxSubType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaxDetailsByTaxSubType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaxDetailsByTaxSubType) UnmarshalBinary(b []byte) error {
	var res TaxDetailsByTaxSubType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}