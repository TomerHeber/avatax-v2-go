// Code generated by go-swagger; DO NOT EDIT.

package notices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// DeleteNoticeResponsibilityTypeReader is a Reader for the DeleteNoticeResponsibilityType structure.
type DeleteNoticeResponsibilityTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNoticeResponsibilityTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNoticeResponsibilityTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteNoticeResponsibilityTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteNoticeResponsibilityTypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteNoticeResponsibilityTypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNoticeResponsibilityTypeOK creates a DeleteNoticeResponsibilityTypeOK with default headers values
func NewDeleteNoticeResponsibilityTypeOK() *DeleteNoticeResponsibilityTypeOK {
	return &DeleteNoticeResponsibilityTypeOK{}
}

/* DeleteNoticeResponsibilityTypeOK describes a response with status code 200, with default header values.

Success
*/
type DeleteNoticeResponsibilityTypeOK struct {
	Payload []*models.ErrorDetail
}

func (o *DeleteNoticeResponsibilityTypeOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/notices/responsibilities/{responsibilityId}][%d] deleteNoticeResponsibilityTypeOK  %+v", 200, o.Payload)
}
func (o *DeleteNoticeResponsibilityTypeOK) GetPayload() []*models.ErrorDetail {
	return o.Payload
}

func (o *DeleteNoticeResponsibilityTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNoticeResponsibilityTypeBadRequest creates a DeleteNoticeResponsibilityTypeBadRequest with default headers values
func NewDeleteNoticeResponsibilityTypeBadRequest() *DeleteNoticeResponsibilityTypeBadRequest {
	return &DeleteNoticeResponsibilityTypeBadRequest{}
}

/* DeleteNoticeResponsibilityTypeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteNoticeResponsibilityTypeBadRequest struct {
}

func (o *DeleteNoticeResponsibilityTypeBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/notices/responsibilities/{responsibilityId}][%d] deleteNoticeResponsibilityTypeBadRequest ", 400)
}

func (o *DeleteNoticeResponsibilityTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNoticeResponsibilityTypeUnauthorized creates a DeleteNoticeResponsibilityTypeUnauthorized with default headers values
func NewDeleteNoticeResponsibilityTypeUnauthorized() *DeleteNoticeResponsibilityTypeUnauthorized {
	return &DeleteNoticeResponsibilityTypeUnauthorized{}
}

/* DeleteNoticeResponsibilityTypeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteNoticeResponsibilityTypeUnauthorized struct {
}

func (o *DeleteNoticeResponsibilityTypeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/notices/responsibilities/{responsibilityId}][%d] deleteNoticeResponsibilityTypeUnauthorized ", 401)
}

func (o *DeleteNoticeResponsibilityTypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNoticeResponsibilityTypeNotFound creates a DeleteNoticeResponsibilityTypeNotFound with default headers values
func NewDeleteNoticeResponsibilityTypeNotFound() *DeleteNoticeResponsibilityTypeNotFound {
	return &DeleteNoticeResponsibilityTypeNotFound{}
}

/* DeleteNoticeResponsibilityTypeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteNoticeResponsibilityTypeNotFound struct {
}

func (o *DeleteNoticeResponsibilityTypeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/notices/responsibilities/{responsibilityId}][%d] deleteNoticeResponsibilityTypeNotFound ", 404)
}

func (o *DeleteNoticeResponsibilityTypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}