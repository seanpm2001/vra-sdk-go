// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteAzureStorageProfileReader is a Reader for the DeleteAzureStorageProfile structure.
type DeleteAzureStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAzureStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAzureStorageProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteAzureStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAzureStorageProfileNoContent creates a DeleteAzureStorageProfileNoContent with default headers values
func NewDeleteAzureStorageProfileNoContent() *DeleteAzureStorageProfileNoContent {
	return &DeleteAzureStorageProfileNoContent{}
}

/*DeleteAzureStorageProfileNoContent handles this case with default header values.

No Content
*/
type DeleteAzureStorageProfileNoContent struct {
}

func (o *DeleteAzureStorageProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/storage-profiles-azure/{id}][%d] deleteAzureStorageProfileNoContent ", 204)
}

func (o *DeleteAzureStorageProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAzureStorageProfileForbidden creates a DeleteAzureStorageProfileForbidden with default headers values
func NewDeleteAzureStorageProfileForbidden() *DeleteAzureStorageProfileForbidden {
	return &DeleteAzureStorageProfileForbidden{}
}

/*DeleteAzureStorageProfileForbidden handles this case with default header values.

Forbidden
*/
type DeleteAzureStorageProfileForbidden struct {
}

func (o *DeleteAzureStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/storage-profiles-azure/{id}][%d] deleteAzureStorageProfileForbidden ", 403)
}

func (o *DeleteAzureStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
