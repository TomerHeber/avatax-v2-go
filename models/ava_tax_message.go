// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AvaTaxMessage Informational or warning messages returned by AvaTax with a transaction
// Example: {"details":"Address: 123 Main Street, Irvine, CA 92615 US, TaxRegion:2127863, Latitude: 33.657808, Longitude: -117.968489","refersTo":"Addresses[0] - TaxDate: 2017-09-21 00:00:00Z","severity":"Success","source":"Avalara.AvaTax.Services.Tax.Steps","summary":"Using JAAS"}
//
// swagger:model AvaTaxMessage
type AvaTaxMessage struct {

	// Detailed information that explains what the summary provided
	// Example: Address: 123 Main Street, Irvine, CA 92615 US, TaxRegion:2127863, Latitude: 33.657808, Longitude: -117.968489
	Details string `json:"details,omitempty"`

	// Information about what object in your request this message refers to
	// Example: Addresses[0] - TaxDate: 2017-09-21 00:00:00Z
	RefersTo string `json:"refersTo,omitempty"`

	// A category that indicates how severely this message affects the results
	// Example: Success
	Severity string `json:"severity,omitempty"`

	// The name of the code or service that generated this message
	// Example: Avalara.AvaTax.Services.Tax.Steps
	Source string `json:"source,omitempty"`

	// A brief summary of what this message tells us
	// Example: Using JAAS
	Summary string `json:"summary,omitempty"`
}

// Validate validates this ava tax message
func (m *AvaTaxMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ava tax message based on context it is used
func (m *AvaTaxMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AvaTaxMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvaTaxMessage) UnmarshalBinary(b []byte) error {
	var res AvaTaxMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}