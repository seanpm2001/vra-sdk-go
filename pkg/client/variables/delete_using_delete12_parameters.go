// Code generated by go-swagger; DO NOT EDIT.

package variables

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

// NewDeleteUsingDELETE12Params creates a new DeleteUsingDELETE12Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUsingDELETE12Params() *DeleteUsingDELETE12Params {
	return &DeleteUsingDELETE12Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUsingDELETE12ParamsWithTimeout creates a new DeleteUsingDELETE12Params object
// with the ability to set a timeout on a request.
func NewDeleteUsingDELETE12ParamsWithTimeout(timeout time.Duration) *DeleteUsingDELETE12Params {
	return &DeleteUsingDELETE12Params{
		timeout: timeout,
	}
}

// NewDeleteUsingDELETE12ParamsWithContext creates a new DeleteUsingDELETE12Params object
// with the ability to set a context for a request.
func NewDeleteUsingDELETE12ParamsWithContext(ctx context.Context) *DeleteUsingDELETE12Params {
	return &DeleteUsingDELETE12Params{
		Context: ctx,
	}
}

// NewDeleteUsingDELETE12ParamsWithHTTPClient creates a new DeleteUsingDELETE12Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUsingDELETE12ParamsWithHTTPClient(client *http.Client) *DeleteUsingDELETE12Params {
	return &DeleteUsingDELETE12Params{
		HTTPClient: client,
	}
}

/* DeleteUsingDELETE12Params contains all the parameters to send to the API endpoint
   for the delete using d e l e t e 12 operation.

   Typically these are written to a http.Request.
*/
type DeleteUsingDELETE12Params struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Name.

	   The name of the Variable
	*/
	Name string

	/* Project.

	   The project the Variable belongs to
	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete using d e l e t e 12 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsingDELETE12Params) WithDefaults() *DeleteUsingDELETE12Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete using d e l e t e 12 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsingDELETE12Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete using d e l e t e 12 params
func (o *DeleteUsingDELETE12Params) WithTimeout(timeout time.Duration) *DeleteUsingDELETE12Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete using d e l e t e 12 params
func (o *DeleteUsingDELETE12Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete using d e l e t e 12 params
func (o *DeleteUsingDELETE12Params) WithContext(ctx context.Context) *DeleteUsingDELETE12Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete using d e l e t e 12 params
func (o *DeleteUsingDELETE12Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete using d e l e t e 12 params
func (o *DeleteUsingDELETE12Params) WithHTTPClient(client *http.Client) *DeleteUsingDELETE12Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete using d e l e t e 12 params
func (o *DeleteUsingDELETE12Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete using d e l e t e 12 params
func (o *DeleteUsingDELETE12Params) WithAuthorization(authorization string) *DeleteUsingDELETE12Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete using d e l e t e 12 params
func (o *DeleteUsingDELETE12Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the delete using d e l e t e 12 params
func (o *DeleteUsingDELETE12Params) WithAPIVersion(aPIVersion string) *DeleteUsingDELETE12Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete using d e l e t e 12 params
func (o *DeleteUsingDELETE12Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithName adds the name to the delete using d e l e t e 12 params
func (o *DeleteUsingDELETE12Params) WithName(name string) *DeleteUsingDELETE12Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete using d e l e t e 12 params
func (o *DeleteUsingDELETE12Params) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the delete using d e l e t e 12 params
func (o *DeleteUsingDELETE12Params) WithProject(project string) *DeleteUsingDELETE12Params {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the delete using d e l e t e 12 params
func (o *DeleteUsingDELETE12Params) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUsingDELETE12Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion

	if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
