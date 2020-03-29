package actions

import (
	"net/http"

	"github.com/wyntre/rpg_api/models"
)

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
