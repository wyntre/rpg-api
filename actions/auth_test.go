package actions

import (
	"net/http"

	"github.com/wyntre/rpg_api/models"
	"github.com/gobuffalo/envy"
)

// test valid auth token generated
func (as *ActionSuite) Test_Auth_Create() {
	// create valid user
	user := &models.User{
		Email: 								"test@test.com",
		Password: 						"test",
		PasswordConfirmation: "test",
	}
	verrs, err := user.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	// correct login
	auth := &AuthRequest{
		Email: "test@test.com",
		Password: "test",
	}

	res := as.JSON("/v1/auth").Post(auth)
	as.Equal(http.StatusAccepted, res.Code)
	atr := &AuthTokenResponse{}
	res.Bind(atr)
	as.NotNil(atr.Token)
}

// test invalid auth attempt
func (as *ActionSuite) Test_Auth_Create_Fail() {
	// create valid user
	user := &models.User{
		Email: 								"test@test.com",
		Password: 						"test",
		PasswordConfirmation: "test",
	}
	verrs, err := user.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	// correct login
	auth := &AuthRequest{
		Email: "test@test.com",
		Password: "fake",
	}

	res := as.JSON("/v1/auth").Post(auth)
	as.Equal(http.StatusUnauthorized, res.Code)
	er := &ErrorResponse{}
	res.Bind(er)
	as.NotNil(er.Error)
	as.Equal("invalid email/password", er.Error)
}

// test invalid auth token request
func (as *ActionSuite) Test_Auth_Verification_Fail() {
	// create a valid user
	token := as.CreateUser("test@test.com", "test")

	// set expiration time so invalid
	envy.Set("JWT_TOKEN_EXPIRATION", "-2h")
	user := &models.User{}
	err := as.DB.Where("email = ?", "test@test.com").First(user)
	as.NoError(err)
	// reset expiration token

	// generate expired token
	token, err = AuthGenerateToken(user)
	as.NoError(err)

	// test expired token
	req := as.JSON("/")
	req.Headers["Authorization"] = token
	res := req.Get()
	as.Equal(http.StatusUnauthorized, res.Code)

	envy.Set("JWT_TOKEN_EXPIRATION", "2h")
}

// test valid auth token revokation
func (as *ActionSuite) Test_Auth_Revokation() {
	// create valid user
	token := as.CreateUser("test@test.com", "test")

	// auth'd request
	req := as.JSON("/")
	req.Headers["Authorization"] = token
	res := req.Get()
	as.Equal(http.StatusOK, res.Code)

	// revoke token
	req = as.JSON("/v1/auth")
	req.Headers["Authorization"] = token
	res = req.Delete()
	as.Equal(http.StatusAccepted, res.Code)

	// revoke token again
	req = as.JSON("/v1/auth")
	req.Headers["Authorization"] = token
	res = req.Delete()
	as.Equal(http.StatusExpectationFailed, res.Code)
}
