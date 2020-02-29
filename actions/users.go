package actions

import (
  "net/http"
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
		return c.Render(http.StatusExpectationFailed, r.JSON(verrs))
	}

	return c.Render(http.StatusCreated, r.JSON(map[string]string{
    "token": "token"
    }))
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Session().Set("redirectURL", c.Request().URL.String())

			err := c.Session().Save()
			if err != nil {
				return errors.WithStack(err)
			}

			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(302, "/auth/new")
		}
		return next(c)
	}
}
