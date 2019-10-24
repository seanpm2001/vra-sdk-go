// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteDeploymentUsingDELETEReader is a Reader for the DeleteDeploymentUsingDELETE structure.
type DeleteDeploymentUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeploymentUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDeploymentUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteDeploymentUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDeploymentUsingDELETEOK creates a DeleteDeploymentUsingDELETEOK with default headers values
func NewDeleteDeploymentUsingDELETEOK() *DeleteDeploymentUsingDELETEOK {
	return &DeleteDeploymentUsingDELETEOK{}
}

/*DeleteDeploymentUsingDELETEOK handles this case with default header values.

OK
*/
type DeleteDeploymentUsingDELETEOK struct {
	Payload *models.DeploymentRequest
}

func (o *DeleteDeploymentUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{depId}][%d] deleteDeploymentUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteDeploymentUsingDELETEOK) GetPayload() *models.DeploymentRequest {
	return o.Payload
}

func (o *DeleteDeploymentUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeploymentUsingDELETEUnauthorized creates a DeleteDeploymentUsingDELETEUnauthorized with default headers values
func NewDeleteDeploymentUsingDELETEUnauthorized() *DeleteDeploymentUsingDELETEUnauthorized {
	return &DeleteDeploymentUsingDELETEUnauthorized{}
}

/*DeleteDeploymentUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteDeploymentUsingDELETEUnauthorized struct {
}

func (o *DeleteDeploymentUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{depId}][%d] deleteDeploymentUsingDELETEUnauthorized ", 401)
}

func (o *DeleteDeploymentUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
