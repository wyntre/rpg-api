package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/wyntre/rpg_api/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (TileType)
// DB Table: Plural (tileTypes)
// Resource: Plural (TileTypes)
// Path: Plural (/tileTypes)
// View Template Folder: Plural (/templates/tileTypes/)

// TileTypesResource is the resource for the TileType model
type TileTypesResource struct {
	buffalo.Resource
}

// List gets the tile types for a tile category
func (v TileTypesResource) List(c buffalo.Context) error {
	tileCategoryID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("bad tile category id"))
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	tileCategory := &models.TileCategory{}

	if err := tx.Eager().Find(tileCategory, tileCategoryID); err != nil {
		return c.Error(http.StatusNotFound, errors.New("tile category not found"))
	}

	return c.Render(http.StatusOK, r.JSON(map[string]models.TileTypes{
		"tile_types": tileCategory.TileTypes,
	}))
}

// Show gets the data for one TileType. This function is tileTypeped to
// the path GET /tileTypes/{tileTypeID}
func (v TileTypesResource) Show(c buffalo.Context) error {
	tileTypeID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("bad tile type id"))
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Allocate an empty TileType
	tileType := &models.TileType{}

	// To find the TileType the parameter tileTypeID is used.
	if err := tx.Eager().Find(tileType, tileTypeID); err != nil {
		return c.Error(http.StatusNotFound, errors.New("tile type not found"))
	}

	return c.Render(http.StatusOK, r.JSON(tileType))
}

// Create adds a TileType to the DB. This function is tileTypeped to the
// path POST /tileTypes
func (v TileTypesResource) Create(c buffalo.Context) error {
	// Allocate an empty TileType
	tileType := &models.TileType{}
	// Bind tileType to the html form elements
	if err := c.Bind(tileType); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(tileType)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return c.Error(http.StatusUnprocessableEntity, verrs)
	}

	return c.Render(http.StatusCreated, r.JSON(tileType))
}

// Update changes a TileType in the DB. This function is tileTypeped to
// the path PUT /tileTypes/{tileTypeID}
func (v TileTypesResource) Update(c buffalo.Context) error {
	tileTypeID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("bad tile type id"))
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	tileType := &models.TileType{}

	if err := tx.Find(tileType, tileTypeID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind TileType to the html form elements
	if err := c.Bind(tileType); err != nil {
		return err
	}

	tileType.ID = tileTypeID

	verrs, err := tx.ValidateAndUpdate(tileType)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return c.Error(http.StatusUnprocessableEntity, verrs)
	}

	return c.Render(http.StatusAccepted, r.JSON(tileType))
}

// Destroy deletes a TileType from the DB. This function is tileTypeped
// to the path DELETE /tileTypes/{tileTypeID}
func (v TileTypesResource) Destroy(c buffalo.Context) error {
	tileTypeID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("bad tile type id"))
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Allocate an empty TileType
	tileType := &models.TileType{}

	// To find the TileType the parameter tileTypeID is used.
	if err := tx.Find(tileType, tileTypeID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(tileType); err != nil {
		return err
	}

	return c.Render(http.StatusAccepted, r.JSON(tileType))
}
