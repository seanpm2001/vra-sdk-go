// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GerritTriggerSpecGerritEventConfiguration GerritTriggerSpec$GerritEventConfiguration
//
// swagger:model GerritTriggerSpecGerritEventConfiguration
type GerritTriggerSpecGerritEventConfiguration struct {

	// Type of the gerrit event.
	// Example: patchset-created
	// Required: true
	EventType *string `json:"eventType"`

	// Comment to be posted to the ChangeSet on execution termination.
	// Example: Pipeline Execution Failed
	FailureComment string `json:"failureComment,omitempty"`

	// Map representing the Input properties for the Pipeline.
	// Example: [{"ip":"10.5.23.84","script":"testScript.sh"}]
	Input map[string]string `json:"input,omitempty"`

	// Pipeline that needs to be triggered on receiving this event.
	// Example: DemoPipeline
	// Required: true
	Pipeline *string `json:"pipeline"`

	// Comment to be posted to the ChangeSet on execution termination.
	// Example: Pipeline Execution Completed
	SuccessComment string `json:"successComment,omitempty"`

	// The label to be posted on Gerrit Server to perform actions.
	// Example: Verified +1
	VerifiedLabel string `json:"verifiedLabel,omitempty"`
}

// Validate validates this gerrit trigger spec gerrit event configuration
func (m *GerritTriggerSpecGerritEventConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePipeline(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GerritTriggerSpecGerritEventConfiguration) validateEventType(formats strfmt.Registry) error {

	if err := validate.Required("eventType", "body", m.EventType); err != nil {
		return err
	}

	return nil
}

func (m *GerritTriggerSpecGerritEventConfiguration) validatePipeline(formats strfmt.Registry) error {

	if err := validate.Required("pipeline", "body", m.Pipeline); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this gerrit trigger spec gerrit event configuration based on context it is used
func (m *GerritTriggerSpecGerritEventConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GerritTriggerSpecGerritEventConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GerritTriggerSpecGerritEventConfiguration) UnmarshalBinary(b []byte) error {
	var res GerritTriggerSpecGerritEventConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
