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

// PostalCodeModel Represents a PostalCode and its associated data like: country, region, effective dates, etc.
// Example: {"country":"US","effDate":"2005-04-01T00:00:00","endDate":"9999-12-31T00:00:00","postalCode":"98104","region":"WA","taxRegionId":0}
//
// swagger:model PostalCodeModel
type PostalCodeModel struct {

	// Country this PostalCode locates in
	// Example: US
	// Read Only: true
	Country string `json:"country,omitempty"`

	// The date when the PostalCode becomes effective
	// Example: 2005-04-01T00:00:00
	// Read Only: true
	// Format: date-time
	EffDate strfmt.DateTime `json:"effDate,omitempty"`

	// The date when the PostalCode becomes expired
	// Example: 9999-12-31T00:00:00
	// Read Only: true
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// The postalCode
	// Example: 98104
	// Read Only: true
	PostalCode string `json:"postalCode,omitempty"`

	// The Region/State/Province this PostalCode locates in
	// Example: WA
	// Read Only: true
	Region string `json:"region,omitempty"`

	// An Avalara assigned TaxRegion Id associated to the PostalCode
	// Example: 0
	// Read Only: true
	TaxRegionID int32 `json:"taxRegionId,omitempty"`
}

// Validate validates this postal code model
func (m *PostalCodeModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffDate(formats); err != nil {
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

func (m *PostalCodeModel) validateEffDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EffDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effDate", "body", "date-time", m.EffDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PostalCodeModel) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this postal code model based on the context it is used
func (m *PostalCodeModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCountry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEffDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostalCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxRegionID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostalCodeModel) contextValidateCountry(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "country", "body", string(m.Country)); err != nil {
		return err
	}

	return nil
}

func (m *PostalCodeModel) contextValidateEffDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "effDate", "body", strfmt.DateTime(m.EffDate)); err != nil {
		return err
	}

	return nil
}

func (m *PostalCodeModel) contextValidateEndDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "endDate", "body", strfmt.DateTime(m.EndDate)); err != nil {
		return err
	}

	return nil
}

func (m *PostalCodeModel) contextValidatePostalCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "postalCode", "body", string(m.PostalCode)); err != nil {
		return err
	}

	return nil
}

func (m *PostalCodeModel) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "region", "body", string(m.Region)); err != nil {
		return err
	}

	return nil
}

func (m *PostalCodeModel) contextValidateTaxRegionID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "taxRegionId", "body", int32(m.TaxRegionID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostalCodeModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostalCodeModel) UnmarshalBinary(b []byte) error {
	var res PostalCodeModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
