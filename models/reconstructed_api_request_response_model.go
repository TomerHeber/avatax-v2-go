// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ReconstructedAPIRequestResponseModel This model contains a reconstructed CreateTransaction request object that could potentially be used
// to recreate this transaction.
//
// Note that the API changes over time, and this reconstructed model is likely different from the exact request
// that was originally used to create this transaction.
// Example: {"request":{"addresses":{"singleLocation":{"city":"Irvine","country":"US","line1":"2000 Main Street","postalCode":"92614","region":"CA"}},"commit":true,"companyCode":"DEFAULT","currencyCode":"USD","customerCode":"ABC","date":"2017-10-12","description":"Yarn","lines":[{"amount":100,"description":"Yarn","itemCode":"Y0001","number":"1","quantity":1,"taxCode":"PS081282"}],"purchaseOrderNo":"2017 - 10 - 12 - 001","type":"SalesInvoice"}}
//
// swagger:model ReconstructedApiRequestResponseModel
type ReconstructedAPIRequestResponseModel struct {

	// API request
	// Example: {"addresses":{"singleLocation":{"city":"Irvine","country":"US","line1":"2000 Main Street","postalCode":"92614","region":"CA"}},"commit":true,"companyCode":"DEFAULT","currencyCode":"USD","customerCode":"ABC","date":"2017-10-12","description":"Yarn","lines":[{"amount":100,"description":"Yarn","itemCode":"Y0001","number":"1","quantity":1,"taxCode":"PS081282"}],"purchaseOrderNo":"2017 - 10 - 12 - 001","type":"SalesInvoice"}
	Request *CreateTransactionModel `json:"request,omitempty"`
}

// Validate validates this reconstructed Api request response model
func (m *ReconstructedAPIRequestResponseModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReconstructedAPIRequestResponseModel) validateRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.Request) { // not required
		return nil
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this reconstructed Api request response model based on the context it is used
func (m *ReconstructedAPIRequestResponseModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReconstructedAPIRequestResponseModel) contextValidateRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.Request != nil {
		if err := m.Request.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReconstructedAPIRequestResponseModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReconstructedAPIRequestResponseModel) UnmarshalBinary(b []byte) error {
	var res ReconstructedAPIRequestResponseModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
