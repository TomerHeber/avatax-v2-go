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

// DeleteTaxCodeReader is a Reader for the DeleteTaxCode structure.
type DeleteTaxCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTaxCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTaxCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteTaxCodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteTaxCodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTaxCodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteTaxCodeOK creates a DeleteTaxCodeOK with default headers values
func NewDeleteTaxCodeOK() *DeleteTaxCodeOK {
	return &DeleteTaxCodeOK{}
}

/* DeleteTaxCodeOK describes a response with status code 200, with default header values.

Success
*/
type DeleteTaxCodeOK struct {
	Payload []*models.ErrorDetail
}

func (o *DeleteTaxCodeOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/taxcodes/{id}][%d] deleteTaxCodeOK  %+v", 200, o.Payload)
}
func (o *DeleteTaxCodeOK) GetPayload() []*models.ErrorDetail {
	return o.Payload
}

func (o *DeleteTaxCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTaxCodeBadRequest creates a DeleteTaxCodeBadRequest with default headers values
func NewDeleteTaxCodeBadRequest() *DeleteTaxCodeBadRequest {
	return &DeleteTaxCodeBadRequest{}
}

/* DeleteTaxCodeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteTaxCodeBadRequest struct {
}

func (o *DeleteTaxCodeBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/taxcodes/{id}][%d] deleteTaxCodeBadRequest ", 400)
}

func (o *DeleteTaxCodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTaxCodeUnauthorized creates a DeleteTaxCodeUnauthorized with default headers values
func NewDeleteTaxCodeUnauthorized() *DeleteTaxCodeUnauthorized {
	return &DeleteTaxCodeUnauthorized{}
}

/* DeleteTaxCodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteTaxCodeUnauthorized struct {
}

func (o *DeleteTaxCodeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/taxcodes/{id}][%d] deleteTaxCodeUnauthorized ", 401)
}

func (o *DeleteTaxCodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTaxCodeNotFound creates a DeleteTaxCodeNotFound with default headers values
func NewDeleteTaxCodeNotFound() *DeleteTaxCodeNotFound {
	return &DeleteTaxCodeNotFound{}
}

/* DeleteTaxCodeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteTaxCodeNotFound struct {
}

func (o *DeleteTaxCodeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/taxcodes/{id}][%d] deleteTaxCodeNotFound ", 404)
}

func (o *DeleteTaxCodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
