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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NotificationConfiguration NotificationConfiguration
//
// Represents collection of different Event configurations.
//
// swagger:discriminator NotificationConfiguration Represents collection of different Event configurations.
type NotificationConfiguration interface {
	runtime.Validatable
	runtime.ContextValidatable

	Email() []EmailEventConfig
	SetEmail([]EmailEventConfig)

	Jira() []JiraEventConfig
	SetJira([]JiraEventConfig)

	Webhook() []WebhookEventConfig
	SetWebhook([]WebhookEventConfig)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type notificationConfiguration struct {
	emailField []EmailEventConfig

	jiraField []JiraEventConfig

	webhookField []WebhookEventConfig
}

// Email gets the email of this polymorphic type
func (m *notificationConfiguration) Email() []EmailEventConfig {
	return m.emailField
}

// SetEmail sets the email of this polymorphic type
func (m *notificationConfiguration) SetEmail(val []EmailEventConfig) {
	m.emailField = val
}

// Jira gets the jira of this polymorphic type
func (m *notificationConfiguration) Jira() []JiraEventConfig {
	return m.jiraField
}

// SetJira sets the jira of this polymorphic type
func (m *notificationConfiguration) SetJira(val []JiraEventConfig) {
	m.jiraField = val
}

// Webhook gets the webhook of this polymorphic type
func (m *notificationConfiguration) Webhook() []WebhookEventConfig {
	return m.webhookField
}

// SetWebhook sets the webhook of this polymorphic type
func (m *notificationConfiguration) SetWebhook(val []WebhookEventConfig) {
	m.webhookField = val
}

// UnmarshalNotificationConfigurationSlice unmarshals polymorphic slices of NotificationConfiguration
func UnmarshalNotificationConfigurationSlice(reader io.Reader, consumer runtime.Consumer) ([]NotificationConfiguration, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []NotificationConfiguration
	for _, element := range elements {
		obj, err := unmarshalNotificationConfiguration(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalNotificationConfiguration unmarshals polymorphic NotificationConfiguration
func UnmarshalNotificationConfiguration(reader io.Reader, consumer runtime.Consumer) (NotificationConfiguration, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalNotificationConfiguration(data, consumer)
}

func unmarshalNotificationConfiguration(data []byte, consumer runtime.Consumer) (NotificationConfiguration, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Represents collection of different Event configurations. property.
	var getType struct {
		RepresentsCollectionOfDifferentEventConfigurations string `json:"Represents collection of different Event configurations."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Represents collection of different Event configurations.", "body", getType.RepresentsCollectionOfDifferentEventConfigurations); err != nil {
		return nil, err
	}

	// The value of Represents collection of different Event configurations. is used to determine which type to create and unmarshal the data into
	switch getType.RepresentsCollectionOfDifferentEventConfigurations {
	case "NotificationConfiguration":
		var result notificationConfiguration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Represents collection of different Event configurations. value: %q", getType.RepresentsCollectionOfDifferentEventConfigurations)
}

// Validate validates this notification configuration
func (m *notificationConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJira(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhook(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *notificationConfiguration) validateEmail(formats strfmt.Registry) error {
	if swag.IsZero(m.Email()) { // not required
		return nil
	}

	for i := 0; i < len(m.Email()); i++ {

		if err := m.emailField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("email" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *notificationConfiguration) validateJira(formats strfmt.Registry) error {
	if swag.IsZero(m.Jira()) { // not required
		return nil
	}

	for i := 0; i < len(m.Jira()); i++ {

		if err := m.jiraField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jira" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *notificationConfiguration) validateWebhook(formats strfmt.Registry) error {
	if swag.IsZero(m.Webhook()) { // not required
		return nil
	}

	for i := 0; i < len(m.Webhook()); i++ {

		if err := m.webhookField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this notification configuration based on the context it is used
func (m *notificationConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJira(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWebhook(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *notificationConfiguration) contextValidateEmail(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Email()); i++ {

		if err := m.emailField[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("email" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *notificationConfiguration) contextValidateJira(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Jira()); i++ {

		if err := m.jiraField[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jira" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *notificationConfiguration) contextValidateWebhook(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Webhook()); i++ {

		if err := m.webhookField[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhook" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}
