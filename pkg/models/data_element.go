// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataElement DataElement
//
// swagger:model DataElement
type DataElement struct {

	// description
	Description string `json:"description,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// properties
	Properties interface{} `json:"properties,omitempty"`
}

// Validate validates this data element
func (m *DataElement) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this data element based on context it is used
func (m *DataElement) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataElement) UnmarshalBinary(b []byte) error {
	var res DataElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
