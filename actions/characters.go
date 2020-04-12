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
//    Error:    409, 500, with errors
func CharactersCreate(c buffalo.Context) error {
  // grab user id from claims
  claims := c.Value("claims").(jwt.MapClaims)
  user_id, err := uuid.FromString(claims["id"].(string))
  if err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("bad user id"))
  }

  character := &models.Character{}

  if err := c.Bind(character); err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("bad character data"))
	}

  character.UserID = user_id

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := tx.ValidateAndCreate(character)
	if err != nil {
    c.Logger().Debug(err)
		return c.Error(http.StatusInternalServerError, errors.New("character not created"))
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
  // grab user id from claims
  claims := c.Value("claims").(jwt.MapClaims)
  user_id, err := uuid.FromString(claims["id"].(string))
  if err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("bad user id"))
  }

  // grab character id from url params
  character_id, err := uuid.FromString(c.Param("id"))
  if err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("bad character id"))
  }

  // find character
  character := &models.Character{}
  tx := c.Value("tx").(*pop.Connection)
  if err := tx.Where("user_id = ?", user_id).Find(character, character_id); err != nil {
    return c.Error(http.StatusNotFound, errors.New("character not found"))
  }

  // delete character
  if err := tx.Destroy(character); err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("character not deleted"))
  }

  // return response
  return c.Render(http.StatusOK, r.JSON(map[string]string{
    "message": "character deleted",
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
  user_id, err := uuid.FromString(claims["id"].(string))
  if err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("bad user id"))
  }

  user := &models.User{}

  tx := c.Value("tx").(*pop.Connection)
  if err := tx.Eager("Characters").Find(user, user_id); err != nil {
    return c.Error(http.StatusNotFound, errors.New("characters not found"))
  }

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
    return c.Error(http.StatusInternalServerError, errors.New("bad user id"))
  }

  character_id, err := uuid.FromString(c.Param("id"))
  if err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("bad character id"))
  }

  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return errors.New("no transaction found")
  }

  character := &models.Character{}
  if err := tx.Where("user_id = ?", user_id).Find(character, character_id); err != nil {
    return c.Error(http.StatusNotFound, errors.New("character not found"))
  }

  return c.Render(http.StatusOK, r.JSON(
    character,
  ))
}

// URI: /v1/characters/:id
// Method: POST
// Data:
//    Name: string
//    Description: string
// Return:
//    Success:  200, return Character JSON
//    Error:    409, with errors
func CharactersUpdate(c buffalo.Context) error {
  claims := c.Value("claims").(jwt.MapClaims)
  user_id, err := uuid.FromString(claims["id"].(string))
  if err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("bad user id"))
  }

  character_id, err := uuid.FromString(c.Param("id"))
  if err != nil {
    return c.Error(http.StatusInternalServerError, errors.New("bad character id"))
  }

  character := &models.Character{}

  // ensure there is a character for the user
  tx := c.Value("tx").(*pop.Connection)
  if err := tx.Where("user_id = ?", user_id).Find(character, character_id); err != nil {
    return c.Error(http.StatusNotFound, errors.New("character not found"))
  }

  if err := c.Bind(character); err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("incorrect character data"))
	}

  character.UserID = user_id

	verrs, err := tx.ValidateAndUpdate(character)
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("character not updated"))
	}

	if verrs.HasAny() {
    c.Set("errors", verrs)
		return c.Error(http.StatusConflict, errors.New(verrs.Error()))
	}

	return c.Render(http.StatusOK, r.JSON(
    character,
  ))
}
