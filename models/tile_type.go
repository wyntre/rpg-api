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
type TileType struct {
	ID             uuid.UUID     `json:"id" db:"id"`
	Name           string        `json:"name" db:"name"`
	TileCategoryID uuid.UUID     `json:"-" db:"tile_category_id"`
	TileCategory   *TileCategory `json:"tile_category" belongs_to:"tile_category"`
	Tiles          Tiles         `json:"tiles,omitempty" has_many:"tiles"`
	CreatedAt      time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (t TileType) String() string {
	jc, _ := json.Marshal(t)
	return string(jc)
}

// Tiles is not required by pop and may be deleted
type TileTypes []TileType

// String is not required by pop and may be deleted
func (t TileTypes) String() string {
	jc, _ := json.Marshal(t)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *TileType) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Name, Name: "Name"},
		&validators.UUIDIsPresent{Field: t.TileCategoryID, Name: "TileCategoryID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *TileType) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *TileType) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
