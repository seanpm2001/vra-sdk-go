// Code generated by go-swagger; DO NOT EDIT.

package location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateZoneReader is a Reader for the UpdateZone structure.
type UpdateZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateZoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateZoneBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateZoneForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateZoneOK creates a UpdateZoneOK with default headers values
func NewUpdateZoneOK() *UpdateZoneOK {
	return &UpdateZoneOK{}
}

/* UpdateZoneOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateZoneOK struct {
	Payload *models.Zone
}

func (o *UpdateZoneOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/zones/{id}][%d] updateZoneOK  %+v", 200, o.Payload)
}
func (o *UpdateZoneOK) GetPayload() *models.Zone {
	return o.Payload
}

func (o *UpdateZoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Zone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateZoneBadRequest creates a UpdateZoneBadRequest with default headers values
func NewUpdateZoneBadRequest() *UpdateZoneBadRequest {
	return &UpdateZoneBadRequest{}
}

/* UpdateZoneBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type UpdateZoneBadRequest struct {
	Payload *models.Error
}

func (o *UpdateZoneBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/zones/{id}][%d] updateZoneBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateZoneBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateZoneBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateZoneForbidden creates a UpdateZoneForbidden with default headers values
func NewUpdateZoneForbidden() *UpdateZoneForbidden {
	return &UpdateZoneForbidden{}
}

/* UpdateZoneForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateZoneForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *UpdateZoneForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/zones/{id}][%d] updateZoneForbidden  %+v", 403, o.Payload)
}
func (o *UpdateZoneForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *UpdateZoneForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
