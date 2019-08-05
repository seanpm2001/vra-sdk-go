// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteAzureCloudAccountReader is a Reader for the DeleteAzureCloudAccount structure.
type DeleteAzureCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAzureCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAzureCloudAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteAzureCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAzureCloudAccountNoContent creates a DeleteAzureCloudAccountNoContent with default headers values
func NewDeleteAzureCloudAccountNoContent() *DeleteAzureCloudAccountNoContent {
	return &DeleteAzureCloudAccountNoContent{}
}

/*DeleteAzureCloudAccountNoContent handles this case with default header values.

No Content
*/
type DeleteAzureCloudAccountNoContent struct {
}

func (o *DeleteAzureCloudAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts-azure/{id}][%d] deleteAzureCloudAccountNoContent ", 204)
}

func (o *DeleteAzureCloudAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAzureCloudAccountForbidden creates a DeleteAzureCloudAccountForbidden with default headers values
func NewDeleteAzureCloudAccountForbidden() *DeleteAzureCloudAccountForbidden {
	return &DeleteAzureCloudAccountForbidden{}
}

/*DeleteAzureCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type DeleteAzureCloudAccountForbidden struct {
}

func (o *DeleteAzureCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts-azure/{id}][%d] deleteAzureCloudAccountForbidden ", 403)
}

func (o *DeleteAzureCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
