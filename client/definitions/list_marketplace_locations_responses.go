// Code generated by go-swagger; DO NOT EDIT.

package definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// ListMarketplaceLocationsReader is a Reader for the ListMarketplaceLocations structure.
type ListMarketplaceLocationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMarketplaceLocationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListMarketplaceLocationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListMarketplaceLocationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListMarketplaceLocationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListMarketplaceLocationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListMarketplaceLocationsOK creates a ListMarketplaceLocationsOK with default headers values
func NewListMarketplaceLocationsOK() *ListMarketplaceLocationsOK {
	return &ListMarketplaceLocationsOK{}
}

/* ListMarketplaceLocationsOK describes a response with status code 200, with default header values.

Success
*/
type ListMarketplaceLocationsOK struct {
	Payload *models.MarketplaceLocationModelFetchResult
}

func (o *ListMarketplaceLocationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/definitions/marketplacelocations][%d] listMarketplaceLocationsOK  %+v", 200, o.Payload)
}
func (o *ListMarketplaceLocationsOK) GetPayload() *models.MarketplaceLocationModelFetchResult {
	return o.Payload
}

func (o *ListMarketplaceLocationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MarketplaceLocationModelFetchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMarketplaceLocationsBadRequest creates a ListMarketplaceLocationsBadRequest with default headers values
func NewListMarketplaceLocationsBadRequest() *ListMarketplaceLocationsBadRequest {
	return &ListMarketplaceLocationsBadRequest{}
}

/* ListMarketplaceLocationsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ListMarketplaceLocationsBadRequest struct {
}

func (o *ListMarketplaceLocationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/definitions/marketplacelocations][%d] listMarketplaceLocationsBadRequest ", 400)
}

func (o *ListMarketplaceLocationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMarketplaceLocationsUnauthorized creates a ListMarketplaceLocationsUnauthorized with default headers values
func NewListMarketplaceLocationsUnauthorized() *ListMarketplaceLocationsUnauthorized {
	return &ListMarketplaceLocationsUnauthorized{}
}

/* ListMarketplaceLocationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListMarketplaceLocationsUnauthorized struct {
}

func (o *ListMarketplaceLocationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/definitions/marketplacelocations][%d] listMarketplaceLocationsUnauthorized ", 401)
}

func (o *ListMarketplaceLocationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListMarketplaceLocationsNotFound creates a ListMarketplaceLocationsNotFound with default headers values
func NewListMarketplaceLocationsNotFound() *ListMarketplaceLocationsNotFound {
	return &ListMarketplaceLocationsNotFound{}
}

/* ListMarketplaceLocationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ListMarketplaceLocationsNotFound struct {
}

func (o *ListMarketplaceLocationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/definitions/marketplacelocations][%d] listMarketplaceLocationsNotFound ", 404)
}

func (o *ListMarketplaceLocationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}