// Code generated by go-swagger; DO NOT EDIT.

package custom_integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// ReleaseByIDAndVersionUsingPOSTReader is a Reader for the ReleaseByIDAndVersionUsingPOST structure.
type ReleaseByIDAndVersionUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReleaseByIDAndVersionUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReleaseByIDAndVersionUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReleaseByIDAndVersionUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReleaseByIDAndVersionUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReleaseByIDAndVersionUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReleaseByIDAndVersionUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReleaseByIDAndVersionUsingPOSTOK creates a ReleaseByIDAndVersionUsingPOSTOK with default headers values
func NewReleaseByIDAndVersionUsingPOSTOK() *ReleaseByIDAndVersionUsingPOSTOK {
	return &ReleaseByIDAndVersionUsingPOSTOK{}
}

/* ReleaseByIDAndVersionUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type ReleaseByIDAndVersionUsingPOSTOK struct {
	Payload models.CustomIntegration
}

func (o *ReleaseByIDAndVersionUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/custom-integrations/{id}/versions/{version}/release][%d] releaseByIdAndVersionUsingPOSTOK  %+v", 200, o.Payload)
}
func (o *ReleaseByIDAndVersionUsingPOSTOK) GetPayload() models.CustomIntegration {
	return o.Payload
}

func (o *ReleaseByIDAndVersionUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalCustomIntegration(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewReleaseByIDAndVersionUsingPOSTUnauthorized creates a ReleaseByIDAndVersionUsingPOSTUnauthorized with default headers values
func NewReleaseByIDAndVersionUsingPOSTUnauthorized() *ReleaseByIDAndVersionUsingPOSTUnauthorized {
	return &ReleaseByIDAndVersionUsingPOSTUnauthorized{}
}

/* ReleaseByIDAndVersionUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type ReleaseByIDAndVersionUsingPOSTUnauthorized struct {
}

func (o *ReleaseByIDAndVersionUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/custom-integrations/{id}/versions/{version}/release][%d] releaseByIdAndVersionUsingPOSTUnauthorized ", 401)
}

func (o *ReleaseByIDAndVersionUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReleaseByIDAndVersionUsingPOSTForbidden creates a ReleaseByIDAndVersionUsingPOSTForbidden with default headers values
func NewReleaseByIDAndVersionUsingPOSTForbidden() *ReleaseByIDAndVersionUsingPOSTForbidden {
	return &ReleaseByIDAndVersionUsingPOSTForbidden{}
}

/* ReleaseByIDAndVersionUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ReleaseByIDAndVersionUsingPOSTForbidden struct {
}

func (o *ReleaseByIDAndVersionUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/custom-integrations/{id}/versions/{version}/release][%d] releaseByIdAndVersionUsingPOSTForbidden ", 403)
}

func (o *ReleaseByIDAndVersionUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReleaseByIDAndVersionUsingPOSTNotFound creates a ReleaseByIDAndVersionUsingPOSTNotFound with default headers values
func NewReleaseByIDAndVersionUsingPOSTNotFound() *ReleaseByIDAndVersionUsingPOSTNotFound {
	return &ReleaseByIDAndVersionUsingPOSTNotFound{}
}

/* ReleaseByIDAndVersionUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ReleaseByIDAndVersionUsingPOSTNotFound struct {
	Payload *models.Error
}

func (o *ReleaseByIDAndVersionUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/custom-integrations/{id}/versions/{version}/release][%d] releaseByIdAndVersionUsingPOSTNotFound  %+v", 404, o.Payload)
}
func (o *ReleaseByIDAndVersionUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ReleaseByIDAndVersionUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReleaseByIDAndVersionUsingPOSTInternalServerError creates a ReleaseByIDAndVersionUsingPOSTInternalServerError with default headers values
func NewReleaseByIDAndVersionUsingPOSTInternalServerError() *ReleaseByIDAndVersionUsingPOSTInternalServerError {
	return &ReleaseByIDAndVersionUsingPOSTInternalServerError{}
}

/* ReleaseByIDAndVersionUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ReleaseByIDAndVersionUsingPOSTInternalServerError struct {
}

func (o *ReleaseByIDAndVersionUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/custom-integrations/{id}/versions/{version}/release][%d] releaseByIdAndVersionUsingPOSTInternalServerError ", 500)
}

func (o *ReleaseByIDAndVersionUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
