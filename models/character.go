package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// Character is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Character struct {
    ID 					uuid.UUID `json:"id" db:"id"`
		UserID			uuid.UUID `json:"-" db:"user_id"`
		User				*User     `json:"user,omitempty" belongs_to:"user"`
    Name 				string 		`json:"name" db:"name"`
    Description string 		`json:"description" db:"description"`
    CreatedAt 	time.Time `json:"created_at" db:"created_at"`
    UpdatedAt 	time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (c Character) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Characters is not required by pop and may be deleted
type Characters []Character

// String is not required by pop and may be deleted
func (c Characters) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Character) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Name, Name: "Name"},
		&validators.StringIsPresent{Field: c.Description, Name: "Description"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Character) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Character) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
