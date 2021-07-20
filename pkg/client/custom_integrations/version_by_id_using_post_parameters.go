// Code generated by go-swagger; DO NOT EDIT.

package custom_integrations

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

// NewVersionByIDUsingPOSTParams creates a new VersionByIDUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVersionByIDUsingPOSTParams() *VersionByIDUsingPOSTParams {
	return &VersionByIDUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVersionByIDUsingPOSTParamsWithTimeout creates a new VersionByIDUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewVersionByIDUsingPOSTParamsWithTimeout(timeout time.Duration) *VersionByIDUsingPOSTParams {
	return &VersionByIDUsingPOSTParams{
		timeout: timeout,
	}
}

// NewVersionByIDUsingPOSTParamsWithContext creates a new VersionByIDUsingPOSTParams object
// with the ability to set a context for a request.
func NewVersionByIDUsingPOSTParamsWithContext(ctx context.Context) *VersionByIDUsingPOSTParams {
	return &VersionByIDUsingPOSTParams{
		Context: ctx,
	}
}

// NewVersionByIDUsingPOSTParamsWithHTTPClient creates a new VersionByIDUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewVersionByIDUsingPOSTParamsWithHTTPClient(client *http.Client) *VersionByIDUsingPOSTParams {
	return &VersionByIDUsingPOSTParams{
		HTTPClient: client,
	}
}

/* VersionByIDUsingPOSTParams contains all the parameters to send to the API endpoint
   for the version by Id using p o s t operation.

   Typically these are written to a http.Request.
*/
type VersionByIDUsingPOSTParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Body.

	   Request object for version of a Custom Integration
	*/
	Body models.VersionRequest

	/* ID.

	   The ID of the Custom Integration
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the version by Id using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VersionByIDUsingPOSTParams) WithDefaults() *VersionByIDUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the version by Id using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VersionByIDUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the version by Id using p o s t params
func (o *VersionByIDUsingPOSTParams) WithTimeout(timeout time.Duration) *VersionByIDUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the version by Id using p o s t params
func (o *VersionByIDUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the version by Id using p o s t params
func (o *VersionByIDUsingPOSTParams) WithContext(ctx context.Context) *VersionByIDUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the version by Id using p o s t params
func (o *VersionByIDUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the version by Id using p o s t params
func (o *VersionByIDUsingPOSTParams) WithHTTPClient(client *http.Client) *VersionByIDUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the version by Id using p o s t params
func (o *VersionByIDUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the version by Id using p o s t params
func (o *VersionByIDUsingPOSTParams) WithAuthorization(authorization string) *VersionByIDUsingPOSTParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the version by Id using p o s t params
func (o *VersionByIDUsingPOSTParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the version by Id using p o s t params
func (o *VersionByIDUsingPOSTParams) WithAPIVersion(aPIVersion string) *VersionByIDUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the version by Id using p o s t params
func (o *VersionByIDUsingPOSTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the version by Id using p o s t params
func (o *VersionByIDUsingPOSTParams) WithBody(body models.VersionRequest) *VersionByIDUsingPOSTParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the version by Id using p o s t params
func (o *VersionByIDUsingPOSTParams) SetBody(body models.VersionRequest) {
	o.Body = body
}

// WithID adds the id to the version by Id using p o s t params
func (o *VersionByIDUsingPOSTParams) WithID(id string) *VersionByIDUsingPOSTParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the version by Id using p o s t params
func (o *VersionByIDUsingPOSTParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VersionByIDUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
