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

// CycleSafeEditRequestModel Options for expiring a filing calendar.
// Example: {"companyId":0,"edits":[{"destination":"","fieldName":"","newValue":"","oldValue":"","questionCode":"","questionId":0}],"filingCalendarId":0,"taxFormCode":""}
//
// swagger:model CycleSafeEditRequestModel
type CycleSafeEditRequestModel struct {

	// Company Identifier
	// Example: 0
	CompanyID int32 `json:"companyId,omitempty"`

	// Filing calendar edits
	// Example: [{"destination":"","fieldName":"","newValue":"","oldValue":"","questionCode":"","questionId":0}]
	Edits []*CycleSafeFilingCalendarEditModel `json:"edits"`

	// Filing Calendar Identifier
	// Example: 0
	FilingCalendarID int64 `json:"filingCalendarId,omitempty"`

	// Tax Form Code
	TaxFormCode string `json:"taxFormCode,omitempty"`
}

// Validate validates this cycle safe edit request model
func (m *CycleSafeEditRequestModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEdits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CycleSafeEditRequestModel) validateEdits(formats strfmt.Registry) error {
	if swag.IsZero(m.Edits) { // not required
		return nil
	}

	for i := 0; i < len(m.Edits); i++ {
		if swag.IsZero(m.Edits[i]) { // not required
			continue
		}

		if m.Edits[i] != nil {
			if err := m.Edits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("edits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cycle safe edit request model based on the context it is used
func (m *CycleSafeEditRequestModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEdits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CycleSafeEditRequestModel) contextValidateEdits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Edits); i++ {

		if m.Edits[i] != nil {
			if err := m.Edits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("edits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("edits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CycleSafeEditRequestModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CycleSafeEditRequestModel) UnmarshalBinary(b []byte) error {
	var res CycleSafeEditRequestModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}