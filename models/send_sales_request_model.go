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

// SendSalesRequestModel SendSales Request Model.
// Example: {"companyId":12345,"date":"2021-12-25T00:00:00+00:00","format":"STANDARD","taxCodes":["P0000000","PC040100"],"type":"Csv"}
//
// swagger:model SendSalesRequestModel
type SendSalesRequestModel struct {

	// The companyId for which the send sales file is being generated.
	// Example: 12345
	// Required: true
	CompanyID *int32 `json:"companyId"`

	// The date for which send sales file is being generated.
	// Example: 2021-12-25T00:00:00+00:00
	// Required: true
	// Format: date-time
	Date *strfmt.DateTime `json:"date"`

	// The send sales file format.
	// Example: STANDARD
	// Enum: [STANDARD DMA DMA_NEW]
	Format string `json:"format,omitempty"`

	// List of taxCodes to be included in send sales file.
	// Example: ["P0000000","PC040100"]
	// Required: true
	TaxCodes []string `json:"taxCodes"`

	// The send sales file type
	// Example: Csv
	// Enum: [Csv Json]
	Type string `json:"type,omitempty"`
}

// Validate validates this send sales request model
func (m *SendSalesRequestModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompanyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxCodes(formats); err != nil {
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

func (m *SendSalesRequestModel) validateCompanyID(formats strfmt.Registry) error {

	if err := validate.Required("companyId", "body", m.CompanyID); err != nil {
		return err
	}

	return nil
}

func (m *SendSalesRequestModel) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

var sendSalesRequestModelTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STANDARD","DMA","DMA_NEW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sendSalesRequestModelTypeFormatPropEnum = append(sendSalesRequestModelTypeFormatPropEnum, v)
	}
}

const (

	// SendSalesRequestModelFormatSTANDARD captures enum value "STANDARD"
	SendSalesRequestModelFormatSTANDARD string = "STANDARD"

	// SendSalesRequestModelFormatDMA captures enum value "DMA"
	SendSalesRequestModelFormatDMA string = "DMA"

	// SendSalesRequestModelFormatDMANEW captures enum value "DMA_NEW"
	SendSalesRequestModelFormatDMANEW string = "DMA_NEW"
)

// prop value enum
func (m *SendSalesRequestModel) validateFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sendSalesRequestModelTypeFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SendSalesRequestModel) validateFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

func (m *SendSalesRequestModel) validateTaxCodes(formats strfmt.Registry) error {

	if err := validate.Required("taxCodes", "body", m.TaxCodes); err != nil {
		return err
	}

	return nil
}

var sendSalesRequestModelTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Csv","Json"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sendSalesRequestModelTypeTypePropEnum = append(sendSalesRequestModelTypeTypePropEnum, v)
	}
}

const (

	// SendSalesRequestModelTypeCsv captures enum value "Csv"
	SendSalesRequestModelTypeCsv string = "Csv"

	// SendSalesRequestModelTypeJSON captures enum value "Json"
	SendSalesRequestModelTypeJSON string = "Json"
)

// prop value enum
func (m *SendSalesRequestModel) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, sendSalesRequestModelTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SendSalesRequestModel) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this send sales request model based on context it is used
func (m *SendSalesRequestModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SendSalesRequestModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendSalesRequestModel) UnmarshalBinary(b []byte) error {
	var res SendSalesRequestModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}