// Code generated by go-swagger; DO NOT EDIT.

package filing_calendars

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/TomerHeber/avatax-v2-go/models"
)

// DeleteCompanyReturnSettingsReader is a Reader for the DeleteCompanyReturnSettings structure.
type DeleteCompanyReturnSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCompanyReturnSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCompanyReturnSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCompanyReturnSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCompanyReturnSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCompanyReturnSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCompanyReturnSettingsOK creates a DeleteCompanyReturnSettingsOK with default headers values
func NewDeleteCompanyReturnSettingsOK() *DeleteCompanyReturnSettingsOK {
	return &DeleteCompanyReturnSettingsOK{}
}

/* DeleteCompanyReturnSettingsOK describes a response with status code 200, with default header values.

Success
*/
type DeleteCompanyReturnSettingsOK struct {
	Payload []*models.CompanyReturnSettingModel
}

func (o *DeleteCompanyReturnSettingsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/filingcalendars/{filingCalendarId}/setting/{companyReturnSettingId}][%d] deleteCompanyReturnSettingsOK  %+v", 200, o.Payload)
}
func (o *DeleteCompanyReturnSettingsOK) GetPayload() []*models.CompanyReturnSettingModel {
	return o.Payload
}

func (o *DeleteCompanyReturnSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCompanyReturnSettingsBadRequest creates a DeleteCompanyReturnSettingsBadRequest with default headers values
func NewDeleteCompanyReturnSettingsBadRequest() *DeleteCompanyReturnSettingsBadRequest {
	return &DeleteCompanyReturnSettingsBadRequest{}
}

/* DeleteCompanyReturnSettingsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DeleteCompanyReturnSettingsBadRequest struct {
}

func (o *DeleteCompanyReturnSettingsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/filingcalendars/{filingCalendarId}/setting/{companyReturnSettingId}][%d] deleteCompanyReturnSettingsBadRequest ", 400)
}

func (o *DeleteCompanyReturnSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCompanyReturnSettingsUnauthorized creates a DeleteCompanyReturnSettingsUnauthorized with default headers values
func NewDeleteCompanyReturnSettingsUnauthorized() *DeleteCompanyReturnSettingsUnauthorized {
	return &DeleteCompanyReturnSettingsUnauthorized{}
}

/* DeleteCompanyReturnSettingsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteCompanyReturnSettingsUnauthorized struct {
}

func (o *DeleteCompanyReturnSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/filingcalendars/{filingCalendarId}/setting/{companyReturnSettingId}][%d] deleteCompanyReturnSettingsUnauthorized ", 401)
}

func (o *DeleteCompanyReturnSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCompanyReturnSettingsNotFound creates a DeleteCompanyReturnSettingsNotFound with default headers values
func NewDeleteCompanyReturnSettingsNotFound() *DeleteCompanyReturnSettingsNotFound {
	return &DeleteCompanyReturnSettingsNotFound{}
}

/* DeleteCompanyReturnSettingsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteCompanyReturnSettingsNotFound struct {
}

func (o *DeleteCompanyReturnSettingsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/companies/{companyId}/filingcalendars/{filingCalendarId}/setting/{companyReturnSettingId}][%d] deleteCompanyReturnSettingsNotFound ", 404)
}

func (o *DeleteCompanyReturnSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}