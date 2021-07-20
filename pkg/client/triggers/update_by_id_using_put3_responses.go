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

// UpdateByIDUsingPUT3Reader is a Reader for the UpdateByIDUsingPUT3 structure.
type UpdateByIDUsingPUT3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateByIDUsingPUT3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateByIDUsingPUT3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateByIDUsingPUT3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateByIDUsingPUT3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateByIDUsingPUT3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateByIDUsingPUT3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateByIDUsingPUT3OK creates a UpdateByIDUsingPUT3OK with default headers values
func NewUpdateByIDUsingPUT3OK() *UpdateByIDUsingPUT3OK {
	return &UpdateByIDUsingPUT3OK{}
}

/* UpdateByIDUsingPUT3OK describes a response with status code 200, with default header values.

'Success' with Gerrit Trigger Update
*/
type UpdateByIDUsingPUT3OK struct {
	Payload models.GerritTrigger
}

func (o *UpdateByIDUsingPUT3OK) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-triggers/{id}][%d] updateByIdUsingPUT3OK  %+v", 200, o.Payload)
}
func (o *UpdateByIDUsingPUT3OK) GetPayload() models.GerritTrigger {
	return o.Payload
}

func (o *UpdateByIDUsingPUT3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritTrigger(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateByIDUsingPUT3Unauthorized creates a UpdateByIDUsingPUT3Unauthorized with default headers values
func NewUpdateByIDUsingPUT3Unauthorized() *UpdateByIDUsingPUT3Unauthorized {
	return &UpdateByIDUsingPUT3Unauthorized{}
}

/* UpdateByIDUsingPUT3Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type UpdateByIDUsingPUT3Unauthorized struct {
}

func (o *UpdateByIDUsingPUT3Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-triggers/{id}][%d] updateByIdUsingPUT3Unauthorized ", 401)
}

func (o *UpdateByIDUsingPUT3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateByIDUsingPUT3Forbidden creates a UpdateByIDUsingPUT3Forbidden with default headers values
func NewUpdateByIDUsingPUT3Forbidden() *UpdateByIDUsingPUT3Forbidden {
	return &UpdateByIDUsingPUT3Forbidden{}
}

/* UpdateByIDUsingPUT3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateByIDUsingPUT3Forbidden struct {
}

func (o *UpdateByIDUsingPUT3Forbidden) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-triggers/{id}][%d] updateByIdUsingPUT3Forbidden ", 403)
}

func (o *UpdateByIDUsingPUT3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateByIDUsingPUT3NotFound creates a UpdateByIDUsingPUT3NotFound with default headers values
func NewUpdateByIDUsingPUT3NotFound() *UpdateByIDUsingPUT3NotFound {
	return &UpdateByIDUsingPUT3NotFound{}
}

/* UpdateByIDUsingPUT3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateByIDUsingPUT3NotFound struct {
	Payload *models.Error
}

func (o *UpdateByIDUsingPUT3NotFound) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-triggers/{id}][%d] updateByIdUsingPUT3NotFound  %+v", 404, o.Payload)
}
func (o *UpdateByIDUsingPUT3NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateByIDUsingPUT3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateByIDUsingPUT3InternalServerError creates a UpdateByIDUsingPUT3InternalServerError with default headers values
func NewUpdateByIDUsingPUT3InternalServerError() *UpdateByIDUsingPUT3InternalServerError {
	return &UpdateByIDUsingPUT3InternalServerError{}
}

/* UpdateByIDUsingPUT3InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdateByIDUsingPUT3InternalServerError struct {
}

func (o *UpdateByIDUsingPUT3InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-triggers/{id}][%d] updateByIdUsingPUT3InternalServerError ", 500)
}

func (o *UpdateByIDUsingPUT3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
