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

// ComplianceTaxRateModel The tax rate model.
// Example: {"effectiveDate":"1900-01-01","endDate":"2011-04-30","id":1,"jurisdictionId":1,"rate":0.05,"rateTypeId":"F","rateTypeTaxTypeMappingId":46,"taxName":"AL STATE TAX","taxRegionId":4574,"taxTypeId":"S","unitOfBasisId":13}
//
// swagger:model ComplianceTaxRateModel
type ComplianceTaxRateModel struct {

	// The date this rate is starts to take effect.
	// Example: 1900-01-01T00:00:00
	// Read Only: true
	// Format: date
	EffectiveDate strfmt.Date `json:"effectiveDate,omitempty"`

	// The date this rate is no longer active.
	// Example: 2011-04-30T00:00:00
	// Read Only: true
	// Format: date
	EndDate strfmt.Date `json:"endDate,omitempty"`

	// The unique id of the rate.
	// Example: 1
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// The id of the jurisdiction.
	// Example: 1
	// Read Only: true
	JurisdictionID int32 `json:"jurisdictionId,omitempty"`

	// The tax rate.
	// Example: 0.05
	// Read Only: true
	Rate float64 `json:"rate,omitempty"`

	// The rate type.
	// Example: F
	// Read Only: true
	RateTypeID string `json:"rateTypeId,omitempty"`

	// The rate type tax type mapping id.
	// Example: 46
	// Read Only: true
	RateTypeTaxTypeMappingID int32 `json:"rateTypeTaxTypeMappingId,omitempty"`

	// The name of the tax.
	// Example: AL STATE TAX
	// Read Only: true
	TaxName string `json:"taxName,omitempty"`

	// The id of the tax region.
	// Example: 4574
	// Read Only: true
	TaxRegionID int32 `json:"taxRegionId,omitempty"`

	// The tax type.
	// Example: S
	// Read Only: true
	TaxTypeID string `json:"taxTypeId,omitempty"`

	// The unit of basis.
	// Example: 13
	// Read Only: true
	UnitOfBasisID int64 `json:"unitOfBasisId,omitempty"`
}

// Validate validates this compliance tax rate model
func (m *ComplianceTaxRateModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffectiveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComplianceTaxRateModel) validateEffectiveDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectiveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveDate", "body", "date", m.EffectiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceTaxRateModel) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this compliance tax rate model based on the context it is used
func (m *ComplianceTaxRateModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEffectiveDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJurisdictionID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRateTypeID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRateTypeTaxTypeMappingID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxRegionID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxTypeID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnitOfBasisID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComplianceTaxRateModel) contextValidateEffectiveDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "effectiveDate", "body", strfmt.Date(m.EffectiveDate)); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceTaxRateModel) contextValidateEndDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "endDate", "body", strfmt.Date(m.EndDate)); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceTaxRateModel) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceTaxRateModel) contextValidateJurisdictionID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "jurisdictionId", "body", int32(m.JurisdictionID)); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceTaxRateModel) contextValidateRate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rate", "body", float64(m.Rate)); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceTaxRateModel) contextValidateRateTypeID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rateTypeId", "body", string(m.RateTypeID)); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceTaxRateModel) contextValidateRateTypeTaxTypeMappingID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rateTypeTaxTypeMappingId", "body", int32(m.RateTypeTaxTypeMappingID)); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceTaxRateModel) contextValidateTaxName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "taxName", "body", string(m.TaxName)); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceTaxRateModel) contextValidateTaxRegionID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "taxRegionId", "body", int32(m.TaxRegionID)); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceTaxRateModel) contextValidateTaxTypeID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "taxTypeId", "body", string(m.TaxTypeID)); err != nil {
		return err
	}

	return nil
}

func (m *ComplianceTaxRateModel) contextValidateUnitOfBasisID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "unitOfBasisId", "body", int64(m.UnitOfBasisID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComplianceTaxRateModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComplianceTaxRateModel) UnmarshalBinary(b []byte) error {
	var res ComplianceTaxRateModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}