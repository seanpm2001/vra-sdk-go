// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreateMachineParams creates a new CreateMachineParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateMachineParams() *CreateMachineParams {
	return &CreateMachineParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMachineParamsWithTimeout creates a new CreateMachineParams object
// with the ability to set a timeout on a request.
func NewCreateMachineParamsWithTimeout(timeout time.Duration) *CreateMachineParams {
	return &CreateMachineParams{
		timeout: timeout,
	}
}

// NewCreateMachineParamsWithContext creates a new CreateMachineParams object
// with the ability to set a context for a request.
func NewCreateMachineParamsWithContext(ctx context.Context) *CreateMachineParams {
	return &CreateMachineParams{
		Context: ctx,
	}
}

// NewCreateMachineParamsWithHTTPClient creates a new CreateMachineParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateMachineParamsWithHTTPClient(client *http.Client) *CreateMachineParams {
	return &CreateMachineParams{
		HTTPClient: client,
	}
}

/* CreateMachineParams contains all the parameters to send to the API endpoint
   for the create machine operation.

   Typically these are written to a http.Request.
*/
type CreateMachineParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   Machine Specification instance
	*/
	Body *models.MachineSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMachineParams) WithDefaults() *CreateMachineParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create machine params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMachineParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create machine params
func (o *CreateMachineParams) WithTimeout(timeout time.Duration) *CreateMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create machine params
func (o *CreateMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create machine params
func (o *CreateMachineParams) WithContext(ctx context.Context) *CreateMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create machine params
func (o *CreateMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create machine params
func (o *CreateMachineParams) WithHTTPClient(client *http.Client) *CreateMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create machine params
func (o *CreateMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create machine params
func (o *CreateMachineParams) WithAPIVersion(aPIVersion *string) *CreateMachineParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create machine params
func (o *CreateMachineParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create machine params
func (o *CreateMachineParams) WithBody(body *models.MachineSpecification) *CreateMachineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create machine params
func (o *CreateMachineParams) SetBody(body *models.MachineSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// query param apiVersion
		var qrAPIVersion string

		if o.APIVersion != nil {
			qrAPIVersion = *o.APIVersion
		}
		qAPIVersion := qrAPIVersion
		if qAPIVersion != "" {

			if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
				return err
			}
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
