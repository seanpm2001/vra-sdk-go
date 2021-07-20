// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EndpointCertificate EndpointCertificate
//
// Represents the complete SSL Certificate information of a FQDN.
//
// swagger:discriminator EndpointCertificate Represents the complete SSL Certificate information of a FQDN.
type EndpointCertificate interface {
	runtime.Validatable
	runtime.ContextValidatable

	Fingerprints() CertificateFingerprint
	SetFingerprints(CertificateFingerprint)

	// Represents the entity issuing the Certificate to the holding body.
	IssuedBy() *CertificateIssuer
	SetIssuedBy(*CertificateIssuer)

	IssuedTo() CertificateIssuedTo
	SetIssuedTo(CertificateIssuedTo)

	PeriodOfValidity() CertificateValidity
	SetPeriodOfValidity(CertificateValidity)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type endpointCertificate struct {
	fingerprintsField CertificateFingerprint

	issuedByField *CertificateIssuer

	issuedToField CertificateIssuedTo

	periodOfValidityField CertificateValidity
}

// Fingerprints gets the fingerprints of this polymorphic type
func (m *endpointCertificate) Fingerprints() CertificateFingerprint {
	return m.fingerprintsField
}

// SetFingerprints sets the fingerprints of this polymorphic type
func (m *endpointCertificate) SetFingerprints(val CertificateFingerprint) {
	m.fingerprintsField = val
}

// IssuedBy gets the issued by of this polymorphic type
func (m *endpointCertificate) IssuedBy() *CertificateIssuer {
	return m.issuedByField
}

// SetIssuedBy sets the issued by of this polymorphic type
func (m *endpointCertificate) SetIssuedBy(val *CertificateIssuer) {
	m.issuedByField = val
}

// IssuedTo gets the issued to of this polymorphic type
func (m *endpointCertificate) IssuedTo() CertificateIssuedTo {
	return m.issuedToField
}

// SetIssuedTo sets the issued to of this polymorphic type
func (m *endpointCertificate) SetIssuedTo(val CertificateIssuedTo) {
	m.issuedToField = val
}

// PeriodOfValidity gets the period of validity of this polymorphic type
func (m *endpointCertificate) PeriodOfValidity() CertificateValidity {
	return m.periodOfValidityField
}

// SetPeriodOfValidity sets the period of validity of this polymorphic type
func (m *endpointCertificate) SetPeriodOfValidity(val CertificateValidity) {
	m.periodOfValidityField = val
}

// UnmarshalEndpointCertificateSlice unmarshals polymorphic slices of EndpointCertificate
func UnmarshalEndpointCertificateSlice(reader io.Reader, consumer runtime.Consumer) ([]EndpointCertificate, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []EndpointCertificate
	for _, element := range elements {
		obj, err := unmarshalEndpointCertificate(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalEndpointCertificate unmarshals polymorphic EndpointCertificate
func UnmarshalEndpointCertificate(reader io.Reader, consumer runtime.Consumer) (EndpointCertificate, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalEndpointCertificate(data, consumer)
}

func unmarshalEndpointCertificate(data []byte, consumer runtime.Consumer) (EndpointCertificate, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Represents the complete SSL Certificate information of a FQDN. property.
	var getType struct {
		RepresentsTheCompleteSSLCertificateInformationOfaFQDN string `json:"Represents the complete SSL Certificate information of a FQDN."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Represents the complete SSL Certificate information of a FQDN.", "body", getType.RepresentsTheCompleteSSLCertificateInformationOfaFQDN); err != nil {
		return nil, err
	}

	// The value of Represents the complete SSL Certificate information of a FQDN. is used to determine which type to create and unmarshal the data into
	switch getType.RepresentsTheCompleteSSLCertificateInformationOfaFQDN {
	case "EndpointCertificate":
		var result endpointCertificate
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Represents the complete SSL Certificate information of a FQDN. value: %q", getType.RepresentsTheCompleteSSLCertificateInformationOfaFQDN)
}

// Validate validates this endpoint certificate
func (m *endpointCertificate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFingerprints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuedTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriodOfValidity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *endpointCertificate) validateFingerprints(formats strfmt.Registry) error {
	if swag.IsZero(m.Fingerprints()) { // not required
		return nil
	}

	if err := m.Fingerprints().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fingerprints")
		}
		return err
	}

	return nil
}

func (m *endpointCertificate) validateIssuedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.IssuedBy()) { // not required
		return nil
	}

	if m.IssuedBy() != nil {
		if err := m.IssuedBy().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issuedBy")
			}
			return err
		}
	}

	return nil
}

func (m *endpointCertificate) validateIssuedTo(formats strfmt.Registry) error {
	if swag.IsZero(m.IssuedTo()) { // not required
		return nil
	}

	if err := m.IssuedTo().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("issuedTo")
		}
		return err
	}

	return nil
}

func (m *endpointCertificate) validatePeriodOfValidity(formats strfmt.Registry) error {
	if swag.IsZero(m.PeriodOfValidity()) { // not required
		return nil
	}

	if err := m.PeriodOfValidity().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("periodOfValidity")
		}
		return err
	}

	return nil
}

// ContextValidate validate this endpoint certificate based on the context it is used
func (m *endpointCertificate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFingerprints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIssuedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIssuedTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeriodOfValidity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *endpointCertificate) contextValidateFingerprints(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Fingerprints().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fingerprints")
		}
		return err
	}

	return nil
}

func (m *endpointCertificate) contextValidateIssuedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.IssuedBy() != nil {
		if err := m.IssuedBy().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("issuedBy")
			}
			return err
		}
	}

	return nil
}

func (m *endpointCertificate) contextValidateIssuedTo(ctx context.Context, formats strfmt.Registry) error {

	if err := m.IssuedTo().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("issuedTo")
		}
		return err
	}

	return nil
}

func (m *endpointCertificate) contextValidatePeriodOfValidity(ctx context.Context, formats strfmt.Registry) error {

	if err := m.PeriodOfValidity().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("periodOfValidity")
		}
		return err
	}

	return nil
}
