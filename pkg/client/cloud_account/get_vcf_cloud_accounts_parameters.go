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

// NewGetVcfCloudAccountsParams creates a new GetVcfCloudAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVcfCloudAccountsParams() *GetVcfCloudAccountsParams {
	return &GetVcfCloudAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVcfCloudAccountsParamsWithTimeout creates a new GetVcfCloudAccountsParams object
// with the ability to set a timeout on a request.
func NewGetVcfCloudAccountsParamsWithTimeout(timeout time.Duration) *GetVcfCloudAccountsParams {
	return &GetVcfCloudAccountsParams{
		timeout: timeout,
	}
}

// NewGetVcfCloudAccountsParamsWithContext creates a new GetVcfCloudAccountsParams object
// with the ability to set a context for a request.
func NewGetVcfCloudAccountsParamsWithContext(ctx context.Context) *GetVcfCloudAccountsParams {
	return &GetVcfCloudAccountsParams{
		Context: ctx,
	}
}

// NewGetVcfCloudAccountsParamsWithHTTPClient creates a new GetVcfCloudAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVcfCloudAccountsParamsWithHTTPClient(client *http.Client) *GetVcfCloudAccountsParams {
	return &GetVcfCloudAccountsParams{
		HTTPClient: client,
	}
}

/* GetVcfCloudAccountsParams contains all the parameters to send to the API endpoint
   for the get vcf cloud accounts operation.

   Typically these are written to a http.Request.
*/
type GetVcfCloudAccountsParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vcf cloud accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcfCloudAccountsParams) WithDefaults() *GetVcfCloudAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vcf cloud accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVcfCloudAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vcf cloud accounts params
func (o *GetVcfCloudAccountsParams) WithTimeout(timeout time.Duration) *GetVcfCloudAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vcf cloud accounts params
func (o *GetVcfCloudAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vcf cloud accounts params
func (o *GetVcfCloudAccountsParams) WithContext(ctx context.Context) *GetVcfCloudAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vcf cloud accounts params
func (o *GetVcfCloudAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vcf cloud accounts params
func (o *GetVcfCloudAccountsParams) WithHTTPClient(client *http.Client) *GetVcfCloudAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vcf cloud accounts params
func (o *GetVcfCloudAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get vcf cloud accounts params
func (o *GetVcfCloudAccountsParams) WithAPIVersion(aPIVersion *string) *GetVcfCloudAccountsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get vcf cloud accounts params
func (o *GetVcfCloudAccountsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetVcfCloudAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
