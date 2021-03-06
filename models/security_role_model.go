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

// SecurityRoleModel Represents a single security role.
// Example: {"description":"AccountAdmin","id":3}
//
// swagger:model SecurityRoleModel
type SecurityRoleModel struct {

	// A description of this security role
	// Example: AccountAdmin
	// Read Only: true
	Description string `json:"description,omitempty"`

	// The unique ID number of this security role.
	// Example: 3
	// Read Only: true
	ID int32 `json:"id,omitempty"`
}

// Validate validates this security role model
func (m *SecurityRoleModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this security role model based on the context it is used
func (m *SecurityRoleModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *SecurityRoleModel) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *SecurityRoleModel) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int32(m.ID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecurityRoleModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecurityRoleModel) UnmarshalBinary(b []byte) error {
	var res SecurityRoleModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
