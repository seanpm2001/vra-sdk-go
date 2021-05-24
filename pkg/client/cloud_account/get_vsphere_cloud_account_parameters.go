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

// NewGetVSphereCloudAccountParams creates a new GetVSphereCloudAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVSphereCloudAccountParams() *GetVSphereCloudAccountParams {
	return &GetVSphereCloudAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVSphereCloudAccountParamsWithTimeout creates a new GetVSphereCloudAccountParams object
// with the ability to set a timeout on a request.
func NewGetVSphereCloudAccountParamsWithTimeout(timeout time.Duration) *GetVSphereCloudAccountParams {
	return &GetVSphereCloudAccountParams{
		timeout: timeout,
	}
}

// NewGetVSphereCloudAccountParamsWithContext creates a new GetVSphereCloudAccountParams object
// with the ability to set a context for a request.
func NewGetVSphereCloudAccountParamsWithContext(ctx context.Context) *GetVSphereCloudAccountParams {
	return &GetVSphereCloudAccountParams{
		Context: ctx,
	}
}

// NewGetVSphereCloudAccountParamsWithHTTPClient creates a new GetVSphereCloudAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVSphereCloudAccountParamsWithHTTPClient(client *http.Client) *GetVSphereCloudAccountParams {
	return &GetVSphereCloudAccountParams{
		HTTPClient: client,
	}
}

/* GetVSphereCloudAccountParams contains all the parameters to send to the API endpoint
   for the get v sphere cloud account operation.

   Typically these are written to a http.Request.
*/
type GetVSphereCloudAccountParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of the Cloud Account
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v sphere cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVSphereCloudAccountParams) WithDefaults() *GetVSphereCloudAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v sphere cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVSphereCloudAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v sphere cloud account params
func (o *GetVSphereCloudAccountParams) WithTimeout(timeout time.Duration) *GetVSphereCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v sphere cloud account params
func (o *GetVSphereCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v sphere cloud account params
func (o *GetVSphereCloudAccountParams) WithContext(ctx context.Context) *GetVSphereCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v sphere cloud account params
func (o *GetVSphereCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v sphere cloud account params
func (o *GetVSphereCloudAccountParams) WithHTTPClient(client *http.Client) *GetVSphereCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v sphere cloud account params
func (o *GetVSphereCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get v sphere cloud account params
func (o *GetVSphereCloudAccountParams) WithAPIVersion(aPIVersion *string) *GetVSphereCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get v sphere cloud account params
func (o *GetVSphereCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get v sphere cloud account params
func (o *GetVSphereCloudAccountParams) WithID(id string) *GetVSphereCloudAccountParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v sphere cloud account params
func (o *GetVSphereCloudAccountParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVSphereCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
