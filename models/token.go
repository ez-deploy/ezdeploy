package models

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Token Token Object
//
// swagger:model Token
type Token struct {
	// id
	// Example: 1
	ID int64 `json:"id,omitempty" db:"id,type=INTEGER,PRIMARY"`

	// token type
	// Enum: [session private public]
	Type string `json:"type,omitempty" db:"type,type=VARCHAR(20)"`

	// user's id
	// Example: 1
	UserID int64 `json:"user_id,omitempty" db:"user_id,type=INTEGER"`

	// token value
	// Example: zxicgoiuasie
	Value string `json:"value,omitempty" db:"value,type=VARCHAR(255)"`

	// create at, unix timestamp
	// Example: 1528894200
	CreateAt int64 `json:"create_at,omitempty" db:"create_at,type=INTEGER"`

	// expired at, unix timestamp
	// Example: 1528994200
	ExpiredAt int64 `json:"expired_at,omitempty" db:"expired_at,type=INTEGER"`

	// create at, unix timestamp
	// Example: 1528894200
	UpdateAt int64 `json:"update_at,omitempty" db:"update_at,type=INTEGER"`
}

// Validate validates this token
func (m *Token) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tokenTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["session","private","public"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tokenTypeTypePropEnum = append(tokenTypeTypePropEnum, v)
	}
}

const (

	// TokenTypeSession captures enum value "session"
	TokenTypeSession string = "session"

	// TokenTypePrivate captures enum value "private"
	TokenTypePrivate string = "private"

	// TokenTypePublic captures enum value "public"
	TokenTypePublic string = "public"
)

// prop value enum
func (m *Token) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tokenTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Token) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this token based on context it is used
func (m *Token) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Token) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Token) UnmarshalBinary(b []byte) error {
	var res Token
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
