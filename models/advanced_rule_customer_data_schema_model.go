// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdvancedRuleCustomerDataSchemaModel Model for retrieving customer data schema
// Example: {"ruleId":"1ddd8c0a73ca4beb888c07ff1aacf2bc"}
//
// swagger:model AdvancedRuleCustomerDataSchemaModel
type AdvancedRuleCustomerDataSchemaModel struct {

	// Customer data schema
	CustomerDataSchema string `json:"customerDataSchema,omitempty"`

	// Unique identifier for the rule
	// Example: 1ddd8c0a73ca4beb888c07ff1aacf2bc
	RuleID string `json:"ruleId,omitempty"`
}

// Validate validates this advanced rule customer data schema model
func (m *AdvancedRuleCustomerDataSchemaModel) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this advanced rule customer data schema model based on context it is used
func (m *AdvancedRuleCustomerDataSchemaModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdvancedRuleCustomerDataSchemaModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdvancedRuleCustomerDataSchemaModel) UnmarshalBinary(b []byte) error {
	var res AdvancedRuleCustomerDataSchemaModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}