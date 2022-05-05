package models

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthorityRole Authority Role Object
//
// swagger:model AuthorityRole
type AuthorityRole struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty" db:"id,type=INTEGER,primary,auto_increment"`

	// user id
	// Example: 1
	UserID int64 `json:"user_id,omitempty" db:"user_id,type=INTEGER"`

	// project id
	// Example: 1
	ProjectID int64 `json:"project_id,omitempty" db:"project_id,type=INTEGER"`

	// role
	// Example: admin
	Role string `json:"role,omitempty" db:"role,type=VARCHAR(255)"`
}

// Validate validates this authority role
func (m *AuthorityRole) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this authority role based on context it is used
func (m *AuthorityRole) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthorityRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorityRole) UnmarshalBinary(b []byte) error {
	var res AuthorityRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
