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

// ReportParametersModel The output model for report parameter definitions
// Example: {"country":"US","currencyCode":"USD","dateFilter":"DD","dateFormat":"MM/dd/yyyy","docType":"S","documentStatus":"Commited","endDate":"2021-12-25T00:00:00+00:00","isLocked":true,"isModifiedDateSameAsDocumentDate":false,"merchantSellerId":"abcd","numberOfPartitions":2,"partition":0,"startDate":"2021-11-25T00:00:00+00:00","state":"All","taxCode":"123","taxGroup":"Alcohol","taxName":"VAT","taxSubType":"Amusement"}
//
// swagger:model ReportParametersModel
type ReportParametersModel struct {

	// The country filter used for your report
	// Example: US
	Country string `json:"country,omitempty"`

	// The currency code used for your report
	// Example: USD
	CurrencyCode string `json:"currencyCode,omitempty"`

	// The code your business application uses to identify a customer or vendor
	CustomerVendorCode string `json:"customerVendorCode,omitempty"`

	// The date type filter used for your report
	// Example: DD
	DateFilter string `json:"dateFilter,omitempty"`

	// The date format used for your report
	// Example: MM/dd/yyyy
	DateFormat string `json:"dateFormat,omitempty"`

	// The doc type filter used for your report
	// Example: S
	DocType string `json:"docType,omitempty"`

	// The Document status filter used for report
	// For documentStatus, accepted values are: Temporary, Saved, Posted, Committed, Cancelled, Adjusted, Queued, PendingApproval
	// Example: Commited
	DocumentStatus string `json:"documentStatus,omitempty"`

	// The end date filter used for your report
	// Example: 2021-12-25T00:00:00+00:00
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// If true, include only documents that are locked.
	// If false, include only documents that are not locked.
	// Defaults to false if not specified.
	// Example: true
	IsLocked bool `json:"isLocked,omitempty"`

	// If true, modified date will be same as document date
	// If false, modified date will not be same as document date
	// Defaults to false if not specified.
	// Example: false
	IsModifiedDateSameAsDocumentDate bool `json:"isModifiedDateSameAsDocumentDate,omitempty"`

	// If set, include only documents associated with this merchantSellerId.
	// Example: abcd
	MerchantSellerID string `json:"merchantSellerId,omitempty"`

	// Number of partitions to split the report into.
	// Example: 2
	NumberOfPartitions int32 `json:"numberOfPartitions,omitempty"`

	// The zero-based partition number to retrieve in this export request.
	// Example: 0
	Partition int32 `json:"partition,omitempty"`

	// The start date filter used for your report
	// Example: 2021-11-25T00:00:00+00:00
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// The state filter used for your report
	// Example: All
	State string `json:"state,omitempty"`

	// The AvaTax tax code or customer tax code associated with the item or SKU in the transaction
	// Example: 123
	TaxCode string `json:"taxCode,omitempty"`

	// TaxGroup is required to support Sales tax (Sales + SellersUse) and VAT (Input+ Output).
	// TaxTypes, such as Lodging, Bottle, LandedCost, Ewaste, BevAlc, etc
	// Example: Alcohol
	TaxGroup string `json:"taxGroup,omitempty"`

	// The description of the tax
	// Example: VAT
	TaxName string `json:"taxName,omitempty"`

	// Defines the individual taxes associated with a TaxType category, such as Lodging TaxType which supports numerous TaxSubTypes, including Hotel, Occupancy, ConventionCenter, Accommotations, etc.
	// Example: Amusement
	TaxSubType string `json:"taxSubType,omitempty"`
}

// Validate validates this report parameters model
func (m *ReportParametersModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportParametersModel) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportParametersModel) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this report parameters model based on context it is used
func (m *ReportParametersModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReportParametersModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportParametersModel) UnmarshalBinary(b []byte) error {
	var res ReportParametersModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
