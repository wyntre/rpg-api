package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
	"time"
)

// Tile is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Tile struct {
	ID         uuid.UUID `json:"id" db:"id"`
	UserID     uuid.UUID `json:"-" db:"user_id"`
	LevelID    uuid.UUID `json:"level_id" db:"level_id"`
	Level      *Level    `json:"level,omitempty" belongs_to:"level"`
	X          int       `json:"x" db:"x"`
	Y          int       `json:"y" db:"y"`
	TileTypeID uuid.UUID `json:"-" db:"tile_type_id"`
	TileType   *TileType `json:"tile_type" belongs_to:"tile_type"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (t Tile) String() string {
	jc, _ := json.Marshal(t)
	return string(jc)
}

// Tiles is not required by pop and may be deleted
type Tiles []Tile

// String is not required by pop and may be deleted
func (t Tiles) String() string {
	jc, _ := json.Marshal(t)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Tile) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.UUIDIsPresent{Field: t.UserID, Name: "UserID"},
		&validators.UUIDIsPresent{Field: t.LevelID, Name: "LevelID"},
		&validators.IntIsPresent{Field: t.X, Name: "X"},
		&validators.IntIsPresent{Field: t.Y, Name: "Y"},
		&validators.UUIDIsPresent{Field: t.TileTypeID, Name: "TileTypeID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Tile) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Tile) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
