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

// CompanyDistanceThresholdModel A company-distance-threshold model indicates the distance between a company
// and the taxing borders of various countries.  Distance thresholds are necessary
// to correctly calculate some value-added taxes.
//
// Distance thresholds only apply to sales of goods in certain countries.  A distance threshold
// is applied for each ship-from/ship-to combination of countries.  The threshold amount is defined by
// the ship-to country.
//
// Generally, if you have exceeded a distance threshold for taxes between a pair of countries, your tax calculation
// will be determined to be the rate in the destination country.  If you have not exceeded the threshold,
// your tax calculation will be determined to be the rate in the origin country.
//
// The amount of a threshold is not tracked or managed in AvaTax, but the decision of your tax compliance department
// as to whether you have exceeded this threshold is maintained in this object.
//
// By default, you are considered to have exceeded tax thresholds. If you wish to change this default, you can create
// a company-distance-threshold object to select the correct behavior for this origin/destination tax calculation process.
// Example: {"companyId":0,"destinationCountry":"BR","id":0,"originCountry":"FR","thresholdExceeded":false,"type":"Sale"}
//
// swagger:model CompanyDistanceThresholdModel
type CompanyDistanceThresholdModel struct {

	// The ID number of the company that defined this distance threshold.
	// Example: 0
	// Read Only: true
	CompanyID int32 `json:"companyId,omitempty"`

	// The destination country for this threshold.
	//
	// This field supports many different country identifiers:
	//  * Two character ISO 3166 codes
	//  * Three character ISO 3166 codes
	//  * Fully spelled out names of the country in ISO supported languages
	//  * Common alternative spellings for many countries
	//
	// For a full list of all supported codes and names, please see the Definitions API `ListCountries`.
	// Example: BR
	// Required: true
	DestinationCountry *string `json:"destinationCountry"`

	// For distance threshold values that change over time, this is the earliest date for which this distance
	// threshold is valid.  If null, this distance threshold is valid for all dates earlier than the `endDate` field.
	// Format: date-time
	EffDate strfmt.DateTime `json:"effDate,omitempty"`

	// For distance threshold values that change over time, this is the latest date for which this distance
	// threshold is valid.  If null, this distance threshold is valid for all dates later than the `effDate` field.
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// A unique ID number representing this distance threshold object.
	// Example: 0
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// The origin country for this threshold.
	//
	// This field supports many different country identifiers:
	//  * Two character ISO 3166 codes
	//  * Three character ISO 3166 codes
	//  * Fully spelled out names of the country in ISO supported languages
	//  * Common alternative spellings for many countries
	//
	// For a full list of all supported codes and names, please see the Definitions API `ListCountries`.
	// Example: FR
	// Required: true
	OriginCountry *string `json:"originCountry"`

	// True if your tax professional has determined that the value-added tax distance threshold is exceeded for this pair of countries.
	//
	// If you set this value to `false`, your value added taxes will be calculated using the origin country.  Otherwise, value
	// added taxes will be calculated on the destination country.
	// Example: false
	ThresholdExceeded bool `json:"thresholdExceeded,omitempty"`

	// Indicates the distance threshold type.
	//
	// This value can be either `Sale` or `Purchase`.
	// Example: Sale
	// Required: true
	// Max Length: 20
	// Min Length: 0
	Type *string `json:"type"`
}

// Validate validates this company distance threshold model
func (m *CompanyDistanceThresholdModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginCountry(formats); err != nil {
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

func (m *CompanyDistanceThresholdModel) validateDestinationCountry(formats strfmt.Registry) error {

	if err := validate.Required("destinationCountry", "body", m.DestinationCountry); err != nil {
		return err
	}

	return nil
}

func (m *CompanyDistanceThresholdModel) validateEffDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EffDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effDate", "body", "date-time", m.EffDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CompanyDistanceThresholdModel) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CompanyDistanceThresholdModel) validateOriginCountry(formats strfmt.Registry) error {

	if err := validate.Required("originCountry", "body", m.OriginCountry); err != nil {
		return err
	}

	return nil
}

func (m *CompanyDistanceThresholdModel) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.MinLength("type", "body", *m.Type, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("type", "body", *m.Type, 20); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this company distance threshold model based on the context it is used
func (m *CompanyDistanceThresholdModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCompanyID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompanyDistanceThresholdModel) contextValidateCompanyID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "companyId", "body", int32(m.CompanyID)); err != nil {
		return err
	}

	return nil
}

func (m *CompanyDistanceThresholdModel) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompanyDistanceThresholdModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompanyDistanceThresholdModel) UnmarshalBinary(b []byte) error {
	var res CompanyDistanceThresholdModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}