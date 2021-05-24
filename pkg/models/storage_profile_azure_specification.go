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

// StorageProfileAzureSpecification Represents a specification of Azure storage profile.
//
// swagger:model StorageProfileAzureSpecification
type StorageProfileAzureSpecification struct {

	// Indicates the caching mechanism for additional disk.
	// Example: None / ReadOnly / ReadWrite
	DataDiskCaching string `json:"dataDiskCaching,omitempty"`

	// Indicates if a storage policy contains default storage properties.
	// Example: true
	DefaultItem bool `json:"defaultItem,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Indicates the id of disk encryption set.
	// Example: /subscriptions/b8ef63/resourceGroups/DiskEncryptionSets/providers/Microsoft.Compute/diskEncryptionSets/MyDES
	DiskEncryptionSetID string `json:"diskEncryptionSetId,omitempty"`

	// Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.
	// Example: Standard_LRS / Premium_LRS
	DiskType string `json:"diskType,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.
	// Example: None / ReadOnly / ReadWrite
	OsDiskCaching string `json:"osDiskCaching,omitempty"`

	// The If of the region that is associated with the storage profile.
	// Example: 31186
	// Required: true
	RegionID *string `json:"regionId"`

	// Id of a storage account where in the disk is placed.
	// Example: aaa82
	StorageAccountID string `json:"storageAccountId,omitempty"`

	// Indicates whether this storage policy should support encryption or not.
	// Example: false
	SupportsEncryption bool `json:"supportsEncryption,omitempty"`

	// A set of tag keys and optional values for a storage policy which define set of specifications for creating a disk.
	// Example: [ { \"key\" : \"tier\", \"value\": \"silver\" } ]
	Tags []*Tag `json:"tags"`
}

// Validate validates this storage profile azure specification
func (m *StorageProfileAzureSpecification) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *StorageProfileAzureSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StorageProfileAzureSpecification) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("regionId", "body", m.RegionID); err != nil {
		return err
	}

	return nil
}

func (m *StorageProfileAzureSpecification) validateTags(formats strfmt.Registry) error {
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

// ContextValidate validate this storage profile azure specification based on the context it is used
func (m *StorageProfileAzureSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageProfileAzureSpecification) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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
func (m *StorageProfileAzureSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageProfileAzureSpecification) UnmarshalBinary(b []byte) error {
	var res StorageProfileAzureSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
