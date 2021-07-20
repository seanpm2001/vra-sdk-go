// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// ResumeOrRerunEventUsingPOSTReader is a Reader for the ResumeOrRerunEventUsingPOST structure.
type ResumeOrRerunEventUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResumeOrRerunEventUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResumeOrRerunEventUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewResumeOrRerunEventUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewResumeOrRerunEventUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewResumeOrRerunEventUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResumeOrRerunEventUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewResumeOrRerunEventUsingPOSTOK creates a ResumeOrRerunEventUsingPOSTOK with default headers values
func NewResumeOrRerunEventUsingPOSTOK() *ResumeOrRerunEventUsingPOSTOK {
	return &ResumeOrRerunEventUsingPOSTOK{}
}

/* ResumeOrRerunEventUsingPOSTOK describes a response with status code 200, with default header values.

'Success' with Re-run/Resume Docker Registry Event
*/
type ResumeOrRerunEventUsingPOSTOK struct {
	Payload models.DockerRegistryEvent
}

func (o *ResumeOrRerunEventUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/registry-events/{id}][%d] resumeOrRerunEventUsingPOSTOK  %+v", 200, o.Payload)
}
func (o *ResumeOrRerunEventUsingPOSTOK) GetPayload() models.DockerRegistryEvent {
	return o.Payload
}

func (o *ResumeOrRerunEventUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalDockerRegistryEvent(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewResumeOrRerunEventUsingPOSTUnauthorized creates a ResumeOrRerunEventUsingPOSTUnauthorized with default headers values
func NewResumeOrRerunEventUsingPOSTUnauthorized() *ResumeOrRerunEventUsingPOSTUnauthorized {
	return &ResumeOrRerunEventUsingPOSTUnauthorized{}
}

/* ResumeOrRerunEventUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type ResumeOrRerunEventUsingPOSTUnauthorized struct {
}

func (o *ResumeOrRerunEventUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/registry-events/{id}][%d] resumeOrRerunEventUsingPOSTUnauthorized ", 401)
}

func (o *ResumeOrRerunEventUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResumeOrRerunEventUsingPOSTForbidden creates a ResumeOrRerunEventUsingPOSTForbidden with default headers values
func NewResumeOrRerunEventUsingPOSTForbidden() *ResumeOrRerunEventUsingPOSTForbidden {
	return &ResumeOrRerunEventUsingPOSTForbidden{}
}

/* ResumeOrRerunEventUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ResumeOrRerunEventUsingPOSTForbidden struct {
}

func (o *ResumeOrRerunEventUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/registry-events/{id}][%d] resumeOrRerunEventUsingPOSTForbidden ", 403)
}

func (o *ResumeOrRerunEventUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResumeOrRerunEventUsingPOSTNotFound creates a ResumeOrRerunEventUsingPOSTNotFound with default headers values
func NewResumeOrRerunEventUsingPOSTNotFound() *ResumeOrRerunEventUsingPOSTNotFound {
	return &ResumeOrRerunEventUsingPOSTNotFound{}
}

/* ResumeOrRerunEventUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ResumeOrRerunEventUsingPOSTNotFound struct {
	Payload *models.Error
}

func (o *ResumeOrRerunEventUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/registry-events/{id}][%d] resumeOrRerunEventUsingPOSTNotFound  %+v", 404, o.Payload)
}
func (o *ResumeOrRerunEventUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ResumeOrRerunEventUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResumeOrRerunEventUsingPOSTInternalServerError creates a ResumeOrRerunEventUsingPOSTInternalServerError with default headers values
func NewResumeOrRerunEventUsingPOSTInternalServerError() *ResumeOrRerunEventUsingPOSTInternalServerError {
	return &ResumeOrRerunEventUsingPOSTInternalServerError{}
}

/* ResumeOrRerunEventUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ResumeOrRerunEventUsingPOSTInternalServerError struct {
}

func (o *ResumeOrRerunEventUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/registry-events/{id}][%d] resumeOrRerunEventUsingPOSTInternalServerError ", 500)
}

func (o *ResumeOrRerunEventUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
