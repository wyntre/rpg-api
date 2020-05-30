package actions

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
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
		return c.Error(http.StatusUnprocessableEntity, errors.New("incorrect auth fields"))
	}

	c.Logger().Debug(envy.Get("JWT_PUBLIC_KEY", ""))
	c.Logger().Debug(envy.Get("JWT_PRIVATE_KEY", ""))

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
		return c.Error(http.StatusInternalServerError, err)
	}

	// confirm that the given password matches the hashed password from the db
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	if err != nil {
		return bad()
	}

	signedToken, err := AuthGenerateToken(u)
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	return c.Render(http.StatusAccepted, r.JSON(map[string]string{
		"token": fmt.Sprintf("Bearer %s", signedToken),
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
	// get token and strip "Bearer " from string
	token := &models.Revokedtoken{}
	token.Token = strings.Split(
		c.Request().Header.Get("Authorization"),
		"Bearer ",
	)[1]

	// add revoked token to database
	tx := c.Value("tx").(*pop.Connection)
	verrs, err := tx.ValidateAndCreate(token)
	if err != nil {
		return c.Error(http.StatusInternalServerError, errors.New("could not invalidate token"))
	}

	if verrs.HasAny() {
		// return clean error message, remove extra data provided by validator
		return c.Error(http.StatusExpectationFailed, errors.New(strings.Split(
			verrs.Error(),
			"%",
		)[0]))
	}

	return c.Render(http.StatusAccepted, r.JSON(map[string]string{
		"message": "token invalidated",
	}))
}

func AuthGenerateToken(u *models.User) (string, error) {
	privateKey := envy.Get("JWT_PRIVATE_KEY", "keys/rsakey.pem")
	key, err := ioutil.ReadFile(privateKey)
	if err != nil {
		return "", errors.New("could not read key file")
	}

	parsedKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		return "", errors.New("error parsing key")
	}

	expireTime := envy.Get("JWT_TOKEN_EXPIRATION", "2h")
	expireToken, err := time.ParseDuration(expireTime)
	if err != nil {
		return "", errors.New("token expiration error")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"id":  u.ID,
		"exp": time.Now().Add(expireToken).Unix(),
	})
	signedToken, err := token.SignedString(parsedKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func checkToken(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		token := strings.Split(
			c.Request().Header.Get("Authorization"),
			"Bearer ",
		)[1]

		rtoken := &models.Revokedtoken{}

		tx := c.Value("tx").(*pop.Connection)
		err := tx.Where("token = ?", token).First(rtoken)
		// if no rows are returned, then token is not revoked
		if err != nil {
			if errors.Cause(err) == sql.ErrNoRows {
				return next(c)
			}
			return c.Error(http.StatusInternalServerError, err)
		}
		return c.Error(http.StatusUnauthorized, errors.New("token unauthorized"))
	}
}
