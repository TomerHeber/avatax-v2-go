// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountModel An AvaTax account.
// Example: {"accountStatusId":"Test","accountTypeId":"Regular","effectiveDate":"2021-12-25T00:00:00+00:00","id":12345,"isSamlEnabled":false,"name":"Test Account"}
//
// swagger:model AccountModel
type AccountModel struct {

	// The current status of this account.
	// Example: Inactive
	// Enum: [Inactive Active Test New]
	AccountStatusID string `json:"accountStatusId,omitempty"`

	// The type of this account.
	// Example: Regular
	// Enum: [Regular Firm FirmClient]
	AccountTypeID string `json:"accountTypeId,omitempty"`

	// The date when this record was created.
	// Read Only: true
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// The User ID of the user who created this record.
	// Read Only: true
	CreatedUserID int32 `json:"createdUserId,omitempty"`

	// For system registrar use only.
	// Max Length: 100
	// Min Length: 0
	Crmid *string `json:"crmid,omitempty"`

	// The earliest date on which this account may be used.
	// Example: 2021-12-25T00:00:00+00:00
	// Format: date-time
	EffectiveDate strfmt.DateTime `json:"effectiveDate,omitempty"`

	// If this account has been closed, this is the last date the account was open.
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// The unique ID number assigned to this account.
	// Example: 12345
	// Required: true
	ID *int32 `json:"id"`

	// Is Saml based authentication used by this account for user to login via AI.
	// Example: false
	IsSamlEnabled bool `json:"isSamlEnabled,omitempty"`

	// The date/time when this record was last modified.
	// Read Only: true
	// Format: date-time
	ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

	// The user ID of the user who last modified this record.
	// Read Only: true
	ModifiedUserID int32 `json:"modifiedUserId,omitempty"`

	// The name of this account.
	// Example: Test Account
	// Required: true
	// Max Length: 50
	// Min Length: 0
	Name *string `json:"name"`

	// Optional: A list of subscriptions granted to this account.  To fetch this list, add the query string "?$include=Subscriptions" to your URL.
	Subscriptions []*SubscriptionModel `json:"subscriptions"`

	// Optional: A list of all the users belonging to this account.  To fetch this list, add the query string "?$include=Users" to your URL.
	Users []*UserModel `json:"users"`
}

// Validate validates this account model
func (m *AccountModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountStatusID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrmid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffectiveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var accountModelTypeAccountStatusIDPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Inactive","Active","Test","New"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountModelTypeAccountStatusIDPropEnum = append(accountModelTypeAccountStatusIDPropEnum, v)
	}
}

const (

	// AccountModelAccountStatusIDInactive captures enum value "Inactive"
	AccountModelAccountStatusIDInactive string = "Inactive"

	// AccountModelAccountStatusIDActive captures enum value "Active"
	AccountModelAccountStatusIDActive string = "Active"

	// AccountModelAccountStatusIDTest captures enum value "Test"
	AccountModelAccountStatusIDTest string = "Test"

	// AccountModelAccountStatusIDNew captures enum value "New"
	AccountModelAccountStatusIDNew string = "New"
)

// prop value enum
func (m *AccountModel) validateAccountStatusIDEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountModelTypeAccountStatusIDPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountModel) validateAccountStatusID(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountStatusID) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccountStatusIDEnum("accountStatusId", "body", m.AccountStatusID); err != nil {
		return err
	}

	return nil
}

var accountModelTypeAccountTypeIDPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Regular","Firm","FirmClient"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountModelTypeAccountTypeIDPropEnum = append(accountModelTypeAccountTypeIDPropEnum, v)
	}
}

const (

	// AccountModelAccountTypeIDRegular captures enum value "Regular"
	AccountModelAccountTypeIDRegular string = "Regular"

	// AccountModelAccountTypeIDFirm captures enum value "Firm"
	AccountModelAccountTypeIDFirm string = "Firm"

	// AccountModelAccountTypeIDFirmClient captures enum value "FirmClient"
	AccountModelAccountTypeIDFirmClient string = "FirmClient"
)

// prop value enum
func (m *AccountModel) validateAccountTypeIDEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountModelTypeAccountTypeIDPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountModel) validateAccountTypeID(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountTypeID) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccountTypeIDEnum("accountTypeId", "body", m.AccountTypeID); err != nil {
		return err
	}

	return nil
}

func (m *AccountModel) validateCreatedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountModel) validateCrmid(formats strfmt.Registry) error {
	if swag.IsZero(m.Crmid) { // not required
		return nil
	}

	if err := validate.MinLength("crmid", "body", *m.Crmid, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("crmid", "body", *m.Crmid, 100); err != nil {
		return err
	}

	return nil
}

func (m *AccountModel) validateEffectiveDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EffectiveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("effectiveDate", "body", "date-time", m.EffectiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountModel) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountModel) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *AccountModel) validateModifiedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountModel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 50); err != nil {
		return err
	}

	return nil
}

func (m *AccountModel) validateSubscriptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Subscriptions) { // not required
		return nil
	}

	for i := 0; i < len(m.Subscriptions); i++ {
		if swag.IsZero(m.Subscriptions[i]) { // not required
			continue
		}

		if m.Subscriptions[i] != nil {
			if err := m.Subscriptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subscriptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subscriptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountModel) validateUsers(formats strfmt.Registry) error {
	if swag.IsZero(m.Users) { // not required
		return nil
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this account model based on the context it is used
func (m *AccountModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubscriptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountModel) contextValidateCreatedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdDate", "body", strfmt.DateTime(m.CreatedDate)); err != nil {
		return err
	}

	return nil
}

func (m *AccountModel) contextValidateCreatedUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdUserId", "body", int32(m.CreatedUserID)); err != nil {
		return err
	}

	return nil
}

func (m *AccountModel) contextValidateModifiedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedDate", "body", strfmt.DateTime(m.ModifiedDate)); err != nil {
		return err
	}

	return nil
}

func (m *AccountModel) contextValidateModifiedUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedUserId", "body", int32(m.ModifiedUserID)); err != nil {
		return err
	}

	return nil
}

func (m *AccountModel) contextValidateSubscriptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Subscriptions); i++ {

		if m.Subscriptions[i] != nil {
			if err := m.Subscriptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subscriptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subscriptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountModel) contextValidateUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Users); i++ {

		if m.Users[i] != nil {
			if err := m.Users[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountModel) UnmarshalBinary(b []byte) error {
	var res AccountModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}