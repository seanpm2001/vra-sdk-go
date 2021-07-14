// Code generated by go-swagger; DO NOT EDIT.

package requests

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

// NewGetEventLogsContentUsingGETParams creates a new GetEventLogsContentUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEventLogsContentUsingGETParams() *GetEventLogsContentUsingGETParams {
	return &GetEventLogsContentUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventLogsContentUsingGETParamsWithTimeout creates a new GetEventLogsContentUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetEventLogsContentUsingGETParamsWithTimeout(timeout time.Duration) *GetEventLogsContentUsingGETParams {
	return &GetEventLogsContentUsingGETParams{
		timeout: timeout,
	}
}

// NewGetEventLogsContentUsingGETParamsWithContext creates a new GetEventLogsContentUsingGETParams object
// with the ability to set a context for a request.
func NewGetEventLogsContentUsingGETParamsWithContext(ctx context.Context) *GetEventLogsContentUsingGETParams {
	return &GetEventLogsContentUsingGETParams{
		Context: ctx,
	}
}

// NewGetEventLogsContentUsingGETParamsWithHTTPClient creates a new GetEventLogsContentUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEventLogsContentUsingGETParamsWithHTTPClient(client *http.Client) *GetEventLogsContentUsingGETParams {
	return &GetEventLogsContentUsingGETParams{
		HTTPClient: client,
	}
}

/* GetEventLogsContentUsingGETParams contains all the parameters to send to the API endpoint
   for the get event logs content using g e t operation.

   Typically these are written to a http.Request.
*/
type GetEventLogsContentUsingGETParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* EventID.

	   Event ID

	   Format: uuid
	*/
	EventID strfmt.UUID

	/* RequestID.

	   Request ID

	   Format: uuid
	*/
	RequestID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get event logs content using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEventLogsContentUsingGETParams) WithDefaults() *GetEventLogsContentUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get event logs content using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEventLogsContentUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get event logs content using get params
func (o *GetEventLogsContentUsingGETParams) WithTimeout(timeout time.Duration) *GetEventLogsContentUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get event logs content using get params
func (o *GetEventLogsContentUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get event logs content using get params
func (o *GetEventLogsContentUsingGETParams) WithContext(ctx context.Context) *GetEventLogsContentUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get event logs content using get params
func (o *GetEventLogsContentUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get event logs content using get params
func (o *GetEventLogsContentUsingGETParams) WithHTTPClient(client *http.Client) *GetEventLogsContentUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get event logs content using get params
func (o *GetEventLogsContentUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get event logs content using get params
func (o *GetEventLogsContentUsingGETParams) WithAPIVersion(aPIVersion *string) *GetEventLogsContentUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get event logs content using get params
func (o *GetEventLogsContentUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithEventID adds the eventID to the get event logs content using get params
func (o *GetEventLogsContentUsingGETParams) WithEventID(eventID strfmt.UUID) *GetEventLogsContentUsingGETParams {
	o.SetEventID(eventID)
	return o
}

// SetEventID adds the eventId to the get event logs content using get params
func (o *GetEventLogsContentUsingGETParams) SetEventID(eventID strfmt.UUID) {
	o.EventID = eventID
}

// WithRequestID adds the requestID to the get event logs content using get params
func (o *GetEventLogsContentUsingGETParams) WithRequestID(requestID strfmt.UUID) *GetEventLogsContentUsingGETParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the get event logs content using get params
func (o *GetEventLogsContentUsingGETParams) SetRequestID(requestID strfmt.UUID) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventLogsContentUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param eventId
	if err := r.SetPathParam("eventId", o.EventID.String()); err != nil {
		return err
	}

	// path param requestId
	if err := r.SetPathParam("requestId", o.RequestID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
