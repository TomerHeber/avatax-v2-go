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

// ListLoginVerifiersReader is a Reader for the ListLoginVerifiers structure.
type ListLoginVerifiersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLoginVerifiersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListLoginVerifiersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListLoginVerifiersOK creates a ListLoginVerifiersOK with default headers values
func NewListLoginVerifiersOK() *ListLoginVerifiersOK {
	return &ListLoginVerifiersOK{}
}

/* ListLoginVerifiersOK describes a response with status code 200, with default header values.

Success
*/
type ListLoginVerifiersOK struct {
	Payload *models.SkyscraperStatusModelFetchResult
}

func (o *ListLoginVerifiersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/definitions/filingcalendars/loginverifiers][%d] listLoginVerifiersOK  %+v", 200, o.Payload)
}
func (o *ListLoginVerifiersOK) GetPayload() *models.SkyscraperStatusModelFetchResult {
	return o.Payload
}

func (o *ListLoginVerifiersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SkyscraperStatusModelFetchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
