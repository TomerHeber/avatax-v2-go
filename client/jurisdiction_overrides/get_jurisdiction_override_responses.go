// Code generated by go-swagger; DO NOT EDIT.

package jurisdiction_overrides

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// GetJurisdictionOverrideReader is a Reader for the GetJurisdictionOverride structure.
type GetJurisdictionOverrideReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJurisdictionOverrideReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJurisdictionOverrideOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetJurisdictionOverrideBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetJurisdictionOverrideUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJurisdictionOverrideNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJurisdictionOverrideOK creates a GetJurisdictionOverrideOK with default headers values
func NewGetJurisdictionOverrideOK() *GetJurisdictionOverrideOK {
	return &GetJurisdictionOverrideOK{}
}

/* GetJurisdictionOverrideOK describes a response with status code 200, with default header values.

Success
*/
type GetJurisdictionOverrideOK struct {
	Payload *models.JurisdictionOverrideModel
}

func (o *GetJurisdictionOverrideOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/accounts/{accountId}/jurisdictionoverrides/{id}][%d] getJurisdictionOverrideOK  %+v", 200, o.Payload)
}
func (o *GetJurisdictionOverrideOK) GetPayload() *models.JurisdictionOverrideModel {
	return o.Payload
}

func (o *GetJurisdictionOverrideOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JurisdictionOverrideModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJurisdictionOverrideBadRequest creates a GetJurisdictionOverrideBadRequest with default headers values
func NewGetJurisdictionOverrideBadRequest() *GetJurisdictionOverrideBadRequest {
	return &GetJurisdictionOverrideBadRequest{}
}

/* GetJurisdictionOverrideBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetJurisdictionOverrideBadRequest struct {
}

func (o *GetJurisdictionOverrideBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/accounts/{accountId}/jurisdictionoverrides/{id}][%d] getJurisdictionOverrideBadRequest ", 400)
}

func (o *GetJurisdictionOverrideBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetJurisdictionOverrideUnauthorized creates a GetJurisdictionOverrideUnauthorized with default headers values
func NewGetJurisdictionOverrideUnauthorized() *GetJurisdictionOverrideUnauthorized {
	return &GetJurisdictionOverrideUnauthorized{}
}

/* GetJurisdictionOverrideUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetJurisdictionOverrideUnauthorized struct {
}

func (o *GetJurisdictionOverrideUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/accounts/{accountId}/jurisdictionoverrides/{id}][%d] getJurisdictionOverrideUnauthorized ", 401)
}

func (o *GetJurisdictionOverrideUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetJurisdictionOverrideNotFound creates a GetJurisdictionOverrideNotFound with default headers values
func NewGetJurisdictionOverrideNotFound() *GetJurisdictionOverrideNotFound {
	return &GetJurisdictionOverrideNotFound{}
}

/* GetJurisdictionOverrideNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetJurisdictionOverrideNotFound struct {
}

func (o *GetJurisdictionOverrideNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/accounts/{accountId}/jurisdictionoverrides/{id}][%d] getJurisdictionOverrideNotFound ", 404)
}

func (o *GetJurisdictionOverrideNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}