package models

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceVersion Service Version Object
//
// swagger:model ServiceVersion
type ServiceVersion struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty" db:"id,type=INTEGER,primary,auto_increment"`

	// version name
	// Example: foobar
	Name string `json:"name,omitempty" db:"name,type=VARCHAR(255)"`

	// version description
	// Example: foobar version
	Description string `json:"description,omitempty" db:"description,type=VARCHAR(255)"`

	// image name
	// Example: busybox:latest
	Image string `json:"image,omitempty" db:"image,type=VARCHAR(255)"`

	// container port
	// Example: 80
	ContainerPort int64 `json:"container_port,omitempty" db:"container_port,type=INTEGER"`

	// create at, unix timestamp
	// Example: 1528894200
	CreateAt int64 `json:"create_at,omitempty" db:"create_at,type=INTEGER"`

	// service id
	// Example: 1
	ServiceID int64 `json:"service_id,omitempty" db:"service_id,type=INTEGER"`
}

// Validate validates this service version
func (m *ServiceVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service version based on context it is used
func (m *ServiceVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceVersion) UnmarshalBinary(b []byte) error {
	var res ServiceVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
