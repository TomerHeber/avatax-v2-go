// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuditTransactionModel Information about a previously created transaction
// Example: {"apiCallStatus":"ReconstructedApiCallAvailable","companyId":0,"reconstructed":{"request":{"addresses":{"singleLocation":{"city":"Irvine","country":"US","line1":"2000 Main Street","postalCode":"92614","region":"CA"}},"commit":true,"companyCode":"DEFAULT","currencyCode":"USD","customerCode":"ABC","date":"2021-12-25","description":"Yarn","lines":[{"amount":100,"description":"Yarn","itemCode":"Y0001","number":"1","quantity":1,"taxCode":"PS081282"}],"purchaseOrderNo":"2021-12-25-001","type":"SalesInvoice"}},"serverTimestamp":"2021-12-25T17:08:47.5120263Z"}
//
// swagger:model AuditTransactionModel
type AuditTransactionModel struct {

	// api call status
	// Example: OriginalApiCallAvailable
	// Enum: [OriginalApiCallAvailable ReconstructedApiCallAvailable Any]
	APICallStatus string `json:"apiCallStatus,omitempty"`

	// Unique ID number of the company that created this transaction
	// Example: 0
	CompanyID int32 `json:"companyId,omitempty"`

	// Original API request/response
	// Example: {"request":{"addresses":{"singleLocation":{"city":"Irvine","country":"US","line1":"2000 Main Street","postalCode":"92614","region":"CA"}},"commit":true,"companyCode":"DEFAULT","currencyCode":"USD","customerCode":"ABC","date":"2021-12-25","description":"Yarn","lines":[{"amount":100,"description":"Yarn","itemCode":"Y0001","number":"1","quantity":1,"taxCode":"PS081282"}],"purchaseOrderNo":"2021-12-25-001","type":"SalesInvoice"},"response":{"addresses":[{"boundaryLevel":"Address","city":"Bainbridge Island","country":"US","id":0,"line1":"100 Ravine Lane Northeast #220","postalCode":"98110","region":"WA","taxRegionId":0,"transactionId":0}],"adjustmentDescription":"","adjustmentReason":"NotAdjusted","code":"64f00cfd-56a9-4025-becd-e54b2ef8053c","companyId":12345,"country":"US","currencyCode":"USD","customerCode":"ABC","customerVendorCode":"ABC","date":"2021-12-25","description":"Yarn","destinationAddressId":123456789,"entityUseCode":"","exchangeRate":2,"exchangeRateCurrencyCode":"USD","exchangeRateEffectiveDate":"2021-12-25","exemptNo":"","id":123456789,"isSellerImporterOfRecord":false,"lines":[{"boundaryOverrideId":0,"description":"Yarn","destinationAddressId":12345,"details":[{"addressId":12345,"country":"US","exemptAmount":0,"exemptReasonId":4,"exemptUnits":62.5,"id":123456789,"inState":false,"jurisCode":"06","jurisName":"CALIFORNIA","jurisType":"STA","jurisdictionId":5000531,"nonTaxableAmount":0,"nonTaxableRuleId":0,"nonTaxableType":"BaseRule","nonTaxableUnits":1000,"rate":0.0625,"rateRuleId":1321915,"rateSourceId":3,"rateType":"General","region":"CA","reportingExemptUnits":125,"reportingNonTaxableUnits":2000,"reportingTax":125,"reportingTaxCalculated":125,"reportingTaxableUnits":125,"serCode":"","signatureCode":"AGAM","sourcing":"Destination","stateAssignedNo":"","stateFIPS":"06","tax":62.5,"taxAuthorityTypeId":45,"taxCalculated":62.5,"taxName":"CA STATE TAX","taxOverride":0,"taxRegionId":2127184,"taxType":"Sales","taxableAmount":1000,"taxableUnits":62.5,"transactionId":123456789,"transactionLineId":123456789}],"discountAmount":100,"discountTypeId":0,"entityUseCode":"","exemptAmount":0,"exemptCertId":0,"exemptNo":"","id":123456789,"isItemTaxable":true,"isSSTP":false,"itemCode":"116292","lineAmount":1000,"lineNumber":"1","originAddressId":123456789,"quantity":1,"ref1":"Note: Deliver to Bob","reportingDate":"2021-12-25","revAccount":"","sourcing":"Destination","tax":62.5,"taxCalculated":62.5,"taxCode":"PS081282","taxDate":"2021-12-25","taxEngine":"","taxIncluded":false,"taxOverrideAmount":0,"taxOverrideReason":"","taxOverrideType":"None","taxableAmount":1000,"transactionId":123456789,"vatNumberTypeId":0}],"locationCode":"DEFAULT","locked":false,"originAddressId":123456789,"reconciled":true,"region":"CA","salespersonCode":"DEF","status":"Committed","taxDate":"2021-12-25T00:00:00+00:00","taxDetailsByTaxType":[{"taxType":"SalesAndUse","totalExempt":0.05,"totalNonTaxable":0,"totalTax":0.625,"totalTaxable":100}],"taxOverrideAmount":0,"taxOverrideReason":"","taxOverrideType":"None","totalAmount":1000,"totalDiscount":0,"totalExempt":0,"totalTax":62.5,"totalTaxCalculated":62.5,"totalTaxable":1000,"type":"SalesInvoice","version":0}}
	Original *OriginalAPIRequestResponseModel `json:"original,omitempty"`

	// Reconstructed API request/response
	// Example: {"request":{"addresses":{"singleLocation":{"city":"Irvine","country":"US","line1":"2000 Main Street","postalCode":"92614","region":"CA"}},"commit":true,"companyCode":"DEFAULT","currencyCode":"USD","customerCode":"ABC","date":"2021-12-25","description":"Yarn","lines":[{"amount":100,"description":"Yarn","itemCode":"Y0001","number":"1","quantity":1,"taxCode":"PS081282"}],"purchaseOrderNo":"2021-12-25-001","type":"SalesInvoice"}}
	Reconstructed *ReconstructedAPIRequestResponseModel `json:"reconstructed,omitempty"`

	// Length of time the original API call took
	// Format: date-time
	ServerDuration strfmt.DateTime `json:"serverDuration,omitempty"`

	// Server timestamp, in UTC, of the date/time when the original transaction was created
	// Example: 2021-12-25T17:08:47.5120263Z
	// Format: date-time
	ServerTimestamp strfmt.DateTime `json:"serverTimestamp,omitempty"`
}

// Validate validates this audit transaction model
func (m *AuditTransactionModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPICallStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReconstructed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var auditTransactionModelTypeAPICallStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OriginalApiCallAvailable","ReconstructedApiCallAvailable","Any"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		auditTransactionModelTypeAPICallStatusPropEnum = append(auditTransactionModelTypeAPICallStatusPropEnum, v)
	}
}

const (

	// AuditTransactionModelAPICallStatusOriginalAPICallAvailable captures enum value "OriginalApiCallAvailable"
	AuditTransactionModelAPICallStatusOriginalAPICallAvailable string = "OriginalApiCallAvailable"

	// AuditTransactionModelAPICallStatusReconstructedAPICallAvailable captures enum value "ReconstructedApiCallAvailable"
	AuditTransactionModelAPICallStatusReconstructedAPICallAvailable string = "ReconstructedApiCallAvailable"

	// AuditTransactionModelAPICallStatusAny captures enum value "Any"
	AuditTransactionModelAPICallStatusAny string = "Any"
)

// prop value enum
func (m *AuditTransactionModel) validateAPICallStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, auditTransactionModelTypeAPICallStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AuditTransactionModel) validateAPICallStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.APICallStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateAPICallStatusEnum("apiCallStatus", "body", m.APICallStatus); err != nil {
		return err
	}

	return nil
}

func (m *AuditTransactionModel) validateOriginal(formats strfmt.Registry) error {
	if swag.IsZero(m.Original) { // not required
		return nil
	}

	if m.Original != nil {
		if err := m.Original.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("original")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("original")
			}
			return err
		}
	}

	return nil
}

func (m *AuditTransactionModel) validateReconstructed(formats strfmt.Registry) error {
	if swag.IsZero(m.Reconstructed) { // not required
		return nil
	}

	if m.Reconstructed != nil {
		if err := m.Reconstructed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reconstructed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reconstructed")
			}
			return err
		}
	}

	return nil
}

func (m *AuditTransactionModel) validateServerDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerDuration) { // not required
		return nil
	}

	if err := validate.FormatOf("serverDuration", "body", "date-time", m.ServerDuration.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuditTransactionModel) validateServerTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.ServerTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("serverTimestamp", "body", "date-time", m.ServerTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this audit transaction model based on the context it is used
func (m *AuditTransactionModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOriginal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReconstructed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditTransactionModel) contextValidateOriginal(ctx context.Context, formats strfmt.Registry) error {

	if m.Original != nil {
		if err := m.Original.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("original")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("original")
			}
			return err
		}
	}

	return nil
}

func (m *AuditTransactionModel) contextValidateReconstructed(ctx context.Context, formats strfmt.Registry) error {

	if m.Reconstructed != nil {
		if err := m.Reconstructed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reconstructed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reconstructed")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditTransactionModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditTransactionModel) UnmarshalBinary(b []byte) error {
	var res AuditTransactionModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
