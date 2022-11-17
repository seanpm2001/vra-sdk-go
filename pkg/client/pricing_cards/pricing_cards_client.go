// Code generated by go-swagger; DO NOT EDIT.

package pricing_cards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pricing cards API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pricing cards API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePolicyUsingPOST2(params *CreatePolicyUsingPOST2Params, opts ...ClientOption) (*CreatePolicyUsingPOST2OK, *CreatePolicyUsingPOST2Created, error)

	DeletePolicyUsingDELETE4(params *DeletePolicyUsingDELETE4Params, opts ...ClientOption) (*DeletePolicyUsingDELETE4OK, *DeletePolicyUsingDELETE4NoContent, error)

	GetPoliciesUsingGET4(params *GetPoliciesUsingGET4Params, opts ...ClientOption) (*GetPoliciesUsingGET4OK, error)

	GetPolicyUsingGET4(params *GetPolicyUsingGET4Params, opts ...ClientOption) (*GetPolicyUsingGET4OK, error)

	UpdatePolicyUsingPUT2(params *UpdatePolicyUsingPUT2Params, opts ...ClientOption) (*UpdatePolicyUsingPUT2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreatePolicyUsingPOST2 creates a new pricing card

Create a new pricing card based on request body and validate its field according to business rules.
*/
func (a *Client) CreatePolicyUsingPOST2(params *CreatePolicyUsingPOST2Params, opts ...ClientOption) (*CreatePolicyUsingPOST2OK, *CreatePolicyUsingPOST2Created, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePolicyUsingPOST2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPolicyUsingPOST_2",
		Method:             "POST",
		PathPattern:        "/price/api/private/pricing-cards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePolicyUsingPOST2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreatePolicyUsingPOST2OK:
		return value, nil, nil
	case *CreatePolicyUsingPOST2Created:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pricing_cards: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeletePolicyUsingDELETE4 deletes the pricing card with specified Id

Deletes the pricing card with the specified id
*/
func (a *Client) DeletePolicyUsingDELETE4(params *DeletePolicyUsingDELETE4Params, opts ...ClientOption) (*DeletePolicyUsingDELETE4OK, *DeletePolicyUsingDELETE4NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePolicyUsingDELETE4Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePolicyUsingDELETE_4",
		Method:             "DELETE",
		PathPattern:        "/price/api/private/pricing-cards/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePolicyUsingDELETE4Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeletePolicyUsingDELETE4OK:
		return value, nil, nil
	case *DeletePolicyUsingDELETE4NoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pricing_cards: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPoliciesUsingGET4 fetches all pricing cards for private policy cloud

Returns a paginated list of pricing cards
*/
func (a *Client) GetPoliciesUsingGET4(params *GetPoliciesUsingGET4Params, opts ...ClientOption) (*GetPoliciesUsingGET4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPoliciesUsingGET4Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPoliciesUsingGET_4",
		Method:             "GET",
		PathPattern:        "/price/api/private/pricing-cards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPoliciesUsingGET4Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPoliciesUsingGET4OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPoliciesUsingGET_4: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetPolicyUsingGET4 finds the pricing card with specified Id

Returns the pricing card with the specified id
*/
func (a *Client) GetPolicyUsingGET4(params *GetPolicyUsingGET4Params, opts ...ClientOption) (*GetPolicyUsingGET4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPolicyUsingGET4Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPolicyUsingGET_4",
		Method:             "GET",
		PathPattern:        "/price/api/private/pricing-cards/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPolicyUsingGET4Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPolicyUsingGET4OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPolicyUsingGET_4: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdatePolicyUsingPUT2 updates the pricing card

Updates the pricing card with the specified Id
*/
func (a *Client) UpdatePolicyUsingPUT2(params *UpdatePolicyUsingPUT2Params, opts ...ClientOption) (*UpdatePolicyUsingPUT2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePolicyUsingPUT2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePolicyUsingPUT_2",
		Method:             "PUT",
		PathPattern:        "/price/api/private/pricing-cards/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePolicyUsingPUT2Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdatePolicyUsingPUT2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePolicyUsingPUT_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
