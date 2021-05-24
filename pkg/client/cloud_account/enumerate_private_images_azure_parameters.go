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
)

// NewEnumeratePrivateImagesAzureParams creates a new EnumeratePrivateImagesAzureParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnumeratePrivateImagesAzureParams() *EnumeratePrivateImagesAzureParams {
	return &EnumeratePrivateImagesAzureParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnumeratePrivateImagesAzureParamsWithTimeout creates a new EnumeratePrivateImagesAzureParams object
// with the ability to set a timeout on a request.
func NewEnumeratePrivateImagesAzureParamsWithTimeout(timeout time.Duration) *EnumeratePrivateImagesAzureParams {
	return &EnumeratePrivateImagesAzureParams{
		timeout: timeout,
	}
}

// NewEnumeratePrivateImagesAzureParamsWithContext creates a new EnumeratePrivateImagesAzureParams object
// with the ability to set a context for a request.
func NewEnumeratePrivateImagesAzureParamsWithContext(ctx context.Context) *EnumeratePrivateImagesAzureParams {
	return &EnumeratePrivateImagesAzureParams{
		Context: ctx,
	}
}

// NewEnumeratePrivateImagesAzureParamsWithHTTPClient creates a new EnumeratePrivateImagesAzureParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnumeratePrivateImagesAzureParamsWithHTTPClient(client *http.Client) *EnumeratePrivateImagesAzureParams {
	return &EnumeratePrivateImagesAzureParams{
		HTTPClient: client,
	}
}

/* EnumeratePrivateImagesAzureParams contains all the parameters to send to the API endpoint
   for the enumerate private images azure operation.

   Typically these are written to a http.Request.
*/
type EnumeratePrivateImagesAzureParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   Id of Azure cloud account to enumerate
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enumerate private images azure params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnumeratePrivateImagesAzureParams) WithDefaults() *EnumeratePrivateImagesAzureParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enumerate private images azure params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnumeratePrivateImagesAzureParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enumerate private images azure params
func (o *EnumeratePrivateImagesAzureParams) WithTimeout(timeout time.Duration) *EnumeratePrivateImagesAzureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enumerate private images azure params
func (o *EnumeratePrivateImagesAzureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enumerate private images azure params
func (o *EnumeratePrivateImagesAzureParams) WithContext(ctx context.Context) *EnumeratePrivateImagesAzureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enumerate private images azure params
func (o *EnumeratePrivateImagesAzureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enumerate private images azure params
func (o *EnumeratePrivateImagesAzureParams) WithHTTPClient(client *http.Client) *EnumeratePrivateImagesAzureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enumerate private images azure params
func (o *EnumeratePrivateImagesAzureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the enumerate private images azure params
func (o *EnumeratePrivateImagesAzureParams) WithAPIVersion(aPIVersion *string) *EnumeratePrivateImagesAzureParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the enumerate private images azure params
func (o *EnumeratePrivateImagesAzureParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the enumerate private images azure params
func (o *EnumeratePrivateImagesAzureParams) WithID(id string) *EnumeratePrivateImagesAzureParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the enumerate private images azure params
func (o *EnumeratePrivateImagesAzureParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EnumeratePrivateImagesAzureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
