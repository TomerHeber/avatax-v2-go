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

// MarketplaceLocationModelFetchResult marketplace location model fetch result
//
// swagger:model MarketplaceLocationModelFetchResult
type MarketplaceLocationModelFetchResult struct {

	// at next link
	AtNextLink string `json:"@nextLink,omitempty"`

	// at recordset count
	AtRecordsetCount int32 `json:"@recordsetCount,omitempty"`

	// page key
	PageKey string `json:"pageKey,omitempty"`

	// value
	Value []*MarketplaceLocationModel `json:"value"`
}

// Validate validates this marketplace location model fetch result
func (m *MarketplaceLocationModelFetchResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MarketplaceLocationModelFetchResult) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	for i := 0; i < len(m.Value); i++ {
		if swag.IsZero(m.Value[i]) { // not required
			continue
		}

		if m.Value[i] != nil {
			if err := m.Value[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("value" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("value" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this marketplace location model fetch result based on the context it is used
func (m *MarketplaceLocationModelFetchResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MarketplaceLocationModelFetchResult) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Value); i++ {

		if m.Value[i] != nil {
			if err := m.Value[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("value" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("value" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MarketplaceLocationModelFetchResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketplaceLocationModelFetchResult) UnmarshalBinary(b []byte) error {
	var res MarketplaceLocationModelFetchResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}