// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NoticeResponsibilityModel NoticeResponsibility Model
// Example: {"description":"Customer-Invalid login (Return filed manual)","id":11,"isActive":true,"sortOrder":55}
//
// swagger:model NoticeResponsibilityModel
type NoticeResponsibilityModel struct {

	// The description name of this notice responsibility
	// Example: Customer-Invalid login (Return filed manual)
	Description string `json:"description,omitempty"`

	// The unique ID number of this notice responsibility.
	// Example: 11
	ID int32 `json:"id,omitempty"`

	// Defines if the responsibility is active
	// Example: true
	IsActive bool `json:"isActive,omitempty"`

	// The sort order of this responsibility
	// Example: 55
	SortOrder int32 `json:"sortOrder,omitempty"`
}

// Validate validates this notice responsibility model
func (m *NoticeResponsibilityModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this notice responsibility model based on context it is used
func (m *NoticeResponsibilityModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NoticeResponsibilityModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NoticeResponsibilityModel) UnmarshalBinary(b []byte) error {
	var res NoticeResponsibilityModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
