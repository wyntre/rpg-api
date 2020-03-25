package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// Revokedtoken is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Revokedtoken struct {
    ID 				uuid.UUID `json:"id" db:"id"`
    Token 		string 		`json:"token" db:"token"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (t *Revokedtoken) Create(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndCreate(t)
}

// String is not required by pop and may be deleted
func (r Revokedtoken) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Revokedtokens is not required by pop and may be deleted
type Revokedtokens []Revokedtoken

// String is not required by pop and may be deleted
func (r Revokedtokens) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *Revokedtoken) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(
		&validators.StringIsPresent{Field: r.Token, Name: "Token"},
		&validators.FuncValidator{
			Field:   r.Token,
			Name:    "Token",
			Message: "Token is already revoked",
			Fn: func() bool {
				var b bool
				q := tx.Where("token = ?", r.Token)
				if r.ID != uuid.Nil {
					q = q.Where("id != ?", r.ID)
				}
				b, err = q.Exists(r)
				if err != nil {
					return false
				}
				return !b
			},
		},
	), err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (r *Revokedtoken) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (r *Revokedtoken) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
