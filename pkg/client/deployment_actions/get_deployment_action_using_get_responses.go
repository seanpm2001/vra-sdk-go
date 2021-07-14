// Code generated by go-swagger; DO NOT EDIT.

package deployment_actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetDeploymentActionUsingGETReader is a Reader for the GetDeploymentActionUsingGET structure.
type GetDeploymentActionUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentActionUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentActionUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDeploymentActionUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeploymentActionUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentActionUsingGETOK creates a GetDeploymentActionUsingGETOK with default headers values
func NewGetDeploymentActionUsingGETOK() *GetDeploymentActionUsingGETOK {
	return &GetDeploymentActionUsingGETOK{}
}

/* GetDeploymentActionUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetDeploymentActionUsingGETOK struct {
	Payload *models.ResourceAction
}

func (o *GetDeploymentActionUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/actions/{actionId}][%d] getDeploymentActionUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetDeploymentActionUsingGETOK) GetPayload() *models.ResourceAction {
	return o.Payload
}

func (o *GetDeploymentActionUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceAction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentActionUsingGETUnauthorized creates a GetDeploymentActionUsingGETUnauthorized with default headers values
func NewGetDeploymentActionUsingGETUnauthorized() *GetDeploymentActionUsingGETUnauthorized {
	return &GetDeploymentActionUsingGETUnauthorized{}
}

/* GetDeploymentActionUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetDeploymentActionUsingGETUnauthorized struct {
}

func (o *GetDeploymentActionUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/actions/{actionId}][%d] getDeploymentActionUsingGETUnauthorized ", 401)
}

func (o *GetDeploymentActionUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeploymentActionUsingGETNotFound creates a GetDeploymentActionUsingGETNotFound with default headers values
func NewGetDeploymentActionUsingGETNotFound() *GetDeploymentActionUsingGETNotFound {
	return &GetDeploymentActionUsingGETNotFound{}
}

/* GetDeploymentActionUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDeploymentActionUsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetDeploymentActionUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/actions/{actionId}][%d] getDeploymentActionUsingGETNotFound  %+v", 404, o.Payload)
}
func (o *GetDeploymentActionUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDeploymentActionUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
