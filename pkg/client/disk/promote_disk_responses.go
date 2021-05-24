// Code generated by go-swagger; DO NOT EDIT.

package disk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// PromoteDiskReader is a Reader for the PromoteDisk structure.
type PromoteDiskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PromoteDiskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPromoteDiskAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPromoteDiskForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPromoteDiskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPromoteDiskAccepted creates a PromoteDiskAccepted with default headers values
func NewPromoteDiskAccepted() *PromoteDiskAccepted {
	return &PromoteDiskAccepted{}
}

/* PromoteDiskAccepted describes a response with status code 202, with default header values.

successful operation
*/
type PromoteDiskAccepted struct {
	Payload *models.RequestTracker
}

func (o *PromoteDiskAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/block-devices/{id}/operations/promote][%d] promoteDiskAccepted  %+v", 202, o.Payload)
}
func (o *PromoteDiskAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *PromoteDiskAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteDiskForbidden creates a PromoteDiskForbidden with default headers values
func NewPromoteDiskForbidden() *PromoteDiskForbidden {
	return &PromoteDiskForbidden{}
}

/* PromoteDiskForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PromoteDiskForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *PromoteDiskForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/block-devices/{id}/operations/promote][%d] promoteDiskForbidden  %+v", 403, o.Payload)
}
func (o *PromoteDiskForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *PromoteDiskForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteDiskNotFound creates a PromoteDiskNotFound with default headers values
func NewPromoteDiskNotFound() *PromoteDiskNotFound {
	return &PromoteDiskNotFound{}
}

/* PromoteDiskNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PromoteDiskNotFound struct {
	Payload *models.Error
}

func (o *PromoteDiskNotFound) Error() string {
	return fmt.Sprintf("[POST /iaas/api/block-devices/{id}/operations/promote][%d] promoteDiskNotFound  %+v", 404, o.Payload)
}
func (o *PromoteDiskNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PromoteDiskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
