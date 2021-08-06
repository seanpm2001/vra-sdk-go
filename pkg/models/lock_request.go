// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LockRequest LockRequest
//
// swagger:model LockRequest
type LockRequest struct {

	// service name
	ServiceName string `json:"serviceName,omitempty"`
}

// Validate validates this lock request
func (m *LockRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this lock request based on context it is used
func (m *LockRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LockRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LockRequest) UnmarshalBinary(b []byte) error {
	var res LockRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
