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

// AuthInfo auth info
//
// swagger:model AuthInfo
type AuthInfo struct {

	// token
	Token *Token `json:"token,omitempty"`

	// user's info
	UserInfo *UserInfo `json:"user_info,omitempty"`
}

// Validate validates this auth info
func (m *AuthInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthInfo) validateToken(formats strfmt.Registry) error {
	if swag.IsZero(m.Token) { // not required
		return nil
	}

	if m.Token != nil {
		if err := m.Token.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

func (m *AuthInfo) validateUserInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.UserInfo) { // not required
		return nil
	}

	if m.UserInfo != nil {
		if err := m.UserInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user_info")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this auth info based on the context it is used
func (m *AuthInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthInfo) contextValidateToken(ctx context.Context, formats strfmt.Registry) error {

	if m.Token != nil {
		if err := m.Token.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

func (m *AuthInfo) contextValidateUserInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.UserInfo != nil {
		if err := m.UserInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthInfo) UnmarshalBinary(b []byte) error {
	var res AuthInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}