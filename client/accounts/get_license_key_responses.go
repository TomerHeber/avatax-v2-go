// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// GetLicenseKeyReader is a Reader for the GetLicenseKey structure.
type GetLicenseKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLicenseKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLicenseKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLicenseKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLicenseKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLicenseKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLicenseKeyOK creates a GetLicenseKeyOK with default headers values
func NewGetLicenseKeyOK() *GetLicenseKeyOK {
	return &GetLicenseKeyOK{}
}

/* GetLicenseKeyOK describes a response with status code 200, with default header values.

Success
*/
type GetLicenseKeyOK struct {
	Payload *models.AccountLicenseKeyModel
}

func (o *GetLicenseKeyOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/accounts/{id}/licensekey/{licensekeyname}][%d] getLicenseKeyOK  %+v", 200, o.Payload)
}
func (o *GetLicenseKeyOK) GetPayload() *models.AccountLicenseKeyModel {
	return o.Payload
}

func (o *GetLicenseKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountLicenseKeyModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseKeyBadRequest creates a GetLicenseKeyBadRequest with default headers values
func NewGetLicenseKeyBadRequest() *GetLicenseKeyBadRequest {
	return &GetLicenseKeyBadRequest{}
}

/* GetLicenseKeyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetLicenseKeyBadRequest struct {
}

func (o *GetLicenseKeyBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/accounts/{id}/licensekey/{licensekeyname}][%d] getLicenseKeyBadRequest ", 400)
}

func (o *GetLicenseKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLicenseKeyUnauthorized creates a GetLicenseKeyUnauthorized with default headers values
func NewGetLicenseKeyUnauthorized() *GetLicenseKeyUnauthorized {
	return &GetLicenseKeyUnauthorized{}
}

/* GetLicenseKeyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetLicenseKeyUnauthorized struct {
}

func (o *GetLicenseKeyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/accounts/{id}/licensekey/{licensekeyname}][%d] getLicenseKeyUnauthorized ", 401)
}

func (o *GetLicenseKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLicenseKeyNotFound creates a GetLicenseKeyNotFound with default headers values
func NewGetLicenseKeyNotFound() *GetLicenseKeyNotFound {
	return &GetLicenseKeyNotFound{}
}

/* GetLicenseKeyNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetLicenseKeyNotFound struct {
}

func (o *GetLicenseKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/accounts/{id}/licensekey/{licensekeyname}][%d] getLicenseKeyNotFound ", 404)
}

func (o *GetLicenseKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}