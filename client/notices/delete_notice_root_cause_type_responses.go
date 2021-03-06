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

// DeleteNoticeRootCauseTypeReader is a Reader for the DeleteNoticeRootCauseType structure.
type DeleteNoticeRootCauseTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNoticeRootCauseTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNoticeRootCauseTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteNoticeRootCauseTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteNoticeRootCauseTypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteNoticeRootCauseTypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNoticeRootCauseTypeOK creates a DeleteNoticeRootCauseTypeOK with default headers values
func NewDeleteNoticeRootCauseTypeOK() *DeleteNoticeRootCauseTypeOK {
	return &DeleteNoticeRootCauseTypeOK{}
}

/* DeleteNoticeRootCauseTypeOK describes a response with status code 200, with default header values.

Success
*/
type DeleteNoticeRootCauseTypeOK struct {
	Payload []*models.ErrorDetail
}

func (o *DeleteNoticeRootCauseTypeOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/notices/rootcauses/{rootCauseId}][%d] deleteNoticeRootCauseTypeOK  %+v", 200, o.Payload)
}
func (o *DeleteNoticeRootCauseTypeOK) GetPayload() []*models.ErrorDetail {
	return o.Payload
}

func (o *DeleteNoticeRootCauseTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNoticeRootCauseTypeBadRequest creates a DeleteNoticeRootCauseTypeBadRequest with default headers values
func NewDeleteNoticeRootCauseTypeBadRequest() *DeleteNoticeRootCauseTypeBadRequest {
	return &DeleteNoticeRootCauseTypeBadRequest{}
}

/* DeleteNoticeRootCauseTypeBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteNoticeRootCauseTypeBadRequest struct {
}

func (o *DeleteNoticeRootCauseTypeBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/notices/rootcauses/{rootCauseId}][%d] deleteNoticeRootCauseTypeBadRequest ", 400)
}

func (o *DeleteNoticeRootCauseTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNoticeRootCauseTypeUnauthorized creates a DeleteNoticeRootCauseTypeUnauthorized with default headers values
func NewDeleteNoticeRootCauseTypeUnauthorized() *DeleteNoticeRootCauseTypeUnauthorized {
	return &DeleteNoticeRootCauseTypeUnauthorized{}
}

/* DeleteNoticeRootCauseTypeUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteNoticeRootCauseTypeUnauthorized struct {
}

func (o *DeleteNoticeRootCauseTypeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/notices/rootcauses/{rootCauseId}][%d] deleteNoticeRootCauseTypeUnauthorized ", 401)
}

func (o *DeleteNoticeRootCauseTypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNoticeRootCauseTypeNotFound creates a DeleteNoticeRootCauseTypeNotFound with default headers values
func NewDeleteNoticeRootCauseTypeNotFound() *DeleteNoticeRootCauseTypeNotFound {
	return &DeleteNoticeRootCauseTypeNotFound{}
}

/* DeleteNoticeRootCauseTypeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteNoticeRootCauseTypeNotFound struct {
}

func (o *DeleteNoticeRootCauseTypeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/notices/rootcauses/{rootCauseId}][%d] deleteNoticeRootCauseTypeNotFound ", 404)
}

func (o *DeleteNoticeRootCauseTypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
