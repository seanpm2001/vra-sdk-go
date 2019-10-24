// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewGetPolicyUsingGETParams creates a new GetPolicyUsingGETParams object
// with the default values initialized.
func NewGetPolicyUsingGETParams() *GetPolicyUsingGETParams {
	var ()
	return &GetPolicyUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPolicyUsingGETParamsWithTimeout creates a new GetPolicyUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPolicyUsingGETParamsWithTimeout(timeout time.Duration) *GetPolicyUsingGETParams {
	var ()
	return &GetPolicyUsingGETParams{

		timeout: timeout,
	}
}

// NewGetPolicyUsingGETParamsWithContext creates a new GetPolicyUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPolicyUsingGETParamsWithContext(ctx context.Context) *GetPolicyUsingGETParams {
	var ()
	return &GetPolicyUsingGETParams{

		Context: ctx,
	}
}

// NewGetPolicyUsingGETParamsWithHTTPClient creates a new GetPolicyUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPolicyUsingGETParamsWithHTTPClient(client *http.Client) *GetPolicyUsingGETParams {
	var ()
	return &GetPolicyUsingGETParams{
		HTTPClient: client,
	}
}

/*GetPolicyUsingGETParams contains all the parameters to send to the API endpoint
for the get policy using g e t operation typically these are written to a http.Request
*/
type GetPolicyUsingGETParams struct {

	/*ComputeStats
	  computeStats

	*/
	ComputeStats *bool
	/*ExpandCriteriaLegend
	  Augments Criteria in returned object with legend data structures for translation between IDs and human readable DataElements

	*/
	ExpandCriteriaLegend *bool
	/*ID
	  Policy ID

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get policy using get params
func (o *GetPolicyUsingGETParams) WithTimeout(timeout time.Duration) *GetPolicyUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get policy using get params
func (o *GetPolicyUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get policy using get params
func (o *GetPolicyUsingGETParams) WithContext(ctx context.Context) *GetPolicyUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get policy using get params
func (o *GetPolicyUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get policy using get params
func (o *GetPolicyUsingGETParams) WithHTTPClient(client *http.Client) *GetPolicyUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get policy using get params
func (o *GetPolicyUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComputeStats adds the computeStats to the get policy using get params
func (o *GetPolicyUsingGETParams) WithComputeStats(computeStats *bool) *GetPolicyUsingGETParams {
	o.SetComputeStats(computeStats)
	return o
}

// SetComputeStats adds the computeStats to the get policy using get params
func (o *GetPolicyUsingGETParams) SetComputeStats(computeStats *bool) {
	o.ComputeStats = computeStats
}

// WithExpandCriteriaLegend adds the expandCriteriaLegend to the get policy using get params
func (o *GetPolicyUsingGETParams) WithExpandCriteriaLegend(expandCriteriaLegend *bool) *GetPolicyUsingGETParams {
	o.SetExpandCriteriaLegend(expandCriteriaLegend)
	return o
}

// SetExpandCriteriaLegend adds the expandCriteriaLegend to the get policy using get params
func (o *GetPolicyUsingGETParams) SetExpandCriteriaLegend(expandCriteriaLegend *bool) {
	o.ExpandCriteriaLegend = expandCriteriaLegend
}

// WithID adds the id to the get policy using get params
func (o *GetPolicyUsingGETParams) WithID(id strfmt.UUID) *GetPolicyUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get policy using get params
func (o *GetPolicyUsingGETParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetPolicyUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ComputeStats != nil {

		// query param computeStats
		var qrComputeStats bool
		if o.ComputeStats != nil {
			qrComputeStats = *o.ComputeStats
		}
		qComputeStats := swag.FormatBool(qrComputeStats)
		if qComputeStats != "" {
			if err := r.SetQueryParam("computeStats", qComputeStats); err != nil {
				return err
			}
		}

	}

	if o.ExpandCriteriaLegend != nil {

		// query param expandCriteriaLegend
		var qrExpandCriteriaLegend bool
		if o.ExpandCriteriaLegend != nil {
			qrExpandCriteriaLegend = *o.ExpandCriteriaLegend
		}
		qExpandCriteriaLegend := swag.FormatBool(qrExpandCriteriaLegend)
		if qExpandCriteriaLegend != "" {
			if err := r.SetQueryParam("expandCriteriaLegend", qExpandCriteriaLegend); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
