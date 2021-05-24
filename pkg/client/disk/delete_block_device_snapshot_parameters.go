// Code generated by go-swagger; DO NOT EDIT.

package disk

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

// NewDeleteBlockDeviceSnapshotParams creates a new DeleteBlockDeviceSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBlockDeviceSnapshotParams() *DeleteBlockDeviceSnapshotParams {
	return &DeleteBlockDeviceSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBlockDeviceSnapshotParamsWithTimeout creates a new DeleteBlockDeviceSnapshotParams object
// with the ability to set a timeout on a request.
func NewDeleteBlockDeviceSnapshotParamsWithTimeout(timeout time.Duration) *DeleteBlockDeviceSnapshotParams {
	return &DeleteBlockDeviceSnapshotParams{
		timeout: timeout,
	}
}

// NewDeleteBlockDeviceSnapshotParamsWithContext creates a new DeleteBlockDeviceSnapshotParams object
// with the ability to set a context for a request.
func NewDeleteBlockDeviceSnapshotParamsWithContext(ctx context.Context) *DeleteBlockDeviceSnapshotParams {
	return &DeleteBlockDeviceSnapshotParams{
		Context: ctx,
	}
}

// NewDeleteBlockDeviceSnapshotParamsWithHTTPClient creates a new DeleteBlockDeviceSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBlockDeviceSnapshotParamsWithHTTPClient(client *http.Client) *DeleteBlockDeviceSnapshotParams {
	return &DeleteBlockDeviceSnapshotParams{
		HTTPClient: client,
	}
}

/* DeleteBlockDeviceSnapshotParams contains all the parameters to send to the API endpoint
   for the delete block device snapshot operation.

   Typically these are written to a http.Request.
*/
type DeleteBlockDeviceSnapshotParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of the block device.
	*/
	ID string

	/* Id1.

	   Snapshot id to delete.
	*/
	Id1 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete block device snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBlockDeviceSnapshotParams) WithDefaults() *DeleteBlockDeviceSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete block device snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBlockDeviceSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete block device snapshot params
func (o *DeleteBlockDeviceSnapshotParams) WithTimeout(timeout time.Duration) *DeleteBlockDeviceSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete block device snapshot params
func (o *DeleteBlockDeviceSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete block device snapshot params
func (o *DeleteBlockDeviceSnapshotParams) WithContext(ctx context.Context) *DeleteBlockDeviceSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete block device snapshot params
func (o *DeleteBlockDeviceSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete block device snapshot params
func (o *DeleteBlockDeviceSnapshotParams) WithHTTPClient(client *http.Client) *DeleteBlockDeviceSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete block device snapshot params
func (o *DeleteBlockDeviceSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete block device snapshot params
func (o *DeleteBlockDeviceSnapshotParams) WithAPIVersion(aPIVersion *string) *DeleteBlockDeviceSnapshotParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete block device snapshot params
func (o *DeleteBlockDeviceSnapshotParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete block device snapshot params
func (o *DeleteBlockDeviceSnapshotParams) WithID(id string) *DeleteBlockDeviceSnapshotParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete block device snapshot params
func (o *DeleteBlockDeviceSnapshotParams) SetID(id string) {
	o.ID = id
}

// WithId1 adds the id1 to the delete block device snapshot params
func (o *DeleteBlockDeviceSnapshotParams) WithId1(id1 string) *DeleteBlockDeviceSnapshotParams {
	o.SetId1(id1)
	return o
}

// SetId1 adds the id1 to the delete block device snapshot params
func (o *DeleteBlockDeviceSnapshotParams) SetId1(id1 string) {
	o.Id1 = id1
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBlockDeviceSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param id1
	if err := r.SetPathParam("id1", o.Id1); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
