// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SyncItemsResponseModel The response returned after an item sync was requested.
// Example: {"status":"Accepted"}
//
// swagger:model SyncItemsResponseModel
type SyncItemsResponseModel struct {

	// The status of the request
	// Example: Accepted
	Status string `json:"status,omitempty"`
}

// Validate validates this sync items response model
func (m *SyncItemsResponseModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sync items response model based on context it is used
func (m *SyncItemsResponseModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SyncItemsResponseModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncItemsResponseModel) UnmarshalBinary(b []byte) error {
	var res SyncItemsResponseModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}