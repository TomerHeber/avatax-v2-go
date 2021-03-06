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

// CommitMultiDocumentModel Commit this MultiDocument object so that all transactions within it can be reported on a tax filing.
// Example: {"code":"3ac79777-3281-4dcb-9f71-eb0c023165d1","commit":true,"type":"SalesInvoice"}
//
// swagger:model CommitMultiDocumentModel
type CommitMultiDocumentModel struct {

	// Represents the unique code of this MultiDocument transaction.
	//
	// A MultiDocument transaction is uniquely identified by its `accountId`, `code`, and `type`.        ///
	// Example: 3ac79777-3281-4dcb-9f71-eb0c023165d1
	// Required: true
	Code *string `json:"code"`

	// Set this value to be `true` to commit this transaction.
	//
	// Committing a transaction allows it to be reported on a tax filing.  Uncommitted transactions will not be reported.
	// Example: true
	// Required: true
	Commit *bool `json:"commit"`

	// Represents the document type of this MultiDocument transaction.  For more information about
	// document types, see [DocumentType](https://developer.avalara.com/api-reference/avatax/rest/v2/models/enums/DocumentType/).
	//
	// A MultiDocument transaction is uniquely identified by its `accountId`, `code`, and `type`.
	// Example: SalesOrder
	// Enum: [SalesOrder SalesInvoice PurchaseOrder PurchaseInvoice ReturnOrder ReturnInvoice InventoryTransferOrder InventoryTransferInvoice ReverseChargeOrder ReverseChargeInvoice CustomsInvoice CustomsOrder Any]
	Type string `json:"type,omitempty"`
}

// Validate validates this commit multi document model
func (m *CommitMultiDocumentModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommit(formats); err != nil {
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

func (m *CommitMultiDocumentModel) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *CommitMultiDocumentModel) validateCommit(formats strfmt.Registry) error {

	if err := validate.Required("commit", "body", m.Commit); err != nil {
		return err
	}

	return nil
}

var commitMultiDocumentModelTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SalesOrder","SalesInvoice","PurchaseOrder","PurchaseInvoice","ReturnOrder","ReturnInvoice","InventoryTransferOrder","InventoryTransferInvoice","ReverseChargeOrder","ReverseChargeInvoice","CustomsInvoice","CustomsOrder","Any"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commitMultiDocumentModelTypeTypePropEnum = append(commitMultiDocumentModelTypeTypePropEnum, v)
	}
}

const (

	// CommitMultiDocumentModelTypeSalesOrder captures enum value "SalesOrder"
	CommitMultiDocumentModelTypeSalesOrder string = "SalesOrder"

	// CommitMultiDocumentModelTypeSalesInvoice captures enum value "SalesInvoice"
	CommitMultiDocumentModelTypeSalesInvoice string = "SalesInvoice"

	// CommitMultiDocumentModelTypePurchaseOrder captures enum value "PurchaseOrder"
	CommitMultiDocumentModelTypePurchaseOrder string = "PurchaseOrder"

	// CommitMultiDocumentModelTypePurchaseInvoice captures enum value "PurchaseInvoice"
	CommitMultiDocumentModelTypePurchaseInvoice string = "PurchaseInvoice"

	// CommitMultiDocumentModelTypeReturnOrder captures enum value "ReturnOrder"
	CommitMultiDocumentModelTypeReturnOrder string = "ReturnOrder"

	// CommitMultiDocumentModelTypeReturnInvoice captures enum value "ReturnInvoice"
	CommitMultiDocumentModelTypeReturnInvoice string = "ReturnInvoice"

	// CommitMultiDocumentModelTypeInventoryTransferOrder captures enum value "InventoryTransferOrder"
	CommitMultiDocumentModelTypeInventoryTransferOrder string = "InventoryTransferOrder"

	// CommitMultiDocumentModelTypeInventoryTransferInvoice captures enum value "InventoryTransferInvoice"
	CommitMultiDocumentModelTypeInventoryTransferInvoice string = "InventoryTransferInvoice"

	// CommitMultiDocumentModelTypeReverseChargeOrder captures enum value "ReverseChargeOrder"
	CommitMultiDocumentModelTypeReverseChargeOrder string = "ReverseChargeOrder"

	// CommitMultiDocumentModelTypeReverseChargeInvoice captures enum value "ReverseChargeInvoice"
	CommitMultiDocumentModelTypeReverseChargeInvoice string = "ReverseChargeInvoice"

	// CommitMultiDocumentModelTypeCustomsInvoice captures enum value "CustomsInvoice"
	CommitMultiDocumentModelTypeCustomsInvoice string = "CustomsInvoice"

	// CommitMultiDocumentModelTypeCustomsOrder captures enum value "CustomsOrder"
	CommitMultiDocumentModelTypeCustomsOrder string = "CustomsOrder"

	// CommitMultiDocumentModelTypeAny captures enum value "Any"
	CommitMultiDocumentModelTypeAny string = "Any"
)

// prop value enum
func (m *CommitMultiDocumentModel) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, commitMultiDocumentModelTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CommitMultiDocumentModel) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this commit multi document model based on context it is used
func (m *CommitMultiDocumentModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommitMultiDocumentModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommitMultiDocumentModel) UnmarshalBinary(b []byte) error {
	var res CommitMultiDocumentModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
