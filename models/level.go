package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// Level is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Level struct {
    ID 					uuid.UUID  `json:"id" db:"id"`
		UserID 			uuid.UUID  `json:"-" db:"user_id"`
    Name 				string 		 `json:"name" db:"name"`
    Description string 		 `json:"description" db:"description"`
    MapID       uuid.UUID  `json:"map_id" db:"map_id"`
    Map         *Map       `json:"map,omitempty" belongs_to:"map"`
    CreatedAt 	time.Time  `json:"created_at" db:"created_at"`
    UpdatedAt 	time.Time  `json:"updated_at" db:"updated_at"`
    SortOrder   int        `json:"sort_order" db:"sort_order"`
}

// String is not required by pop and may be deleted
func (l Level) String() string {
	jc, _ := json.Marshal(l)
	return string(jc)
}

// Levels is not required by pop and may be deleted
type Levels []Level

// String is not required by pop and may be deleted
func (l Levels) String() string {
	jc, _ := json.Marshal(l)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (l *Level) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: l.UserID, Name: "UserID"},
		&validators.StringIsPresent{Field: l.Name, Name: "Name"},
		&validators.StringIsPresent{Field: l.Description, Name: "Description"},
    &validators.UUIDIsPresent{Field: l.MapID, Name: "MapID"},
    &validators.IntIsPresent{Field: l.SortOrder, Name: "SortOrder"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (l *Level) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (l *Level) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
