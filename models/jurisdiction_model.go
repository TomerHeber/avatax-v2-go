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

// JurisdictionModel Represents information about a single legal taxing jurisdiction
// Example: {"code":"53","name":"WASHINGTON","rate":0.06,"region":"WA","salesRate":0.06,"signatureCode":"","type":"State","useRate":0.06}
//
// swagger:model JurisdictionModel
type JurisdictionModel struct {

	// The city name of this jurisdiction
	City string `json:"city,omitempty"`

	// The code that is used to identify this jurisdiction
	// Example: 53
	// Required: true
	Code *string `json:"code"`

	// Name or ISO 3166 code identifying the country of this jurisdiction.
	//
	// This field supports many different country identifiers:
	//  * Two character ISO 3166 codes
	//  * Three character ISO 3166 codes
	//  * Fully spelled out names of the country in ISO supported languages
	//  * Common alternative spellings for many countries
	//
	// For a full list of all supported codes and names, please see the Definitions API `ListCountries`.
	Country string `json:"country,omitempty"`

	// The county name of this jurisdiction
	County string `json:"county,omitempty"`

	// County FIPS code
	CountyFips string `json:"countyFips,omitempty"`

	// The date this jurisdiction starts to take effect on tax calculations
	// Format: date-time
	EffectiveDate strfmt.DateTime `json:"effectiveDate,omitempty"`

	// The date this jurisdiction stops to take effect on tax calculations
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// Unique AvaTax Id of this Jurisdiction
	ID int32 `json:"id,omitempty"`

	// The name of this jurisdiction
	// Example: WASHINGTON
	// Required: true
	Name *string `json:"name"`

	// City FIPS code
	PlaceFips string `json:"placeFips,omitempty"`

	// The base rate of tax specific to this jurisdiction.
	// Example: 0.06
	Rate float64 `json:"rate,omitempty"`

	// Name or ISO 3166 code identifying the region within the country.
	//
	// This field supports many different region identifiers:
	//  * Two and three character ISO 3166 region codes
	//  * Fully spelled out names of the region in ISO supported languages
	//  * Common alternative spellings for many regions
	//
	// For a full list of all supported codes and names, please see the Definitions API `ListRegions`.
	// Example: WA
	Region string `json:"region,omitempty"`

	// The "Sales" tax rate specific to this jurisdiction.
	// Example: 0.06
	SalesRate float64 `json:"salesRate,omitempty"`

	// A short name of the jurisidiction
	ShortName string `json:"shortName,omitempty"`

	// The Avalara-supplied signature code for this jurisdiction.
	// Required: true
	SignatureCode *string `json:"signatureCode"`

	// State FIPS code
	StateFips string `json:"stateFips,omitempty"`

	// The type of the jurisdiction, indicating whether it is a country, state/region, city, for example.
	// Example: Country
	// Required: true
	// Enum: [Country State County City Special]
	Type *string `json:"type"`

	// The "Seller's Use" tax rate specific to this jurisdiction.
	// Example: 0.06
	UseRate float64 `json:"useRate,omitempty"`
}

// Validate validates this jurisdiction model
func (m *JurisdictionModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignatureCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JurisdictionModel) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionModel) validateEffectiveDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectiveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveDate", "body", "date-time", m.EffectiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionModel) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionModel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionModel) validateSignatureCode(formats strfmt.Registry) error {

	if err := validate.Required("signatureCode", "body", m.SignatureCode); err != nil {
		return err
	}

	return nil
}

var jurisdictionModelTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Country","State","County","City","Special"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jurisdictionModelTypeTypePropEnum = append(jurisdictionModelTypeTypePropEnum, v)
	}
}

const (

	// JurisdictionModelTypeCountry captures enum value "Country"
	JurisdictionModelTypeCountry string = "Country"

	// JurisdictionModelTypeState captures enum value "State"
	JurisdictionModelTypeState string = "State"

	// JurisdictionModelTypeCounty captures enum value "County"
	JurisdictionModelTypeCounty string = "County"

	// JurisdictionModelTypeCity captures enum value "City"
	JurisdictionModelTypeCity string = "City"

	// JurisdictionModelTypeSpecial captures enum value "Special"
	JurisdictionModelTypeSpecial string = "Special"
)

// prop value enum
func (m *JurisdictionModel) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, jurisdictionModelTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JurisdictionModel) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this jurisdiction model based on context it is used
func (m *JurisdictionModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JurisdictionModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JurisdictionModel) UnmarshalBinary(b []byte) error {
	var res JurisdictionModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}