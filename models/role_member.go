package models

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RoleMember Authority Role Member Object
//
// swagger:model RoleMember
type RoleMember struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty" db:"id,type=INTEGER,primary,auto_increment"`

	// role id
	// Example: 1
	RoleID int64 `json:"role_id,omitempty" db:"role_id,type=INTEGER"`

	// user id
	// Example: 1
	UserID int64 `json:"user_id,omitempty" db:"user_id,type=INTEGER"`
}

// Validate validates this role member
func (m *RoleMember) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this role member based on context it is used
func (m *RoleMember) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RoleMember) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleMember) UnmarshalBinary(b []byte) error {
	var res RoleMember
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
