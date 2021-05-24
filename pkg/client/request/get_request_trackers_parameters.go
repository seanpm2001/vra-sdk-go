// Code generated by go-swagger; DO NOT EDIT.

package request

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

// NewGetRequestTrackersParams creates a new GetRequestTrackersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRequestTrackersParams() *GetRequestTrackersParams {
	return &GetRequestTrackersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRequestTrackersParamsWithTimeout creates a new GetRequestTrackersParams object
// with the ability to set a timeout on a request.
func NewGetRequestTrackersParamsWithTimeout(timeout time.Duration) *GetRequestTrackersParams {
	return &GetRequestTrackersParams{
		timeout: timeout,
	}
}

// NewGetRequestTrackersParamsWithContext creates a new GetRequestTrackersParams object
// with the ability to set a context for a request.
func NewGetRequestTrackersParamsWithContext(ctx context.Context) *GetRequestTrackersParams {
	return &GetRequestTrackersParams{
		Context: ctx,
	}
}

// NewGetRequestTrackersParamsWithHTTPClient creates a new GetRequestTrackersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRequestTrackersParamsWithHTTPClient(client *http.Client) *GetRequestTrackersParams {
	return &GetRequestTrackersParams{
		HTTPClient: client,
	}
}

/* GetRequestTrackersParams contains all the parameters to send to the API endpoint
   for the get request trackers operation.

   Typically these are written to a http.Request.
*/
type GetRequestTrackersParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get request trackers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRequestTrackersParams) WithDefaults() *GetRequestTrackersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get request trackers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRequestTrackersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get request trackers params
func (o *GetRequestTrackersParams) WithTimeout(timeout time.Duration) *GetRequestTrackersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get request trackers params
func (o *GetRequestTrackersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get request trackers params
func (o *GetRequestTrackersParams) WithContext(ctx context.Context) *GetRequestTrackersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get request trackers params
func (o *GetRequestTrackersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get request trackers params
func (o *GetRequestTrackersParams) WithHTTPClient(client *http.Client) *GetRequestTrackersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get request trackers params
func (o *GetRequestTrackersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get request trackers params
func (o *GetRequestTrackersParams) WithAPIVersion(aPIVersion *string) *GetRequestTrackersParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get request trackers params
func (o *GetRequestTrackersParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetRequestTrackersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
