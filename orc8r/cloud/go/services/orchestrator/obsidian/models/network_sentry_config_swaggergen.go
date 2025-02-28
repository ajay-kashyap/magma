// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetworkSentryConfig Sentry.io configuration
// swagger:model network_sentry_config
type NetworkSentryConfig struct {

	// log message patterns that are excluded (regex substring match)
	// Min Items: 0
	ExclusionPatterns []string `json:"exclusion_patterns"`

	// sample rate
	// Maximum: 1
	// Minimum: 0
	SampleRate *float32 `json:"sample_rate,omitempty"`

	// upload mme log
	UploadMmeLog bool `json:"upload_mme_log,omitempty"`

	// url native
	// Min Length: 0
	// Format: uri
	URLNative strfmt.URI `json:"url_native,omitempty"`

	// url python
	// Min Length: 0
	// Format: uri
	URLPython strfmt.URI `json:"url_python,omitempty"`
}

func (m *NetworkSentryConfig) UnmarshalJSON(b []byte) error {
	type NetworkSentryConfigAlias NetworkSentryConfig
	var t NetworkSentryConfigAlias
	if err := json.Unmarshal([]byte("{\"exclusion_patterns\":[\"GetServiceInfo\",\"GetOperationalStates\",\"ConnectionError\",\"CheckinError\",\"Checkin Error\",\"GetChallenge error!\",\"Connection to FluentBit\",\"\\\\[SyncRPC\\\\]\",\"Metrics upload error\",\"Streaming from the cloud failed!\",\"Fetch subscribers error!\",\"GRPC call failed for state replication\"],\"sample_rate\":1,\"upload_mme_log\":false}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = NetworkSentryConfig(t)
	return nil
}

// Validate validates this network sentry config
func (m *NetworkSentryConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExclusionPatterns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSampleRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURLNative(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURLPython(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkSentryConfig) validateExclusionPatterns(formats strfmt.Registry) error {

	if swag.IsZero(m.ExclusionPatterns) { // not required
		return nil
	}

	iExclusionPatternsSize := int64(len(m.ExclusionPatterns))

	if err := validate.MinItems("exclusion_patterns", "body", iExclusionPatternsSize, 0); err != nil {
		return err
	}

	return nil
}

func (m *NetworkSentryConfig) validateSampleRate(formats strfmt.Registry) error {

	if swag.IsZero(m.SampleRate) { // not required
		return nil
	}

	if err := validate.Minimum("sample_rate", "body", float64(*m.SampleRate), 0, false); err != nil {
		return err
	}

	if err := validate.Maximum("sample_rate", "body", float64(*m.SampleRate), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *NetworkSentryConfig) validateURLNative(formats strfmt.Registry) error {

	if swag.IsZero(m.URLNative) { // not required
		return nil
	}

	if err := validate.MinLength("url_native", "body", string(m.URLNative), 0); err != nil {
		return err
	}

	if err := validate.FormatOf("url_native", "body", "uri", m.URLNative.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NetworkSentryConfig) validateURLPython(formats strfmt.Registry) error {

	if swag.IsZero(m.URLPython) { // not required
		return nil
	}

	if err := validate.MinLength("url_python", "body", string(m.URLPython), 0); err != nil {
		return err
	}

	if err := validate.FormatOf("url_python", "body", "uri", m.URLPython.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkSentryConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkSentryConfig) UnmarshalBinary(b []byte) error {
	var res NetworkSentryConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
