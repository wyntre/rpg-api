package actions

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/envy"
	"github.com/wyntre/rpg_api/models"
)

// test valid auth token generated
func (as *ActionSuite) Test_Auth_Create() {
	// create valid user
	user := &models.User{
		Email:                "test@test.com",
		Password:             "test",
		PasswordConfirmation: "test",
	}
	verrs, err := user.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	// correct login
	auth := &AuthRequest{
		Email:    "test@test.com",
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
		Email:                "test@test.com",
		Password:             "test",
		PasswordConfirmation: "test",
	}
	verrs, err := user.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	// incorrect password
	auth := &AuthRequest{
		Email:    "test@test.com",
		Password: "fake",
	}

	res := as.JSON("/v1/auth").Post(auth)
	as.Equal(http.StatusUnauthorized, res.Code)
	er := &ErrorResponse{}
	res.Bind(er)
	as.NotNil(er.Error)
	as.Equal("invalid email/password", er.Error)

	// incorrect email
	auth = &AuthRequest{
		Email:    "test@fake.com",
		Password: "test",
	}

	res = as.JSON("/v1/auth").Post(auth)
	as.Equal(http.StatusUnauthorized, res.Code)
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
	as.Equal(http.StatusUnauthorized, res.Code)
}

func (as *ActionSuite) Test_Auth_Revised_Token() {
	token := as.CreateUser("test@test.com", "test")

	// get claims from token
	base64_claims := strings.Split(token, ".")
	claims, err := jwt.DecodeSegment(base64_claims[1])
	as.NoError(err)

	// change claims
	revised_claims := strings.Replace(string(claims), "id", "ids", 1)
	base64_claims[1] = jwt.EncodeSegment([]byte(revised_claims))
	token = strings.Join(base64_claims, ".")

	// auth'd request
	req := as.JSON("/")
	req.Headers["Authorization"] = token
	res := req.Get()
	as.Equal(http.StatusUnauthorized, res.Code)

	// ensure verification error
	er := &ErrorResponse{}
	res.Bind(er)
	as.Equal("crypto/rsa: verification error", er.Error)
}

func (as *ActionSuite) Test_Auth_Revoked_Token() {
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

	// auth'd request
	req = as.JSON("/")
	req.Headers["Authorization"] = token
	res = req.Get()
	as.Equal(http.StatusUnauthorized, res.Code)
}
