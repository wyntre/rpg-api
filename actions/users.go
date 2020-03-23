package actions

import (
  "net/http"
  "strings"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/wyntre/rpg_api/models"
)

// UsersCreate registers a new user with the application.
// HTTP Method: POST
// Expected Data:
//   JSON Object:
//    {
//      "email": string,
//      "password": string
//    }
//
// Return:
//  401 with verrs
//  201 with JWT token
func UsersCreate(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := u.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
    c.Set("errors", verrs)
		return c.Error(http.StatusConflict, errors.New(verrs.Error()))
	}

	return c.Render(http.StatusCreated, r.JSON(map[string]string{
    "token": "token",
    }))
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
    // get Authorization header, strip out Bearer and check if empty
		if token := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", 1); token == "" {
			return c.Render(http.StatusUnauthorized, r.JSON(map[string]string{
          "message": "unauthoried access",
        }))
    // if not empty, see if the token has been revoked
    } else {
      t := &models.Revokedtoken{}
      tx := c.Value("tx").(*pop.Connection)

      if err := tx.Where("token = ?", token).First(t); err != nil {
        return c.Render(http.StatusUnauthorized, r.JSON(map[string]string{
            "message": "unauthoried access",
          }))
      }
    }
    // if not empty and not revoked, validate token
		return next(c)
	}
}
