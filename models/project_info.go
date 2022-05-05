package models

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectInfo Project Object
//
// swagger:model ProjectInfo
type ProjectInfo struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty" db:"id,type=INTEGER,primary,auto_increment"`

	// owner_id
	// Example: 1
	OwnerID int64 `json:"owner_id,omitempty" db:"owner_id,type=INTEGER"`

	// project name
	// Example: foobar
	Name string `json:"name,omitempty" db:"name,type=VARCHAR(255)"`

	// project description
	// Example: foobar project
	Description string `json:"description,omitempty" db:"description,type=VARCHAR(255)"`

	// create at, unix timestamp
	// Example: 1528894200
	CreateAt int64 `json:"create_at,omitempty" db:"create_at,type=INTEGER"`

	// update at, unix timestamp
	// Example: 1528894200
	UpdateAt int64 `json:"update_at,omitempty" db:"update_at,type=INTEGER"`
}

// Validate validates this project info
func (m *ProjectInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project info based on context it is used
func (m *ProjectInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectInfo) UnmarshalBinary(b []byte) error {
	var res ProjectInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
