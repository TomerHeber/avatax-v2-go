// Code generated by go-swagger; DO NOT EDIT.

package upcs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// DeleteUPCReader is a Reader for the DeleteUPC structure.
type DeleteUPCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUPCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUPCOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUPCBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteUPCUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUPCNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUPCOK creates a DeleteUPCOK with default headers values
func NewDeleteUPCOK() *DeleteUPCOK {
	return &DeleteUPCOK{}
}

/* DeleteUPCOK describes a response with status code 200, with default header values.

Success
*/
type DeleteUPCOK struct {
	Payload []*models.ErrorDetail
}

func (o *DeleteUPCOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/upcs/{id}][%d] deleteUPCOK  %+v", 200, o.Payload)
}
func (o *DeleteUPCOK) GetPayload() []*models.ErrorDetail {
	return o.Payload
}

func (o *DeleteUPCOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUPCBadRequest creates a DeleteUPCBadRequest with default headers values
func NewDeleteUPCBadRequest() *DeleteUPCBadRequest {
	return &DeleteUPCBadRequest{}
}

/* DeleteUPCBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteUPCBadRequest struct {
}

func (o *DeleteUPCBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/upcs/{id}][%d] deleteUPCBadRequest ", 400)
}

func (o *DeleteUPCBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUPCUnauthorized creates a DeleteUPCUnauthorized with default headers values
func NewDeleteUPCUnauthorized() *DeleteUPCUnauthorized {
	return &DeleteUPCUnauthorized{}
}

/* DeleteUPCUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteUPCUnauthorized struct {
}

func (o *DeleteUPCUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/upcs/{id}][%d] deleteUPCUnauthorized ", 401)
}

func (o *DeleteUPCUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUPCNotFound creates a DeleteUPCNotFound with default headers values
func NewDeleteUPCNotFound() *DeleteUPCNotFound {
	return &DeleteUPCNotFound{}
}

/* DeleteUPCNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteUPCNotFound struct {
}

func (o *DeleteUPCNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/upcs/{id}][%d] deleteUPCNotFound ", 404)
}

func (o *DeleteUPCNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}