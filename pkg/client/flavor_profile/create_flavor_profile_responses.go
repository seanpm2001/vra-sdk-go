// Code generated by go-swagger; DO NOT EDIT.

package flavor_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateFlavorProfileReader is a Reader for the CreateFlavorProfile structure.
type CreateFlavorProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFlavorProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateFlavorProfileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateFlavorProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateFlavorProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateFlavorProfileCreated creates a CreateFlavorProfileCreated with default headers values
func NewCreateFlavorProfileCreated() *CreateFlavorProfileCreated {
	return &CreateFlavorProfileCreated{}
}

/* CreateFlavorProfileCreated describes a response with status code 201, with default header values.

successful operation
*/
type CreateFlavorProfileCreated struct {
	Payload *models.FlavorProfile
}

func (o *CreateFlavorProfileCreated) Error() string {
	return fmt.Sprintf("[POST /iaas/api/flavor-profiles][%d] createFlavorProfileCreated  %+v", 201, o.Payload)
}
func (o *CreateFlavorProfileCreated) GetPayload() *models.FlavorProfile {
	return o.Payload
}

func (o *CreateFlavorProfileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlavorProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFlavorProfileBadRequest creates a CreateFlavorProfileBadRequest with default headers values
func NewCreateFlavorProfileBadRequest() *CreateFlavorProfileBadRequest {
	return &CreateFlavorProfileBadRequest{}
}

/* CreateFlavorProfileBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type CreateFlavorProfileBadRequest struct {
	Payload *models.Error
}

func (o *CreateFlavorProfileBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/flavor-profiles][%d] createFlavorProfileBadRequest  %+v", 400, o.Payload)
}
func (o *CreateFlavorProfileBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateFlavorProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFlavorProfileForbidden creates a CreateFlavorProfileForbidden with default headers values
func NewCreateFlavorProfileForbidden() *CreateFlavorProfileForbidden {
	return &CreateFlavorProfileForbidden{}
}

/* CreateFlavorProfileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateFlavorProfileForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *CreateFlavorProfileForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/flavor-profiles][%d] createFlavorProfileForbidden  %+v", 403, o.Payload)
}
func (o *CreateFlavorProfileForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *CreateFlavorProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
