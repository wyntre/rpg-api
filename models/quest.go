package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// Quest is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Quest struct {
    ID 					uuid.UUID  `json:"id" db:"id"`
		UserID 			uuid.UUID  `json:"-" db:"user_id"`
    Name 				string 		 `json:"name" db:"name"`
    Description string 		 `json:"description" db:"description"`
    CampaignID  uuid.UUID  `json:"campaign_id" db:"campaign_id"`
    Campaign    *Campaign  `json:"campaign,omitempty" belongs_to:"campaign"`
		Maps				Maps			 `json:"maps,omitempty" has_many:"maps"`
    CreatedAt 	time.Time  `json:"created_at" db:"created_at"`
    UpdatedAt 	time.Time  `json:"updated_at" db:"updated_at"`
    SortOrder   int        `json:"sort_order" db:"sort_order"`
}

// String is not required by pop and may be deleted
func (q Quest) String() string {
	jc, _ := json.Marshal(q)
	return string(jc)
}

// Quests is not required by pop and may be deleted
type Quests []Quest

// String is not required by pop and may be deleted
func (q Quests) String() string {
	jc, _ := json.Marshal(q)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (q *Quest) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: q.UserID, Name: "UserID"},
		&validators.StringIsPresent{Field: q.Name, Name: "Name"},
		&validators.StringIsPresent{Field: q.Description, Name: "Description"},
    &validators.UUIDIsPresent{Field: q.CampaignID, Name: "CampaignID"},
    &validators.IntIsPresent{Field: q.SortOrder, Name: "SortOrder"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (q *Quest) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (q *Quest) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
