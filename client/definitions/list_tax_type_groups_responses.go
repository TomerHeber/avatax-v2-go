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

// ListTaxTypeGroupsReader is a Reader for the ListTaxTypeGroups structure.
type ListTaxTypeGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTaxTypeGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTaxTypeGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTaxTypeGroupsOK creates a ListTaxTypeGroupsOK with default headers values
func NewListTaxTypeGroupsOK() *ListTaxTypeGroupsOK {
	return &ListTaxTypeGroupsOK{}
}

/* ListTaxTypeGroupsOK describes a response with status code 200, with default header values.

Success
*/
type ListTaxTypeGroupsOK struct {
	Payload *models.TaxTypeGroupModelFetchResult
}

func (o *ListTaxTypeGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/definitions/taxtypegroups][%d] listTaxTypeGroupsOK  %+v", 200, o.Payload)
}
func (o *ListTaxTypeGroupsOK) GetPayload() *models.TaxTypeGroupModelFetchResult {
	return o.Payload
}

func (o *ListTaxTypeGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaxTypeGroupModelFetchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
