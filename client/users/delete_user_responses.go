// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// DeleteUserReader is a Reader for the DeleteUser structure.
type DeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserOK creates a DeleteUserOK with default headers values
func NewDeleteUserOK() *DeleteUserOK {
	return &DeleteUserOK{}
}

/* DeleteUserOK describes a response with status code 200, with default header values.

Success
*/
type DeleteUserOK struct {
	Payload []*models.ErrorDetail
}

func (o *DeleteUserOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/accounts/{accountId}/users/{id}][%d] deleteUserOK  %+v", 200, o.Payload)
}
func (o *DeleteUserOK) GetPayload() []*models.ErrorDetail {
	return o.Payload
}

func (o *DeleteUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserBadRequest creates a DeleteUserBadRequest with default headers values
func NewDeleteUserBadRequest() *DeleteUserBadRequest {
	return &DeleteUserBadRequest{}
}

/* DeleteUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteUserBadRequest struct {
}

func (o *DeleteUserBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/accounts/{accountId}/users/{id}][%d] deleteUserBadRequest ", 400)
}

func (o *DeleteUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserUnauthorized creates a DeleteUserUnauthorized with default headers values
func NewDeleteUserUnauthorized() *DeleteUserUnauthorized {
	return &DeleteUserUnauthorized{}
}

/* DeleteUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteUserUnauthorized struct {
}

func (o *DeleteUserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/accounts/{accountId}/users/{id}][%d] deleteUserUnauthorized ", 401)
}

func (o *DeleteUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserNotFound creates a DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() *DeleteUserNotFound {
	return &DeleteUserNotFound{}
}

/* DeleteUserNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteUserNotFound struct {
}

func (o *DeleteUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/accounts/{accountId}/users/{id}][%d] deleteUserNotFound ", 404)
}

func (o *DeleteUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
