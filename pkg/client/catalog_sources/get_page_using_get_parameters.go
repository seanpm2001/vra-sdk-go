// Code generated by go-swagger; DO NOT EDIT.

package catalog_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPageUsingGETParams creates a new GetPageUsingGETParams object
// with the default values initialized.
func NewGetPageUsingGETParams() *GetPageUsingGETParams {
	var ()
	return &GetPageUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPageUsingGETParamsWithTimeout creates a new GetPageUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPageUsingGETParamsWithTimeout(timeout time.Duration) *GetPageUsingGETParams {
	var ()
	return &GetPageUsingGETParams{

		timeout: timeout,
	}
}

// NewGetPageUsingGETParamsWithContext creates a new GetPageUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPageUsingGETParamsWithContext(ctx context.Context) *GetPageUsingGETParams {
	var ()
	return &GetPageUsingGETParams{

		Context: ctx,
	}
}

// NewGetPageUsingGETParamsWithHTTPClient creates a new GetPageUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPageUsingGETParamsWithHTTPClient(client *http.Client) *GetPageUsingGETParams {
	var ()
	return &GetPageUsingGETParams{
		HTTPClient: client,
	}
}

/*GetPageUsingGETParams contains all the parameters to send to the API endpoint
for the get page using g e t operation typically these are written to a http.Request
*/
type GetPageUsingGETParams struct {

	/*DollarOrderby
	  Sorting criteria in the format: property (asc|desc). Default sort order is ascending. Multiple sort criteria are supported.

	*/
	DollarOrderby []string
	/*DollarSkip
	  Number of records you want to skip

	*/
	DollarSkip *int32
	/*DollarTop
	  Number of records you want

	*/
	DollarTop *int32
	/*ProjectID
	  Find sources which contains items that can be requested in the given projectId

	*/
	ProjectID *string
	/*Search
	  Matches will have this string in their name or description.

	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get page using get params
func (o *GetPageUsingGETParams) WithTimeout(timeout time.Duration) *GetPageUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get page using get params
func (o *GetPageUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get page using get params
func (o *GetPageUsingGETParams) WithContext(ctx context.Context) *GetPageUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get page using get params
func (o *GetPageUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get page using get params
func (o *GetPageUsingGETParams) WithHTTPClient(client *http.Client) *GetPageUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get page using get params
func (o *GetPageUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the get page using get params
func (o *GetPageUsingGETParams) WithDollarOrderby(dollarOrderby []string) *GetPageUsingGETParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the get page using get params
func (o *GetPageUsingGETParams) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSkip adds the dollarSkip to the get page using get params
func (o *GetPageUsingGETParams) WithDollarSkip(dollarSkip *int32) *GetPageUsingGETParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get page using get params
func (o *GetPageUsingGETParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get page using get params
func (o *GetPageUsingGETParams) WithDollarTop(dollarTop *int32) *GetPageUsingGETParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get page using get params
func (o *GetPageUsingGETParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithProjectID adds the projectID to the get page using get params
func (o *GetPageUsingGETParams) WithProjectID(projectID *string) *GetPageUsingGETParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get page using get params
func (o *GetPageUsingGETParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithSearch adds the search to the get page using get params
func (o *GetPageUsingGETParams) WithSearch(search *string) *GetPageUsingGETParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get page using get params
func (o *GetPageUsingGETParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetPageUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesDollarOrderby := o.DollarOrderby

	joinedDollarOrderby := swag.JoinByFormat(valuesDollarOrderby, "multi")
	// query array param $orderby
	if err := r.SetQueryParam("$orderby", joinedDollarOrderby...); err != nil {
		return err
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

	if o.ProjectID != nil {

		// query param projectId
		var qrProjectID string
		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID
		if qProjectID != "" {
			if err := r.SetQueryParam("projectId", qProjectID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
