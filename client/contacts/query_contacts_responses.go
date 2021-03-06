// Code generated by go-swagger; DO NOT EDIT.

package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// QueryContactsReader is a Reader for the QueryContacts structure.
type QueryContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQueryContactsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueryContactsOK creates a QueryContactsOK with default headers values
func NewQueryContactsOK() *QueryContactsOK {
	return &QueryContactsOK{}
}

/* QueryContactsOK describes a response with status code 200, with default header values.

Success
*/
type QueryContactsOK struct {
	Payload *models.ContactModelFetchResult
}

func (o *QueryContactsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/contacts][%d] queryContactsOK  %+v", 200, o.Payload)
}
func (o *QueryContactsOK) GetPayload() *models.ContactModelFetchResult {
	return o.Payload
}

func (o *QueryContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactModelFetchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryContactsBadRequest creates a QueryContactsBadRequest with default headers values
func NewQueryContactsBadRequest() *QueryContactsBadRequest {
	return &QueryContactsBadRequest{}
}

/* QueryContactsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryContactsBadRequest struct {
}

func (o *QueryContactsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/contacts][%d] queryContactsBadRequest ", 400)
}

func (o *QueryContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueryContactsUnauthorized creates a QueryContactsUnauthorized with default headers values
func NewQueryContactsUnauthorized() *QueryContactsUnauthorized {
	return &QueryContactsUnauthorized{}
}

/* QueryContactsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type QueryContactsUnauthorized struct {
}

func (o *QueryContactsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/contacts][%d] queryContactsUnauthorized ", 401)
}

func (o *QueryContactsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
