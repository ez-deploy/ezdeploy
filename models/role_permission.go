package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RolePermission Role Permissions Object
//
// swagger:model RolePermission
type RolePermission struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty" db:"id,type=INTEGER,primary,auto_increment"`

	// permission
	// Enum: [read write delete deploy]
	Permission string `json:"permission,omitempty" db:"permission,type=VARCHAR(255)"`

	// AuthorityRole id
	// Example: 1
	RoleID int64 `json:"role_id,omitempty" db:"role_id,type=INTEGER"`
}

// Validate validates this role permission
func (m *RolePermission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rolePermissionTypePermissionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["read","write","delete","deploy"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rolePermissionTypePermissionPropEnum = append(rolePermissionTypePermissionPropEnum, v)
	}
}

const (

	// RolePermissionPermissionRead captures enum value "read"
	RolePermissionPermissionRead string = "read"

	// RolePermissionPermissionWrite captures enum value "write"
	RolePermissionPermissionWrite string = "write"

	// RolePermissionPermissionDelete captures enum value "delete"
	RolePermissionPermissionDelete string = "delete"

	// RolePermissionPermissionDeploy captures enum value "deploy"
	RolePermissionPermissionDeploy string = "deploy"
)

// prop value enum
func (m *RolePermission) validatePermissionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rolePermissionTypePermissionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RolePermission) validatePermission(formats strfmt.Registry) error {
	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	// value enum
	if err := m.validatePermissionEnum("permission", "body", m.Permission); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this role permission based on context it is used
func (m *RolePermission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RolePermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RolePermission) UnmarshalBinary(b []byte) error {
	var res RolePermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
