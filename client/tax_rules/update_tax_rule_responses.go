// Code generated by go-swagger; DO NOT EDIT.

package tax_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// UpdateTaxRuleReader is a Reader for the UpdateTaxRule structure.
type UpdateTaxRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTaxRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTaxRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTaxRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateTaxRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTaxRuleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateTaxRuleOK creates a UpdateTaxRuleOK with default headers values
func NewUpdateTaxRuleOK() *UpdateTaxRuleOK {
	return &UpdateTaxRuleOK{}
}

/* UpdateTaxRuleOK describes a response with status code 200, with default header values.

Success
*/
type UpdateTaxRuleOK struct {
	Payload *models.TaxRuleModel
}

func (o *UpdateTaxRuleOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/taxrules/{id}][%d] updateTaxRuleOK  %+v", 200, o.Payload)
}
func (o *UpdateTaxRuleOK) GetPayload() *models.TaxRuleModel {
	return o.Payload
}

func (o *UpdateTaxRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaxRuleModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTaxRuleBadRequest creates a UpdateTaxRuleBadRequest with default headers values
func NewUpdateTaxRuleBadRequest() *UpdateTaxRuleBadRequest {
	return &UpdateTaxRuleBadRequest{}
}

/* UpdateTaxRuleBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateTaxRuleBadRequest struct {
}

func (o *UpdateTaxRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/taxrules/{id}][%d] updateTaxRuleBadRequest ", 400)
}

func (o *UpdateTaxRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTaxRuleUnauthorized creates a UpdateTaxRuleUnauthorized with default headers values
func NewUpdateTaxRuleUnauthorized() *UpdateTaxRuleUnauthorized {
	return &UpdateTaxRuleUnauthorized{}
}

/* UpdateTaxRuleUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateTaxRuleUnauthorized struct {
}

func (o *UpdateTaxRuleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/taxrules/{id}][%d] updateTaxRuleUnauthorized ", 401)
}

func (o *UpdateTaxRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateTaxRuleNotFound creates a UpdateTaxRuleNotFound with default headers values
func NewUpdateTaxRuleNotFound() *UpdateTaxRuleNotFound {
	return &UpdateTaxRuleNotFound{}
}

/* UpdateTaxRuleNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateTaxRuleNotFound struct {
}

func (o *UpdateTaxRuleNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/companies/{companyId}/taxrules/{id}][%d] updateTaxRuleNotFound ", 404)
}

func (o *UpdateTaxRuleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
