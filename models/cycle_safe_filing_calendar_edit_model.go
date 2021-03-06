// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CycleSafeFilingCalendarEditModel Filing Calendar Edit
// Example: {"destination":"","fieldName":"","newValue":"","oldValue":"","questionCode":"","questionId":0}
//
// swagger:model CycleSafeFilingCalendarEditModel
type CycleSafeFilingCalendarEditModel struct {

	// Destination is used to identify filing questions' type Other or Settings.
	Destination string `json:"destination,omitempty"`

	// Field To Edit
	FieldName string `json:"fieldName,omitempty"`

	// New Value
	NewValue interface{} `json:"newValue,omitempty"`

	// Old Value
	OldValue interface{} `json:"oldValue,omitempty"`

	// The filing question code.
	QuestionCode string `json:"questionCode,omitempty"`

	// Question
	// Example: 0
	QuestionID int64 `json:"questionId,omitempty"`
}

// Validate validates this cycle safe filing calendar edit model
func (m *CycleSafeFilingCalendarEditModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cycle safe filing calendar edit model based on context it is used
func (m *CycleSafeFilingCalendarEditModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CycleSafeFilingCalendarEditModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CycleSafeFilingCalendarEditModel) UnmarshalBinary(b []byte) error {
	var res CycleSafeFilingCalendarEditModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
