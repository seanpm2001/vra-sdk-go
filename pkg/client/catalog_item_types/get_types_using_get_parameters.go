// Code generated by go-swagger; DO NOT EDIT.

package catalog_item_types

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
	"github.com/go-openapi/swag"
)

// NewGetTypesUsingGETParams creates a new GetTypesUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTypesUsingGETParams() *GetTypesUsingGETParams {
	return &GetTypesUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTypesUsingGETParamsWithTimeout creates a new GetTypesUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetTypesUsingGETParamsWithTimeout(timeout time.Duration) *GetTypesUsingGETParams {
	return &GetTypesUsingGETParams{
		timeout: timeout,
	}
}

// NewGetTypesUsingGETParamsWithContext creates a new GetTypesUsingGETParams object
// with the ability to set a context for a request.
func NewGetTypesUsingGETParamsWithContext(ctx context.Context) *GetTypesUsingGETParams {
	return &GetTypesUsingGETParams{
		Context: ctx,
	}
}

// NewGetTypesUsingGETParamsWithHTTPClient creates a new GetTypesUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTypesUsingGETParamsWithHTTPClient(client *http.Client) *GetTypesUsingGETParams {
	return &GetTypesUsingGETParams{
		HTTPClient: client,
	}
}

/* GetTypesUsingGETParams contains all the parameters to send to the API endpoint
   for the get types using g e t operation.

   Typically these are written to a http.Request.
*/
type GetTypesUsingGETParams struct {

	/* DollarOrderby.

	   Sorting criteria in the format: property (asc|desc). Default sort order is ascending. Multiple sort criteria are supported.
	*/
	DollarOrderby []string

	/* DollarSkip.

	   Number of records you want to skip

	   Format: int32
	*/
	DollarSkip *int32

	/* DollarTop.

	   Number of records you want

	   Format: int32
	*/
	DollarTop *int32

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get types using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTypesUsingGETParams) WithDefaults() *GetTypesUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get types using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTypesUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get types using get params
func (o *GetTypesUsingGETParams) WithTimeout(timeout time.Duration) *GetTypesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get types using get params
func (o *GetTypesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get types using get params
func (o *GetTypesUsingGETParams) WithContext(ctx context.Context) *GetTypesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get types using get params
func (o *GetTypesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get types using get params
func (o *GetTypesUsingGETParams) WithHTTPClient(client *http.Client) *GetTypesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get types using get params
func (o *GetTypesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the get types using get params
func (o *GetTypesUsingGETParams) WithDollarOrderby(dollarOrderby []string) *GetTypesUsingGETParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get types using get params
func (o *GetTypesUsingGETParams) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the get types using get params
func (o *GetTypesUsingGETParams) WithDollarSkip(dollarSkip *int32) *GetTypesUsingGETParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get types using get params
func (o *GetTypesUsingGETParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get types using get params
func (o *GetTypesUsingGETParams) WithDollarTop(dollarTop *int32) *GetTypesUsingGETParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get types using get params
func (o *GetTypesUsingGETParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get types using get params
func (o *GetTypesUsingGETParams) WithAPIVersion(aPIVersion *string) *GetTypesUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get types using get params
func (o *GetTypesUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetTypesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarOrderby != nil {

		// binding items for $orderby
		joinedDollarOrderby := o.bindParamDollarOrderby(reg)

		// query array param $orderby
		if err := r.SetQueryParam("$orderby", joinedDollarOrderby...); err != nil {
			return err
		}
	}

	if o.DollarSkip != nil {

		// query param $skip
		var qrDollarSkip int32

		if o.DollarSkip != nil {
			qrDollarSkip = *o.DollarSkip
		}
		qDollarSkip := swag.FormatInt32(qrDollarSkip)
		if qDollarSkip != "" {

			if err := r.SetQueryParam("$skip", qDollarSkip); err != nil {
				return err
			}
		}
	}

	if o.DollarTop != nil {

		// query param $top
		var qrDollarTop int32

		if o.DollarTop != nil {
			qrDollarTop = *o.DollarTop
		}
		qDollarTop := swag.FormatInt32(qrDollarTop)
		if qDollarTop != "" {

			if err := r.SetQueryParam("$top", qDollarTop); err != nil {
				return err
			}
		}
	}

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

// bindParamGetTypesUsingGET binds the parameter $orderby
func (o *GetTypesUsingGETParams) bindParamDollarOrderby(formats strfmt.Registry) []string {
	dollarOrderbyIR := o.DollarOrderby

	var dollarOrderbyIC []string
	for _, dollarOrderbyIIR := range dollarOrderbyIR { // explode []string

		dollarOrderbyIIV := dollarOrderbyIIR // string as string
		dollarOrderbyIC = append(dollarOrderbyIC, dollarOrderbyIIV)
	}

	// items.CollectionFormat: "multi"
	dollarOrderbyIS := swag.JoinByFormat(dollarOrderbyIC, "multi")

	return dollarOrderbyIS
}
