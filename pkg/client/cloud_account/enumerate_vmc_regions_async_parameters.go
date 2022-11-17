// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

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

// NewEnumerateVmcRegionsAsyncParams creates a new EnumerateVmcRegionsAsyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnumerateVmcRegionsAsyncParams() *EnumerateVmcRegionsAsyncParams {
	return &EnumerateVmcRegionsAsyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnumerateVmcRegionsAsyncParamsWithTimeout creates a new EnumerateVmcRegionsAsyncParams object
// with the ability to set a timeout on a request.
func NewEnumerateVmcRegionsAsyncParamsWithTimeout(timeout time.Duration) *EnumerateVmcRegionsAsyncParams {
	return &EnumerateVmcRegionsAsyncParams{
		timeout: timeout,
	}
}

// NewEnumerateVmcRegionsAsyncParamsWithContext creates a new EnumerateVmcRegionsAsyncParams object
// with the ability to set a context for a request.
func NewEnumerateVmcRegionsAsyncParamsWithContext(ctx context.Context) *EnumerateVmcRegionsAsyncParams {
	return &EnumerateVmcRegionsAsyncParams{
		Context: ctx,
	}
}

// NewEnumerateVmcRegionsAsyncParamsWithHTTPClient creates a new EnumerateVmcRegionsAsyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnumerateVmcRegionsAsyncParamsWithHTTPClient(client *http.Client) *EnumerateVmcRegionsAsyncParams {
	return &EnumerateVmcRegionsAsyncParams{
		HTTPClient: client,
	}
}

/*
EnumerateVmcRegionsAsyncParams contains all the parameters to send to the API endpoint

	for the enumerate vmc regions async operation.

	Typically these are written to a http.Request.
*/
type EnumerateVmcRegionsAsyncParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion string

	/* Body.

	   VMC region enumeration specification
	*/
	Body *models.CloudAccountVmcRegionEnumerationSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enumerate vmc regions async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnumerateVmcRegionsAsyncParams) WithDefaults() *EnumerateVmcRegionsAsyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enumerate vmc regions async params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnumerateVmcRegionsAsyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enumerate vmc regions async params
func (o *EnumerateVmcRegionsAsyncParams) WithTimeout(timeout time.Duration) *EnumerateVmcRegionsAsyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enumerate vmc regions async params
func (o *EnumerateVmcRegionsAsyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enumerate vmc regions async params
func (o *EnumerateVmcRegionsAsyncParams) WithContext(ctx context.Context) *EnumerateVmcRegionsAsyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enumerate vmc regions async params
func (o *EnumerateVmcRegionsAsyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enumerate vmc regions async params
func (o *EnumerateVmcRegionsAsyncParams) WithHTTPClient(client *http.Client) *EnumerateVmcRegionsAsyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enumerate vmc regions async params
func (o *EnumerateVmcRegionsAsyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the enumerate vmc regions async params
func (o *EnumerateVmcRegionsAsyncParams) WithAPIVersion(aPIVersion string) *EnumerateVmcRegionsAsyncParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the enumerate vmc regions async params
func (o *EnumerateVmcRegionsAsyncParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the enumerate vmc regions async params
func (o *EnumerateVmcRegionsAsyncParams) WithBody(body *models.CloudAccountVmcRegionEnumerationSpecification) *EnumerateVmcRegionsAsyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the enumerate vmc regions async params
func (o *EnumerateVmcRegionsAsyncParams) SetBody(body *models.CloudAccountVmcRegionEnumerationSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EnumerateVmcRegionsAsyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {

		if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
			return err
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
