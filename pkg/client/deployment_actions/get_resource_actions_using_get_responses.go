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

// GetResourceActionsUsingGETReader is a Reader for the GetResourceActionsUsingGET structure.
type GetResourceActionsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceActionsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceActionsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetResourceActionsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResourceActionsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourceActionsUsingGETOK creates a GetResourceActionsUsingGETOK with default headers values
func NewGetResourceActionsUsingGETOK() *GetResourceActionsUsingGETOK {
	return &GetResourceActionsUsingGETOK{}
}

/* GetResourceActionsUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetResourceActionsUsingGETOK struct {
	Payload []*models.ResourceAction
}

func (o *GetResourceActionsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources/{resourceId}/actions][%d] getResourceActionsUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetResourceActionsUsingGETOK) GetPayload() []*models.ResourceAction {
	return o.Payload
}

func (o *GetResourceActionsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceActionsUsingGETUnauthorized creates a GetResourceActionsUsingGETUnauthorized with default headers values
func NewGetResourceActionsUsingGETUnauthorized() *GetResourceActionsUsingGETUnauthorized {
	return &GetResourceActionsUsingGETUnauthorized{}
}

/* GetResourceActionsUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetResourceActionsUsingGETUnauthorized struct {
}

func (o *GetResourceActionsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources/{resourceId}/actions][%d] getResourceActionsUsingGETUnauthorized ", 401)
}

func (o *GetResourceActionsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResourceActionsUsingGETNotFound creates a GetResourceActionsUsingGETNotFound with default headers values
func NewGetResourceActionsUsingGETNotFound() *GetResourceActionsUsingGETNotFound {
	return &GetResourceActionsUsingGETNotFound{}
}

/* GetResourceActionsUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetResourceActionsUsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetResourceActionsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources/{resourceId}/actions][%d] getResourceActionsUsingGETNotFound  %+v", 404, o.Payload)
}
func (o *GetResourceActionsUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetResourceActionsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
