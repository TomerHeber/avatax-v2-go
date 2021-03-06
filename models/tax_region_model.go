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
	"github.com/go-openapi/validate"
)

// TaxRegionModel The tax region model.
// Example: {"code":"2000000000","country":"US","effectiveDate":"1900-01-01","endDate":"2013-11-30","id":120,"isAcm":false,"isSst":false,"jurisdictions":[{"effectiveDate":"1900-01-01T00:00:00","endDate":"2013-11-30T00:00:00","isAcm":true,"isSst":true,"jurisCode":"20","jurisName":"KANSAS","jurisType":"State","jurisdictionId":20,"taxAuthorityId":280}],"name":"KANSAS","region":"KS","signatureCode":"AUIO"}
//
// swagger:model TaxRegionModel
type TaxRegionModel struct {

	// The name of the city.
	// Read Only: true
	City string `json:"city,omitempty"`

	// The code of the tax region.
	// Example: 2000000000
	// Read Only: true
	Code string `json:"code,omitempty"`

	// Name or ISO 3166 code identifying the country of this jurisdiction.
	//
	// This field supports many different country identifiers:
	//  * Two character ISO 3166 codes
	//  * Three character ISO 3166 codes
	//  * Fully spelled out names of the country in ISO supported languages
	//  * Common alternative spellings for many countries
	//
	// For a full list of all supported codes and names, please see the Definitions API `ListCountries`.
	// Example: US
	// Read Only: true
	Country string `json:"country,omitempty"`

	// The name of the county.
	// Read Only: true
	County string `json:"county,omitempty"`

	// The date this tax region starts to take effect.
	// Example: 1900-01-01T00:00:00
	// Read Only: true
	// Format: date
	EffectiveDate strfmt.Date `json:"effectiveDate,omitempty"`

	// The date this tax region stops to take effect.
	// Example: 2013-11-30T00:00:00
	// Read Only: true
	// Format: date
	EndDate strfmt.Date `json:"endDate,omitempty"`

	// The id of the tax region.
	// Example: 120
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// Is Acm flag.
	// Example: false
	// Read Only: true
	IsAcm *bool `json:"isAcm,omitempty"`

	// Is SST flag.
	// Example: false
	// Read Only: true
	IsSst *bool `json:"isSst,omitempty"`

	// List of jurisdictions associated with this tax region.
	// Example: [{"effectiveDate":"1900-01-01T00:00:00","endDate":"2013-11-30T00:00:00","isAcm":true,"isSst":true,"jurisCode":"20","jurisName":"KANSAS","jurisType":"State","jurisdictionId":20,"taxAuthorityId":280}]
	Jurisdictions []*DenormalizedJurisModel `json:"jurisdictions"`

	// The name of the tax region.
	// Example: KANSAS
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Name or ISO 3166 code identifying the region within the country.
	//
	// This field supports many different region identifiers:
	//  * Two and three character ISO 3166 region codes
	//  * Fully spelled out names of the region in ISO supported languages
	//  * Common alternative spellings for many regions
	//
	// For a full list of all supported codes and names, please see the Definitions API `ListRegions`.
	// Example: KS
	// Read Only: true
	Region string `json:"region,omitempty"`

	// The ser code.
	// Read Only: true
	SerCode string `json:"serCode,omitempty"`

	// The tax region signature code.
	// Example: AUIO
	// Read Only: true
	SignatureCode string `json:"signatureCode,omitempty"`
}

// Validate validates this tax region model
func (m *TaxRegionModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffectiveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJurisdictions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxRegionModel) validateEffectiveDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectiveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveDate", "body", "date", m.EffectiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaxRegionModel) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TaxRegionModel) validateJurisdictions(formats strfmt.Registry) error {
	if swag.IsZero(m.Jurisdictions) { // not required
		return nil
	}

	for i := 0; i < len(m.Jurisdictions); i++ {
		if swag.IsZero(m.Jurisdictions[i]) { // not required
			continue
		}

		if m.Jurisdictions[i] != nil {
			if err := m.Jurisdictions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jurisdictions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("jurisdictions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tax region model based on the context it is used
func (m *TaxRegionModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCounty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEffectiveDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsAcm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIsSst(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJurisdictions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSerCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignatureCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaxRegionModel) contextValidateCity(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "city", "body", string(m.City)); err != nil {
		return err
	}

	return nil
}

func (m *TaxRegionModel) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "code", "body", string(m.Code)); err != nil {
		return err
	}

	return nil
}

func (m *TaxRegionModel) contextValidateCountry(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "country", "body", string(m.Country)); err != nil {
		return err
	}

	return nil
}

func (m *TaxRegionModel) contextValidateCounty(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "county", "body", string(m.County)); err != nil {
		return err
	}

	return nil
}

func (m *TaxRegionModel) contextValidateEffectiveDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "effectiveDate", "body", strfmt.Date(m.EffectiveDate)); err != nil {
		return err
	}

	return nil
}

func (m *TaxRegionModel) contextValidateEndDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "endDate", "body", strfmt.Date(m.EndDate)); err != nil {
		return err
	}

	return nil
}

func (m *TaxRegionModel) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *TaxRegionModel) contextValidateIsAcm(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "isAcm", "body", m.IsAcm); err != nil {
		return err
	}

	return nil
}

func (m *TaxRegionModel) contextValidateIsSst(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "isSst", "body", m.IsSst); err != nil {
		return err
	}

	return nil
}

func (m *TaxRegionModel) contextValidateJurisdictions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Jurisdictions); i++ {

		if m.Jurisdictions[i] != nil {
			if err := m.Jurisdictions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jurisdictions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("jurisdictions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TaxRegionModel) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *TaxRegionModel) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "region", "body", string(m.Region)); err != nil {
		return err
	}

	return nil
}

func (m *TaxRegionModel) contextValidateSerCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "serCode", "body", string(m.SerCode)); err != nil {
		return err
	}

	return nil
}

func (m *TaxRegionModel) contextValidateSignatureCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "signatureCode", "body", string(m.SignatureCode)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaxRegionModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaxRegionModel) UnmarshalBinary(b []byte) error {
	var res TaxRegionModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
