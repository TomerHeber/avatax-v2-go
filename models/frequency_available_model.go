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

// FrequencyAvailableModel Frequency Available object
// Example: {"availableCycles":[{"cycleName":"","filingDueDate":"2021-12-25T17:08:42.5250831+00:00","transactionalPeriodEnd":"2022-01-24T17:08:42.5250841+00:00","transactionalPeriodStart":"2021-12-25T17:08:42.5250844+00:00"}],"frequencyCode":"","frequencyName":"","reason":""}
//
// swagger:model FrequencyAvailableModel
type FrequencyAvailableModel struct {

	// Expired Calendar End Date
	AvailableCycles []*AvailableCycleModel `json:"availableCycles"`

	// Frequency Code
	FrequencyCode string `json:"frequencyCode,omitempty"`

	// Frequency Name
	FrequencyName string `json:"frequencyName,omitempty"`

	// Reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this frequency available model
func (m *FrequencyAvailableModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableCycles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FrequencyAvailableModel) validateAvailableCycles(formats strfmt.Registry) error {
	if swag.IsZero(m.AvailableCycles) { // not required
		return nil
	}

	for i := 0; i < len(m.AvailableCycles); i++ {
		if swag.IsZero(m.AvailableCycles[i]) { // not required
			continue
		}

		if m.AvailableCycles[i] != nil {
			if err := m.AvailableCycles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availableCycles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("availableCycles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this frequency available model based on the context it is used
func (m *FrequencyAvailableModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvailableCycles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FrequencyAvailableModel) contextValidateAvailableCycles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AvailableCycles); i++ {

		if m.AvailableCycles[i] != nil {
			if err := m.AvailableCycles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availableCycles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("availableCycles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FrequencyAvailableModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrequencyAvailableModel) UnmarshalBinary(b []byte) error {
	var res FrequencyAvailableModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
