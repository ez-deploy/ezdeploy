package models

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceInfo Service Object
//
// swagger:model ServiceInfo
type ServiceInfo struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty" db:"id,type=INTEGER,primary,auto_increment"`

	// project id
	// Example: 1
	ProjectID int64 `json:"project_id,omitempty" db:"project_id,type=INTEGER"`

	// service name
	// Example: foobar
	Name string `json:"name,omitempty" db:"name,type=VARCHAR(255)"`

	// service description
	// Example: foobar service
	Description string `json:"description,omitempty" db:"description,type=VARCHAR(255)"`

	// service is running or not
	// Example: false
	Running bool `json:"running,omitempty" db:"running,type=BOOLEAN"`

	// version id
	// Example: 1
	VersionID int64 `json:"version_id,omitempty" db:"version_id,type=INTEGER"`

	// number of pod replicas
	// Example: 16
	Replica int64 `json:"replica,omitempty" db:"replica,type=INTEGER"`

	// expose service or not
	// Example: service
	// Enum: [incluster nodeport none]
	ExposeType string `json:"expose_type,omitempty" db:"expose_type,type=VARCHAR(255)"`

	// in-cluster expose port
	// Example: 80
	InClusterPort int64 `json:"in_cluster_port,omitempty" db:"in_cluster_port,type=INTEGER"`

	// node-port expose port
	// Example: 80
	NodePort int64 `json:"node_port,omitempty" db:"node_port,type=INTEGER"`

	// create at, unix timestamp
	// Example: 1528894200
	CreateAt int64 `json:"create_at,omitempty" db:"create_at,type=INTEGER"`

	// update at, unix timestamp
	// Example: 1528894200
	UpdateAt int64 `json:"update_at,omitempty" db:"update_at,type=INTEGER"`
}

// Validate validates this service info
func (m *ServiceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExposeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serviceInfoTypeExposeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["incluster","nodeport","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceInfoTypeExposeTypePropEnum = append(serviceInfoTypeExposeTypePropEnum, v)
	}
}

const (

	// ServiceInfoExposeTypeIncluster captures enum value "incluster"
	ServiceInfoExposeTypeIncluster string = "incluster"

	// ServiceInfoExposeTypeNodeport captures enum value "nodeport"
	ServiceInfoExposeTypeNodeport string = "nodeport"

	// ServiceInfoExposeTypeNone captures enum value "none"
	ServiceInfoExposeTypeNone string = "none"
)

// prop value enum
func (m *ServiceInfo) validateExposeTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, serviceInfoTypeExposeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ServiceInfo) validateExposeType(formats strfmt.Registry) error {
	if swag.IsZero(m.ExposeType) { // not required
		return nil
	}

	// value enum
	fmt.Println("get valid type")
	if err := m.validateExposeTypeEnum("expose_type", "body", m.ExposeType); err != nil {
		return err
	}

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
