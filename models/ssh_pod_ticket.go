package models

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SSHPodTicket SSH Pod Ticket Object
//
// swagger:model SSHPodTicket
type SSHPodTicket struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty" db:"id,type=INTEGER,primary,auto_increment"`

	// namespace name
	// Example: foobar
	NamespaceName string `json:"namespace_name,omitempty" db:"namespace_name,type=VARCHAR(255)"`

	// pod name
	// Example: foobar-9zqb2
	PodName string `json:"pod_name,omitempty" db:"pod_name,type=VARCHAR(255)"`

	// ticket
	// Example: AISBJFCOIZXUF==
	Ticket string `json:"ticket,omitempty" db:"ticket,type=VARCHAR(255)"`

	// user id
	// Example: 1
	UserID int64 `json:"user_id,omitempty" db:"user_id,type=INTEGER"`

	// create at, unix timestamp
	CreateAt int64 `json:"create_at,omitempty" db:"create_at,type=INTEGER"`
}

// Validate validates this SSH pod ticket
func (m *SSHPodTicket) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this SSH pod ticket based on context it is used
func (m *SSHPodTicket) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SSHPodTicket) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SSHPodTicket) UnmarshalBinary(b []byte) error {
	var res SSHPodTicket
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
