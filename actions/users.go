package actions

import (
  "fmt"
  "net/http"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/wyntre/rpg_api/models"
  "github.com/gobuffalo/envy"
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

  c.Logger().Info(envy.Get("JWT_PUBLIC_KEY", ""))
	c.Logger().Info(envy.Get("JWT_PRIVATE_KEY", ""))
  data := *u
  data.Password = "*"
  c.Logger().Info(data)

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := u.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
    c.Set("errors", verrs)
		return c.Error(http.StatusConflict, errors.New(verrs.Error()))
	}

  signedToken, err := AuthGenerateToken(u)
  if err != nil {
    return err
  }

	return c.Render(http.StatusCreated, r.JSON(map[string]string{
    "token": fmt.Sprintf("Bearer %s", signedToken),
    }))
}
