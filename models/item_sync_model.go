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

// ItemSyncModel An abridged item model used for syncing product catalogs with AvaTax.
// Example: {"description":"Potato Chips","itemCode":"CHIP1041","itemGroup":"Packaged Food","taxCode":"PF051578"}
//
// swagger:model ItemSyncModel
type ItemSyncModel struct {

	// A friendly description of the item. If your company has enrolled in Streamlined Sales Tax, this description must be auditable.
	// Example: Potato Chips
	// Required: true
	Description *string `json:"description"`

	// A unique code representing this item.
	// Example: CHIP1041
	// Required: true
	ItemCode *string `json:"itemCode"`

	// A group to which the item belongs.
	// Example: Packaged Food
	ItemGroup string `json:"itemGroup,omitempty"`

	// The tax code of the item (optional)
	// Example: PF051578
	TaxCode string `json:"taxCode,omitempty"`
}

// Validate validates this item sync model
func (m *ItemSyncModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemSyncModel) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ItemSyncModel) validateItemCode(formats strfmt.Registry) error {

	if err := validate.Required("itemCode", "body", m.ItemCode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this item sync model based on context it is used
func (m *ItemSyncModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ItemSyncModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemSyncModel) UnmarshalBinary(b []byte) error {
	var res ItemSyncModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
