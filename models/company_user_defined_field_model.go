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

// CompanyUserDefinedFieldModel company user defined field model
// Example: {"companyId":12345,"createdDate":"2021-12-25T00:00:00+00:00","createdUserId":123,"dataType":"String","friendlyName":"General Ledger Account Number","id":1,"modifiedDate":"2021-12-25T00:00:00+00:00","modifiedUserId":12,"name":"UDF1","userDefinedFieldType":"Document"}
//
// swagger:model CompanyUserDefinedFieldModel
type CompanyUserDefinedFieldModel struct {

	// The id of the company to which the datasource belongs to.
	// Example: 12345
	CompanyID int32 `json:"companyId,omitempty"`

	// The date when this record was created.
	// Example: 2021-12-25T00:00:00+00:00
	// Read Only: true
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// The User ID of the user who created this record.
	// Example: 123
	// Read Only: true
	CreatedUserID int32 `json:"createdUserId,omitempty"`

	// The unique ID number of this connection.
	// Example: String
	// Enum: [String Number Date Boolean]
	DataType string `json:"dataType,omitempty"`

	// The unique ID number of this connection.
	// Example: General Ledger Account Number
	// Required: true
	// Max Length: 4096
	// Min Length: 0
	FriendlyName *string `json:"friendlyName"`

	// The id of the datasource.
	// Example: 1
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// The date/time when this record was last modified.
	// Example: 2021-12-25T00:00:00+00:00
	// Read Only: true
	// Format: date-time
	ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

	// The user ID of the user who last modified this record.
	// Example: 12
	// Read Only: true
	ModifiedUserID int32 `json:"modifiedUserId,omitempty"`

	// The extractor/connector id.
	// Example: UDF1
	// Required: true
	// Max Length: 50
	// Min Length: 0
	Name *string `json:"name"`

	// The category of user defined type For Example: Document level or Line level UDF.
	// Example: Document
	// Enum: [Document Line]
	UserDefinedFieldType string `json:"userDefinedFieldType,omitempty"`
}

// Validate validates this company user defined field model
func (m *CompanyUserDefinedFieldModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFriendlyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserDefinedFieldType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CompanyUserDefinedFieldModel) validateCreatedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var companyUserDefinedFieldModelTypeDataTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["String","Number","Date","Boolean"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		companyUserDefinedFieldModelTypeDataTypePropEnum = append(companyUserDefinedFieldModelTypeDataTypePropEnum, v)
	}
}

const (

	// CompanyUserDefinedFieldModelDataTypeString captures enum value "String"
	CompanyUserDefinedFieldModelDataTypeString string = "String"

	// CompanyUserDefinedFieldModelDataTypeNumber captures enum value "Number"
	CompanyUserDefinedFieldModelDataTypeNumber string = "Number"

	// CompanyUserDefinedFieldModelDataTypeDate captures enum value "Date"
	CompanyUserDefinedFieldModelDataTypeDate string = "Date"

	// CompanyUserDefinedFieldModelDataTypeBoolean captures enum value "Boolean"
	CompanyUserDefinedFieldModelDataTypeBoolean string = "Boolean"
)

// prop value enum
func (m *CompanyUserDefinedFieldModel) validateDataTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, companyUserDefinedFieldModelTypeDataTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CompanyUserDefinedFieldModel) validateDataType(formats strfmt.Registry) error {
	if swag.IsZero(m.DataType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataTypeEnum("dataType", "body", m.DataType); err != nil {
		return err
	}

	return nil
}

func (m *CompanyUserDefinedFieldModel) validateFriendlyName(formats strfmt.Registry) error {

	if err := validate.Required("friendlyName", "body", m.FriendlyName); err != nil {
		return err
	}

	if err := validate.MinLength("friendlyName", "body", *m.FriendlyName, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("friendlyName", "body", *m.FriendlyName, 4096); err != nil {
		return err
	}

	return nil
}

func (m *CompanyUserDefinedFieldModel) validateModifiedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CompanyUserDefinedFieldModel) validateName(formats strfmt.Registry) error {

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

var companyUserDefinedFieldModelTypeUserDefinedFieldTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Document","Line"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		companyUserDefinedFieldModelTypeUserDefinedFieldTypePropEnum = append(companyUserDefinedFieldModelTypeUserDefinedFieldTypePropEnum, v)
	}
}

const (

	// CompanyUserDefinedFieldModelUserDefinedFieldTypeDocument captures enum value "Document"
	CompanyUserDefinedFieldModelUserDefinedFieldTypeDocument string = "Document"

	// CompanyUserDefinedFieldModelUserDefinedFieldTypeLine captures enum value "Line"
	CompanyUserDefinedFieldModelUserDefinedFieldTypeLine string = "Line"
)

// prop value enum
func (m *CompanyUserDefinedFieldModel) validateUserDefinedFieldTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, companyUserDefinedFieldModelTypeUserDefinedFieldTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CompanyUserDefinedFieldModel) validateUserDefinedFieldType(formats strfmt.Registry) error {
	if swag.IsZero(m.UserDefinedFieldType) { // not required
		return nil
	}

	// value enum
	if err := m.validateUserDefinedFieldTypeEnum("userDefinedFieldType", "body", m.UserDefinedFieldType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this company user defined field model based on the context it is used
func (m *CompanyUserDefinedFieldModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedUserID(ctx, formats); err != nil {
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

func (m *CompanyUserDefinedFieldModel) contextValidateCreatedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdDate", "body", strfmt.DateTime(m.CreatedDate)); err != nil {
		return err
	}

	return nil
}

func (m *CompanyUserDefinedFieldModel) contextValidateCreatedUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdUserId", "body", int32(m.CreatedUserID)); err != nil {
		return err
	}

	return nil
}

func (m *CompanyUserDefinedFieldModel) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CompanyUserDefinedFieldModel) contextValidateModifiedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedDate", "body", strfmt.DateTime(m.ModifiedDate)); err != nil {
		return err
	}

	return nil
}

func (m *CompanyUserDefinedFieldModel) contextValidateModifiedUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedUserId", "body", int32(m.ModifiedUserID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CompanyUserDefinedFieldModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CompanyUserDefinedFieldModel) UnmarshalBinary(b []byte) error {
	var res CompanyUserDefinedFieldModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
