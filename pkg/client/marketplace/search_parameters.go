// Code generated by go-swagger; DO NOT EDIT.

package marketplace

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

// NewSearchParams creates a new SearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchParams() *SearchParams {
	return &SearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchParamsWithTimeout creates a new SearchParams object
// with the ability to set a timeout on a request.
func NewSearchParamsWithTimeout(timeout time.Duration) *SearchParams {
	return &SearchParams{
		timeout: timeout,
	}
}

// NewSearchParamsWithContext creates a new SearchParams object
// with the ability to set a context for a request.
func NewSearchParamsWithContext(ctx context.Context) *SearchParams {
	return &SearchParams{
		Context: ctx,
	}
}

// NewSearchParamsWithHTTPClient creates a new SearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchParamsWithHTTPClient(client *http.Client) *SearchParams {
	return &SearchParams{
		HTTPClient: client,
	}
}

/* SearchParams contains all the parameters to send to the API endpoint
   for the search operation.

   Typically these are written to a http.Request.
*/
type SearchParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information, please refer to /content/api/about
	*/
	APIVersion *string

	/* Search.

	   Full-text search term
	*/
	Search *string

	/* SourceID.

	   Content Source Id

	   Format: uuid
	*/
	SourceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchParams) WithDefaults() *SearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search params
func (o *SearchParams) WithTimeout(timeout time.Duration) *SearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search params
func (o *SearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search params
func (o *SearchParams) WithContext(ctx context.Context) *SearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search params
func (o *SearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search params
func (o *SearchParams) WithHTTPClient(client *http.Client) *SearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search params
func (o *SearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the search params
func (o *SearchParams) WithAPIVersion(aPIVersion *string) *SearchParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the search params
func (o *SearchParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithSearch adds the search to the search params
func (o *SearchParams) WithSearch(search *string) *SearchParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the search params
func (o *SearchParams) SetSearch(search *string) {
	o.Search = search
}

// WithSourceID adds the sourceID to the search params
func (o *SearchParams) WithSourceID(sourceID strfmt.UUID) *SearchParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the search params
func (o *SearchParams) SetSourceID(sourceID strfmt.UUID) {
	o.SourceID = sourceID
}

// WriteToRequest writes these params to a swagger request
func (o *SearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Search != nil {

		// query param search
		var qrSearch string

		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {

			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
		}
	}

	// path param sourceId
	if err := r.SetPathParam("sourceId", o.SourceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
