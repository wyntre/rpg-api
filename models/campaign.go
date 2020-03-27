package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// Campaign is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Campaign struct {
    ID 					uuid.UUID `json:"id" db:"id"`
		UserID 			uuid.UUID `json:"-" db:"user_id"`
		User				*User     `json:"user,omitempty" belongs_to:"user"`
    Name 				string 		`json:"name" db:"name"`
    Description string 		`json:"description" db:"description"`
    CreatedAt 	time.Time `json:"created_at" db:"created_at"`
    UpdatedAt 	time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (c Campaign) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Campaigns is not required by pop and may be deleted
type Campaigns []Campaign

// String is not required by pop and may be deleted
func (c Campaigns) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Campaign) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: c.UserID, Name: "UserID"},
		&validators.StringIsPresent{Field: c.Name, Name: "Name"},
		&validators.StringIsPresent{Field: c.Description, Name: "Description"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Campaign) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Campaign) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
