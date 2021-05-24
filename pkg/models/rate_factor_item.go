// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RateFactorItem RateFactorItem
//
// swagger:model RateFactorItem
type RateFactorItem struct {

	// key
	Key string `json:"key,omitempty"`

	// rate factor
	RateFactor *RateFactor `json:"rateFactor,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this rate factor item
func (m *RateFactorItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRateFactor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RateFactorItem) validateRateFactor(formats strfmt.Registry) error {
	if swag.IsZero(m.RateFactor) { // not required
		return nil
	}

	if m.RateFactor != nil {
		if err := m.RateFactor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rateFactor")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rate factor item based on the context it is used
func (m *RateFactorItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRateFactor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RateFactorItem) contextValidateRateFactor(ctx context.Context, formats strfmt.Registry) error {

	if m.RateFactor != nil {
		if err := m.RateFactor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rateFactor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RateFactorItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RateFactorItem) UnmarshalBinary(b []byte) error {
	var res RateFactorItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
