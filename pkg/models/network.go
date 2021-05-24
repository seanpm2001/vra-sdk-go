// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Network The network object is an opaque reference to a logical network that network interfaces are attached to.<br>
// Based on settings specified by your cloud administrator, it may be a reference to an existing network, or be backed by an on-demand network created for isolation, or a security group that will be attached to machines as part of provisioning.<br>
// Networks are a limited resource, when it is not needed it should be deleted.<br>**HATEOAS** links:<br>**self** - Network - Self link to this network
//
// swagger:model Network
type Network struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// IPv4 address range of the network in CIDR format
	// Example: 10.10.0.0/16
	// Required: true
	Cidr *string `json:"cidr"`

	// Set of ids of the cloud accounts this resource belongs to.
	// Example: [9e49]
	// Unique: true
	CloudAccountIds []string `json:"cloudAccountIds"`

	// Date when the entity was created. The date is in ISO 8601 and UTC.
	// Example: 2012-09-27
	CreatedAt string `json:"createdAt,omitempty"`

	// Additional properties that may be used to extend the base resource.
	// Example: { \"property\" : \"value\" }
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// Deployment id that is associated with this resource.
	// Example: 123e4567-e89b-12d3-a456-426655440000
	DeploymentID string `json:"deploymentId,omitempty"`

	// A human-friendly description.
	// Example: my-description
	Description string `json:"description,omitempty"`

	// External entity Id on the provider side.
	// Example: i-cfe4-e241-e53b-756a9a2e25d2
	ExternalID string `json:"externalId,omitempty"`

	// The external regionId of the resource.
	// Example: us-east-1
	// Required: true
	ExternalRegionID *string `json:"externalRegionId"`

	// The external zoneId of the resource.
	// Example: us-east-1a
	// Required: true
	ExternalZoneID *string `json:"externalZoneId"`

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

	// The id of the project this resource belongs to.
	// Example: 9e49
	ProjectID string `json:"projectId,omitempty"`

	// A set of tag keys and optional values that were set on this resource.
	// Example: [ { \"key\" : \"ownedBy\", \"value\": \"Rainpole\" } ]
	Tags []*Tag `json:"tags"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this network
func (m *Network) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudAccountIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalRegionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalZoneID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Network) validateLinks(formats strfmt.Registry) error {

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

func (m *Network) validateCidr(formats strfmt.Registry) error {

	if err := validate.Required("cidr", "body", m.Cidr); err != nil {
		return err
	}

	return nil
}

func (m *Network) validateCloudAccountIds(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudAccountIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("cloudAccountIds", "body", m.CloudAccountIds); err != nil {
		return err
	}

	return nil
}

func (m *Network) validateExternalRegionID(formats strfmt.Registry) error {

	if err := validate.Required("externalRegionId", "body", m.ExternalRegionID); err != nil {
		return err
	}

	return nil
}

func (m *Network) validateExternalZoneID(formats strfmt.Registry) error {

	if err := validate.Required("externalZoneId", "body", m.ExternalZoneID); err != nil {
		return err
	}

	return nil
}

func (m *Network) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Network) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this network based on the context it is used
func (m *Network) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Network) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Network) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Network) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Network) UnmarshalBinary(b []byte) error {
	var res Network
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
