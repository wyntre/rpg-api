package actions

import (
	"net/http"
	"testing"

	"github.com/gobuffalo/suite/v3"
	"github.com/wyntre/rpg_api/models"
)

type ActionSuite struct {
	*suite.Action
}

type AuthRequest struct {
	Email    string `json:email`
	Password string `json:password`
}

type AuthTokenResponse struct {
	Token string `json:token`
}

type ErrorResponse struct {
	Error string `json:error`
	Trace string `json:trace`
	Code  string `json:code`
}

func Test_ActionSuite(t *testing.T) {
	action := suite.NewAction(App())

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}

func (as *ActionSuite) ObtainToken(email string, password string) string {
	res := as.JSON("/v1/auth/").Post(&AuthRequest{
		Email:    email,
		Password: password,
	})

	as.Equal(http.StatusAccepted, res.Code)
	atr := &AuthTokenResponse{}
	res.Bind(atr)
	as.NotNil(atr.Token)

	return atr.Token
}

func (as *ActionSuite) CreateUser(email string, password string) string {
	user := &models.User{
		Email:                email,
		Password:             password,
		PasswordConfirmation: password,
	}
	verrs, err := user.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())

	return as.ObtainToken(email, password)
}
