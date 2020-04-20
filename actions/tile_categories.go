package actions

import (
  "net/http"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop"
  "github.com/wyntre/rpg_api/models"
  "github.com/gofrs/uuid"
  "github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (TileCategory)
// DB Table: Plural (tile_categories)
// Resource: Plural (TileCategories)
// Path: Plural (/tile_categories)
// View Template Folder: Plural (/templates/tile_categorys/)

// TileCategorysResource is the resource for the TileCategory model
type TileCategoriesResource struct{
  buffalo.Resource
}

// List gets all TileCategorys. This function is tile_categoryped to the path
// GET /tile_categorys/{tile category id}
func (v TileCategoriesResource) List(c buffalo.Context) error {
  tile_categories := &models.TileCategories{}

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return errors.New("no transaction found")
  }

  // Retrieve all TileCategorys from the DB
  if err := tx.All(tile_categories); err != nil {
    return c.Error(http.StatusNotFound, errors.New("tile categories not found"))
  }

  return c.Render(http.StatusOK, r.JSON(tile_categories))
}

// Show gets the data for one TileCategory. This function is tile_categoryped to
// the path GET /tile_categorys/{tile_category_id}
func (v TileCategoriesResource) Show(c buffalo.Context) error {
  tile_category_id, err := uuid.FromString(c.Param("id"))
  if err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("bad tile category id"))
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return errors.New("no transaction found")
  }

  // Allocate an empty TileCategory
  tile_category := &models.TileCategory{}

  // To find the TileCategory the parameter tile_category_id is used.
  if err := tx.Eager().Find(tile_category, tile_category_id); err != nil {
    return c.Error(http.StatusNotFound, errors.New("tile category not found"))
  }

  return c.Render(http.StatusOK, r.JSON(tile_category))
}

// Create adds a TileCategory to the DB. This function is tile_categoryped to the
// path POST /tile_categorys
func (v TileCategoriesResource) Create(c buffalo.Context) error {
  // Allocate an empty TileCategory
  tile_category := &models.TileCategory{}
  // Bind tile_category to the html form elements
  if err := c.Bind(tile_category); err != nil {
    return err
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return errors.New("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(tile_category)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return c.Error(http.StatusUnprocessableEntity, verrs)
  }

  return c.Render(http.StatusCreated, r.JSON(tile_category))
}

// Update changes a TileCategory in the DB. This function is tile_categoryped to
// the path PUT /tile_categorys/{tile_category_id}
func (v TileCategoriesResource) Update(c buffalo.Context) error {
  tile_category_id, err := uuid.FromString(c.Param("id"))
  if err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("bad tile category id"))
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return errors.New("no transaction found")
  }

  tile_category := &models.TileCategory{}
  // Bind TileCategory to the html form elements
  if err := c.Bind(tile_category); err != nil {
    return err
  }

  tile_category.ID = tile_category_id

  verrs, err := tx.ValidateAndUpdate(tile_category)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
      return c.Error(http.StatusUnprocessableEntity, verrs)
  }

  return c.Render(http.StatusOK, r.JSON(tile_category))
}

// Destroy deletes a TileCategory from the DB. This function is tile_categoryped
// to the path DELETE /tile_categorys/{tile_category_id}
func (v TileCategoriesResource) Destroy(c buffalo.Context) error {
  tile_category_id, err := uuid.FromString(c.Param("id"))
  if err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("bad tile category id"))
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return errors.New("no transaction found")
  }

  // Allocate an empty TileCategory
  tile_category := &models.TileCategory{}

  // To find the TileCategory the parameter tile_category_id is used.
  if err := tx.Find(tile_category, tile_category_id); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  if err := tx.Destroy(tile_category); err != nil {
    return err
  }

  return c.Render(http.StatusOK, r.JSON(tile_category))
}