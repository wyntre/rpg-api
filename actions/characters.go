package actions

import (
  "net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/wyntre/rpg_api/models"
  "github.com/dgrijalva/jwt-go"
  "github.com/gofrs/uuid"
)

// URI: /v1/characters/new
// Method: POST
// Data:
//    Name: string
//    Description: string
// Return:
//    Success:  201, return Character JSON
//    Error:    409, with errors
func CharactersCreate(c buffalo.Context) error {
  // grab user id from claims
  claims := c.Value("claims").(jwt.MapClaims)
  user_id, err := uuid.FromString(claims["id"].(string))
  if err != nil {
    return errors.WithStack(err)
  }

  character := &models.Character{}

  if err := c.Bind(character); err != nil {
		return errors.WithStack(err)
	}

  character.UserID = user_id

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := tx.ValidateAndCreate(character)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
    c.Set("errors", verrs)
		return c.Error(http.StatusConflict, errors.New(verrs.Error()))
	}

	return c.Render(http.StatusCreated, r.JSON(
    character,
  ))
}

func CharactersDestroy(c buffalo.Context) error {
  return c.Render(http.StatusNotFound, r.JSON(map[string]string{
    "message": "page not found",
    }))
}

// URI: /v1/characters/
// Method: GET
// Data:
//    None
// Return:
//    Success:  200, return List: Character JSON
//    Error:    409, with errors
func CharactersList(c buffalo.Context) error {
  claims := c.Value("claims").(jwt.MapClaims)
  c.Logger().Info(claims["id"])
  user_id, err := uuid.FromString(claims["id"].(string))
  if err != nil {
    return errors.WithStack(err)
  }

  user := &models.User{}

  tx := c.Value("tx").(*pop.Connection)
  if err := tx.Eager("Characters").Find(user, user_id); err != nil {
    return errors.WithStack(err)
  }

  c.Logger().Info(user)

  return c.Render(http.StatusOK, r.JSON(map[string]models.Characters{
    "characters": user.Characters,
    }))
}

// URI: /v1/characters/:id
// Method: GET
// Data:
//    id
// Return:
//    Success:  200, return Character JSON
//    Error:    409, with errors
func CharactersShow(c buffalo.Context) error {
  // grab user id from claims
  claims := c.Value("claims").(jwt.MapClaims)
  user_id, err := uuid.FromString(claims["id"].(string))
  if err != nil {
    return errors.WithStack(err)
  }

  character_id, err := uuid.FromString(c.Param("id"))

  character := &models.Character{}

  tx := c.Value("tx").(*pop.Connection)
  if err := tx.Where("user_id = ?", user_id).Find(character, character_id); err != nil {
    return errors.WithStack(err)
  }

  return c.Render(http.StatusOK, r.JSON(
    character,
  ))
}

func CharactersUpdate(c buffalo.Context) error {
  return c.Render(http.StatusNotFound, r.JSON(map[string]string{
    "message": "page not found",
    }))
}
