package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
	"time"
)

// Campaign is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Campaign struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	UserID      uuid.UUID  `json:"-" db:"user_id"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description" db:"description"`
	Characters  Characters `json:"characters,omitempty" has_many:"characters" order_by:"created_at asc"`
	Quests      Quests     `json:"quests,omitempty" has_many:"quests" order_by:"sort_order asc"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
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
