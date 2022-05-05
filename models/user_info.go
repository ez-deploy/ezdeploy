package models

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserInfo User's Info
//
// swagger:model UserInfo
type UserInfo struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty" db:"id,type=INTEGER,primary,auto_increment"`

	// email, unique
	// Example: foo@bar.com
	Email string `json:"email,omitempty" db:"email,type=VARCHAR(255)"`

	// password
	// Example: foobar123
	Password string `json:"password,omitempty" db:"password,type=VARCHAR(255)"`

	// user's name, not unique
	// Example: foobar
	UserName string `json:"user_name,omitempty" db:"user_name,type=VARCHAR(255)"`
}

// Validate validates this user info
func (m *UserInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this user info based on context it is used
func (m *UserInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserInfo) UnmarshalBinary(b []byte) error {
	var res UserInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
