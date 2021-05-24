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

// NetworkProfileSpecification network profile specification
//
// swagger:model NetworkProfileSpecification
type NetworkProfileSpecification struct {

	// Additional properties that may be used to extend the Network Profile object that is produced from this specification.  For isolationType security group, datastoreId identifies the Compute Resource Edge datastore. computeCluster and resourcePoolId identify the Compute Resource Edge cluster. For isolationType subnet, distributedLogicalRouterStateLink identifies the on-demand network distributed local router.  onDemandNetworkIPAssignmentType identifies the on-demand network IP range assignment type static, dynamic, or mixed.
	// Example: { \"resourcePoolId\" : \"resource-pool-1\", \"datastoreId\" : \"StoragePod:group-p87839\", \"computeCluster\" : \"/resources/compute/1234\", \"distributedLogicalRouterStateLink\" : \"/resources/routers/1234\", \"onDemandNetworkIPAssignmentType\" : \"dynamic\"}
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// List of external IP blocks coming from an external IPAM provider that can be used to create subnetworks inside them
	// Example: [\"3e2bb9bc-6a6a-11ea-bc55-0242ac130003\"]
	ExternalIPBlockIds []string `json:"externalIpBlockIds"`

	// A list of fabric network Ids which are assigned to the network profile.
	// Example: [ \"6543\" ]
	FabricNetworkIds []string `json:"fabricNetworkIds"`

	// The CIDR prefix length to be used for the isolated networks that are created with the network profile.
	// Example: 24
	IsolatedNetworkCIDRPrefix int32 `json:"isolatedNetworkCIDRPrefix,omitempty"`

	// The Id of the fabric network used for outbound access.
	// Example: 1234
	IsolationExternalFabricNetworkID string `json:"isolationExternalFabricNetworkId,omitempty"`

	// CIDR of the isolation network domain.
	// Example: 10.10.10.0/24
	IsolationNetworkDomainCIDR string `json:"isolationNetworkDomainCIDR,omitempty"`

	// The Id of the network domain used for creating isolated networks.
	// Example: 1234
	IsolationNetworkDomainID string `json:"isolationNetworkDomainId,omitempty"`

	// Specifies the isolation type e.g. none, subnet or security group
	// Example: SUBNET
	// Enum: [NONE SUBNET SECURITY_GROUP]
	IsolationType string `json:"isolationType,omitempty"`

	// A list of load balancers which are assigned to the network profile.
	// Example: [ \"6545\" ]
	LoadBalancerIds []string `json:"loadBalancerIds"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// The Id of the region for which this profile is created
	// Example: 9e49
	// Required: true
	RegionID *string `json:"regionId"`

	// A list of security group Ids which are assigned to the network profile.
	// Example: [ \"6545\" ]
	SecurityGroupIds []string `json:"securityGroupIds"`

	// A set of tag keys and optional values that should be set on any resource that is produced from this specification.
	// Example: [ { \"key\" : \"dev\", \"value\": \"hard\" } ]
	Tags []*Tag `json:"tags"`
}

// Validate validates this network profile specification
func (m *NetworkProfileSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsolationType(formats); err != nil {
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

var networkProfileSpecificationTypeIsolationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","SUBNET","SECURITY_GROUP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkProfileSpecificationTypeIsolationTypePropEnum = append(networkProfileSpecificationTypeIsolationTypePropEnum, v)
	}
}

const (

	// NetworkProfileSpecificationIsolationTypeNONE captures enum value "NONE"
	NetworkProfileSpecificationIsolationTypeNONE string = "NONE"

	// NetworkProfileSpecificationIsolationTypeSUBNET captures enum value "SUBNET"
	NetworkProfileSpecificationIsolationTypeSUBNET string = "SUBNET"

	// NetworkProfileSpecificationIsolationTypeSECURITYGROUP captures enum value "SECURITY_GROUP"
	NetworkProfileSpecificationIsolationTypeSECURITYGROUP string = "SECURITY_GROUP"
)

// prop value enum
func (m *NetworkProfileSpecification) validateIsolationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkProfileSpecificationTypeIsolationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkProfileSpecification) validateIsolationType(formats strfmt.Registry) error {
	if swag.IsZero(m.IsolationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateIsolationTypeEnum("isolationType", "body", m.IsolationType); err != nil {
		return err
	}

	return nil
}

func (m *NetworkProfileSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NetworkProfileSpecification) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("regionId", "body", m.RegionID); err != nil {
		return err
	}

	return nil
}

func (m *NetworkProfileSpecification) validateTags(formats strfmt.Registry) error {
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

// ContextValidate validate this network profile specification based on the context it is used
func (m *NetworkProfileSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkProfileSpecification) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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
func (m *NetworkProfileSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkProfileSpecification) UnmarshalBinary(b []byte) error {
	var res NetworkProfileSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
