// Code generated by go-swagger; DO NOT EDIT.

package settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// QuerySettingsReader is a Reader for the QuerySettings structure.
type QuerySettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuerySettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuerySettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQuerySettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewQuerySettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQuerySettingsOK creates a QuerySettingsOK with default headers values
func NewQuerySettingsOK() *QuerySettingsOK {
	return &QuerySettingsOK{}
}

/* QuerySettingsOK describes a response with status code 200, with default header values.

Success
*/
type QuerySettingsOK struct {
	Payload *models.SettingModelFetchResult
}

func (o *QuerySettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/settings][%d] querySettingsOK  %+v", 200, o.Payload)
}
func (o *QuerySettingsOK) GetPayload() *models.SettingModelFetchResult {
	return o.Payload
}

func (o *QuerySettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SettingModelFetchResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuerySettingsBadRequest creates a QuerySettingsBadRequest with default headers values
func NewQuerySettingsBadRequest() *QuerySettingsBadRequest {
	return &QuerySettingsBadRequest{}
}

/* QuerySettingsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QuerySettingsBadRequest struct {
}

func (o *QuerySettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/settings][%d] querySettingsBadRequest ", 400)
}

func (o *QuerySettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQuerySettingsUnauthorized creates a QuerySettingsUnauthorized with default headers values
func NewQuerySettingsUnauthorized() *QuerySettingsUnauthorized {
	return &QuerySettingsUnauthorized{}
}

/* QuerySettingsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type QuerySettingsUnauthorized struct {
}

func (o *QuerySettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/settings][%d] querySettingsUnauthorized ", 401)
}

func (o *QuerySettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}