package actions

import (
  "fmt"
  "net/http"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop"
  "github.com/wyntre/rpg_api/models"
  "github.com/gofrs/uuid"
  "github.com/dgrijalva/jwt-go"
  "github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Campaign)
// DB Table: Plural (campaigns)
// Resource: Plural (Campaigns)
// Path: Plural (/campaigns)
// View Template Folder: Plural (/templates/campaigns/)

// CampaignsResource is the resource for the Campaign model
type CampaignsResource struct{
  buffalo.Resource
}

// List gets all Campaigns. This function is mapped to the path
// GET /campaigns
func (v CampaignsResource) List(c buffalo.Context) error {
  claims := c.Value("claims").(jwt.MapClaims)
  user_id, err := uuid.FromString(claims["id"].(string))
  if err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("bad user id"))
  }

  user := &models.User{}

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Retrieve all Campaigns from the DB
  if err := tx.Eager("Campaigns").Find(user, user_id); err != nil {
    return c.Error(http.StatusNotFound, errors.New("campaigns not found"))
  }

  return c.Render(http.StatusOK, r.JSON(map[string]models.Campaigns{
      "campaigns": user.Campaigns,
  }))
}

// Show gets the data for one Campaign. This function is mapped to
// the path GET /campaigns/{campaign_id}
func (v CampaignsResource) Show(c buffalo.Context) error {
  claims := c.Value("claims").(jwt.MapClaims)
  user_id, err := uuid.FromString(claims["id"].(string))
  if err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("bad user id"))
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Campaign
  campaign := &models.Campaign{}

  // To find the Campaign the parameter campaign_id is used.
  if err := tx.Where("user_id = ?", user_id).Find(campaign, c.Param("campaign_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  return c.Render(200, r.JSON(campaign))
}

// Create adds a Campaign to the DB. This function is mapped to the
// path POST /campaigns
func (v CampaignsResource) Create(c buffalo.Context) error {
  claims := c.Value("claims").(jwt.MapClaims)
  user_id, err := uuid.FromString(claims["id"].(string))
  if err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("bad user id"))
  }

  // Allocate an empty Campaign
  campaign := &models.Campaign{}

  // Bind campaign to the html form elements
  if err := c.Bind(campaign); err != nil {
    return err
  }

  campaign.UserID = user_id

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(campaign)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    return c.Error(http.StatusUnprocessableEntity, verrs)
  }

  return c.Render(http.StatusCreated, r.JSON(campaign))
}

// Update changes a Campaign in the DB. This function is mapped to
// the path PUT /campaigns/{campaign_id}
func (v CampaignsResource) Update(c buffalo.Context) error {
  claims := c.Value("claims").(jwt.MapClaims)
  user_id, err := uuid.FromString(claims["id"].(string))
  if err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("bad user id"))
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Campaign
  campaign := &models.Campaign{}

  campaign.UserID = user_id

  if err := tx.Where("user_id = ?", user_id).Find(campaign, c.Param("campaign_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  // Bind Campaign to the html form elements
  if err := c.Bind(campaign); err != nil {
    return err
  }

  verrs, err := tx.ValidateAndUpdate(campaign)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
      return c.Error(http.StatusUnprocessableEntity, verrs)
  }

  return c.Render(http.StatusOK, r.JSON(campaign))
}

// Destroy deletes a Campaign from the DB. This function is mapped
// to the path DELETE /campaigns/{campaign_id}
func (v CampaignsResource) Destroy(c buffalo.Context) error {
  claims := c.Value("claims").(jwt.MapClaims)
  user_id, err := uuid.FromString(claims["id"].(string))
  if err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("bad user id"))
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Campaign
  campaign := &models.Campaign{}

  // To find the Campaign the parameter campaign_id is used.
  if err := tx.Where("user_id = ?", user_id).Find(campaign, c.Param("campaign_id")); err != nil {
    return c.Error(http.StatusNotFound, err)
  }

  if err := tx.Destroy(campaign); err != nil {
    return err
  }

  return c.Render(http.StatusOK, r.JSON(campaign))
}
