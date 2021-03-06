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

// CompanyInitializationModel Company Initialization Model
// Example: {"city":"Irvine","companyCode":"DEFAULT","country":"US","email":"bob@example.org","firstName":"Bob","isFein":false,"lastName":"Example","line1":"2000 Main Street","mobileNumber":"714 555-1212","name":"Bob's Artisan Pottery","phoneNumber":"714 555-2121","postalCode":"92614","region":"CA","taxpayerIdNumber":"12-3456789","title":"Owner"}
//
// swagger:model CompanyInitializationModel
type CompanyInitializationModel struct {

	// City
	// Example: Irvine
	// Required: true
	// Max Length: 50
	// Min Length: 0
	City *string `json:"city"`

	// Company Code - used to distinguish between companies within your accounting system
	// Example: DEFAULT
	// Max Length: 25
	// Min Length: 0
	CompanyCode *string `json:"companyCode,omitempty"`

	// Name or ISO 3166 code identifying the country.
	//
	// This field supports many different country identifiers:
	//  * Two character ISO 3166 codes
	//  * Three character ISO 3166 codes
	//  * Fully spelled out names of the country in ISO supported languages
	//  * Common alternative spellings for many countries
	//
	// For a full list of all supported codes and names, please see the Definitions API `ListCountries`.
	// Example: US
	// Required: true
	Country *string `json:"country"`

	// Email
	// Example: bob@example.org
	// Required: true
	// Max Length: 50
	// Min Length: 0
	Email *string `json:"email"`

	// Fax Number
	// Max Length: 25
	// Min Length: 0
	FaxNumber *string `json:"faxNumber,omitempty"`

	// First Name
	// Example: Bob
	// Required: true
	// Max Length: 50
	// Min Length: 0
	FirstName *string `json:"firstName"`

	// Set this field to true if the taxPayerIdNumber is a FEIN.
	IsFein bool `json:"isFein,omitempty"`

	// Last Name
	// Example: Example
	// Required: true
	// Max Length: 50
	// Min Length: 0
	LastName *string `json:"lastName"`

	// Address Line1
	// Example: 2000 Main Street
	// Required: true
	// Max Length: 50
	// Min Length: 0
	Line1 *string `json:"line1"`

	// Line2
	// Max Length: 50
	// Min Length: 0
	Line2 *string `json:"line2,omitempty"`

	// Line3
	// Max Length: 50
	// Min Length: 0
	Line3 *string `json:"line3,omitempty"`

	// Mobile Number
	// Example: 714 555-1212
	// Max Length: 25
	// Min Length: 0
	MobileNumber *string `json:"mobileNumber,omitempty"`

	// Company Name
	// Example: Bob's Artisan Pottery
	// Required: true
	// Max Length: 50
	// Min Length: 0
	Name *string `json:"name"`

	// Parent Company ID
	ParentCompanyID int32 `json:"parentCompanyId,omitempty"`

	// Phone Number
	// Example: 714 555-2121
	// Required: true
	// Max Length: 25
	// Min Length: 0
	PhoneNumber *string `json:"phoneNumber"`

	// Postal Code
	// Example: 92614
	// Required: true
	// Max Length: 10
	// Min Length: 0
	PostalCode *string `json:"postalCode"`

	// Name or ISO 3166 code identifying the region within the country.
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

	// United States Taxpayer ID number, usually your Employer Identification Number if you are a business or your
	// Social Security Number if you are an individual.
	// This value is required if the address provided is inside the US and if you subscribed to the Avalara Managed Returns or SST Certified Service Provider service. Otherwise it is optional.
	// Example: 12-3456789
	// Max Length: 11
	// Min Length: 0
	TaxpayerIDNumber *string `json:"taxpayerIdNumber,omitempty"`

	// Title
	// Example: Owner
	// Max Length: 50
	// Min Length: 0
	Title *string `json:"title,omitempty"`

	// Vat Registration Id - leave blank if not known.
	// Max Length: 25
	// Min Length: 0
	VatRegistrationID *string `json:"vatRegistrationId,omitempty"`
}

// Validate validates this company initialization model
func (m *CompanyInitializationModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompanyCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFaxNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLine1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLine2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLine3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMobileNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxpayerIDNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVatRegistrationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompanyInitializationModel) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	if err := validate.MinLength("city", "body", *m.City, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("city", "body", *m.City, 50); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validateCompanyCode(formats strfmt.Registry) error {
	if swag.IsZero(m.CompanyCode) { // not required
		return nil
	}

	if err := validate.MinLength("companyCode", "body", *m.CompanyCode, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("companyCode", "body", *m.CompanyCode, 25); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.MinLength("email", "body", *m.Email, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("email", "body", *m.Email, 50); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validateFaxNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.FaxNumber) { // not required
		return nil
	}

	if err := validate.MinLength("faxNumber", "body", *m.FaxNumber, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("faxNumber", "body", *m.FaxNumber, 25); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	if err := validate.MinLength("firstName", "body", *m.FirstName, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("firstName", "body", *m.FirstName, 50); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	if err := validate.MinLength("lastName", "body", *m.LastName, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("lastName", "body", *m.LastName, 50); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validateLine1(formats strfmt.Registry) error {

	if err := validate.Required("line1", "body", m.Line1); err != nil {
		return err
	}

	if err := validate.MinLength("line1", "body", *m.Line1, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("line1", "body", *m.Line1, 50); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validateLine2(formats strfmt.Registry) error {
	if swag.IsZero(m.Line2) { // not required
		return nil
	}

	if err := validate.MinLength("line2", "body", *m.Line2, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("line2", "body", *m.Line2, 50); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validateLine3(formats strfmt.Registry) error {
	if swag.IsZero(m.Line3) { // not required
		return nil
	}

	if err := validate.MinLength("line3", "body", *m.Line3, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("line3", "body", *m.Line3, 50); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validateMobileNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.MobileNumber) { // not required
		return nil
	}

	if err := validate.MinLength("mobileNumber", "body", *m.MobileNumber, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("mobileNumber", "body", *m.MobileNumber, 25); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 50); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validatePhoneNumber(formats strfmt.Registry) error {

	if err := validate.Required("phoneNumber", "body", m.PhoneNumber); err != nil {
		return err
	}

	if err := validate.MinLength("phoneNumber", "body", *m.PhoneNumber, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("phoneNumber", "body", *m.PhoneNumber, 25); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validatePostalCode(formats strfmt.Registry) error {

	if err := validate.Required("postalCode", "body", m.PostalCode); err != nil {
		return err
	}

	if err := validate.MinLength("postalCode", "body", *m.PostalCode, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("postalCode", "body", *m.PostalCode, 10); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validateTaxpayerIDNumber(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxpayerIDNumber) { // not required
		return nil
	}

	if err := validate.MinLength("taxpayerIdNumber", "body", *m.TaxpayerIDNumber, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("taxpayerIdNumber", "body", *m.TaxpayerIDNumber, 11); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validateTitle(formats strfmt.Registry) error {
	if swag.IsZero(m.Title) { // not required
		return nil
	}

	if err := validate.MinLength("title", "body", *m.Title, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 50); err != nil {
		return err
	}

	return nil
}

func (m *CompanyInitializationModel) validateVatRegistrationID(formats strfmt.Registry) error {
	if swag.IsZero(m.VatRegistrationID) { // not required
		return nil
	}

	if err := validate.MinLength("vatRegistrationId", "body", *m.VatRegistrationID, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("vatRegistrationId", "body", *m.VatRegistrationID, 25); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this company initialization model based on context it is used
func (m *CompanyInitializationModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CompanyInitializationModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompanyInitializationModel) UnmarshalBinary(b []byte) error {
	var res CompanyInitializationModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
