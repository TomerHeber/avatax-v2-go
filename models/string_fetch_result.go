// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StringFetchResult string fetch result
//
// swagger:model StringFetchResult
type StringFetchResult struct {

	// at next link
	AtNextLink string `json:"@nextLink,omitempty"`

	// at recordset count
	AtRecordsetCount int32 `json:"@recordsetCount,omitempty"`

	// page key
	PageKey string `json:"pageKey,omitempty"`

	// value
	Value []string `json:"value"`
}

// Validate validates this string fetch result
func (m *StringFetchResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this string fetch result based on context it is used
func (m *StringFetchResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StringFetchResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StringFetchResult) UnmarshalBinary(b []byte) error {
	var res StringFetchResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}