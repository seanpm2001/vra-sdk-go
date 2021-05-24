// Code generated by go-swagger; DO NOT EDIT.

package blueprint_terraform_integrations

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

// NewGetTerraformConfigurationSourcesUsingGET1Params creates a new GetTerraformConfigurationSourcesUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTerraformConfigurationSourcesUsingGET1Params() *GetTerraformConfigurationSourcesUsingGET1Params {
	return &GetTerraformConfigurationSourcesUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTerraformConfigurationSourcesUsingGET1ParamsWithTimeout creates a new GetTerraformConfigurationSourcesUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetTerraformConfigurationSourcesUsingGET1ParamsWithTimeout(timeout time.Duration) *GetTerraformConfigurationSourcesUsingGET1Params {
	return &GetTerraformConfigurationSourcesUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetTerraformConfigurationSourcesUsingGET1ParamsWithContext creates a new GetTerraformConfigurationSourcesUsingGET1Params object
// with the ability to set a context for a request.
func NewGetTerraformConfigurationSourcesUsingGET1ParamsWithContext(ctx context.Context) *GetTerraformConfigurationSourcesUsingGET1Params {
	return &GetTerraformConfigurationSourcesUsingGET1Params{
		Context: ctx,
	}
}

// NewGetTerraformConfigurationSourcesUsingGET1ParamsWithHTTPClient creates a new GetTerraformConfigurationSourcesUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetTerraformConfigurationSourcesUsingGET1ParamsWithHTTPClient(client *http.Client) *GetTerraformConfigurationSourcesUsingGET1Params {
	return &GetTerraformConfigurationSourcesUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetTerraformConfigurationSourcesUsingGET1Params contains all the parameters to send to the API endpoint
   for the get terraform configuration sources using get1 operation.

   Typically these are written to a http.Request.
*/
type GetTerraformConfigurationSourcesUsingGET1Params struct {

	/* Projects.

	   A comma-separated list of project IDs. Results will be filtered so only configuration sources accessible from one or more of these projects will be returned.
	*/
	Projects []string

	/* Search.

	   A search parameter to search based on configuration source name or repository.
	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get terraform configuration sources using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformConfigurationSourcesUsingGET1Params) WithDefaults() *GetTerraformConfigurationSourcesUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get terraform configuration sources using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformConfigurationSourcesUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get terraform configuration sources using get1 params
func (o *GetTerraformConfigurationSourcesUsingGET1Params) WithTimeout(timeout time.Duration) *GetTerraformConfigurationSourcesUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get terraform configuration sources using get1 params
func (o *GetTerraformConfigurationSourcesUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get terraform configuration sources using get1 params
func (o *GetTerraformConfigurationSourcesUsingGET1Params) WithContext(ctx context.Context) *GetTerraformConfigurationSourcesUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get terraform configuration sources using get1 params
func (o *GetTerraformConfigurationSourcesUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get terraform configuration sources using get1 params
func (o *GetTerraformConfigurationSourcesUsingGET1Params) WithHTTPClient(client *http.Client) *GetTerraformConfigurationSourcesUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get terraform configuration sources using get1 params
func (o *GetTerraformConfigurationSourcesUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjects adds the projects to the get terraform configuration sources using get1 params
func (o *GetTerraformConfigurationSourcesUsingGET1Params) WithProjects(projects []string) *GetTerraformConfigurationSourcesUsingGET1Params {
	o.SetProjects(projects)
	return o
}

// SetProjects adds the projects to the get terraform configuration sources using get1 params
func (o *GetTerraformConfigurationSourcesUsingGET1Params) SetProjects(projects []string) {
	o.Projects = projects
}

// WithSearch adds the search to the get terraform configuration sources using get1 params
func (o *GetTerraformConfigurationSourcesUsingGET1Params) WithSearch(search *string) *GetTerraformConfigurationSourcesUsingGET1Params {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get terraform configuration sources using get1 params
func (o *GetTerraformConfigurationSourcesUsingGET1Params) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetTerraformConfigurationSourcesUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Projects != nil {

		// binding items for projects
		joinedProjects := o.bindParamProjects(reg)

		// query array param projects
		if err := r.SetQueryParam("projects", joinedProjects...); err != nil {
			return err
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

// bindParamGetTerraformConfigurationSourcesUsingGET1 binds the parameter projects
func (o *GetTerraformConfigurationSourcesUsingGET1Params) bindParamProjects(formats strfmt.Registry) []string {
	projectsIR := o.Projects

	var projectsIC []string
	for _, projectsIIR := range projectsIR { // explode []string

		projectsIIV := projectsIIR // string as string
		projectsIC = append(projectsIC, projectsIIV)
	}

	// items.CollectionFormat: "multi"
	projectsIS := swag.JoinByFormat(projectsIC, "multi")

	return projectsIS
}
