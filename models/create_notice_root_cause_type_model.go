// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateNoticeRootCauseTypeModel Model to create a new tax notice root cause type.
// Example: {"description":"Customer-Invalid login (Return filed manual)","isActive":true,"sortOrder":55}
//
// swagger:model CreateNoticeRootCauseTypeModel
type CreateNoticeRootCauseTypeModel struct {

	// The description name of this notice RootCause
	// Example: Customer-Invalid login (Return filed manual)
	Description string `json:"description,omitempty"`

	// Defines if the RootCause is active
	// Example: true
	IsActive bool `json:"isActive,omitempty"`

	// The sort order of this RootCause
	// Example: 55
	SortOrder int32 `json:"sortOrder,omitempty"`
}

// Validate validates this create notice root cause type model
func (m *CreateNoticeRootCauseTypeModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create notice root cause type model based on context it is used
func (m *CreateNoticeRootCauseTypeModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateNoticeRootCauseTypeModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateNoticeRootCauseTypeModel) UnmarshalBinary(b []byte) error {
	var res CreateNoticeRootCauseTypeModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}