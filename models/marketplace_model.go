// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MarketplaceModel Marketplace Location Output model
// Example: {"marketplace":"Amazon","marketplaceId":"a0n0b00000ODPh2AAH"}
//
// swagger:model MarketplaceModel
type MarketplaceModel struct {

	// Marketplace Location
	// Example: Amazon
	Marketplace string `json:"marketplace,omitempty"`

	// Marketplace Location Id
	// Example: a0n0b00000ODPh2AAH
	MarketplaceID string `json:"marketplaceId,omitempty"`
}

// Validate validates this marketplace model
func (m *MarketplaceModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this marketplace model based on context it is used
func (m *MarketplaceModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MarketplaceModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketplaceModel) UnmarshalBinary(b []byte) error {
	var res MarketplaceModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}