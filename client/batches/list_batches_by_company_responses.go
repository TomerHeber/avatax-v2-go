// Code generated by go-swagger; DO NOT EDIT.

package batches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// ListBatchesByCompanyReader is a Reader for the ListBatchesByCompany structure.
type ListBatchesByCompanyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBatchesByCompanyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBatchesByCompanyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListBatchesByCompanyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListBatchesByCompanyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListBatchesByCompanyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListBatchesByCompanyOK creates a ListBatchesByCompanyOK with default headers values
func NewListBatchesByCompanyOK() *ListBatchesByCompanyOK {
	return &ListBatchesByCompanyOK{}
}

/* ListBatchesByCompanyOK describes a response with status code 200, with default header values.

Success
*/
type ListBatchesByCompanyOK struct {
	Payload *models.BatchModelFetchResult
}

func (o *ListBatchesByCompanyOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/batches][%d] listBatchesByCompanyOK  %+v", 200, o.Payload)
}
func (o *ListBatchesByCompanyOK) GetPayload() *models.BatchModelFetchResult {
	return o.Payload
}

func (o *ListBatchesByCompanyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BatchModelFetchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBatchesByCompanyBadRequest creates a ListBatchesByCompanyBadRequest with default headers values
func NewListBatchesByCompanyBadRequest() *ListBatchesByCompanyBadRequest {
	return &ListBatchesByCompanyBadRequest{}
}

/* ListBatchesByCompanyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListBatchesByCompanyBadRequest struct {
}

func (o *ListBatchesByCompanyBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/batches][%d] listBatchesByCompanyBadRequest ", 400)
}

func (o *ListBatchesByCompanyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListBatchesByCompanyUnauthorized creates a ListBatchesByCompanyUnauthorized with default headers values
func NewListBatchesByCompanyUnauthorized() *ListBatchesByCompanyUnauthorized {
	return &ListBatchesByCompanyUnauthorized{}
}

/* ListBatchesByCompanyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListBatchesByCompanyUnauthorized struct {
}

func (o *ListBatchesByCompanyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/batches][%d] listBatchesByCompanyUnauthorized ", 401)
}

func (o *ListBatchesByCompanyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListBatchesByCompanyNotFound creates a ListBatchesByCompanyNotFound with default headers values
func NewListBatchesByCompanyNotFound() *ListBatchesByCompanyNotFound {
	return &ListBatchesByCompanyNotFound{}
}

/* ListBatchesByCompanyNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ListBatchesByCompanyNotFound struct {
}

func (o *ListBatchesByCompanyNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/batches][%d] listBatchesByCompanyNotFound ", 404)
}

func (o *ListBatchesByCompanyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
