package models

import (
	"encoding/json"
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
	"time"
)

// Character is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Character struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	UserID      uuid.UUID  `json:"-" db:"user_id"`
	User        *User      `json:"-" belongs_to:"user"`
	Name        string     `json:"name" db:"name"`
	Description string     `json:"description" db:"description"`
	CampaignID  nulls.UUID `json:"-" db:"campaign_id"`
	Campaign    *Campaign  `json:"campaign,omitempty" belongs_to:"campaign"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
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
		&validators.UUIDIsPresent{Field: c.UserID, Name: "UserID"},
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
