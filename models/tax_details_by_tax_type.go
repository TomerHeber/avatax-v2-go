// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TaxDetailsByTaxType Tax Details by Tax Type
// Example: {"taxType":"SalesAndUse","totalExempt":0.05,"totalNonTaxable":0,"totalTax":0.625,"totalTaxable":100}
//
// swagger:model TaxDetailsByTaxType
type TaxDetailsByTaxType struct {

	// Tax subtype details
	TaxSubTypeDetails []*TaxDetailsByTaxSubType `json:"taxSubTypeDetails"`

	// Tax Type
	// Example: SalesAndUse
	TaxType string `json:"taxType,omitempty"`

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

// Validate validates this tax details by tax type
func (m *TaxDetailsByTaxType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaxSubTypeDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxDetailsByTaxType) validateTaxSubTypeDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxSubTypeDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.TaxSubTypeDetails); i++ {
		if swag.IsZero(m.TaxSubTypeDetails[i]) { // not required
			continue
		}

		if m.TaxSubTypeDetails[i] != nil {
			if err := m.TaxSubTypeDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taxSubTypeDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taxSubTypeDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tax details by tax type based on the context it is used
func (m *TaxDetailsByTaxType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTaxSubTypeDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxDetailsByTaxType) contextValidateTaxSubTypeDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TaxSubTypeDetails); i++ {

		if m.TaxSubTypeDetails[i] != nil {
			if err := m.TaxSubTypeDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taxSubTypeDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taxSubTypeDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaxDetailsByTaxType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaxDetailsByTaxType) UnmarshalBinary(b []byte) error {
	var res TaxDetailsByTaxType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}