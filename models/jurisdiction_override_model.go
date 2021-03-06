// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JurisdictionOverrideModel Represents an override of tax jurisdictions for a specific address.
//
// During the time period represented by EffDate through EndDate, all tax decisions for addresses matching
// this override object will be assigned to the list of jurisdictions designated in this object.
// Example: {"accountId":0,"boundaryLevel":"Address","city":"Irvine","country":"US","description":"a test JO","effectiveDate":"2006-08-01","endDate":"2099-01-12","id":0,"isDefault":false,"jurisdictions":[{"code":"53","name":"WASHINGTON","rate":0.06,"region":"WA","salesRate":0.06,"signatureCode":"","type":"State","useRate":0.06}],"line1":"5500 Irvine Center Dr","postalCode":"92618","region":"CA","taxRegionId":0}
//
// swagger:model JurisdictionOverrideModel
type JurisdictionOverrideModel struct {

	// The unique ID number assigned to this account.
	// Example: 0
	AccountID int32 `json:"accountId,omitempty"`

	// The boundary level of this override
	// Example: Address
	// Read Only: true
	// Enum: [Address Zip9 Zip5]
	BoundaryLevel string `json:"boundaryLevel,omitempty"`

	// The city address of the physical location affected by this override.
	// Example: Irvine
	// Max Length: 50
	// Min Length: 0
	City *string `json:"city,omitempty"`

	// The two character ISO-3166 country code of the country affected by this override.
	//
	// Note that only United States addresses are affected by the jurisdiction override system.
	// Example: US
	// Read Only: true
	Country string `json:"country,omitempty"`

	// The date when this record was created.
	// Read Only: true
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// The User ID of the user who created this record.
	// Example: 0
	// Read Only: true
	CreatedUserID int32 `json:"createdUserId,omitempty"`

	// A description of why this jurisdiction override was created.
	// Example: a test JO
	// Required: true
	// Max Length: 50
	// Min Length: 0
	Description *string `json:"description"`

	// The date when this override first takes effect.  Set this value to null to affect all dates up to the end date.
	// Example: 2006-08-01T00:00:00
	// Format: date
	EffectiveDate strfmt.Date `json:"effectiveDate,omitempty"`

	// The date when this override will cease to take effect.  Set this value to null to never expire.
	// Example: 2099-01-12T00:00:00
	// Format: date
	EndDate strfmt.Date `json:"endDate,omitempty"`

	// The unique ID number of this override.
	// Example: 0
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// True if this is a default boundary
	IsDefault bool `json:"isDefault,omitempty"`

	// A list of the tax jurisdictions that will be assigned to this overridden address.
	// Required: true
	Jurisdictions []*JurisdictionModel `json:"jurisdictions"`

	// The street address of the physical location affected by this override.
	// Example: 5500 Irvine Center Dr
	// Max Length: 50
	// Min Length: 0
	Line1 *string `json:"line1,omitempty"`

	// The date/time when this record was last modified.
	// Read Only: true
	// Format: date-time
	ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

	// The user ID of the user who last modified this record.
	// Example: 0
	// Read Only: true
	ModifiedUserID int32 `json:"modifiedUserId,omitempty"`

	// The postal code of the physical location affected by this override.
	// Example: 92618
	// Required: true
	// Max Length: 11
	// Min Length: 0
	PostalCode *string `json:"postalCode"`

	// Name or ISO 3166 code identifying the region within the country to be affected by this override.
	//
	// Note that only United States addresses are affected by the jurisdiction override system.
	//
	// This field supports many different region identifiers:
	//  * Two and three character ISO 3166 region codes
	//  * Fully spelled out names of the region in ISO supported languages
	//  * Common alternative spellings for many regions
	//
	// For a full list of all supported codes and names, please see the Definitions API `ListRegions`.
	// Example: CA
	// Required: true
	Region *string `json:"region"`

	// The TaxRegionId of the new location affected by this jurisdiction override.
	// Example: 0
	// Required: true
	TaxRegionID *int32 `json:"taxRegionId"`
}

// Validate validates this jurisdiction override model
func (m *JurisdictionOverrideModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoundaryLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.validateJurisdictions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLine1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxRegionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var jurisdictionOverrideModelTypeBoundaryLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Address","Zip9","Zip5"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jurisdictionOverrideModelTypeBoundaryLevelPropEnum = append(jurisdictionOverrideModelTypeBoundaryLevelPropEnum, v)
	}
}

const (

	// JurisdictionOverrideModelBoundaryLevelAddress captures enum value "Address"
	JurisdictionOverrideModelBoundaryLevelAddress string = "Address"

	// JurisdictionOverrideModelBoundaryLevelZip9 captures enum value "Zip9"
	JurisdictionOverrideModelBoundaryLevelZip9 string = "Zip9"

	// JurisdictionOverrideModelBoundaryLevelZip5 captures enum value "Zip5"
	JurisdictionOverrideModelBoundaryLevelZip5 string = "Zip5"
)

// prop value enum
func (m *JurisdictionOverrideModel) validateBoundaryLevelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, jurisdictionOverrideModelTypeBoundaryLevelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JurisdictionOverrideModel) validateBoundaryLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.BoundaryLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateBoundaryLevelEnum("boundaryLevel", "body", m.BoundaryLevel); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) validateCity(formats strfmt.Registry) error {
	if swag.IsZero(m.City) { // not required
		return nil
	}

	if err := validate.MinLength("city", "body", *m.City, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("city", "body", *m.City, 50); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) validateCreatedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", *m.Description, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", *m.Description, 50); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) validateEffectiveDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectiveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveDate", "body", "date", m.EffectiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) validateJurisdictions(formats strfmt.Registry) error {

	if err := validate.Required("jurisdictions", "body", m.Jurisdictions); err != nil {
		return err
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

func (m *JurisdictionOverrideModel) validateLine1(formats strfmt.Registry) error {
	if swag.IsZero(m.Line1) { // not required
		return nil
	}

	if err := validate.MinLength("line1", "body", *m.Line1, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("line1", "body", *m.Line1, 50); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) validateModifiedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) validatePostalCode(formats strfmt.Registry) error {

	if err := validate.Required("postalCode", "body", m.PostalCode); err != nil {
		return err
	}

	if err := validate.MinLength("postalCode", "body", *m.PostalCode, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("postalCode", "body", *m.PostalCode, 11); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) validateTaxRegionID(formats strfmt.Registry) error {

	if err := validate.Required("taxRegionId", "body", m.TaxRegionID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this jurisdiction override model based on the context it is used
func (m *JurisdictionOverrideModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBoundaryLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountry(ctx, formats); err != nil {
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

	if err := m.contextValidateJurisdictions(ctx, formats); err != nil {
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

func (m *JurisdictionOverrideModel) contextValidateBoundaryLevel(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "boundaryLevel", "body", string(m.BoundaryLevel)); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) contextValidateCountry(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "country", "body", string(m.Country)); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) contextValidateCreatedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdDate", "body", strfmt.DateTime(m.CreatedDate)); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) contextValidateCreatedUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdUserId", "body", int32(m.CreatedUserID)); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) contextValidateJurisdictions(ctx context.Context, formats strfmt.Registry) error {

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

func (m *JurisdictionOverrideModel) contextValidateModifiedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedDate", "body", strfmt.DateTime(m.ModifiedDate)); err != nil {
		return err
	}

	return nil
}

func (m *JurisdictionOverrideModel) contextValidateModifiedUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedUserId", "body", int32(m.ModifiedUserID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JurisdictionOverrideModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JurisdictionOverrideModel) UnmarshalBinary(b []byte) error {
	var res JurisdictionOverrideModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
