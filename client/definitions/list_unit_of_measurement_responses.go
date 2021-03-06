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

// ListUnitOfMeasurementReader is a Reader for the ListUnitOfMeasurement structure.
type ListUnitOfMeasurementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUnitOfMeasurementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUnitOfMeasurementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListUnitOfMeasurementOK creates a ListUnitOfMeasurementOK with default headers values
func NewListUnitOfMeasurementOK() *ListUnitOfMeasurementOK {
	return &ListUnitOfMeasurementOK{}
}

/* ListUnitOfMeasurementOK describes a response with status code 200, with default header values.

Success
*/
type ListUnitOfMeasurementOK struct {
	Payload *models.UomModelFetchResult
}

func (o *ListUnitOfMeasurementOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/definitions/unitofmeasurements][%d] listUnitOfMeasurementOK  %+v", 200, o.Payload)
}
func (o *ListUnitOfMeasurementOK) GetPayload() *models.UomModelFetchResult {
	return o.Payload
}

func (o *ListUnitOfMeasurementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UomModelFetchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
