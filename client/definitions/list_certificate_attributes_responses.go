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

// ListCertificateAttributesReader is a Reader for the ListCertificateAttributes structure.
type ListCertificateAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCertificateAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCertificateAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCertificateAttributesOK creates a ListCertificateAttributesOK with default headers values
func NewListCertificateAttributesOK() *ListCertificateAttributesOK {
	return &ListCertificateAttributesOK{}
}

/* ListCertificateAttributesOK describes a response with status code 200, with default header values.

Success
*/
type ListCertificateAttributesOK struct {
	Payload *models.CertificateAttributeModelFetchResult
}

func (o *ListCertificateAttributesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/definitions/certificateattributes][%d] listCertificateAttributesOK  %+v", 200, o.Payload)
}
func (o *ListCertificateAttributesOK) GetPayload() *models.CertificateAttributeModelFetchResult {
	return o.Payload
}

func (o *ListCertificateAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateAttributeModelFetchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}