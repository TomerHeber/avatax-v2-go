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

// ProductSystemCountryModel Represents a System Country.
// Example: {"country":"US","effDate":"2021-12-25T00:00:00+00:00","endDate":"9999-12-31T00:00:00","systemCountryId":1,"systemId":1}
//
// swagger:model ProductSystemCountryModel
type ProductSystemCountryModel struct {

	// string value of country code for SystemCountry
	// Example: US
	Country string `json:"country,omitempty"`

	// DateTime as EffDate for SystemCountry
	// Example: 2021-12-25T00:00:00+00:00
	// Format: date-time
	EffDate strfmt.DateTime `json:"effDate,omitempty"`

	// DateTime as EffDate for SystemCountry
	// Example: 9999-12-31T00:00:00
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// Its Integer SystemCountryId value for SystemCountry
	// Example: 1
	SystemCountryID int32 `json:"systemCountryId,omitempty"`

	// Its Integer SystemId value for SystemCountry
	// Example: 1
	SystemID int32 `json:"systemId,omitempty"`
}

// Validate validates this product system country model
func (m *ProductSystemCountryModel) Validate(formats strfmt.Registry) error {
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

func (m *ProductSystemCountryModel) validateEffDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EffDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effDate", "body", "date-time", m.EffDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProductSystemCountryModel) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this product system country model based on context it is used
func (m *ProductSystemCountryModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProductSystemCountryModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductSystemCountryModel) UnmarshalBinary(b []byte) error {
	var res ProductSystemCountryModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}