// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FilingAdjustmentModel A model for return adjustments.
// Example: {"accountType":"AccountsReceivableAccountsPayable","amount":100,"id":15634,"isCalculated":false,"period":"CurrentPeriod","type":"Discount"}
//
// swagger:model FilingAdjustmentModel
type FilingAdjustmentModel struct {

	// The account type of the adjustment.
	// Example: None
	// Required: true
	// Enum: [None AccountsReceivableAccountsPayable AccountsReceivable AccountsPayable]
	AccountType *string `json:"accountType"`

	// The adjustment amount.
	// Example: 100
	// Required: true
	Amount *float64 `json:"amount"`

	// The date when this record was created.
	// Read Only: true
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// The User ID of the user who created this record.
	// Read Only: true
	CreatedUserID int32 `json:"createdUserId,omitempty"`

	// The filing return id that this applies too
	// Read Only: true
	FilingID int64 `json:"filingId,omitempty"`

	// The unique ID number for the adjustment.
	// Example: 15634
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Whether or not the adjustment has been calculated.
	// Example: false
	IsCalculated bool `json:"isCalculated,omitempty"`

	// The date/time when this record was last modified.
	// Read Only: true
	// Format: date-time
	ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

	// The user ID of the user who last modified this record.
	// Read Only: true
	ModifiedUserID int32 `json:"modifiedUserId,omitempty"`

	// The filing period the adjustment is applied to.
	// Example: None
	// Required: true
	// Enum: [None CurrentPeriod NextPeriod]
	Period *string `json:"period"`

	// A descriptive reason for creating this adjustment.
	Reason string `json:"reason,omitempty"`

	// The type of the adjustment.
	// Example: Discount
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this filing adjustment model
func (m *FilingAdjustmentModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var filingAdjustmentModelTypeAccountTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","AccountsReceivableAccountsPayable","AccountsReceivable","AccountsPayable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		filingAdjustmentModelTypeAccountTypePropEnum = append(filingAdjustmentModelTypeAccountTypePropEnum, v)
	}
}

const (

	// FilingAdjustmentModelAccountTypeNone captures enum value "None"
	FilingAdjustmentModelAccountTypeNone string = "None"

	// FilingAdjustmentModelAccountTypeAccountsReceivableAccountsPayable captures enum value "AccountsReceivableAccountsPayable"
	FilingAdjustmentModelAccountTypeAccountsReceivableAccountsPayable string = "AccountsReceivableAccountsPayable"

	// FilingAdjustmentModelAccountTypeAccountsReceivable captures enum value "AccountsReceivable"
	FilingAdjustmentModelAccountTypeAccountsReceivable string = "AccountsReceivable"

	// FilingAdjustmentModelAccountTypeAccountsPayable captures enum value "AccountsPayable"
	FilingAdjustmentModelAccountTypeAccountsPayable string = "AccountsPayable"
)

// prop value enum
func (m *FilingAdjustmentModel) validateAccountTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, filingAdjustmentModelTypeAccountTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FilingAdjustmentModel) validateAccountType(formats strfmt.Registry) error {

	if err := validate.Required("accountType", "body", m.AccountType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAccountTypeEnum("accountType", "body", *m.AccountType); err != nil {
		return err
	}

	return nil
}

func (m *FilingAdjustmentModel) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *FilingAdjustmentModel) validateCreatedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FilingAdjustmentModel) validateModifiedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var filingAdjustmentModelTypePeriodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","CurrentPeriod","NextPeriod"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		filingAdjustmentModelTypePeriodPropEnum = append(filingAdjustmentModelTypePeriodPropEnum, v)
	}
}

const (

	// FilingAdjustmentModelPeriodNone captures enum value "None"
	FilingAdjustmentModelPeriodNone string = "None"

	// FilingAdjustmentModelPeriodCurrentPeriod captures enum value "CurrentPeriod"
	FilingAdjustmentModelPeriodCurrentPeriod string = "CurrentPeriod"

	// FilingAdjustmentModelPeriodNextPeriod captures enum value "NextPeriod"
	FilingAdjustmentModelPeriodNextPeriod string = "NextPeriod"
)

// prop value enum
func (m *FilingAdjustmentModel) validatePeriodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, filingAdjustmentModelTypePeriodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FilingAdjustmentModel) validatePeriod(formats strfmt.Registry) error {

	if err := validate.Required("period", "body", m.Period); err != nil {
		return err
	}

	// value enum
	if err := m.validatePeriodEnum("period", "body", *m.Period); err != nil {
		return err
	}

	return nil
}

func (m *FilingAdjustmentModel) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this filing adjustment model based on the context it is used
func (m *FilingAdjustmentModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilingID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilingAdjustmentModel) contextValidateCreatedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdDate", "body", strfmt.DateTime(m.CreatedDate)); err != nil {
		return err
	}

	return nil
}

func (m *FilingAdjustmentModel) contextValidateCreatedUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdUserId", "body", int32(m.CreatedUserID)); err != nil {
		return err
	}

	return nil
}

func (m *FilingAdjustmentModel) contextValidateFilingID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "filingId", "body", int64(m.FilingID)); err != nil {
		return err
	}

	return nil
}

func (m *FilingAdjustmentModel) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *FilingAdjustmentModel) contextValidateModifiedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedDate", "body", strfmt.DateTime(m.ModifiedDate)); err != nil {
		return err
	}

	return nil
}

func (m *FilingAdjustmentModel) contextValidateModifiedUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedUserId", "body", int32(m.ModifiedUserID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilingAdjustmentModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilingAdjustmentModel) UnmarshalBinary(b []byte) error {
	var res FilingAdjustmentModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
