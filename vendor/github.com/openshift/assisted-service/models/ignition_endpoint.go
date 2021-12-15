// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IgnitionEndpoint Explicit ignition endpoint overrides the default ignition endpoint.
//
// swagger:model ignition-endpoint
type IgnitionEndpoint struct {

	// A CA certficate to be used when contacting the URL via https.
	CaCertificate *string `json:"ca_certificate,omitempty"`

	// The URL for the ignition endpoint.
	URL *string `json:"url,omitempty"`
}

// Validate validates this ignition endpoint
func (m *IgnitionEndpoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ignition endpoint based on context it is used
func (m *IgnitionEndpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IgnitionEndpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgnitionEndpoint) UnmarshalBinary(b []byte) error {
	var res IgnitionEndpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
