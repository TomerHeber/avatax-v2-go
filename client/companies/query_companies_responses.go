// Code generated by go-swagger; DO NOT EDIT.

package companies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// QueryCompaniesReader is a Reader for the QueryCompanies structure.
type QueryCompaniesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryCompaniesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryCompaniesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryCompaniesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryCompaniesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueryCompaniesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryCompaniesOK creates a QueryCompaniesOK with default headers values
func NewQueryCompaniesOK() *QueryCompaniesOK {
	return &QueryCompaniesOK{}
}

/* QueryCompaniesOK describes a response with status code 200, with default header values.

Success
*/
type QueryCompaniesOK struct {
	Payload *models.CompanyModelFetchResult
}

func (o *QueryCompaniesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies][%d] queryCompaniesOK  %+v", 200, o.Payload)
}
func (o *QueryCompaniesOK) GetPayload() *models.CompanyModelFetchResult {
	return o.Payload
}

func (o *QueryCompaniesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CompanyModelFetchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryCompaniesBadRequest creates a QueryCompaniesBadRequest with default headers values
func NewQueryCompaniesBadRequest() *QueryCompaniesBadRequest {
	return &QueryCompaniesBadRequest{}
}

/* QueryCompaniesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryCompaniesBadRequest struct {
}

func (o *QueryCompaniesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies][%d] queryCompaniesBadRequest ", 400)
}

func (o *QueryCompaniesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueryCompaniesUnauthorized creates a QueryCompaniesUnauthorized with default headers values
func NewQueryCompaniesUnauthorized() *QueryCompaniesUnauthorized {
	return &QueryCompaniesUnauthorized{}
}

/* QueryCompaniesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type QueryCompaniesUnauthorized struct {
}

func (o *QueryCompaniesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies][%d] queryCompaniesUnauthorized ", 401)
}

func (o *QueryCompaniesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueryCompaniesNotFound creates a QueryCompaniesNotFound with default headers values
func NewQueryCompaniesNotFound() *QueryCompaniesNotFound {
	return &QueryCompaniesNotFound{}
}

/* QueryCompaniesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type QueryCompaniesNotFound struct {
}

func (o *QueryCompaniesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies][%d] queryCompaniesNotFound ", 404)
}

func (o *QueryCompaniesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}