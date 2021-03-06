// Code generated by go-swagger; DO NOT EDIT.

package tax_codes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// GetTaxCodeReader is a Reader for the GetTaxCode structure.
type GetTaxCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaxCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaxCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTaxCodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTaxCodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTaxCodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTaxCodeOK creates a GetTaxCodeOK with default headers values
func NewGetTaxCodeOK() *GetTaxCodeOK {
	return &GetTaxCodeOK{}
}

/* GetTaxCodeOK describes a response with status code 200, with default header values.

Success
*/
type GetTaxCodeOK struct {
	Payload *models.TaxCodeModel
}

func (o *GetTaxCodeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/taxcodes/{id}][%d] getTaxCodeOK  %+v", 200, o.Payload)
}
func (o *GetTaxCodeOK) GetPayload() *models.TaxCodeModel {
	return o.Payload
}

func (o *GetTaxCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaxCodeModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaxCodeBadRequest creates a GetTaxCodeBadRequest with default headers values
func NewGetTaxCodeBadRequest() *GetTaxCodeBadRequest {
	return &GetTaxCodeBadRequest{}
}

/* GetTaxCodeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetTaxCodeBadRequest struct {
}

func (o *GetTaxCodeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/taxcodes/{id}][%d] getTaxCodeBadRequest ", 400)
}

func (o *GetTaxCodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTaxCodeUnauthorized creates a GetTaxCodeUnauthorized with default headers values
func NewGetTaxCodeUnauthorized() *GetTaxCodeUnauthorized {
	return &GetTaxCodeUnauthorized{}
}

/* GetTaxCodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetTaxCodeUnauthorized struct {
}

func (o *GetTaxCodeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/taxcodes/{id}][%d] getTaxCodeUnauthorized ", 401)
}

func (o *GetTaxCodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTaxCodeNotFound creates a GetTaxCodeNotFound with default headers values
func NewGetTaxCodeNotFound() *GetTaxCodeNotFound {
	return &GetTaxCodeNotFound{}
}

/* GetTaxCodeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetTaxCodeNotFound struct {
}

func (o *GetTaxCodeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/taxcodes/{id}][%d] getTaxCodeNotFound ", 404)
}

func (o *GetTaxCodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
