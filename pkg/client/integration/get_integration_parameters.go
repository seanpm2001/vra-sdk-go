// Code generated by go-swagger; DO NOT EDIT.

package integration

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

// NewGetIntegrationParams creates a new GetIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIntegrationParams() *GetIntegrationParams {
	return &GetIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIntegrationParamsWithTimeout creates a new GetIntegrationParams object
// with the ability to set a timeout on a request.
func NewGetIntegrationParamsWithTimeout(timeout time.Duration) *GetIntegrationParams {
	return &GetIntegrationParams{
		timeout: timeout,
	}
}

// NewGetIntegrationParamsWithContext creates a new GetIntegrationParams object
// with the ability to set a context for a request.
func NewGetIntegrationParamsWithContext(ctx context.Context) *GetIntegrationParams {
	return &GetIntegrationParams{
		Context: ctx,
	}
}

// NewGetIntegrationParamsWithHTTPClient creates a new GetIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIntegrationParamsWithHTTPClient(client *http.Client) *GetIntegrationParams {
	return &GetIntegrationParams{
		HTTPClient: client,
	}
}

/*
GetIntegrationParams contains all the parameters to send to the API endpoint

	for the get integration operation.

	Typically these are written to a http.Request.
*/
type GetIntegrationParams struct {

	/* DollarSelect.

	   Select a subset of properties to include in the response.
	*/
	DollarSelect *string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion string

	/* ID.

	   The ID of the Integration
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationParams) WithDefaults() *GetIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get integration params
func (o *GetIntegrationParams) WithTimeout(timeout time.Duration) *GetIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get integration params
func (o *GetIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get integration params
func (o *GetIntegrationParams) WithContext(ctx context.Context) *GetIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get integration params
func (o *GetIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get integration params
func (o *GetIntegrationParams) WithHTTPClient(client *http.Client) *GetIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get integration params
func (o *GetIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarSelect adds the dollarSelect to the get integration params
func (o *GetIntegrationParams) WithDollarSelect(dollarSelect *string) *GetIntegrationParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the get integration params
func (o *GetIntegrationParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
}

// WithAPIVersion adds the aPIVersion to the get integration params
func (o *GetIntegrationParams) WithAPIVersion(aPIVersion string) *GetIntegrationParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get integration params
func (o *GetIntegrationParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get integration params
func (o *GetIntegrationParams) WithID(id string) *GetIntegrationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get integration params
func (o *GetIntegrationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarSelect != nil {

		// query param $select
		var qrDollarSelect string

		if o.DollarSelect != nil {
			qrDollarSelect = *o.DollarSelect
		}
		qDollarSelect := qrDollarSelect
		if qDollarSelect != "" {

			if err := r.SetQueryParam("$select", qDollarSelect); err != nil {
				return err
			}
		}
	}

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
