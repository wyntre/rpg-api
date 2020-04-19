package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
	"time"
)

// Map is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Map struct {
	ID          uuid.UUID `json:"id" db:"id"`
	UserID      uuid.UUID `json:"-" db:"user_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	QuestID     uuid.UUID `json:"quest_id" db:"quest_id"`
	Quest       *Quest    `json:"quest,omitempty" belongs_to:"quest"`
	Levels      Levels    `json:"levels,omitempty" has_many:"levels" order_by:"sort_order asc"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	SortOrder   int       `json:"sort_order" db:"sort_order"`
}

// String is not required by pop and may be deleted
func (m Map) String() string {
	jc, _ := json.Marshal(m)
	return string(jc)
}

// Maps is not required by pop and may be deleted
type Maps []Map

// String is not required by pop and may be deleted
func (m Maps) String() string {
	jc, _ := json.Marshal(m)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (m *Map) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: m.UserID, Name: "UserID"},
		&validators.StringIsPresent{Field: m.Name, Name: "Name"},
		&validators.StringIsPresent{Field: m.Description, Name: "Description"},
		&validators.UUIDIsPresent{Field: m.QuestID, Name: "QuestID"},
		&validators.IntIsPresent{Field: m.SortOrder, Name: "SortOrder"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (m *Map) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (m *Map) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
