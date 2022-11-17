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

// NewDeleteVmcCloudAccountParams creates a new DeleteVmcCloudAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVmcCloudAccountParams() *DeleteVmcCloudAccountParams {
	return &DeleteVmcCloudAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVmcCloudAccountParamsWithTimeout creates a new DeleteVmcCloudAccountParams object
// with the ability to set a timeout on a request.
func NewDeleteVmcCloudAccountParamsWithTimeout(timeout time.Duration) *DeleteVmcCloudAccountParams {
	return &DeleteVmcCloudAccountParams{
		timeout: timeout,
	}
}

// NewDeleteVmcCloudAccountParamsWithContext creates a new DeleteVmcCloudAccountParams object
// with the ability to set a context for a request.
func NewDeleteVmcCloudAccountParamsWithContext(ctx context.Context) *DeleteVmcCloudAccountParams {
	return &DeleteVmcCloudAccountParams{
		Context: ctx,
	}
}

// NewDeleteVmcCloudAccountParamsWithHTTPClient creates a new DeleteVmcCloudAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVmcCloudAccountParamsWithHTTPClient(client *http.Client) *DeleteVmcCloudAccountParams {
	return &DeleteVmcCloudAccountParams{
		HTTPClient: client,
	}
}

/*
DeleteVmcCloudAccountParams contains all the parameters to send to the API endpoint

	for the delete vmc cloud account operation.

	Typically these are written to a http.Request.
*/
type DeleteVmcCloudAccountParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion string

	/* ID.

	   The ID of the Cloud Account
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete vmc cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVmcCloudAccountParams) WithDefaults() *DeleteVmcCloudAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete vmc cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVmcCloudAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete vmc cloud account params
func (o *DeleteVmcCloudAccountParams) WithTimeout(timeout time.Duration) *DeleteVmcCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vmc cloud account params
func (o *DeleteVmcCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vmc cloud account params
func (o *DeleteVmcCloudAccountParams) WithContext(ctx context.Context) *DeleteVmcCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vmc cloud account params
func (o *DeleteVmcCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vmc cloud account params
func (o *DeleteVmcCloudAccountParams) WithHTTPClient(client *http.Client) *DeleteVmcCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vmc cloud account params
func (o *DeleteVmcCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete vmc cloud account params
func (o *DeleteVmcCloudAccountParams) WithAPIVersion(aPIVersion string) *DeleteVmcCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete vmc cloud account params
func (o *DeleteVmcCloudAccountParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete vmc cloud account params
func (o *DeleteVmcCloudAccountParams) WithID(id string) *DeleteVmcCloudAccountParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete vmc cloud account params
func (o *DeleteVmcCloudAccountParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVmcCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
