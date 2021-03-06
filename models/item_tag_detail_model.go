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

// ItemTagDetailModel Represents a tag for an item in your company's product catalog.
// Example: {"companyId":8010687,"createdDate":"0001-01-01T00:00:00","itemId":151284220,"itemTagDetailId":51,"tagName":"CLASSIFICATION_INITIATIONERROR"}
//
// swagger:model ItemTagDetailModel
type ItemTagDetailModel struct {

	// The unique ID number of the company that owns this item.
	// Example: 8010687
	CompanyID int32 `json:"companyId,omitempty"`

	// The date when this record was created.
	// Example: 0001-01-01T00:00:00
	// Read Only: true
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// The unique ID number of this item.
	// Example: 151284220
	ItemID int64 `json:"itemId,omitempty"`

	// The unique ID number of the item-tag relation.
	// Example: 51
	// Read Only: true
	ItemTagDetailID int32 `json:"itemTagDetailId,omitempty"`

	// The unique tag Id for the tags.
	// Read Only: true
	TagID int32 `json:"tagId,omitempty"`

	// The tag name.
	// Example: CLASSIFICATION_INITIATIONERROR
	// Required: true
	TagName *string `json:"tagName"`
}

// Validate validates this item tag detail model
func (m *ItemTagDetailModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemTagDetailModel) validateCreatedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ItemTagDetailModel) validateTagName(formats strfmt.Registry) error {

	if err := validate.Required("tagName", "body", m.TagName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this item tag detail model based on the context it is used
func (m *ItemTagDetailModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedDate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateItemTagDetailID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTagID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ItemTagDetailModel) contextValidateCreatedDate(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdDate", "body", strfmt.DateTime(m.CreatedDate)); err != nil {
		return err
	}

	return nil
}

func (m *ItemTagDetailModel) contextValidateItemTagDetailID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "itemTagDetailId", "body", int32(m.ItemTagDetailID)); err != nil {
		return err
	}

	return nil
}

func (m *ItemTagDetailModel) contextValidateTagID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "tagId", "body", int32(m.TagID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ItemTagDetailModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemTagDetailModel) UnmarshalBinary(b []byte) error {
	var res ItemTagDetailModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
