// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExemptionReasonModel An exemption reason defines why a certificate allows a customer to be exempt
// for purposes of tax calculation.  For a full list of defined exemption reasons,
// please call the `ListCertificateExemptionReasons` API.
// Example: {"name":"EXPOSURE"}
//
// swagger:model ExemptionReasonModel
type ExemptionReasonModel struct {

	// A unique ID number representing this exemption reason.
	ID int32 `json:"id,omitempty"`

	// A friendly name describing this exemption reason.
	// Example: EXPOSURE
	Name string `json:"name,omitempty"`
}

// Validate validates this exemption reason model
func (m *ExemptionReasonModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this exemption reason model based on context it is used
func (m *ExemptionReasonModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExemptionReasonModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExemptionReasonModel) UnmarshalBinary(b []byte) error {
	var res ExemptionReasonModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}