// Code generated by go-swagger; DO NOT EDIT.

package executions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetAllUsingGET3Reader is a Reader for the GetAllUsingGET3 structure.
type GetAllUsingGET3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllUsingGET3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllUsingGET3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllUsingGET3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllUsingGET3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllUsingGET3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllUsingGET3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllUsingGET3OK creates a GetAllUsingGET3OK with default headers values
func NewGetAllUsingGET3OK() *GetAllUsingGET3OK {
	return &GetAllUsingGET3OK{}
}

/* GetAllUsingGET3OK describes a response with status code 200, with default header values.

'Success' with the requested Executions
*/
type GetAllUsingGET3OK struct {
	Payload models.Executions
}

func (o *GetAllUsingGET3OK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/executions][%d] getAllUsingGET3OK  %+v", 200, o.Payload)
}
func (o *GetAllUsingGET3OK) GetPayload() models.Executions {
	return o.Payload
}

func (o *GetAllUsingGET3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecutions(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetAllUsingGET3Unauthorized creates a GetAllUsingGET3Unauthorized with default headers values
func NewGetAllUsingGET3Unauthorized() *GetAllUsingGET3Unauthorized {
	return &GetAllUsingGET3Unauthorized{}
}

/* GetAllUsingGET3Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetAllUsingGET3Unauthorized struct {
}

func (o *GetAllUsingGET3Unauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/executions][%d] getAllUsingGET3Unauthorized ", 401)
}

func (o *GetAllUsingGET3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUsingGET3Forbidden creates a GetAllUsingGET3Forbidden with default headers values
func NewGetAllUsingGET3Forbidden() *GetAllUsingGET3Forbidden {
	return &GetAllUsingGET3Forbidden{}
}

/* GetAllUsingGET3Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllUsingGET3Forbidden struct {
}

func (o *GetAllUsingGET3Forbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/executions][%d] getAllUsingGET3Forbidden ", 403)
}

func (o *GetAllUsingGET3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUsingGET3NotFound creates a GetAllUsingGET3NotFound with default headers values
func NewGetAllUsingGET3NotFound() *GetAllUsingGET3NotFound {
	return &GetAllUsingGET3NotFound{}
}

/* GetAllUsingGET3NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllUsingGET3NotFound struct {
	Payload *models.Error
}

func (o *GetAllUsingGET3NotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/executions][%d] getAllUsingGET3NotFound  %+v", 404, o.Payload)
}
func (o *GetAllUsingGET3NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllUsingGET3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllUsingGET3InternalServerError creates a GetAllUsingGET3InternalServerError with default headers values
func NewGetAllUsingGET3InternalServerError() *GetAllUsingGET3InternalServerError {
	return &GetAllUsingGET3InternalServerError{}
}

/* GetAllUsingGET3InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetAllUsingGET3InternalServerError struct {
}

func (o *GetAllUsingGET3InternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/executions][%d] getAllUsingGET3InternalServerError ", 500)
}

func (o *GetAllUsingGET3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
