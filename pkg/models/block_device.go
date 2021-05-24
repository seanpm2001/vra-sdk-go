// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BlockDevice State object representing a block device.<br>**HATEOAS** links:<br>**cloud-accounts** - array[CloudAccount] - Cloud accounts where this disk is provisioned.<br>**self** - BlockDevice - Self link to this block device
//
// swagger:model BlockDevice
type BlockDevice struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Capacity of the block device in GB.
	// Example: 10
	// Required: true
	CapacityInGB *int32 `json:"capacityInGB"`

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

	// Indicates whether the block device survives a delete action.
	// Example: true
	Persistent bool `json:"persistent,omitempty"`

	// The id of the project this resource belongs to.
	// Example: 9e49
	ProjectID string `json:"projectId,omitempty"`

	// Status of the block device
	// Example: ATTACHED
	// Required: true
	// Enum: [DETACHED ATTACHED AVAILABLE]
	Status *string `json:"status"`

	// A set of tag keys and optional values that were set on this resource.
	// Example: [ { \"key\" : \"ownedBy\", \"value\": \"Rainpole\" } ]
	Tags []*Tag `json:"tags"`

	// Type of the block device
	// Example: HDD
	// Enum: [SSD HDD CDROM FLOPPY]
	Type string `json:"type,omitempty"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	// Example: 2012-09-27
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this block device
func (m *BlockDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacityInGB(formats); err != nil {
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

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockDevice) validateLinks(formats strfmt.Registry) error {

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

func (m *BlockDevice) validateCapacityInGB(formats strfmt.Registry) error {

	if err := validate.Required("capacityInGB", "body", m.CapacityInGB); err != nil {
		return err
	}

	return nil
}

func (m *BlockDevice) validateCloudAccountIds(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudAccountIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("cloudAccountIds", "body", m.CloudAccountIds); err != nil {
		return err
	}

	return nil
}

func (m *BlockDevice) validateExternalRegionID(formats strfmt.Registry) error {

	if err := validate.Required("externalRegionId", "body", m.ExternalRegionID); err != nil {
		return err
	}

	return nil
}

func (m *BlockDevice) validateExternalZoneID(formats strfmt.Registry) error {

	if err := validate.Required("externalZoneId", "body", m.ExternalZoneID); err != nil {
		return err
	}

	return nil
}

func (m *BlockDevice) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var blockDeviceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DETACHED","ATTACHED","AVAILABLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		blockDeviceTypeStatusPropEnum = append(blockDeviceTypeStatusPropEnum, v)
	}
}

const (

	// BlockDeviceStatusDETACHED captures enum value "DETACHED"
	BlockDeviceStatusDETACHED string = "DETACHED"

	// BlockDeviceStatusATTACHED captures enum value "ATTACHED"
	BlockDeviceStatusATTACHED string = "ATTACHED"

	// BlockDeviceStatusAVAILABLE captures enum value "AVAILABLE"
	BlockDeviceStatusAVAILABLE string = "AVAILABLE"
)

// prop value enum
func (m *BlockDevice) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, blockDeviceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BlockDevice) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *BlockDevice) validateTags(formats strfmt.Registry) error {
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

var blockDeviceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SSD","HDD","CDROM","FLOPPY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		blockDeviceTypeTypePropEnum = append(blockDeviceTypeTypePropEnum, v)
	}
}

const (

	// BlockDeviceTypeSSD captures enum value "SSD"
	BlockDeviceTypeSSD string = "SSD"

	// BlockDeviceTypeHDD captures enum value "HDD"
	BlockDeviceTypeHDD string = "HDD"

	// BlockDeviceTypeCDROM captures enum value "CDROM"
	BlockDeviceTypeCDROM string = "CDROM"

	// BlockDeviceTypeFLOPPY captures enum value "FLOPPY"
	BlockDeviceTypeFLOPPY string = "FLOPPY"
)

// prop value enum
func (m *BlockDevice) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, blockDeviceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BlockDevice) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this block device based on the context it is used
func (m *BlockDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *BlockDevice) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BlockDevice) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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
func (m *BlockDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockDevice) UnmarshalBinary(b []byte) error {
	var res BlockDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
