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

// NoticeResponsibilityDetailModel NoticeResponsibility Model
// Example: {"description":"Avalara-Compliance","id":0,"noticeId":4,"taxNoticeResponsibilityId":4}
//
// swagger:model NoticeResponsibilityDetailModel
type NoticeResponsibilityDetailModel struct {

	// The description name of this filing frequency
	// Example: Avalara-Compliance
	// Read Only: true
	Description string `json:"description,omitempty"`

	// The unique ID number of this filing frequency.
	// Example: 0
	// Read Only: true
	ID int32 `json:"id,omitempty"`

	// TaxNoticeId
	// Example: 4
	// Required: true
	NoticeID *int32 `json:"noticeId"`

	// TaxNoticeResponsibilityId
	// Example: 4
	// Required: true
	TaxNoticeResponsibilityID *int32 `json:"taxNoticeResponsibilityId"`
}

// Validate validates this notice responsibility detail model
func (m *NoticeResponsibilityDetailModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNoticeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxNoticeResponsibilityID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NoticeResponsibilityDetailModel) validateNoticeID(formats strfmt.Registry) error {

	if err := validate.Required("noticeId", "body", m.NoticeID); err != nil {
		return err
	}

	return nil
}

func (m *NoticeResponsibilityDetailModel) validateTaxNoticeResponsibilityID(formats strfmt.Registry) error {

	if err := validate.Required("taxNoticeResponsibilityId", "body", m.TaxNoticeResponsibilityID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this notice responsibility detail model based on the context it is used
func (m *NoticeResponsibilityDetailModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDescription(ctx, formats); err != nil {
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

func (m *NoticeResponsibilityDetailModel) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *NoticeResponsibilityDetailModel) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NoticeResponsibilityDetailModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NoticeResponsibilityDetailModel) UnmarshalBinary(b []byte) error {
	var res NoticeResponsibilityDetailModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}