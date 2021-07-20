// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomIntegrations CustomIntegrations
//
// List of Custom Integration instances.
//
// swagger:discriminator CustomIntegrations List of Custom Integration instances.
type CustomIntegrations interface {
	runtime.Validatable
	runtime.ContextValidatable

	// Number of resources within the current page.
	Count() int64
	SetCount(int64)

	// Details of the queried resources.
	// Example: \"documents\": {\n        \"/codestream/api/executions/b80254a7-fcff-4918-ad88-501d08096337\": {\n            \"project\": \"test-project\",\n            \"id\": \"b80254a7-fcff-4918-ad88-501d08096337\",\n            \"name\": \"test\",\n            \"updatedAt\": \"2019-09-23 13:48:54.483\",\n            \"tags\": [],\n            \"_link\": \"/codestream/api/executions/b80254a7-fcff-4918-ad88-501d08096337\",\n            \"_updateTimeInMicros\": 1569226734483000,\n            \"_createTimeInMicros\": 1569226720988000,\n            \"index\": 1,\n            \"notifications\": [],\n            \"comments\": \"\",\n            \"starred\": {},\n            \"input\": {},\n            \"output\": {},\n            \"stageOrder\": [],\n            \"stages\": {},\n            \"status\": \"COMPLETED\",\n            \"statusMessage\": \"Execution Completed.\",\n            \"_durationInMicros\": 13495000,\n            \"_requestTimeInMicros\": 1569226720988000,\n            \"_executedBy\": \"exampleuser\",\n            \"_pipelineLink\": \"/codestream/api/pipelines/b49898f9-d42d-4f19-8bda-e77a373c41b9\",\n            \"_nested\": false,\n            \"_rollback\": false,\n            \"_inputMeta\": {},\n            \"_outputMeta\": {},\n            \"workspaceResults\": []\n        }\n    }
	Documents() map[string]CustomIntegration
	SetDocuments(map[string]CustomIntegration)

	// Partial URLs representing the links to the queried resources.
	// Example: /codestream/api/executions/b80254a7-fcff-4918-ad88-501d08096337
	Links() []string
	SetLinks([]string)

	// Number of resources across all pages.
	TotalCount() int64
	SetTotalCount(int64)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type customIntegrations struct {
	countField int64

	documentsField map[string]CustomIntegration

	linksField []string

	totalCountField int64
}

// Count gets the count of this polymorphic type
func (m *customIntegrations) Count() int64 {
	return m.countField
}

// SetCount sets the count of this polymorphic type
func (m *customIntegrations) SetCount(val int64) {
	m.countField = val
}

// Documents gets the documents of this polymorphic type
func (m *customIntegrations) Documents() map[string]CustomIntegration {
	return m.documentsField
}

// SetDocuments sets the documents of this polymorphic type
func (m *customIntegrations) SetDocuments(val map[string]CustomIntegration) {
	m.documentsField = val
}

// Links gets the links of this polymorphic type
func (m *customIntegrations) Links() []string {
	return m.linksField
}

// SetLinks sets the links of this polymorphic type
func (m *customIntegrations) SetLinks(val []string) {
	m.linksField = val
}

// TotalCount gets the total count of this polymorphic type
func (m *customIntegrations) TotalCount() int64 {
	return m.totalCountField
}

// SetTotalCount sets the total count of this polymorphic type
func (m *customIntegrations) SetTotalCount(val int64) {
	m.totalCountField = val
}

// UnmarshalCustomIntegrationsSlice unmarshals polymorphic slices of CustomIntegrations
func UnmarshalCustomIntegrationsSlice(reader io.Reader, consumer runtime.Consumer) ([]CustomIntegrations, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []CustomIntegrations
	for _, element := range elements {
		obj, err := unmarshalCustomIntegrations(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalCustomIntegrations unmarshals polymorphic CustomIntegrations
func UnmarshalCustomIntegrations(reader io.Reader, consumer runtime.Consumer) (CustomIntegrations, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalCustomIntegrations(data, consumer)
}

func unmarshalCustomIntegrations(data []byte, consumer runtime.Consumer) (CustomIntegrations, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the List of Custom Integration instances. property.
	var getType struct {
		ListOfCustomIntegrationInstances string `json:"List of Custom Integration instances."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("List of Custom Integration instances.", "body", getType.ListOfCustomIntegrationInstances); err != nil {
		return nil, err
	}

	// The value of List of Custom Integration instances. is used to determine which type to create and unmarshal the data into
	switch getType.ListOfCustomIntegrationInstances {
	case "CustomIntegrations":
		var result customIntegrations
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid List of Custom Integration instances. value: %q", getType.ListOfCustomIntegrationInstances)
}

// Validate validates this custom integrations
func (m *customIntegrations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocuments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *customIntegrations) validateDocuments(formats strfmt.Registry) error {
	if swag.IsZero(m.Documents()) { // not required
		return nil
	}

	for k := range m.Documents() {

		if err := validate.Required("documents"+"."+k, "body", m.Documents()[k]); err != nil {
			return err
		}
		if val, ok := m.Documents()[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this custom integrations based on the context it is used
func (m *customIntegrations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDocuments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *customIntegrations) contextValidateDocuments(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Documents() {

		if val, ok := m.Documents()[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}
