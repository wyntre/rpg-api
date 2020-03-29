package actions

import (
	"net/http"

	"github.com/wyntre/rpg_api/models"
)

func (as *ActionSuite) Test_HomeHandler() {
	res := as.JSON("/").Get()
	as.Equal(http.StatusUnauthorized, res.Code)
}

func (as *ActionSuite) Test_HomeHandler_Authorized() {
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

	// auth'd request
	req := as.JSON("/")
	req.Headers["Authorization"] = atr.Token
	res = req.Get()
	as.Equal(http.StatusOK, res.Code)
}
