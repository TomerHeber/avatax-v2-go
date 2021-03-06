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
)

// FilingsCheckupAuthorityModel Cycle Safe Expiration results.
// Example: {"jurisdictionId":42,"suggestedForms":[{"country":"US","region":"NY","taxAuthorityId":71,"taxFormCode":"NYAU11"}],"tax":0,"taxAuthorityId":71,"taxAuthorityName":"NEW YORK","taxAuthorityTypeId":45,"taxTypeId":"U"}
//
// swagger:model FilingsCheckupAuthorityModel
type FilingsCheckupAuthorityModel struct {

	// Jurisdiction Id of the tax authority
	JurisdictionID int32 `json:"jurisdictionId,omitempty"`

	// Location Code of the tax authority
	LocationCode string `json:"locationCode,omitempty"`

	// Suggested forms to file due to tax collected
	SuggestedForms []*FilingsCheckupSuggestedFormModel `json:"suggestedForms"`

	// Amount of tax collected in this tax authority
	Tax float64 `json:"tax,omitempty"`

	// Unique ID of the tax authority
	TaxAuthorityID int32 `json:"taxAuthorityId,omitempty"`

	// Name of the tax authority
	TaxAuthorityName string `json:"taxAuthorityName,omitempty"`

	// Type Id of the tax authority
	TaxAuthorityTypeID int32 `json:"taxAuthorityTypeId,omitempty"`

	// Tax Type collected in the tax authority
	TaxTypeID string `json:"taxTypeId,omitempty"`
}

// Validate validates this filings checkup authority model
func (m *FilingsCheckupAuthorityModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSuggestedForms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilingsCheckupAuthorityModel) validateSuggestedForms(formats strfmt.Registry) error {
	if swag.IsZero(m.SuggestedForms) { // not required
		return nil
	}

	for i := 0; i < len(m.SuggestedForms); i++ {
		if swag.IsZero(m.SuggestedForms[i]) { // not required
			continue
		}

		if m.SuggestedForms[i] != nil {
			if err := m.SuggestedForms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("suggestedForms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("suggestedForms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this filings checkup authority model based on the context it is used
func (m *FilingsCheckupAuthorityModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSuggestedForms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilingsCheckupAuthorityModel) contextValidateSuggestedForms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SuggestedForms); i++ {

		if m.SuggestedForms[i] != nil {
			if err := m.SuggestedForms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("suggestedForms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("suggestedForms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilingsCheckupAuthorityModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilingsCheckupAuthorityModel) UnmarshalBinary(b []byte) error {
	var res FilingsCheckupAuthorityModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
