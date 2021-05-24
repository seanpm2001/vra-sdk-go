// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// PowerOffMachineReader is a Reader for the PowerOffMachine structure.
type PowerOffMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PowerOffMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPowerOffMachineAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPowerOffMachineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPowerOffMachineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPowerOffMachineAccepted creates a PowerOffMachineAccepted with default headers values
func NewPowerOffMachineAccepted() *PowerOffMachineAccepted {
	return &PowerOffMachineAccepted{}
}

/* PowerOffMachineAccepted describes a response with status code 202, with default header values.

successful operation
*/
type PowerOffMachineAccepted struct {
	Payload *models.RequestTracker
}

func (o *PowerOffMachineAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/power-off][%d] powerOffMachineAccepted  %+v", 202, o.Payload)
}
func (o *PowerOffMachineAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *PowerOffMachineAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPowerOffMachineForbidden creates a PowerOffMachineForbidden with default headers values
func NewPowerOffMachineForbidden() *PowerOffMachineForbidden {
	return &PowerOffMachineForbidden{}
}

/* PowerOffMachineForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PowerOffMachineForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *PowerOffMachineForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/power-off][%d] powerOffMachineForbidden  %+v", 403, o.Payload)
}
func (o *PowerOffMachineForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *PowerOffMachineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPowerOffMachineNotFound creates a PowerOffMachineNotFound with default headers values
func NewPowerOffMachineNotFound() *PowerOffMachineNotFound {
	return &PowerOffMachineNotFound{}
}

/* PowerOffMachineNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PowerOffMachineNotFound struct {
	Payload *models.Error
}

func (o *PowerOffMachineNotFound) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/power-off][%d] powerOffMachineNotFound  %+v", 404, o.Payload)
}
func (o *PowerOffMachineNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PowerOffMachineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
