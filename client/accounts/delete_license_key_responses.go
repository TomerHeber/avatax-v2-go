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

// DeleteLicenseKeyReader is a Reader for the DeleteLicenseKey structure.
type DeleteLicenseKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLicenseKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLicenseKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLicenseKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteLicenseKeyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLicenseKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLicenseKeyOK creates a DeleteLicenseKeyOK with default headers values
func NewDeleteLicenseKeyOK() *DeleteLicenseKeyOK {
	return &DeleteLicenseKeyOK{}
}

/* DeleteLicenseKeyOK describes a response with status code 200, with default header values.

Success
*/
type DeleteLicenseKeyOK struct {
	Payload []*models.ErrorDetail
}

func (o *DeleteLicenseKeyOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/accounts/{id}/licensekey/{licensekeyname}][%d] deleteLicenseKeyOK  %+v", 200, o.Payload)
}
func (o *DeleteLicenseKeyOK) GetPayload() []*models.ErrorDetail {
	return o.Payload
}

func (o *DeleteLicenseKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLicenseKeyBadRequest creates a DeleteLicenseKeyBadRequest with default headers values
func NewDeleteLicenseKeyBadRequest() *DeleteLicenseKeyBadRequest {
	return &DeleteLicenseKeyBadRequest{}
}

/* DeleteLicenseKeyBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteLicenseKeyBadRequest struct {
}

func (o *DeleteLicenseKeyBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/accounts/{id}/licensekey/{licensekeyname}][%d] deleteLicenseKeyBadRequest ", 400)
}

func (o *DeleteLicenseKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLicenseKeyUnauthorized creates a DeleteLicenseKeyUnauthorized with default headers values
func NewDeleteLicenseKeyUnauthorized() *DeleteLicenseKeyUnauthorized {
	return &DeleteLicenseKeyUnauthorized{}
}

/* DeleteLicenseKeyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteLicenseKeyUnauthorized struct {
}

func (o *DeleteLicenseKeyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/accounts/{id}/licensekey/{licensekeyname}][%d] deleteLicenseKeyUnauthorized ", 401)
}

func (o *DeleteLicenseKeyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLicenseKeyNotFound creates a DeleteLicenseKeyNotFound with default headers values
func NewDeleteLicenseKeyNotFound() *DeleteLicenseKeyNotFound {
	return &DeleteLicenseKeyNotFound{}
}

/* DeleteLicenseKeyNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteLicenseKeyNotFound struct {
}

func (o *DeleteLicenseKeyNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/accounts/{id}/licensekey/{licensekeyname}][%d] deleteLicenseKeyNotFound ", 404)
}

func (o *DeleteLicenseKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}