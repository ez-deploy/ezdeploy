package models

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EnvironmentVariable Environment Variable for service deploy
//
// swagger:model EnvironmentVariable
type EnvironmentVariable struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty" db:"id,type=INTEGER,primary,auto_increment"`

	// service id
	// Example: 1
	ServiceID int64 `json:"service_id,omitempty" db:"service_id,type=INTEGER"`

	// environment variable name
	// Example: foobar
	Name string `json:"name,omitempty" db:"name,type=VARCHAR(255)"`

	// environment variable value
	// Example: foobar
	Value string `json:"value,omitempty" db:"value,type=VARCHAR(255)"`
}

// Validate validates this environment variable
func (m *EnvironmentVariable) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this environment variable based on context it is used
func (m *EnvironmentVariable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentVariable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentVariable) UnmarshalBinary(b []byte) error {
	var res EnvironmentVariable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
