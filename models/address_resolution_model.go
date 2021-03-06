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

// AddressResolutionModel Address Resolution Model
// Example: {"address":{"city":"Irvine","country":"US","line1":"2000 Main Street","postalCode":"92614","region":"CA"},"coordinates":{"latitude":33.684884,"longitude":-117.851321},"messages":[{"details":"The physical location exists but there are no homes on this street. One reason might be railroad tracks or rivers running alongside this street, as they would prevent construction of homes in this location.","refersTo":"address","severity":"Error","source":"Avalara.AvaTax.Services.Address","summary":"The address is not deliverable."}],"resolutionQuality":"Intersection","taxAuthorities":[{"avalaraId":"AGAM","jurisdictionName":"CALIFORNIA","jurisdictionType":"State","signatureCode":"AGAM"}]}
//
// swagger:model AddressResolutionModel
type AddressResolutionModel struct {

	// The original address
	// Example: {"city":"Irvine","country":"US","line1":"2000 Main Street","postalCode":"92614","region":"CA"}
	Address *AddressInfo `json:"address,omitempty"`

	// The geospatial coordinates of this address
	// Example: {"latitude":33.684884,"longitude":-117.851321}
	Coordinates *CoordinateInfo `json:"coordinates,omitempty"`

	// List of informational and warning messages regarding this address
	// Example: [{"details":"The physical location exists but there are no homes on this street. One reason might be railroad tracks or rivers running alongside this street, as they would prevent construction of homes in this location.","refersTo":"address","severity":"Error","source":"Avalara.AvaTax.Services.Address","summary":"The address is not deliverable."}]
	Messages []*AvaTaxMessage `json:"messages"`

	// The resolution quality of the geospatial coordinates
	// Example: NotCoded
	// Enum: [NotCoded External CountryCentroid RegionCentroid PartialCentroid PostalCentroidGood PostalCentroidBetter PostalCentroidBest Intersection Interpolated Rooftop Constant]
	ResolutionQuality string `json:"resolutionQuality,omitempty"`

	// List of informational and warning messages regarding this address
	// Example: [{"avalaraId":"AGAM","jurisdictionName":"CALIFORNIA","jurisdictionType":"State","signatureCode":"AGAM"}]
	TaxAuthorities []*TaxAuthorityInfo `json:"taxAuthorities"`

	// The validated address or addresses
	ValidatedAddresses []*ValidatedAddressInfo `json:"validatedAddresses"`
}

// Validate validates this address resolution model
func (m *AddressResolutionModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoordinates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolutionQuality(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaxAuthorities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidatedAddresses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddressResolutionModel) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *AddressResolutionModel) validateCoordinates(formats strfmt.Registry) error {
	if swag.IsZero(m.Coordinates) { // not required
		return nil
	}

	if m.Coordinates != nil {
		if err := m.Coordinates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coordinates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("coordinates")
			}
			return err
		}
	}

	return nil
}

func (m *AddressResolutionModel) validateMessages(formats strfmt.Registry) error {
	if swag.IsZero(m.Messages) { // not required
		return nil
	}

	for i := 0; i < len(m.Messages); i++ {
		if swag.IsZero(m.Messages[i]) { // not required
			continue
		}

		if m.Messages[i] != nil {
			if err := m.Messages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("messages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var addressResolutionModelTypeResolutionQualityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NotCoded","External","CountryCentroid","RegionCentroid","PartialCentroid","PostalCentroidGood","PostalCentroidBetter","PostalCentroidBest","Intersection","Interpolated","Rooftop","Constant"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addressResolutionModelTypeResolutionQualityPropEnum = append(addressResolutionModelTypeResolutionQualityPropEnum, v)
	}
}

const (

	// AddressResolutionModelResolutionQualityNotCoded captures enum value "NotCoded"
	AddressResolutionModelResolutionQualityNotCoded string = "NotCoded"

	// AddressResolutionModelResolutionQualityExternal captures enum value "External"
	AddressResolutionModelResolutionQualityExternal string = "External"

	// AddressResolutionModelResolutionQualityCountryCentroid captures enum value "CountryCentroid"
	AddressResolutionModelResolutionQualityCountryCentroid string = "CountryCentroid"

	// AddressResolutionModelResolutionQualityRegionCentroid captures enum value "RegionCentroid"
	AddressResolutionModelResolutionQualityRegionCentroid string = "RegionCentroid"

	// AddressResolutionModelResolutionQualityPartialCentroid captures enum value "PartialCentroid"
	AddressResolutionModelResolutionQualityPartialCentroid string = "PartialCentroid"

	// AddressResolutionModelResolutionQualityPostalCentroidGood captures enum value "PostalCentroidGood"
	AddressResolutionModelResolutionQualityPostalCentroidGood string = "PostalCentroidGood"

	// AddressResolutionModelResolutionQualityPostalCentroidBetter captures enum value "PostalCentroidBetter"
	AddressResolutionModelResolutionQualityPostalCentroidBetter string = "PostalCentroidBetter"

	// AddressResolutionModelResolutionQualityPostalCentroidBest captures enum value "PostalCentroidBest"
	AddressResolutionModelResolutionQualityPostalCentroidBest string = "PostalCentroidBest"

	// AddressResolutionModelResolutionQualityIntersection captures enum value "Intersection"
	AddressResolutionModelResolutionQualityIntersection string = "Intersection"

	// AddressResolutionModelResolutionQualityInterpolated captures enum value "Interpolated"
	AddressResolutionModelResolutionQualityInterpolated string = "Interpolated"

	// AddressResolutionModelResolutionQualityRooftop captures enum value "Rooftop"
	AddressResolutionModelResolutionQualityRooftop string = "Rooftop"

	// AddressResolutionModelResolutionQualityConstant captures enum value "Constant"
	AddressResolutionModelResolutionQualityConstant string = "Constant"
)

// prop value enum
func (m *AddressResolutionModel) validateResolutionQualityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addressResolutionModelTypeResolutionQualityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AddressResolutionModel) validateResolutionQuality(formats strfmt.Registry) error {
	if swag.IsZero(m.ResolutionQuality) { // not required
		return nil
	}

	// value enum
	if err := m.validateResolutionQualityEnum("resolutionQuality", "body", m.ResolutionQuality); err != nil {
		return err
	}

	return nil
}

func (m *AddressResolutionModel) validateTaxAuthorities(formats strfmt.Registry) error {
	if swag.IsZero(m.TaxAuthorities) { // not required
		return nil
	}

	for i := 0; i < len(m.TaxAuthorities); i++ {
		if swag.IsZero(m.TaxAuthorities[i]) { // not required
			continue
		}

		if m.TaxAuthorities[i] != nil {
			if err := m.TaxAuthorities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taxAuthorities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taxAuthorities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AddressResolutionModel) validateValidatedAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.ValidatedAddresses) { // not required
		return nil
	}

	for i := 0; i < len(m.ValidatedAddresses); i++ {
		if swag.IsZero(m.ValidatedAddresses[i]) { // not required
			continue
		}

		if m.ValidatedAddresses[i] != nil {
			if err := m.ValidatedAddresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validatedAddresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validatedAddresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this address resolution model based on the context it is used
func (m *AddressResolutionModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCoordinates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaxAuthorities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValidatedAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddressResolutionModel) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.Address != nil {
		if err := m.Address.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *AddressResolutionModel) contextValidateCoordinates(ctx context.Context, formats strfmt.Registry) error {

	if m.Coordinates != nil {
		if err := m.Coordinates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coordinates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("coordinates")
			}
			return err
		}
	}

	return nil
}

func (m *AddressResolutionModel) contextValidateMessages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Messages); i++ {

		if m.Messages[i] != nil {
			if err := m.Messages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("messages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AddressResolutionModel) contextValidateTaxAuthorities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TaxAuthorities); i++ {

		if m.TaxAuthorities[i] != nil {
			if err := m.TaxAuthorities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("taxAuthorities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("taxAuthorities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AddressResolutionModel) contextValidateValidatedAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ValidatedAddresses); i++ {

		if m.ValidatedAddresses[i] != nil {
			if err := m.ValidatedAddresses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validatedAddresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validatedAddresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddressResolutionModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddressResolutionModel) UnmarshalBinary(b []byte) error {
	var res AddressResolutionModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
