// Code generated by go-swagger; DO NOT EDIT.

package funding_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// FundingRequestStatusReader is a Reader for the FundingRequestStatus structure.
type FundingRequestStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FundingRequestStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFundingRequestStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFundingRequestStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFundingRequestStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFundingRequestStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFundingRequestStatusOK creates a FundingRequestStatusOK with default headers values
func NewFundingRequestStatusOK() *FundingRequestStatusOK {
	return &FundingRequestStatusOK{}
}

/* FundingRequestStatusOK describes a response with status code 200, with default header values.

Success
*/
type FundingRequestStatusOK struct {
	Payload *models.FundingStatusModel
}

func (o *FundingRequestStatusOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/fundingrequests/{id}][%d] fundingRequestStatusOK  %+v", 200, o.Payload)
}
func (o *FundingRequestStatusOK) GetPayload() *models.FundingStatusModel {
	return o.Payload
}

func (o *FundingRequestStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FundingStatusModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFundingRequestStatusBadRequest creates a FundingRequestStatusBadRequest with default headers values
func NewFundingRequestStatusBadRequest() *FundingRequestStatusBadRequest {
	return &FundingRequestStatusBadRequest{}
}

/* FundingRequestStatusBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type FundingRequestStatusBadRequest struct {
}

func (o *FundingRequestStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/fundingrequests/{id}][%d] fundingRequestStatusBadRequest ", 400)
}

func (o *FundingRequestStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFundingRequestStatusUnauthorized creates a FundingRequestStatusUnauthorized with default headers values
func NewFundingRequestStatusUnauthorized() *FundingRequestStatusUnauthorized {
	return &FundingRequestStatusUnauthorized{}
}

/* FundingRequestStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type FundingRequestStatusUnauthorized struct {
}

func (o *FundingRequestStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/fundingrequests/{id}][%d] fundingRequestStatusUnauthorized ", 401)
}

func (o *FundingRequestStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFundingRequestStatusNotFound creates a FundingRequestStatusNotFound with default headers values
func NewFundingRequestStatusNotFound() *FundingRequestStatusNotFound {
	return &FundingRequestStatusNotFound{}
}

/* FundingRequestStatusNotFound describes a response with status code 404, with default header values.

Not Found
*/
type FundingRequestStatusNotFound struct {
}

func (o *FundingRequestStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/fundingrequests/{id}][%d] fundingRequestStatusNotFound ", 404)
}

func (o *FundingRequestStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}