// Code generated by go-swagger; DO NOT EDIT.

package items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// UpdateItemClassificationReader is a Reader for the UpdateItemClassification structure.
type UpdateItemClassificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateItemClassificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateItemClassificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateItemClassificationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateItemClassificationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateItemClassificationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateItemClassificationOK creates a UpdateItemClassificationOK with default headers values
func NewUpdateItemClassificationOK() *UpdateItemClassificationOK {
	return &UpdateItemClassificationOK{}
}

/* UpdateItemClassificationOK describes a response with status code 200, with default header values.

Success
*/
type UpdateItemClassificationOK struct {
	Payload *models.ItemClassificationOutputModel
}

func (o *UpdateItemClassificationOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/items/{itemId}/classifications/{id}][%d] updateItemClassificationOK  %+v", 200, o.Payload)
}
func (o *UpdateItemClassificationOK) GetPayload() *models.ItemClassificationOutputModel {
	return o.Payload
}

func (o *UpdateItemClassificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ItemClassificationOutputModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateItemClassificationBadRequest creates a UpdateItemClassificationBadRequest with default headers values
func NewUpdateItemClassificationBadRequest() *UpdateItemClassificationBadRequest {
	return &UpdateItemClassificationBadRequest{}
}

/* UpdateItemClassificationBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateItemClassificationBadRequest struct {
}

func (o *UpdateItemClassificationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/items/{itemId}/classifications/{id}][%d] updateItemClassificationBadRequest ", 400)
}

func (o *UpdateItemClassificationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateItemClassificationUnauthorized creates a UpdateItemClassificationUnauthorized with default headers values
func NewUpdateItemClassificationUnauthorized() *UpdateItemClassificationUnauthorized {
	return &UpdateItemClassificationUnauthorized{}
}

/* UpdateItemClassificationUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateItemClassificationUnauthorized struct {
}

func (o *UpdateItemClassificationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/items/{itemId}/classifications/{id}][%d] updateItemClassificationUnauthorized ", 401)
}

func (o *UpdateItemClassificationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateItemClassificationNotFound creates a UpdateItemClassificationNotFound with default headers values
func NewUpdateItemClassificationNotFound() *UpdateItemClassificationNotFound {
	return &UpdateItemClassificationNotFound{}
}

/* UpdateItemClassificationNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateItemClassificationNotFound struct {
}

func (o *UpdateItemClassificationNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/items/{itemId}/classifications/{id}][%d] updateItemClassificationNotFound ", 404)
}

func (o *UpdateItemClassificationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}