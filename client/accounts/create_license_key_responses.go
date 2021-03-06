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

// CreateLicenseKeyReader is a Reader for the CreateLicenseKey structure.
type CreateLicenseKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLicenseKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateLicenseKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateLicenseKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateLicenseKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateLicenseKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateLicenseKeyOK creates a CreateLicenseKeyOK with default headers values
func NewCreateLicenseKeyOK() *CreateLicenseKeyOK {
	return &CreateLicenseKeyOK{}
}

/* CreateLicenseKeyOK describes a response with status code 200, with default header values.

Success
*/
type CreateLicenseKeyOK struct {
	Payload *models.LicenseKeyModel
}

func (o *CreateLicenseKeyOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/accounts/{id}/licensekey][%d] createLicenseKeyOK  %+v", 200, o.Payload)
}
func (o *CreateLicenseKeyOK) GetPayload() *models.LicenseKeyModel {
	return o.Payload
}

func (o *CreateLicenseKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicenseKeyModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLicenseKeyBadRequest creates a CreateLicenseKeyBadRequest with default headers values
func NewCreateLicenseKeyBadRequest() *CreateLicenseKeyBadRequest {
	return &CreateLicenseKeyBadRequest{}
}

/* CreateLicenseKeyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateLicenseKeyBadRequest struct {
}

func (o *CreateLicenseKeyBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/accounts/{id}/licensekey][%d] createLicenseKeyBadRequest ", 400)
}

func (o *CreateLicenseKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateLicenseKeyUnauthorized creates a CreateLicenseKeyUnauthorized with default headers values
func NewCreateLicenseKeyUnauthorized() *CreateLicenseKeyUnauthorized {
	return &CreateLicenseKeyUnauthorized{}
}

/* CreateLicenseKeyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateLicenseKeyUnauthorized struct {
}

func (o *CreateLicenseKeyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/accounts/{id}/licensekey][%d] createLicenseKeyUnauthorized ", 401)
}

func (o *CreateLicenseKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateLicenseKeyNotFound creates a CreateLicenseKeyNotFound with default headers values
func NewCreateLicenseKeyNotFound() *CreateLicenseKeyNotFound {
	return &CreateLicenseKeyNotFound{}
}

/* CreateLicenseKeyNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateLicenseKeyNotFound struct {
}

func (o *CreateLicenseKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/accounts/{id}/licensekey][%d] createLicenseKeyNotFound ", 404)
}

func (o *CreateLicenseKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
