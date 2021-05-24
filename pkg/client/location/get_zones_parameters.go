// Code generated by go-swagger; DO NOT EDIT.

package location

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

// NewGetZonesParams creates a new GetZonesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetZonesParams() *GetZonesParams {
	return &GetZonesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetZonesParamsWithTimeout creates a new GetZonesParams object
// with the ability to set a timeout on a request.
func NewGetZonesParamsWithTimeout(timeout time.Duration) *GetZonesParams {
	return &GetZonesParams{
		timeout: timeout,
	}
}

// NewGetZonesParamsWithContext creates a new GetZonesParams object
// with the ability to set a context for a request.
func NewGetZonesParamsWithContext(ctx context.Context) *GetZonesParams {
	return &GetZonesParams{
		Context: ctx,
	}
}

// NewGetZonesParamsWithHTTPClient creates a new GetZonesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetZonesParamsWithHTTPClient(client *http.Client) *GetZonesParams {
	return &GetZonesParams{
		HTTPClient: client,
	}
}

/* GetZonesParams contains all the parameters to send to the API endpoint
   for the get zones operation.

   Typically these are written to a http.Request.
*/
type GetZonesParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetZonesParams) WithDefaults() *GetZonesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get zones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetZonesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get zones params
func (o *GetZonesParams) WithTimeout(timeout time.Duration) *GetZonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get zones params
func (o *GetZonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get zones params
func (o *GetZonesParams) WithContext(ctx context.Context) *GetZonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get zones params
func (o *GetZonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get zones params
func (o *GetZonesParams) WithHTTPClient(client *http.Client) *GetZonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get zones params
func (o *GetZonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get zones params
func (o *GetZonesParams) WithAPIVersion(aPIVersion *string) *GetZonesParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get zones params
func (o *GetZonesParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetZonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
