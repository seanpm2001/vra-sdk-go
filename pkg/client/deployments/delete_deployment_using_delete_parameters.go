// Code generated by go-swagger; DO NOT EDIT.

package deployments

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

// NewDeleteDeploymentUsingDELETEParams creates a new DeleteDeploymentUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDeploymentUsingDELETEParams() *DeleteDeploymentUsingDELETEParams {
	return &DeleteDeploymentUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDeploymentUsingDELETEParamsWithTimeout creates a new DeleteDeploymentUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteDeploymentUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteDeploymentUsingDELETEParams {
	return &DeleteDeploymentUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteDeploymentUsingDELETEParamsWithContext creates a new DeleteDeploymentUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteDeploymentUsingDELETEParamsWithContext(ctx context.Context) *DeleteDeploymentUsingDELETEParams {
	return &DeleteDeploymentUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteDeploymentUsingDELETEParamsWithHTTPClient creates a new DeleteDeploymentUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDeploymentUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteDeploymentUsingDELETEParams {
	return &DeleteDeploymentUsingDELETEParams{
		HTTPClient: client,
	}
}

/* DeleteDeploymentUsingDELETEParams contains all the parameters to send to the API endpoint
   for the delete deployment using d e l e t e operation.

   Typically these are written to a http.Request.
*/
type DeleteDeploymentUsingDELETEParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* DepID.

	   Deployment ID

	   Format: uuid
	*/
	DepID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete deployment using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeploymentUsingDELETEParams) WithDefaults() *DeleteDeploymentUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete deployment using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDeploymentUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete deployment using d e l e t e params
func (o *DeleteDeploymentUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteDeploymentUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete deployment using d e l e t e params
func (o *DeleteDeploymentUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete deployment using d e l e t e params
func (o *DeleteDeploymentUsingDELETEParams) WithContext(ctx context.Context) *DeleteDeploymentUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete deployment using d e l e t e params
func (o *DeleteDeploymentUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete deployment using d e l e t e params
func (o *DeleteDeploymentUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteDeploymentUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete deployment using d e l e t e params
func (o *DeleteDeploymentUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete deployment using d e l e t e params
func (o *DeleteDeploymentUsingDELETEParams) WithAPIVersion(aPIVersion *string) *DeleteDeploymentUsingDELETEParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete deployment using d e l e t e params
func (o *DeleteDeploymentUsingDELETEParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDepID adds the depID to the delete deployment using d e l e t e params
func (o *DeleteDeploymentUsingDELETEParams) WithDepID(depID strfmt.UUID) *DeleteDeploymentUsingDELETEParams {
	o.SetDepID(depID)
	return o
}

// SetDepID adds the depId to the delete deployment using d e l e t e params
func (o *DeleteDeploymentUsingDELETEParams) SetDepID(depID strfmt.UUID) {
	o.DepID = depID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDeploymentUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
