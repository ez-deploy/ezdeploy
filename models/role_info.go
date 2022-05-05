package models

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RoleInfo Authority Role Object
//
// swagger:model RoleInfo
type RoleInfo struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty" db:"id,type=INTEGER,primary,auto_increment"`

	// project id
	// Example: 1
	ProjectID int64 `json:"project_id,omitempty" db:"project_id,type=INTEGER"`

	// role
	// Example: admin
	Role string `json:"role,omitempty" db:"role,type=VARCHAR(255)"`
}

// Validate validates this role info
func (m *RoleInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this role info based on context it is used
func (m *RoleInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RoleInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleInfo) UnmarshalBinary(b []byte) error {
	var res RoleInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
