package actions

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/pkg/errors"
	"github.com/wyntre/rpg_api/models"
	"golang.org/x/crypto/bcrypt"
)

// AuthCreate attempts to log the user in with an existing account.
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
//  202 with JWT token
func AuthCreate(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)

	// find a user with the email
	err := tx.Where("email = ?", strings.ToLower(strings.TrimSpace(u.Email))).First(u)

	// helper function to handle bad attempts
	bad := func() error {
		verrs := validate.NewErrors()
		verrs.Add("email", "invalid email/password")

		return c.Error(http.StatusUnauthorized, errors.New(verrs.Error()))
	}

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// couldn't find an user with the supplied email address.
			return bad()
		}
		return errors.WithStack(err)
	}

	// confirm that the given password matches the hashed password from the db
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	if err != nil {
		return bad()
	}

	return c.Render(http.StatusAccepted, r.JSON(map[string]string{
    "token": "token",
    }))
}

// AuthDestroy clears the session and logs a user out
// HTTP Method: DELETE
// Expected Data:
//  Header: Authorization
//
// Return:
//  417 with verrs
//  202
func AuthDestroy(c buffalo.Context) error {
	token := &models.Revokedtoken{}
	token.Token = c.Request().Header.Get("Authorization")

  tx := c.Value("tx").(*pop.Connection)
  verrs, err := token.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

  if verrs.HasAny() {
		return c.Error(http.StatusExpectationFailed, errors.New(verrs.Error()))
	}

  return c.Render(http.StatusAccepted, r.JSON(map[string]string{
    "message": "token invalidated",
    }))
}
