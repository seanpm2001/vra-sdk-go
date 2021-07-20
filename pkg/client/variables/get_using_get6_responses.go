// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetUsingGET6Reader is a Reader for the GetUsingGET6 structure.
type GetUsingGET6Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsingGET6Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsingGET6OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsingGET6Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsingGET6Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsingGET6NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsingGET6InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsingGET6OK creates a GetUsingGET6OK with default headers values
func NewGetUsingGET6OK() *GetUsingGET6OK {
	return &GetUsingGET6OK{}
}

/* GetUsingGET6OK describes a response with status code 200, with default header values.

'Success' with the requested Variable
*/
type GetUsingGET6OK struct {
	Payload *models.Variable
}

func (o *GetUsingGET6OK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/variables/{project}/{name}][%d] getUsingGET6OK  %+v", 200, o.Payload)
}
func (o *GetUsingGET6OK) GetPayload() *models.Variable {
	return o.Payload
}

func (o *GetUsingGET6OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Variable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsingGET6Unauthorized creates a GetUsingGET6Unauthorized with default headers values
func NewGetUsingGET6Unauthorized() *GetUsingGET6Unauthorized {
	return &GetUsingGET6Unauthorized{}
}

/* GetUsingGET6Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetUsingGET6Unauthorized struct {
}

func (o *GetUsingGET6Unauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/variables/{project}/{name}][%d] getUsingGET6Unauthorized ", 401)
}

func (o *GetUsingGET6Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGET6Forbidden creates a GetUsingGET6Forbidden with default headers values
func NewGetUsingGET6Forbidden() *GetUsingGET6Forbidden {
	return &GetUsingGET6Forbidden{}
}

/* GetUsingGET6Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUsingGET6Forbidden struct {
}

func (o *GetUsingGET6Forbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/variables/{project}/{name}][%d] getUsingGET6Forbidden ", 403)
}

func (o *GetUsingGET6Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGET6NotFound creates a GetUsingGET6NotFound with default headers values
func NewGetUsingGET6NotFound() *GetUsingGET6NotFound {
	return &GetUsingGET6NotFound{}
}

/* GetUsingGET6NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetUsingGET6NotFound struct {
	Payload *models.Error
}

func (o *GetUsingGET6NotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/variables/{project}/{name}][%d] getUsingGET6NotFound  %+v", 404, o.Payload)
}
func (o *GetUsingGET6NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUsingGET6NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsingGET6InternalServerError creates a GetUsingGET6InternalServerError with default headers values
func NewGetUsingGET6InternalServerError() *GetUsingGET6InternalServerError {
	return &GetUsingGET6InternalServerError{}
}

/* GetUsingGET6InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetUsingGET6InternalServerError struct {
}

func (o *GetUsingGET6InternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/variables/{project}/{name}][%d] getUsingGET6InternalServerError ", 500)
}

func (o *GetUsingGET6InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
