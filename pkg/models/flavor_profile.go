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

// FlavorProfile Represents a structure that holds flavor mappings defined for the corresponding cloud end-point region.<br>**HATEOAS** links:<br>**region** - Region - Region for the profile.<br>**self** - FlavorProfile - Self link to this flavor profile
//
// swagger:model FlavorProfile
type FlavorProfile struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Id of the cloud account this flavor profile belongs to.
	// Example: [9e49]
	CloudAccountID string `json:"cloudAccountId,omitempty"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// The id of the region for which this profile is defined
	// Example: us-east-1
	ExternalRegionID string `json:"externalRegionId,omitempty"`

	// A list of the flavor mappings defined for the corresponding cloud end-point region
	// Required: true
	FlavorMappings *FlavorMapping `json:"flavorMappings"`

	// The id of this resource instance
	// Example: 9e49
	// Required: true
	ID *string `json:"id"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Example: my-name
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	// Example: 9e49
	OrgID string `json:"orgId,omitempty"`

	// This field is deprecated. Use orgId instead. The id of the organization this entity belongs to.
	// Example: deprecated
	OrganizationID string `json:"organizationId,omitempty"`

	// Email of the user that owns the entity.
	// Example: csp@vmware.com
	Owner string `json:"owner,omitempty"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this flavor profile
func (m *FlavorProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlavorMappings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlavorProfile) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	for k := range m.Links {

		if err := validate.Required("_links"+"."+k, "body", m.Links[k]); err != nil {
			return err
		}
		if val, ok := m.Links[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *FlavorProfile) validateFlavorMappings(formats strfmt.Registry) error {

	if err := validate.Required("flavorMappings", "body", m.FlavorMappings); err != nil {
		return err
	}

	if m.FlavorMappings != nil {
		if err := m.FlavorMappings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flavorMappings")
			}
			return err
		}
	}

	return nil
}

func (m *FlavorProfile) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this flavor profile based on the context it is used
func (m *FlavorProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFlavorMappings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlavorProfile) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	for k := range m.Links {

		if val, ok := m.Links[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *FlavorProfile) contextValidateFlavorMappings(ctx context.Context, formats strfmt.Registry) error {

	if m.FlavorMappings != nil {
		if err := m.FlavorMappings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flavorMappings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlavorProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlavorProfile) UnmarshalBinary(b []byte) error {
	var res FlavorProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
