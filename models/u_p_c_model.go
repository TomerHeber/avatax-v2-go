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

// UPCModel One Universal Product Code object as defined for your company.
// Example: {"companyId":1234567,"description":"Yarn","id":123456789,"legacyTaxCode":"PS081282","upc":"023032550992"}
//
// swagger:model UPCModel
type UPCModel struct {

	// The unique ID number of the company to which this UPC belongs.
	// Example: 1234567
	// Read Only: true
	CompanyID int32 `json:"companyId,omitempty"`

	// The date when this record was created.
	// Read Only: true
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// The User ID of the user who created this record.
	// Read Only: true
	CreatedUserID int32 `json:"createdUserId,omitempty"`

	// Description of the product to which this UPC applies.
	// Example: Yarn
	// Required: true
	// Max Length: 255
	// Min Length: 0
	Description *string `json:"description"`

	// If this UPC became effective on a certain date, this contains the first date on which the UPC was effective.
	// Format: date
	EffectiveDate strfmt.Date `json:"effectiveDate,omitempty"`

	// If this UPC expired or will expire on a certain date, this contains the last date on which the UPC was effective.
	// Format: date
	EndDate strfmt.Date `json:"endDate,omitempty"`

	// The unique ID number for this UPC.
	// Example: 123456789
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// A flag indicating whether this UPC code is attached to the AvaTax system or to a company.
	IsSystem int32 `json:"isSystem,omitempty"`

	// Legacy Tax Code applied to any product sold with this UPC.
	// Example: PS081282
	// Max Length: 50
	// Min Length: 0
	LegacyTaxCode *string `json:"legacyTaxCode,omitempty"`

	// The date/time when this record was last modified.
	// Read Only: true
	// Format: date-time
	ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

	// The user ID of the user who last modified this record.
	// Read Only: true
	ModifiedUserID int32 `json:"modifiedUserId,omitempty"`

	// The 12-14 character Universal Product Code, European Article Number, or Global Trade Identification Number.
	// Example: 023032550992
	// Required: true
	// Max Length: 18
	// Min Length: 12
	Upc *string `json:"upc"`

	// A usage identifier for this UPC code.
	Usage int32 `json:"usage,omitempty"`
}

// Validate validates this u p c model
func (m *UPCModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLegacyTaxCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UPCModel) validateCreatedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UPCModel) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", *m.Description, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", *m.Description, 255); err != nil {
		return err
	}

	return nil
}

func (m *UPCModel) validateEffectiveDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectiveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveDate", "body", "date", m.EffectiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UPCModel) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UPCModel) validateLegacyTaxCode(formats strfmt.Registry) error {
	if swag.IsZero(m.LegacyTaxCode) { // not required
		return nil
	}

	if err := validate.MinLength("legacyTaxCode", "body", *m.LegacyTaxCode, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("legacyTaxCode", "body", *m.LegacyTaxCode, 50); err != nil {
		return err
	}

	return nil
}

func (m *UPCModel) validateModifiedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UPCModel) validateUpc(formats strfmt.Registry) error {

	if err := validate.Required("upc", "body", m.Upc); err != nil {
		return err
	}

	if err := validate.MinLength("upc", "body", *m.Upc, 12); err != nil {
		return err
	}

	if err := validate.MaxLength("upc", "body", *m.Upc, 18); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this u p c model based on the context it is used
func (m *UPCModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCompanyID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UPCModel) contextValidateCompanyID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "companyId", "body", int32(m.CompanyID)); err != nil {
		return err
	}

	return nil
}

func (m *UPCModel) contextValidateCreatedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdDate", "body", strfmt.DateTime(m.CreatedDate)); err != nil {
		return err
	}

	return nil
}

func (m *UPCModel) contextValidateCreatedUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdUserId", "body", int32(m.CreatedUserID)); err != nil {
		return err
	}

	return nil
}

func (m *UPCModel) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *UPCModel) contextValidateModifiedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedDate", "body", strfmt.DateTime(m.ModifiedDate)); err != nil {
		return err
	}

	return nil
}

func (m *UPCModel) contextValidateModifiedUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedUserId", "body", int32(m.ModifiedUserID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UPCModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UPCModel) UnmarshalBinary(b []byte) error {
	var res UPCModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
