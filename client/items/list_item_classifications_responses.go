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

// ListItemClassificationsReader is a Reader for the ListItemClassifications structure.
type ListItemClassificationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListItemClassificationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListItemClassificationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListItemClassificationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListItemClassificationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListItemClassificationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListItemClassificationsOK creates a ListItemClassificationsOK with default headers values
func NewListItemClassificationsOK() *ListItemClassificationsOK {
	return &ListItemClassificationsOK{}
}

/* ListItemClassificationsOK describes a response with status code 200, with default header values.

Success
*/
type ListItemClassificationsOK struct {
	Payload *models.ItemClassificationOutputModelFetchResult
}

func (o *ListItemClassificationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/items/{itemId}/classifications][%d] listItemClassificationsOK  %+v", 200, o.Payload)
}
func (o *ListItemClassificationsOK) GetPayload() *models.ItemClassificationOutputModelFetchResult {
	return o.Payload
}

func (o *ListItemClassificationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ItemClassificationOutputModelFetchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListItemClassificationsBadRequest creates a ListItemClassificationsBadRequest with default headers values
func NewListItemClassificationsBadRequest() *ListItemClassificationsBadRequest {
	return &ListItemClassificationsBadRequest{}
}

/* ListItemClassificationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListItemClassificationsBadRequest struct {
}

func (o *ListItemClassificationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/items/{itemId}/classifications][%d] listItemClassificationsBadRequest ", 400)
}

func (o *ListItemClassificationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListItemClassificationsUnauthorized creates a ListItemClassificationsUnauthorized with default headers values
func NewListItemClassificationsUnauthorized() *ListItemClassificationsUnauthorized {
	return &ListItemClassificationsUnauthorized{}
}

/* ListItemClassificationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListItemClassificationsUnauthorized struct {
}

func (o *ListItemClassificationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/items/{itemId}/classifications][%d] listItemClassificationsUnauthorized ", 401)
}

func (o *ListItemClassificationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListItemClassificationsNotFound creates a ListItemClassificationsNotFound with default headers values
func NewListItemClassificationsNotFound() *ListItemClassificationsNotFound {
	return &ListItemClassificationsNotFound{}
}

/* ListItemClassificationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ListItemClassificationsNotFound struct {
}

func (o *ListItemClassificationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/companies/{companyId}/items/{itemId}/classifications][%d] listItemClassificationsNotFound ", 404)
}

func (o *ListItemClassificationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
