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
	"github.com/go-openapi/validate"
)

// DeleteErrorTransactionsRequestModel Request model for when a user is deleting multiple error transaction
// Example: {"models":[{"documentCode":"2e604c50-82f9-43e0-90ba-5e95fd27625f","documentType":"SalesOrder"}]}
//
// swagger:model DeleteErrorTransactionsRequestModel
type DeleteErrorTransactionsRequestModel struct {

	// List of error transactions to be deleted
	// Example: [{"documentCode":"2e604c50-82f9-43e0-90ba-5e95fd27625f","documentType":"SalesOrder"}]
	// Required: true
	Models []*ErrorTransactionModelBase `json:"models"`
}

// Validate validates this delete error transactions request model
func (m *DeleteErrorTransactionsRequestModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteErrorTransactionsRequestModel) validateModels(formats strfmt.Registry) error {

	if err := validate.Required("models", "body", m.Models); err != nil {
		return err
	}

	for i := 0; i < len(m.Models); i++ {
		if swag.IsZero(m.Models[i]) { // not required
			continue
		}

		if m.Models[i] != nil {
			if err := m.Models[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("models" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("models" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this delete error transactions request model based on the context it is used
func (m *DeleteErrorTransactionsRequestModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteErrorTransactionsRequestModel) contextValidateModels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Models); i++ {

		if m.Models[i] != nil {
			if err := m.Models[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("models" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("models" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeleteErrorTransactionsRequestModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteErrorTransactionsRequestModel) UnmarshalBinary(b []byte) error {
	var res DeleteErrorTransactionsRequestModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}