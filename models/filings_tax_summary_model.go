// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FilingsTaxSummaryModel Represents a listing of all tax calculation data for filings and for accruing to future filings.
// Example: {"collectAmount":10,"nonTaxableAccrualAmount":100,"nonTaxableAmount":100,"remittanceAmount":10,"reportableNonTaxableAmount":50,"reportableSalesAmount":100,"reportableTaxAmount":5,"reportableTaxableAmount":5,"salesAccrualAmount":100,"salesAmount":200,"taxAccrualAmount":10,"taxAmount":10,"taxableAccrualAmount":0,"taxableAmount":0}
//
// swagger:model FilingsTaxSummaryModel
type FilingsTaxSummaryModel struct {

	// The collect amount
	// Example: 10
	CollectAmount float64 `json:"collectAmount,omitempty"`

	// The nontaxable accrual amount
	// Example: 100
	NonTaxableAccrualAmount float64 `json:"nonTaxableAccrualAmount,omitempty"`

	// The nontaxable amount
	// Example: 100
	NonTaxableAmount float64 `json:"nonTaxableAmount,omitempty"`

	// The remittance amount
	// Example: 10
	RemittanceAmount float64 `json:"remittanceAmount,omitempty"`

	// reportableNonTaxableAmount
	// Example: 50
	ReportableNonTaxableAmount float64 `json:"reportableNonTaxableAmount,omitempty"`

	// reportableSalesAmount
	// Example: 100
	ReportableSalesAmount float64 `json:"reportableSalesAmount,omitempty"`

	// reportableTaxAmount
	// Example: 5
	ReportableTaxAmount float64 `json:"reportableTaxAmount,omitempty"`

	// reportableTaxableAmount
	// Example: 5
	ReportableTaxableAmount float64 `json:"reportableTaxableAmount,omitempty"`

	// The sales accrual amount
	// Example: 100
	SalesAccrualAmount float64 `json:"salesAccrualAmount,omitempty"`

	// The total sales amount
	// Example: 200
	SalesAmount float64 `json:"salesAmount,omitempty"`

	// The tax accrual amount
	// Example: 10
	TaxAccrualAmount float64 `json:"taxAccrualAmount,omitempty"`

	// The tax amount
	// Example: 10
	TaxAmount float64 `json:"taxAmount,omitempty"`

	// The taxable sales accrual amount
	// Example: 0
	TaxableAccrualAmount float64 `json:"taxableAccrualAmount,omitempty"`

	// The taxable amount
	// Example: 0
	TaxableAmount float64 `json:"taxableAmount,omitempty"`
}

// Validate validates this filings tax summary model
func (m *FilingsTaxSummaryModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this filings tax summary model based on context it is used
func (m *FilingsTaxSummaryModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FilingsTaxSummaryModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilingsTaxSummaryModel) UnmarshalBinary(b []byte) error {
	var res FilingsTaxSummaryModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}