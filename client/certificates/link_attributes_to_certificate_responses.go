// Code generated by go-swagger; DO NOT EDIT.

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// LinkAttributesToCertificateReader is a Reader for the LinkAttributesToCertificate structure.
type LinkAttributesToCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LinkAttributesToCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLinkAttributesToCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLinkAttributesToCertificateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewLinkAttributesToCertificateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLinkAttributesToCertificateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLinkAttributesToCertificateOK creates a LinkAttributesToCertificateOK with default headers values
func NewLinkAttributesToCertificateOK() *LinkAttributesToCertificateOK {
	return &LinkAttributesToCertificateOK{}
}

/* LinkAttributesToCertificateOK describes a response with status code 200, with default header values.

Success
*/
type LinkAttributesToCertificateOK struct {
	Payload *models.CertificateAttributeModelFetchResult
}

func (o *LinkAttributesToCertificateOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/certificates/{id}/attributes/link][%d] linkAttributesToCertificateOK  %+v", 200, o.Payload)
}
func (o *LinkAttributesToCertificateOK) GetPayload() *models.CertificateAttributeModelFetchResult {
	return o.Payload
}

func (o *LinkAttributesToCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateAttributeModelFetchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLinkAttributesToCertificateBadRequest creates a LinkAttributesToCertificateBadRequest with default headers values
func NewLinkAttributesToCertificateBadRequest() *LinkAttributesToCertificateBadRequest {
	return &LinkAttributesToCertificateBadRequest{}
}

/* LinkAttributesToCertificateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type LinkAttributesToCertificateBadRequest struct {
}

func (o *LinkAttributesToCertificateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/certificates/{id}/attributes/link][%d] linkAttributesToCertificateBadRequest ", 400)
}

func (o *LinkAttributesToCertificateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLinkAttributesToCertificateUnauthorized creates a LinkAttributesToCertificateUnauthorized with default headers values
func NewLinkAttributesToCertificateUnauthorized() *LinkAttributesToCertificateUnauthorized {
	return &LinkAttributesToCertificateUnauthorized{}
}

/* LinkAttributesToCertificateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type LinkAttributesToCertificateUnauthorized struct {
}

func (o *LinkAttributesToCertificateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/certificates/{id}/attributes/link][%d] linkAttributesToCertificateUnauthorized ", 401)
}

func (o *LinkAttributesToCertificateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLinkAttributesToCertificateNotFound creates a LinkAttributesToCertificateNotFound with default headers values
func NewLinkAttributesToCertificateNotFound() *LinkAttributesToCertificateNotFound {
	return &LinkAttributesToCertificateNotFound{}
}

/* LinkAttributesToCertificateNotFound describes a response with status code 404, with default header values.

Not Found
*/
type LinkAttributesToCertificateNotFound struct {
}

func (o *LinkAttributesToCertificateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/companies/{companyId}/certificates/{id}/attributes/link][%d] linkAttributesToCertificateNotFound ", 404)
}

func (o *LinkAttributesToCertificateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
