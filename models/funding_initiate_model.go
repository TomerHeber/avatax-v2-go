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

// FundingInitiateModel funding initiate model
// Example: {"fundingEmailRecipient":"user@example.org","requestEmail":true}
//
// swagger:model FundingInitiateModel
type FundingInitiateModel struct {

	// If you have requested an email for funding setup, this is the recipient who will receive an
	// email inviting them to setup funding configuration for Avalara Managed Returns.  The recipient can
	// then click on a link in the email and setup funding configuration for this company.
	// Example: user@example.org
	// Required: true
	FundingEmailRecipient *string `json:"fundingEmailRecipient"`

	// Set this value to true to request an email to the recipient
	// Example: true
	RequestEmail bool `json:"requestEmail,omitempty"`

	// Set this value to true to request an HTML-based funding widget that can be embedded within an
	// existing user interface.  A user can then interact with the HTML-based funding widget to set up
	// funding information for the company.
	RequestWidget bool `json:"requestWidget,omitempty"`
}

// Validate validates this funding initiate model
func (m *FundingInitiateModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFundingEmailRecipient(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FundingInitiateModel) validateFundingEmailRecipient(formats strfmt.Registry) error {

	if err := validate.Required("fundingEmailRecipient", "body", m.FundingEmailRecipient); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this funding initiate model based on context it is used
func (m *FundingInitiateModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FundingInitiateModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FundingInitiateModel) UnmarshalBinary(b []byte) error {
	var res FundingInitiateModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}