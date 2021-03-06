// Code generated by go-swagger; DO NOT EDIT.

package definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// ListRateTypesByCountryTaxTypeTaxSubTypeReader is a Reader for the ListRateTypesByCountryTaxTypeTaxSubType structure.
type ListRateTypesByCountryTaxTypeTaxSubTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRateTypesByCountryTaxTypeTaxSubTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRateTypesByCountryTaxTypeTaxSubTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRateTypesByCountryTaxTypeTaxSubTypeOK creates a ListRateTypesByCountryTaxTypeTaxSubTypeOK with default headers values
func NewListRateTypesByCountryTaxTypeTaxSubTypeOK() *ListRateTypesByCountryTaxTypeTaxSubTypeOK {
	return &ListRateTypesByCountryTaxTypeTaxSubTypeOK{}
}

/* ListRateTypesByCountryTaxTypeTaxSubTypeOK describes a response with status code 200, with default header values.

Success
*/
type ListRateTypesByCountryTaxTypeTaxSubTypeOK struct {
	Payload *models.RateTypesModelFetchResult
}

func (o *ListRateTypesByCountryTaxTypeTaxSubTypeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/definitions/countries/{country}/taxtypes/{taxTypeId}/taxsubtypes/{taxSubTypeId}/ratetypes][%d] listRateTypesByCountryTaxTypeTaxSubTypeOK  %+v", 200, o.Payload)
}
func (o *ListRateTypesByCountryTaxTypeTaxSubTypeOK) GetPayload() *models.RateTypesModelFetchResult {
	return o.Payload
}

func (o *ListRateTypesByCountryTaxTypeTaxSubTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RateTypesModelFetchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
