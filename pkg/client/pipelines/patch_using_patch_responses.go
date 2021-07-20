// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// PatchUsingPATCHReader is a Reader for the PatchUsingPATCH structure.
type PatchUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchUsingPATCHNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchUsingPATCHInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchUsingPATCHOK creates a PatchUsingPATCHOK with default headers values
func NewPatchUsingPATCHOK() *PatchUsingPATCHOK {
	return &PatchUsingPATCHOK{}
}

/* PatchUsingPATCHOK describes a response with status code 200, with default header values.

'Success' with the updated Pipeline
*/
type PatchUsingPATCHOK struct {
	Payload models.Pipeline
}

func (o *PatchUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{id}][%d] patchUsingPATCHOK  %+v", 200, o.Payload)
}
func (o *PatchUsingPATCHOK) GetPayload() models.Pipeline {
	return o.Payload
}

func (o *PatchUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalPipeline(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewPatchUsingPATCHUnauthorized creates a PatchUsingPATCHUnauthorized with default headers values
func NewPatchUsingPATCHUnauthorized() *PatchUsingPATCHUnauthorized {
	return &PatchUsingPATCHUnauthorized{}
}

/* PatchUsingPATCHUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type PatchUsingPATCHUnauthorized struct {
}

func (o *PatchUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{id}][%d] patchUsingPATCHUnauthorized ", 401)
}

func (o *PatchUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchUsingPATCHForbidden creates a PatchUsingPATCHForbidden with default headers values
func NewPatchUsingPATCHForbidden() *PatchUsingPATCHForbidden {
	return &PatchUsingPATCHForbidden{}
}

/* PatchUsingPATCHForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchUsingPATCHForbidden struct {
}

func (o *PatchUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{id}][%d] patchUsingPATCHForbidden ", 403)
}

func (o *PatchUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchUsingPATCHNotFound creates a PatchUsingPATCHNotFound with default headers values
func NewPatchUsingPATCHNotFound() *PatchUsingPATCHNotFound {
	return &PatchUsingPATCHNotFound{}
}

/* PatchUsingPATCHNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchUsingPATCHNotFound struct {
	Payload *models.Error
}

func (o *PatchUsingPATCHNotFound) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{id}][%d] patchUsingPATCHNotFound  %+v", 404, o.Payload)
}
func (o *PatchUsingPATCHNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchUsingPATCHNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsingPATCHInternalServerError creates a PatchUsingPATCHInternalServerError with default headers values
func NewPatchUsingPATCHInternalServerError() *PatchUsingPATCHInternalServerError {
	return &PatchUsingPATCHInternalServerError{}
}

/* PatchUsingPATCHInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PatchUsingPATCHInternalServerError struct {
}

func (o *PatchUsingPATCHInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{id}][%d] patchUsingPATCHInternalServerError ", 500)
}

func (o *PatchUsingPATCHInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
