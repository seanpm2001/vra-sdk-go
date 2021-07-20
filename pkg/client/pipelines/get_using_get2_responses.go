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

// GetUsingGET2Reader is a Reader for the GetUsingGET2 structure.
type GetUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsingGET2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsingGET2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsingGET2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsingGET2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsingGET2OK creates a GetUsingGET2OK with default headers values
func NewGetUsingGET2OK() *GetUsingGET2OK {
	return &GetUsingGET2OK{}
}

/* GetUsingGET2OK describes a response with status code 200, with default header values.

'Success' with the requested Pipeline
*/
type GetUsingGET2OK struct {
	Payload models.Pipeline
}

func (o *GetUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}][%d] getUsingGET2OK  %+v", 200, o.Payload)
}
func (o *GetUsingGET2OK) GetPayload() models.Pipeline {
	return o.Payload
}

func (o *GetUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalPipeline(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetUsingGET2Unauthorized creates a GetUsingGET2Unauthorized with default headers values
func NewGetUsingGET2Unauthorized() *GetUsingGET2Unauthorized {
	return &GetUsingGET2Unauthorized{}
}

/* GetUsingGET2Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetUsingGET2Unauthorized struct {
}

func (o *GetUsingGET2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}][%d] getUsingGET2Unauthorized ", 401)
}

func (o *GetUsingGET2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGET2Forbidden creates a GetUsingGET2Forbidden with default headers values
func NewGetUsingGET2Forbidden() *GetUsingGET2Forbidden {
	return &GetUsingGET2Forbidden{}
}

/* GetUsingGET2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUsingGET2Forbidden struct {
}

func (o *GetUsingGET2Forbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}][%d] getUsingGET2Forbidden ", 403)
}

func (o *GetUsingGET2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGET2NotFound creates a GetUsingGET2NotFound with default headers values
func NewGetUsingGET2NotFound() *GetUsingGET2NotFound {
	return &GetUsingGET2NotFound{}
}

/* GetUsingGET2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetUsingGET2NotFound struct {
	Payload *models.Error
}

func (o *GetUsingGET2NotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}][%d] getUsingGET2NotFound  %+v", 404, o.Payload)
}
func (o *GetUsingGET2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUsingGET2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsingGET2InternalServerError creates a GetUsingGET2InternalServerError with default headers values
func NewGetUsingGET2InternalServerError() *GetUsingGET2InternalServerError {
	return &GetUsingGET2InternalServerError{}
}

/* GetUsingGET2InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetUsingGET2InternalServerError struct {
}

func (o *GetUsingGET2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}][%d] getUsingGET2InternalServerError ", 500)
}

func (o *GetUsingGET2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
