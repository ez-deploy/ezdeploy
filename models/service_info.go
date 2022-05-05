package models

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceInfo Service Object
//
// swagger:model ServiceInfo
type ServiceInfo struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty" db:"id,type=INTEGER,PRIMARY"`

	// project id
	// Example: 1
	ProjectID int64 `json:"project_id,omitempty" db:"project_id,type=INTEGER"`

	// service name
	// Example: foobar
	Name string `json:"name,omitempty" db:"name,type=VARCHAR(255)"`

	// service description
	// Example: foobar service
	Description string `json:"description,omitempty" db:"description,type=VARCHAR(255)"`

	// version id
	// Example: 1
	VersionID int64 `json:"version_id,omitempty" db:"version_id,type=INTEGER"`

	// number of pod replicas
	// Example: 1
	Replica int64 `json:"replica,omitempty" db:"replica,type=INTEGER"`

	// expose port
	// Example: 80
	ExposePort int64 `json:"expose_port,omitempty" db:"expose_port,type=INTEGER"`

	// create at, unix timestamp
	// Example: 1528894200
	CreateAt int64 `json:"create_at,omitempty" db:"create_at,type=INTEGER"`

	// update at, unix timestamp
	// Example: 1528894200
	UpdateAt int64 `json:"update_at,omitempty" db:"update_at,type=INTEGER"`
}

// Validate validates this service info
func (m *ServiceInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service info based on context it is used
func (m *ServiceInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceInfo) UnmarshalBinary(b []byte) error {
	var res ServiceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
