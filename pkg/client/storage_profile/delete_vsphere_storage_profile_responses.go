// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteVSphereStorageProfileReader is a Reader for the DeleteVSphereStorageProfile structure.
type DeleteVSphereStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVSphereStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteVSphereStorageProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteVSphereStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVSphereStorageProfileNoContent creates a DeleteVSphereStorageProfileNoContent with default headers values
func NewDeleteVSphereStorageProfileNoContent() *DeleteVSphereStorageProfileNoContent {
	return &DeleteVSphereStorageProfileNoContent{}
}

/* DeleteVSphereStorageProfileNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteVSphereStorageProfileNoContent struct {
}

func (o *DeleteVSphereStorageProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/storage-profiles-vsphere/{id}][%d] deleteVSphereStorageProfileNoContent ", 204)
}

func (o *DeleteVSphereStorageProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteVSphereStorageProfileForbidden creates a DeleteVSphereStorageProfileForbidden with default headers values
func NewDeleteVSphereStorageProfileForbidden() *DeleteVSphereStorageProfileForbidden {
	return &DeleteVSphereStorageProfileForbidden{}
}

/* DeleteVSphereStorageProfileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteVSphereStorageProfileForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *DeleteVSphereStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/storage-profiles-vsphere/{id}][%d] deleteVSphereStorageProfileForbidden  %+v", 403, o.Payload)
}
func (o *DeleteVSphereStorageProfileForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *DeleteVSphereStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
