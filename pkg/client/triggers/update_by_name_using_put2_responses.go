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

// UpdateByNameUsingPUT2Reader is a Reader for the UpdateByNameUsingPUT2 structure.
type UpdateByNameUsingPUT2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateByNameUsingPUT2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateByNameUsingPUT2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateByNameUsingPUT2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateByNameUsingPUT2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateByNameUsingPUT2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateByNameUsingPUT2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateByNameUsingPUT2OK creates a UpdateByNameUsingPUT2OK with default headers values
func NewUpdateByNameUsingPUT2OK() *UpdateByNameUsingPUT2OK {
	return &UpdateByNameUsingPUT2OK{}
}

/* UpdateByNameUsingPUT2OK describes a response with status code 200, with default header values.

'Success' with Gerrit Listener Update
*/
type UpdateByNameUsingPUT2OK struct {
	Payload models.GerritListener
}

func (o *UpdateByNameUsingPUT2OK) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-listeners/{project}/{name}][%d] updateByNameUsingPUT2OK  %+v", 200, o.Payload)
}
func (o *UpdateByNameUsingPUT2OK) GetPayload() models.GerritListener {
	return o.Payload
}

func (o *UpdateByNameUsingPUT2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritListener(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateByNameUsingPUT2Unauthorized creates a UpdateByNameUsingPUT2Unauthorized with default headers values
func NewUpdateByNameUsingPUT2Unauthorized() *UpdateByNameUsingPUT2Unauthorized {
	return &UpdateByNameUsingPUT2Unauthorized{}
}

/* UpdateByNameUsingPUT2Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type UpdateByNameUsingPUT2Unauthorized struct {
}

func (o *UpdateByNameUsingPUT2Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-listeners/{project}/{name}][%d] updateByNameUsingPUT2Unauthorized ", 401)
}

func (o *UpdateByNameUsingPUT2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateByNameUsingPUT2Forbidden creates a UpdateByNameUsingPUT2Forbidden with default headers values
func NewUpdateByNameUsingPUT2Forbidden() *UpdateByNameUsingPUT2Forbidden {
	return &UpdateByNameUsingPUT2Forbidden{}
}

/* UpdateByNameUsingPUT2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateByNameUsingPUT2Forbidden struct {
}

func (o *UpdateByNameUsingPUT2Forbidden) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-listeners/{project}/{name}][%d] updateByNameUsingPUT2Forbidden ", 403)
}

func (o *UpdateByNameUsingPUT2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateByNameUsingPUT2NotFound creates a UpdateByNameUsingPUT2NotFound with default headers values
func NewUpdateByNameUsingPUT2NotFound() *UpdateByNameUsingPUT2NotFound {
	return &UpdateByNameUsingPUT2NotFound{}
}

/* UpdateByNameUsingPUT2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateByNameUsingPUT2NotFound struct {
	Payload *models.Error
}

func (o *UpdateByNameUsingPUT2NotFound) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-listeners/{project}/{name}][%d] updateByNameUsingPUT2NotFound  %+v", 404, o.Payload)
}
func (o *UpdateByNameUsingPUT2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateByNameUsingPUT2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateByNameUsingPUT2InternalServerError creates a UpdateByNameUsingPUT2InternalServerError with default headers values
func NewUpdateByNameUsingPUT2InternalServerError() *UpdateByNameUsingPUT2InternalServerError {
	return &UpdateByNameUsingPUT2InternalServerError{}
}

/* UpdateByNameUsingPUT2InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdateByNameUsingPUT2InternalServerError struct {
}

func (o *UpdateByNameUsingPUT2InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/gerrit-listeners/{project}/{name}][%d] updateByNameUsingPUT2InternalServerError ", 500)
}

func (o *UpdateByNameUsingPUT2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
