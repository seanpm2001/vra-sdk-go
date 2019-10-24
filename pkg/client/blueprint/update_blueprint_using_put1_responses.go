// Code generated by go-swagger; DO NOT EDIT.

package blueprint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateBlueprintUsingPUT1Reader is a Reader for the UpdateBlueprintUsingPUT1 structure.
type UpdateBlueprintUsingPUT1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBlueprintUsingPUT1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateBlueprintUsingPUT1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateBlueprintUsingPUT1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateBlueprintUsingPUT1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateBlueprintUsingPUT1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateBlueprintUsingPUT1OK creates a UpdateBlueprintUsingPUT1OK with default headers values
func NewUpdateBlueprintUsingPUT1OK() *UpdateBlueprintUsingPUT1OK {
	return &UpdateBlueprintUsingPUT1OK{}
}

/*UpdateBlueprintUsingPUT1OK handles this case with default header values.

OK
*/
type UpdateBlueprintUsingPUT1OK struct {
	Payload *models.Blueprint
}

func (o *UpdateBlueprintUsingPUT1OK) Error() string {
	return fmt.Sprintf("[PUT /blueprint/api/blueprints/{blueprintId}][%d] updateBlueprintUsingPUT1OK  %+v", 200, o.Payload)
}

func (o *UpdateBlueprintUsingPUT1OK) GetPayload() *models.Blueprint {
	return o.Payload
}

func (o *UpdateBlueprintUsingPUT1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Blueprint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBlueprintUsingPUT1Unauthorized creates a UpdateBlueprintUsingPUT1Unauthorized with default headers values
func NewUpdateBlueprintUsingPUT1Unauthorized() *UpdateBlueprintUsingPUT1Unauthorized {
	return &UpdateBlueprintUsingPUT1Unauthorized{}
}

/*UpdateBlueprintUsingPUT1Unauthorized handles this case with default header values.

Unauthorized
*/
type UpdateBlueprintUsingPUT1Unauthorized struct {
}

func (o *UpdateBlueprintUsingPUT1Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /blueprint/api/blueprints/{blueprintId}][%d] updateBlueprintUsingPUT1Unauthorized ", 401)
}

func (o *UpdateBlueprintUsingPUT1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBlueprintUsingPUT1Forbidden creates a UpdateBlueprintUsingPUT1Forbidden with default headers values
func NewUpdateBlueprintUsingPUT1Forbidden() *UpdateBlueprintUsingPUT1Forbidden {
	return &UpdateBlueprintUsingPUT1Forbidden{}
}

/*UpdateBlueprintUsingPUT1Forbidden handles this case with default header values.

Forbidden
*/
type UpdateBlueprintUsingPUT1Forbidden struct {
}

func (o *UpdateBlueprintUsingPUT1Forbidden) Error() string {
	return fmt.Sprintf("[PUT /blueprint/api/blueprints/{blueprintId}][%d] updateBlueprintUsingPUT1Forbidden ", 403)
}

func (o *UpdateBlueprintUsingPUT1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBlueprintUsingPUT1NotFound creates a UpdateBlueprintUsingPUT1NotFound with default headers values
func NewUpdateBlueprintUsingPUT1NotFound() *UpdateBlueprintUsingPUT1NotFound {
	return &UpdateBlueprintUsingPUT1NotFound{}
}

/*UpdateBlueprintUsingPUT1NotFound handles this case with default header values.

Not Found
*/
type UpdateBlueprintUsingPUT1NotFound struct {
}

func (o *UpdateBlueprintUsingPUT1NotFound) Error() string {
	return fmt.Sprintf("[PUT /blueprint/api/blueprints/{blueprintId}][%d] updateBlueprintUsingPUT1NotFound ", 404)
}

func (o *UpdateBlueprintUsingPUT1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
