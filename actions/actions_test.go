package actions

import (
	"testing"

	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/suite"
)

type ActionSuite struct {
	*suite.Action
}

type AuthRequest struct {
	Email string `json:email`
	Password string `json:password`
}

type AuthTokenResponse struct {
	Token string `json:token`
}

type ErrorResponse struct {
	Error string `json:error`
	Trace string `json:trace`
	Code 	string `json:code`
}

func Test_ActionSuite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(App(), packr.New("Test_ActionSuite", "../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}
