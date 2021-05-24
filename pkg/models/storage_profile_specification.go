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

// StorageProfileSpecification Represents a specification of generic storage profile.
//
// swagger:model StorageProfileSpecification
type StorageProfileSpecification struct {

	// Indicates if a storage profile is a default profile.
	// Required: true
	DefaultItem *bool `json:"defaultItem"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Map of storage properties that are to be applied on disk while provisioning.
	// Example: { \"diskProperties\": {\n                    \"provisioningType\": \"thin\",\n                    \"sharesLevel\": \"low\",\n                    \"shares\": \"500\",\n                    \"limitIops\": \"500\"\n                    \"diskType\": \"firstClass\"\n                    \"deviceType\": \"ebs\"\n                    \"volumeType\": \"gp2\"\n                    \"azureDataDiskCaching\": \"ReadWrite\"\n                    \"azureOsDiskCaching\": \"ReadWrite\"\n                    \"azureManagedDiskType\": \"Standard_LRS\"\n                } }
	DiskProperties map[string]string `json:"diskProperties,omitempty"`

	// Map of storage placements to know where the disk is provisioned.
	// Example: { \"diskTargetProperties\": {\n                    \"storageAccountId\": \"27dhbf7\",\n                    \"storagePolicyId\": \"7fhfj9f\",\n                    \"datastoreId\": \"638nfjd8\",\n                } }
	DiskTargetProperties map[string]string `json:"diskTargetProperties,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// The Id of the region that is associated with the storage profile.
	// Example: 31186
	// Required: true
	RegionID *string `json:"regionId"`

	// Indicates whether this storage profile supports encryption or not.
	SupportsEncryption bool `json:"supportsEncryption,omitempty"`

	// A list of tags that represent the capabilities of this storage profile
	// Example: [ { \"key\" : \"tier\", \"value\": \"silver\" } ]
	Tags []*Tag `json:"tags"`
}

// Validate validates this storage profile specification
func (m *StorageProfileSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionID(formats); err != nil {
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

func (m *StorageProfileSpecification) validateDefaultItem(formats strfmt.Registry) error {

	if err := validate.Required("defaultItem", "body", m.DefaultItem); err != nil {
		return err
	}

	return nil
}

func (m *StorageProfileSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StorageProfileSpecification) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("regionId", "body", m.RegionID); err != nil {
		return err
	}

	return nil
}

func (m *StorageProfileSpecification) validateTags(formats strfmt.Registry) error {
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

// ContextValidate validate this storage profile specification based on the context it is used
func (m *StorageProfileSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageProfileSpecification) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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
func (m *StorageProfileSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageProfileSpecification) UnmarshalBinary(b []byte) error {
	var res StorageProfileSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
