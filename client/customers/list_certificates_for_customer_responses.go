// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// ListCertificatesForCustomerReader is a Reader for the ListCertificatesForCustomer structure.
type ListCertificatesForCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCertificatesForCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCertificatesForCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCertificatesForCustomerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListCertificatesForCustomerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListCertificatesForCustomerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCertificatesForCustomerOK creates a ListCertificatesForCustomerOK with default headers values
func NewListCertificatesForCustomerOK() *ListCertificatesForCustomerOK {
	return &ListCertificatesForCustomerOK{}
}

/* ListCertificatesForCustomerOK describes a response with status code 200, with default header values.

Success
*/
type ListCertificatesForCustomerOK struct {
	Payload *models.CertificateModelFetchResult
}

func (o *ListCertificatesForCustomerOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/customers/{customerCode}/certificates][%d] listCertificatesForCustomerOK  %+v", 200, o.Payload)
}
func (o *ListCertificatesForCustomerOK) GetPayload() *models.CertificateModelFetchResult {
	return o.Payload
}

func (o *ListCertificatesForCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateModelFetchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCertificatesForCustomerBadRequest creates a ListCertificatesForCustomerBadRequest with default headers values
func NewListCertificatesForCustomerBadRequest() *ListCertificatesForCustomerBadRequest {
	return &ListCertificatesForCustomerBadRequest{}
}

/* ListCertificatesForCustomerBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListCertificatesForCustomerBadRequest struct {
}

func (o *ListCertificatesForCustomerBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/customers/{customerCode}/certificates][%d] listCertificatesForCustomerBadRequest ", 400)
}

func (o *ListCertificatesForCustomerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCertificatesForCustomerUnauthorized creates a ListCertificatesForCustomerUnauthorized with default headers values
func NewListCertificatesForCustomerUnauthorized() *ListCertificatesForCustomerUnauthorized {
	return &ListCertificatesForCustomerUnauthorized{}
}

/* ListCertificatesForCustomerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListCertificatesForCustomerUnauthorized struct {
}

func (o *ListCertificatesForCustomerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/customers/{customerCode}/certificates][%d] listCertificatesForCustomerUnauthorized ", 401)
}

func (o *ListCertificatesForCustomerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListCertificatesForCustomerNotFound creates a ListCertificatesForCustomerNotFound with default headers values
func NewListCertificatesForCustomerNotFound() *ListCertificatesForCustomerNotFound {
	return &ListCertificatesForCustomerNotFound{}
}

/* ListCertificatesForCustomerNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ListCertificatesForCustomerNotFound struct {
}

func (o *ListCertificatesForCustomerNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/customers/{customerCode}/certificates][%d] listCertificatesForCustomerNotFound ", 404)
}

func (o *ListCertificatesForCustomerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}