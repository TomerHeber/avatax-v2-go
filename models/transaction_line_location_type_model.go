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

// TransactionLineLocationTypeModel Represents information about location types stored in a line
// Example: {"documentAddressId":789,"documentLineId":456,"documentLineLocationTypeId":123,"locationTypeCode":"SingleLocation"}
//
// swagger:model TransactionLineLocationTypeModel
type TransactionLineLocationTypeModel struct {

	// The address ID corresponding to this model
	// Example: 789
	// Read Only: true
	DocumentAddressID int64 `json:"documentAddressId,omitempty"`

	// The unique ID number of the document line associated with this line location address model
	// Example: 456
	// Read Only: true
	DocumentLineID int64 `json:"documentLineId,omitempty"`

	// The unique ID number of this line location address model
	// Example: 123
	// Read Only: true
	DocumentLineLocationTypeID int64 `json:"documentLineLocationTypeId,omitempty"`

	// The location type code corresponding to this model
	// Example: SingleLocation
	// Read Only: true
	LocationTypeCode string `json:"locationTypeCode,omitempty"`
}

// Validate validates this transaction line location type model
func (m *TransactionLineLocationTypeModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this transaction line location type model based on the context it is used
func (m *TransactionLineLocationTypeModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDocumentAddressID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDocumentLineID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDocumentLineLocationTypeID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocationTypeCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionLineLocationTypeModel) contextValidateDocumentAddressID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "documentAddressId", "body", int64(m.DocumentAddressID)); err != nil {
		return err
	}

	return nil
}

func (m *TransactionLineLocationTypeModel) contextValidateDocumentLineID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "documentLineId", "body", int64(m.DocumentLineID)); err != nil {
		return err
	}

	return nil
}

func (m *TransactionLineLocationTypeModel) contextValidateDocumentLineLocationTypeID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "documentLineLocationTypeId", "body", int64(m.DocumentLineLocationTypeID)); err != nil {
		return err
	}

	return nil
}

func (m *TransactionLineLocationTypeModel) contextValidateLocationTypeCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "locationTypeCode", "body", string(m.LocationTypeCode)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionLineLocationTypeModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionLineLocationTypeModel) UnmarshalBinary(b []byte) error {
	var res TransactionLineLocationTypeModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}