// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIVipConnectivityResponse api vip connectivity response
//
// swagger:model api_vip_connectivity_response
type APIVipConnectivityResponse struct {

	// Ignition fetched from the specified API VIP
	Ignition string `json:"ignition,omitempty"`

	// Ignition downloadability check result.
	IsSuccess bool `json:"is_success,omitempty"`
}

// Validate validates this api vip connectivity response
func (m *APIVipConnectivityResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this api vip connectivity response based on context it is used
func (m *APIVipConnectivityResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIVipConnectivityResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIVipConnectivityResponse) UnmarshalBinary(b []byte) error {
	var res APIVipConnectivityResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
