// Code generated by go-swagger; DO NOT EDIT.

package catalog_items

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

// NewGetVersionByIDUsingGETParams creates a new GetVersionByIDUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVersionByIDUsingGETParams() *GetVersionByIDUsingGETParams {
	return &GetVersionByIDUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVersionByIDUsingGETParamsWithTimeout creates a new GetVersionByIDUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetVersionByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetVersionByIDUsingGETParams {
	return &GetVersionByIDUsingGETParams{
		timeout: timeout,
	}
}

// NewGetVersionByIDUsingGETParamsWithContext creates a new GetVersionByIDUsingGETParams object
// with the ability to set a context for a request.
func NewGetVersionByIDUsingGETParamsWithContext(ctx context.Context) *GetVersionByIDUsingGETParams {
	return &GetVersionByIDUsingGETParams{
		Context: ctx,
	}
}

// NewGetVersionByIDUsingGETParamsWithHTTPClient creates a new GetVersionByIDUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVersionByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetVersionByIDUsingGETParams {
	return &GetVersionByIDUsingGETParams{
		HTTPClient: client,
	}
}

/* GetVersionByIDUsingGETParams contains all the parameters to send to the API endpoint
   for the get version by Id using g e t operation.

   Typically these are written to a http.Request.
*/
type GetVersionByIDUsingGETParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* ID.

	   Catalog Item ID

	   Format: uuid
	*/
	ID strfmt.UUID

	/* VersionID.

	   Catalog Item Version ID
	*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get version by Id using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionByIDUsingGETParams) WithDefaults() *GetVersionByIDUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get version by Id using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionByIDUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get version by Id using get params
func (o *GetVersionByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetVersionByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get version by Id using get params
func (o *GetVersionByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get version by Id using get params
func (o *GetVersionByIDUsingGETParams) WithContext(ctx context.Context) *GetVersionByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get version by Id using get params
func (o *GetVersionByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get version by Id using get params
func (o *GetVersionByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetVersionByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get version by Id using get params
func (o *GetVersionByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get version by Id using get params
func (o *GetVersionByIDUsingGETParams) WithAPIVersion(aPIVersion *string) *GetVersionByIDUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get version by Id using get params
func (o *GetVersionByIDUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get version by Id using get params
func (o *GetVersionByIDUsingGETParams) WithID(id strfmt.UUID) *GetVersionByIDUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get version by Id using get params
func (o *GetVersionByIDUsingGETParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithVersionID adds the versionID to the get version by Id using get params
func (o *GetVersionByIDUsingGETParams) WithVersionID(versionID string) *GetVersionByIDUsingGETParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get version by Id using get params
func (o *GetVersionByIDUsingGETParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetVersionByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param versionId
	if err := r.SetPathParam("versionId", o.VersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
