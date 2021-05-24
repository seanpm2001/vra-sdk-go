// Code generated by go-swagger; DO NOT EDIT.

package blueprint

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

// NewGetBlueprintUsingGET1Params creates a new GetBlueprintUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBlueprintUsingGET1Params() *GetBlueprintUsingGET1Params {
	return &GetBlueprintUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBlueprintUsingGET1ParamsWithTimeout creates a new GetBlueprintUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetBlueprintUsingGET1ParamsWithTimeout(timeout time.Duration) *GetBlueprintUsingGET1Params {
	return &GetBlueprintUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetBlueprintUsingGET1ParamsWithContext creates a new GetBlueprintUsingGET1Params object
// with the ability to set a context for a request.
func NewGetBlueprintUsingGET1ParamsWithContext(ctx context.Context) *GetBlueprintUsingGET1Params {
	return &GetBlueprintUsingGET1Params{
		Context: ctx,
	}
}

// NewGetBlueprintUsingGET1ParamsWithHTTPClient creates a new GetBlueprintUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetBlueprintUsingGET1ParamsWithHTTPClient(client *http.Client) *GetBlueprintUsingGET1Params {
	return &GetBlueprintUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetBlueprintUsingGET1Params contains all the parameters to send to the API endpoint
   for the get blueprint using get1 operation.

   Typically these are written to a http.Request.
*/
type GetBlueprintUsingGET1Params struct {

	/* DollarSelect.

	   Fields to include in content. Use * to get all fields.
	*/
	DollarSelect []string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* BlueprintID.

	   blueprintId

	   Format: uuid
	*/
	BlueprintID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get blueprint using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBlueprintUsingGET1Params) WithDefaults() *GetBlueprintUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get blueprint using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBlueprintUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get blueprint using get1 params
func (o *GetBlueprintUsingGET1Params) WithTimeout(timeout time.Duration) *GetBlueprintUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get blueprint using get1 params
func (o *GetBlueprintUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get blueprint using get1 params
func (o *GetBlueprintUsingGET1Params) WithContext(ctx context.Context) *GetBlueprintUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get blueprint using get1 params
func (o *GetBlueprintUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get blueprint using get1 params
func (o *GetBlueprintUsingGET1Params) WithHTTPClient(client *http.Client) *GetBlueprintUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get blueprint using get1 params
func (o *GetBlueprintUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarSelect adds the dollarSelect to the get blueprint using get1 params
func (o *GetBlueprintUsingGET1Params) WithDollarSelect(dollarSelect []string) *GetBlueprintUsingGET1Params {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the get blueprint using get1 params
func (o *GetBlueprintUsingGET1Params) SetDollarSelect(dollarSelect []string) {
	o.DollarSelect = dollarSelect
}

// WithAPIVersion adds the aPIVersion to the get blueprint using get1 params
func (o *GetBlueprintUsingGET1Params) WithAPIVersion(aPIVersion *string) *GetBlueprintUsingGET1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get blueprint using get1 params
func (o *GetBlueprintUsingGET1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBlueprintID adds the blueprintID to the get blueprint using get1 params
func (o *GetBlueprintUsingGET1Params) WithBlueprintID(blueprintID strfmt.UUID) *GetBlueprintUsingGET1Params {
	o.SetBlueprintID(blueprintID)
	return o
}

// SetBlueprintID adds the blueprintId to the get blueprint using get1 params
func (o *GetBlueprintUsingGET1Params) SetBlueprintID(blueprintID strfmt.UUID) {
	o.BlueprintID = blueprintID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBlueprintUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarSelect != nil {

		// binding items for $select
		joinedDollarSelect := o.bindParamDollarSelect(reg)

		// query array param $select
		if err := r.SetQueryParam("$select", joinedDollarSelect...); err != nil {
			return err
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

	// path param blueprintId
	if err := r.SetPathParam("blueprintId", o.BlueprintID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetBlueprintUsingGET1 binds the parameter $select
func (o *GetBlueprintUsingGET1Params) bindParamDollarSelect(formats strfmt.Registry) []string {
	dollarSelectIR := o.DollarSelect

	var dollarSelectIC []string
	for _, dollarSelectIIR := range dollarSelectIR { // explode []string

		dollarSelectIIV := dollarSelectIIR // string as string
		dollarSelectIC = append(dollarSelectIC, dollarSelectIIV)
	}

	// items.CollectionFormat: "multi"
	dollarSelectIS := swag.JoinByFormat(dollarSelectIC, "multi")

	return dollarSelectIS
}
