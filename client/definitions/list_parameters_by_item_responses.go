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

// ListParametersByItemReader is a Reader for the ListParametersByItem structure.
type ListParametersByItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListParametersByItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListParametersByItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListParametersByItemOK creates a ListParametersByItemOK with default headers values
func NewListParametersByItemOK() *ListParametersByItemOK {
	return &ListParametersByItemOK{}
}

/* ListParametersByItemOK describes a response with status code 200, with default header values.

Success
*/
type ListParametersByItemOK struct {
	Payload *models.ParameterModelFetchResult
}

func (o *ListParametersByItemOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/definitions/parameters/byitem/{companyCode}/{itemCode}][%d] listParametersByItemOK  %+v", 200, o.Payload)
}
func (o *ListParametersByItemOK) GetPayload() *models.ParameterModelFetchResult {
	return o.Payload
}

func (o *ListParametersByItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParameterModelFetchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}