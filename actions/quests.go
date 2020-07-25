package actions

import (
	"database/sql"
	"net/http"

	"github.com/dgrijalva/jwt-go"
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
// Model: Singular (Quest)
// DB Table: Plural (quests)
// Resource: Plural (Quests)
// Path: Plural (/quests)
// View Template Folder: Plural (/templates/quests/)

// QuestsResource is the resource for the Quest model
type QuestsResource struct {
	buffalo.Resource
}

// List gets all Quests. This function is mapped to the path
// GET /quests/{campaign_id}
func (v QuestsResource) List(c buffalo.Context) error {
	claims := c.Value("claims").(jwt.MapClaims)
	userID, err := uuid.FromString(claims["id"].(string))
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("bad user id"))
	}

	campaignID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("bad campaign id"))
	}

	campaign := &models.Campaign{}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Retrieve all Quests from the DB
	if err := tx.Eager("Quests").Where("user_id = ?", userID).Find(campaign, campaignID); err != nil {
		return c.Error(http.StatusNotFound, errors.New("quests not found"))
	}

	return c.Render(http.StatusOK, r.JSON(map[string]models.Quests{
		"quests": campaign.Quests,
	}))
}

// Show gets the data for one Quest. This function is mapped to
// the path GET /quests/{quest_id}
func (v QuestsResource) Show(c buffalo.Context) error {
	claims := c.Value("claims").(jwt.MapClaims)
	userID, err := uuid.FromString(claims["id"].(string))
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("bad user id"))
	}

	questID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("bad quest id"))
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Allocate an empty Quest
	quest := &models.Quest{}

	// To find the Quest the parameter quest_id is used.
	if err := tx.Eager().Where("user_id = ?", userID).Find(quest, questID); err != nil {
		return c.Error(http.StatusNotFound, errors.New("quest not found"))
	}

	return c.Render(200, r.JSON(quest))
}

// Create adds a Quest to the DB. This function is mapped to the
// path POST /quests
func (v QuestsResource) Create(c buffalo.Context) error {
	claims := c.Value("claims").(jwt.MapClaims)
	userID, err := uuid.FromString(claims["id"].(string))
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("bad user id"))
	}

	// Allocate an empty Quest
	quest := &models.Quest{}

	// Bind quest to the html form elements
	if err := c.Bind(quest); err != nil {
		return err
	}

	quest.UserID = userID

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Get Quest of Highest Sort Order
	lastQuest := &models.Quest{}
	err = tx.Where("user_id = ?", userID).Where("campaign_id = ?", quest.CampaignID).Order("sort_order desc").First(lastQuest)
	if err != nil {
		if errors.Cause(err) != sql.ErrNoRows {
			c.Logger().Error(err)
			return errors.New("transaction error")
		}
	}

	// lastQuest.SortOrder defaults to 0 if no quests found
	quest.SortOrder = lastQuest.SortOrder + 1

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(quest)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return c.Error(http.StatusUnprocessableEntity, verrs)
	}

	return c.Render(http.StatusCreated, r.JSON(quest))
}

// Update changes a Quest in the DB. This function is mapped to
// the path PUT /quests/{quest_id}
func (v QuestsResource) Update(c buffalo.Context) error {
	claims := c.Value("claims").(jwt.MapClaims)
	userID, err := uuid.FromString(claims["id"].(string))
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("bad user id"))
	}

	questID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("bad quest id"))
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Allocate an empty Quest
	quest := &models.Quest{}

	quest.UserID = userID

	if err := tx.Where("user_id = ?", userID).Find(quest, questID); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Quest to the html form elements
	if err := c.Bind(quest); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(quest)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return c.Error(http.StatusUnprocessableEntity, verrs)
	}

	return c.Render(http.StatusOK, r.JSON(quest))
}

// Destroy deletes a Quest from the DB. This function is mapped
// to the path DELETE /quests/{quest_id}
// This API should destroy all child elements (maps)
// This API should accept the destroy command, any UI should implement a user check ("Are you sure?")
func (v QuestsResource) Destroy(c buffalo.Context) error {
	claims := c.Value("claims").(jwt.MapClaims)
	userID, err := uuid.FromString(claims["id"].(string))
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("bad user id"))
	}

	questID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("bad quest id"))
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.New("no transaction found")
	}

	// Allocate an empty Quest
	quest := &models.Quest{}

	// To find the Quest the parameter quest_id is used.
	if err := tx.Where("user_id = ?", userID).Find(quest, questID); err != nil {
		return c.Error(http.StatusNotFound, errors.New("quest not found"))
	}

	// start to destroy maps
	maps := models.Maps{}
	if err := tx.Where("user_id = ?", userID).Where("quest_id = ?", quest.ID).All(&maps); err != nil {
		if errors.Cause(err) != sql.ErrNoRows {
			return c.Error(http.StatusInternalServerError, errors.New("cannot select maps"))
		}
	}

	for i := range maps {
		// start to destroy levels
		levels := models.Levels{}
		if err := tx.Where("user_id = ?", userID).Where("map_id = ?", maps[i].ID).All(&levels); err != nil {
			if errors.Cause(err) != sql.ErrNoRows {
				return c.Error(http.StatusInternalServerError, errors.New("cannot select levels"))
			}
		}
		if err := tx.Destroy(levels); err != nil {
			return c.Error(http.StatusInternalServerError, errors.New("could not destroy levels"))
		}
		// end to destroy levels
	}

	if err := tx.Destroy(maps); err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("could not destroy maps"))
	}
	// end to destroy maps

	// destroy quest
	if err := tx.Destroy(quest); err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("could not destroy quest"))
	}

	return c.Render(http.StatusOK, r.JSON(quest))
}
