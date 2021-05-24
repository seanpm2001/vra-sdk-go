// Code generated by go-swagger; DO NOT EDIT.

package deployment_actions

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

// NewGetResourceActionsUsingGETParams creates a new GetResourceActionsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResourceActionsUsingGETParams() *GetResourceActionsUsingGETParams {
	return &GetResourceActionsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceActionsUsingGETParamsWithTimeout creates a new GetResourceActionsUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetResourceActionsUsingGETParamsWithTimeout(timeout time.Duration) *GetResourceActionsUsingGETParams {
	return &GetResourceActionsUsingGETParams{
		timeout: timeout,
	}
}

// NewGetResourceActionsUsingGETParamsWithContext creates a new GetResourceActionsUsingGETParams object
// with the ability to set a context for a request.
func NewGetResourceActionsUsingGETParamsWithContext(ctx context.Context) *GetResourceActionsUsingGETParams {
	return &GetResourceActionsUsingGETParams{
		Context: ctx,
	}
}

// NewGetResourceActionsUsingGETParamsWithHTTPClient creates a new GetResourceActionsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResourceActionsUsingGETParamsWithHTTPClient(client *http.Client) *GetResourceActionsUsingGETParams {
	return &GetResourceActionsUsingGETParams{
		HTTPClient: client,
	}
}

/* GetResourceActionsUsingGETParams contains all the parameters to send to the API endpoint
   for the get resource actions using g e t operation.

   Typically these are written to a http.Request.
*/
type GetResourceActionsUsingGETParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* DepID.

	   Deployment ID

	   Format: uuid
	*/
	DepID strfmt.UUID

	/* ResourceID.

	   Resource ID

	   Format: uuid
	*/
	ResourceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get resource actions using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceActionsUsingGETParams) WithDefaults() *GetResourceActionsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resource actions using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourceActionsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) WithTimeout(timeout time.Duration) *GetResourceActionsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) WithContext(ctx context.Context) *GetResourceActionsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) WithHTTPClient(client *http.Client) *GetResourceActionsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) WithAPIVersion(aPIVersion *string) *GetResourceActionsUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDepID adds the depID to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) WithDepID(depID strfmt.UUID) *GetResourceActionsUsingGETParams {
	o.SetDepID(depID)
	return o
}

// SetDepID adds the depId to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) SetDepID(depID strfmt.UUID) {
	o.DepID = depID
}

// WithResourceID adds the resourceID to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) WithResourceID(resourceID strfmt.UUID) *GetResourceActionsUsingGETParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the get resource actions using get params
func (o *GetResourceActionsUsingGETParams) SetResourceID(resourceID strfmt.UUID) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceActionsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param depId
	if err := r.SetPathParam("depId", o.DepID.String()); err != nil {
		return err
	}

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
